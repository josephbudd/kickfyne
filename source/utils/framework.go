package utils

// NoFrameworkMessage builds a message for the user.
// The message tells the user to create the framework before attempting the userAction param.
func NoFrameworkMessage(userAction string) (message string) {
	message = "You must build the framework before " + userAction + ".\n" +
		"Use the following command.\n$ kickfyne framework"
	return
}
