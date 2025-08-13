package handlers

import "time"

type Handlers struct {
	TTL         time.Duration
	ServerToken string
	AdminToken  string
}

func NewHandlers(ServerToken string, Token string, duration time.Duration) *Handlers {
	return &Handlers{AdminToken: Token,
		ServerToken: ServerToken,
		TTL:         duration}
}
