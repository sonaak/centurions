package testing

import (
	"github.com/pkg/errors"
	"github.com/sonaak/centurions/core"
	"sync"
)

type Recorder struct {
	store map[string]interface{}
	lock  sync.RWMutex
}

func NewRecorder() *Recorder {
	return &Recorder{
		store: make(map[string]interface{}),
	}
}

func (recorder *Recorder) Get(key string) (value interface{}, err error) {
	recorder.lock.RLock()
	defer recorder.lock.RUnlock()

	value, in := recorder.store[key]
	if !in {
		err = errors.Errorf("no key %s in recorder", key)
	}

	return
}

func (recorder *Recorder) Set(key string, value interface{}) {
	recorder.lock.Lock()
	defer recorder.lock.Unlock()

	recorder.store[key] = value
}

const (
	run    = "run"
	cancel = "cancel"
)

type MockStep struct {
	recorder *Recorder

	MockStatus   core.Status
	CancelResult bool
}

func NewMockStep(cancelResult bool, mockStatus core.Status) *MockStep {
	return &MockStep{
		recorder:     NewRecorder(),
		MockStatus:   mockStatus,
		CancelResult: cancelResult,
	}
}

func (step *MockStep) Run() core.Status {
	step.recorder.Set(run, true)
	return step.MockStatus
}

func (step *MockStep) Cancel() bool {
	step.recorder.Set(cancel, true)
	return step.CancelResult
}

func (step *MockStep) hasDone(fName string) bool {
	val, err := step.recorder.Get(fName)
	if err != nil {
		return false
	}

	hasDone, ok := val.(bool)
	return ok && hasDone
}

func (step *MockStep) HasRun() bool {
	return step.hasDone(run)
}

func (step *MockStep) HasCanceled() bool {
	return step.hasDone(cancel)
}
