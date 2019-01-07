package unmo

import (
	"math/rand"

	"github.com/katsuyuki-nakamura/unmo/responder"
	patternresponder "github.com/katsuyuki-nakamura/unmo/responder/pattern"
	randomresponder "github.com/katsuyuki-nakamura/unmo/responder/random"
	whatresponder "github.com/katsuyuki-nakamura/unmo/responder/what"
)

type Unmo struct {
	name              string
	responder_what    responder.Responder
	responder_random  responder.Responder
	responder_pattern responder.Responder
	responder         responder.Responder
}

func NewUnmo(name string) *Unmo {
	u := new(Unmo)
	u.name = name
	u.responder_what = whatresponder.NewWhatResponder()
	u.responder_random = randomresponder.NewRandomResponder()
	u.responder_pattern = patternresponder.NewPatternResponder()
	u.responder = u.responder_pattern
	return u
}

// 構造体のメンバーを変更する場合は、ポインタ渡し
func (u *Unmo) Dialogue(input string) string {
	num := rand.Intn(100)
	switch {
	case (num >= 0 && num < 60):
		u.responder = u.responder_what
	case (num > 60 && num < 90):
		u.responder = u.responder_random
	default:
		u.responder = u.responder_pattern
	}
	return u.responder.Response(input)
}

func (u Unmo) ResponderName() string {
	return u.responder.Name()
}

func (u Unmo) Name() string {
	return u.name
}
