package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorFormat(t *testing.T) {
	e := New(CodeNotFound, "record missing")
	assert.Contains(t, e.Error(), "record missing")
	assert.Equal(t, CodeNotFound, e.Code)
}

func TestErrorUnwrap(t *testing.T) {
	inner := errors.New("boom")
	e := Wrap(CodeInternal, "outer", inner)
	assert.ErrorIs(t, e, inner)
}
