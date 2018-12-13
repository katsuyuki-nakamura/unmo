package whatresponder

import (
	"github.com/katsuyuki-nakamura/unmo/responder"
)

// responder
type WhatResponder struct {
	name string
}

func NewWhatResponder() responder.Responder {
	r := &(WhatResponder{name: "What"})
	return r
}

func (r WhatResponder) Response(input string) string {
	return input + "ってなに？"
}

func (r WhatResponder) Name() string {
	return r.name
}
