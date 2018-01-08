package builder

import (
	"bytes"
	"fmt"
	"io"

	"github.com/pkg/errors"
)

// Writer defines the interface
type Writer interface {
	io.Writer
	Append(...interface{})
}

var _ Writer = NewWriter()

// BytesWriter implments Writer and save expression in bytes.Buffer
type BytesWriter struct {
	writer *bytes.Buffer
	buffer []byte
	args   []interface{}
}

// NewWriter creates a new string writer
func NewWriter() *BytesWriter {
	w := &BytesWriter{}
	w.writer = bytes.NewBuffer(w.buffer)
	return w
}

// Write writes data to Writer
func (s *BytesWriter) Write(buf []byte) (int, error) {
	n, err := s.writer.Write(buf)
	return n, errors.WithStack(err)
}

// Append appends args to Writer
func (s *BytesWriter) Append(args ...interface{}) {
	s.args = append(s.args, args...)
}

// Cond defines an interface
type Cond interface {
	WriteTo(Writer) error
	And(...Cond) Cond
	Or(...Cond) Cond
	IsValid() bool
}

type CondTopic struct {
	Topic string
}

var _ Cond = CondTopic{}

// NewCond creates an empty condition
func NewCond() Cond {
	return CondTopic{}
}

func (c CondTopic) inTopicFormat() string {
	return fmt.Sprintf("'%s' in topics", c.Topic)
}

func (c CondTopic) WriteTo(w Writer) error {
	_, err := fmt.Fprintf(w, c.inTopicFormat())
	return errors.WithStack(err)
}

func (c CondTopic) And(conds ...Cond) Cond {
	return And(c, And(conds...))
}

func (c CondTopic) Or(conds ...Cond) Cond {
	return Or(c, Or(conds...))
}

func (c CondTopic) IsValid() bool {
	return c.Topic != ""
}

// ToCondition convert a builder or cond to condition string and args
func ToCondition(cond Cond) (string, []interface{}, error) {
	return condToExpr(cond.(Cond))
}

func condToExpr(cond Cond) (string, []interface{}, error) {
	if cond == nil || !cond.IsValid() {
		return "", nil, nil
	}

	w := NewWriter()
	if err := cond.WriteTo(w); err != nil {
		return "", nil, err
	}
	return w.writer.String(), w.args, nil
}
