package shared

type sharedTemplateData struct {
	ImportPrefix string
}

const (
	sharedFileName = "shared.go"

	sharedTemplate = `package shared

import (
	"context"
	"fmt"

	"{{ .ImportPrefix }}/shared/paths"
	"{{ .ImportPrefix }}/shared/store"
)

// Start starts the shared.
func Start(ctx context.Context, ctxCancel context.CancelFunc) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("shared.Start: %w", err)
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
	err = stores.Close()
	return
}
`
)
