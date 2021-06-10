# Copyright 2021 Alex Szakaly
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

@healthcheck
@liveness

Feature: Kubernetes Liveness Probe
  As an administrator
  In order to track application liveness
  I need to be able to have endpoint for liveness probe

  Background:
    Given there is a running Pod in namespace "kubelet-serving-cert-approver" with label "app.kubernetes.io/name=kubelet-serving-cert-approver"
    And the Pod shall provide "/healthz" endpoint at port 8080

  Scenario: Application shall provide liveness probe endpoint
    Then response shall contain "ok"
