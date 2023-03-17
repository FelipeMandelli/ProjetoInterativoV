package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/FelipeMandelli/ProjetoInterativoV/cmd/api/internal/config"
	"github.com/FelipeMandelli/ProjetoInterativoV/cmd/api/internal/services"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/sync/errgroup"
)

func main() {
	zapConfig := zap.NewProductionConfig()

	zapConfig.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder

	logger, err := zapConfig.Build()
	if err != nil {
		log.Fatal("error creating logger")
		os.Exit(1)
	}

	logger.Info("This is the API appliction!")

	ctx, stopCtx := context.WithCancel(context.Background())

	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGHUP, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

		signalReceived := <-sig

		logger.Warn("stopping context", zap.String("signal received", signalReceived.String()))
		stopCtx()
	}()

	err = config.SetupConfigurations()
	if err != nil {
		log.Fatal("error setting up configurations: ", err)
	}

	errorGroup, ctx := errgroup.WithContext(ctx)

	httpServer := &http.Server{
		Addr:    viper.GetString(config.AddressKey),
		Handler: services.CreateRouter(),
	}

	errorGroup.Go(func() error {
		logger.Info("serving API on " + httpServer.Addr)
		return httpServer.ListenAndServe()
	})

	errorGroup.Go(func() error {
		<-ctx.Done()
		logger.Info("context has stopped")
		return httpServer.Shutdown(context.Background())
	})

	if err := errorGroup.Wait(); err != nil {
		logger.Fatal("exit reason", zap.Error(err))
	}
=
}
