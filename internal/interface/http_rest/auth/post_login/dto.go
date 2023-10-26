package post_login

import (
	"link-shortener/internal/application/auth/login"
)

type requestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *requestBody) toUsecasePayload() *login.Payload {
	return &login.Payload{
		Username: r.Username,
		Password: r.Password,
	}
}

type token struct {
	Token string `json:"token"`
}

type responseBody struct {
	Access  *token `json:"access_token"`
	Refresh *token `json:"refreh_token"`
}

func responseFromResult(r *login.Result) *responseBody {
	return &responseBody{
		Access: &token{
			Token: r.Access,
		},
		Refresh: &token{
			Token: r.Refresh,
		},
	}
}
