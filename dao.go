package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Dao struct {
	Db *gorm.DB
}

func NewDao() (*Dao, error) {

	// 做 database 連線
	db, err := gorm.Open(sqlite.Open("database/test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 做個重新初始化
	db.AutoMigrate(&Post{})

	return &Dao{
		Db: db,
	}, nil
}

func (d *Dao) CreatePostDao(post *Post) (*Post, error) {

	result := d.Db.Create(&post)
	if result.Error != nil {
		return nil, result.Error
	}

	return post, nil
}

func (d *Dao) FindByIdPostDao(id int) (*Post, error) {
	var post Post

	result := d.Db.First(&post, id)
	if result.Error != nil {
		return &post, result.Error
	}

	return &post, nil
}
