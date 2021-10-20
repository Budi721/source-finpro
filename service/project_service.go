package service

import (
	"github.com/itp-backend/backend-a-co-create/dto"
	"github.com/itp-backend/backend-a-co-create/model"
	"github.com/itp-backend/backend-a-co-create/repository"
	log "github.com/sirupsen/logrus"
)

type IProjectService interface {
	CreateProject(project *dto.Project) (*model.Project, error)
	GetDetailProject(projectId int) (*model.Project, error)
	DeleteProject(projectId int) error
	GetProjectByInvitedUser(invitedId int) ([]*model.Project, error)
	UpdateInvitation(project dto.ProjectInvitation) (*model.Project, error)
}

type projectService struct {
	repo   repository.IProjectRepository
}

func CreateProject(project *dto.Project) (*model.Project, error) {
	projectToCreate, err := repository.CreateProject(project)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return projectToCreate, nil
}

func GetDetailProject(projectId int) (*model.Project, error) {
	project, err := repository.FindProjectById(projectId)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return project, nil
}

func DeleteProject(projectId int) error {
	err := repository.DeleteProject(projectId)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func GetProjectByInvitedUser(invitedId int) ([]*model.Project, error) {
	project, err := repository.FindProjectByInvitedUserId(invitedId)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return project, nil
}

func UpdateInvitation(project dto.ProjectInvitation) (*model.Project, error) {
	projectUpdated, err := repository.UpdateInvitationProject(project)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return projectUpdated, nil
}


func GetAllProject() ([]*model.Project, error) {
	projects, err := repository.FindAllProject()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return projects, nil
}
