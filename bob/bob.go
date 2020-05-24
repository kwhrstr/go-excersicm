// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	realRemark := strings.TrimSpace(remark)
	if realRemark == "" {
		return "Fine. Be that way!"
	}

	if isShout(realRemark) && strings.HasSuffix(realRemark, "?"){
		return "Calm down, I know what I'm doing!"
	}
	if isShout(realRemark) {
		return "Whoa, chill out!"
	}
	if strings.HasSuffix(realRemark, "?") {
		return "Sure."
	}

	return "Whatever."
}

func isShout(remark string) bool {
	return HasLetter(remark) && strings.ToUpper(remark) == remark
}

func HasLetter(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}