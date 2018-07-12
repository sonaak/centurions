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

type testObject struct {
	Integer int
	Float64 float64
	Message string

	privateValue uint
}

func (obj *testObject) GetInteger() int {
	return obj.Integer
}


var _ = Describe("Template", func() {
	var testObj testObject

	BeforeEach(func() {
		testObj = testObject {
			Integer: -20,
			Float64: 0.321131,
			Message: "hello, 世界",
		}
	})

	It("should template correctly with public int attributes", func() {
		template, _ := ParseTemplate("{{ .Integer }}")
		Expect(template.Template(testObj)).To(Equal("-20"))
	})

	It("should template correctly with public float attributes", func() {
		template, _ := ParseTemplate("{{ .Float64 }}")
		Expect(template.Template(testObj)).To(Equal("0.321131"))
	})

	It("should template correctly with public string attributes", func() {
		template, _ := ParseTemplate("{{ .Message }}")
		Expect(template.Template(testObj)).To(Equal("hello, 世界"))
	})

	It("should not template correctly with private attributes", func() {
		template, _ := ParseTemplate("{{ .privateValue }}")
		str, err := template.Template(testObj)
		Expect(err).ToNot(BeNil())
		Expect(str).To(Equal(""))
	})

	It("should template correctly with public attributes", func() {
		template, _ := ParseTemplate("{{ .GetInteger }}")
		str, err := template.Template(&testObj)
		Expect(err).To(BeNil())
		Expect(str).To(Equal("-20"))
	})
})