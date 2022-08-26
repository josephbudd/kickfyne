package storing

const (
	stateFileName = "state.go"

	stateTemplate = `package storing

	import (
		"bytes"
		"fmt"
		"sync"
	
		"fyne.io/fyne/v2"
		"fyne.io/fyne/v2/storage"
		"gopkg.in/yaml.v3"
	
		"{{ .ImportPrefix }}/shared/paths"
		"{{ .ImportPrefix }}/shared/store/record"
	)
	
	type StateData struct {
		LastID  uint64
		Records []*record.State
	}
	
	// StateStore is the API of the State store.
	// It is the implementation of the interface in /domain/store/storer/State.go.
	type StateStore struct {
		uri  fyne.URI
		lock sync.Mutex
		data StateData
	}
	
	// NewStateStore constructs a new StateStore.
	// Param db is an open bolt data-store.
	// Returns a pointer to the new StateStore.
	func NewStateStore() (store *StateStore) {
		store = &StateStore{
			uri: paths.StoreURI("state.yaml"),
		}
		return
	}
	
	// Open opens the bolt data-store.
	// Keeps the file in memory.
	// Returns the error.
	func (store *StateStore) Open() (err error) {
	
		defer func() {
			if err != nil {
				err = fmt.Errorf("StateStore.Open: %w", err)
			}
		}()
	
		store.lock.Lock()
		defer store.lock.Unlock()
	
		err = store.readAll()
		return
	}
	
	// Close closes the data-store.
	// Returns the error.
	func (store *StateStore) Close() (err error) {
	
		defer func() {
			if err != nil {
				err = fmt.Errorf("StateStore.Close: %w", err)
			}
		}()
	
		// The store is always closed.
		return
	}
	
	// Get retrieves one *record.State from the data-store.
	// Param id is the record ID.
	// Returns a *record.State and error.
	// When no record is found, the returned *record.State is nil and the returned error is nil.
	func (store *StateStore) Get(id uint64) (r *record.State, err error) {
	
		defer func() {
			if err != nil {
				err = fmt.Errorf("StateStore.Get: %w", err)
			}
		}()
	
		store.lock.Lock()
		defer store.lock.Unlock()
	
		for _, r = range store.data.Records {
			if r.ID == id {
				return
			}
		}
		// Not found. No error.
		r = nil
		return
	}
	
	// GetAll retrieves all of the *record.State records from the data-store.
	// Returns a slice of *record.State and error.
	// When no records are found, the returned slice length is 0 and the returned error is nil.
	func (store *StateStore) GetAll() (rr []*record.State, err error) {
	
		defer func() {
			if err != nil {
				err = fmt.Errorf("StateStore.GetAll: %w", err)
			}
		}()
	
		store.lock.Lock()
		defer store.lock.Unlock()
	
		rr = store.data.Records
		return
	}
	
	// Update updates the *record.State in the data-store.
	// Param newR is the *record.State to be updated.
	// If newR is a new record then r.ID is updated as well.
	// Returns the error.
	func (store *StateStore) Update(newR *record.State) (err error) {
	
		defer func() {
			if err != nil {
				err = fmt.Errorf("StateStore.Update: %w", err)
			}
		}()
	
		store.lock.Lock()
		defer store.lock.Unlock()
	
		if len(store.data.Records) == 0 {
			newR.ID = 1
			store.data.Records = append(store.data.Records, newR)
		} else {
			if newR.ID != 1 {
				err = fmt.Errorf("invalid record id:%d", newR.ID)
				return
			}
			store.data.Records[0] = newR
		}
	
		// Write the file backend.
		err = store.writeAll()
		return
	}
	
	// UpdateAll updates a slice of *record.State in the data-store.
	// Param newRR is the slice of *record.State to be updated.
	// If any record in newRR is new then it's ID is updated as well.
	// Returns the error.
	func (store *StateStore) UpdateAll(newRR []*record.State) (err error) {
	
		defer func() {
			if err != nil {
				err = fmt.Errorf("StateStore.UpdateAll: %w", err)
			}
		}()
	
		err = fmt.Errorf("deprecated use Update")
		return
	}
	
	// Remove removes the record.State from the data-store.
	// Param id is the record ID of the record.State to be removed.
	// If the record is not found returns a nil error.
	// Returns the error.
	func (store *StateStore) Remove(id uint64) (err error) {
	
		defer func() {
			if err != nil {
				err = fmt.Errorf("StateStore.Remove: %w", err)
			}
		}()
	
		err = fmt.Errorf("can not remove the state record")
		return
	}
	
	func (store *StateStore) readAll() (err error) {
	
		defer func() {
			if err != nil {
				err = fmt.Errorf("StateStore.readAll: %w", err)
			}
		}()
	
		// If the file doesn't exists then setup the data.
		var exists bool
		if exists, err = storage.Exists(store.uri); err != nil {
			return
		}
		if !exists {
			store.data.Records = make([]*record.State, 0, 1024)
			return
		}
	
		// Open.
		var rc fyne.URIReadCloser
		if rc, err = storage.Reader(store.uri); err != nil {
			return
		}
		defer func() {
			closeErr := rc.Close()
			if err == nil {
				err = closeErr
			}
		}()
	
		// Read.
		buffer := bytes.Buffer{}
		if _, err = buffer.ReadFrom(rc); err != nil {
			return
		}
		err = yaml.Unmarshal(buffer.Bytes(), &store.data)
		return
	}
	
	func (store *StateStore) writeAll() (err error) {
	
		defer func() {
			if err != nil {
				err = fmt.Errorf("StateStore.writeAll: %w", err)
			}
		}()
	
		// Open.
		var wc fyne.URIWriteCloser
		if wc, err = storage.Writer(store.uri); err != nil {
			return
		}
		defer func() {
			closeErr := wc.Close()
			if err == nil {
				err = closeErr
			}
		}()
	
		// Convert.
		var bb []byte
		if bb, err = yaml.Marshal(&store.data); err != nil {
			return
		}
	
		// Write.
		_, err = wc.Write(bb)
		return
	}`
)

type stateTemplateData struct {
	ImportPrefix string
}
