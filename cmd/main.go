package main

import (
  "log"
  "github.com/Sedokovka/simple-to-do-app"
  "github.com/Sedokovka/simple-to-do-app/pkg/handler"
  "github.com/Sedokovka/simple-to-do-app/pkg/repository"
  "github.com/Sedokovka/simple-to-do-app/pkg/service"
)

func main() {
  repos := repository.NewRepository()
  services := service.NewService(repos)
  handlers := handler.NewHandler(services)

  srv := new(todo.Server)
  if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
    log.Fatal("Mistake")
  }
}
