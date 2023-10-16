package usecase

import (
	"employee_app/internal/app/employee/repository"
	"employee_app/internal/app/models"
)

type EmployeeUseCase interface {
	CreateEmployee(employee *models.Employee) error
	GetEmployees() ([]models.Employee, error)
	GetEmployeeByID(id uint) (*models.Employee, error)
	UpdateEmployee(employee *models.Employee) error
	DeleteEmployee(id uint) error
}

type employeeUseCase struct {
	employeeRepo repository.EmployeeRepository
}

func NewEmployeeUseCase(employeeRepo repository.EmployeeRepository) *employeeUseCase {
	return &employeeUseCase{employeeRepo: employeeRepo}
}

func (uc *employeeUseCase) CreateEmployee(employee *models.Employee) error {
	return uc.employeeRepo.Create(employee)
}

func (uc *employeeUseCase) GetEmployees() ([]models.Employee, error) {
	return uc.employeeRepo.FindAll()
}

func (uc *employeeUseCase) GetEmployeeByID(id uint) (*models.Employee, error) {
	return uc.employeeRepo.FindByID(id)
}

func (uc *employeeUseCase) UpdateEmployee(employee *models.Employee) error {
	return uc.employeeRepo.Update(employee)
}

func (uc *employeeUseCase) DeleteEmployee(id uint) error {
	return uc.employeeRepo.Delete(id)
}
