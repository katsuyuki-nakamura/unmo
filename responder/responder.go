package responder

type Responder interface {
	Response(input string) string
	Name() string
}
