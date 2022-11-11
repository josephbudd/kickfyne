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

	"{{ .ImportPrefix }}/backend/txrx"
	"{{ .ImportPrefix }}/backend/store"
)

// Start starts the backend.
func Start(ctx context.Context, ctxCancel context.CancelFunc) (stores *store.Stores, err error) {
	// Opens the stores.
	if stores, err = store.New(); err != nil {
		return
	}
	if err = stores.Open(); err != nil {
		return
	}
	// Receive messages from the front-end.
	txrx.StartReceiver(ctx, ctxCancel, stores)
	return
}

`
