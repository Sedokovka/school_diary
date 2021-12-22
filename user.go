package crud

type User struct {
  Id int `json:"-" db:"id"`
  Login string `json:"login" binding:"required"`
  Name string `json:"name" binding:"required"`
  Password string `json:"password" binding:"required"`
  Salt string `json:"salt" binding:"required"`
}
