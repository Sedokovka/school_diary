package main

import (
  "log"
  "github.com/spf13/viper"
  "github.com/joho/godotenv"
  "github.com/Sedokovka/simple-to-do-app"
  "github.com/Sedokovka/simple-to-do-app/pkg/handler"
  "github.com/Sedokovka/simple-to-do-app/pkg/repository"
  "github.com/Sedokovka/simple-to-do-app/pkg/service"
  "os"
)

func main() {
  if err := initConfig(); err != nil {
    log.Fatal("error initializing configs: %s", err.Error())
  }

  if err := godotenv.Load(); err != nil {
      log.Fatal("error load configs: %s", err.Error())
  }

  db, err := repository.NewMysqlDB(repository.Config{
    Host: viper.GetString("db.host"),
    Port: viper.GetString("db.port"),
    Username: viper.GetString("db.username"),
    Password:os.Getenv("DB_PASSWORD"),
    DBName: viper.GetString("db.dbname"),
    SSlMode: viper.GetString("db.sslmode"),
  })

  if err != nil {
    log.Fatal(err)
  }

  repos := repository.NewRepository(db)
  services := service.NewService(repos)
  handlers := handler.NewHandler(services)

  srv := new(crud.Server)
  if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
    log.Fatal("Mistake")
  }
}

func initConfig() error {
  viper.AddConfigPath("configs")
  viper.SetConfigName("config")
  return viper.ReadInConfig()
}
