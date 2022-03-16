package underscore_test

import (
	"errors"
	"testing"

	u "github.com/rjNemo/underscore"
	"github.com/stretchr/testify/assert"
)

func TestSuccess(t *testing.T) {
	res := isAnswerToLife(42)
	assert.True(t, res.IsSuccess())
}

func TestFailure(t *testing.T) {
	res := isAnswerToLife(13)
	assert.False(t, res.IsSuccess())
}

func TestIsOK(t *testing.T) {
	res, err := isAnswerToLife(42).ToValue()
	assert.NoError(t, err)
	assert.Equal(t, "You get it", *res)
}

func TestIsError(t *testing.T) {
	life := isAnswerToLife(13)
	res, err := life.ToValue()
	assert.Error(t, err)
	assert.Equal(t, "nope", life.(u.Err[string]).Error())
	assert.Nil(t, res)
}

func isAnswerToLife(num int) u.Result[string] {
	if num == 42 {
		res := "You get it"
		return u.ToResult(&res, nil)
	}
	return u.ToResult[string](nil, errors.New("nope")) // u.Err[string]{Err: errors.New("nope")}
}
