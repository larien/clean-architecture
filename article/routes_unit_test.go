package article

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bxcodec/faker/v3"
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
			is := assert.New(t)

			c := new(MockController)
			defer c.AssertExpectations(t)

			c.On("List", mock.Anything).
				Return(nil, nil).
				Once()

			req := httptest.NewRequest(http.MethodGet, "/articles", nil)
			rec := httptest.NewRecorder()

			handler := http.HandlerFunc(list(c))
			handler.ServeHTTP(rec, req)

			is.Equal(http.StatusNotFound, rec.Code)
		})
		t.Run("and articles were found", func(t *testing.T) {
			t.Parallel()
			is := assert.New(t)

			c := new(MockController)
			defer c.AssertExpectations(t)

			article := Article{
				Title:   faker.Sentence(),
				Content: faker.Paragraph(),
				Author:  faker.Name(),
			}

			var articles []Article
			articles = append(articles, article)

			body, _ := json.Marshal(articles)

			c.On("List", mock.Anything).
				Return(&articles, nil).
				Once()

			req := httptest.NewRequest(http.MethodGet, "/articles", bytes.NewBuffer(body))
			rec := httptest.NewRecorder()

			handler := http.HandlerFunc(list(c))
			handler.ServeHTTP(rec, req)

			is.Equal(http.StatusOK, rec.Code)
			var result []Article
			is.Nil(json.Unmarshal(rec.Body.Bytes(), &result))
			is.Equal(articles, result)
		})
	})
}

func TestRoutes_detail(t *testing.T) {
	t.Parallel()
	t.Run("when ID is invalid", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		c := new(MockController)
		defer c.AssertExpectations(t)

		id := faker.UUIDDigit()
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/articles/%s/detail", id), nil)
		rec := httptest.NewRecorder()

		handler := http.HandlerFunc(detail(c))
		handler.ServeHTTP(rec, req)

		is.Equal(http.StatusBadRequest, rec.Code)
	})
	t.Run("when controller fails", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		c := new(MockController)
		defer c.AssertExpectations(t)

		c.On("Detail", mock.Anything).
			Return(nil, errors.New("error")).
			Once()

		id := faker.UnixTime()
		fmt.Println(id)
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/articles/%d/detail", id), nil)
		rec := httptest.NewRecorder()

		handler := http.HandlerFunc(detail(c))
		handler.ServeHTTP(rec, req)

		is.Equal(http.StatusInternalServerError, rec.Code)
	})
	t.Run("when controller succeeds", func(t *testing.T) {
		t.Parallel()
		t.Run("and the article isn't found", func(t *testing.T) {
			t.Parallel()
		})
		t.Run("and the article is found", func(t *testing.T) {
			t.Parallel()
		})
	})
}
