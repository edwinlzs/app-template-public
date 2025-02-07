package utils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONResponse(t *testing.T) {
	t.Run("successful response", func(t *testing.T) {
		w := httptest.NewRecorder()
		data := map[string]string{"key": "value"}

		JSONResponse(w, data)

		assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
		assert.Equal(t, http.StatusOK, w.Code)

		var result map[string]string
		err := json.NewDecoder(w.Body).Decode(&result)
		assert.NoError(t, err)
		assert.Equal(t, data, result)
	})

	t.Run("handles non-encodable data", func(t *testing.T) {
		w := httptest.NewRecorder()
		data := make(chan int) // channels cannot be JSON encoded

		JSONResponse(w, data)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
