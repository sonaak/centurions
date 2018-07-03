package testing

import (
	"testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sonaak/centurions/core"
)

func TestRecorder_Get(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestRecorder Get")
}

var _ = Describe("TestRecorder Get", func() {
	var recorder *Recorder

	BeforeEach(func() {
		recorder = NewRecorder()
	})

	It("should be able to retrieve 1 uint without distorting type", func() {
		recorder.Set("a", uint(1))
		val, err := recorder.Get("a")
		Expect(err).To(BeNil())
		Expect(val).To(Equal(uint(1)))
	})

	It("should be able to retrieve string without distorting type", func() {
		recorder.Set("a", "a string value")
		val, err := recorder.Get("a")
		Expect(err).To(BeNil())
		Expect(val).To(Equal("a string value"))
	})

	It("should be able to retrieve booleans without distorting type", func() {
		recorder.Set("True", true)
		recorder.Set("False", false)
		val, err := recorder.Get("True")
		Expect(err).To(BeNil())
		Expect(val).To(Equal(true))

		val, err = recorder.Get("False")
		Expect(err).To(BeNil())
		Expect(val).To(Equal(false))
	})

	It("should respond with error when there is no such key", func(){
		recorder.Set("foo", 1)

		val, err := recorder.Get("bar")
		Expect(err).ToNot(BeNil())
		Expect(val).To(BeNil())
	})
})

func TestMockStep_Run(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestMockStep Run")
}

var _ = Describe("TestMockStep Run", func(){
	It("should record the attempt", func(){
		step := NewMockStep(true, core.PASSED)
		_ = step.Run()
		Expect(step.HasRun()).To(Equal(true))
	})

	It("should respond with the correct status", func(){
		step := NewMockStep(true, core.PASSED)
		Expect(step.Run()).To(Equal(core.PASSED))
	})
})