package initializations

import (
	messagebroker "bmt_mail_service/initilizations/message_broker"
)

func Run() {
	loadConfigs()
	initLogger()

	messagebroker.InitReaders()

	select {}
}
