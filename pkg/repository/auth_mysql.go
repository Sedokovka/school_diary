package repository

import (
  "fmt"
  "github.com/jmoiron/sqlx"
  "github.com/Sedokovka/simple-to-do-app"
)

type AuthMysql struct {
  db *sqlx.DB
}

func NewAuthMysql(db *sqlx.DB) *AuthMysql {
  return &AuthMysql{db: db}
}

func (r *AuthMysql) CreateUser(user crud.User) (int, error) {
  var id int

  queryInsert := fmt.Sprintf("INSERT INTO %s (login, name, password_hash) VALUES ('%s', '%s', '%s')", userTable, user.Login, user.Name, user.Password)
  _, err := r.db.Query(queryInsert)
  if err != nil {
    return 0, err
  }
  query := fmt.Sprintf("SELECT id FROM %s WHERE login = '%s' AND name = '%s'", userTable, user.Login, user.Name);
  row := r.db.QueryRow(query)
  if err := row.Scan(&id); err != nil {
    return 0, err
  }
  return id, nil;
}

func (r *AuthMysql) GetUser(username, password string)(crud.User, error) {
  var user crud.User
  query := fmt.Sprintf("SELECT id FROM %s WHERE login = '%s' AND password_hash = '%s'", userTable, username, password);
  err := r.db.Get(&user, query)

  return user, err;
}
