package database

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"github.com/baxiang/go-miro-practice/internal/pkg/models"
)


var ProviderSet = wire.NewSet(NewDatabase,NewOptions)

type Options struct {
	URL   string `yaml:"url"`
	Debug bool
}

func NewOptions(v *viper.Viper, logger *zap.Logger) (*Options, error) {
   o :=&Options{}
   var err  error
   if err =v.UnmarshalKey("db",o);err!=nil{
   	 return nil,errors.Wrap(err,"unmarshal db option error")
   }
	logger.Info("load database options success", zap.String("url", o.URL))
	return o, err
}

func NewDatabase(o Options)(*gorm.DB,error){
	var err error
	db, err := gorm.Open("mysql", o.URL)
	if err!=nil{
		return nil,errors.Wrap(err,"gorm open database connection error")
	}
	if o.Debug{
		db = db.Debug()
	}
	db.AutoMigrate(&models.Detail{})
	db.AutoMigrate(&models.Rating{})
	db.AutoMigrate(&models.Review{})
	return db,nil
}
