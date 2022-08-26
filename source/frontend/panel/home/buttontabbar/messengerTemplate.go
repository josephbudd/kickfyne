package buttontabbar

import "github.com/josephbudd/kickfyne/source/utils"

type messengerTemplateData struct {
	ImportPrefix   string
	PackageName    string
	HomeButtonName string
	GroupName      string
	Funcs          utils.Funcs
}

const (
	messengerFileName = "messageHandler.go"

	messengerTemplate = `package {{ .PackageName }}

	import (
		"fmt"
	
		"{{ .ImportPrefix }}/frontend/panel"
	)
	
	const (
		groupName = "{{ .GroupName }}"
	)
	
	var (
		messenger *messageHandler = &messageHandler{}
		groupID                   = panel.NextGroupID()
	)
	
	type messageHandler struct{}
	
	func (m *messageHandler) listen() (err error) {
	
		defer func() {
			if err != nil {
				err = fmt.Errorf("messageHandler.listen: %w", err)
			}
		}()
	
		/* KICKFYNE TODO:
		Add listeners for each message that is used by this package.
		The messages are in shared/message/ folder.
		Use a message's ID.
	
		Example:
		if err = txrx.AddListener(message.ContactForEditID, m); err != nil {
			return
		}
	
		*/
		return
	}
	
	// Listen listens for dispatched messages.
	// It is part of the txrx.Listener interface implementation.
	func (m *messageHandler) Listen(msg interface{}) {
	
		/* KICKFYNE TODO:
		Add a switch with cases for each message type and corresonding handler.
	
		Example:
		switch msg := msg.(type) {
		case *message.ContactAdd:
			m.contactAddRX(msg)
		}
	
		*/
	}
	
	// GroupName returns this panel group's name.
	// It is part of the txrx.Listener interface implementation.
	func (m *messageHandler) GroupName() (name string) {
		name = groupName
		return
	}
	
	/* KICKFYNE TODO:
	Add senders and receivers for each message used in this package.
	
	Example:
	// ContactAdd message.
	
	// contactAddTX sends the contact add message to the back end.
	func (m *messageHandler) contactAddTX(r *record.ContactAdd) {
		message.FrontEndToBackEnd <- message.NewContactAdd(groupID, r)
	}
	
	// contactAddRX receives the contact add message from the back end.
	func (m *messageHandler) contactAddRX(msg *message.ContactAdd) {
		if msg.GroupID != groupID {
			return
		}
		if msg.Error {
			dialog.ShowInformation("Error", msg.ErrorMessage, window)
			return
		}
		dialog.ShowInformation("Success", "Contact added.", window)
		fPanel.form.Clear()
	}
	
	*/

`
)
