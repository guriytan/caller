package caller

import (
	"errors"
	"fmt"
)

var (
	ErrCloseBody = errors.New("body has close")
	ErrRequest   = errors.New("request is wrong")
	ErrServer    = errors.New("server is wrong")
)

func UnWarpHttpError(err error) (*HttpError, error) {
	var callError *HttpError
	if errors.As(err, &callError) {
		return callError, nil
	}
	return nil, errors.New("not match error type")
}

type NoRetryError struct {
	error
}

func newNoRetryError(err error) *NoRetryError {
	return &NoRetryError{error: err}
}

type HttpError struct {
	code    int
	message string
	err     error
}

func newHttpError(code int, message string, err error) *HttpError {
	return &HttpError{code: code, message: message, err: err}
}

func (h *HttpError) Error() string {
	return fmt.Sprintf("call http failed, status code: %d, msg: %s, err: %v", h.code, h.message, h.err)
}

func (h *HttpError) StatusCode() int {
	return h.code
}

func (h *HttpError) Message() string {
	return h.message
}

type RetryError struct {
	retry int
	err   error
}

func newRetryError(retry int, err error) *RetryError {
	return &RetryError{retry: retry, err: err}
}

func (e *RetryError) Error() string {
	return fmt.Sprintf("retry failed, time: %d, err: %v", e.retry, e.err)
}

type ResultError struct {
	message string
	err     error
}

func newResultError(message string, err error) *ResultError {
	return &ResultError{message: message, err: err}
}

func (e *ResultError) Error() string {
	return fmt.Sprintf("get result failed, msg: %v, err: %v", e.message, e.err)
}
