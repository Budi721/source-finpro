package service

import (
    "context"
    "github.com/google/uuid"
    "github.com/itp-backend/backend-a-co-create/common/errors"
    "github.com/itp-backend/backend-a-co-create/config"
    "github.com/itp-backend/backend-a-co-create/contract"
    "github.com/itp-backend/backend-a-co-create/external/jwt_client"
    "github.com/itp-backend/backend-a-co-create/model"
    "github.com/itp-backend/backend-a-co-create/repository"
    "time"
)

type IUserService interface {
    Register(ctx context.Context, namaLengkap, username, password, topik string) (*model.Enrollment, error)
    Login(ctx context.Context, username, password, loginAs string) (*model.User, error)
    ChangePassword(ctx context.Context, email, password string) error
    BuildProfile(ctx context.Context, user *model.User) (*model.User, error)
    GetUserProfile(ctx context.Context, email string) (*model.User, error)
    IsValid(user *model.User) (bool, error)
    GetRepo() repository.UserRepository
}

func NewUserService(repo repository.UserRepository, appConfig *config.Config, jwtClient jwt_client.JWTClientInterface) IUserService {
    return &userService{
        r:         repo,
        appConfig: appConfig,
        jwtClient: jwtClient,
    }
}

type userService struct {
    r repository.UserRepository
    appConfig *config.Config
    jwtClient jwt_client.JWTClientInterface
}

func (service userService) Register(ctx context.Context, namaLengkap, username, password, topik string) (*model.Enrollment, error) {
    exists, err := service.r.DoesUsernameExist(ctx, username)
    if err != nil {
        return nil, err
    }
    if exists {
        return nil, errors.New("User already exists")
    }

    return service.r.CreateMinimal(ctx, namaLengkap, username, password, topik)
}

func (service userService) Login(ctx context.Context, username, password, loginAs string) (*model.User, error) {
    user, err := service.r.FindByUsername(ctx, username, password, loginAs)
    if err != nil {
       return &model.User{}, err
    }

    atClaims := contract.JWTMapClaim{
        Authorized: true,
        RequestID:  uuid.New().String(),
    }

    atClaims.Subject = user.Username
    atClaims.ExpiresAt =  time.Now().Add(time.Hour * 24).Unix()

    token, err := service.jwtClient.GenerateTokenStringWithClaims(atClaims, service.appConfig.JWTSecret)
    if err != nil {
        return &model.User{}, errors.NewBadRequestError(err)
    }

    return &model.User{Id: user.Id, AuthToken: token}, err
}

func (service userService) ChangePassword(ctx context.Context, email, password string) error {
    panic("implement me")
}

func (service userService) BuildProfile(ctx context.Context, user *model.User) (*model.User, error) {
    panic("implement me")
}

func (service userService) GetUserProfile(ctx context.Context, email string) (*model.User, error) {
    panic("implement me")
}

func (service userService) IsValid(user *model.User) (bool, error) {
    panic("implement me")
}

func (service userService) GetRepo() repository.UserRepository {
    panic("implement me")
}

