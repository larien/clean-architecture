package post

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/mock"
	assert "github.com/stretchr/testify/require"
)

func TestPost_NewRouter(t *testing.T) {
	t.Parallel()
	t.Run("when Router is created", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		c := new(MockController)

		router := NewRoutes(c)
		is.NotNil(router)
		is.NotEqual(0, len(router.Routes()))
	})
}

func TestPost_create(t *testing.T) {
	t.Parallel()
	t.Run("when JSON is invalid", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		c := new(MockController)

		invalidJSON := `{ invalid_json:"" }`
		req := httptest.NewRequest(
			http.MethodPost,
			"/posts",
			bytes.NewBufferString(invalidJSON),
		)
		rec := httptest.NewRecorder()

		handler := http.HandlerFunc(create(c))
		handler.ServeHTTP(rec, req)

		is.Equal(http.StatusBadRequest, rec.Code)
	})
	t.Run("when JSON is valid", func(t *testing.T) {
		t.Parallel()
		t.Run("and controller failed", func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			c := new(MockController)
			defer c.AssertExpectations(t)

			c.On("Create", mock.Anything, mock.Anything).
				Return(errors.New("error")).
				Once()

			post := &Post{}
			err := faker.FakeData(post)
			is.Nil(err)

			body, _ := json.Marshal(post)
			req := httptest.NewRequest(http.MethodPost, "/posts", bytes.NewBuffer(body))
			rec := httptest.NewRecorder()

			handler := http.HandlerFunc(create(c))
			handler.ServeHTTP(rec, req)

			is.Equal(http.StatusInternalServerError, rec.Code)
		})
		t.Run("and controller succeeded", func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			c := new(MockController)
			defer c.AssertExpectations(t)

			c.On("Create", mock.Anything, mock.Anything).
				Return(nil).
				Once()

			post := &Post{}
			err := faker.FakeData(post)
			is.Nil(err)

			body, _ := json.Marshal(post)
			req := httptest.NewRequest(http.MethodPost, "/posts", bytes.NewBuffer(body))
			rec := httptest.NewRecorder()

			handler := http.HandlerFunc(create(c))
			handler.ServeHTTP(rec, req)

			is.Equal(http.StatusOK, rec.Code)
		})
	})
}

// invalid JSON

// valid json
