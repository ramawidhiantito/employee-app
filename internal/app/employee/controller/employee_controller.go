package controller

import (
	"employee_app/internal/app/employee/usecase"
	"employee_app/internal/app/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// EmployeeController represents the controller for employee-related operations
type EmployeeController struct {
	useCase usecase.EmployeeUseCase
}

// NewEmployeeController creates a new EmployeeController with the provided use case
func NewEmployeeController(useCase usecase.EmployeeUseCase) *EmployeeController {
	return &EmployeeController{useCase: useCase}
}

// @Summary Get a list of employees
// @Description Get a list of all employees
// @ID get-employees
// @Produce json
// @Success 200 {array} models.Employee "Successful response with a list of Employees"
// @Router /employees/ [get]
func (ec *EmployeeController) GetEmployees(c *gin.Context) {
	employees, err := ec.useCase.GetEmployees()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch employees"})
		return
	}

	c.JSON(http.StatusOK, employees)
}

// @Summary Get an employee by ID
// @Description Get an employee by ID
// @ID get-employee
// @Produce json
// @Param id path int true "Employee ID"
// @Success 200 {object} models.Employee "Successful response with an Employee"
// @Router /employees/{id} [get]
func (ec *EmployeeController) GetEmployee(c *gin.Context) {
	employeeID, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	employee, err := ec.useCase.GetEmployeeByID(uint(employeeID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	c.JSON(http.StatusOK, employee)
}

// @Summary Create a new employee
// @Description Create a new employee
// @ID create-employee
// @Accept json
// @Param employee body models.Employee true "Employee object"
// @Success 201 {object} models.Employee "Successful response with a created Employee"
// @Router /employees/ [post]
func (ec *EmployeeController) CreateEmployee(c *gin.Context) {
	var employee models.Employee

	if err := ec.useCase.CreateEmployee(&employee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create employee"})
		return
	}

	c.JSON(http.StatusCreated, employee)
}

// @Summary Update an employee by ID
// @Description Update an employee by ID
// @ID update-employee
// @Accept json
// @Param id path int true "Employee ID"
// @Param employee body models.Employee true "Employee object"
// @Success 200 {object} models.Employee "Successful response with an updated Employee"
// @Router /employees/{id} [put]
func (ec *EmployeeController) UpdateEmployee(c *gin.Context) {
	employeeID, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	employee.ID = uint(employeeID)

	if err := ec.useCase.UpdateEmployee(&employee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update employee"})
		return
	}

	c.JSON(http.StatusOK, employee)
}

// @Summary Delete an employee by ID
// @Description Delete an employee by ID
// @ID delete-employee
// @Param id path int true "Employee ID"
// @Success 204 "No Content" "Successful response with no content"
// @Router /employees/{id} [delete]
func (ec *EmployeeController) DeleteEmployee(c *gin.Context) {
	employeeID, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	if err := ec.useCase.DeleteEmployee(uint(employeeID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete employee"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
