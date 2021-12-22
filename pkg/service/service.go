package service

import (
  "github.com/Sedokovka/simple-to-do-app"
  "github.com/Sedokovka/simple-to-do-app/pkg/repository"
)

type Authorization interface {
  CreateUser(user crud.User) (int, error)
  GenerateToken(username, password string) (string, error)
  ParseToken(token string) (int, error)
}

type Teacher interface {

}
type User interface {

}
type Pupil interface {
  CreatePupil(userId int, pupil crud.Pupil) (int, error)

}
type Parent interface {

}

type Service struct {
  Authorization
  Teacher
  Parent
  Pupil
  User
}

func NewService(repos *repository.Repository) *Service {
  return &Service{
    Authorization: NewAuthService(repos.Authorization),
    Pupil: NewPupilService(repos.Authorization),
  }
}
