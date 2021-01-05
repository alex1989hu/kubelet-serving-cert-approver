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

analytics_settings(False)

settings = read_json('./hack/tilt-config.json', default={})

docker_build(settings.get('default_registry') + '/' +  settings.get('image_repository') + '/kubelet-serving-cert-approver',
    '.',
    dockerfile='Dockerfile',
    ignore=[
        './.github/',
        './.idea/',
        './.vscode/',
        './deploy/',
        './hack/',
        './*.md',
        './codecov.yml',
        './.gitattributes',
        './.gitignore',
        './.golangci.yml',
    ]
)

k8s_yaml('./deploy/ha-install.yaml')

k8s_resource(
    workload='kubelet-serving-cert-approver',
    port_forwards=[8080, 9090],
    objects=['kubelet-serving-cert-approver:namespace',
        'kubelet-serving-cert-approver:serviceaccount',
        'kubelet-serving-cert-approver:podsecuritypolicy',
        'leader-election\\:kubelet-serving-cert-approver:role',
        'certificates\\:kubelet-serving-cert-approver:clusterrole',
        'events\\:kubelet-serving-cert-approver:clusterrole',
        'psp\\:kubelet-serving-cert-approver:clusterrole',
        'events\\:kubelet-serving-cert-approver:rolebinding',
        'leader-election\\:kubelet-serving-cert-approver:rolebinding',
        'psp\\:kubelet-serving-cert-approver:rolebinding',
        'kubelet-serving-cert-approver:clusterrolebinding'
    ],
    new_name='cluster-setup',
)
