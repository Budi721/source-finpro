package repository

import (
	"errors"
	"github.com/itp-backend/backend-a-co-create/dto"
	"github.com/itp-backend/backend-a-co-create/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type IArticleRepository interface {
	CreateArticle(article *dto.Article) (*model.Article, error)
	DeleteArticle(idArticle int) error
	FindArticleById(idArticle int) (*model.Article, error)
	FindAllArticle() ([]*model.Article, error)
}

type articleRepository struct {
	DB *gorm.DB
}

func CreateArticle(article *dto.Article) (*model.Article, error) {
	a := &model.Article{
		PostingDate: time.Now(),
		Kategori:    article.Kategori,
		Judul:       article.Judul,
		IsiArtikel:  article.IsiArtikel,
		IdUser:      article.IdUser,
	}

	result := db.Create(&a)
	if result.Error != nil {
		log.Error(result.Error)
		return nil, result.Error
	}
	return a, nil
}

func DeleteArticle(idArticle int) error {
	var article model.Article
	article.IdArtikel = idArticle

	err := db.Where("id_artikel = ?", idArticle).First(&article).Error

	switch err {
	case nil:
		db.Delete(&article)
		return nil
	case gorm.ErrRecordNotFound:
		return errors.New("error: not found")
	default:
		return errors.New("error: database error")
	}
}

func FindArticleById(idArticle int) (*model.Article, error) {
	var article model.Article
	article.IdArtikel = idArticle

	if err := db.First(&article).Error; err != nil {
		log.Error(err)
		return &article, err
	}

	return &article, nil
}

func FindAllArticle() ([]*model.Article, error) {
	var articles []*model.Article
	if err := db.Table("articles").Find(&articles).Error; err != nil {
		log.Error(err)
		return articles, err
	}

	return articles, nil
}


