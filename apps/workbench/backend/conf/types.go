package conf

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	Charset  string `yaml:"charset"`
}

type Config struct {
	Database         DBConfig         `yaml:"database"`
	EmailConfig      EmailConfig      `yaml:"email_config"`
	PermissionConfig PermissionConfig `yaml:"permission_config"`
}

type EmailConfig struct {
	Path string `yaml:"path"`
}

type PermissionConfig struct {
	Path string `yaml:"path"`
}
