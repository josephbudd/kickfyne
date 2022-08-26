package message

type messageTemplateData struct {
	MessageName string
}

var messageTemplate = `package message

var {{ .MessageName }}ID = NextID()

type {{ .MessageName }} struct {
	id      uint64
	name    string
	GroupID uint64 // both ways

	/* KICKFYNE TODO:
	Complete this {{ .MessageName }} struct definition.
	*/

	Error        bool   // to front
	Fatal        bool   // to front
	ErrorMessage string // to front
}

// Build{{ .MessageName }} constructs a new New {{ .MessageName }} message.
func Build{{ .MessageName }}(groupID uint64) (msg {{ .MessageName }}) {
	msg = {{ .MessageName }}{
		id:                  {{ .MessageName }}ID,
		name:                "{{ .MessageName }}",
		GroupID:             groupID,

		/* KICKFYNE TODO:
		Complete this {{ .MessageName }} constructor if needed.
		*/
	}
	return
}

// {{ .MessageName }} implements the MSGer interface with ID and MSG and FatalError.

// ID returns the message's id
func (msg *{{ .MessageName }}) ID() (id uint64) {
	id = msg.id
	return
}

// Name returns the message's id
func (msg *{{ .MessageName }}) Name() (name string) {
	name = msg.name
	return
}

// Message returns the message's id
func (msg *{{ .MessageName }}) MSG() (m interface{}) {
	m = msg
	return
}

// IsFatal return if there was a fatal error and it's message.
func (msg *{{ .MessageName }}) FatalError() (fatal bool, message string) {
	fatal = msg.Fatal
	message = msg.ErrorMessage
	return
}

`
