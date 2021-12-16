package service

import (
  "github.com/Sedokovka/simple-to-do-app/pkg/repository"
)

type Authorization interface {

}

type Teacher interface {

}
type User interface {

}
type Pupil interface {

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
  return &Service{}
}
