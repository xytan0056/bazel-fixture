// Package strutil provides string helpers, including a replacements table
// loaded from an embedded JSON file.
package strutil

import (
	_ "embed"
	"encoding/json"
	"strings"
)

//go:embed replacements.json
var replacementsData []byte

var replacements map[string]string

func init() {
	if err := json.Unmarshal(replacementsData, &replacements); err != nil {
		panic("strutil: failed to parse embedded replacements.json: " + err.Error())
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
