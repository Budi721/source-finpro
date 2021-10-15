package service

import (
    "github.com/itp-backend/backend-a-co-create/model"
    "github.com/itp-backend/backend-a-co-create/repository"
    log "github.com/sirupsen/logrus"
)

type EnrollmentService interface {
    GetEnrollmentByStatus(status string) ([]model.Enrollment, error)
    ApproveEnrollment(idUsers []uint)
}

func GetEnrollmentByStatus(status string) ([]*model.Enrollment, error) {
    enrollments, err := repository.FindAllByStatus(status)
    if err != nil {
        log.Error(err)
        return []*model.Enrollment{}, err
    }

    return enrollments, nil
}

func ApproveEnrollment(idUsers []uint) ([]*model.Enrollment, error) {
    enrollments, err := repository.UpdateStatusEnrollment(idUsers)
    if err != nil {
        log.Error(err)
        return []*model.Enrollment{}, err
    }

    return enrollments, nil
}
