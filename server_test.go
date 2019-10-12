package workday

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/kataras/iris/httptest"
)

// Test the GET /employees endpoint 200
func TestShouldGetEmployeesEndpointStatusOK(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	app := IrisHTTPServer(db)
	if err != nil {
		t.Errorf("Iris HTTP server failed: %v", err.Error())
	}

	rows := sqlmock.NewRows([]string{"ID", "Firstname", "Lastname", "Role", "Password"})
	mock.ExpectQuery("^SELECT (.+) FROM employee$").WillReturnRows(rows)

	e := httptest.New(t, app)
	e.GET("/employees").Expect().Status(httptest.StatusOK)
}

// Test the GET /employees endpoint 500
func TestShouldGetEmployeesEndpointStatusInternalError(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	app := IrisHTTPServer(db)

	e := httptest.New(t, app)
	e.GET("/employees").Expect().Status(httptest.StatusInternalServerError)
}
