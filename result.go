package caller

import (
	"bytes"
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

	err error
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
	return p.readByte()
}

func (p *result) String() (string, error) {
	read, err := p.readByte()
	if err != nil {
		return "", err
	}
	return bytesToString(read), err
}

func (p *result) Raw() (io.ReadCloser, error) {
	if len(p.bytes) != 0 {
		return io.NopCloser(bytes.NewReader(p.bytes)), nil
	}

	if p.Err() != nil {
		return nil, p.Err()
	}

	return p.body, nil
}

func (p *result) Parse(receive interface{}) error {
	return p.ParseWithFunc(receive, p.parseFunc)
}

func (p *result) ParseWithFunc(receive interface{}, parse ParseFunc) error {
	read, err := p.readByte()
	if err != nil {
		return err
	}

	if err = parse(read, receive); err != nil {
		return newResultError("parse response body failed", err)
	}
	return nil
}

func (p *result) readByte() ([]byte, error) {
	if len(p.bytes) != 0 {
		return p.bytes, nil
	}

	if p.Err() != nil {
		return nil, p.Err()
	}
	defer func() { _ = p.body.Close() }()

	var err error
	p.bytes, err = readerToBytes(p.body)
	if err != nil {
		return nil, newResultError("read response body failed", err)
	}

	return p.bytes, nil
}
