package message

type messageTemplateData struct {
	MessageName string
}

var messageTemplate = `package message

var {{ .MessageName }}ID = NextID()

type {{ .MessageName }} struct {
	id      uint64
	name    string
	ScreenPackage string // both ways

	/* KICKFYNE TODO:
	Complete this {{ .MessageName }} struct definition.
	Add the members that you want this message to have.
	*/

	Error        bool   // to front
	Fatal        bool   // to front
	ErrorMessage string // to front
}

// New{{ .MessageName }} returns a *{{ .MessageName }} message.
func New{{ .MessageName }}(screenPackage string) (msg *{{ .MessageName }}) {
	msg = &{{ .MessageName }}{
		id:            {{ .MessageName }}ID,
		name:          "{{ .MessageName }}",
		ScreenPackage: screenPackage,

		/* KICKFYNE TODO:
		Complete New{{ .MessageName }} as needed.
		*/
	}
	return
}

// {{ .MessageName }} implements the MSGer interface with ID and AsInterface and FatalError.

// ID returns the message's id
func (msg *{{ .MessageName }}) ID() (id uint64) {
	id = msg.id
	return
}

// Name returns the message's Name.
func (msg *{{ .MessageName }}) Name() (name string) {
	name = msg.name
	return
}

// AsInterface returns the message as an interface{}.
func (msg *{{ .MessageName }}) AsInterface() (m interface{}) {
	m = msg
	return
}

// FatalError return if there was a fatal error and it's message.
func (msg *{{ .MessageName }}) FatalError() (fatal bool, message, screenPackage string) {
	fatal = msg.Fatal
	message = msg.ErrorMessage
	screenPackage = msg.ScreenPackage
	return
}

`
