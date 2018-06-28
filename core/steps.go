package core


type Status uint8


type Runnable interface {
	Run() Status
	Cancel() bool
}
