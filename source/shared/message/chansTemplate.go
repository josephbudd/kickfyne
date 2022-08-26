package message

const (
	chansFileName = "chans.go"
)

var chansTemplate = `package message

type MSGer interface {
	ID() (id uint64)
	MSG() (msg interface{})
	Name() (name string)
	FatalError() (fatal bool, message string)
}

var FrontEndToBackEnd = make(chan MSGer, 255)
var BackEndToFrontEnd = make(chan MSGer, 255)
var messageID uint64

func NextID() (id uint64) {
	id = messageID
	messageID++
	return
}

func IsValidID(id uint64) (isvalid bool) {
	isvalid = (id < messageID)
	return
}

`
