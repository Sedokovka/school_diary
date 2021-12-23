package crud

type Teacher struct {
  Id int `json:"-"`
  Name string `json:"name"`
  Rfid_card_id string `json:"rfid_card_id"`
  User_id int `json:"user_id"`
}

type Pupil struct {
  Id int `json:"-" db:"id"`
  Name string `json:"name" binding:"required" db:"name"`
  Class_id int `json:"class_id" db:"class_id"`
  Birthdate string `json:"birthdate" db:"birthdate"`
  Rfid_card_id string `json:"rfid_card_id"  binding:"required" db:"rfid_card_id"`
  User_id int `json:"user_id" db:"user_id"`
}

type Parent struct {
  Id int `json:"-"`
  Title string `json:"title"`
  Description string `json:"description"`
  Done bool `json:"done"`
}
