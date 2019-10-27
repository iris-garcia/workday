package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	. "github.com/iris-garcia/workday/api"
)

// Test the GET /status endpoint 200
func TestShouldGetStatusEndpointStatusOK(t *testing.T) {
	assert := assert.New(t)
	db, _, err := sqlmock.New()
	defer db.Close()

	assert.Nil(err)
	assert.NotNil(db)

	router := GinRouter(db)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/status", nil)
	router.ServeHTTP(w, req)

	assert.Equal(200, w.Code)
}

// Test the GET /employees endpoint 200
func TestShouldGetEmployeesEndpointStatusOK(t *testing.T) {
	assert := assert.New(t)
	db, mock, err := sqlmock.New()
	defer db.Close()

	assert.Nil(err)
	assert.NotNil(db)

	router := GinRouter(db)

	rows := sqlmock.NewRows([]string{"ID", "Firstname", "Lastname", "Role", "Password"})
	mock.ExpectQuery("^SELECT (.+) FROM employee$").WillReturnRows(rows)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/employees", nil)
	router.ServeHTTP(w, req)

	assert.Equal(200, w.Code)
}

// Test the GET /employees endpoint 500
func TestShouldGetEmployeesEndpointStatusInternalError(t *testing.T) {
	assert := assert.New(t)
	db, _, err := sqlmock.New()
	defer db.Close()

	assert.Nil(err)
	assert.NotNil(db)

	router := GinRouter(db)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/employees", nil)
	router.ServeHTTP(w, req)

	assert.Equal(500, w.Code)
}
