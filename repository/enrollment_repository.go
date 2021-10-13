package repository

import (
    "github.com/itp-backend/backend-a-co-create/model"
    log "github.com/sirupsen/logrus"
)

type nrollmentRepository interface {
    FindAllByStatus(status string) ([]*model.Enrollment, error)
    UpdateStatusEnrollment(idUsers []uint) ([]*model.Enrollment, error)
}

func FindAllByStatus(status string) ([]*model.Enrollment, error) {
    var enrollments []*model.Enrollment
    if err := db.Where("enrollment_status = ?", status).Find(&enrollments).Error; err != nil {
        log.Error(err)
        return enrollments, err
    }

    return enrollments, nil
}

func UpdateStatusEnrollment(idUsers []uint) ([]*model.Enrollment, error) {
    var enrollments []*model.Enrollment
    db.Where("id_user IN ?", idUsers).Find(&enrollments)

    if err := db.Table("enrollments").
       Where("id_user IN ?", idUsers).
       Updates(model.Enrollment{EnrollmentStatus: 1}).
       Error; err != nil {
           log.Error(err)
           return enrollments, err
    }

    return enrollments, nil
}

