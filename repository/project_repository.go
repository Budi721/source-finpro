package repository

import (
    "context"
    "github.com/itp-backend/backend-a-co-create/model"
    log "github.com/sirupsen/logrus"
    "gorm.io/gorm"
)

type IProjectRepository interface {
    Create(ctx context.Context, project *model.Project) (*model.Project, error)
    FindById(ctx context.Context, idProject int) (*model.Project, error)
    Delete(ctx context.Context, idProject int) error
    // Approve Invited Project
    FindByInvitedUserId(ctx context.Context, invitedId int) (*model.Project, error)
}

func NewProjectRepository(db *gorm.DB) IProjectRepository {
    return projectRepository{DB: db}
}

type projectRepository struct {
    DB *gorm.DB
}

func (repo projectRepository) Create(ctx context.Context, project *model.Project) (*model.Project, error) {
    var invitedUserId []model.User
    var collaboratorUserId []model.User

    if len(project.InvitedUserId) > 0 {
        repo.DB.Find(&invitedUserId, project.InvitedUserId)
    }
    if len(project.CollaboratorUserId) > 0 {
        repo.DB.Find(&collaboratorUserId, project.CollaboratorUserId)
    }

    p := &model.Project{
        KategoriProject:    project.KategoriProject,
        NamaProject:        project.NamaProject,
        TanggalMulai:       project.TanggalMulai,
        LinkTrello:         project.LinkTrello,
        DeskripsiProject:   project.DeskripsiProject,
        InvitedUserId:      project.InvitedUserId,
        CollaboratorUserId: project.CollaboratorUserId,
        Admin:              project.Admin,
        UsersInvited:       invitedUserId,
        UsersCollaborator:  collaboratorUserId,
    }

    result := repo.DB.Create(&p)
    if result.Error != nil {
        log.Error(result.Error)
        return nil, result.Error
    }
    return p, nil
}

func (repo projectRepository) FindById(ctx context.Context, idProject int) (*model.Project, error) {
    var project model.Project
    project.IdProject = idProject

    if err := repo.DB.First(&project).Error; err != nil {
        log.Error(err)
        return nil, err
    }

    return &project, nil
}

func (repo projectRepository) Delete(ctx context.Context, idProject int) error {
    var project model.Project
    project.IdProject = idProject

    if err := repo.DB.Delete(&project).Error; err != nil {
        log.Error(err)
        return err
    }

    return nil
}

func (repo projectRepository) FindByInvitedUserId(ctx context.Context, invitedId int) (*model.Project, error) {
    var project model.Project
    /**
    @todo
    ubah data dari json ke slice
    filtering by invited project id
     */
    if err := repo.DB.First(&project).Error; err != nil {
        log.Error(err)
        return nil, err
    }

    return &project, nil
}


