package logging

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	Logger *zap.Logger
)

func InitLog() {
	// 日志切割
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./error.log",
		MaxSize:    10,
		MaxBackups: 30,
		MaxAge:     30,
		Compress:   false,
	}

	encodingConfig := zap.NewProductionEncoderConfig()
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encodingConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger)),
		zapcore.DebugLevel,
	)

	zap.ReplaceGlobals(zap.New(core, zap.AddCaller()))
	Logger = zap.L()
}
