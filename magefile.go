// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/magefile/mage/mg"
)

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

// Build builds the binary
func Build() error {
	mg.Deps(InstallDeps)
	os.Setenv("CGO_ENABLED", "0")
	os.Setenv("GOOS", "linux")
	fmt.Println("Building...")
	cmd := exec.Command("go", "build", "-o", "api_server", "./cmd/api/server.go")
	return cmd.Run()
}

// Install installs the binary in /usr/local/bin
func Install() error {
	mg.Deps(Build)
	fmt.Println("Installing...")
	return os.Rename("./api_server", "/usr/bin/api_server")
}

// InstallDeps installs all the needed dependencies
func InstallDeps() error {
	fmt.Println("Installing Deps...")
	cmd := exec.Command("go", "get", "./...")
	return cmd.Run()
}

// Clean cleans up everything.
func Clean() {
	fmt.Println("Cleaning...")
}

// Start starts the API HTPP server. Currently using pm2 as process manager
func Start() error {
	mg.Deps(InstallDeps, Build)
	cmd := exec.Command("pm2", "start", "./api_server", "--name", "api")
	return cmd.Run()
}

// Start starts the API HTPP server. Currently using pm2 as process manager
func StartAPI() error {
	mg.Deps(InstallDeps, Build)
	os.Setenv("GIN_MODE", "release")
	cmd := exec.Command("./api_server")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Stop stops the API HTTP server.
func Stop() error {
	cmd := exec.Command("pm2", "stop", "api")
	return cmd.Run()
}

// StartDev startdev bootstraps a working dev environment.
func StartDev() error {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error: ", err.Error())
	}
	dataVol := fmt.Sprintf("%s/internal/db/data:/var/lib/mysql", cwd)
	sqlVol := fmt.Sprintf("%s/internal/db/sql:/sql", cwd)
	cmd := exec.Command("docker", "run", "--name", "mariadb", "-e", "MYSQL_RANDOM_ROOT_PASSWORD=yes", "-e",
		"MYSQL_PASSWORD=workday", "-e", "MYSQL_USER=workday", "-e", "MYSQL_DATABASE=workday", "-v",
		dataVol, "-v", sqlVol, "-p", "3306:3306", "-d", "mariadb:10.4",
	)
	err = cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return nil
}

// CreateVM creates a vagrant box already provisioned
func CreateVM() error {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error: ", err.Error())
	}
	cmd := exec.Command("packer", "build", "-var", "'version=10.2.0'", "debian10.json")
	cmd.Dir = cwd + "/.packer"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// ProvisionVM runs an Ansible playbook to provision the configured host.
func ProvisionVM() error {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error: ", err.Error())
	}
	cmd := exec.Command("ansible-playbook", "-i", "env/packer", "workday.yml")
	cmd.Dir = cwd + "/.packer/ansible"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Test runs the tests suite.
func Test() error {
	fmt.Println("Running tests")
	os.Setenv("GIN_MODE", "test")
	cmd := exec.Command("go", "test", "-v", "./...")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return nil
}

// TestVerbose runs the tests suite with verbosity
func TestVerbose() error {
	fmt.Println("Running tests")
	os.Setenv("GIN_MODE", "test")
	cmd := exec.Command("ginkgo", "-v", "test", "./...")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return nil
}

// TestAndCoverage runs tests and generate the code coverage.
func TestAndCoverage() error {
	fmt.Println("Running tests and generating code coverage")
	os.Setenv("GIN_MODE", "test")
	cmd := exec.Command("go", "test", "-v", "-cover", "-coverprofile=workday.coverprofile", "./...")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	cmd = exec.Command("go", "tool", "cover", "-html=workday.coverprofile", "-o", "docs/static/coverage.html")
	return cmd.Run()
}
