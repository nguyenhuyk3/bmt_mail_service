package dispatchers

import (
	"bmt_mail_service/dto/messages"
	"bmt_mail_service/global"
	"bmt_mail_service/utils/sender"
	"encoding/json"
	"log"
)

func DispatchEmail(emailMessage messages.EmailMessage) {
	switch emailMessage.Type {
	case global.REGISTRATION_OTP_EMAIL:
		var otpMsg messages.OtpMessage
		if err := json.Unmarshal(emailMessage.Payload, &otpMsg); err != nil {
			log.Printf("failed to unmarshal OTP payload: %v\n", err)
			return
		}
		err := sender.SendTemplateEmailOtp(
			[]string{otpMsg.Email},
			global.Config.Server.FromEmail,
			"registration_otp_email.html",
			global.REGISTRATION_PURPOSE,
			map[string]interface{}{
				"otp":             otpMsg.Otp,
				"from_email":      global.Config.Server.FromEmail,
				"expiration_time": otpMsg.ExpirationTime,
			})
		if err != nil {
			log.Printf("failed to send OTP email to %s: %v\n", otpMsg.Email, err)
		} else {
			log.Printf("successfully sent OTP email to %s\n", otpMsg.Email)
		}
	default:
		log.Printf("unrecognized email type: %s\n", emailMessage.Type)
	}
}
