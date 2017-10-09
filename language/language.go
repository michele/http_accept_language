package language

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Languages struct {
	languages []string
	locales   []string
	header    string
}

type locale struct {
	locale  string
	lang    string
	quality float64
}

var Default string
var whiteSpace = regexp.MustCompile(`\s+`)
var validate = regexp.MustCompile(`(?i)^[a-z\-0-9]+|\*$`)
var matchTerritory = regexp.MustCompile(`(?i)-[a-z0-9]+$`)

func ParseHeader(header string) (lang Languages) {
	lang.header = whiteSpace.ReplaceAllString(header, "")
	lang.languages = []string{}
	lang.locales = []string{}
	temps := strings.Split(lang.header, ",")
	locales := []locale{}
	for _, l := range temps {
		ss := strings.Split(l, ";q=")
		if validate.MatchString(ss[0]) {
			if ss[0] == "*" {
				continue
			}
			loc := locale{
				locale: matchTerritory.ReplaceAllStringFunc(strings.ToLower(ss[0]), func(terr string) string {
					return strings.ToUpper(terr)
				}),
				lang: strings.ToLower(strings.Split(ss[0], "-")[0]),
			}
			if len(ss) > 1 {
				q, err := strconv.ParseFloat(ss[1], 32)
				if err == nil {
					loc.quality = q
				}
			} else {
				loc.quality = 1.0
			}

			locales = append(locales, loc)
		}
	}
	sort.Slice(locales, func(i, j int) bool {
		return locales[i].quality > locales[j].quality
	})
	presenceLang := map[string]interface{}{}
	presenceLoc := map[string]interface{}{}
	for _, loc := range locales {
		if _, ok := presenceLoc[loc.locale]; ok {
			continue
		}
		lang.locales = append(lang.locales, loc.locale)
		presenceLoc[loc.locale] = nil
		if _, ok := presenceLang[loc.lang]; ok {
			continue
		}
		lang.languages = append(lang.languages, loc.lang)
		presenceLang[loc.lang] = nil
	}
	return
}

func (l Languages) Preferred() string {
	if len(l.languages) == 0 {
		return Default
	}
	return l.languages[0]
}

func (l Languages) PreferredLocale() string {
	if len(l.locales) == 0 {
		return Default
	}
	return l.locales[0]
}

func (l Languages) All() []string {
	if len(l.languages) == 0 {
		return []string{Default}
	}
	return append([]string(nil), l.languages...)
}

func (l Languages) AllLocales() []string {
	if len(l.locales) == 0 {
		return []string{Default}
	}
	return append([]string(nil), l.locales...)
}
