package crud

type User struct {
  Id int `json:"-"`
  Login string `json:"login" binding:"required"`
  Name string `json:"name" binding:"required"`
  Username string `json:"username" binding:"required"`
  Password_hash string `json:"password_hash" binding:"required"`
  Salt string `json:"salt" binding:"required"`
}
