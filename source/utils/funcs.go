package utils

import (
	"fmt"
	"strings"
)

const (
	ucAlphabet  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits      = "1234567890"
	asciiSpaces = " \t\r\n\v\f"
)

// inValidNames must be all lowercase.
var inValidNames = []string{
	"package",
}

var lcAlphabet = strings.ToLower(ucAlphabet)

type Funcs struct {
	LowerCase      func(string) string
	UpperCase      func(string) string
	LabelToName    func(string) string
	LabelToVarName func(string) string
	DeCap          func(string) string
	Inc            func(i int) (newI int)
	Comment        func(string) string
	Cap            func(string) string
}

func GetFuncs() (funcs Funcs) {
	funcs = Funcs{
		LowerCase:      strings.ToLower,
		UpperCase:      strings.ToUpper,
		LabelToName:    LabelToName,
		LabelToVarName: LabelToVarName,
		DeCap:          DeCap,
		Inc:            IncInt,
		Comment:        comment,
		Cap:            Cap,
	}
	return
}

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

func Cap(name string) (capped string) {
	capped = strings.ToUpper(name[:1]) + name[1:]
	return
}

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

// validateName returns is the name is valid.
// userMessage contains the error message for the user.
func validateName(name string) (isValid bool, userMessage string) {
	lines := make([]string, 0, 2)
	var msg string
	s := name[:1]
	isValid = true
	if !strings.Contains(ucAlphabet, s) {
		msg = "A name must begin with a capital letter."
		lines = append(lines, msg)
		isValid = false
	}
	if strings.ContainsAny(name, asciiSpaces) {
		msg = "A name must not contain any white space."
		lines = append(lines, msg)
		isValid = false
	}
	lcName := strings.ToLower(name)
	for _, inValidName := range inValidNames {
		if inValidName == lcName {
			msg = fmt.Sprintf("The name %q can not be used.", name)
			lines = append(lines, msg)
			isValid = false
		}
	}
	if !isValid {
		userMessage = strings.Join(lines, "\n")
	}
	return
}

func IncInt(i int) (newI int) {
	newI = i + 1
	return
}
