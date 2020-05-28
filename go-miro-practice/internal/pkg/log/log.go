package log

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var ProviderSet = wire.NewSet(NewLogger, NewOptions)

type Options struct {
	Filename string
	MaxSize int
	MaxBackups int
	MaxAge int
	Level string
	Stdout bool
}

func NewOptions(v *viper.Viper)(*Options,error){
	options := &Options{}
	var err error
	if err =v.UnmarshalKey("log",options);err!=nil{
		return nil, err
	}
	return options, err
}


func NewLogger(option *Options)(*zap.Logger,error){
	var err error
	level := zap.NewAtomicLevel()
	level.UnmarshalText([]byte(option.Level))
	var logger *zap.Logger

	fw := zapcore.AddSync(&lumberjack.Logger{
		Filename:   option.Filename,
		MaxSize:    option.MaxSize,
		MaxAge:     option.MaxAge,
		MaxBackups: option.MaxBackups,
	})
	cw := zapcore.Lock(os.Stdout)
	cores := make([]zapcore.Core, 0, 2)
	je := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	cores = append(cores, zapcore.NewCore(je, fw, level))

	if option.Stdout {
		ce := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		cores = append(cores, zapcore.NewCore(ce, cw, level))
	}

	core := zapcore.NewTee(cores...)
	logger = zap.New(core)
	zap.ReplaceGlobals(logger)
	return logger, err
}