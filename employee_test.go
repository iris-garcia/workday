package workday

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

type MyCustomError struct {
	message string
}

func (e MyCustomError) Error() string {
	return e.message
}

func assertJSON(a interface{}, b interface{}, t *testing.T) {
	actual, err := json.Marshal(a)
	expected, err := json.Marshal(b)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
	}

	if bytes.Compare(expected, actual) != 0 {
		t.Errorf("the expected json: %s is different from actual %s", expected, actual)
	}
}

func TestShouldCreateEmployee(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	if err != nil {
		t.Errorf("Iris HTTP server failed: %v", err.Error())
	}

	mock.ExpectExec("INSERT INTO employee").WithArgs("Iris", "Garcia", 1, "secret").
		WillReturnResult(sqlmock.NewResult(1, 1))

	emp := Employee{Firstname: "Iris", Lastname: "Garcia", Role: 1, Password: "secret"}

	id, rows, err := CreateEmployee(db, emp)

	if id != 1 {
		t.Error("lastID should be 1")
	}

	if rows != 1 {
		t.Error("rowsAffected should be 1")
	}

	// make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestShouldFailWhenDBExecCreateEmployee(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	if err != nil {
		t.Errorf("Iris HTTP server failed: %v", err.Error())
	}

	mock.ExpectExec("INSERT INTO employee").WithArgs("Iris", "Garcia", 1, "secret").
		WillReturnError(MyCustomError{message: "Error running db.Exec"})

	emp := Employee{Firstname: "Iris", Lastname: "Garcia", Role: 1, Password: "secret"}

	_, _, err = CreateEmployee(db, emp)

	if err == nil {
		t.Error("Expected an error when the db.Exec fails")
	}

	// make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestShouldFailWhenCreateEmployee(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	if err != nil {
		t.Errorf("Iris HTTP server failed: %v", err.Error())
	}

	mock.ExpectExec("INSERT INTO employee").WithArgs("Iris", "Garcia", 1, "secret").
		WillReturnResult(sqlmock.NewErrorResult(MyCustomError{message: "DB error"}))

	emp := Employee{Firstname: "Iris", Lastname: "Garcia", Role: 1, Password: "secret"}

	id, rows, err := CreateEmployee(db, emp)

	if err == nil {
		t.Error("Expected an error when a new employee is created and the lastID is 0")
	}

	if id == 1 {
		t.Error("Should not get any lastID inserted")
	}

	if rows == 1 {
		t.Error("Should not get any row affected")
	}

	// make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestShouldGetAllEmployees(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// expect required DB actions
	rows := sqlmock.NewRows([]string{"ID", "Firstname", "Lastname", "Role", "Password"}).
		AddRow(1, "Iris", "Garcia", 1, "secret").
		AddRow(2, "Name", "Lastname", 2, "changeme")

	mock.ExpectQuery("^SELECT (.+) FROM employee$").WillReturnRows(rows)

	expected := []Employee{
		Employee{ID: 1, Firstname: "Iris", Lastname: "Garcia", Role: 1, Password: "secret"},
		Employee{ID: 2, Firstname: "Name", Lastname: "Lastname", Role: 2, Password: "changeme"},
	}

	emps, err := GetAllEmployees(db)

	assertJSON(emps, expected, t)

	// make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
