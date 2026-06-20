package repositories

import (
	"employee-api/models"

	"gorm.io/gorm"
)

type EmployeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{
		db: db,
	}
}

func (r *EmployeeRepository) Create(employee *models.Employee) error {
	return r.db.Create(employee).Error
}

func (r *EmployeeRepository) FindAll() ([]models.Employee, error) {
	var employees []models.Employee

	err := r.db.Find(&employees).Error

	return employees, err
}

func (r *EmployeeRepository) FindByID(id string) (*models.Employee, error) {

	var employee models.Employee

	err := r.db.First(&employee, id).Error

	return &employee, err
}

func (r *EmployeeRepository) Update(
	employee *models.Employee,
	data map[string]interface{},
) error {

	return r.db.
		Model(employee).
		Updates(data).
		Error
}

func (r *EmployeeRepository) Delete(
	employee *models.Employee,
) error {

	return r.db.Delete(employee).Error
}
