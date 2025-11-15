package acceptance_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestUserRegistrationScenarios(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "User Registration Acceptance Scenarios Suite")
}

