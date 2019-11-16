package api_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/iris-garcia/workday/api"
)

var _ = Describe("Config", func() {
	It("Should read a config file correctly", func() {
		cfg, err := LoadDBConfig("../db_config.toml")
		expected := DBConfig{
			Host:     "127.0.0.1",
			Database: "workday",
			User:     "workday",
			Password: "workday",
		}

		Expect(err).NotTo(HaveOccurred())
		Expect(cfg).To(Equal(expected))
	})

	It("Should get a DBConfig reading environment variables", func() {
		os.Setenv("MARIADB_SERVICE_HOST", "1.1.1.1")
		os.Setenv("WORKDAY_DB_NAME", "name")
		os.Setenv("WORKDAY_DB_USER", "user")
		os.Setenv("WORKDAY_DB_PASSWORD", "pass")
		cfg, err := LoadDBConfig("./db_config.toml")
		expected := DBConfig{
			Host:     "1.1.1.1",
			Database: "name",
			User:     "user",
			Password: "pass",
		}

		Expect(err).NotTo(HaveOccurred())
		Expect(cfg).To(Equal(expected))

		// Clean up environment vars
		os.Unsetenv("MARIADB_SERVICE_HOST")
		os.Unsetenv("WORKDAY_DB_NAME")
		os.Unsetenv("WORKDAY_DB_USER")
		os.Unsetenv("WORKDAY_DB_PASSWORD")
	})

	It("Should return an error when the config file does not exists", func() {
		_, err := LoadDBConfig("./random_file.toml")
		Expect(err).To(HaveOccurred())
	})
})
