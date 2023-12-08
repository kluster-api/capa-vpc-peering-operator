// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package errors

import (
	"fmt"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
)

const (
	levelError = "error"
)

type tfError struct {
	message string
}

type applyFailed struct {
	*tfError
}

// TerraformLog represents relevant fields of a Terraform CLI JSON-formatted log line
type TerraformLog struct {
	Level      string        `json:"@level"`
	Message    string        `json:"@message"`
	Diagnostic LogDiagnostic `json:"diagnostic"`
}

// LogDiagnostic represents relevant fields of a Terraform CLI JSON-formatted
// log line diagnostic info
type LogDiagnostic struct {
	Severity string `json:"severity"`
	Summary  string `json:"summary"`
	Detail   string `json:"detail"`
}

func (t *tfError) Error() string {
	return t.message
}

func newTFError(message string, logs []byte) (string, *tfError) {
	tfError := &tfError{
		message: message,
	}

	tfLogs, err := parseTerraformLogs(logs)
	if err != nil {
		return err.Error(), tfError
	}

	messages := make([]string, 0, len(tfLogs))
	for _, l := range tfLogs {
		// only use error logs
		if l == nil || l.Level != levelError {
			continue
		}
		m := l.Message
		if l.Diagnostic.Severity == levelError && l.Diagnostic.Summary != "" {
			m = fmt.Sprintf("%s: %s", l.Diagnostic.Summary, l.Diagnostic.Detail)
		}
		messages = append(messages, m)
	}
	tfError.message = fmt.Sprintf("%s: %s", message, strings.Join(messages, "\n"))
	return "", tfError
}

func parseTerraformLogs(logs []byte) ([]*TerraformLog, error) {
	logLines := strings.Split(string(logs), "\n")
	tfLogs := make([]*TerraformLog, 0, len(logLines))
	for _, l := range logLines {
		log := &TerraformLog{}
		l := strings.TrimSpace(l)
		if l == "" {
			continue
		}
		if err := jsoniter.ConfigCompatibleWithStandardLibrary.UnmarshalFromString(l, log); err != nil {
			return nil, err
		}
		tfLogs = append(tfLogs, log)
	}
	return tfLogs, nil
}

// NewApplyFailed returns a new apply failure error with given logs.
func NewApplyFailed(logs []byte) error {
	parseError, tfError := newTFError("apply failed", logs)
	result := &applyFailed{tfError: tfError}
	if parseError == "" {
		return result
	}
	return errors.WithMessage(result, parseError)
}

// IsApplyFailed returns whether error is due to failure of an apply operation.
func IsApplyFailed(err error) bool {
	r := &applyFailed{}
	return errors.As(err, &r)
}

type destroyFailed struct {
	*tfError
}

// NewDestroyFailed returns a new destroy failure error with given logs.
func NewDestroyFailed(logs []byte) error {
	parseError, tfError := newTFError("destroy failed", logs)
	result := &destroyFailed{tfError: tfError}
	if parseError == "" {
		return result
	}
	return errors.WithMessage(result, parseError)
}

// IsDestroyFailed returns whether error is due to failure of a destroy operation.
func IsDestroyFailed(err error) bool {
	r := &destroyFailed{}
	return errors.As(err, &r)
}

type refreshFailed struct {
	*tfError
}

// NewRefreshFailed returns a new destroy failure error with given logs.
func NewRefreshFailed(logs []byte) error {
	parseError, tfError := newTFError("refresh failed", logs)
	result := &refreshFailed{tfError: tfError}
	if parseError == "" {
		return result
	}
	return errors.WithMessage(result, parseError)
}

// IsRefreshFailed returns whether error is due to failure of a destroy operation.
func IsRefreshFailed(err error) bool {
	r := &refreshFailed{}
	return errors.As(err, &r)
}

type planFailed struct {
	*tfError
}

// NewPlanFailed returns a new destroy failure error with given logs.
func NewPlanFailed(logs []byte) error {
	parseError, tfError := newTFError("plan failed", logs)
	result := &planFailed{tfError: tfError}
	if parseError == "" {
		return result
	}
	return errors.WithMessage(result, parseError)
}

// IsPlanFailed returns whether error is due to failure of a destroy operation.
func IsPlanFailed(err error) bool {
	r := &planFailed{}
	return errors.As(err, &r)
}

type retrySchedule struct {
	invocationCount int
	ttl             int
}

func NewRetryScheduleError(invocationCount, ttl int) error {
	return &retrySchedule{
		invocationCount: invocationCount,
		ttl:             ttl,
	}
}

func (r *retrySchedule) Error() string {
	return fmt.Sprintf("native provider reuse budget has been exceeded: invocationCount: %d, ttl: %d", r.invocationCount, r.ttl)
}

// IsRetryScheduleError returns whether the error is a retry error
// for the scheduler.
func IsRetryScheduleError(err error) bool {
	r := &retrySchedule{}
	return errors.As(err, &r)
}

type asyncCreateFailed struct {
	error
}

// NewAsyncCreateFailed returns a new async crate failure.
func NewAsyncCreateFailed(err error) error {
	if err == nil {
		return nil
	}
	return &asyncCreateFailed{
		error: errors.Wrap(err, "async create failed"),
	}
}

// IsAsyncCreateFailed returns whether error is due to failure of
// an async create operation.
func IsAsyncCreateFailed(err error) bool {
	r := &asyncCreateFailed{}
	return errors.As(err, &r)
}

type asyncUpdateFailed struct {
	error
}

// NewAsyncUpdateFailed returns a new async update failure.
func NewAsyncUpdateFailed(err error) error {
	if err == nil {
		return nil
	}
	return &asyncUpdateFailed{
		error: errors.Wrap(err, "async update failed"),
	}
}

// IsAsyncUpdateFailed returns whether error is due to failure of
// an async update operation.
func IsAsyncUpdateFailed(err error) bool {
	r := &asyncUpdateFailed{}
	return errors.As(err, &r)
}

type asyncDeleteFailed struct {
	error
}

// NewAsyncDeleteFailed returns a new async delete failure.
func NewAsyncDeleteFailed(err error) error {
	if err == nil {
		return nil
	}
	return &asyncDeleteFailed{
		error: errors.Wrap(err, "async delete failed"),
	}
}

// IsAsyncDeleteFailed returns whether error is due to failure of
// an async delete operation.
func IsAsyncDeleteFailed(err error) bool {
	r := &asyncDeleteFailed{}
	return errors.As(err, &r)
}
