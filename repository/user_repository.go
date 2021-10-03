package repository

import (
    "context"
    "github.com/itp-backend/backend-a-co-create/common/errors"
    "github.com/itp-backend/backend-a-co-create/model"
    errCheck "github.com/pkg/errors"
    "gorm.io/gorm"
)

type UserRepository interface {
    FindByID(ctx context.Context, id uint) (*model.User, error)
    BuildProfile(ctx context.Context, user *model.User) (*model.User, error)
    CreateMinimal(ctx context.Context, namaLengkap, username, password, topik string) (*model.Enrollment, error)
    FindByUsername(ctx context.Context, username, password, loginAs string) (*model.User, error)
    DoesUsernameExist(ctx context.Context, username string) (bool, error)
    ChangePassword(ctx context.Context, email, password string) error
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &IUserRepository{DB: db}
}

type IUserRepository struct {
    DB *gorm.DB
}

func (repo *IUserRepository) FindByID(ctx context.Context, id uint) (*model.User, error) {
    panic("implement me")
}

func (repo *IUserRepository) BuildProfile(ctx context.Context, user *model.User) (*model.User, error) {
    panic("implement me")
}

func (repo *IUserRepository) CreateMinimal(ctx context.Context, namaLengkap, username, password, topik string) (*model.Enrollment, error) {
    u := &model.User{
        Username: username,
    }
    u.Password = u.SetPassword(password)

    tx := repo.DB.Begin()
    result := tx.Create(&u)
    if result.Error != nil {
        tx.Rollback()
        return nil, result.Error
    }

    enrollment := &model.Enrollment{
        IdUser:           u.Id,
        NamaLengkap:      namaLengkap,
        Username:         username,
        TopikDiminati:    topik,
        EnrollmentStatus: 0,
    }
    enrollment.Password = u.SetPassword(password)

    result = tx.Create(&enrollment)
    if result.Error != nil {
        tx.Rollback()
        return nil, result.Error
    }
    tx.Commit()
    return enrollment, nil
}

func (repo *IUserRepository) FindByUsername(ctx context.Context, username, password, loginAs string) (*model.User, error) {
    u := &model.User{}
    result := repo.DB.Where("username = ? AND login_as = ?", username, loginAs).First(u)

    if err := u.ComparePassword(password); err != nil {
        return nil, errors.NewInternalError(err, "Error: not found")
    }

    switch result.Error {
    case nil:
        return u, nil
    case gorm.ErrRecordNotFound:
        return nil, errors.NewInternalError(result.Error, "Error: not found")
    default:
        return nil, errors.NewInternalError(result.Error, "Error: database error")
    }
}

func (repo *IUserRepository) DoesUsernameExist(ctx context.Context, username string) (bool, error) {
    u := &model.User{}
    if err := repo.DB.Where("username = ?", username).First(u).Error; errCheck.Is(err, gorm.ErrRecordNotFound) {
        return false, nil
    }

    return true, nil
}

func (repo *IUserRepository) ChangePassword(ctx context.Context, email, password string) error {
    panic("implement me")
}
