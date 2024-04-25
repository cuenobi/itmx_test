package delivery

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"itmx_test/domain"
	"itmx_test/service/entity"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCustomerService struct {
	mock.Mock
}

func (m *MockCustomerService) CreateCustomer(customer *entity.Customer) error {
	args := m.Called(customer)
	return args.Error(0)
}

func (m *MockCustomerService) GetCustomerByID(id string) (*entity.Customer, error) {
	args := m.Called(id)
	return &entity.Customer{}, args.Error(1)
}

func (m *MockCustomerService) UpdateCustomerByID(customer *entity.Customer, id string) error {
	args := m.Called(customer, id)
	return args.Error(0)
}

func (m *MockCustomerService) DelCustomerByID(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreateCustomerHandler(t *testing.T) {
	mockService := new(MockCustomerService)
	handler := &CustomerHandler{cu: mockService}

	app := fiber.New()
	NewCustomerHandler(app, mockService)
	app.Post("/customers", handler.CreateCustomer)

	t.Run("successful customer creation", func(t *testing.T) {
		// Mock behavior for CreateCustomer
		mockService.On("CreateCustomer", mock.Anything).Return(nil)

		// Make request to create a customer
		reqBody := `{"name": "John Doe", "age": 30}`
		req := httptest.NewRequest("POST", "/customers", bytes.NewBufferString(reqBody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		// Assert that there were no errors
		assert.NoError(t, err)

		// Assert that the HTTP status code is correct
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)

		// Assert the response body contains the expected message
		expectedResponse := `{"Message":"create customer successful"}`
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		assert.NoError(t, err)
		assert.Equal(t, expectedResponse, string(bodyBytes))

		// Assert that the expected method was called
		mockService.AssertCalled(t, "CreateCustomer", mock.Anything)
	})

	t.Run("error invalid body create customer", func(t *testing.T) {
		// Mock behavior for CreateCustomer
		mockService.On("CreateCustomer", mock.Anything).Return(errors.New("failed to create customer"))

		// Make request to create a customer
		reqBody := `{"name": "", "age": -30}`
		req := httptest.NewRequest("POST", "/customers", bytes.NewBufferString(reqBody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		// Assert that there were no errors
		assert.NoError(t, err)

		// Assert that the HTTP status code is correct
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

		// Assert the response body contains the expected error message
		expectedResponse := `{"Age":"Age is min","Name":"Name is required"}`
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		assert.NoError(t, err)
		assert.Equal(t, expectedResponse, string(bodyBytes))

		// Assert that the expected method was called
		mockService.AssertCalled(t, "CreateCustomer", mock.Anything)
	})
}

func TestGetCustomerHandler(t *testing.T) {
	mockService := new(MockCustomerService)
	handler := &CustomerHandler{cu: mockService}

	app := fiber.New()
	NewCustomerHandler(app, mockService)
	app.Get("/customers/:id", handler.GetCustomer)

	t.Run("successful get customer", func(t *testing.T) {
		// Mock behavior for GetCustomerByID
		expectedCustomer := &entity.Customer{Name: "test", Age: 11}
		mockService.On("GetCustomerByID", "1").Return(expectedCustomer, nil)

		// Make request to get a customer
		req := httptest.NewRequest("GET", "/customers/1", nil)
		resp, err := app.Test(req)

		// Assert that there were no errors
		assert.NoError(t, err)

		// Assert that the HTTP status code is correct
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		// Assert that the response body matches the expected customer
		var responseBody entity.Customer
		err = json.NewDecoder(resp.Body).Decode(&responseBody)
		assert.NoError(t, err)

		// Assert that the expected method was called
		mockService.AssertExpectations(t)
	})

	t.Run("error getting customer", func(t *testing.T) {
		// Mock behavior for GetCustomerByID
		mockService.On("GetCustomerByID", "invalid_id").Return(nil, domain.ErrNotFound)

		// Make request to get a customer with invalid ID
		req := httptest.NewRequest("GET", "/customers/invalid_id", nil)
		resp, err := app.Test(req)

		// Assert that there were no errors
		assert.NoError(t, err)

		// Assert that the HTTP status code is correct
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)

		// Assert that the response body matches the expected error message
		var responseBody ResponseError
		err = json.NewDecoder(resp.Body).Decode(&responseBody)
		assert.NoError(t, err)
		assert.Equal(t, "your requested Item is not found", responseBody.Message)

		// Assert that the expected method was called
		mockService.AssertExpectations(t)
	})
}

func TestUpdateCustomerHandler(t *testing.T) {
	mockService := new(MockCustomerService)
	handler := &CustomerHandler{cu: mockService}

	app := fiber.New()
	NewCustomerHandler(app, mockService)
	app.Put("/customers/:id", handler.UpdateCustomer)

	t.Run("valid request body", func(t *testing.T) {
		// Prepare a valid request body
		validReqBody := `{"name": "John Doe", "age": 30}`

		// Mock service response
		mockService.On("UpdateCustomerByID", mock.AnythingOfType("*entity.Customer"), mock.AnythingOfType("string")).Return(nil)

		// Make request with valid request body
		req := httptest.NewRequest("PUT", "/customers/1", bytes.NewBufferString(validReqBody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		// Assert that there were no errors
		assert.NoError(t, err)

		// Assert that the HTTP status code is correct
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		// Assert that the expected method was called
		mockService.AssertExpectations(t)
	})

	t.Run("invalid request body", func(t *testing.T) {
		// Prepare an invalid request body (age out of range)
		invalidReqBody := `{"name": "John Doe", "age": 150}`

		// Make request with invalid request body
		req := httptest.NewRequest("PUT", "/customers/1", bytes.NewBufferString(invalidReqBody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)

		// Assert that there were no errors
		assert.NoError(t, err)

		// Assert that the HTTP status code is correct
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

		// Assert that the expected error message is returned
		expectedErrorMessage := `{"Age":"Age is max"}`
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		assert.NoError(t, err)
		assert.Equal(t, expectedErrorMessage, string(bodyBytes))

		// Assert that the expected method was called
		mockService.AssertExpectations(t)
	})

	// t.Run("not found user id", func(t *testing.T) {
	// 	// Prepare a valid request body
	// 	validReqBody := `{"name": "John Doe", "age": 30}`

	// 	// Mock service response
	// 	mockService.On("UpdateCustomerByID", mock.AnythingOfType("*entity.Customer"), mock.AnythingOfType("string")).Return(domain.ErrNotFound)

	// 	// Make request with valid request body
	// 	req := httptest.NewRequest("PUT", "/customers/invalid_id", bytes.NewBufferString(validReqBody))
	// 	req.Header.Set("Content-Type", "application/json")
	// 	resp, err := app.Test(req)

	// 	// Assert that there were no errors
	// 	assert.NoError(t, err)

	// 	// Assert that the HTTP status code is correct
	// 	assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)

	// 	// Assert that the response body matches the expected error message
	// 	var responseBody ResponseError
	// 	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	// 	assert.NoError(t, err)
	// 	assert.Equal(t, "your requested Item is not found", responseBody.Message)

	// 	// Assert that the expected method was called
	// 	mockService.AssertExpectations(t)
	// })
}

func TestDeleteCustomerHandler(t *testing.T) {
	mockService := new(MockCustomerService)
	handler := &CustomerHandler{cu: mockService}

	app := fiber.New()
	NewCustomerHandler(app, mockService)
	app.Delete("/customers/:id", handler.DeleteCustomer)

	t.Run("delete customer by existing id", func(t *testing.T) {
		// Mock service response
		mockService.On("DelCustomerByID", mock.AnythingOfType("string")).Return(nil)
	
		// Make request with valid id
		req := httptest.NewRequest("DELETE", "/customers/existing_id", nil)
		resp, err := app.Test(req)
	
		// Assert that there were no errors
		assert.NoError(t, err)
	
		// Assert that the HTTP status code is correct
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	
		// Assert that the expected method was called
		mockService.AssertExpectations(t)
	})
}