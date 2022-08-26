package panel

const (
	groupIDFileName = "groupid.go"
)

var groupIDTemplate = `package panel

import "sync"

var groupid uint64
var lock sync.Mutex

func NextGroupID() (id uint64) {
	lock.Lock()
	defer lock.Unlock()

	id = groupid
	groupid++
	return
}

`
