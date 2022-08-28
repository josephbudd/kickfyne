package storer

type templateData struct {
	ImportPrefix    string
	RecordName      string
	StoringFilePath string
}

var template = `package storer

import (
	"{{ .ImportPrefix }}/shared/store/record"
)

/* KICKFYNE NOTE WELL:
This defines the interface which is implemented by {{ .RecordName }}Store in {{ .StoringFilePath }}.
There fore, any changes made here must be reflected in {{ .StoringFilePath }}.
*/

// {{ .RecordName }}Storer defines the behavior (API) of a store of /shared/store/record.{{ .RecordName }} records.
type {{ .RecordName }}Storer interface {

	// Open opens the data-store.
	// Keeps the file in memory.
	// Returns the error.
	Open() (err error)

	// Close closes the data-store.
	// Returns the error.
	Close() (err error)

	// IsFull returns if the store is full.
	IsFull() (isFull bool)

	// NextID returns the next available ID or an error.
	NextID() (nextID uint64, err error)

	// Get retrieves one record.{{ .RecordName }} from the data-store.
	// Param id is the record ID.
	// Returns a record.{{ .RecordName }} and error.
	// When no record is found, the returned record.{{ .RecordName }} is nil and the returned error is nil.
	Get(id uint64) (r record.{{ .RecordName }}, err error)

	// GetAll retrieves all of the record.{{ .RecordName }} records from the data-store.
	// Returns a slice of record.{{ .RecordName }} and error.
	// When no records are found, the returned slice length is 0 and the returned error is nil.
	GetAll() (rr []record.{{ .RecordName }}, err error)

	// Update updates the record.{{ .RecordName }} in the data-store.
	// Param newR is the record.{{ .RecordName }} to be updated.
	// If newR is a new record then r.ID is updated as well.
	// Returns the error.
	Update(newR record.{{ .RecordName }}) (updatedR record.{{ .RecordName }}, err error)

	// UpdateAll updates a slice of record.{{ .RecordName }} in the data-store.
	// Param newRR is the slice of record.{{ .RecordName }} to be updated.
	// If any record in newRR is new then it's ID is updated as well.
	// Returns the error.
	UpdateAll(newRR []record.{{ .RecordName }}) (updatedRR []record.{{ .RecordName }}, err error)

	// Remove removes the record.{{ .RecordName }} from the data-store.
	// Param id is the record ID of the record.{{ .RecordName }} to be removed.
	// If the record is not found returns a nil error.
	// Returns the error.
	Remove(id uint64) (err error)
}

`
