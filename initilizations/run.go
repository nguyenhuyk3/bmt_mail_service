package initializations

import (
	messagebroker "bmt_mail_service/initilizations/message_broker"
	"log"
)

func Run() {
	loadConfigs()

	if err := messagebroker.InitReader(); err != nil {
		log.Fatalf("failed to start consumer: %v", err)
	}
}
