package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	Turn "github.com/Ocas17/Kursach_UD"
	"github.com/Ocas17/Kursach_UD/internal/handler"
	"github.com/Ocas17/Kursach_UD/internal/repository"
	"github.com/Ocas17/Kursach_UD/internal/service"
	"github.com/Ocas17/Kursach_UD/pkg/postgres"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)

// @title Kursach_UD API
// @version 1.0
// @description API Server for List Application

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := InitConfig(); err != nil {
		logrus.Fatalf("config init error:%s", err)
	}

	if err := gotenv.Load(); err != nil {
		logrus.Fatalf("gotenv.Load err :%s", err)
	}
	db,err:=postgres.Newpostgresdb(postgres.Config{
		Host:    viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode: viper.GetString("db.sslmode"),

	})
	if err != nil {
		logrus.Fatalf("db init error:%s", err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(Turn.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}