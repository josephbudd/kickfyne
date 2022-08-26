package shared

import "github.com/josephbudd/kickfyne/source/utils"

const (
	MainFileName = "main.go"
)

type mainTemplateData struct {
	ImportPrefix string
	AppName      string
	Funcs        utils.Funcs
}

var mainTemplate = `{{ $lCAppName := call .Funcs.LowerCase .AppName }}package main

import (
	"context"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"{{ .ImportPrefix }}/backend"
	"{{ .ImportPrefix }}/frontend"
)

const (
	envTrue  = "1"
	envFalse = ""
)

func main() {

	var err error
	defer func() {
		if err != nil {
			log.Printf("main: err is %s", err.Error())
			os.Exit(1)
		}
	}()

	if len(os.Getenv("FYNE_SCALE")) == 0 {
		os.Setenv("FYNE_SCALE", "1")
	}
	if len(os.Getenv("FYNE_THEME")) == 0 {
		os.Setenv("FYNE_THEME", "dark")
	}
	os.Setenv("USETESTPATH", envFalse)
	os.Setenv("CWT_TESTING", envFalse)

	a := app.New()
	w := a.NewWindow("{{ .AppName }}")

	// Cancel.
	ctx, ctxCancel := context.WithCancel(context.Background())
	w.SetCloseIntercept(
		ctxCancel,
	)
	errCh := make(chan error, 2)
	go monitor(w, ctx, ctxCancel, errCh)

	w.Show()
	size := size16x9(1000, 0)
	w.Resize(size)
	w.CenterOnScreen()

	// Start the back end.
	if err = backend.Start(ctx, ctxCancel, errCh); err != nil {
		return
	}

	// Start the front end.
	if err = frontend.Start(ctx, ctxCancel, a, w); err != nil {
		return
	}

	// Build the front end content for the window.
	var viewContent *fyne.Container
	if viewContent, err = frontend.Content(); err != nil {
		return
	}
	w.SetContent(viewContent)
	a.Run()
}

func monitor(w fyne.Window, ctx context.Context, ctxCancel context.CancelFunc, errCh chan error) {
	select {
	case <-ctx.Done():
		w.Close()
		os.Exit(0)
		return
	case err := <-errCh:
		log.Println(err)
		w.Close()
		os.Exit(1)
		return
	}
}

func size16x9(width, height int) (size fyne.Size) {
	var newWidth float32
	var newHeight float32
	switch {
	case width != 0:
		if width < 0 {
			width = 0 - width
		}
		r := width / 16
		newWidth = float32(r * 16)
		newHeight = float32(r * 9)
	case height != 0:
		if height < 0 {
			height = 0 - height
		}
		r := height / 9
		newWidth = float32(r * 16)
		newHeight = float32(r * 9)
	default:
		// default to 720 width.
		r := 720 / 16
		newWidth = float32(r * 16)
		newHeight = float32(r * 9)
	}
	size = fyne.Size{Width: newWidth, Height: newHeight}
	return
}

`
