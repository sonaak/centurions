package testing

import (
	"testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRecorder_Get(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestRecorder Get")
}

var _ = Describe("TestRecorder Get", func() {
	It("should be able to retrieve a value without distorting type", func() {
		recorder := Recorder {
			store: make(map[string]interface{}),
		}

		recorder.Set("a", uint(1))
		val, err := recorder.Get("a")
		Expect(err).To(BeNil())
		Expect(val).To(Equal(uint(1)))
	})
})
