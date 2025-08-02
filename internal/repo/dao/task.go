package dao

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type TaskDao struct {
	db *gorm.DB
}

func NewTaskDao(db *gorm.DB) *TaskDao {
	return &TaskDao{db: db}
}

// Save 创建一个 task 或者去更新一个 task
func (t *TaskDao) Save(ctx context.Context, task Task) error {
	now := time.Now().Unix()
	task.Ctime = now
	task.Utime = now
	return t.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "uuid"}},
		DoUpdates: clause.Assignments(map[string]any{
			"state":  task.State,
			"utime":  task.Utime,
			"result": task.Result,
		}),
	}).Create(&task).Error
}

// GetTask 获取对应的 task
func (t *TaskDao) GetTask(ctx context.Context, uuid string) (Task, error) {
	var task Task
	err := t.db.WithContext(ctx).Where("uuid = ?", uuid).First(&task).Error
	return task, err
}

type Task struct {
	Id      int64  `gorm:"column:id;autoIncrement;primaryKey"`
	Uid     int64  `gorm:"column:uid"`
	UUID    string `gorm:"column:uuid;uniqueIndex;type:varchar(36)"`
	Content string `gorm:"column:content"`
	Result  string `gorm:"column:result"`
	State   string `gorm:"column:state"`
	Ctime   int64  `gorm:"column:ctime"`
	Utime   int64  `gorm:"column:utime"`
}
