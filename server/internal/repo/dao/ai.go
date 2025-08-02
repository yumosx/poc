package dao

import "gorm.io/gorm"

type AIDao struct {
	db *gorm.DB
}

func NewAIDao(db *gorm.DB) *AIDao {
	return &AIDao{db: db}
}

type Chat struct {
	Id      int64  `gorm:"id"`
	Content string `gorm:"column:content"`
	Ctime   int64  `gorm:"column:ctime"`
}
