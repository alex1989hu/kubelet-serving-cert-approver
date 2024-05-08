// Copyright 2021 Alex Szakaly
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package certificatesigningrequest

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.uber.org/zap"
	authorizationv1 "k8s.io/api/authorization/v1"
	certificatesv1 "k8s.io/api/certificates/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	k8sclient "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	ctrlclient "sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/alex1989hu/kubelet-serving-cert-approver/metrics"
)

const eventWarningReason = "ApproveFailed"

var errSubjectAccessReview = errors.New("could not perform Subject Access Review")

// SigningReconciler reconciles a CertificateSigningRequest object.
type SigningReconciler struct {
	Client        ctrlclient.Client
	ClientSet     k8sclient.Interface
	Scheme        *runtime.Scheme
	EventRecorder record.EventRecorder
	Logger        *zap.Logger
}

// Reconcile processes request and returns the result.
//
//nolint:gocyclo
func (r *SigningReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	reqLogger := r.Logger.With(zap.String("csr.name", req.Name))

	var csr certificatesv1.CertificateSigningRequest

	if err := r.Client.Get(ctx, req.NamespacedName, &csr); ctrlclient.IgnoreNotFound(err) != nil {
		message := "Unable to to get Certificate Signing Requests"
		reqLogger.Error(message, zap.Error(err))

		return ctrl.Result{}, fmt.Errorf("%s %w", message, err)
	}

	switch {
	case csr.Spec.SignerName != certificatesv1.KubeletServingSignerName:
		if ce := reqLogger.Check(zap.DebugLevel,
			"Certificate Signing Request is not Kubelet serving Certificate"); ce != nil {
			ce.Write(
				zap.String("csr.signer", csr.Spec.SignerName),
			)
		}
	case !csr.DeletionTimestamp.IsZero():
		if ce := reqLogger.Check(zap.DebugLevel,
			"Certificate Signing Request has been deleted"); ce != nil {
			ce.Write(
				zap.Time("csr.deleted", csr.DeletionTimestamp.Time),
			)
		}
	case csr.Status.Certificate != nil:
		if ce := reqLogger.Check(zap.DebugLevel,
			"Certificate Signing Request is already signed"); ce != nil {
			ce.Write(
				zap.String("csr.signer", csr.Spec.SignerName),
			)
		}
	case len(csr.Status.Conditions) != 0:
		if ce := reqLogger.Check(zap.DebugLevel,
			"Certificate Signing Request already has approval condition"); ce != nil {
			ce.Write(
				zap.Any("csr.conditions", csr.Status.Conditions),
			)
		}
	default:
		x509cr, err := parseCSR(csr.Spec.Request)
		if err != nil {
			message := "Unable to parse Certificate Signing Request"

			reqLogger.Error(message, zap.Error(err))
			metrics.NumberOfInvalidCertificateSigningRequests.Inc()
			r.EventRecorder.Event(&csr, corev1.EventTypeWarning, eventWarningReason,
				message+": "+csr.Name+"): "+err.Error())

			return ctrl.Result{}, err
		}

		if errIsKubeletServingCert := isRequestConform(reqLogger, csr, x509cr); errIsKubeletServingCert != nil {
			message := "Unable to recognize the Certificate Signing Request"

			reqLogger.Error(message, zap.Error(err))
			metrics.NumberOfInvalidCertificateSigningRequests.Inc()
			r.EventRecorder.Event(&csr, corev1.EventTypeWarning, eventWarningReason,
				message+": "+csr.Name+"): "+errIsKubeletServingCert.Error())

			return ctrl.Result{}, fmt.Errorf(
				"the Certificate Signing Request does not conform with expectation: %w", errIsKubeletServingCert)
		}

		authorized, err := r.authorize(&csr)
		if err != nil {
			message := "Unable to get authorization of Certificate Signing Request"

			reqLogger.Error(message, zap.Error(err))
			metrics.NumberOfInvalidCertificateSigningRequests.Inc()
			r.EventRecorder.Event(&csr, corev1.EventTypeWarning, eventWarningReason,
				message+": "+csr.Name+err.Error())

			return ctrl.Result{}, fmt.Errorf("%s: %w", message, err)
		}

		if authorized {
			appendApprovalCondition(&csr)

			_, err = r.ClientSet.CertificatesV1().CertificateSigningRequests().UpdateApproval(
				context.TODO(), csr.Name, &csr, metav1.UpdateOptions{})
			if err != nil {
				message := "Unable to perform UpdateApproval"

				reqLogger.Error(message, zap.Error(err))
				metrics.NumberOfInvalidCertificateSigningRequests.Inc()
				r.EventRecorder.Event(&csr, corev1.EventTypeWarning, eventWarningReason,
					message+"("+csr.Name+"): "+err.Error())

				return ctrl.Result{}, fmt.Errorf("%s: %w", message, err)
			}

			reqLogger.Info("The Certificate Signing Request has been approved",
				zap.Strings("csr.request.dns", x509cr.DNSNames),
				zap.Any("csr.request.ip", x509cr.IPAddresses))

			metrics.NumberOfApprovedCertificateRequests.Inc()

			r.EventRecorder.Event(&csr, corev1.EventTypeNormal, "Approved",
				"The Certificate Signing Request has been approved: "+csr.Name)
		} else {
			message := "Node is not authorized. Unable to perform Subject Access Review, "

			reqLogger.Error(message)
			metrics.NumberOfInvalidCertificateSigningRequests.Inc()
			r.EventRecorder.Event(&csr, corev1.EventTypeWarning, eventWarningReason,
				message+": "+csr.Name)

			return ctrl.Result{}, errSubjectAccessReview
		}
	}

	return ctrl.Result{}, nil
}

// Validate that the given node has authorization to actually create CSRs.
func (r *SigningReconciler) authorize(csr *certificatesv1.CertificateSigningRequest) (bool, error) {
	log := r.Logger.With(zap.String("csr.name", csr.Name))

	extra := make(map[string]authorizationv1.ExtraValue, len(csr.Spec.Extra))

	for k, v := range csr.Spec.Extra {
		extra[k] = authorizationv1.ExtraValue(v)
	}

	sar := authorizationv1.SubjectAccessReview{
		Spec: authorizationv1.SubjectAccessReviewSpec{
			User:   csr.Spec.Username,
			UID:    csr.Spec.UID,
			Groups: csr.Spec.Groups,
			Extra:  extra,
			ResourceAttributes: &authorizationv1.ResourceAttributes{
				Group:    certificatesv1.GroupName,
				Resource: "certificatesigningrequests",
				Verb:     "create",
			},
		},
	}

	res, err := r.ClientSet.AuthorizationV1().SubjectAccessReviews().Create(context.TODO(), &sar, metav1.CreateOptions{})
	if err != nil {
		log.Error("Can not create SubjectAccessReviews resource", zap.Error(err))

		return false, fmt.Errorf("can not perform Subject Access Review action: %w", err)
	}

	return res.Status.Allowed, nil
}

// SetupWithManager configures controller for manager to handle CertificateSigningRequest.
func (r *SigningReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr). //nolint:wrapcheck
							For(&certificatesv1.CertificateSigningRequest{}).
							Complete(r)
}

// appendApprovalCondition sets fields for audit purpose.
func appendApprovalCondition(csr *certificatesv1.CertificateSigningRequest) {
	csr.Status.Conditions = append(csr.Status.Conditions, certificatesv1.CertificateSigningRequestCondition{
		Type:           certificatesv1.CertificateApproved,
		Status:         corev1.ConditionTrue,
		Reason:         "Approved by Kubelet Serving Certificate Approver",
		LastUpdateTime: metav1.Time{Time: time.Now().UTC()},
		Message:        "Auto approving Kubelet Serving Certificate after Subject Access Review.",
	})
}
