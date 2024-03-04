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

//go:build e2e

//nolint:wrapcheck
package e2e_test

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/cucumber/godog"
	"github.com/prometheus/common/expfmt"
	"github.com/stretchr/testify/assert"
	certificatesv1 "k8s.io/api/certificates/v1"
	corev1 "k8s.io/api/core/v1"
	eventsv1 "k8s.io/api/events/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientgokubernetes "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/portforward"
	"k8s.io/client-go/tools/remotecommand"
	"k8s.io/client-go/transport"
	"k8s.io/client-go/transport/spdy"
)

const expectationDoesNotMeetMessage = "expectation does not meet"

// commandResult holds command execution result.
type commandResult struct {
	Error  error
	StdOut string
	StdErr string
}

type ApproverInstance struct {
	Clientset                     *clientgokubernetes.Clientset
	RestConfig                    *rest.Config
	Pod                           corev1.Pod
	CommandResult                 commandResult
	CertificateSigningRequestList []certificatesv1.CertificateSigningRequest
	Events                        []eventsv1.Event
	Metrics                       []string
	Response                      []byte
}

// InitializeScenario sets context and defines steps being used in scenarios.
func InitializeScenario(s *godog.ScenarioContext) {
	var instance ApproverInstance

	s.Before(func(ctx context.Context, _ *godog.Scenario) (context.Context, error) {
		clientset, restConfig, err := createNewClientSet()
		if err != nil {
			return ctx, err
		}

		instance = ApproverInstance{Clientset: clientset, RestConfig: restConfig}

		return ctx, nil
	})

	// Steps for testing features related to Certificate Signing Request Approval
	s.Step(`^there are "([^"]*)" Certificate Signing Requests$`,
		instance.thereAreCertificateSigningRequests)

	s.Step(`^Certificate Signing Requests shall have approval condition$`,
		instance.certificateSigningRequestsShallHaveApprovalCondition)

	s.Step(`^Certificate Signing Requests shall be approved$`,
		instance.certificateSigningRequestsShallBeApproved)

	// Steps for testing features related to Event Recording
	s.Step(`^there are events related to Certificate Signing Requests$`,
		instance.thereAreEventsRelatedToCertificateSigningRequests)

	s.Step(`^approval events shall contain "([^"]*)" reason$`,
		instance.approvalEventsShallContainReason)

	s.Step(`^approval events shall have "([^"]*)" message$`,
		instance.approvalEventsShallHaveMessage)

	// Steps for testing features related to Prometheus Metrics
	s.Step(`^there is a running Pod in namespace "([^"]*)" with label "([^"]*)"$`,
		instance.thereIsARunningPodInNamespaceWithLabel)

	s.Step(`^the Pod shall provide "([^"]*)" endpoint at port (\d+)$`,
		instance.thePodShallProvideEndpointAtPort,
	)

	s.Step(`^response shall be parseable Prometheus Metrics$`,
		instance.responseShallBeParseablePrometheusMetrics,
	)

	s.Step(`^metrics shall contain "([^"]*)" metric$`,
		instance.metricsShallContainMetric)

	// Steps for testing features related to Health Check
	s.Step(`^response shall contain "([^"]*)"$`,
		instance.responseShallContain,
	)

	// Steps for testing security hardening features
	s.Step(`^I execute command "([^"]*)" in the running Pod$`,
		instance.iExecuteCommandInTheRunningPod,
	)

	s.Step(`^command execution shall report error$`,
		instance.commandExecutionShallReportError,
	)

	s.Step(`^command execution shall not report any error$`,
		instance.commandExecutionShallNotReportAnyError,
	)

	s.Step(`^command execution error message shall contain:$`,
		instance.commandExecutionErrorMessageShallContain,
	)
}

// thereAreCertificateSigningRequests ensures that there is already existing Certificate Signing Request.
// +feature: certificatesigningrequest
func (c *ApproverInstance) thereAreCertificateSigningRequests(signer string) error {
	csrList, errListCertificates := c.Clientset.CertificatesV1().CertificateSigningRequests().List(
		context.TODO(), metav1.ListOptions{
			FieldSelector: "spec.signerName=" + signer,
		})

	if err := assertActual(assert.Nil, errListCertificates); err != nil {
		return fmt.Errorf("%s: %w", expectationDoesNotMeetMessage, err)
	}

	if err := assertExpectedAndActual(assert.GreaterOrEqual, len(csrList.Items),
		1, "There shall be at least one Certificate Signing Request"); err != nil {
		return fmt.Errorf("%s: %w", expectationDoesNotMeetMessage, err)
	}

	c.CertificateSigningRequestList = csrList.Items

	return nil
}

// certificateSigningRequestsShallHaveApprovalCondition ensures that each Certificate Signing Requests have
// approval condition.
// feature: certificatesigningrequest
func (c *ApproverInstance) certificateSigningRequestsShallHaveApprovalCondition() error {
	for _, csr := range c.CertificateSigningRequestList {
		return assertActual(assert.NotNil, csr.Status.Conditions)
	}

	return nil
}

// certificateSigningRequestsShallBeApproved ensures that each Certificate Signing Requests is approved.
// feature: certificatesigningrequest
func (c *ApproverInstance) certificateSigningRequestsShallBeApproved() error {
	for _, csr := range c.CertificateSigningRequestList {
		for _, condition := range csr.Status.Conditions {
			if err := assertExpectedAndActual(assert.Equal, corev1.ConditionTrue, condition.Status); err != nil {
				return fmt.Errorf("%s: %w", expectationDoesNotMeetMessage, err)
			}

			if err := assertExpectedAndActual(assert.Equal, certificatesv1.CertificateApproved, condition.Type); err != nil {
				return fmt.Errorf("%s: %w", expectationDoesNotMeetMessage, err)
			}
		}
	}

	return nil
}

// thereAreEventsRelatedToCertificateSigningRequests ensures that there is already existing event
// related to Certificate Signing Request.
// feature: eventrecorder
func (c *ApproverInstance) thereAreEventsRelatedToCertificateSigningRequests() error {
	events, errListEvents := c.Clientset.EventsV1().Events("default").List(context.TODO(), metav1.ListOptions{})
	if err := assertActual(assert.Nil, errListEvents); err != nil {
		return fmt.Errorf("can not list events: %w", err)
	}

	if err := assertExpectedAndActual(assert.GreaterOrEqual, len(events.Items), 1,
		"There shall be events"); err != nil {
		return fmt.Errorf("%s: %w", expectationDoesNotMeetMessage, err)
	}

	var actualEvents []eventsv1.Event

	for _, event := range events.Items {
		if event.Regarding.Kind == "CertificateSigningRequest" {
			actualEvents = append(actualEvents, event)
		}
	}

	if err := assertExpectedAndActual(assert.GreaterOrEqual, len(actualEvents), 1,
		"There shall be at least one event related to Certificate Signing Request"); err != nil {
		return fmt.Errorf("%s: %w", expectationDoesNotMeetMessage, err)
	}

	c.Events = actualEvents

	return nil
}

// approvalEventsShallContainReason ensures that each event shall have a specific reason.
// feature: eventrecorder
func (c *ApproverInstance) approvalEventsShallContainReason(reason string) error {
	for _, event := range c.Events {
		event := event

		if err := assertExpectedAndActual(assert.Contains, event.Reason, reason); err != nil {
			return fmt.Errorf("%s: %w", expectationDoesNotMeetMessage, err)
		}
	}

	return nil
}

// approvalEventsShallHaveMessage ensures that each each event shall have a specific message.
// feature: eventrecorder
func (c *ApproverInstance) approvalEventsShallHaveMessage(message string) error {
	for _, event := range c.Events {
		event := event

		if err := assertExpectedAndActual(assert.Contains, event.Note, message); err != nil {
			return fmt.Errorf("%s: %w", expectationDoesNotMeetMessage, err)
		}
	}

	return nil
}

// thereIsARunningPodInNamespaceWithLabel ensures that there is already running Pod within given namespace
// with given label.
// feature: metrics
func (c *ApproverInstance) thereIsARunningPodInNamespaceWithLabel(namespace, label string) error {
	podList, errList := c.Clientset.CoreV1().Pods(namespace).List(context.TODO(),
		metav1.ListOptions{
			LabelSelector: label,
		},
	)

	if err := assertActual(assert.Nil, errList); err != nil {
		return fmt.Errorf("can not list pods in %s namespace: %w", namespace, errList)
	}

	if err := assertExpectedAndActual(assert.GreaterOrEqual, len(podList.Items), 1); err != nil {
		return fmt.Errorf("%s: %w", expectationDoesNotMeetMessage, errList)
	}

	c.Pod = podList.Items[0]

	return nil
}

// thePodShallProvideEndpointAtPort ensures that there is existing endpoint at given port.
// feature: healthcheck, metrics
func (c *ApproverInstance) thePodShallProvideEndpointAtPort(endpoint string, port int) error {
	resp, errRequest := proxyRequestToPod(c.RestConfig, c.Pod.Namespace, c.Pod.Name, "http", endpoint, port)

	if err := assertActual(assert.Nil, errRequest); err != nil {
		return fmt.Errorf("can not list pods in %s namespace: %w", c.Pod.Namespace, err)
	}

	c.Response = resp

	return nil
}

// metricsShallContainMetric ensures that the given Prometheus Metric exist.
// feature: metrics
func (c *ApproverInstance) metricsShallContainMetric(metric string) error {
	if err := assertExpectedAndActual(assert.Contains, c.Metrics, metric); err != nil {
		return fmt.Errorf("%s: %w", expectationDoesNotMeetMessage, err)
	}

	return nil
}

// responseShallBeParseablePrometheusMetrics ensures that the response is parseable Prometheus Metrics.
// feature: metrics
func (c *ApproverInstance) responseShallBeParseablePrometheusMetrics() error {
	metrics, errParse := parseMetricNames(c.Response)
	if err := assertActual(assert.Nil, errParse); err != nil {
		return fmt.Errorf("can not parse Prometheus metrics: %w", errParse)
	}

	c.Metrics = metrics

	return nil
}

// responseShallContain ensures that the saved response contains given string.
// feature: healthcheck
func (c *ApproverInstance) responseShallContain(expected string) error {
	if err := assertActual(assert.NotNil, c.Response); err != nil {
		return fmt.Errorf("%s: %w", expectationDoesNotMeetMessage, err)
	}

	actualResponseText := string(c.Response)

	if err := assertExpectedAndActual(assert.GreaterOrEqual, len(actualResponseText), 1); err != nil {
		return fmt.Errorf("%s: %w", expectationDoesNotMeetMessage, err)
	}

	if err := assertExpectedAndActual(assert.Contains, actualResponseText, expected); err != nil {
		return fmt.Errorf("%s: %w", expectationDoesNotMeetMessage, err)
	}

	return nil
}

// execOption holds options for command execution.
type execOption struct {
	Stdin         io.Reader
	Namespace     string
	PodName       string
	ContainerName string
	Command       []string
	CaptureStdout bool
	CaptureStderr bool
}

// ExecuteCommandInContainer executes command in specified container and return stdout, stderr and error.
func (c *ApproverInstance) ExecuteCommandInContainer(options execOption,
	timeout time.Duration,
) (string, string, error) {
	const (
		container   = "container"
		tty         = false
		resource    = "pods"
		subResource = "exec"
	)

	request := c.Clientset.CoreV1().RESTClient().Post().
		Resource(resource).
		Name(options.PodName).
		Namespace(options.Namespace).
		SubResource(subResource).
		Param(container, options.ContainerName)

	request.VersionedParams(&corev1.PodExecOptions{
		Container: options.ContainerName,
		Command:   options.Command,
		Stdin:     options.Stdin != nil,
		Stdout:    options.CaptureStdout,
		Stderr:    options.CaptureStderr,
		TTY:       tty,
	}, scheme.ParameterCodec)

	var stdout, stderr bytes.Buffer

	ctx, cancel := context.WithTimeout(context.TODO(), timeout)

	defer cancel()

	err := execute(ctx, http.MethodPost, request.URL(), c.RestConfig, options.Stdin, &stdout, &stderr, tty)

	return strings.TrimSpace(stdout.String()), strings.TrimSpace(stderr.String()), err
}

// iExecuteCommandInTheRunningPod executes given command in existing Pod.
// feature: security
func (c *ApproverInstance) iExecuteCommandInTheRunningPod(command string) error {
	if err := assertExpectedLenAndActual(assert.Len, c.Pod.Spec.Containers, 1); err != nil {
		return fmt.Errorf("%s: %w", expectationDoesNotMeetMessage, err)
	}

	stdout, stderr, err := c.ExecuteCommandInContainer(execOption{
		Command:       []string{command},
		Namespace:     c.Pod.Namespace,
		PodName:       c.Pod.Name,
		ContainerName: c.Pod.Spec.Containers[0].Name,
		Stdin:         nil,
		CaptureStdout: true,
		CaptureStderr: true,
	}, 30*time.Second)

	c.CommandResult = commandResult{
		StdOut: stdout,
		StdErr: stderr,
		Error:  err,
	}

	return nil
}

// commandExecutionShallReportError ensures that executed command reported error.
// feature: security
func (c *ApproverInstance) commandExecutionShallReportError() error {
	if err := assertActual(assert.NotNil, c.CommandResult.Error); err != nil {
		return fmt.Errorf("%s: %w", expectationDoesNotMeetMessage, err)
	}

	return nil
}

// commandExecutionShallNotReportAnyError ensures that executed command did not report any error.
// feature: security
func (c *ApproverInstance) commandExecutionShallNotReportAnyError() error {
	if err := assertActual(assert.Nil, c.CommandResult.Error); err != nil {
		return fmt.Errorf("%s: %w", expectationDoesNotMeetMessage, err)
	}

	return nil
}

// commandExecutionErrorMessageShallContain ensures that executed command error message contains expected string.
// feature: security
func (c *ApproverInstance) commandExecutionErrorMessageShallContain(expected *godog.DocString) error {
	if err := assertExpectedAndActual(assert.Contains, c.CommandResult.Error.Error(), expected.Content); err != nil {
		return fmt.Errorf("%s: %w", expectationDoesNotMeetMessage, err)
	}

	return nil
}

func parseMetricNames(data []byte) ([]string, error) {
	parser := expfmt.TextParser{}

	mfs, err := parser.TextToMetricFamilies(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("can not parse Prometheus Metrics: %w", err)
	}

	ms := make([]string, 0, len(mfs))

	for key := range mfs {
		ms = append(ms, key)
	}

	return ms, nil
}

func proxyRequestToPod(config *rest.Config, namespace, podname, scheme, path string,
	port int,
) ([]byte, error) {
	cancel, err := setupForwarding(config, namespace, port, podname)
	if err != nil {
		return nil, fmt.Errorf("can not setup port forwarding: %w", err)
	}

	defer cancel()

	var query string

	if strings.Contains(path, "?") {
		elm := strings.SplitN(path, "?", 2)
		path = elm[0]
		query = elm[1]
	}

	reqURL := url.URL{
		Scheme:   scheme,
		Path:     path,
		RawQuery: query,
		Host:     fmt.Sprintf("127.0.0.1:%d", port),
	}

	resp, err := sendRequest(config, reqURL.String())
	if err != nil {
		return nil, fmt.Errorf("can not send request: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("can not read response: %w", err)
	}

	return body, nil
}

func setupForwarding(config *rest.Config, namespace string, port int,
	podname string,
) (cancel func(), err error) {
	hostIP := strings.TrimPrefix(config.Host, "https://")

	trans, upgrader, err := spdy.RoundTripperFor(config)
	if err != nil {
		return noop, fmt.Errorf("can not configure RoundTripper: %w", err)
	}

	dialer := spdy.NewDialer(
		upgrader,
		&http.Client{
			Transport: trans,
		},
		http.MethodPost,
		&url.URL{
			Scheme: "https",
			Path:   fmt.Sprintf("/api/v1/namespaces/%s/pods/%s/portforward", namespace, podname),
			Host:   hostIP,
		},
	)

	var berr, bout bytes.Buffer

	buffErr := bufio.NewWriter(&berr)
	buffOut := bufio.NewWriter(&bout)

	// stopCh controls the port forwarding lifecycle. When it gets closed the port forward will terminate.
	stopCh := make(chan struct{}, 1)
	// readyCh communicate when the port forward is ready to get traffic.
	readyCh := make(chan struct{}, 1)

	portForwarder, err := portforward.New(dialer, []string{strconv.Itoa(port)}, stopCh, readyCh, buffOut, buffErr)
	if err != nil {
		return noop, fmt.Errorf("can not create new portforwarding: %w", err)
	}

	go func() {
		if err := portForwarder.ForwardPorts(); err != nil {
			panic(err)
		}
	}()
	<-readyCh

	return func() {
		stopCh <- struct{}{}
	}, nil
}

func noop() {
}

func sendRequest(config *rest.Config, url string) (*http.Response, error) {
	tsConfig, err := config.TransportConfig()
	if err != nil {
		return nil, fmt.Errorf("can not configure transport %w", err)
	}

	ts, err := transport.New(tsConfig)
	if err != nil {
		return nil, fmt.Errorf("can not create new transport %w", err)
	}

	client := &http.Client{Transport: ts}

	request, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("can not create new transport: %w", err)
	}

	return client.Do(request)
}

// execute executes command in given container of a Pod.
func execute(ctx context.Context, method string, url *url.URL, config *rest.Config, stdin io.Reader,
	stdout, stderr io.Writer, tty bool,
) error {
	exec, err := remotecommand.NewSPDYExecutor(config, method, url)
	if err != nil {
		return err
	}

	return exec.StreamWithContext(ctx, remotecommand.StreamOptions{
		Stdin:  stdin,
		Stdout: stdout,
		Stderr: stderr,
		Tty:    tty,
	})
}

// createNewClientSet creates a client to be used to communicate with Kubernetes API Server.
func createNewClientSet() (*clientgokubernetes.Clientset, *rest.Config, error) {
	config, errConfig := clientcmd.NewDefaultClientConfigLoadingRules().Load()
	if errConfig != nil {
		return nil, nil, fmt.Errorf("can not create new default client loading rules: %w", errConfig)
	}

	restConfig, errRestConfig := clientcmd.NewDefaultClientConfig(*config, &clientcmd.ConfigOverrides{}).ClientConfig()
	if errRestConfig != nil {
		return nil, nil, fmt.Errorf("can not create new default client configuration: %w", errRestConfig)
	}

	client, errNewClientset := clientgokubernetes.NewForConfig(restConfig)
	if errNewClientset != nil {
		return nil, nil, fmt.Errorf("can not create clientset for the given config: %w", errNewClientset)
	}

	return client, restConfig, nil
}
