package message

const (
	chansFileName = "chans.go"
)

var chansTemplate = `package message

type MSGer interface {
	ID() (id uint64)
	Name() (name string)
	FatalError() (fatal bool, message, screenPackage string)
	AsInterface() (msg interface{})
}

var FrontEndToBackEnd = make(chan MSGer, 255)
var BackEndToFrontEnd = make(chan MSGer, 255)
var messageID = uint64(1)

func NextID() (id uint64) {
	id = messageID
	messageID++
	return
}

func IsValidID(id uint64) (isvalid bool) {
	isvalid = (id > 0 && id < messageID)
	return
}

`
