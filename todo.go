package todo

type TodoList struct {
  id int `json:"-"`
  Title string `json:"title"`
  Description string `json:"description"`
}

type UserList struct {
  id int
  UserId int
  ListId int
}

type TodoItem struct {
  id int `json:"-"`
  Title string `json:"title"`
  Description string `json:"description"`
  Done bool `json:"done"`
}

type ListsItem struct {
  id int
  ItemId int
  ListId int
}
