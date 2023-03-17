package main

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	zapConfig := zap.NewProductionConfig()

	zapConfig.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder

	logger, err := zapConfig.Build()
	if err != nil {
		log.Fatal("error creating logger")
	}

	logger.Info("This is the core application!")
}
