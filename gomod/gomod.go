package gomod

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

const (
	fileName = "go.mod"
)

func Read(pathWD string) (module string, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("gomod.Read: %w", err)
			return
		}
	}()

	var content []byte
	modGoPath := filepath.Join(pathWD, fileName)
	if content, err = ioutil.ReadFile(modGoPath); err != nil {
		fmt.Printf("Unable to find a %q file.\n", fileName)
		fmt.Println(`You need to run the command "go mod init" in this folder.`)
		return
	}
	lines := bytes.Split(content, []byte("\n"))
	if len(lines) == 0 {
		err = fmt.Errorf("bad lines")
		return
	}
	line := strings.TrimSpace(string(lines[0]))
	parts := strings.Split(line, " ")
	if parts[0] != "module" {
		err = fmt.Errorf("no module")
		return
	}
	if len(parts) != 2 {
		err = fmt.Errorf("bad parts")
		return
	}
	module = parts[1]
	return
}
