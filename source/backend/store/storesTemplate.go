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
{{ $lenRecords := len .RecordNames }}
{{ $haveRecords := ne $lenRecords 0 -}}
{{ $recordNamesPadded := call .Funcs.PadSlice .RecordNames -}}
{{ $recordNamesColonPadded := call .Funcs.SuffixPadSlice .RecordNames ":" -}}
package store

import (
	"fmt"
	"strings"
	"sync"
	"time"
{{- if $haveRecords }}

	"{{ .ImportPrefix }}/backend/store/storing"
{{- end }}
)

var (
	_stores                 *Stores
	storesErr               error
	ErrStoresAlreadyClosing = fmt.Errorf("the stores are already waiting to close")
)

type storeState uint64

const (
	stateOpening storeState = 1 << iota
	stateClosing
	shortDuration = time.Second / time.Duration(2)
)

// Stores is each of the applicatioinlication's storers.
type Stores struct {
{{- if $haveRecords }}

	// Storers.
 {{- range $i, $padded := $recordNamesPadded }}
	{{ $padded }} *storing.{{ index $DOT.RecordNames $i }}Store
 {{- end }}

{{- end }}
	closeLockCount int
	lock         sync.Mutex
	closeLock    sync.Mutex
	state        storeState
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

	_stores = &Stores{}
{{- range $recordName := .RecordNames }}
	var {{ call $DOT.Funcs.DeCap $recordName }}Store *storing.{{ $recordName }}Store
	if {{ call $DOT.Funcs.DeCap $recordName }}Store, err = storing.New{{ $recordName }}Store(_stores.BlockClose, _stores.UnBlockClose); err != nil {
		return
	}
	_stores.{{ $recordName }} = {{ call $DOT.Funcs.DeCap $recordName }}Store
{{- end }}
	stores = _stores
	return
}

// BlockClose signals that the stores cannot be closed until the transaction ends.
func (stores *Stores) BlockClose() {
	stores.lock.Lock()
	if stores.closeLockCount++; stores.closeLockCount == 1 {
		stores.closeLock.Lock()
	}
	stores.lock.Unlock()
}

// UnBlockClose signals that the transaction has ended and the store may be closed if required.
func (stores *Stores) UnBlockClose() {
	stores.lock.Lock()
	defer stores.lock.Unlock()

	if stores.closeLockCount--; stores.closeLockCount > 0 {
		return
	}
	if stores.closeLockCount < 0 {
		stores.closeLockCount = 0
		return
	}
	stores.closeLock.Unlock()
}

// removeAllCloseBlocks signals that the transaction has ended and the store may be closed if required.
func (stores *Stores) removeAllCloseBlocks() (err error) {
	if stores.closeLockCount == 0 {
		// Transactions are already ended.
		return
	}
	stores.closeLockCount = 0
	if stores.state&stateClosing == stateClosing {
		stores.closeLock.Unlock()
	} else {
		err = stores.close()
	}
	return
}

// Open opens every store.
// It returns all of the errors as one single error.
func (stores *Stores) Open() (err error) {
	if err = storesErr; err != nil {
		return
	}
	defer func() {
		if err != nil {
			err = fmt.Errorf("stores.Open: %w", err)
		}
	}()

	stores.lock.Lock()
	stores.open()
	stores.lock.Unlock()
	err = storesErr
	return
}

// open opens every store.
// It sets storesErr to the open error.
// Uses no locks.
func (stores *Stores) open() {
	if storesErr != nil {
		return
	}
	errList := make([]string, 0, 10)
	stores.state |= stateOpening
	defer func() {
		stores.state &= ^stateOpening
		if len(errList) > 0 {
			msg := strings.Join(errList, "\n")
			storesErr = fmt.Errorf("stores.open: %s", msg)
		}
	}()
{{- if $haveRecords }}

	// Open each store.
	var err error
 {{- range $name := .RecordNames }}
	if err = stores.{{ $name }}.Open(); err != nil {
		errList = append(errList, err.Error())
	}
 {{- end }}
{{- end }}
}

func (stores *Stores) Close() (err error) {

	defer func() {
		stores.state &= ^stateClosing
		if err != nil {
			err = fmt.Errorf("stores.Close: %w", err)
		}
	}()

	if stores.state&^stateClosing == stateClosing {
		// Already waiting to close.
		err = ErrStoresAlreadyClosing
		return
	}
	stores.BlockClose()
	defer stores.UnBlockClose()
	stores.state |= stateClosing
	err = stores.close()
	return
}

// CloseTimeout closes every store.
// If a closer is blocked then it unblocks the close after waiting the param seconds.
func (stores *Stores) CloseTimeout(seconds int) (err error) {
	var timer *time.Timer

	defer func() {
		if err != nil {
			err = fmt.Errorf("stores.CloseTimeout: %w", err)
		}
	}()

	if stores.state&stateClosing != stateClosing {
		// There is no closer waiting to close, so close now.
		err = stores.Close()
		return
	}

	// There is a blocked closer waiting.
	if seconds == 0 {
		// Force a close now.
		// Unblock the waiting closer now, allowing it to close.
		stores.closeLock.Unlock()
		return
	}
	// Wait to close but check periodically.
	longDuration := time.Second * time.Duration(seconds)
	var accum time.Duration
	timer = time.NewTimer(shortDuration)
	for accum < longDuration {
		<-timer.C
		if stores.state&stateClosing != stateClosing {
			break
		}
		timer.Reset(shortDuration)
		accum += shortDuration
	}
	// End all closeLockCount and Unblock the waiting closer now, allowing it to close.
	err = stores.removeAllCloseBlocks()
	return
}

// close closes every store.
// It returns all of the errors as one single error.
// stores.Close is called by main.waitAndClose when the application shuts down.
// Uses no locks.
func (stores *Stores) close() (err error) {
	if err = storesErr; err != nil {
		return
	}

	errList := make([]string, 0, 4)
	defer func() {
		if len(errList) > 0 {
			msg := strings.Join(errList, "\n")
			err = fmt.Errorf("stores.close: %s", msg)
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
