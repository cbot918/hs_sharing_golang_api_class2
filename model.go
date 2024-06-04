package main

type Post struct {
	Id    int    `gorm:"primaryKey"`
	Title string `json:"title"`
	Body  string `json:"body"`
}
