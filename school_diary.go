package crud

type Teacher struct {
  id int `json:"-"`
  name string `json:"name"`
  rfid_card_id string `json:"rfid_card_id"`
  user_id int `json:"user_id"`
}

type Pupil struct {
  id int `json:"-"`
  name string `json:"name" binding:"required"`
  class_id int `json:"class_id"`
  birthdate string `json:"birthdate"`
  rfid_card_id string `json:"rfid_card_id"  binding:"required"`
  user_id string `json:"user_id"`
}

type Parent struct {
  id int `json:"-"`
  Title string `json:"title"`
  Description string `json:"description"`
  Done bool `json:"done"`
}
