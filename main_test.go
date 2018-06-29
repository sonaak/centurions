package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestHello(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Main Suite")
}

var _ = Describe("Main Suite", func() {
	Describe("Hello function", func() {
		It("should say respond with 'Hello'", func() {
			Expect(Hello()).To(Equal("Hello world!"))
		})
	})
})
