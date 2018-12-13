package unmo

import (
	"math/rand"

	"github.com/katsuyuki-nakamura/unmo/responder"
	randomresponder "github.com/katsuyuki-nakamura/unmo/responder/random"
	whatresponder "github.com/katsuyuki-nakamura/unmo/responder/what"
)

type Unmo struct {
	name       string
	responders []responder.Responder
	responder  responder.Responder
}

func NewUnmo(name string) *Unmo {
	u := &(Unmo{
		name:       name,
		responders: []responder.Responder{whatresponder.NewWhatResponder(), randomresponder.NewRandomResponder()},
	})
	u.responder = u.responders[0]
	return u
}

func (u Unmo) Dialogue(input string) string {
	u.responder = u.responders[rand.Intn(2)]
	return u.responder.Response(input)
}

func (u Unmo) ResponderName() string {
	return u.responder.Name()
}

func (u Unmo) Name() string {
	return u.name
}
