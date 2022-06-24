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

package certificatesigningrequest //nolint:testpackage // Need to reach functions.

import (
	"testing"

	"go.uber.org/goleak"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//nolint:gochecknoglobals
var TestLogger *zap.Logger

func TestMain(m *testing.M) {
	// Create logger with enabled debug logging to reveal any possible issue due to zap CheckedEntry logging
	TestLogger, _ = zap.Config{ //nolint:errcheck
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
		Development: true,
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}.Build()

	// TestMain is needed due to t.Parallel() incompatibility of goleak.
	// https://github.com/uber-go/goleak/issues/16
	goleak.VerifyTestMain(m,
		// controller-runtime's log.init intentionally waits: https://github.com/kubernetes-sigs/controller-runtime/pull/1309
		goleak.IgnoreTopFunction("time.Sleep"),
		// flushDaemon leaks: https://github.com/kubernetes/client-go/issues/900
		goleak.IgnoreTopFunction("k8s.io/klog/v2.(*loggingT).flushDaemon"),
	)
}
