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

// Init implements the MSGer interface with ID and MSG.

// ID returns the message's id
func (msg *Init) ID() (id uint64) {
	id = msg.id
	return
}

// Name returns the message's id
func (msg *Init) Name() (name string) {
	name = msg.name
	return
}

// Message returns the message's id
func (msg *Init) MSG() (m interface{}) {
	m = msg
	return
}

// IsFatal return if there was a fatal error and it's message.
func (msg *Init) FatalError() (fatal bool, message string) {
	fatal = msg.Fatal
	message = msg.ErrorMessage
	return
}

`
