package txrx

const (
	docFileName = "doc.go"
)

var docTemplate = `// Package txrx receives messages from the front end.
// It also sends messages to the front end.
// The dispatcher is func Listen.
// Every other file has a specific message handler.
package txrx

`
