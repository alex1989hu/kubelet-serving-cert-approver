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

@csr
@certificatesigningrequest

Feature: Kubelet Serving TLS Certificate Signing Request Approval
  As an administrator
  In order to securely communicate with kubelet
  I need to have a valid certificate issued by Kubernetes Root Certificate Authority

  Scenario: Kubelet Serving TLS Certificate Signing Requests shall have approval condition
    Given there are "kubernetes.io/kubelet-serving" Certificate Signing Requests
    Then Certificate Signing Requests shall have approval condition

  Scenario: Kubelet Serving TLS Certificate Signing Requests shall be approved
    Given there are "kubernetes.io/kubelet-serving" Certificate Signing Requests
    Then Certificate Signing Requests shall be approved
