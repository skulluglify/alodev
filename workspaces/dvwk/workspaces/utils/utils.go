package utils

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

func CapitalizeWord(word string) string {
	if len(word) == 0 {
		return ""
	}
	word = strings.Trim(word, " ")
	return strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
}

func CapitalizeEachWord(word string) string {
	if len(word) == 0 {
		return ""
	}
	word = strings.Trim(word, " ")
	parts := strings.Split(word, " ")
	for i := 0; i < len(parts); i++ {
		parts[i] = CapitalizeWord(parts[i])
	}
	return strings.Join(parts, " ")
}

func Title(word string) string {
	title := cases.Title(language.Und)
	return title.String(word)
}
