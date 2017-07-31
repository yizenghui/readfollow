package conf

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

var (
	// Conf 常量
	Conf Config
	// DefaultConfigFile  默认匹配位置
	DefaultConfigFile = "conf.toml"
)

//Config 配置
type Config struct {
	ReleaseMode bool `toml:"release_mode"` //发布模式

	// 应用配置
	App app `toml:"app"`

	// PGSQL
	DB database `toml:"database"`

	Wechat wechat `toml:"wechat"`
}

type app struct {
	Name string `toml:"name"`
	Host string `toml:"host"`
	Port string `toml:"port"`
}

type database struct {
	DBName   string `toml:"dbname"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	Type     string `toml:"type"`
}

type wechat struct {
	AppID     string `toml:"appid"`
	AppSecret string `toml:"appsecret"`
	OriID     string `toml:"oriid"`
	Token     string `toml:"token"`
	AesKey    string `toml:"aeskey"`
}

func init() {
	InitConfig("")
}

// InitConfig initializes the app configuration by first setting defaults,
// then overriding settings from the app config file, then overriding
// It returns an error if any.
func InitConfig(configFile string) error {
	var retErr error
	if configFile == "" {
		configFile = DefaultConfigFile
	}

	// Set defaults.
	Conf = Config{
		ReleaseMode: false,
	}

	if _, err := os.Stat(configFile); err != nil {

		retErr = errors.New("config file err:" + err.Error())

	} else {

		configBytes, err := ioutil.ReadFile(configFile)
		if err != nil {
			retErr = errors.New("config load err:" + err.Error())
		}
		_, err = toml.Decode(string(configBytes), &Conf)
		if err != nil {
			retErr = errors.New("config decode err:" + err.Error())
		}
	}

	return retErr
}
