package human

type Human struct {
	token string
}

func New(token string) *Human {
	return &Human{token: token}
}

func (h *Human) GetToken() string {
	return h.token
}
