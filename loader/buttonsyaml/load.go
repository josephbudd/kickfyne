package buttonsyaml

import (
	"fmt"
	"os"

	"github.com/josephbudd/kickfyne/source/utils"
	"gopkg.in/yaml.v3"
)

// Load reads a yaml file and returns the ManifestHome and error.
func Load(fpath string, folderPaths *utils.FolderPaths) (buttonsYAML YAML, isOk bool, failureMessage string, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("panelloader.Read: %w", err)
		}
	}()

	var content []byte
	if content, err = os.ReadFile(fpath); err != nil {
		if os.IsNotExist(err) {
			failureMessage = fmt.Sprintf("%q does not exist", fpath)
			isOk = false
			err = nil
		}
		return
	}
	if err = yaml.Unmarshal(content, &buttonsYAML); err != nil {
		return
	}
	isOk, failureMessage, err = Check(buttonsYAML, folderPaths)
	return
}
