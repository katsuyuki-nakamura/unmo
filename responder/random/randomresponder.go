package randomresponder

import (
	"math/rand"

	"github.com/katsuyuki-nakamura/unmo/dictionary"
	"github.com/katsuyuki-nakamura/unmo/responder"
)

// responder
type RandomResponder struct {
	name      string
	responses []string
}

func selectRandom(responses []string) string {
	return responses[rand.Intn(len(responses))]
}
func NewRandomResponder() responder.Responder {

	r := new(RandomResponder)
	r.name = "Random"
	r.responses = dictionary.GetInstance().Random()
	return r
}

func (r RandomResponder) Response(input string) string {
	return selectRandom(r.responses)
}

func (r RandomResponder) Name() string {
	return r.name
}
