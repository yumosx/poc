package repo

import (
	"context"
	"github.com/google/uuid"
	"github.com/yumosx/poc/internal/domain"
	dao2 "github.com/yumosx/poc/internal/repo/dao"
)

type AIRepo struct {
	chatDao *dao2.AIDao
	taskDao *dao2.TaskDao
}

func NewAIRepo(dao *dao2.AIDao, taskDao *dao2.TaskDao) *AIRepo {
	return &AIRepo{chatDao: dao, taskDao: taskDao}
}

func (repo *AIRepo) GetTask(ctx context.Context, uuid string) (domain.Task, error) {
	task, err := repo.taskDao.GetTask(ctx, uuid)
	if err != nil {
		return domain.Task{}, err
	}
	return domain.Task{UUID: task.UUID, Content: task.Content, State: task.State, Result: task.Result}, nil
}

func (repo *AIRepo) SaveTask(ctx context.Context, task domain.Task) (string, error) {
	if task.UUID == "" {
		task.UUID = uuid.New().String()
	}
	err := repo.taskDao.Save(ctx, dao2.Task{
		UUID:    task.UUID,
		Content: task.Content,
		State:   task.State,
		Result:  task.Result,
	})

	if err != nil {
		return "", err
	}
	return task.UUID, nil
}
