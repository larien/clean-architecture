package article

import (
	"github.com/stretchr/testify/mock"
)

// MockRepository injects mock dependency into repository
type MockRepository struct {
	mock.Mock
}

// Create represents the mocked method for Create feature in Repository layer
func (m *MockRepository) Create(article *Article) error {
	args := m.Called(article)
	return args.Error(0)
}

// List represents the mocked method for List feature in Repository layer
func (m *MockRepository) List() (*[]Article, error) {
	args := m.Called(nil)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*[]Article), nil
}
