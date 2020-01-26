package article

import (
	"github.com/stretchr/testify/mock"
)

// MockController injects mock dependency into controller
type MockController struct {
	mock.Mock
}

// create represents the mocked method for Create feature in Controller layer
func (m *MockController) Create(article *Article) error {
	args := m.Called(article)
	return args.Error(0)
}
