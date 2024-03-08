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

FROM golang:1.21.8 as builder

# To let GitHub CI driven buildx pass build arguments
ARG TARGETOS
ARG TARGETARCH

# Set necessary environment variables
ENV CGO_ENABLED=0 \
    GOOS=$TARGETOS \
    GOARCH=$TARGETARCH

# All these steps will be cached
WORKDIR /build
COPY go.mod .
COPY go.sum .

# Get dependencies - will also be cached if we won't change mod/sum
RUN go mod download
# Copy the source code as the last step
COPY . .

# Build the binary
RUN GIT_COMMIT=$(git rev-parse --short=8 HEAD || echo "dev" ) && \
    GIT_BRANCH=$(git rev-parse --symbolic-full-name --abbrev-ref HEAD || echo "dirty" ) && \
    BUILD_TIME=$(date -u) && \
    go build -trimpath -o /app/kubelet-serving-cert-approver \
      -ldflags "-buildid= -w -s \
        -X 'github.com/alex1989hu/kubelet-serving-cert-approver/build.GitBranch=$GIT_BRANCH' \
        -X 'github.com/alex1989hu/kubelet-serving-cert-approver/build.GitCommit=$GIT_COMMIT' \
        -X 'github.com/alex1989hu/kubelet-serving-cert-approver/build.Time=$BUILD_TIME'" && \
    if [ "$GOARCH" = "amd64" ]; then CGO_ENABLED=1 go test -race ./... -v ; else go test ./... -v ; fi;

# Production image
FROM gcr.io/distroless/static-debian12@sha256:16f75ae7665b13825daffba81f12d6b1a16d0e1217c562fadfce0ba77ca7b891

COPY --from=builder /app/kubelet-serving-cert-approver /app/kubelet-serving-cert-approver

WORKDIR /app

USER 65534:65534

EXPOSE 8080 9090

ENTRYPOINT ["/app/kubelet-serving-cert-approver"]
