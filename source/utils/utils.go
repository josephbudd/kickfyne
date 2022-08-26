package utils

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"github.com/josephbudd/kickfyne/gomod"
)

var (
	FMode = os.FileMode(0664)
	DMode = os.FileMode(0775)
)

// ProcessTemplate processes a template.
func ProcessTemplate(
	templateName string,
	writePath string,
	templateString string,
	templateParams interface{},
) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("utils.ProcessTemplate: %w", err)
		}
	}()

	bb := new(bytes.Buffer)
	t := template.Must(template.New(writePath).Parse(templateString))
	if err = t.Execute(bb, templateParams); err != nil {
		return
	}
	err = os.WriteFile(writePath, bb.Bytes(), FMode)
	return
}

// WriteFile writes a file.
func WriteFile(writePath string, content []byte) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("utils.WriteFile: %w", err)
		}
	}()

	err = os.WriteFile(writePath, content, FMode)
	return
}

func CopyFile(srcPath, dstPath string) (err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("utils.CopyFile: %w", err)
		}
	}()

	var content []byte
	if content, err = os.ReadFile(srcPath); err != nil {
		return
	}

	err = os.WriteFile(dstPath, content, 0644)
	return
}

// FolderHasFolders returns if childFolders are in parentFolder.
func FolderHasFolders(parentFolderPath string, requiredFolderNames ...string) (hasFolders bool, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("source.FolderHasFolders: %w", err)
		}
	}()

	var currentFolderNames []string
	if currentFolderNames, err = FolderNames(parentFolderPath); err != nil {
		return
	}

	// Check for required child folder names.
	var foundCount int
	for _, currentFolderName := range currentFolderNames {
		for _, requiredFolderName := range requiredFolderNames {
			if requiredFolderName == currentFolderName {
				foundCount++
				break
			}
		}
	}
	hasFolders = foundCount == len(requiredFolderNames)

	return
}

// FolderNames returns childFolder names.
func FolderNames(parentFolderPath string) (names []string, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("source.FolderNames(%q): %w", parentFolderPath, err)
		}
	}()

	var dirEntrys []os.DirEntry
	if dirEntrys, err = os.ReadDir(parentFolderPath); err != nil {
		return
	}
	names = make([]string, 0, len(dirEntrys))
	for _, dirEntry := range dirEntrys {
		if dirEntry.IsDir() {
			name := dirEntry.Name()
			names = append(names, name)
		}
	}

	return
}

// FileNames returns a folder's file names.
func FileNames(parentFolderPath string) (names []string, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("utils.FileNames(%q): %w", parentFolderPath, err)
		}
	}()

	var dirEntrys []os.DirEntry
	if dirEntrys, err = os.ReadDir(parentFolderPath); err != nil {
		return
	}
	names = make([]string, 0, len(dirEntrys))
	for _, dirEntry := range dirEntrys {
		if !dirEntry.IsDir() {
			name := dirEntry.Name()
			names = append(names, name)
		}
	}

	return
}

// gotBuilt returns if the framework was built in this folder.
// It does so by checking for 3 folders.
func gotBuilt(appPath string) (isBuilt bool, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("utils.isBuilt: %w", err)
		}
	}()

	isBuilt, err = FolderHasFolders(
		appPath,
		FolderNameFrontend, FolderNameBackend, FolderNameShared,
	)
	return
}

// Setup
func Setup(pathWD string, dumperCh chan string) (isBuilt bool, importPrefix string, folderPaths *FolderPaths, err error) {
	if isBuilt, err = gotBuilt(pathWD); err != nil {
		return
	}
	// Get the import prefix from the go.mod file.
	if importPrefix, err = gomod.Read(pathWD, dumperCh); err != nil {
		return
	}
	if folderPaths, err = BuildFolderPaths(pathWD); err != nil {
		return
	}
	return
}
