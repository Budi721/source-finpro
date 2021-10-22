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
	FindProjectByCollaboratorUserId(collaboratorId uint64) ([]*model.Project, error)
	UpdateInvitationProject(project dto.ProjectInvitation) (*model.Project, error)
}

type projectRepository struct {
	DB *gorm.DB
}

func CreateProject(project *dto.Project) (*model.Project, error) {
	var invitedUserId []model.User
	var collaboratorUserId []model.User

	if len(project.InvitedUserId) > 0 {
		db.Find(&invitedUserId, project.InvitedUserId)
	}
	db.Find(&collaboratorUserId, project.Creator)

	projectToCreate := &model.Project{
		KategoriProject:  project.KategoriProject,
		NamaProject:      project.NamaProject,
		StartDate:        project.Date,
		LinkTrello:       project.LinkTrello,
		DeskripsiProject: project.DeskripsiProject,
		InvitedUserId:    project.InvitedUserId,
		CollaboratorUserId: []uint64{project.Creator},
		Creator:          project.Creator,
		UsersInvited:     invitedUserId,
		UsersCollaborator: collaboratorUserId,
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


	for _, invited := range project.UsersInvited {
		var user *model.User
		db.First(&user, invited.ID)
		project.InvitedUserId = append(project.InvitedUserId, invited.ID)
		project.InvitedUserName = append(project.InvitedUserName, user.Name)
	}

	for _, collaborator := range project.UsersCollaborator {
		var user *model.User
		db.First(&user, collaborator.ID)
		project.CollaboratorUserId = append(project.CollaboratorUserId, collaborator.ID)
		project.CollaboratorUserName = append(project.CollaboratorUserName, user.Name)
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

	err := db.Joins("JOIN user_invited on user_invited.project_id_project=projects.id_project").
			Where("user_invited.user_id = ?", invitedId).Preload("UsersInvited").Preload("UsersCollaborator").
			Find(&projects).Error


	if err != nil {
		log.Error(err)
		return nil, err
	}

	for _, project := range projects {

		for _, invited := range project.UsersInvited {
			var user *model.User
			db.First(&user, invited.ID)
			project.InvitedUserId = append(project.InvitedUserId, invited.ID)
			project.InvitedUserName = append(project.InvitedUserName, user.Name)
		}

		for _, collaborator := range project.UsersCollaborator {
			var user *model.User
			db.First(&user, collaborator.ID)
			project.CollaboratorUserId = append(project.CollaboratorUserId, collaborator.ID)
			project.CollaboratorUserName = append(project.CollaboratorUserName, user.Name)
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

	for _, invited := range projectUpdated.UsersInvited {
		var user *model.User
		db.First(&user, invited.ID)
		projectUpdated.InvitedUserId = append(projectUpdated.InvitedUserId, invited.ID)
		projectUpdated.InvitedUserName = append(projectUpdated.InvitedUserName, user.Name)
	}

	for _, collaborator := range projectUpdated.UsersCollaborator {
		var user *model.User
		db.First(&user, collaborator.ID)
		projectUpdated.CollaboratorUserId = append(projectUpdated.CollaboratorUserId, collaborator.ID)
		projectUpdated.CollaboratorUserName = append(projectUpdated.CollaboratorUserName, user.Name)
	}
	return &projectUpdated, nil
}

func FindAllProject() ([]*model.Project, error) {
	var projects []*model.Project
	if err := db.Table("projects").Find(&projects).Error; err != nil {
		log.Error(err)
		return projects, err
	}

	return projects, nil
}

func FindProjectByCollaboratorUserId(collaboratorId uint64) ([]*model.Project, error) {
	var projects []*model.Project

	err := db.Joins("JOIN user_collaborator on user_collaborator.project_id_project=projects.id_project").
			Where("user_collaborator.user_id = ?", collaboratorId).Preload("UsersInvited").Preload("UsersCollaborator").
			Find(&projects).Error


	if err != nil {
		log.Error(err)
		return nil, err
	}

	for _, project := range projects {
	
		for _, invited := range project.UsersInvited {
			var user *model.User
			db.First(&user, invited.ID)
			project.InvitedUserId = append(project.InvitedUserId, invited.ID)
			project.InvitedUserName = append(project.InvitedUserName, user.Name)
		}

		for _, collaborator := range project.UsersCollaborator {
			var user *model.User
			db.First(&user, collaborator.ID)
			project.CollaboratorUserId = append(project.CollaboratorUserId, collaborator.ID)
			project.CollaboratorUserName = append(project.CollaboratorUserName, user.Name)
		}
		
	}

	return projects, nil
}