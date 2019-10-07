package workday

import (
	"encoding/json"
	"testing"

	"github.com/kataras/iris"
	"github.com/kataras/iris/httptest"
)

// Test the creation of the HTTPServer
func TestIrisHTTPServer(t *testing.T) {
	_, err := IrisHTTPServer()
	if err != nil {
		t.Errorf("Iris HTTP server failed: %v", err.Error())
	}
}

// Test the GET /employees endpoint
func TestGetEmployees(t *testing.T) {
	app, err := IrisHTTPServer()
	if err != nil {
		t.Errorf("Iris HTTP server failed: %v", err.Error())
	}
	e := httptest.New(t, app)
	empmap := iris.Map{"message": "GET to employess endpoint"}
	expected, _ := json.Marshal(empmap)
	e.GET("/employees").Expect().Body().Equal(string(expected))
}
