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

//nolint:testpackage,wrapcheck // Need to reach functions.
package certificatesigningrequest

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	authorizationv1 "k8s.io/api/authorization/v1"
	certificatesv1 "k8s.io/api/certificates/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/fake"
	clientgotesting "k8s.io/client-go/testing"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"
	ctrlfake "sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

const (
	name      = "csr-e7dbe"
	namespace = "test-namespace"
)

//nolint:gochecknoglobals
var (
	// rsaPrivateKey provides RSA private key for Certificate Signing Request generation.
	rsaPrivateKey *rsa.PrivateKey

	// errMockGet defines an error during mocked client-go Get(...) function.
	errMockGet = errors.New("mocked get error")

	// errAuthorization defines an error during mocked Subject Access Review.
	errAuthorization = errors.New("mocked authorization error")

	// errApprovalUpdate defines an error during Certificate Signing Request Approval update.
	errApprovalUpdate = errors.New("mocked update error")
)

// Client is a mock for the controller-runtime dynamic client interface.
type Client struct {
	mapper     meta.RESTMapper
	scheme     *runtime.Scheme
	StatusMock *StatusClient
	mock.Mock
}

func (c *Client) SubResource(subResource string) client.SubResourceClient {
	args := c.Called(subResource)

	return args.Get(0).(client.SubResourceClient)
}

// NewMockClient creates a new mock controller-runtime client.
func NewMockClient() *Client {
	return &Client{
		StatusMock: &StatusClient{},
	}
}

// Status fulfills StatusClient interface.
func (c *Client) Status() client.StatusWriter {
	return c.StatusMock
}

// Get fulfills Reader interface.
func (c *Client) Get(ctx context.Context, key client.ObjectKey, obj client.Object, opts ...client.GetOption) error {
	args := c.Called(ctx, key, obj, opts)

	return args.Error(0)
}

// List fulfills Reader interface.
func (c *Client) List(context.Context, client.ObjectList, ...client.ListOption) error {
	args := c.Called()

	return args.Error(0)
}

// Create fulfills Writer interface.
func (c *Client) Create(context.Context, client.Object, ...client.CreateOption) error {
	args := c.Called()

	return args.Error(0)
}

// Delete fulfills Writer interface.
func (c *Client) Delete(context.Context, client.Object, ...client.DeleteOption) error {
	args := c.Called()

	return args.Error(0)
}

// Update fulfills Writer interface.
func (c *Client) Update(ctx context.Context, obj client.Object, opts ...client.UpdateOption) error {
	args := c.Called(ctx, obj, opts)

	return args.Error(0)
}

// Patch fulfills Writer interface.
func (c *Client) Patch(context.Context, client.Object, client.Patch, ...client.PatchOption) error {
	args := c.Called()

	return args.Error(0)
}

// DeleteAllOf fulfills Writer interface.
func (c *Client) DeleteAllOf(context.Context, client.Object, ...client.DeleteAllOfOption) error {
	args := c.Called()

	return args.Error(0)
}

func (c *Client) Scheme() *runtime.Scheme {
	return c.scheme
}

func (c *Client) RESTMapper() meta.RESTMapper {
	return c.mapper
}

type StatusClient struct {
	mock.Mock
}

// Create fulfills SubResourceWriter interface.
func (c *StatusClient) Create(ctx context.Context, obj client.Object, subResource client.Object,
	opts ...client.SubResourceCreateOption,
) error {
	args := c.Called(ctx, obj, subResource, opts)

	return args.Error(0)
}

// Update fulfills SubResourceWriter interface.
func (c *StatusClient) Update(ctx context.Context, obj client.Object, opts ...client.SubResourceUpdateOption) error {
	args := c.Called(ctx, obj, opts)

	return args.Error(0)
}

// Patch fulfills SubResourceWriter interface.
func (c *StatusClient) Patch(ctx context.Context, obj client.Object, patch client.Patch,
	opts ...client.SubResourcePatchOption,
) error {
	args := c.Called(ctx, obj, patch, opts)

	return args.Error(0)
}

func TestReconcileClientGetError(t *testing.T) {
	t.Parallel()

	mockClient := NewMockClient()

	mockClient.On("Get", mock.Anything, types.NamespacedName{},
		mock.Anything, mock.AnythingOfType("[]client.GetOption")).Return(errMockGet).Times(1)

	signingReconciler := &SigningReconciler{Client: mockClient, Scheme: runtime.NewScheme(), Logger: TestLogger}
	req := reconcile.Request{NamespacedName: types.NamespacedName{}}
	res, err := signingReconciler.Reconcile(context.TODO(), req)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), errMockGet.Error())
	assert.Equal(t, reconcile.Result{}, res)

	mockClient.AssertExpectations(t)
}

func TestReconcileClientGetNotFoundError(t *testing.T) {
	t.Parallel()

	mockClient := NewMockClient()
	errNotFound := k8serrors.NewNotFound(schema.GroupResource{Group: "group", Resource: "resource"}, "bela")

	mockClient.On("Get", mock.Anything, types.NamespacedName{},
		mock.Anything, mock.AnythingOfType("[]client.GetOption")).Return(errNotFound).Times(1)

	signingReconciler := &SigningReconciler{Client: mockClient, Scheme: runtime.NewScheme(), Logger: TestLogger}
	req := reconcile.Request{NamespacedName: types.NamespacedName{}}
	res, err := signingReconciler.Reconcile(context.TODO(), req)

	assert.Nil(t, err)
	assert.Equal(t, reconcile.Result{}, res)

	mockClient.AssertExpectations(t)
}

func TestReconcileSwitchCasesNegativePath(t *testing.T) {
	t.Parallel()

	tables := []struct {
		goal string
		csr  certificatesv1.CertificateSigningRequest
	}{
		{
			goal: "Not Kubelet serving Certificate SigningRequest",
			csr: certificatesv1.CertificateSigningRequest{
				Spec: certificatesv1.CertificateSigningRequestSpec{
					SignerName: "foobar",
				},
			},
		},
		{
			goal: "Deleted CertificateSigningRequest",
			csr: certificatesv1.CertificateSigningRequest{
				ObjectMeta: metav1.ObjectMeta{
					DeletionTimestamp: &metav1.Time{
						Time: time.Now().UTC(),
					},
					Name:      name,
					Namespace: namespace,
				},
				Spec: certificatesv1.CertificateSigningRequestSpec{
					SignerName: certificatesv1.KubeletServingSignerName,
				},
			},
		},
		{
			goal: "Already signed CertificateSigningRequest",
			csr: certificatesv1.CertificateSigningRequest{
				ObjectMeta: metav1.ObjectMeta{
					Name:      name,
					Namespace: namespace,
				},
				Spec: certificatesv1.CertificateSigningRequestSpec{
					SignerName: certificatesv1.KubeletServingSignerName,
				},
				Status: certificatesv1.CertificateSigningRequestStatus{
					Certificate: []byte(`foo`),
				},
			},
		},
		{
			goal: "Already has approval condition: approved",
			csr: certificatesv1.CertificateSigningRequest{
				ObjectMeta: metav1.ObjectMeta{
					Name:      name,
					Namespace: namespace,
				},
				Spec: certificatesv1.CertificateSigningRequestSpec{
					SignerName: certificatesv1.KubeletServingSignerName,
				},
				Status: certificatesv1.CertificateSigningRequestStatus{
					Conditions: []certificatesv1.CertificateSigningRequestCondition{
						{
							Type: certificatesv1.CertificateApproved,
						},
					},
				},
			},
		},
		{
			goal: "Already has approval condition: denied",
			csr: certificatesv1.CertificateSigningRequest{
				ObjectMeta: metav1.ObjectMeta{
					Name:      name,
					Namespace: namespace,
				},
				Spec: certificatesv1.CertificateSigningRequestSpec{
					SignerName: certificatesv1.KubeletServingSignerName,
				},
				Status: certificatesv1.CertificateSigningRequestStatus{
					Conditions: []certificatesv1.CertificateSigningRequestCondition{
						{
							Type: certificatesv1.CertificateDenied,
						},
					},
				},
			},
		},
		{
			goal: "Already has approval condition: failed",
			csr: certificatesv1.CertificateSigningRequest{
				ObjectMeta: metav1.ObjectMeta{
					Name:      name,
					Namespace: namespace,
				},
				Spec: certificatesv1.CertificateSigningRequestSpec{
					SignerName: certificatesv1.KubeletServingSignerName,
				},
				Status: certificatesv1.CertificateSigningRequestStatus{
					Conditions: []certificatesv1.CertificateSigningRequestCondition{
						{
							Type: certificatesv1.CertificateFailed,
						},
					},
				},
			},
		},
	}

	for _, table := range tables {
		table := table // pin!

		t.Run(fmt.Sprint(table.goal), func(t *testing.T) {
			t.Parallel()

			fakeClient := ctrlfake.NewClientBuilder().WithRuntimeObjects([]runtime.Object{&table.csr}...).Build()
			r := &SigningReconciler{
				Client: fakeClient,
				Scheme: runtime.NewScheme(),
				Logger: TestLogger,
			}
			req := reconcile.Request{
				NamespacedName: types.NamespacedName{
					Name:      table.csr.GetName(),
					Namespace: table.csr.GetNamespace(),
				},
			}

			res, err := r.Reconcile(context.TODO(), req)

			assert.Nil(t, err)
			assert.Equal(t, reconcile.Result{}, res)
		})
	}
}

func TestReconcileValidCSR(t *testing.T) {
	t.Parallel()

	csr := certificatesv1.CertificateSigningRequest{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: certificatesv1.CertificateSigningRequestSpec{
			Extra: map[string]certificatesv1.ExtraValue{
				"Kubernetes Tiers": []string{"control-plane", "worker"},
			},
			Usages:     validUsages,
			Username:   validUsername,
			SignerName: certificatesv1.KubeletServingSignerName,
			Request:    generatePEMEncodedCSR(t),
		},
	}

	fakeEventRecorder := record.NewFakeRecorder(1)
	fakeClientset := fake.Clientset{}

	// Provide authorization by fake k8s clientset
	fakeClientset.Fake.PrependReactor(
		"create",
		"subjectaccessreviews",
		func(action clientgotesting.Action) (handled bool, ret runtime.Object, err error) {
			sar := &authorizationv1.SubjectAccessReview{
				Status: authorizationv1.SubjectAccessReviewStatus{
					Allowed: true,
					Reason:  "test",
				},
			}

			return true, sar, nil
		})

	fakeClient := ctrlfake.NewClientBuilder().WithRuntimeObjects([]runtime.Object{&csr}...).Build()
	r := &SigningReconciler{
		Client:        fakeClient,
		ClientSet:     &fakeClientset,
		EventRecorder: fakeEventRecorder,
		Scheme:        runtime.NewScheme(),
		Logger:        TestLogger,
	}
	req := reconcile.Request{
		NamespacedName: types.NamespacedName{
			Name:      csr.GetName(),
			Namespace: csr.GetNamespace(),
		},
	}

	res, err := r.Reconcile(context.TODO(), req)

	assert.Nil(t, err)
	assert.Equal(t, reconcile.Result{}, res)
	assert.Len(t, fakeEventRecorder.Events, 1)
}

func TestReconcileParseCSRError(t *testing.T) {
	t.Parallel()

	csr := certificatesv1.CertificateSigningRequest{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: certificatesv1.CertificateSigningRequestSpec{
			Usages:     validUsages,
			Username:   validUsername,
			SignerName: certificatesv1.KubeletServingSignerName,
			Request:    []byte(`foobar`),
		},
	}

	fakeEventRecorder := record.NewFakeRecorder(1)
	fakeClientset := fake.Clientset{}

	fakeClientset.Fake.PrependReactor(
		"create",
		"subjectaccessreviews",
		func(action clientgotesting.Action) (handled bool, ret runtime.Object, err error) {
			sar := &authorizationv1.SubjectAccessReview{
				Status: authorizationv1.SubjectAccessReviewStatus{
					Allowed: true,
					Reason:  "test",
				},
			}

			return true, sar, nil
		})

	fakeClient := ctrlfake.NewClientBuilder().WithRuntimeObjects([]runtime.Object{&csr}...).Build()
	r := &SigningReconciler{
		Client:        fakeClient,
		ClientSet:     &fakeClientset,
		EventRecorder: fakeEventRecorder,
		Scheme:        runtime.NewScheme(),
		Logger:        TestLogger,
	}

	req := reconcile.Request{
		NamespacedName: types.NamespacedName{
			Name:      csr.GetName(),
			Namespace: csr.GetNamespace(),
		},
	}

	res, err := r.Reconcile(context.TODO(), req)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "PEM Block")
	assert.Equal(t, reconcile.Result{}, res)
	assert.Len(t, fakeEventRecorder.Events, 1)
}

func TestReconcileRecognizeError(t *testing.T) {
	t.Parallel()

	csr := certificatesv1.CertificateSigningRequest{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: certificatesv1.CertificateSigningRequestSpec{
			Usages:     validUsages,
			Username:   "foo-system",
			SignerName: certificatesv1.KubeletServingSignerName,
			Request:    generatePEMEncodedCSR(t),
		},
	}

	fakeEventRecorder := record.NewFakeRecorder(1)
	fakeClientset := fake.Clientset{}

	fakeClientset.Fake.PrependReactor(
		"create",
		"subjectaccessreviews",
		func(action clientgotesting.Action) (handled bool, ret runtime.Object, err error) {
			return true, &authorizationv1.SubjectAccessReview{}, nil
		})

	fakeClient := ctrlfake.NewClientBuilder().WithRuntimeObjects([]runtime.Object{&csr}...).Build()
	r := &SigningReconciler{
		Client:        fakeClient,
		ClientSet:     &fakeClientset,
		EventRecorder: fakeEventRecorder,
		Scheme:        runtime.NewScheme(),
		Logger:        TestLogger,
	}

	req := reconcile.Request{
		NamespacedName: types.NamespacedName{
			Name:      csr.GetName(),
			Namespace: csr.GetNamespace(),
		},
	}

	res, err := r.Reconcile(context.TODO(), req)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "x509 Common Name")
	assert.Equal(t, reconcile.Result{}, res)
	assert.Len(t, fakeEventRecorder.Events, 1)
}

func TestReconcileAuthorizationError(t *testing.T) {
	t.Parallel()

	csr := certificatesv1.CertificateSigningRequest{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: certificatesv1.CertificateSigningRequestSpec{
			Usages:     validUsages,
			Username:   validUsername,
			SignerName: certificatesv1.KubeletServingSignerName,
			Request:    generatePEMEncodedCSR(t),
		},
	}

	fakeEventRecorder := record.NewFakeRecorder(1)
	fakeClientset := fake.Clientset{}

	fakeClientset.Fake.PrependReactor(
		"create",
		"subjectaccessreviews",
		func(action clientgotesting.Action) (handled bool, ret runtime.Object, err error) {
			return true, &authorizationv1.SubjectAccessReview{}, errAuthorization
		})

	fakeClient := ctrlfake.NewClientBuilder().WithRuntimeObjects([]runtime.Object{&csr}...).Build()
	r := &SigningReconciler{
		Client:        fakeClient,
		ClientSet:     &fakeClientset,
		EventRecorder: fakeEventRecorder,
		Scheme:        runtime.NewScheme(),
		Logger:        TestLogger,
	}

	req := reconcile.Request{
		NamespacedName: types.NamespacedName{
			Name:      csr.GetName(),
			Namespace: csr.GetNamespace(),
		},
	}

	res, err := r.Reconcile(context.TODO(), req)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), errAuthorization.Error())
	assert.Equal(t, reconcile.Result{}, res)
	assert.Len(t, fakeEventRecorder.Events, 1)
}

func TestReconcileAuthorizationDenied(t *testing.T) {
	t.Parallel()

	csr := certificatesv1.CertificateSigningRequest{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: certificatesv1.CertificateSigningRequestSpec{
			Usages:     validUsages,
			Username:   validUsername,
			SignerName: certificatesv1.KubeletServingSignerName,
			Request:    generatePEMEncodedCSR(t),
		},
	}

	fakeEventRecorder := record.NewFakeRecorder(1)
	fakeClientset := fake.Clientset{}

	fakeClientset.Fake.PrependReactor(
		"create",
		"subjectaccessreviews",
		func(action clientgotesting.Action) (handled bool, ret runtime.Object, err error) {
			sar := &authorizationv1.SubjectAccessReview{
				Status: authorizationv1.SubjectAccessReviewStatus{
					Allowed: false,
					Reason:  "test",
				},
			}

			return true, sar, nil
		})

	fakeClient := ctrlfake.NewClientBuilder().WithRuntimeObjects([]runtime.Object{&csr}...).Build()
	r := &SigningReconciler{
		Client:        fakeClient,
		ClientSet:     &fakeClientset,
		EventRecorder: fakeEventRecorder,
		Scheme:        runtime.NewScheme(),
		Logger:        TestLogger,
	}

	req := reconcile.Request{
		NamespacedName: types.NamespacedName{
			Name:      csr.GetName(),
			Namespace: csr.GetNamespace(),
		},
	}

	res, err := r.Reconcile(context.TODO(), req)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "Subject Access Review")
	assert.Equal(t, reconcile.Result{}, res)
	assert.Len(t, fakeEventRecorder.Events, 1)
}

func TestReconcileUpdateApprovalError(t *testing.T) {
	t.Parallel()

	csr := certificatesv1.CertificateSigningRequest{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: certificatesv1.CertificateSigningRequestSpec{
			Usages:     validUsages,
			Username:   validUsername,
			SignerName: certificatesv1.KubeletServingSignerName,
			Request:    generatePEMEncodedCSR(t),
		},
	}

	fakeEventRecorder := record.NewFakeRecorder(1)
	fakeClientset := fake.Clientset{}

	// Provide authorization by fake k8s clientset
	fakeClientset.Fake.PrependReactor(
		"create",
		"subjectaccessreviews",
		func(action clientgotesting.Action) (handled bool, ret runtime.Object, err error) {
			sar := &authorizationv1.SubjectAccessReview{
				Status: authorizationv1.SubjectAccessReviewStatus{
					Allowed: true,
					Reason:  "test",
				},
			}

			return true, sar, nil
		})

	fakeClientset.Fake.PrependReactor(
		"update",
		"certificatesigningrequests",
		func(action clientgotesting.Action) (handled bool, ret runtime.Object, err error) {
			return true, &certificatesv1.CertificateSigningRequest{}, errApprovalUpdate
		})

	fakeClient := ctrlfake.NewClientBuilder().WithRuntimeObjects([]runtime.Object{&csr}...).Build()
	r := &SigningReconciler{
		Client:        fakeClient,
		ClientSet:     &fakeClientset,
		EventRecorder: fakeEventRecorder,
		Scheme:        runtime.NewScheme(),
		Logger:        TestLogger,
	}

	req := reconcile.Request{
		NamespacedName: types.NamespacedName{
			Name:      csr.GetName(),
			Namespace: csr.GetNamespace(),
		},
	}

	res, err := r.Reconcile(context.TODO(), req)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), errApprovalUpdate.Error())
	assert.Equal(t, reconcile.Result{}, res)
	assert.Len(t, fakeEventRecorder.Events, 1)
}

func TestReconcileNilManager(t *testing.T) {
	t.Parallel()

	r := SigningReconciler{}
	assert.Error(t, r.SetupWithManager(nil))
}

// TestAppendApprovalOptions ensures that the approval options will not change accidentally.
func TestAppendApprovalOptions(t *testing.T) {
	t.Parallel()

	csr := certificatesv1.CertificateSigningRequest{}

	appendApprovalCondition(&csr)
	assert.Len(t, csr.Status.Conditions, 1)

	condition := csr.Status.Conditions[0]
	assert.Equal(t, certificatesv1.CertificateApproved, condition.Type)
	assert.Contains(t, condition.Reason, "Kubelet Serving Certificate Approver")
	assert.Contains(t, condition.Message, "Auto approving Kubelet Serving Certificate")
}

func init() {
	res, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	rsaPrivateKey = res
}

// generatePEMEncodedCSR creates PEM encoded Certificate Signing Request.
func generatePEMEncodedCSR(t *testing.T) []byte {
	t.Helper()

	csrTemplate := x509.CertificateRequest{
		Subject: pkix.Name{
			Organization: []string{validOrganization},
			CommonName:   validUsername,
		},
		DNSNames:           validDNSNames,
		SignatureAlgorithm: x509.SHA256WithRSA,
	}

	csrCertificate, err := x509.CreateCertificateRequest(rand.Reader, &csrTemplate, rsaPrivateKey)
	if err != nil {
		t.Fatalf("Can not create Certificate Request %v", err)
	}

	csr := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE REQUEST",
		Bytes: csrCertificate,
	})

	return csr
}
