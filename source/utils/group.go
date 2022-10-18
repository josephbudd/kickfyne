package utils

const (
	groupNameSep     = ":"
	LandingGroupName = "landing"
)

func GroupName(parentGroupName, childPackageName string) (groupName string) {
	groupName = parentGroupName + groupNameSep + childPackageName
	return
}
