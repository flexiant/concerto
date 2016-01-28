package utils

import (
	"github.com/stretchr/testify/mock"
)

// MockConcertoService  service manager.
type MockConcertoService struct {
	mock.Mock
}

// Post mocks POST request to Concerto API
func (m *MockConcertoService) Post(path string, payload *map[string]string) ([]byte, int, error) {
	args := m.Called(path, payload)
	return args.Get(0).([]byte), args.Int(1), args.Error(2)
}

// Put mocks PUT request to Concerto API
func (m *MockConcertoService) Put(path string, payload *map[string]string) ([]byte, int, error) {
	args := m.Called(path, payload)
	return args.Get(0).([]byte), args.Int(1), args.Error(2)
}

// Delete mocks DELETE request to Concerto API
func (m *MockConcertoService) Delete(path string) ([]byte, int, error) {
	args := m.Called(path)
	return args.Get(0).([]byte), args.Int(1), args.Error(2)
}

// Get mocks GET request to Concerto API
func (m *MockConcertoService) Get(path string) ([]byte, int, error) {
	args := m.Called(path)
	return args.Get(0).([]byte), args.Int(1), args.Error(2)
}

// GetFile sends GET request to Concerto API and receives a file
func (m *MockConcertoService) GetFile(path string, directoryPath string) (string, int, error) {
	args := m.Called(path, directoryPath)
	return args.String(0), args.Int(1), args.Error(2)
}
