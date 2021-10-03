package handler

import (
    "github.com/itp-backend/backend-a-co-create/common/errors"
    "github.com/itp-backend/backend-a-co-create/common/responder"
    "github.com/itp-backend/backend-a-co-create/service"
    log "github.com/sirupsen/logrus"
    "net/http"
)

func EnrollmentRequestHandler(service service.IEnrollmentService) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodGet {
            responder.NewHttpResponse(r, w, http.StatusMethodNotAllowed, nil, errors.New("Error: Method is not allowed"))
            return
        }

        status := r.FormValue("status")
        enrollments, err := service.GetByStatus(r.Context(), status)
        if err != nil {
            log.Error(err)
            responder.NewHttpResponse(r, w, http.StatusInternalServerError, nil, err)
            return
        }

        responder.NewHttpResponse(r, w, http.StatusOK, enrollments, nil)
        return
    }
}
