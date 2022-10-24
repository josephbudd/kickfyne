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
	"{{ .ImportPrefix }}/shared/store"
)

// Start starts the backend.
func Start(ctx context.Context, ctxCancel context.CancelFunc) {
	stores := store.New()
	// Messages.
	txrx.StartReceiver(ctx, ctxCancel, stores)
}

`
