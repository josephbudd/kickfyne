package utils

import (
	"fmt"
	"path/filepath"
	"strings"
)

const (
	groupNameSep  = ":"
	HomeGroupName = "home"
)

// validateGroupName returns is the name is valid.
// userMessage contains the error message for the user.
func validateGroupName(groupName string) (isValid bool, userMessage string) {
	lines := make([]string, 0, len(groupName))
	var msg string
	isValid = true
	if !strings.Contains(lcAlphabet, groupName[:1]) {
		msg = "A package name must begin with a lower case letter."
		lines = append(lines, msg)
		isValid = false
	}
	l := len(groupName)
	isValid = true
	for i := 1; i < l; i++ {
		ch := groupName[i : i+1]
		switch {
		case strings.Contains(lcAlphabet, ch):
			continue
		case strings.Contains(ucAlphabet, ch):
			continue
		case strings.Contains(digits, ch):
			continue
		default:
			isValid = false
			msg = fmt.Sprintf(`A package name must not contains the character %q.`, ch)
			lines = append(lines, msg)
		}
	}
	if !isValid {
		userMessage = strings.Join(lines, "\n")
	}
	return
}

func GroupName(parentGroupName, childPackageName string) (groupName string) {
	groupName = parentGroupName + groupNameSep + childPackageName
	return
}

func GroupNamePath(groupName string) (path string) {
	names := strings.Split(groupName, groupNameSep)
	path = filepath.Join(names...)
	return
}
