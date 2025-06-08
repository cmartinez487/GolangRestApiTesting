package service

import (
	"github.com/cmartinez/GolangRestApiTesting/internal/domain"
	"github.com/cmartinez/GolangRestApiTesting/internal/repository"
)

type TaskService struct {
	repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) GetTasks() ([]domain.Task, error) {
	return s.repo.GetTasks()
}
func (s *TaskService) GetTaskByID(id uint) (*domain.Task, error) {
	return s.repo.GetTaskByID(id)
}
func (s *TaskService) CreateTask(task *domain.Task) error {
	return s.repo.CreateTask(task)
}
func (s *TaskService) UpdateTask(task *domain.Task) error {
	return s.repo.UpdateTask(task)
}
func (s *TaskService) DeleteTask(id uint) error {
	return s.repo.DeleteTask(id)
}
