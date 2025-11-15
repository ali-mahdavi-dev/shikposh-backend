package products_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestProductCommandHandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ProductCommandHandler Suite")
}
