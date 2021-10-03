package handler

import (
	"encoding/json"
	"github.com/itp-backend/backend-a-co-create/common/errors"
	"github.com/itp-backend/backend-a-co-create/common/responder"
	"github.com/itp-backend/backend-a-co-create/model"
	"github.com/itp-backend/backend-a-co-create/service"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func Login(IUserService service.IUserService) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            responder.NewHttpResponse(r, w, http.StatusMethodNotAllowed, nil, errors.New("Error: Method is not allowed"))
            return
        }
        var user model.User
        if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
            log.Error(err)
            responder.NewHttpResponse(r, w, http.StatusInternalServerError, nil, err)
            return
        }

        u, err := IUserService.Login(r.Context(), user.Username, user.Password, user.LoginAs)
        if err != nil {
            responder.NewHttpResponse(r, w, http.StatusUnauthorized, nil, errors.NewUnauthorizedError("invalid username and password"))
            return
        }

        cookie := http.Cookie{
            Name:       "jwt",
            Value:      u.AuthToken,
            Expires:    time.Now().Add(time.Hour * 24),
        }

        http.SetCookie(w, &cookie)

        responder.NewHttpResponse(r, w, http.StatusOK, u, nil)
        return
    }
}

func Register(IUserService service.IUserService) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            responder.NewHttpResponse(r, w, http.StatusMethodNotAllowed, nil, errors.New("Error: Method is not allowed"))
            return
        }

        var enrollment model.Enrollment

        if err := json.NewDecoder(r.Body).Decode(&enrollment); err != nil {
            log.Warning(err)
            responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
            return
        }

        u, err := IUserService.Register(r.Context(), enrollment.NamaLengkap, enrollment.Username, enrollment.Password, enrollment.TopikDiminati)
        if err != nil {
            log.Warning(err)
            responder.NewHttpResponse(r, w, http.StatusInternalServerError, nil, err)
            return
        }

        w.WriteHeader(http.StatusCreated)

        responder.NewHttpResponse(r, w, http.StatusCreated, u, nil)
        return
    }
}

func Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie := http.Cookie{
			Name:    "jwt",
			Value:   "",
			Expires: time.Now().Add(-time.Hour),
		}

		http.SetCookie(w, &cookie)

        w.WriteHeader(http.StatusOK)

        responder.NewHttpResponse(r, w, http.StatusOK, "", nil)
        return
	}
}
