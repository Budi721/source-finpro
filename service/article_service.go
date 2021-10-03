package service

import (
    "context"
    "github.com/itp-backend/backend-a-co-create/model"
    "github.com/itp-backend/backend-a-co-create/repository"
    log "github.com/sirupsen/logrus"
    "time"
)

type IArticleService interface {
    CreateArticle(ctx context.Context, article *model.Article) (*model.Article, error)
    DeleteArticle(ctx context.Context, idArticle int) error
    GetArticleById(ctx context.Context, idArticle int) (*model.Article, error)
    GetAllArticle(ctx context.Context) ([]*model.Article, error)
}

func NewArticleService(articleRepository repository.IArticleRepository) IArticleService {
    return &articleService{repo: articleRepository}
}

type articleService struct {
    repo repository.IArticleRepository
}

func (service articleService) CreateArticle(ctx context.Context, article *model.Article) (*model.Article, error) {
    article.PostingDate = time.Now().UnixMilli()
    article, err := service.repo.Create(ctx, article)
    if err != nil {
        log.Error(err)
        return nil, err
    }
    return article, nil
}

func (service articleService) DeleteArticle(ctx context.Context, idArticle int) error {
    err := service.repo.Delete(ctx, idArticle)
    if err != nil {
        log.Error(err)
        return err
    }
    return nil
}

func (service articleService) GetArticleById(ctx context.Context, idArticle int) (*model.Article, error) {
    article, err := service.repo.FindById(ctx, idArticle)
    if err != nil {
        log.Error(err)
        return nil, err
    }
    return article, nil
}

func (service articleService) GetAllArticle(ctx context.Context) ([]*model.Article, error) {
    articles, err := service.repo.FindAll(ctx)
    if err != nil {
        log.Error(err)
        return nil, err
    }
    return articles, nil
}
