package responder

// responder
type Responder struct {
	name string
}

func NewResponder(name string) *Responder {
	r := new(Responder)
	r.name = name
	return r
}

func (r Responder) Response(input string) string {
	return input + "ってなに？"
}

func (r Responder) Name() string {
	return r.name
}
