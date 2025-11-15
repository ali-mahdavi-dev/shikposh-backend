package e2e_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestProductAPIE2E(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Product API E2E Suite")
}

