package randomresponder

import "math/rand"

// responder
type RandomResponder struct {
	name      string
	responses []string
}

func NewRandomResponder() *RandomResponder {
	r := &(RandomResponder{name: "Random", responses: []string{"今日はさむいね", "チョコたべたい", "きのう10円ひろった"}})
	return r
}

func (r RandomResponder) Response(input string) string {
	return r.responses[rand.Intn(len(r.responses))]
}

func (r RandomResponder) Name() string {
	return r.name
}
