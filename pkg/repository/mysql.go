package repository

import ( "fmt"
   _"github.com/go-sql-driver/mysql"
      "github.com/jmoiron/sqlx")

type Config struct{
  Host string
  Port string
  Username string
  Password string
  DBName string
  SSlMode string
}

const (
  userTable = "users"
  teacherTable = "teachers"
  pupilTable = "pupils"
  parentTable = "parent"
)

func NewMysqlDB(cfg Config) (*sqlx.DB, error) {
  dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
  fmt.Println(dsn)
  db, err := sqlx.Open("mysql", dsn)

    if err != nil {
      return nil, err
    }

    err = db.Ping()
    if err != nil {
      return nil, err
    }

    return db, nil
}
