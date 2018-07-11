package files

import (
	"testing"
	. "github.com/onsi/gomega"
	. "github.com/onsi/ginkgo"
)

func TestParseTemplate(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ParseTemplate")
}

var _ = Describe("ParseTemplate", func(){
	It("should parse normal string correctly", func() {
		template, err := ParseTemplate("hello, world")
		Expect(err).To(BeNil())
		Expect(template).ToNot(BeNil())
		Expect(template.tpl).ToNot(BeNil())
	})

	It("should parse empty string correctly", func() {
		template, err := ParseTemplate("")
		Expect(err).To(BeNil())
		Expect(template).ToNot(BeNil())
		Expect(template.tpl).ToNot(BeNil())
	})

	It("should parse a valid template correctly", func() {
		template, err := ParseTemplate("{{ .Hello }} - {{ .World }}")
		Expect(err).To(BeNil())
		Expect(template).ToNot(BeNil())
		Expect(template.tpl).ToNot(BeNil())
	})

	It("should not parse an invalid, and should return an error", func() {
		template, err := ParseTemplate("{{ .Hello } - {{ .World }}")
		Expect(err).ToNot(BeNil())
		Expect(template).To(BeNil())
	})
})
