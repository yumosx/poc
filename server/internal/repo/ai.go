package repo

import (
	"context"
	"github.com/google/uuid"
	"github.com/yumosx/poc/server/internal/domain"
	"github.com/yumosx/poc/server/internal/repo/dao"
)

type AIRepo struct {
	chatDao *dao.AIDao
	taskDao *dao.TaskDao
}

func NewAIRepo(dao *dao.AIDao, taskDao *dao.TaskDao) *AIRepo {
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
	err := repo.taskDao.Save(ctx, dao.Task{
		UUID:    task.UUID,
		Content: task.Content,
		State:   task.State,
		Type:    task.Type,
		Result:  task.Result,
	})

	if err != nil {
		return "", err
	}
	return task.UUID, nil
}
