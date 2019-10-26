package api

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
func CreateEmployee(db *sql.DB, emp *Employee) (uint, uint, error) {
	return insertNewEmployee(db, emp)
}

// GetAllEmployees retrieves all employees from DB.
func GetAllEmployees(db *sql.DB) ([]Employee, error) {
	return findAllEmployees(db)
}

// GetEmployee retrieves an employee from DB.
func GetEmployee(db *sql.DB, id uint) (Employee, error) {
	return findEmployee(db, id)
}

// DeleteEmployee removes an employee from DB.
func DeleteEmployee(db *sql.DB, id uint) (uint, error) {
	return removeEmployee(db, id)
}
