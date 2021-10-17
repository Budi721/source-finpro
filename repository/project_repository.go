package repository

import (
	"errors"
	"github.com/itp-backend/backend-a-co-create/dto"
	"github.com/itp-backend/backend-a-co-create/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IProjectRepository interface {
	CreateProject(project *dto.Project) (*model.Project, error)
	FindProjectById(idProject int) (*model.Project, error)
	DeleteProject(idProject int) error
	FindProjectByInvitedUserId(invitedId int) ([]*model.Project, error)
	UpdateInvitationProject(project dto.ProjectInvitation) (*model.Project, error)
}

type projectRepository struct {
	DB *gorm.DB
}

func CreateProject(project *dto.Project) (*model.Project, error) {
	var invitedUserId []model.User

	if len(project.InvitedUserId) > 0 {
		db.Find(&invitedUserId, project.InvitedUserId)
	}

	projectToCreate := &model.Project{
		KategoriProject:  project.KategoriProject,
		NamaProject:      project.NamaProject,
		StartDate:        project.Date,
		LinkTrello:       project.LinkTrello,
		DeskripsiProject: project.DeskripsiProject,
		InvitedUserId:    project.InvitedUserId,
		Creator:          project.Creator,
		UsersInvited:     invitedUserId,
	}
	result := db.Create(&projectToCreate)
	if result.Error != nil {
		log.Error(result.Error)
		return nil, result.Error
	}
	return projectToCreate, nil
}

func FindProjectById(idProject int) (*model.Project, error) {
	var project model.Project
	project.IdProject = idProject

	if err := db.Preload("UsersInvited").Preload("UsersCollaborator").First(&project).Error; err != nil {
		log.Error(err)
		return nil, err
	}

	for _, collaborator := range project.UsersCollaborator {
		project.CollaboratorUserId = append(project.CollaboratorUserId, collaborator.ID)
	}

	for _, invited := range project.UsersInvited {
		project.InvitedUserId = append(project.InvitedUserId, invited.ID)
	}
	return &project, nil
}

func DeleteProject(idProject int) error {
	var project model.Project
	project.IdProject = idProject

	err := db.Model(&project).Association("UsersInvited").Clear()
	if err != nil {
		log.Error(err)
		return err
	}

	err = db.Model(&project).Association("UsersCollaborator").Clear()
	if err != nil {
		log.Error(err)
		return err
	}

	if err := db.First(&project).Error; err != nil {
		log.Error(err)
		return err
	}

	if err := db.Delete(&project).Error; err != nil {
		log.Error(err)
		return errors.New("cannot delete record")
	}

	return nil
}

func FindProjectByInvitedUserId(invitedId int) ([]*model.Project, error) {
	var projects []*model.Project
	var user []model.User
	db.Find(&user, invitedId)
	if err := db.Where(&model.Project{UsersInvited: user}).Preload("UsersInvited").Preload("UsersCollaborator").Find(&projects).Error; err != nil {
		log.Error(err)
		return nil, err
	}

	for _, project := range projects {
		for _, collaborator := range project.UsersCollaborator {
			project.CollaboratorUserId = append(project.CollaboratorUserId, collaborator.ID)
		}

		for _, invited := range project.UsersInvited {
			project.InvitedUserId = append(project.InvitedUserId, invited.ID)
		}
	}

	return projects, nil
}

func UpdateInvitationProject(project dto.ProjectInvitation) (*model.Project, error) {
	var projectUpdated model.Project
	var invitedUserId []model.User

	db.Find(&invitedUserId, project.IdUser)
	projectUpdated.IdProject = project.IdProject

	db.Model(&projectUpdated).Association("UsersInvited").Delete(&model.User{
		GormModel: model.GormModel{
			ID: uint64(project.IdUser),
		},
	})
	db.Model(&projectUpdated).Association("UsersCollaborator").Append(&invitedUserId)

	if err := db.Preload("UsersInvited").Preload("UsersCollaborator").First(&projectUpdated).Error; err != nil {
		log.Error(err)
		return &model.Project{}, err
	}

	for _, collaborator := range projectUpdated.UsersCollaborator {
		projectUpdated.CollaboratorUserId = append(projectUpdated.CollaboratorUserId, collaborator.ID)
	}

	for _, invited := range projectUpdated.UsersInvited {
		projectUpdated.InvitedUserId = append(projectUpdated.InvitedUserId, invited.ID)
	}
	return &projectUpdated, nil
}
