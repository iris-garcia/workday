package workday

import (
	"testing"
)

// Test the creation of the HTTPServer
func TestIrisHTTPServer(t *testing.T) {
	_, err := IrisHTTPServer()
	if err != nil {
		t.Errorf("Iris HTTP server failed: %v", err.Error())
	}
}
