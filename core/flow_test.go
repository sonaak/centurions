package core

import (
	"testing"
	. "github.com/onsi/gomega"
	. "github.com/onsi/ginkgo"
)

func TestFlowNode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FlowNode")
}

var _ = Describe("FlowNode", func(){
	It("should instantiate correctly with NewFlowNode with all nils", func(){
		flowNode := NewFlowNode(nil, nil, nil, nil)
		Expect(flowNode).NotTo(BeNil())
	})
})

func TestFlowNodeStatus(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FlowNodeStatus")
}

var _ = Describe("FlowNodeStatus", func(){
	It("should instantiate correctly with NewFlowNodeStatus", func(){
		Expect(NewFlowNodeStatus(FAILED, nil)).NotTo(BeNil())
	})
})

func TestFlowState(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FlowState")
}

var _ = Describe("FlowState", func(){
	It("should instantiate correctly with NewFlowState", func(){
		criteria := make([]FlowNodeStatus, 0, 0)
		Expect(NewFlowState(criteria)).NotTo(BeNil())
	})
})
