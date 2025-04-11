package settings

type Config struct {
	Server         serverSetting
	ServiceSetting serviceSetting
}

type serviceSetting struct {
	MailSetting   mailSetting   `mapstructure:"mail"`
	KafkaSetting  kafkaSetting  `mapstructure:"kafka"`
	LoggerSetting LoggerSetting `mapstructure:"zap_log"`
}

type serverSetting struct {
	ServerPort string `mapstructure:"SERVER_PORT"`
	FromEmail  string `mapstructure:"FROM_EMAIL"`
}

type mailSetting struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username,omitempty"`
	Password string `mapstructure:"password,omitempty"`
}

type kafkaSetting struct {
	KafkaBroker_1 string `mapstructure:"kafka_broker_1"`
	KafkaBroker_2 string `mapstructure:"kafka_broker_2"`
	KafkaBroker_3 string `mapstructure:"kafka_broker_3"`
}

type LoggerSetting struct {
	LogLevel    string `mapstructure:"log_level"` // examples: "debug", "info", "warn", "error"
	FileLogName string `mapstructure:"file_log_name"`
	MaxBackups  int    `mapstructure:"max_backups"`
	MaxAge      int    `mapstructure:"max_age"` // number of days to keep log
	MaxSize     int    `mapstructure:"max_size"`
	Compress    bool   `mapstructure:"compress"` // compress old logs
}
