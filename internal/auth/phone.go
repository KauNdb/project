package auth

import (
	"math/rand"

	"gorm.io/gorm"
)

type Phone struct {
	gorm.Model
	Phone     string `json:"phone"`
	SessionId string `json:"session_id"`
	Code      int    `json:"code"`
}

func NewPhone(phone string) *Phone {
	return &Phone{
		Phone:     phone,
		SessionId: GenerateSessId(15),
		Code:      GenerateCode(),
	}
}

func GenerateCode() int {
	return rand.Intn(9000) + 1000
}

var letterRunes = []rune("abcdefghijklmnoprstuvwxyzABCDEFGHIJKLMNOPRSTUVWXYZ123456789")

func GenerateSessId(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}
