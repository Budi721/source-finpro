package service

import (
    "context"
    "github.com/itp-backend/backend-a-co-create/model"
    "github.com/itp-backend/backend-a-co-create/repository"
    log "github.com/sirupsen/logrus"
)

type IEnrollmentService interface {
    GetByStatus(ctx context.Context, status string) ([]*model.Enrollment, error)
}

func NewEnrollmentService(repo repository.IEnrollmentRepository) IEnrollmentService {
    return &enrollmentService{r: repo}
}

type enrollmentService struct {
    r repository.IEnrollmentRepository
}

func (service *enrollmentService) GetByStatus(ctx context.Context, status string) ([]*model.Enrollment, error) {
    enrollments, err := service.r.FindAllByStatus(ctx, status)
    if err != nil {
        log.Error(err)
        return []*model.Enrollment{}, err
    }

    return enrollments, nil
}

