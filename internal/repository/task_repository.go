package repository

import (
	"github.com/cmartinez/GolangRestApiTesting/internal/domain"
	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) GetTasks() ([]domain.Task, error) {
	var tasks []domain.Task
	if err := r.db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TaskRepository) GetTaskByID(id uint) (*domain.Task, error) {
	var task domain.Task
	if err := r.db.First(&task, id).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepository) CreateTask(task *domain.Task) error {
	return r.db.Create(task).Error
}

func (r *TaskRepository) UpdateTask(task *domain.Task) error {
	return r.db.Model(task).Updates(task).Error
}

func (r *TaskRepository) DeleteTask(id uint) error {
	return r.db.Delete(&domain.Task{}, id).Error
}
