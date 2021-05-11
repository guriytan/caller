package caller

import (
	"io"
)

type Result interface {
	Err() error
	Byte() ([]byte, error)
	String() (string, error)
	Raw() (io.ReadCloser, error)
	Parse(receive interface{}) error
	ParseWithFunc(receive interface{}, parse ParseFunc) error
}

type result struct {
	body      io.ReadCloser
	bytes     []byte
	parseFunc ParseFunc

	err   error
	close bool
}

func newErrResult(err error) Result {
	return &result{err: err}
}

func newResult(body io.ReadCloser, parse ParseFunc) Result {
	return &result{body: body, parseFunc: parse}
}

func (p *result) Err() error {
	return p.err
}

func (p *result) Byte() ([]byte, error) {
	return p.read()
}

func (p *result) String() (string, error) {
	read, err := p.read()
	if err != nil {
		return "", err
	}
	return bytesToString(read), err
}

func (p *result) Raw() (io.ReadCloser, error) {
	if err := p.check(); err != nil {
		return nil, err
	}

	return p.body, nil
}

func (p *result) Parse(receive interface{}) error {
	return p.ParseWithFunc(receive, p.parseFunc)
}

func (p *result) ParseWithFunc(receive interface{}, parse ParseFunc) error {
	read, err := p.read()
	if err != nil {
		return err
	}

	if err = parse(read, receive); err != nil {
		return newResultError("parse response body failed", err)
	}
	return nil
}

func (p *result) read() ([]byte, error) {
	if p.bytes != nil {
		return p.bytes, nil
	}

	if err := p.check(); err != nil {
		return nil, err
	}
	defer func() { _ = p.body.Close(); p.close = true }()

	var err error
	p.bytes, err = readerToBytes(p.body)
	if err != nil {
		return nil, newResultError("read response body failed", err)
	}

	return p.bytes, nil
}

func (p *result) check() error {
	if p.err != nil {
		return p.err
	}
	if p.close {
		return ErrCloseBody
	}
	return nil
}
