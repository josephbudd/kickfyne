package utils

import "strings"

// Base.

// CapSlice returns a new slice in which each item's first character is upper cased.
func CapSlice(slice []string) (capped []string) {
	capped = make([]string, len(slice))
	for i, s := range slice {
		capped[i] = Cap(s)
	}
	return
}

// DeCapSlice returns a new slice in which each item's first character is lower cased.
func DeCapSlice(slice []string) (decapped []string) {
	decapped = make([]string, len(slice))
	for i, s := range slice {
		decapped[i] = DeCap(s)
	}
	return
}

// LowerCaseSlice returns a new slice in which each item is lower cased.
func LowerCaseSlice(slice []string) (lowerCased []string) {
	lowerCased = make([]string, len(slice))
	for i, s := range slice {
		lowerCased[i] = strings.ToLower(s)
	}
	return
}

// PadSlice returns a new slice in which each item is of equal length and right padded with spaces.
func PadSlice(slice []string) (padded []string) {
	var maxL int
	for _, s := range slice {
		l := len(s)
		if l > maxL {
			maxL = l
		}
	}
	var b strings.Builder
	for i := 0; i < maxL; i++ {
		b.WriteRune(' ')
	}
	padding := b.String()
	padded = make([]string, len(slice))
	for i, s := range slice {
		b.Reset()
		b.WriteString(s)
		l := len(s)
		if l < maxL {
			// Add some padding.
			b.WriteString(padding[:(maxL - l)])
		}
		padded[i] = b.String()
	}
	return
}

// PrefixSlice returns a new slice in which each item is prefixed.
func PrefixSlice(slice []string, prefix string) (prefixed []string) {
	prefixed = make([]string, len(slice))
	for i, s := range slice {
		prefixed[i] = prefix + s
	}
	return
}

// SuffixSlice returns a new slice in which each item is suffixed.
func SuffixSlice(slice []string, suffix string) (suffixed []string) {
	suffixed = make([]string, len(slice))
	for i, s := range slice {
		suffixed[i] = s + suffix
	}
	return
}

// PrefixSuffixSlice returns a new slice in which each item is prefixed and suffixed.
func PrefixSuffixSlice(slice []string, prefix, suffix string) (fixed []string) {
	fixed = make([]string, len(slice))
	for i, s := range slice {
		fixed[i] = prefix + s + suffix
	}
	return
}

// DeCapPadSlice returns a new slice in which each item is decapped and padded.
func DeCapPadSlice(slice []string) (padded []string) {
	decapped := DeCapSlice(slice)
	padded = PadSlice(decapped)
	return
}

// CapPadSlice returns a new slice in which each item is capped and padded.
func CapPadSlice(slice []string) (padded []string) {
	capped := CapSlice(slice)
	padded = PadSlice(capped)
	return
}

// Suffix combinations.

// SuffixPadSlice returns a new slice in which each item is suffixed and padded.
func SuffixPadSlice(slice []string, suffix string) (padded []string) {
	suffixed := SuffixSlice(slice, suffix)
	padded = PadSlice(suffixed)
	return
}

// SuffixLowerCaseSlice returns a new slice in which each item is suffixed and lower cased.
func SuffixLowerCaseSlice(slice []string, suffix string) (lowerCased []string) {
	suffixed := SuffixSlice(slice, suffix)
	lowerCased = LowerCaseSlice(suffixed)
	return
}

// CapSuffixSlice returns a new slice in which each item is capped and suffixed.
func CapSuffixSlice(slice []string, suffix string) (capped []string) {
	suffixed := SuffixSlice(slice, suffix)
	capped = CapSlice(suffixed)
	return
}

// DeCapSuffixSlice returns a new slice in which each item is decapped and suffixed.
func DeCapSuffixSlice(slice []string, suffix string) (decapped []string) {
	suffixed := SuffixSlice(slice, suffix)
	decapped = DeCapSlice(suffixed)
	return
}

// CapSuffixPadSlice returns a new slice in which each item is capped, suffixed and padded.
func CapSuffixPadSlice(slice []string, suffix string) (padded []string) {
	capped := CapSlice(slice)
	suffixed := SuffixSlice(capped, suffix)
	padded = PadSlice(suffixed)
	return
}

// DeCapSuffixPadSlice returns a new slice in which each item is decapped, suffixed and padded.
func DeCapSuffixPadSlice(slice []string, suffix string) (padded []string) {
	decapped := DeCapSlice(slice)
	suffixed := SuffixSlice(decapped, suffix)
	padded = PadSlice(suffixed)
	return
}

// Prefix combinations.

// PrefixPadSlice returns a new slice in which each item is prefixed and padded.
func PrefixPadSlice(slice []string, prefix string) (padded []string) {
	prefixed := PrefixSlice(slice, prefix)
	padded = PadSlice(prefixed)
	return
}

// PrefixLowerCaseSlice returns a new slice in which each item is prefixed and lower cased.
func PrefixLowerCaseSlice(slice []string, prefix string) (lowerCased []string) {
	prefixed := PrefixSlice(slice, prefix)
	lowerCased = LowerCaseSlice(prefixed)
	return
}

// PrefixDeCapPadSlice returns a new slice in which each item is prefixed, decapped and padded.
func PrefixDeCapPadSlice(slice []string, prefix string) (padded []string) {
	prefixed := PrefixSlice(slice, prefix)
	decapped := DeCapSlice(prefixed)
	padded = PadSlice(decapped)
	return
}

// PrefixCapPadSlice returns a new slice in which each item is prefixed, capped and padded.
func PrefixCapPadSlice(slice []string, prefix string) (padded []string) {
	prefixed := PrefixSlice(slice, prefix)
	capped := CapSlice(prefixed)
	padded = PadSlice(capped)
	return
}

// PrefixCapSlice returns a new slice in which each item is prefixed and capped.
func PrefixCapSlice(slice []string, prefix string) (capped []string) {
	prefixed := PrefixSlice(slice, prefix)
	capped = CapSlice(prefixed)
	return
}

// PrefixDeCapSlice returns a new slice in which each item is prefixed and decapped.
func PrefixDeCapSlice(slice []string, prefix string) (decapped []string) {
	prefixed := PrefixSlice(slice, prefix)
	decapped = DeCapSlice(prefixed)
	return
}

// Prefix suffix combinations.

// PrefixSuffixPadSlice returns a new slice in which each item is prefixed, suffixed and padded.
func PrefixSuffixPadSlice(slice []string, prefix, suffix string) (padded []string) {
	fixed := PrefixSuffixSlice(slice, prefix, suffix)
	padded = PadSlice(fixed)
	return
}

// PrefixCapSuffixSlice returns a new slice in which each item is prefixed, suffixed and capped.
func PrefixCapSuffixSlice(slice []string, prefix, suffix string) (capped []string) {
	fixed := PrefixSuffixSlice(slice, prefix, suffix)
	capped = CapSlice(fixed)
	return
}

// PrefixDeCapSuffixSlice returns a new slice in which each item is prefixed, suffixed and decapped.
func PrefixDeCapSuffixSlice(slice []string, prefix, suffix string) (decapped []string) {
	fixed := PrefixSuffixSlice(slice, prefix, suffix)
	decapped = DeCapSlice(fixed)
	return
}

// PrefixCapSuffixPadSlice returns a new slice in which each item is prefixed, suffixed, capped and padded.
func PrefixCapSuffixPadSlice(slice []string, prefix, suffix string) (padded []string) {
	fixed := PrefixSuffixSlice(slice, prefix, suffix)
	capped := CapSlice(fixed)
	padded = PadSlice(capped)
	return
}

// PrefixDeCapSuffixPadSlice returns a new slice in which each item is prefixed, suffixed, decapped and padded.
func PrefixDeCapSuffixPadSlice(slice []string, prefix, suffix string) (padded []string) {
	fixed := PrefixSuffixSlice(slice, prefix, suffix)
	decapped := DeCapSlice(fixed)
	padded = PadSlice(decapped)
	return
}
