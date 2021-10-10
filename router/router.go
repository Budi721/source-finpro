package router

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/itp-backend/backend-a-co-create/config"
	"github.com/itp-backend/backend-a-co-create/controller"
	"github.com/itp-backend/backend-a-co-create/middleware"
)

func AllRouters() *gin.Engine {
	r := gin.Default()

	r.GET("/", controller.TestRouter)

	apiRoutes := r.Group("/api/v1")
	{
		authRouter := apiRoutes.Group("/auth")
		{
			authRouter.POST("/login", controller.Login)
			authRouter.POST("/register", controller.Register)
			authRouter.DELETE("/logout", controller.TestRouter)
		}

		adminRouter := apiRoutes.Group("/admin", middleware.AuthorizeJWT())
		{
			adminRouter.GET("/all-users", controller.GetAllUser)
		}

		userRouter := apiRoutes.Group("/user", middleware.AuthorizeJWT())
		{
			userRouter.GET("/profile", controller.TestRouter)
			userRouter.PUT("/update", controller.UpdateUser)
			userRouter.PUT("/change-password", controller.ChangePassword)
		}

		// with middleware jwt
		enrollRouter := apiRoutes.Group("/enroll")
		{
			enrollRouter.GET("/requests", controller.TestRouter)
			enrollRouter.POST("/approve", controller.TestRouter)
		}

		// with middleware jwt
		acceptEnrollRouter := apiRoutes.Group("/accept-enroll")
		{
			acceptEnrollRouter.POST("/", controller.TestRouter)
			acceptEnrollRouter.POST("/approve", controller.TestRouter)
		}

		// with middleware jwt
		projectRouter := apiRoutes.Group("/project")
		{
			projectRouter.GET("/", controller.TestRouter)
			projectRouter.POST("/create", controller.TestRouter)
			projectRouter.GET("/detail/:id", controller.TestRouter)
			projectRouter.DELETE("/delete/:id", controller.TestRouter)
		}

		articleRouter := apiRoutes.Group("/artikel")
		{
			articleRouter.GET("/list", controller.TestRouter)
			articleRouter.POST("/create", controller.TestRouter)
			articleRouter.GET("/detail/:id", controller.TestRouter)
			articleRouter.DELETE("/delete/:id", controller.TestRouter)
		}
	}

	return r
}

func RunRouter() {
	port := config.Init().AppPort
	if port == "" {
		port = "8080"
	}

	mode := config.Init().Environment
	gin.SetMode(mode)
	r := AllRouters()

	log.Println("Starting server at", port)
	log.Println("Quit the server with CTRL-C.")
	log.Fatal(r.Run(":" + port))
}
