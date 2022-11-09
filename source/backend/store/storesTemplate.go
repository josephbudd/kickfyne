package store

import "github.com/josephbudd/kickfyne/source/utils"

const (
	storesFileName   = "stores.go"
	storesFolderName = "stores"
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
	"strings"{{ if $haveRecords }}

	"{{ .ImportPrefix }}/shared/store/storing"{{ end }}
)

// Stores is each of the application's storers.
type Stores struct {
	// Local yaml stores.
{{- range $i, $padded := $recordNamesPadded }}
	{{ $padded }} *storing.{{ index $DOT.RecordNames $i }}Store
{{- end }}
}

var _stores *Stores

// New constructs a new Stores.
func New() (stores *Stores) {
	if _stores != nil {
		stores = _stores
		return
	}
	_stores = &Stores{
{{- range $i, $colonPadded := $recordNamesColonPadded }}
		{{ $colonPadded }} storing.New{{ index $DOT.RecordNames $i }}Store(),
{{- end }}
	}
	stores = _stores
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
func (stores *Stores) Close() (err error) {

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
