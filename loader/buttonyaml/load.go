package buttonyaml

import (
	"fmt"
	"os"

	"github.com/josephbudd/kickfyne/source/utils"
	"gopkg.in/yaml.v3"
)

// Load reads a yaml file and returns the ManifestHome and error.
func Load(fpath string, folderPaths *utils.FolderPaths) (buttonYAML YAML, isOK bool, failureMessage string, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("buttonyaml.Load: %w", err)
		}
	}()

	var content []byte
	if content, err = os.ReadFile(fpath); err != nil {
		if os.IsNotExist(err) {
			failureMessage = fmt.Sprintf("%q does not exist", fpath)
			err = nil
		}
		return
	}
	if err = yaml.Unmarshal(content, &buttonYAML); err != nil {
		return
	}
	isOK, failureMessage, err = Check(buttonYAML, folderPaths)
	return
}
