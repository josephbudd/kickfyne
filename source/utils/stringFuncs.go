package utils

import (
	"fmt"
	"strings"
)

// Base.

// comment splits a string into lines and comments each line.
// Returns commented lines joined back into a single string.
func comment(desc string) (comment string) {
	descs := strings.Split(desc, "\n")
	comments := make([]string, 0, len(descs))
	// Find the last line.
	var last int
	for last = len(descs) - 1; last >= 0; last-- {
		if d := descs[last]; len(d) > 0 {
			break
		}
	}
	// Comment each line.
	for i := 0; i <= last; i++ {
		d := strings.TrimRight(descs[i], asciiSpaces)
		c := fmt.Sprintf("// " + d)
		comments = append(comments, c)
	}
	comment = strings.Join(comments, "\n")
	return
}

// Cap capitalizes the first character of a string.
func Cap(name string) (capped string) {
	capped = strings.ToUpper(name[:1]) + name[1:]
	return
}

// DeCap un capitalizes the first character of a string.
func DeCap(name string) (decapped string) {
	decapped = strings.ToLower(name[:1]) + name[1:]
	return
}

// LabelToVarName converts a label to a valid varName
// Ex: "helllo world." to "helloWorld"
func LabelToVarName(label string) (varName string) {
	name := LabelToName(label)
	// Lower case camel.
	varName = strings.ToLower(name[:1]) + name[1:]
	return
}

// LabelToName converts a label to a valid name.
// Ex: "helllo world." to "HelloWorld"
// err is nil or contains the error message for the user.
func LabelToName(label string) (name string) {
	validNameChars := make([]string, 0, len(label))
	var started bool
	var followingSpace bool
	for i := range label {
		ch := label[i : i+1]
		switch {
		case strings.Contains(asciiSpaces, ch):
			followingSpace = true
			// No white spaces and punctuation allowed.
		case !started:
			// The first character must be a capital letter.
			switch {
			case strings.Contains(ucAlphabet, ch):
				started = true
				followingSpace = false
				validNameChars = append(validNameChars, ch)
			case strings.Contains(lcAlphabet, ch):
				started = true
				followingSpace = false
				validNameChars = append(validNameChars, strings.ToUpper(ch))
			}
		case started:
			// Only characters and numbers follow the first capitalized letter.
			switch {
			case strings.Contains(ucAlphabet, ch):
				validNameChars = append(validNameChars, ch)
				followingSpace = false
			case strings.Contains(lcAlphabet, ch):
				if followingSpace {
					// Enforce camel case.
					ch = strings.ToUpper(ch)
					followingSpace = false
				}
				validNameChars = append(validNameChars, ch)
			case strings.Contains(digits, ch):
				validNameChars = append(validNameChars, ch)
				followingSpace = false
			}
		}
	}
	name = strings.Join(validNameChars, "")
	return
}

// validateMessageName returns if the message name is valid.
// userMessage contains the error message for the user.
// Param name is the user created name.
func validateMessageName(name string) (isValid bool, userMessage string) {
	isValid, userMessage = validateCamelCaseName(name, "message")
	return
}

// validateRecordName returns if the record name is valid.
// userMessage contains the error message for the user.
// Param name is the user created name.
func validateRecordName(name string) (isValid bool, userMessage string) {
	isValid, userMessage = validateCamelCaseName(name, "record")
	return
}

// validateScreenPanelName returns if the screen panel name is valid.
// userMessage contains the error message for the user.
// Param name is the user created name.
func validateScreenPanelName(name string) (isValid bool, userMessage string) {
	if isValid, userMessage = validateLowerCamelCaseName(name, "screen panel"); !isValid {
		return
	}
	trimmed := strings.TrimSuffix(name, "Panel")
	if isValid = trimmed != name; !isValid {
		userMessage = `A panel name must end with the suffix "Panel".`
	}
	return
}

// validateScreenName returns if the screen name is valid.
// userMessage contains the error message for the user.
// Param name is the user created name.
func validateScreenName(name string) (isValid bool, userMessage string) {
	isValid, userMessage = validateLowerCaseName(name, "screen")
	return
}

// validateLowerCaseName returns if the name is valid.
// userMessage contains the error message for the user.
// Param name is the user created name.
// Param nameType is what the name is created for. ("record", "panel")
func validateLowerCaseName(name, nameType string) (isValid bool, userMessage string) {
	if strings.ToLower(name) != name {
		userMessage = fmt.Sprintf("The %s name %q must be lower case.", nameType, name)
		return
	}
	isValid, userMessage = validateName(name, nameType)
	return
}

// validateCamelCaseName returns if the name is valid.
// userMessage contains the error message for the user.
// Param name is the user created name.
// Param nameType is what the name is created for. ("record", "panel")
func validateCamelCaseName(name, nameType string) (isValid bool, userMessage string) {
	s := name[:1]
	if !strings.Contains(ucAlphabet, s) {
		userMessage = fmt.Sprintf("A %s name must begin with an upper case letter.", nameType)
		isValid = false
		return
	}
	isValid, userMessage = validateName(name, nameType)
	return
}

// validateLowerCamelCaseName returns if the name is valid.
// userMessage contains the error message for the user.
// Param name is the user created name.
// Param nameType is what the name is created for. ("record", "panel")
func validateLowerCamelCaseName(name, nameType string) (isValid bool, userMessage string) {
	s := name[:1]
	if !strings.Contains(lcAlphabet, s) {
		userMessage = fmt.Sprintf("A %s name must begin with a lower case letter.", nameType)
		isValid = false
		return
	}
	isValid, userMessage = validateName(name, nameType)
	return
}

// validateName returns if the name is valid.
// userMessage contains the error message for the user.
// Param name is the user created name.
// Param nameType is what the name is created for. ("record", "panel")
func validateName(name, nameType string) (isValid bool, userMessage string) {
	lines := make([]string, 0, 2)
	var msg string
	isValid = true
	if strings.ContainsAny(name, asciiSpaces) {
		msg = fmt.Sprintf("A %s name must not contain any white space.", nameType)
		lines = append(lines, msg)
		isValid = false
	}
	if strings.ContainsAny(name, puncts) {
		msg = fmt.Sprintf("A %s name must not contain any punctuation.", nameType)
		lines = append(lines, msg)
		isValid = false
	}
	lcName := strings.ToLower(name)
	for _, inValidName := range inValidNames {
		if inValidName == lcName {
			msg = fmt.Sprintf("The %s name %q can not be used.", nameType, name)
			lines = append(lines, msg)
			isValid = false
		}
	}
	if !isValid {
		userMessage = strings.Join(lines, "\n")
	}
	return
}

// Prefix returns a prefixed version of a string.
func Prefix(s string, prefix string) (prefixed string) {
	prefixed = prefix + s
	return
}

// Suffix returns a suffixed version of a string.
func Suffix(s string, suffix string) (suffixed string) {
	suffixed = s + suffix
	return
}

// Prefix returns a prefixed and suffixed version of a string.
func PrefixSuffix(s string, prefix, suffix string) (fixed string) {
	fixed = prefix + s + suffix
	return
}

// Suffix combinations.

func SuffixLowerCase(s string, suffix string) (lowerCased string) {
	suffixed := Suffix(s, suffix)
	lowerCased = strings.ToLower(suffixed)
	return
}

// CapSuffix returns a suffixed and then capped version of a string.
func CapSuffix(s string, suffix string) (capped string) {
	suffixed := Suffix(s, suffix)
	capped = Cap(suffixed)
	return
}

// DeCapSuffix returns a suffixed and then decapped version of a string.
func DeCapSuffix(s string, suffix string) (decapped string) {
	suffixed := Suffix(s, suffix)
	decapped = DeCap(suffixed)
	return
}

// Prefix combinations.

// PrefixLowerCase returns a prefixed and lowercased version of a string.
func PrefixLowerCase(s string, prefix string) (lowerCased string) {
	prefixed := Prefix(s, prefix)
	lowerCased = strings.ToLower(prefixed)
	return
}

// PrefixCap returns a prefixed and capped version of a string.
func PrefixCap(s string, prefix string) (capped string) {
	prefixed := Prefix(s, prefix)
	capped = Cap(prefixed)
	return
}

// PrefixDeCap returns a prefixed and decapped version of a string.
func PrefixDeCap(s string, prefix string) (decapped string) {
	prefixed := Prefix(s, prefix)
	decapped = DeCap(prefixed)
	return
}

// Prefix suffix combinations.

// PrefixCapSuffix returns a prefixed, capped and suffixed version of a string.
func PrefixCapSuffix(s string, prefix, suffix string) (capped string) {
	fixed := PrefixSuffix(s, prefix, suffix)
	capped = Cap(fixed)
	return
}

// PrefixCapSuffix returns a prefixed, decapped and suffixed version of a string.
func PrefixDeCapSuffix(s string, prefix, suffix string) (decapped string) {
	fixed := PrefixSuffix(s, prefix, suffix)
	decapped = DeCap(fixed)
	return
}
