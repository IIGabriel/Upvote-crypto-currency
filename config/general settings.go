package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

func GetEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		zap.L().Warn("Error GetEnv():", zap.Error(err))
	}
	return os.Getenv(key)
}
func createDirectoryIfNotExist() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if _, err := os.Stat(fmt.Sprintf("%s/logs", path)); os.IsNotExist(err) {
		_ = os.Mkdir("logs", os.ModePerm)
	}
}

func getLogWriter() zapcore.WriteSyncer {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	file, err := os.OpenFile(fmt.Sprintf("/%s/logs/%s.txt", path, time.Now().Format("2006-01-02")),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	return zapcore.AddSync(file)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Local().Format("2006-01-02 15:04:05"))
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func InitLogger() {
	createDirectoryIfNotExist()
	core := zapcore.NewCore(getEncoder(), getLogWriter(), zapcore.DebugLevel)
	logg := zap.New(core, zap.AddCaller())

	zap.ReplaceGlobals(logg)
}
