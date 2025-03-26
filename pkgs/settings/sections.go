package settings

type Config struct {
	Server         serverSetting
	ServiceSetting serviceSetting
}

type serviceSetting struct {
	MailSetting mailSetting `mapstructure:"mail"`
}

type serverSetting struct {
	ServerPort string `mapstructure:"SERVER_PORT"`
	FromEmail  string `mapstructure:"FROM_EMAIL"`
	APIKey     string `mapstructure:"API_KEY"`
	Issuer     string `mapstructure:"ISS"`
}

type mailSetting struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username,omitempty"`
	Password string `mapstructure:"password,omitempty"`
}
