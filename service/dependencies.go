package service

import (
	"github.com/itp-backend/backend-a-co-create/app"
	"github.com/itp-backend/backend-a-co-create/external/jwt_client"
	"github.com/itp-backend/backend-a-co-create/external/minio"
	"github.com/itp-backend/backend-a-co-create/external/mysql"
	"github.com/itp-backend/backend-a-co-create/external/redis"
	"github.com/itp-backend/backend-a-co-create/repository"
	"gorm.io/gorm"
)

type Dependencies struct {
	DBConn            *gorm.DB
	AuthService       AuthServiceInterface
	CheckService      CheckService
	UserService       IUserService
	EnrollmentService IEnrollmentService
	ProjectService    IProjectService
	ArticleService    IArticleService
}

func InstantiateDependencies(application *app.Application) Dependencies {
	jwtWrapper := jwt_client.New()
	authService := NewAuthService(application.Config, jwtWrapper)
	redisClient := redis.NewRedisClient(application.Config.RedisAddress)
	mysqlClient := mysql.NewMysqlClient(mysql.ClientConfig{
		Username: application.Config.DBUsername,
		Password: application.Config.DBPassword,
		Host:     application.Config.DBHost,
		Port:     application.Config.DBPort,
		DBName:   application.Config.DBName,
	})
	minioClient := minio.NewMinioClient(minio.ClientConfig{
		Endpoint:   application.Config.MinioEndpoint,
		AccessKey:  application.Config.MinioAccessKey,
		SecretKey:  application.Config.MinioSecretKey,
		Region:     application.Config.MinioRegion,
		BucketName: application.Config.MinioBucket,
	})
	checkService := NewCheckService(redisClient, mysqlClient, minioClient)

	dbConn := mysql.NewDBConnection(mysqlClient)

	userRepo := repository.NewUserRepository(dbConn)
	userService := NewUserService(userRepo, application.Config, jwtWrapper)

	enrollmentRepo := repository.NewEnrollmentRepository(dbConn)
	enrollmentService := NewEnrollmentService(enrollmentRepo)

	projectRepo := repository.NewProjectRepository(dbConn)
	projectService := NewProjectService(projectRepo)

	articleRepo := repository.NewArticleRepository(dbConn)
	articleService := NewArticleService(articleRepo)

	return Dependencies{
		AuthService:       authService,
		CheckService:      checkService,
		UserService:       userService,
		EnrollmentService: enrollmentService,
		ProjectService:    projectService,
		ArticleService:    articleService,
	}
}
