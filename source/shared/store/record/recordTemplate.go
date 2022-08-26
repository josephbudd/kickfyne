package record

import "github.com/josephbudd/kickfyne/source/utils"

const (
	folderName = "record"
)

type templateData struct {
	RecordName string
	Funcs      utils.Funcs
}

var template = `package record

type {{ .RecordName }} struct {
	ID uint64

	/* KICKFYNE TODO:
	Complete this {{ .RecordName }} struct definition.
	*/
}

// Build{{ .RecordName }} builds a new {{ .RecordName }} record.
func Build{{ .RecordName }}() ({{ call .Funcs.DeCap .RecordName }} {{ .RecordName -}} ) {
	{{ call .Funcs.DeCap .RecordName }} = {{ .RecordName }}{
		/* KICKFYNE TODO:
		Complete this {{ .RecordName }} constructor if needed.
		*/
	}
	return
}

// IsZero returns if the record hasn't been given an id.
// A record from the store has an ID > 0.
// Use to check the record returned by {{ .RecordName }}Store.Get(id)
func ({{ call .Funcs.DeCap .RecordName }} {{ .RecordName }}) IsZero() (isZero bool) {
	isZero = {{ call .Funcs.DeCap .RecordName }}.ID == 0
	return
}

`
