package slovnik

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewClient(t *testing.T) {
	const expectedBaseURL = "http://example.com/"
	c, err := NewClient(expectedBaseURL)
	if err != nil {
		t.Fatalf("NewClient, got error '%q', want no error", err)
	}

	if c.baseURL.String() != expectedBaseURL {
		t.Fatalf("NewClient, baseURL = %q, want = %q", c.baseURL, expectedBaseURL)
	}

	if c.client == nil {
		t.Fatalf("NewClient, client = %q, want to be defined", c.client)
	}
}

func TestNewClientBadURL(t *testing.T) {
	const url = "ht\tp://localhost"
	c, err := NewClient(url)

	if c != nil {
		t.Fatalf("NewClient, client = %q, want nil", c)
	}

	if err == nil {
		t.Fatal("NewClient, want an error")
	}
}

func TestTranslate(t *testing.T) {
	const expectedWord = "hlavni"
	expectedTranslations := []string{"Главный"}
	expectedSynonyms := []string{"ústřední"}
	expectedAntonyms := []string{"vedlejší"}
	expectedDerivedWords := []string{"hlavně"}
	expectedWordType := "přídavné jméno"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if r.Method != http.MethodGet {
			t.Fatalf("Got %q, want %q", r.Method, http.MethodGet)
		}

		expectedPath := "/api/translate/hlavni"

		if r.URL.Path != expectedPath {
			t.Fatalf("Translate got %q, want %q", r.URL.Path, expectedPath)
		}

		result := Word{
			Word:         expectedWord,
			Translations: expectedTranslations,
			Synonyms:     expectedSynonyms,
			Antonyms:     expectedAntonyms,
			DerivedWords: expectedDerivedWords,
			WordType:     expectedWordType,
		}

		json.NewEncoder(w).Encode(result)
	}))
	defer ts.Close()

	slovnikURL := ts.URL
	c, _ := NewClient(slovnikURL)

	w, err := c.Translate(expectedWord)
	if err != nil {
		t.Fatalf("Translate, got error '%q', want no error", err)
	}

	if w.Word != expectedWord {
		t.Fatalf("Translate, word = %q, want = %q", w.Word, expectedWord)
	}

	if w.WordType != expectedWordType {
		t.Fatalf("Translate, wordType = %q, want = %q", w.WordType, expectedWordType)
	}

	if len(w.Translations) != len(expectedTranslations) {
		t.Errorf("Translate len(translation) == %d, want %d", len(w.Translations), len(expectedTranslations))
	}

	for i, trans := range w.Translations {
		if trans != expectedTranslations[i] {
			t.Errorf("Translate translation == %q, want %q", trans, expectedTranslations[i])
		}
	}

	if w.WordType != expectedWordType {
		t.Errorf("Translate wordType == %q, want %q", w.WordType, expectedWordType)
	}

	if len(w.Synonyms) != len(expectedSynonyms) {
		t.Errorf("Translate len(synonyms) == %d, want %d", len(w.Synonyms), len(expectedSynonyms))
	}

	for i, synonym := range w.Synonyms {
		if synonym != expectedSynonyms[i] {
			t.Errorf("Translate synonym == %q, want %q", synonym, expectedSynonyms[i])
		}
	}

	if len(w.Antonyms) != len(expectedAntonyms) {
		t.Errorf("Translate len(antonyms) == %d, want %d", len(w.Antonyms), len(expectedAntonyms))
	}

	for i, antonym := range w.Antonyms {
		if antonym != expectedAntonyms[i] {
			t.Errorf("Translate antonym == %q, want %q", antonym, expectedAntonyms[i])
		}
	}

	if len(w.DerivedWords) != len(expectedDerivedWords) {
		t.Errorf("Translate len(derivedWords) == %d, want %d", len(w.DerivedWords), len(expectedDerivedWords))
	}

	for i, derived := range w.DerivedWords {
		if derived != expectedDerivedWords[i] {
			t.Errorf("Translate derivedWord == %q, want %q", derived, expectedDerivedWords[i])
		}
	}
}

func TestTranslateError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer ts.Close()

	slovnikURL := ts.URL
	c, _ := NewClient(slovnikURL)
	w, err := c.Translate("hlavni")

	if w != nil {
		t.Fatal("Translate, word expected to be nil")
	}

	if err == nil {
		t.Fatal("Translate expected to return error")
	}

}
