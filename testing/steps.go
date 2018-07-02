package testing

import (
	"github.com/sonaak/centurions/core"
	"github.com/pkg/errors"
	"sync"
)

type Recorder struct {
	store map[string]interface{}
	lock sync.RWMutex
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


type MockStep struct {
	recorder Recorder

	MockStatus core.Status
	CancelResult bool
}

func (step *MockStep) Run() core.Status {
	step.recorder.Set("run", true)
	return step.MockStatus
}

func (step *MockStep) Cancel() bool {
	step.recorder.Set("cancel", true)
	return step.CancelResult
}
