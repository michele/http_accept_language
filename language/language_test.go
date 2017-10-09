package language_test

import (
	"testing"

	"github.com/michele/http_accept_language/language"
)

var lang language.Languages

var header = " en-us;q=0.3, en-GB;q=2.0, fr , IT, * , #not-valid, en-GB "
var goldenLangs = []string{"en", "fr", "it"}
var goldenLocs = []string{"en-GB", "fr", "it", "en-US"}

func TestPrepare(t *testing.T) {
	language.Default = "xx"

	lang = language.ParseHeader(header)
}

func TestCorrectLanguagesAreParsed(t *testing.T) {
	parsed := lang.All()
	if len(parsed) != len(goldenLangs) {
		t.Error("Parsed languages and golden haven't got the same length")
		t.Errorf("Expected %+v\nGot %+v", goldenLangs, parsed)
		t.FailNow()
	}

	for i, l := range goldenLangs {
		if l != parsed[i] {
			t.Error("Parsed languages and golden don't match")
		}
	}
}

func TestCorrectLocalesAreParsed(t *testing.T) {
	parsed := lang.AllLocales()
	if len(parsed) != len(goldenLocs) {
		t.Error("Parsed locales and golden haven't got the same length")
		t.Errorf("Expected %+v\nGot %+v", goldenLocs, parsed)
		t.FailNow()
	}

	for i, l := range goldenLocs {
		if l != parsed[i] {
			t.Error("Parsed locales and golden don't match")
		}
	}
}

func TestPreferredLanguage(t *testing.T) {
	if lang.Preferred() != goldenLangs[0] {
		t.Error("Preferred language is not correct")
	}
}

func TestPreferredLocale(t *testing.T) {
	if lang.PreferredLocale() != goldenLocs[0] {
		t.Error("Preferred locale is not correct")
	}
}

func TestDefaultIfNoLanguageIsPassed(t *testing.T) {
	l := language.ParseHeader("*, #not-valid")
	if l.Preferred() != language.Default {
		t.Error("Default should've been passed")
	}
	if l.PreferredLocale() != language.Default {
		t.Error("Default should've been passed")
	}
	if l.All()[0] != language.Default {
		t.Error("Languages should've included Default")
	}
	if l.AllLocales()[0] != language.Default {
		t.Error("Locales should've included Default")
	}
}
