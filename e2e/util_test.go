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

package e2e_test

import (
	"fmt"

	"github.com/stretchr/testify/assert"
)

// assertExpectedAndActual is a helper function to allow the step function to call
// assertion functions where you want to compare an expected and an actual value.
func assertExpectedAndActual(a expectedAndActualAssertion,
	expected, actual interface{}, msgAndArgs ...interface{},
) error {
	var t asserter

	a(&t, expected, actual, msgAndArgs...)

	return t.err
}

type expectedAndActualAssertion func(t assert.TestingT,
	expected, actual interface{}, msgAndArgs ...interface{}) bool

// assertActual is a helper function to allow the step function to call
// assertion functions where you want to compare an actual value to a
// predefined state like nil, empty or true/false.
func assertActual(a actualAssertion, actual interface{}) error {
	var t asserter

	a(&t, actual)

	return t.err
}

// assertExpectedLenAndActual is a helper function to allow the step function to call
// assertion functions where you want to compare an actual value to a
// has specific length.
func assertExpectedLenAndActual(a func(t assert.TestingT, object interface{},
	length int, msgAndArgs ...interface{}) bool, actual interface{}, length int, msgAndArgs ...interface{},
) error {
	var t asserter

	a(&t, actual, length, msgAndArgs)

	return t.err
}

type actualAssertion func(t assert.TestingT, actual interface{}, msgAndArgs ...interface{}) bool

// asserter is used to be able to retrieve the error reported by the called assertion.
type asserter struct {
	err error
}

// Errorf is used by the called assertion to report an error.
//
//nolint:err113
func (a *asserter) Errorf(format string, args ...interface{}) {
	a.err = fmt.Errorf(format, args...)
}
