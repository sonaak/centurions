package core


type Status uint8

const (
	RUNNING = Status(0)
	SKIPPED = Status(1)
	PASSED = Status(2)
	REDIRECT = Status(3)
	FAILED = Status(4)
	ERRORED = Status(5)
	CANCELED = Status(6)
	RETRYING = Status(7)
)

type Runnable interface {
	Run() Status
	Cancel() bool
}
