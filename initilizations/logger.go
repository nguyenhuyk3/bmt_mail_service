package initializations

import (
	"bmt_mail_service/global"
	"bmt_mail_service/pkgs/loggers"
)

func initLogger() {
	global.Logger = loggers.NewLogger(global.Config.ServiceSetting.LoggerSetting)
}
