package Company

import "fmt"

type Employee interface {
	GetDetails() string
}

type FullTimeEmployee struct {
	ID     uint64
	Name   string
	Salary float64
}

type PartTimeEmployee struct {
	ID          uint64
	Name        string
	HourlyRate  float64
	HoursWorked int
}

func (p *PartTimeEmployee) GetDetails() string {
	return fmt.Sprintf("id: %d, name: %s, hourly rate: %.2f, hours worked: %d", p.ID, p.Name, p.HourlyRate, p.HoursWorked)
}

func (f *FullTimeEmployee) GetDetails() string {
	return fmt.Sprintf("id: %d, name: %s, salary: %.2f", f.ID, f.Name, f.Salary)
}

type Company struct {
	Employees map[uint64]Employee
}

func New() *Company {
	return &Company{
		Employees: make(map[uint64]Employee),
	}
}

func (c *Company) AddEmployee(e Employee) {
	switch emp := e.(type) {
	case *FullTimeEmployee:
		c.Employees[emp.ID] = emp
	case *PartTimeEmployee:
		c.Employees[emp.ID] = emp

	}
}

func (c *Company) ListEmployees() {
	for _, employee := range c.Employees {
		fmt.Println(employee.GetDetails())
	}
}
