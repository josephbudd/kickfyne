package message

const (
	initFileName = "init.go"
)

var initTemplate = `package message

var InitID = NextID()

type Init struct {
	id           uint64
	name         string
	Message      string // to front
	Error        bool   // to front
	Fatal        bool   // to front
	ErrorMessage string // to front
}

// NewInit constructs a new NewInit message.
func NewInit() (msg *Init) {
	msg = &Init{
		id:   InitID,
		name: "Init",
	}
	return
}

// Init implements the MSGer interface with ID and AsInterface.

// ID returns the message's id.
func (msg *Init) ID() (id uint64) {
	id = msg.id
	return
}

// Name returns the message's name.
func (msg *Init) Name() (name string) {
	name = msg.name
	return
}

// FatalError returns if there was a fatal error and it's message.
func (msg *Init) FatalError() (fatal bool, message, screenPackage string) {
	fatal = msg.Fatal
	message = msg.ErrorMessage
	// Init has no screen package because its sent by the front-end not a screen.
	return
}

// AsInterface returns msg as an interface{}.
func (msg *Init) AsInterface() (m interface{}) {
	m = msg
	return
}

`
