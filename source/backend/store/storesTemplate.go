package store

import "github.com/josephbudd/kickfyne/source/utils"

const (
	storesFileName = "stores.go"
)

type storesTemplateData struct {
	ImportPrefix string
	RecordNames  []string
	Funcs        utils.Funcs
}

var storesTemplate = `{{ $DOT := . -}}
{{ $haveRecords := ne (len .RecordNames) 0 -}}
{{ $recordNamesPadded := call .Funcs.PadSlice .RecordNames -}}
{{ $recordNamesColonPadded := call .Funcs.SuffixPadSlice .RecordNames ":" -}}
package store

import (
	"fmt"
	"log"
	"strings"
{{- if $haveRecords }}

	"{{ .ImportPrefix }}/backend/store/storing"
{{- end }}
)

var (
	_stores   *Stores
	storesErr error
)

// Stores is each of the application's storers.
type Stores struct {
	// Local yaml stores.
{{- range $i, $padded := $recordNamesPadded }}
	{{ $padded }} *storing.{{ index $DOT.RecordNames $i }}Store
{{- end }}
}

// New constructs a new Stores.
func New() (stores *Stores, err error) {

	if err = storesErr; err != nil {
		return
	}
	if stores = _stores; stores != nil {
		return
	}

	defer func() {
		if err != nil {
			err = fmt.Errorf("store.New: %w", err)
			storesErr = err
		}
	}()

{{- range $recordName := .RecordNames }}
	var {{ call $DOT.Funcs.DeCap $recordName }}Store *storing.{{ $recordName }}Store
	if {{ call $DOT.Funcs.DeCap $recordName }}Store, err = storing.New{{ $recordName }}Store(); err != nil {
		return
	}
{{- end }}
	_stores = &Stores{
{{- range $i, $colonPadded := $recordNamesColonPadded }}
		{{ $colonPadded }} {{ call $DOT.Funcs.DeCap (index $DOT.RecordNames $i) }}Store,
{{- end }}
	}
	stores = _stores
	return
}

// Open opens every store.
// It returns all of the errors as one single error.
func (stores *Stores) Open() (err error) {

	if err = storesErr; err != nil {
		return
	}

	errList := make([]string, 0, 10)
	defer func() {
		if len(errList) > 0 {
			msg := strings.Join(errList, "\n")
			err = fmt.Errorf("stores.Open: %s", msg)
			storesErr = err
		}
	}()
{{- if $haveRecords }}

	// Local yaml stores.
 {{- range $name := .RecordNames }}
	if err = stores.{{ $name }}.Open(); err != nil {
		errList = append(errList, err.Error())
	}
 {{- end }}
{{- end }}

	return
}

// Close closes every store.
// It returns all of the errors as one single error.
// stores.Close is called by main.waitAndClose when the app shuts down.
func (stores *Stores) Close() (err error) {
	log.Println("closing stores")

	if err = storesErr; err != nil {
		return
	}

	errList := make([]string, 0, 4)
	defer func() {
		if len(errList) > 0 {
			msg := strings.Join(errList, "\n")
			err = fmt.Errorf("stores.Close: %s", msg)
		}
	}()
{{- if $haveRecords }}

	// Local yaml stores.
 {{- range $name := .RecordNames }}
	if err = stores.{{ $name }}.Close(); err != nil {
		errList = append(errList, err.Error())
	}
 {{- end }}
{{- end }}

	return
}

`
