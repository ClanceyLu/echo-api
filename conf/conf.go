package conf

import (
	"fmt"

	"github.com/spf13/viper"
)

// Conf 配置文件
var Conf *viper.Viper

func init() {
	Conf = viper.New()
	Conf.SetConfigType("toml")
	Conf.AddConfigPath("./conf")
	Conf.SetConfigName("conf")
	err := Conf.ReadInConfig() // 搜索路径，并读取配置数据
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
}
