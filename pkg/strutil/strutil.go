// Package strutil provides string helpers, including a replacements table
// and greetings table loaded from embedded JSON files.
package strutil

import (
	_ "embed"
	"encoding/json"
	"strings"
)

//go:embed replacements.json
var replacementsData []byte

//go:embed greetings.json
var greetingsData []byte

var (
	replacements map[string]string
	greetings    map[string]string
)

func init() {
	if err := json.Unmarshal(replacementsData, &replacements); err != nil {
		panic("strutil: failed to parse embedded replacements.json: " + err.Error())
	}
	if err := json.Unmarshal(greetingsData, &greetings); err != nil {
		panic("strutil: failed to parse embedded greetings.json: " + err.Error())
	}
}

// Sanitize replaces symbols in s with their word equivalents from the
// embedded replacements table.
func Sanitize(s string) string {
	for symbol, word := range replacements {
		s = strings.ReplaceAll(s, symbol, word)
	}
	return s
}

// TitleWords upper-cases the first letter of every whitespace-delimited word.
func TitleWords(s string) string {
	words := strings.Fields(s)
	for i, w := range words {
		if len(w) > 0 {
			words[i] = strings.ToUpper(w[:1]) + w[1:]
		}
	}
	return strings.Join(words, " ")
}

// Greet returns the greeting for the given locale, or the English greeting
// when locale is unknown.
func Greet(locale string) string {
	if g, ok := greetings[locale]; ok {
		return g
	}
	return greetings["en"]
}

// Reverse returns s with its runes reversed.
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
