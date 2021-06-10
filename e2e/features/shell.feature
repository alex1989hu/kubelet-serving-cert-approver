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

@security
@shell

Feature: Security Hardening
  As an administrator
  In order to reduce attack surface
  I need to not to have shell in the container

  Background:
    Given there is a running Pod in namespace "kubelet-serving-cert-approver" with label "app.kubernetes.io/name=kubelet-serving-cert-approver"

  Scenario: Container shall not provide any shell
    When I execute command <command> in the running Pod
    Then command execution shall report error
    And command execution error message shall contain:
      """
      failed to exec in container: failed to start exec
      """

    Examples:
      | command         |
      | ""              |
      | "apt"           |
      | "bash"          |
      | "curl"          |
      | "cut"           |
      | "echo"          |
      | "ls"            |
      | "printf"        |
      | "sh"            |
      | "wget"          |
      | "zsh"           |
      | "/bin/bash"     |
      | "/bin/sh"       |
      | "/bin/zsh"      |
      | "/usr/bin/bash" |
      | "/usr/bin/sh"   |
      | "/usr/bin/zsh"  |

  Scenario: Container shall allow execution of the application itself but nothing besides that
    When I execute command "/app/kubelet-serving-cert-approver" in the running Pod
    Then command execution shall not report any error
