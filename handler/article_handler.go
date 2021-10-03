package handler

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/itp-backend/backend-a-co-create/common/errors"
    "github.com/itp-backend/backend-a-co-create/common/responder"
    "github.com/itp-backend/backend-a-co-create/model"
    "github.com/itp-backend/backend-a-co-create/service"
    log "github.com/sirupsen/logrus"
    "net/http"
    "strconv"
)

func CreateArticleHandler(service service.IArticleService) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            responder.NewHttpResponse(r, w, http.StatusMethodNotAllowed, nil, errors.New("Error: Method is not allowed"))
            return
        }

        var articleRequest *model.Article
        if err := json.NewDecoder(r.Body).Decode(&articleRequest); err != nil {
            log.Error(err)
            responder.NewHttpResponse(r, w, http.StatusInternalServerError, nil, err)
            return
        }

        article, err := service.CreateArticle(r.Context(), articleRequest)
        if err != nil {
            log.Error(err)
            responder.NewHttpResponse(r, w, http.StatusInternalServerError, nil, err)
            return
        }

        responder.NewHttpResponse(r, w, http.StatusCreated, article, nil)
        return
    }
}

func DeleteArticleHandler(service service.IArticleService) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodDelete {
            log.Println("MASUK KE SALAH METOD")
            responder.NewHttpResponse(r, w, http.StatusMethodNotAllowed, "method not allowed", errors.New("Error: Method is not allowed"))
            return
        }

        vars := mux.Vars(r)
        id, err := strconv.Atoi(vars["id"])
        if err != nil {
            log.Error(err)
            responder.NewHttpResponse(r, w, http.StatusInternalServerError, nil, err)
            return
        }

        err = service.DeleteArticle(r.Context(), id)
        if err != nil {
            log.Error(err)
            responder.NewHttpResponse(r, w, http.StatusInternalServerError, nil, err)
            return
        }

        responder.NewHttpResponse(r, w, http.StatusNoContent, "", nil)
        return
    }
}

func GetArticleByIdHandler(service service.IArticleService) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodGet {
            responder.NewHttpResponse(r, w, http.StatusMethodNotAllowed, nil, errors.New("Error: Method is not allowed"))
            return
        }

        vars := mux.Vars(r)
        id, err := strconv.Atoi(vars["id"])
        if err != nil {
            log.Error(err)
            responder.NewHttpResponse(r, w, http.StatusInternalServerError, nil, err)
            return
        }

        article, err := service.GetArticleById(r.Context(), id)
        if err != nil {
            log.Error(err)
            responder.NewHttpResponse(r, w, http.StatusInternalServerError, nil, err)
            return
        }

        responder.NewHttpResponse(r, w, http.StatusOK, article, nil)
        return
    }
}

func GetAllArticleHandler(service service.IArticleService) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodGet {
            responder.NewHttpResponse(r, w, http.StatusMethodNotAllowed, nil, errors.New("Error: Method is not allowed"))
            return
        }

        articles, err := service.GetAllArticle(r.Context())
        if err != nil {
            log.Error(err)
            responder.NewHttpResponse(r, w, http.StatusInternalServerError, nil, err)
            return
        }

        responder.NewHttpResponse(r, w, http.StatusOK, articles, nil)
        return
    }
}

