package dto

type SignUpDto struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type SignInDto struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
