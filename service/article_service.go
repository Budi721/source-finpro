package service

import (
	"github.com/itp-backend/backend-a-co-create/dto"
	"github.com/itp-backend/backend-a-co-create/model"
	"github.com/itp-backend/backend-a-co-create/repository"
	log "github.com/sirupsen/logrus"
)

type IArticleService interface {
	CreateArticle(article *dto.Article) (*model.Article, error)
	DeleteArticle(idArticle int) error
	GetArticleById(idArticle int) (*model.Article, error)
	GetAllArticle() ([]*model.Article, error)
	GetArticleByIdUser(idUser uint64) (*model.Article, error)
}

func CreateArticle(article *dto.Article) (*model.Article, error) {
	articleToCreate, err := repository.CreateArticle(article)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return articleToCreate, nil
}

func DeleteArticle(idArticle int) error {
	err := repository.DeleteArticle(idArticle)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func GetArticleById(idArticle int) (*model.Article, error) {
	article, err := repository.FindArticleById(idArticle)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return article, nil
}

func GetAllArticle() ([]*model.Article, error) {
	articles, err := repository.FindAllArticle()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return articles, nil
}

func GetArticleByIdUser(idUser uint64) ([]*model.Article, error) {
	articles, err := repository.FindArticleByIdUser(idUser)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return articles, nil
}