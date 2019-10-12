package workday

import (
	"database/sql"
)

// Represents an employee
type Employee struct {
	ID        uint   `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Role      uint   `json:role`
	Password  string `json:password`
}

// CreateEmployee saves a new employee in DB.
func CreateEmployee(db *sql.DB, employee Employee) (id uint, rows uint, err error) {
	return insertNewEmployee(db, employee)
}

// GetAllEmployees retrieves all employees from DB.
func GetAllEmployees(db *sql.DB) (employees []Employee, err error) {
	return findAllEmployees(db)
}
