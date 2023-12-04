package conf

import (
	"ProxyHttp/app/global"
	"bytes"
	"github.com/spf13/viper"
)
import _ "embed"

//go:embed conf.yaml
var conf []byte

func InitConfig() {
	viper.SetConfigType("yaml")

	err := viper.ReadConfig(bytes.NewReader(conf))

	if nil != err {
		panic(err)
	}

	global.ProxyUrl = viper.Get("proxy.url").(string)
	global.AppPort = viper.Get("app.port").(int)
}
