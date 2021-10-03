package repository

import (
    "context"
    "github.com/itp-backend/backend-a-co-create/model"
    log "github.com/sirupsen/logrus"
    "gorm.io/gorm"
)

type IEnrollmentRepository interface {
    FindAllByStatus(ctx context.Context, status string) ([]*model.Enrollment, error)
    // Catatan flag apa yang digunakan untuk menandai sudah diapprove
}

func NewEnrollmentRepository(db *gorm.DB) IEnrollmentRepository {
    return enrollmentRepository{DB: db}
}

type enrollmentRepository struct {
    DB *gorm.DB
}

func (repo enrollmentRepository) FindAllByStatus(ctx context.Context, status string) ([]*model.Enrollment, error) {
    var enrollments []*model.Enrollment
    if err := repo.DB.Table("enrollments").Find(&enrollments).Error; err != nil {
        log.Error(err)
        return enrollments, err
    }

    return enrollments, nil
}


