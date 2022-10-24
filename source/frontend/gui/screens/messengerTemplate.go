package screens

type messengerTemplateData struct {
	PackageName  string
	ImportPrefix string
}

const (
	messengerFileName = "messageHandler.go"

	messengerTemplate = `package {{ .PackageName }}

import (
	"fmt"
)

const (
	packageName = "{{ .PackageName }}"
)

// type messageHandler communicates with the backend using messages.
// It is an implementation of the txrx.Receiver interface.
type messageHandler struct{
	screen *screenComponents
}

// newMessageHandler constructs this message handler.
func newMessageHandler(screen *screenComponents) (messenger *messageHandler) {
	messenger = &messageHandler{
		screen: screen,
	}
	messenger.startReceiving()
	return
}

// ScreenPackage returns this screen's package name.
// It is part of the txrx.Receiver interface implementation.
func (m *messageHandler) ScreenPackage() (name string) {
	name = packageName
	return
}

// startReceiving begins the receiving process.
// It adds this message receiver as a listener to messages from the back-end.
func (m *messageHandler) startReceiving() (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("messageHandler.listen: %w", err)
		}
	}()

	/* KICKFYNE TODO:
	Add this listener for each message that is used by this package.
	The messages are in {{ .ImportPrefix }}/shared/message/ folder.
	Use a message's ID.

	Example:
	if err = txrx.AddReceiver(
		m,
		message.GetSomethingID,
		message.DoSomethingID,
	); err != nil {
		return
	}

	*/

	return
}

// Receive receives dispatched messages.
// It is part of the txrx.Receiver interface implementation.
func (m *messageHandler) Receive(msg interface{}) {

	/* KICKFYNE TODO:
	Add a switch with cases for each message type and corresonding receiver.

	Example:
	switch msg := msg.(type) {
	case *message.GetSomething:
		m.receiveGetSomething(msg)
	case *message.DoSomething:
		m.receiveDoSomething(msg)
	}

	*/
}

/* KICKFYNE TODO:
Add send funcs for each message sent.
Add receiver funcs for each message received.

Example:
// GetSomething message.

// sendGetSomething sends a GetSomething to the back-end.
func (m *messageHandler) sendGetSomething(r *record.GetSomething) {
	msg := message.HelloTX(packageName, r)
	txrx.Send(msg)
}

// receiveGetSomething handles a received GetSomething message from the back-end.
func (m *messageHandler) receiveGetSomething(msg *message.GetSomething) {
	if msg.Error {
		if msg.ScreenPackage != packageName {
			// This message was sent by another screen.
			// That screen will deal with the error.
			return
		}
		// This screen sent the message so receive the error here.
		dialog.ShowInformation("Error", msg.ErrorMessage, m.screen.window)
		return
	}
	m.screen.panels.editPanel.FillForm(msg.Something)
	m.screen.panels.editPanel.show()
}

// DoSomething message.

// sendDoSomething sends an DoSomething message to the back-end.
func (m *messageHandler) sendDoSomething(r *record.ContactEdit) {
	msg := message.NewDoSomething(packageName, r)
	txrx.Send(msg)

}

// receiveDoSomething handles a received DoSomething message from the back-end.
func (m *messageHandler) receiveDoSomething(msg *message.DoSomething) {
	if msg.ScreenPackage != packageName {
		// This particular message is ignored here if not sent by this screen.
		return
	}
	// This screen sent the message so it will receive it.
	if msg.Error {
		dialog.ShowInformation("Error", msg.ErrorMessage, window)
		return
	}
	dialog.ShowInformation("Success", "Did something.", m.screen.window)
	// Go back to the select panel.
	m.screen.panels.selectPanel.show()
}

*/

`
)
