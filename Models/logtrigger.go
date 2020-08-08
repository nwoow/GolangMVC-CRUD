package models

type Book struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}
type Logger struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Title    string `json:"title"`
	Filename string `json:"filename"`
	Error    string `json:"error"`
}

type Logconfigure struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Title    string `json:"title"`
	Filename string `json:"filename"`
	Logtype  string `json:"logtype"`
	Loglevel string `json:"loglevel"`
	Sitename string `json:"sitename"`
}
