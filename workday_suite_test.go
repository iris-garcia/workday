package workday_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestWorkday(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Workday Suite")
}
