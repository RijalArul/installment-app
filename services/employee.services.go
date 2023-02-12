package services

import (
	"test-kr-sigma/models/entities"
	"test-kr-sigma/models/web"
	"test-kr-sigma/repositories"
)

func ResponseEmployeeBody(employee entities.Employee) web.EmployeeBodyResponse {
	return web.EmployeeBodyResponse{
		Email:    employee.Email,
		Username: employee.Username,
		Role:     employee.Role,
	}
}

type EmployeeService interface {
	Register(employeeDTO web.RegisterEmployeeDTO) (web.EmployeeBodyResponse, error)
	Login(employeeDTO web.LoginRequestDTO) (*entities.Employee, error)
}

type EmployeeServiceImpl struct {
	employeeRepo repositories.EmployeeRepository
}

func NewEmployeeService(empRepo repositories.EmployeeRepository) EmployeeService {
	return &EmployeeServiceImpl{employeeRepo: empRepo}
}

func (employeeService *EmployeeServiceImpl) Register(employeeDTO web.RegisterEmployeeDTO) (web.EmployeeBodyResponse, error) {
	employee := entities.Employee{
		Email:    employeeDTO.Email,
		Password: employeeDTO.Password,
		Username: employeeDTO.Password,
		Role:     employeeDTO.Role,
	}
	newEmployee, err := employeeService.employeeRepo.Register(employee)
	respEmployee := ResponseEmployeeBody(*newEmployee)
	return respEmployee, err
}

func (employeeService *EmployeeServiceImpl) Login(employeeDTO web.LoginRequestDTO) (*entities.Employee, error) {
	user, err := employeeService.employeeRepo.FindByEmail(employeeDTO.Email)
	return user, err
}
