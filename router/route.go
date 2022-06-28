package router

import (
	"fehu/conf"
	"fehu/controller"
	"fehu/docs"
	"fehu/middleware"

	_ "fehu/docs"

	"github.com/gin-contrib/pprof"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationurl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationurl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

// @x-extension-openapi {"example": "value on a json format"}

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {

	router := gin.New()
	router.ForwardedByClientIP = true

	// 错误中间件
	router.Use(middleware.RecoveryMiddleware())

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))

	if conf.Swag.Enable {
		logrus.Info("加载gin swagger")
		docs.SwaggerInfo.Version = conf.Swag.Version
		docs.SwaggerInfo.BasePath = conf.Swag.BasePath
		docs.SwaggerInfo.Host = conf.Swag.Host
		docs.SwaggerInfo.Schemes = conf.Swag.Schemes
		docs.SwaggerInfo.Title = conf.Swag.Title
		docs.SwaggerInfo.Description = conf.Swag.Description
		url := ginSwagger.URL(conf.Swag.Url)
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}

	router.Use(middleware.LogMiddleware(), middleware.HeaderAuthMiddleware())

	pprof.Register(router)

	r := router.Group("")
	r2 := router.Group("", middleware.JwtAuthMiddleware())
	{
		controller.OrderRegister(r, r2)
	}

	return router
}
