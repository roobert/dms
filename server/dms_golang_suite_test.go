package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDmsGolang(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DmsGolang Suite")
}
