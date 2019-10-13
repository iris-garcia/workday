package workday_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/kataras/iris/httptest"
	"github.com/stretchr/testify/assert"

	. "github.com/iris-garcia/workday"
)

// Test the GET /employees endpoint 200
func TestShouldGetEmployeesEndpointStatusOK(t *testing.T) {
	assert := assert.New(t)
	db, mock, err := sqlmock.New()
	defer db.Close()

	assert.Nil(err)
	assert.NotNil(db)

	app := IrisHTTPServer(db)

	rows := sqlmock.NewRows([]string{"ID", "Firstname", "Lastname", "Role", "Password"})
	mock.ExpectQuery("^SELECT (.+) FROM employee$").WillReturnRows(rows)

	e := httptest.New(t, app)
	e.GET("/employees").Expect().Status(httptest.StatusOK)
}

// Test the GET /employees endpoint 500
func TestShouldGetEmployeesEndpointStatusInternalError(t *testing.T) {
	assert := assert.New(t)
	db, _, err := sqlmock.New()
	defer db.Close()

	assert.Nil(err)
	assert.NotNil(db)

	app := IrisHTTPServer(db)

	e := httptest.New(t, app)
	e.GET("/employees").Expect().Status(httptest.StatusInternalServerError)
}
