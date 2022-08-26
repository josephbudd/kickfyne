package storer

const (
	stateFileName = "state.go"

	stateTemplate = `package storer

import (
	"{{ .ImportPrefix }}/shared/store/record"
)

// StateStorer defines the behavior (API) of a store of /shared/store/record.State records.
type StateStorer interface {

	// Open opens the bolt data-store.
	// Keeps the file in memory.
	// Returns the error.
	Open() (err error)

	// Close closes the data-store.
	// Returns the error.
	Close() (err error)

	// Get retrieves one *record.State from the data-store.
	// Param id is the record ID.
	// Returns a *record.State and error.
	// When no record is found, the returned *record.State is nil and the returned error is nil.
	Get(id uint64) (r *record.State, err error)

	// GetAll retrieves all of the *record.State records from the data-store.
	// Returns a slice of *record.State and error.
	// When no records are found, the returned slice length is 0 and the returned error is nil.
	GetAll() (rr []*record.State, err error)

	// Update updates the *record.State in the data-store.
	// Param newR is the *record.State to be updated.
	// If newR is a new record then r.ID is updated as well.
	// Returns the error.
	Update(newR *record.State) (err error)

	// UpdateAll updates a slice of *record.State in the data-store.
	// Param newRR is the slice of *record.State to be updated.
	// If any record in newRR is new then it's ID is updated as well.
	// Returns the error.
	UpdateAll(newRR []*record.State) (err error)

	// Remove removes the record.State from the data-store.
	// Param id is the record ID of the record.State to be removed.
	// If the record is not found returns a nil error.
	// Returns the error.
	Remove(id uint64) (err error)
}

`
)

type stateTemplateData struct {
	ImportPrefix string
}
