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

// Builds the binary
func Build() error {
	mg.Deps(InstallDeps)
	fmt.Println("Building...")
	cmd := exec.Command("go", "build", "-o", "api_server", "./cmd/api/server.go")
	return cmd.Run()
}

// Install the binary in /usr/local/bin
func Install() error {
	mg.Deps(Build)
	fmt.Println("Installing...")
	return os.Rename("./api_server", "/usr/bin/api_server")
}

// Install dependencies
func InstallDeps() error {
	fmt.Println("Installing Deps...")
	cmd := exec.Command("go", "get", "./api/")
	return cmd.Run()
}

// Clean up after yourself
func Clean() {
	fmt.Println("Cleaning...")
}

// Start the api_server binary which is an HTTP server.
func Start() error {
	mg.Deps(InstallDeps, Build)
	cmd := exec.Command("./api_server")
	return cmd.Run()
}

// StartDev will bootstrap a working dev environment
// TODO:
//  - Use realize to automatically reload the api server on fiche changes
//  - Check if the docker container for mariadb is already running
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

// Run tests.
func Test() error {
	fmt.Println("Running tests")
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

// Run tests and generate the code coverage.
func TestAndCoverage() error {
	fmt.Println("Running tests and generating code coverage")
	cmd := exec.Command("go", "test", "-v", "-cover", "-coverprofile=workday.coverprofile", "./...")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	cmd = exec.Command("go", "tool", "cover", "-html=workday.coverprofile", "-o", "coverage.html")
	return cmd.Run()
}
