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

package logger_test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
	"go.uber.org/zap"

	"github.com/alex1989hu/kubelet-serving-cert-approver/logger"
)

// viperLock helps viper to avoid viper.Set() concurrency issues - only used in test code.
//nolint: gochecknoglobals
var viperLock sync.Mutex

func TestLoggingConfiguration(t *testing.T) {
	t.Parallel()

	tables := []struct {
		debugEnabled bool
	}{
		{
			debugEnabled: false,
		},
		{
			debugEnabled: true,
		},
	}

	for _, table := range tables { //nolint: paralleltest // Disable due to linter bug.
		table := table // scopelint, pin!

		t.Run(fmt.Sprintf("Debug level enabled: (%t)", table.debugEnabled), func(t *testing.T) {
			t.Parallel()

			viperLock.Lock()
			defer viperLock.Unlock()

			viper.Set("debug", table.debugEnabled)

			zapLog := logger.CreateLogger()
			assert.NotNil(t, zapLog)
			assert.Equal(t, table.debugEnabled, zapLog.Core().Enabled(zap.DebugLevel))
		})
	}
}

// TestMain is needed due to t.Parallel() incompatibility of goleak.
// https://github.com/uber-go/goleak/issues/16
func TestMain(m *testing.M) { //nolint: interfacer
	// flushDaemon leaks: https://github.com/kubernetes/client-go/issues/900
	goleak.VerifyTestMain(m, goleak.IgnoreTopFunction("k8s.io/klog/v2.(*loggingT).flushDaemon"))
}
