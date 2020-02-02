package article

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

func TestRoutes_NewRouter(t *testing.T) {
	t.Parallel()
	t.Run("when Routes is created", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		c := new(MockController)

		router := NewRoutes(c)
		is.NotNil(router)
		is.NotEqual(0, len(router.Routes()))
	})
}

func TestRoutes_create(t *testing.T) {
	t.Parallel()
	t.Run("when JSON is invalid", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		c := new(MockController)

		invalidJSON := `{ invalid_json:"" }`
		req := httptest.NewRequest(
			http.MethodPost,
			"/articles",
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

			article := &Article{}
			err := faker.FakeData(article)
			is.Nil(err)

			body, _ := json.Marshal(article)
			req := httptest.NewRequest(http.MethodPost, "/articles", bytes.NewBuffer(body))
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

			article := &Article{}
			err := faker.FakeData(article)
			is.Nil(err)

			body, _ := json.Marshal(article)
			req := httptest.NewRequest(http.MethodPost, "/articles", bytes.NewBuffer(body))
			rec := httptest.NewRecorder()

			handler := http.HandlerFunc(create(c))
			handler.ServeHTTP(rec, req)

			is.Equal(http.StatusOK, rec.Code)
		})
	})
}

func TestRoutes_list(t *testing.T) {
	t.Parallel()
	t.Run("when controller fails", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		c := new(MockController)
		defer c.AssertExpectations(t)

		c.On("List", mock.Anything).
			Return(nil, errors.New("error")).
			Once()

		req := httptest.NewRequest(http.MethodGet, "/articles", nil)
		rec := httptest.NewRecorder()

		handler := http.HandlerFunc(list(c))
		handler.ServeHTTP(rec, req)

		is.Equal(http.StatusInternalServerError, rec.Code)
	})
	t.Run("when controller succeeds", func(t *testing.T) {
		t.Parallel()
		t.Run("and no article was found", func(t *testing.T) {
			t.Parallel()
		})
		t.Run("and articles were found", func(t *testing.T) {
			t.Parallel()
		})
	})

}
