package e2e_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestUserAPIE2E(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "User API E2E Suite")
}

