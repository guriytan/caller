package caller

import (
	"errors"
	"fmt"
)

var (
	ErrCloseBody = errors.New("body has close")
	ErrRequest   = errors.New("request is wrong")
	ErrServer    = errors.New("server is wrong")
	ErrNotMatch  = errors.New("not match error type")
)

func UnWarpError(err error) (*ResultError, error) {
	var callError *ResultError
	if errors.As(err, &callError) {
		return callError, nil
	}
	return nil, ErrNotMatch
}

type NoRetryError struct {
	error
}

func newNoRetryError(err error) *NoRetryError {
	return &NoRetryError{error: err}
}

func (n *NoRetryError) Unwrap() error {
	return n.error
}

type ResultError struct {
	code    int
	message string
	err     error
}

func newHttpError(code int, message string, err error) *ResultError {
	return &ResultError{code: code, message: message, err: err}
}

func newResultError(message string, err error) *ResultError {
	return &ResultError{message: message, err: err}
}

func (h *ResultError) Error() string {
	return fmt.Sprintf("call http failed, status code: %d, msg: %s, err: %v", h.code, h.message, h.err)
}

func (h *ResultError) StatusCode() int {
	return h.code
}

func (h *ResultError) Message() string {
	return h.message
}

func (h *ResultError) Unwrap() error {
	return h.err
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

func (e *RetryError) Unwrap() error {
	return e.err
}
