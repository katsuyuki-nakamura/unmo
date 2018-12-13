package patternresponder

import (
	"math/rand"
	"regexp"
	"strings"

	"github.com/katsuyuki-nakamura/unmo/dictionary"

	"github.com/katsuyuki-nakamura/unmo/responder"
)

// responder
type PatternResponder struct {
	name    string
	pattern map[string]string
}

func NewPatternResponder() responder.Responder {
	r := new(PatternResponder)
	r.name = "Pattern"
	r.pattern = dictionary.GetInstance().Pattern()
	return r
}

func selectRandom(responses []string) string {
	return responses[rand.Intn(len(responses))]
}

func (r PatternResponder) Response(input string) string {
	for pattern, phrase := range r.pattern {
		reg := regexp.MustCompile(pattern)
		if reg.MatchString(input) {
			response := selectRandom(strings.Split(phrase, "|"))
			return response
		}
	}
	return ""
}

func (r PatternResponder) Name() string {
	return r.name
}
