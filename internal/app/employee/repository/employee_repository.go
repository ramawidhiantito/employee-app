package repository

import (
	"employee_app/internal/app/models"

	"github.com/jinzhu/gorm"
)

type EmployeeRepository struct {
	DB *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{DB: db}
}

func (r *EmployeeRepository) Create(employee *models.Employee) error {
	return r.DB.Create(employee).Error
}

func (r *EmployeeRepository) FindAll() ([]models.Employee, error) {
	var employees []models.Employee
	return employees, r.DB.Find(&employees).Error
}

func (r *EmployeeRepository) FindByID(id uint) (*models.Employee, error) {
	var employee models.Employee
	return &employee, r.DB.First(&employee, id).Error
}

func (r *EmployeeRepository) Update(employee *models.Employee) error {
	return r.DB.Save(employee).Error
}

func (r *EmployeeRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Employee{}, id).Error
}
