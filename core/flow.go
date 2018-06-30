package core


type FlowNode struct {
	Step Runnable
	Redirect *FlowNode
	Retry *FlowNode
	DependsOn *FlowState
}


func NewFlowNode(step Runnable, redirect *FlowNode, retry *FlowNode, dependsOn *FlowState) *FlowNode {
	return &FlowNode {
		Step: step,
		Redirect: redirect,
		Retry: retry,
		DependsOn: dependsOn,
	}
}


type FlowNodeStatus struct {
	Status Status
	Node *FlowNode
}

func NewFlowNodeStatus(status Status, node *FlowNode) *FlowNodeStatus {
	return nil
}


type FlowState struct {
	Criteria []FlowNodeStatus
}


func NewFlowState(criteria []FlowNodeStatus) *FlowState {
	return nil
}


type Flow struct {
	Steps []FlowNode
	SuccessCriteria Criteria
}