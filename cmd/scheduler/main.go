package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
	"scheduler/internal/configs"
	"scheduler/internal/domain"
	"scheduler/internal/infra/api"
	"scheduler/internal/infra/http"
	"scheduler/internal/infra/telegram"
	logger2 "scheduler/pkg/logger"
	"time"
)

func init()  {
	if err := godotenv.Load("configs/.env.local"); err != nil {
		log.Print("No .env file found")
	}
}

func main()  {
	config := configs.New()
	logger, err := logger2.NewLogger(config)
	if err != nil {
		panic(err)
	}

	service := domain.NewService(logger)
	api := api.NewAdapter(service)
	httpAdapter := http.NewAdapter(logger, config, api)
	shutdown := make(chan error, 1)

	go func() {
		shutdown <- httpAdapter.Start()
	}()

	botAdapter := telegram.NewAdapter(logger, config, service)

	if err := botAdapter.Start(); err != nil {
		fmt.Println(err)
	}


	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)
	select {
	case s := <-sig:
		logger.Info("Signal received!", zap.Any("signal", s))
	case err := <-shutdown:
		logger.Error("Error while running an application!", zap.Error(err))
	}

	defer logger.Info("Graceful shutdown finished!")

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	if err := httpAdapter.Shutdown(ctx); err != nil {
		log.Panic(err)
	}

}
