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
	"go.uber.org/zap"

	"github.com/alex1989hu/kubelet-serving-cert-approver/logger"
)

//nolint:gochecknoglobals
var log = logger.CreateLogger().With(zap.String("controller", "certificatesigningrequest"))

// SetLogger sets the package logger
func SetLogger(logger *zap.Logger) {
	log = logger
}
