package repository

import (
   _"github.com/go-sql-driver/mysql"
      "github.com/jmoiron/sqlx"
      "github.com/Sedokovka/simple-to-do-app"
)

type Authorization interface {
  CreateUser(user crud.User)(int, error)
  GetUser(username, password string)(crud.User, error)
}

type Teacher interface {

}
type User interface {

}
type Pupil interface {
  CreatePupil(userId int, pupil crud.Pupil) (int, error)
  GetPupilById(userId, pupilId int) (crud.Pupil, error)
  GetAllPupils(userId int) ([]crud.Pupil, error)
}
type Parent interface {

}

type Repository struct {
  Authorization
  Teacher
  Parent
  Pupil
  User
}


func NewRepository(db *sqlx.DB) *Repository {
  return &Repository{
    Authorization: NewAuthMysql(db),
    Pupil: NewPupilMysql(db),
  }
}
