package auth

type PhoneResponseToken struct {
	JWT string `json:"token"`
}

type PhoneRequestWithCode struct {
	SessionId string `json:"session_id" validate:"required"`
	Code      int    `json:"code" validate:"required"`
}

type PhoneRequest struct {
	Phone string `json:"phone" validate:"required"`
}

type PhoneResponse struct {
	SessionId string `json:"sessionId"`
}
