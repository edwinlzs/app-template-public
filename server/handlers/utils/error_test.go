package utils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatusError(t *testing.T) {
	t.Run("implements Error interface", func(t *testing.T) {
		err := errors.New("test error")
		statusErr := StatusError{Code: 400, Err: err}

		assert.Equal(t, err.Error(), statusErr.Error())
		assert.Equal(t, 400, statusErr.Status())
	})
}
