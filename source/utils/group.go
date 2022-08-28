package utils

import (
	"path/filepath"
	"strings"
)

const (
	groupNameSep  = ":"
	HomeGroupName = "home"
)

func GroupName(parentGroupName, childPackageName string) (groupName string) {
	groupName = parentGroupName + groupNameSep + childPackageName
	return
}

func GroupNamePath(groupName string) (path string) {
	names := strings.Split(groupName, groupNameSep)
	path = filepath.Join(names...)
	return
}
