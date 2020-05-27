package config
import (
	"fmt"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

var ProviderSet = wire.NewSet(NewConfig)

// Init 初始化viper
func NewConfig(path string)(*viper.Viper,error){
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigFile(path)
	var err error
	if err =v.ReadInConfig();err==nil{
		fmt.Printf("use config file -> %s\n", v.ConfigFileUsed())
	}else {
		return nil,err
	}
	return v,err
}

