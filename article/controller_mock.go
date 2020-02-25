package article

import (
	"github.com/stretchr/testify/mock"
)

// MockController injects mock dependency into controller
type MockController struct {
	mock.Mock
}

// Create represents the mocked method for Create feature in Controller layer
func (m *MockController) Create(article *Article) error {
	args := m.Called(article)
	return args.Error(0)
}

// List represents the mocked method for List feature in Controller layer
func (m *MockController) List() (*[]Article, error) {
	args := m.Called(nil)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*[]Article), nil
}

// Detail represents the mocked method for Detail feature in Controller layer
func (m *MockController) Detail(id uint) (*Article, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Article), nil
}
