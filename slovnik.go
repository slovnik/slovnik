package slovnik

import (
	"fmt"
	"strings"
)

// Word defines a structure with the word itself and possible translations of that word
type Word struct {
	Word         string   `json:"word"`
	Translations []string `json:"translations"`
	WordType     string   `json:"wordType"`
	Synonyms     []string `json:"synonyms"`
	Antonyms     []string `json:"antonyms"`
	DerivedWords []string `json:"derivedWords"`
}

// Method for transforming Word struct to string
func (w Word) String() string {
	out := fmt.Sprintf("*%s* - %s\n\n", w.Word, strings.Join(w.Translations, ", "))
	out += fmt.Sprintf("*%s*\n", w.WordType)

	if len(w.Synonyms) > 0 {
		out += fmt.Sprintln("\n*Synonyms:*")
		out += fmt.Sprintln(strings.Join(w.Synonyms, ", "))
	}
	if len(w.Antonyms) > 0 {
		out += fmt.Sprintln("\n*Antonyms:*")
		out += fmt.Sprintln(strings.Join(w.Antonyms, ", "))
	}

	if len(w.DerivedWords) > 0 {
		out += fmt.Sprintln("\n*Derived words:*")
		out += fmt.Sprintln(strings.Join(w.DerivedWords, ", "))
	}
	return out
}
