package handler

import (
    "encoding/json"
    "github.com/itp-backend/backend-a-co-create/common/errors"
    "github.com/itp-backend/backend-a-co-create/common/responder"
    "github.com/itp-backend/backend-a-co-create/model"
    "github.com/itp-backend/backend-a-co-create/service"
    log "github.com/sirupsen/logrus"
    "net/http"
)


func CreateProjectHandler(service service.IProjectService) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            responder.NewHttpResponse(r, w, http.StatusMethodNotAllowed, nil, errors.New("Error: Method is not allowed"))
            return
        }

        var project model.Project
        if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
            log.Error(err)
            responder.NewHttpResponse(r, w, http.StatusInternalServerError, nil, err)
            return
        }
        /**
        {
        "kategori_project":"Go Green",
        "nama_project":"Reboisasi",
        "tanggal_mulai":"09/09/2021",
        "link_trello":"https://trello.com/c/EuJ1zjbJ/94-devlist-backend",
        "deskripsi_project":"Deskripsi",
        "invited_user_id":[1,2]
        }
         */
        retrievedProject, err := service.CreateProject(r.Context(), &project)
        if err != nil {
            log.Error(err)
            responder.NewHttpResponse(r, w, http.StatusInternalServerError, nil, err)
            return
        }

        responder.NewHttpResponse(r, w, http.StatusCreated, retrievedProject, nil)
        return
    }
}

func DetailProjectHandler(service service.IProjectService) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {

    }
}

func DeleteProjectHandler(service service.IProjectService) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {

    }
}

func ProjectByInvitedUserIdHandler(service service.IProjectService) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {

    }
}
