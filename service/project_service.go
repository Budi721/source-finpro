package service

import (
    "context"
    "github.com/itp-backend/backend-a-co-create/model"
    "github.com/itp-backend/backend-a-co-create/repository"
    log "github.com/sirupsen/logrus"
)

type IProjectService interface {
	CreateProject(ctx context.Context, project *model.Project) (*model.Project, error)
	GetDetailProject(ctx context.Context, projectId int) (*model.Project, error)
	DeleteProject(ctx context.Context, projectId int)
	// Terima undangan
	GetProjectByInvitedUser(ctx context.Context, invitedId int) (*model.Project, error)
}

func NewProjectService(projectRepository repository.IProjectRepository) IProjectService {
	return projectService{
		repo:   projectRepository,
	}
}

type projectService struct {
	repo   repository.IProjectRepository
}

func (service projectService) CreateProject(ctx context.Context, project *model.Project) (*model.Project, error) {
	project, err := service.repo.Create(ctx, project)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return project, nil
}

func (service projectService) GetDetailProject(ctx context.Context, projectId int) (*model.Project, error) {
	panic("implement me")
}

func (service projectService) DeleteProject(ctx context.Context, projectId int) {
	panic("implement me")
}

func (service projectService) GetProjectByInvitedUser(ctx context.Context, invitedId int) (*model.Project, error) {
	panic("implement me")
}
