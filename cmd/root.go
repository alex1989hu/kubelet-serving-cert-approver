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
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/alex1989hu/kubelet-serving-cert-approver/build"
)

const defaultNamespace = "kubelet-serving-cert-approver"

//nolint: gochecknoglobals
var (
	// namespace represents where to record Kubernetes Events.
	namespace string

	// isDebug represents the debug level of logger.
	isDebug bool

	// enableLeaderElection represents highly available (HA) mode.
	enableLeaderElection bool
)

// rootCmd represents the base command when called without any subcommands.
//nolint: gochecknoglobals
var rootCmd = &cobra.Command{
	Use:     "kubelet-serving-cert-approver",
	Version: fmt.Sprintf("(%s/%s)", build.GitBranch, build.GitCommit),
	Short:   "",
	Long: `Approves kubernetes.io/kubelet-serving Certificate Signing Request
that kubelet use to serve TLS endpoints.
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Execution failed: %v", err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().BoolVarP(&isDebug, "debug", "d", false,
		"Enable debug logging")

	rootCmd.PersistentFlags().BoolVar(&enableLeaderElection, "enable-leader-election", false,
		"Enable leader election for highly available (HA) mode")

	rootCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", defaultNamespace,
		"Configure namespace where to execute")

	if err := viper.BindPFlags(rootCmd.PersistentFlags()); err != nil {
		log.Fatalf("Can not bind flags: %v", err)
	}
}

// initConfig reads environment variables if set.
func initConfig() {
	viper.AutomaticEnv()

	namespace = viper.GetString("namespace")
}
