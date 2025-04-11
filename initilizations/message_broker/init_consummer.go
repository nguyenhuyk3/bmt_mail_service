package messagebroker

import (
	"bmt_mail_service/dto/messages"
	"bmt_mail_service/global"
	"bmt_mail_service/utils/dispatchers"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

var topics = []string{
	global.REGISTRATION_OTP_EMAIL_TOPIC,
	global.FORGOT_PASSWORD_OTP_EMAIL_TOPIC,
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
		GroupID:        global.MAIL_SERVICE_GROUP,
		Topic:          topic,
		CommitInterval: time.Second * 5,
	})
	defer reader.Close()

	for {
		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			global.Logger.Error("error reading message", zap.Any("err", err))
			continue
		}

		var emailMessage messages.EmailMessage
		err = json.Unmarshal(message.Value, &emailMessage)
		if err != nil {
			global.Logger.Error("failed to unmarshal message", zap.Any("err", err))
			continue
		}

		switch topic {
		case global.REGISTRATION_OTP_EMAIL_TOPIC:
			go dispatchers.SendRegistrationOtpEmail(emailMessage)
		case global.FORGOT_PASSWORD_OTP_EMAIL_TOPIC:
			go dispatchers.SendForgotPasswordOtpEmail(emailMessage)
		default:
			global.Logger.Error(fmt.Sprintf("unknown topic: %s", topic))
		}
	}
}
