package state

const (
	stateFileName = "state.go"

	stateTemplate = `package state

	import (
		"fmt"
		"sync"
	
		"github.com/josephbudd/okp/shared/store"
		"github.com/josephbudd/okp/shared/store/record"
	)
	
	type State struct {
		state             *record.State
		stores            *store.Stores

		/* KICKFYNE TODO:
		Complete this definition by adding your own elements that are part of state.

		Example:
		currentCourse     *record.Course
		finalLessonNumber uint64
	
		*/
	}
	
	var appState *State
	var lock sync.Mutex
	
	func lockState() {
		lock.Lock()
	}
	
	func unlockState() {
		lock.Unlock()
	}
	
	// Init initializes the backend bestate.
	// Called by func Start in backend.go.
	func Init(stores *store.Stores) (err error) {
		if appState != nil {
			return
		}
		appState = &State{
			stores: stores,
		}
		err = syncToStores()
		return
	}
	
	func syncToStores() (err error) {
	
		defer func() {
			if err != nil {
				err = fmt.Errorf("appState.syncToStores: %w", err)
			}
		}()
	
		lockState()
		defer unlockState()
	
		if appState.state, err = appState.stores.State.Get(1); err != nil {
			return
		}
		if appState.state == nil {
			err = fmt.Errorf("state not found in stores")
			return
		}

		/* KICKFYNE TODO:
		Add you own initialization of state from the stores.

		Example:
		if appState.currentCourse, err = appState.stores.Course.Get(appState.state.CurrentCourseID); err != nil {
			return
		}
		if appState.currentCourse == nil {
			err = fmt.Errorf("currentCourse not found in stores")
			return
		}
		var last int
		if last = len(appState.currentCourse.HomeWorks); last > 0 {
			last--
		}
		appState.finalLessonNumber = appState.currentCourse.HomeWorks[last].LessonNumber
		*/

		return
	}
	
	`
)

type stateTemplateData struct {
	ImportPrefix string
}
