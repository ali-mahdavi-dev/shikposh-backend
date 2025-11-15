package acceptance_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestProductManagementScenarios(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Product Management Acceptance Scenarios Suite")
}
