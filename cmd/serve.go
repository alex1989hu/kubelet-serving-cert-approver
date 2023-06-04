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

package cmd

import (
	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"github.com/spf13/cobra"
	uberzap "go.uber.org/zap"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgokubernetes "k8s.io/client-go/kubernetes"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/klog/v2"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	ctrllog "sigs.k8s.io/controller-runtime/pkg/log"
	ctrlmetrics "sigs.k8s.io/controller-runtime/pkg/metrics"

	"github.com/alex1989hu/kubelet-serving-cert-approver/build"
	"github.com/alex1989hu/kubelet-serving-cert-approver/controller/certificatesigningrequest"
	"github.com/alex1989hu/kubelet-serving-cert-approver/logger"
	"github.com/alex1989hu/kubelet-serving-cert-approver/metrics"
)

//nolint:gochecknoglobals
var (
	scheme = runtime.NewScheme()

	// serveCmd represents the generate command.
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Starts operator for Kubelet Serving Certificate Approver",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			startServer()
		},
	}
)

func init() {
	rootCmd.AddCommand(serveCmd)

	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
}

func startServer() {
	setupLog := *logger.CreateLogger()

	setupLog.Info("Kubelet Serving Certificate Approver has been started",
		uberzap.String("GitBranch", build.GitBranch),
		uberzap.String("GitCommit", build.GitCommit),
		uberzap.String("Time", build.Time),
		uberzap.String("namespace", namespace),
		uberzap.Bool("ha", enableLeaderElection),
		uberzap.Bool("debug", isDebug))

	// Forward client-go klog calls to zap
	klog.SetLogger(zapr.NewLogger(&setupLog))

	setupLog.Info("Try to talk to Kubernetes API Server, will exit in case of failure")

	pProfBindAddress := "0"

	if isDebug {
		pProfBindAddress = ":8081"

		setupLog.Info("pprof will be enabled", uberzap.String("port", pProfBindAddress))
	}

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:                     scheme,
		MetricsBindAddress:         ":9090",
		HealthProbeBindAddress:     ":8080",
		LeaderElection:             enableLeaderElection,
		LeaderElectionNamespace:    namespace,
		LeaderElectionResourceLock: "leases",
		LeaderElectionID:           "kubelet-serving-certificate-approver",
		// Set NullLogger: https://github.com/kubernetes-sigs/controller-runtime/issues/1122
		Logger:           logr.New(ctrllog.NullLogSink{}),
		PprofBindAddress: pProfBindAddress,
	})
	if err != nil {
		setupLog.Fatal("Unable to start manager", uberzap.Error(err))
	}

	setupLog.Info("Successfully connected to Kubernetes API Server")

	// Add readiness probe
	if errReadyzCheck := mgr.AddReadyzCheck("ready-ping", healthz.Ping); errReadyzCheck != nil {
		setupLog.Fatal("Unable to add readiness check", uberzap.Error(errReadyzCheck))
	}

	// Add liveness probe
	if errHealthzCheck := mgr.AddHealthzCheck("health-ping", healthz.Ping); errHealthzCheck != nil {
		setupLog.Fatal("Unable to add health check", uberzap.Error(errHealthzCheck))
	}

	if err = (&certificatesigningrequest.SigningReconciler{
		Client:        mgr.GetClient(),
		ClientSet:     clientgokubernetes.NewForConfigOrDie(mgr.GetConfig()),
		Scheme:        mgr.GetScheme(),
		EventRecorder: mgr.GetEventRecorderFor("kubelet-serving-cert-aprover"),
		Logger:        logger.CreateLogger().With(uberzap.String("controller", "certificatesigningrequest")),
	}).SetupWithManager(mgr); err != nil {
		setupLog.Fatal("Unable to create controller", uberzap.Error(err))
	}

	// Add our Prometheus metrics
	ctrlmetrics.Registry.MustRegister(metrics.NumberOfApprovedCertificateRequests)
	ctrlmetrics.Registry.MustRegister(metrics.NumberOfInvalidCertificateSigningRequests)

	setupLog.Debug("Certificate Signing Request Reconciler has ben added")

	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Fatal("Unable to start manager", uberzap.Error(err))
	}
}
