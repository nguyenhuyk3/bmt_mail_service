package global

import (
	"bmt_mail_service/pkgs/loggers"
	"bmt_mail_service/pkgs/settings"
)

var (
	Config settings.Config
	Logger *loggers.LoggerZap
)
