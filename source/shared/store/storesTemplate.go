package store

const (
	storesFileName   = "stores.go"
	storesFolderName = "stores"
)

type storesTemplateData struct {
	ImportPrefix           string
	RecordNamesPadded      map[string]string
	RecordNamesColonPadded map[string]string
}

var storesTemplate = `{{ $DOT := . }}{{ $haveRecords := ne (len .RecordNamesPadded) 0 }}package store

import (
	"fmt"
	"strings"{{ if $haveRecords }}

	"{{ .ImportPrefix }}/shared/store/storing"{{ end }}
)

// Stores is each of the application's storers.
type Stores struct {
	// Local yaml stores. {{- range $name, $padded := .RecordNamesPadded }}
	{{ $padded }} *storing.{{ $name }}Store{{ end }}
}

// New constructs a new Stores.
func New() (stores *Stores) {
	stores = &Stores{ {{- range $name, $colonPadded := .RecordNamesColonPadded }}
		{{ $colonPadded }} storing.New{{ $name }}Store(),{{ end }}
	}
	return
}

// Open opens every store.
// It returns all of the errors as one single error.
func (stores *Stores) Open() (err error) {

	errList := make([]string, 0, 10)
	defer func() {
		if len(errList) > 0 {
			msg := strings.Join(errList, "\n")
			err = fmt.Errorf("stores.Open: %s", msg)
		}
	}(){{ if $haveRecords }}

	// Local yaml stores.{{ range $name, $colonPadded := .RecordNamesColonPadded }}
	if err = stores.{{ $name }}.Open(); err != nil {
		errList = append(errList, err.Error())
	}{{ end }}{{ end }}

	return
}

// Close closes every store.
// It returns all of the errors as one single error.
func (stores *Stores) Close() (err error) {

	errList := make([]string, 0, 4)
	defer func() {
		if len(errList) > 0 {
			msg := strings.Join(errList, "\n")
			err = fmt.Errorf("stores.Close: %s", msg)
		}
	}(){{ if $haveRecords }}

	// Local yaml stores. {{- range $name, $colonPadded := .RecordNamesColonPadded }}
	if err = stores.{{ $name }}.Close(); err != nil {
		errList = append(errList, err.Error())
	}{{ end }}{{ end }}

	return
}

`
