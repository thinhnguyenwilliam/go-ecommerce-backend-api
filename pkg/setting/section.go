// go-ecommerce-backend-api/pkg/setting/section.go
package setting

type Config struct {
	Server  Server        `mapstructure:"server"`
	DBMySQL DBMySQL       `mapstructure:"dbmysql"`
	Logger  LoggerSetting `mapstructure:"logger"`
}

type LoggerSetting struct {
	LogLevel    string `mapstructure:"logLevel"`
	FileLogName string `mapstructure:"fileLogName"`
	MaxSize     int    `mapstructure:"maxSize"`
	MaxBackups  int    `mapstructure:"maxBackups"`
	MaxAge      int    `mapstructure:"maxAge"`
	Compress    bool   `mapstructure:"compress"`
}

type Server struct {
	Port int `mapstructure:"port"`
}

type DBMySQL struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	DBName          string `mapstructure:"dbname"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime"`
}
