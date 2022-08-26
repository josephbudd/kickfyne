package backend

const (
	fileName = "backend.go"
)

type templateData struct {
	ImportPrefix string
}

var template = `package backend

import (
	"context"
	"fmt"

	"{{ .ImportPrefix }}/backend/txrx"
	"{{ .ImportPrefix }}/shared/paths"
	"{{ .ImportPrefix }}/shared/store"
)

// Start starts the backend.
func Start(ctx context.Context, ctxCancel context.CancelFunc, errCh chan error) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("backend.Start: %w", err)
		}
	}()


	// App paths.
	if err = paths.Init(); err != nil {
		return
	}

	// Stores.
	// Open loads the store's records into memory.
	// Gets are read from the memory.
	// Updates are written to the memory. Then the store's file is opened for a write from memory and then closed.
	stores := store.New()
	if err = stores.Open(); err != nil {
		return
	}
	if err = stores.Close(); err != nil {
		return
	}

	// Messages.
	txrx.Listen(ctx, ctxCancel, stores)
	return
}

`
