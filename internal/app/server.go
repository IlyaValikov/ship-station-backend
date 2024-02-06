package app

import (
	"fmt"
	"log"

	"backend/docs"
	"backend/internal/pkg/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func (app *Application) Run() {
    r := gin.Default()  

	docs.SwaggerInfo.Title = "ShipStation RestAPI"
	docs.SwaggerInfo.Description = "API server for ShipStation application"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    ShipGroup := r.Group("/ship")
    {   
        ShipGroup.GET("/", middleware.Guest(app.Repository.GetRedisClient(), []byte("AccessSecretKey"), app.Repository), app.Handler.GetShips)
        ShipGroup.GET("/:shipID", middleware.Guest(app.Repository.GetRedisClient(), []byte("AccessSecretKey"), app.Repository), app.Handler.GetShipByID)
        ShipGroup.DELETE("/:shipID", middleware.Authenticate(app.Repository.GetRedisClient(), []byte("AccessSecretKey"), app.Repository), app.Handler.DeleteShip)
        ShipGroup.POST("/", middleware.Authenticate(app.Repository.GetRedisClient(), []byte("AccessSecretKey"), app.Repository), app.Handler.CreateShip)
        ShipGroup.PUT("/:shipID", middleware.Authenticate(app.Repository.GetRedisClient(), []byte("AccessSecretKey"), app.Repository), app.Handler.UpdateShip)
        ShipGroup.POST("/:shipID/request", middleware.Authenticate(app.Repository.GetRedisClient(), []byte("AccessSecretKey"), app.Repository), app.Handler.AddShipToRequest)
        ShipGroup.DELETE("/:shipID/request", middleware.Authenticate(app.Repository.GetRedisClient(), []byte("AccessSecretKey"), app.Repository), app.Handler.RemoveShipFromRequest)
        ShipGroup.POST("/:shipID/image", middleware.Authenticate(app.Repository.GetRedisClient(), []byte("AccessSecretKey"), app.Repository), app.Handler.AddShipImage)
    }
    
    RequestGroup := r.Group("/request").Use(middleware.Authenticate(app.Repository.GetRedisClient(), []byte("AccessSecretKey"), app.Repository))
    {
        RequestGroup.GET("/", app.Handler.GetRequests)
        RequestGroup.GET("/:requestID", app.Handler.GetRequestByID)
        RequestGroup.DELETE("/:requestID", app.Handler.DeleteRequest)
        RequestGroup.PUT("/:requestID/status/user", app.Handler.UpdateRequestStatusUser)  
        RequestGroup.PUT("/:requestID/status/moderator", app.Handler.UpdateRequestStatusModerator)  
        RequestGroup.PUT("/check", app.Handler.CheckRequestUser)  
    }

    UserGroup := r.Group("/user")
    {
        UserGroup.POST("/register", app.Handler.Register)
        UserGroup.POST("/login", app.Handler.Login)
        UserGroup.POST("/logout", middleware.Authenticate(app.Repository.GetRedisClient(), []byte("AccessSecretKey"), app.Repository), app.Handler.Logout)
    }
    addr := fmt.Sprintf("%s:%d", app.Config.ServiceHost, app.Config.ServicePort)
    r.Run(addr)
    log.Println("Server down")
}