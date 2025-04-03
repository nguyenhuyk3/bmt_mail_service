package messagebroker

import (
	"bmt_mail_service/dto/messages"
	"bmt_mail_service/global"
	"bmt_mail_service/utils/dispatchers"
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

var topics = []string{
	global.REGISTRATION_OTP_EMAIL,
	global.FORGOT_PASSWORD_OTP_EMAIL,
}

func InitReaders() {
	log.Println("=============== Mail Service is listening for messages... ===============")

	for _, topic := range topics {
		go startReader(topic)
	}
}

func startReader(topic string) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{
			global.Config.ServiceSetting.KafkaSetting.KafkaBroker_1,
		},
		GroupID: global.MAIL_SERVICE_GROUP,
		Topic:   topic,
	})
	defer reader.Close()

	for {
		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("error reading message: %v\n", err)
			continue
		}

		var emailMessage messages.EmailMessage
		err = json.Unmarshal(message.Value, &emailMessage)
		if err != nil {
			log.Printf("failed to unmarshal message: %v\n", err)
			continue
		}

		switch topic {
		case global.REGISTRATION_OTP_EMAIL:
			go dispatchers.SendRegistrationOtpEmail(emailMessage)
		case global.FORGOT_PASSWORD_OTP_EMAIL:
			go dispatchers.SendForgotPasswordOtpEmail(emailMessage)
		default:
			log.Printf("unknown topic: %s", topic)
		}
	}
}
