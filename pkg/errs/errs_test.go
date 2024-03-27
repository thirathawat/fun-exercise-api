package errs_test

import (
	"testing"

	"github.com/KKGo-Software-engineering/fun-exercise-api/pkg/errs"
	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	err := errs.New("this is error message")
	assert.Equal(t, err.Error(), "this is error message")
}
