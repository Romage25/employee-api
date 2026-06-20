package services

import (
	"employee-api/models"
	"employee-api/repositories"
)

type EmployeeService struct {
	repo *repositories.EmployeeRepository
}

func NewEmployeeService(
	repo *repositories.EmployeeRepository,
) *EmployeeService {

	return &EmployeeService{
		repo: repo,
	}
}

func (s *EmployeeService) Create(
	employee *models.Employee,
) error {

	return s.repo.Create(employee)
}

func (s *EmployeeService) GetAll() (
	[]models.Employee,
	error,
) {

	return s.repo.FindAll()
}

func (s *EmployeeService) GetByID(
	id string,
) (*models.Employee, error) {

	return s.repo.FindByID(id)
}

func (s *EmployeeService) Update(
	id string,
	data map[string]interface{},
) (*models.Employee, error) {

	employee, err := s.repo.FindByID(id)

	if err != nil {
		return nil, err
	}

	err = s.repo.Update(employee, data)

	return employee, err
}

func (s *EmployeeService) Delete(
	id string,
) error {

	employee, err := s.repo.FindByID(id)

	if err != nil {
		return err
	}

	return s.repo.Delete(employee)
}
