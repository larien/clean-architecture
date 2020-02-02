package article

import (
	"errors"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/mock"
	assert "github.com/stretchr/testify/require"
)

func TestController_NewController(t *testing.T) {
	t.Parallel()
	t.Run("when Controller is created", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		r := new(MockRepository)

		controller := NewController(r)
		is.NotNil(controller)
	})
}

func TestController_Create(t *testing.T) {
	t.Parallel()
	t.Run("when repository fails", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		r := new(MockRepository)
		defer r.AssertExpectations(t)

		r.On("Create", mock.Anything).
			Return(errors.New("error")).
			Once()

		c := NewController(r)

		article := &Article{}
		is.Nil(faker.FakeData(article))
		is.NotNil(c.Create(article))
	})
	t.Run("when repository succeeds", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		r := new(MockRepository)
		defer r.AssertExpectations(t)

		r.On("Create", mock.Anything).
			Return(nil).
			Once()

		c := NewController(r)

		article := &Article{}
		is.Nil(faker.FakeData(article))
		is.Nil(c.Create(article))
	})
}

func TestController_List(t *testing.T) {
	t.Parallel()
	t.Run("when repository fails", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		r := new(MockRepository)
		defer r.AssertExpectations(t)

		r.On("List", mock.Anything).
			Return(nil, errors.New("error")).
			Once()

		c := NewController(r)

		articles, err := c.List()
		is.NotNil(err)
		is.Nil(articles)
	})
	t.Run("when repository succeeds", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)

		r := new(MockRepository)
		defer r.AssertExpectations(t)

		article := &Article{
			Title:   faker.Sentence(),
			Content: faker.Paragraph(),
			Author:  faker.Name(),
		}

		var articles []*Article
		articles = append(articles, article)

		r.On("List", mock.Anything).
			Return(articles, nil).
			Once()

		c := NewController(r)

		result, err := c.List()
		is.Nil(err)
		is.Equal(articles, result)
	})
}
