package unmo

import (
	"github.com/katsuyuki-nakamura/unmo/responder"
)

type Unmo struct {
	name      string
	responder *responder.Responder
}

func NewUnmo(name string) *Unmo {
	u := new(Unmo)
	u.name = name
	u.responder = responder.NewResponder("What")
	return u
}

func (u Unmo) Dialogue(input string) string {
	return u.responder.Response(input)
}

func (u Unmo) ResponderName() string {
	return u.responder.Name()
}

func (u Unmo) Name() string {
	return u.name
}
