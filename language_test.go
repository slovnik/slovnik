package slovnik

import "testing"

func TestDetectLanguage(t *testing.T) {
	cases := []struct {
		in   string
		lang Language
	}{
		{"hlavní", Cz},
		{"привет", Ru},
		{"sиniy", Ru},
	}

	for _, c := range cases {
		got := DetectLanguage(c.in)
		if got != c.lang {
			t.Errorf("DetectLanguage(%q) == %q, want %q", c.in, got, c.lang)
		}
	}
}
