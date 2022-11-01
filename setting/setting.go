package setting

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Name string `mapstructure:"name"`
	Mode string `mapstructure:"mode"`
	Version string `mapstructure:"version"`
	Port int `mapstructure:"port"`
    // 添加时间
	StartTime string `mapstructure:"start_time"`
	MachineID int64 `mapstructure:"machine_id"`
	EmailSecretCode string `mapstructure:"email_secret_code"`
	AccessKey string `mapstructure:"access_key"`
	SecretKey string `mapstructure:"secret_key"`
	Bucket string `mapstructure:"bucket"`
	ImgUrl string `mapstructure:"img_url"`
	*LogConfig `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level string `mapstructure:"level"`
	Filename string `mapstructure:"filename"`
	MaxSize int `mapstructure:"max_size"`
	MaxAge int `mapstructure:"max_age"`
	MaxBackups int `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host string `mapstructure:"host"`
	User string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DbName string `mapstructure:"dbname"`
	Port int `mapstructure:"port"`
	MaxOpenConns int `mapstructure:"max_open_conns"`
	MaxIdleConns int `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port int `mapstructure:"port"`
	DB int `mapstructure:"db"`
	PoolSize int `mapstructure:"pool_size"`
}

func Init(filePath string) (err error) {
	viper.SetConfigFile(filePath)
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("viper.ReadInConfig() faild, err:%v\n", err)
		return
	}
	// 把读取到的配置信息反序列化到conf变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal faild, err:%v\n", err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func (in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal faild, err:%v\n", err)
		}
	})
	return
}