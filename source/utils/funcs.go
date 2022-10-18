package utils

import (
	"strings"
)

const (
	ucAlphabet  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits      = "1234567890"
	asciiSpaces = " \t\r\n\v\f"
	puncts      = "`~!@#$%^&*()_-+={[]}|\\:;\"'<>,.?/"
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

	Prefix            func(s string, prefix string) (prefixed string)
	Suffix            func(s string, suffix string) (suffixed string)
	PrefixSuffix      func(s string, prefix, suffix string) (fixed string)
	SuffixLowerCase   func(s string, suffix string) (lowerCased string)
	CapSuffix         func(s string, suffix string) (capped string)
	DeCapSuffix       func(s string, suffix string) (decapped string)
	PrefixLowerCase   func(s string, prefix string) (lowerCased string)
	PrefixCap         func(s string, prefix string) (capped string)
	PrefixDeCap       func(s string, prefix string) (decapped string)
	PrefixCapSuffix   func(s string, prefix, suffix string) (capped string)
	PrefixDeCapSuffix func(s string, prefix, suffix string) (decapped string)

	PadSlice             func(slice []string) (padded []string)
	SuffixSlice          func(slice []string, suffix string) (suffixed []string)
	SuffixPadSlice       func(slice []string, suffix string) (padded []string)
	SuffixLowerCaseSlice func(slice []string, suffix string) (lowerCased []string)
	DeCapSuffixSlice     func(slice []string, suffix string) (suffixed []string)
	CapSuffixSlice       func(slice []string, suffix string) (suffixed []string)
	DeCapPadSlice        func(slice []string) (padded []string)
	CapPadSlice          func(slice []string) (padded []string)
	DeCapSuffixPadSlice  func(slice []string, suffix string) (suffixed []string)
	CapSuffixPadSlice    func(slice []string, suffix string) (suffixed []string)

	PrefixPadSlice            func(slice []string, prefix string) (padded []string)
	PrefixLowerCaseSlice      func(slice []string, prefix string) (lowerCased []string)
	PrefixSuffixSlice         func(slice []string, prefix, suffix string) (suffixed []string)
	PrefixSuffixPadSlice      func(slice []string, prefix, suffix string) (padded []string)
	PrefixDeCapSuffixSlice    func(slice []string, prefix, suffix string) (suffixed []string)
	PrefixCapSuffixSlice      func(slice []string, prefix, suffix string) (suffixed []string)
	PrefixDeCapPadSlice       func(slice []string, prefix string) (padded []string)
	PrefixCapPadSlice         func(slice []string, prefix string) (padded []string)
	PrefixDeCapSuffixPadSlice func(slice []string, prefix, suffix string) (suffixed []string)
	PrefixCapSuffixPadSlice   func(slice []string, prefix, suffix string) (suffixed []string)
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

		Prefix:            Prefix,
		Suffix:            Suffix,
		PrefixSuffix:      PrefixSuffix,
		SuffixLowerCase:   SuffixLowerCase,
		CapSuffix:         CapSuffix,
		DeCapSuffix:       DeCapSuffix,
		PrefixLowerCase:   PrefixLowerCase,
		PrefixCap:         PrefixCap,
		PrefixDeCap:       PrefixDeCap,
		PrefixCapSuffix:   PrefixCapSuffix,
		PrefixDeCapSuffix: PrefixDeCapSuffix,

		PadSlice:             PadSlice,
		SuffixSlice:          SuffixSlice,
		SuffixPadSlice:       SuffixPadSlice,
		SuffixLowerCaseSlice: SuffixLowerCaseSlice,
		CapSuffixSlice:       CapSuffixSlice,
		DeCapSuffixSlice:     DeCapSuffixSlice,
		CapSuffixPadSlice:    CapSuffixPadSlice,
		DeCapPadSlice:        DeCapPadSlice,
		CapPadSlice:          CapPadSlice,
		DeCapSuffixPadSlice:  DeCapSuffixPadSlice,

		PrefixPadSlice:            PrefixPadSlice,
		PrefixLowerCaseSlice:      PrefixLowerCaseSlice,
		PrefixSuffixSlice:         PrefixSuffixSlice,
		PrefixSuffixPadSlice:      PrefixSuffixPadSlice,
		PrefixCapSuffixSlice:      PrefixCapSuffixSlice,
		PrefixDeCapSuffixSlice:    PrefixDeCapSuffixSlice,
		PrefixCapSuffixPadSlice:   PrefixCapSuffixPadSlice,
		PrefixDeCapPadSlice:       PrefixDeCapPadSlice,
		PrefixCapPadSlice:         PrefixCapPadSlice,
		PrefixDeCapSuffixPadSlice: PrefixDeCapSuffixPadSlice,
	}
	return
}

func IncInt(i int) (newI int) {
	newI = i + 1
	return
}
