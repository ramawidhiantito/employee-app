package controller_test

import (
	"bytes"
	"employee_app/internal/app/employee/controller"
	"employee_app/internal/app/models"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockEmployeeUseCase struct {
	mock.Mock
}

func (m *MockEmployeeUseCase) CreateEmployee(employee *models.Employee) error {
	args := m.Called(employee)
	return args.Error(0)
}

func (m *MockEmployeeUseCase) GetEmployees() ([]models.Employee, error) {
	args := m.Called()
	return args.Get(0).([]models.Employee), args.Error(1)
}

func (m *MockEmployeeUseCase) GetEmployeeByID(id uint) (*models.Employee, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Employee), args.Error(1)
}

func (m *MockEmployeeUseCase) UpdateEmployee(employee *models.Employee) error {
	args := m.Called(employee)
	return args.Error(0)
}

func (m *MockEmployeeUseCase) DeleteEmployee(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestGetEmployee(t *testing.T) {
	// Create a test Gin router
	router := gin.Default()

	// Create a mock use case and controller
	mockUseCase := &MockEmployeeUseCase{}
	controller := controller.NewEmployeeController(mockUseCase)

	// Set up the test route
	router.GET("/employees/:id", controller.GetEmployee)

	//Scenario Test 1 : Valid employee ID
	t.Run("Valid Employee ID", func(t *testing.T) {
		testEmployeeID := uint(1)
		expectedEmployee := &models.Employee{
			ID:        testEmployeeID,
			FirstName: "Rama",
			LastName:  "Widhiantito",
			Email:     "rama.widhiantito@gmail.com",
			HireDate:  time.Date(2023, 12, 16, 0, 0, 0, 0, time.UTC),
		}

		// Mock the use case to return the expected employee
		mockUseCase.On("GetEmployeeByID", testEmployeeID).Return(expectedEmployee, nil)

		// Create a test request
		req := httptest.NewRequest("GET", "/employees/1", nil)
		w := httptest.NewRecorder()

		// Perform the request
		router.ServeHTTP(w, req)

		// Assert the response
		assert.Equal(t, http.StatusOK, w.Code)

		// Parse the JSON response and assert its contents
		var responseEmployee models.Employee
		err := json.NewDecoder(w.Body).Decode(&responseEmployee)
		assert.NoError(t, err)
		assert.Equal(t, expectedEmployee, &responseEmployee)
	})

	//Scenario Test 2 : Invalid Employee ID
	t.Run("Invalid Employee ID", func(t *testing.T) {
		// Create a test request with an invalid ID
		req := httptest.NewRequest("GET", "/employees/invalid", nil)
		w := httptest.NewRecorder()

		// Perform the request
		router.ServeHTTP(w, req)

		// Assert the response
		assert.Equal(t, http.StatusBadRequest, w.Code)

		// Parse the JSON response and assert the error message
		var responseJSON map[string]string
		err := json.NewDecoder(w.Body).Decode(&responseJSON)
		assert.NoError(t, err)
		assert.Contains(t, responseJSON["error"], "Invalid employee ID")
	})

	//Scenario Test 3 : Employee not found
	t.Run("Employee Not Found", func(t *testing.T) {
		testEmployeeID := uint(999)

		expectedEmployee := &models.Employee{}
		// Mock the use case to return an error indicating that the employee was not found
		mockUseCase.On("GetEmployeeByID", testEmployeeID).Return(expectedEmployee, errors.New("Employee not found"))

		// Create a test request with a valid ID
		req := httptest.NewRequest("GET", "/employees/999", nil)
		w := httptest.NewRecorder()

		// Perform the request
		router.ServeHTTP(w, req)

		// Assert the response
		assert.Equal(t, http.StatusNotFound, w.Code)

		// Parse the JSON response and assert the error message
		var responseJSON map[string]string
		err := json.NewDecoder(w.Body).Decode(&responseJSON)
		assert.NoError(t, err)
		assert.Contains(t, responseJSON["error"], "Employee not found")
	})
}

func TestUpdateEmployee(t *testing.T) {
	// Create a test Gin router
	router := gin.Default()

	// Create a mock use case and controller
	mockUseCase := &MockEmployeeUseCase{}
	controller := controller.NewEmployeeController(mockUseCase)

	// Set up the test route
	router.PUT("/employees/:id", controller.UpdateEmployee)

	//Scenario 1 Invalid Employee ID
	t.Run("Invalid Employee ID", func(t *testing.T) {
		// Create a request with an invalid employee ID
		req := httptest.NewRequest("PUT", "/employees/invalid_id", nil)
		w := httptest.NewRecorder()

		// Perform the request
		router.ServeHTTP(w, req)

		// Check the response status code and the response body
		assert.Equal(t, http.StatusBadRequest, w.Code)

		var responseBody map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "Invalid employee ID", responseBody["error"])
	})

	//Scenario 2 Valid Employee ID
	t.Run("Valid Employee ID", func(t *testing.T) {
		employee := models.Employee{
			ID:        uint(1),
			FirstName: "RamaUpdate",
			LastName:  "Widhiantito",
			Email:     "rama.widhiantito@gmail.com",
			HireDate:  time.Date(2023, 12, 16, 0, 0, 0, 0, time.UTC),
		}
		payload, _ := json.Marshal(employee)

		// Mock the use case to return updated data
		mockUseCase.On("UpdateEmployee", mock.MatchedBy(func(e *models.Employee) bool {
			return e.ID == 1 && e.FirstName == "RamaUpdate" && e.LastName == "Widhiantito" && e.Email == "rama.widhiantito@gmail.com"
		})).Return(nil)

		// Create a request with a valid employee ID and payload
		req := httptest.NewRequest("PUT", "/employees/1", bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		// Perform the request
		router.ServeHTTP(w, req)

		// Check the response status code and response body for the valid scenario
		assert.Equal(t, http.StatusOK, w.Code)

		// Parse the response body and check if it matches the expected employee
		var responseEmployee models.Employee
		err := json.Unmarshal(w.Body.Bytes(), &responseEmployee)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, employee, responseEmployee)
	})
}

func TestGetEmployees(t *testing.T) {
	// Create a test Gin router
	router := gin.Default()

	// Create a mock use case and controller
	mockUseCase := &MockEmployeeUseCase{}
	controller := controller.NewEmployeeController(mockUseCase)

	// Set up the test route
	router.GET("/employees", controller.GetEmployees)

	//Scenario 1 Sucess get all
	t.Run("Success", func(t *testing.T) {
		// Define a sample list of employees
		employees := []models.Employee{
			{
				ID:        1,
				FirstName: "John",
				LastName:  "Doe",
				Email:     "john.doe@example.com",
				HireDate:  time.Date(2023, 12, 16, 0, 0, 0, 0, time.UTC),
			},
			{
				ID:        2,
				FirstName: "Jane",
				LastName:  "Smith",
				Email:     "jane.smith@example.com",
				HireDate:  time.Date(2023, 12, 16, 0, 0, 0, 0, time.UTC),
			},
		}

		// Mock the use case to indicate a successful fetch of employees
		mockUseCase.On("GetEmployees").Return(employees, nil)

		// Create a request to get employees
		req := httptest.NewRequest("GET", "/employees", nil)
		w := httptest.NewRecorder()

		// Perform the request
		router.ServeHTTP(w, req)

		// Check the response for the "Success" scenario
		assert.Equal(t, http.StatusOK, w.Code)

		// Parse the JSON response and check the list of employees
		var responseEmployees []models.Employee
		err := json.NewDecoder(w.Body).Decode(&responseEmployees)
		assert.NoError(t, err)
		assert.Equal(t, employees, responseEmployees)
		mockUseCase.AssertExpectations(t)
	})
}

func TestDeleteEmployee(t *testing.T) {
	// Create a test Gin router
	router := gin.Default()

	// Create a mock use case and controller
	mockUseCase := &MockEmployeeUseCase{}
	controller := controller.NewEmployeeController(mockUseCase)

	// Set up the test route
	router.DELETE("/employees/:id", controller.DeleteEmployee)

	//Scenario 1 Valid employee ID
	t.Run("Valid Employee ID", func(t *testing.T) {
		// Define a test employee ID
		testEmployeeID := "1"

		// Mock the use case to indicate a successful deletion
		mockUseCase.On("DeleteEmployee", uint(1)).Return(nil)

		// Create a request with a valid employee ID
		req := httptest.NewRequest("DELETE", "/employees/"+testEmployeeID, nil)
		w := httptest.NewRecorder()

		// Perform the request
		router.ServeHTTP(w, req)

		// Check the response for the "Valid Employee ID" scenario
		assert.Equal(t, http.StatusNoContent, w.Code)
		mockUseCase.AssertExpectations(t)
	})

	//Scenario 2 Invalid Employee ID
	t.Run("Invalid Employee ID", func(t *testing.T) {
		// Create a request with an invalid employee ID
		req := httptest.NewRequest("DELETE", "/employees/invalid", nil)
		w := httptest.NewRecorder()

		// Perform the request
		router.ServeHTTP(w, req)

		// Check the response for the "Invalid Employee ID" scenario
		assert.Equal(t, http.StatusBadRequest, w.Code)
		// Parse the JSON response and assert the error message
		var responseJSON map[string]string
		err := json.NewDecoder(w.Body).Decode(&responseJSON)
		assert.NoError(t, err)
		assert.Contains(t, responseJSON["error"], "Invalid employee ID")
	})

	//Scenario 3 Failed to delete
	t.Run("Failed Deletion", func(t *testing.T) {
		// Define a test employee ID
		testEmployeeID := "2"

		// Mock the use case to indicate a failed deletion
		mockUseCase.On("DeleteEmployee", uint(2)).Return(errors.New("Failed to delete employee"))

		// Create a request with a valid employee ID
		req := httptest.NewRequest("DELETE", "/employees/"+testEmployeeID, nil)
		w := httptest.NewRecorder()

		// Perform the request
		router.ServeHTTP(w, req)

		// Check the response for the "Failed Deletion" scenario
		assert.Equal(t, http.StatusInternalServerError, w.Code)

		// Parse the JSON response and assert the error message
		var responseJSON map[string]string
		err := json.NewDecoder(w.Body).Decode(&responseJSON)
		assert.NoError(t, err)
		assert.Contains(t, responseJSON["error"], "Failed to delete employee")
	})
}
