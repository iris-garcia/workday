package api

import (
	"database/sql"
)

// Employee Represents an employee
type Employee struct {
	ID        uint   `json:"id"`
	Firstname string `binding:"required" json:"firstname"`
	Lastname  string `binding:"required" json:"lastname"`
	Role      uint   `binding:"required" json:"role"`
	Password  string `binding:"required" json:"password"`
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

// EditEmployee updates an employee from DB.
func EditEmployee(db *sql.DB, emp *Employee) (uint, error) {
	return updateEmployee(db, emp)
}
