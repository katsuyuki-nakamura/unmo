package whatresponder

// responder
type WhatResponder struct {
	name string
}

func NewWhatResponder() *WhatResponder {
	r := &(WhatResponder{name: "What"})
	return r
}

func (r WhatResponder) Response(input string) string {
	return input + "ってなに？"
}

func (r WhatResponder) Name() string {
	return r.name
}
