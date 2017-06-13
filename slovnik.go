package slovnik

// Word defines a structure with the word itself and possible translations of that word
type Word struct {
	Word         string      `json:"word"`
	Translations []string    `json:"translations"`
	WordType     string      `json:"wordType"`
	Synonyms     []string    `json:"synonyms"`
	Antonyms     []string    `json:"antonyms"`
	DerivedWords []string    `json:"derivedWords"`
	Samples      []SampleUse `json:"samples"`
}

// SampleUse describes example phrase in which word can be used
type SampleUse struct {
	Keyword     string
	Phrase      string
	Translation string
}
