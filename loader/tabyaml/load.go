package tabyaml

import (
	"fmt"
	"os"

	"github.com/josephbudd/kickfyne/source/utils"
	"gopkg.in/yaml.v3"
)

// Load reads a yaml file and returns the ManifestHome and error.
func Load(fpath string, folderPaths *utils.FolderPaths) (tabYAML YAML, isOk bool, userMessage string, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("tab.Load: %w", err)
		}
	}()

	var content []byte
	if content, err = os.ReadFile(fpath); err != nil {
		if os.IsNotExist(err) {
			userMessage = fmt.Sprintf("%q does not exist", fpath)
			err = nil
		}
		return
	}
	if err = yaml.Unmarshal(content, &tabYAML); err != nil {
		return
	}
	isOk, userMessage, err = Check(tabYAML, folderPaths)
	return
}
