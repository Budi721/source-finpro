package repository

import (
    "context"
    "github.com/itp-backend/backend-a-co-create/common/errors"
    "github.com/itp-backend/backend-a-co-create/model"
    log "github.com/sirupsen/logrus"
    "gorm.io/gorm"
)

type IArticleRepository interface {
    Create(ctx context.Context, article *model.Article) (*model.Article, error)
    Delete(ctx context.Context, idArticle int) error
    FindById(ctx context.Context, idArticle int) (*model.Article, error)
    FindAll(ctx context.Context) ([]*model.Article, error)
}

func NewArticleRepository(db *gorm.DB) IArticleRepository {
    return articleRepository{DB: db}
}

type articleRepository struct {
    DB *gorm.DB
}

func (repo articleRepository) Create(ctx context.Context, article *model.Article) (*model.Article, error) {
    a := &model.Article{
        PostingDate: article.PostingDate,
        Kategori:    article.Kategori,
        Judul:       article.Judul,
        IsiArtikel:  article.IsiArtikel,
        IdUser:      article.IdUser,
    }

    result := repo.DB.Create(&a)
    if result.Error != nil {
        log.Error(result.Error)
        return nil, result.Error
    }
    return a, nil
}

func (repo articleRepository) Delete(ctx context.Context, idArticle int) error {
    var article model.Article
    article.IdArtikel = idArticle

    err := repo.DB.Where("id_artikel = ?", idArticle).First(&article).Error

    switch err {
    case nil:
        repo.DB.Delete(&article)
        return nil
    case gorm.ErrRecordNotFound:
        return errors.NewInternalError(err, "Error: not found")
    default:
        return errors.NewInternalError(err, "Error: database error")
    }
}

func (repo articleRepository) FindById(ctx context.Context, idArticle int) (*model.Article, error) {
    var article model.Article
    article.IdArtikel = idArticle

    if err := repo.DB.First(&article).Error; err != nil {
        log.Error(err)
        return &article, err
    }

    return &article, nil
}

func (repo articleRepository) FindAll(ctx context.Context) ([]*model.Article, error) {
    var articles []*model.Article
    if err := repo.DB.Table("articles").Find(&articles).Error; err != nil {
        log.Error(err)
        return articles, err
    }

    return articles, nil
}

