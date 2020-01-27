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
