package core

type Criteria interface {
	Evaluate() bool
}
