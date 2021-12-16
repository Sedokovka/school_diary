package repository

import (
   _"github.com/go-sql-driver/mysql"
      "github.com/jmoiron/sqlx"
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

type Repository struct {
  Authorization
  Teacher
  Parent
  Pupil
  User
}


func NewRepository(db *sqlx.DB) *Repository {
  return &Repository{}
}
