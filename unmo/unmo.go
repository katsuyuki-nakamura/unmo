package unmo

import (
	"math/rand"

	"github.com/katsuyuki-nakamura/unmo/responder"
	patternresponder "github.com/katsuyuki-nakamura/unmo/responder/pattern"
	randomresponder "github.com/katsuyuki-nakamura/unmo/responder/random"
	whatresponder "github.com/katsuyuki-nakamura/unmo/responder/what"
)

type Unmo struct {
	name       string
	responders []responder.Responder
	responder  responder.Responder
}

func NewUnmo(name string) *Unmo {
	u := new(Unmo)
	u.name = name
	u.responders = append(u.responders, whatresponder.NewWhatResponder())
	u.responders = append(u.responders, randomresponder.NewRandomResponder())
	u.responders = append(u.responders, patternresponder.NewPatternResponder())
	u.responder = u.responders[0]
	return u
}

// 構造体のメンバーを変更する場合は、ポインタ渡し
func (u *Unmo) Dialogue(input string) string {
	u.responder = u.responders[rand.Intn(len(u.responders))]
	return u.responder.Response(input)
}

func (u Unmo) ResponderName() string {
	return u.responder.Name()
}

func (u Unmo) Name() string {
	return u.name
}
