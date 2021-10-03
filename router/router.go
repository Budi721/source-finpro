package router

import (
    "fmt"
    "net/http"
	"os"

	"github.com/itp-backend/backend-a-co-create/handler"
	"github.com/itp-backend/backend-a-co-create/service"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func NewRouter(dependencies service.Dependencies) http.Handler {
	r := mux.NewRouter()

	setAuthRouter(r, dependencies.AuthService)
	setCheckRouter(r, dependencies.CheckService)
    setUserRouter(r, dependencies.UserService)
    setEnrollmentRouter(r, dependencies.EnrollmentService)
    setProjectRouter(r, dependencies.ProjectService)
    setArticleRouter(r, dependencies.ArticleService)

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	return loggedRouter
}

func setAuthRouter(router *mux.Router, dependencies service.AuthServiceInterface) {
	router.Methods(http.MethodGet).Path("/auth/token").Handler(handler.GetToken(dependencies))
	router.Methods(http.MethodPost).Path("/auth/token/validate").Handler(handler.ValidateToken(dependencies))
}

func setCheckRouter(router *mux.Router, checkService service.CheckService) {
	router.Methods(http.MethodGet).Path("/check/redis").Handler(handler.CheckRedis(checkService))
	router.Methods(http.MethodGet).Path("/check/mysql").Handler(handler.CheckMysql(checkService))
	router.Methods(http.MethodGet).Path("/check/minio").Handler(handler.CheckMinio(checkService))
}

func setUserRouter(router *mux.Router, dependencies service.IUserService) {
    router.Methods(http.MethodPost).Path("/register").Handler(handler.Register(dependencies))
    router.Methods(http.MethodPost).Path("/login").Handler(handler.Login(dependencies))
    router.Methods(http.MethodPost).Path("/logout").Handler(handler.Logout())

    // Testing Authentication
    router.Methods(http.MethodGet).Path("/").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
        _, _ = fmt.Fprint(writer, "Hello Auth")
    })
}

func setEnrollmentRouter(router *mux.Router, dependencies service.IEnrollmentService)  {
    router.Methods(http.MethodGet).Path("/enrollment_requests").Queries("status", "{status}").Handler(handler.EnrollmentRequestHandler(dependencies))
    // Kurang yang approve
}

func setProjectRouter(router *mux.Router, dependencies service.IProjectService) {
    router.Methods(http.MethodPost).Path("/project/create").Handler(handler.CreateProjectHandler(dependencies))
    router.Methods(http.MethodPost).Path("/project/detail/{id:[0-9]+}").Handler(handler.DetailProjectHandler(dependencies))
    router.Methods(http.MethodPost).Path("/project/delete/{id:[0-9]+}").Handler(handler.DeleteProjectHandler(dependencies))
    // Kurang yang approve
    router.Methods(http.MethodPost).Path("/project").Queries("invited_user_id", "{invited_user_id}").Handler(handler.ProjectByInvitedUserIdHandler(dependencies))
}

func setArticleRouter(router *mux.Router, dependencies service.IArticleService) {
    router.Methods(http.MethodPost).Path("/artikel/create").Handler(handler.CreateArticleHandler(dependencies))
    router.Methods(http.MethodDelete).Path("/artikel/delete/{id:[0-9]+}").Handler(handler.DeleteArticleHandler(dependencies))
    router.Methods(http.MethodGet).Path("/artikel/detail/{id:[0-9]+}").Handler(handler.GetArticleByIdHandler(dependencies))
    router.Methods(http.MethodGet).Path("/list_artikel").Handler(handler.GetAllArticleHandler(dependencies))
}
