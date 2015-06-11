package garbageman_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGarbageman(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pez Garbageman Suite")
}
