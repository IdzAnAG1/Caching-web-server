package handlers

type Handlers struct {
	Token string
}

func NewHandlers(Token string) *Handlers {
	return &Handlers{Token: Token}
}
