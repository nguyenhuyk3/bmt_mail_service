package messages

import "encoding/json"

type EmailMessage struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type OtpMessage struct {
	Email          string `json:"email" binding:"required"`
	Otp            string `json:"otp" binding:"required"`
	ExpirationTime int    `json:"expiration_time" binding:"required"`
}
