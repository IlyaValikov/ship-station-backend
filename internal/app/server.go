package app

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Run запускает приложение.
func (app *Application) Run() {
    r := gin.Default()
    r.Use(cors.Default())  

    ShipGroup := r.Group("/ship")
    {   
        ShipGroup.GET("/", app.Handler.GetShips)
        ShipGroup.GET("/:shipID", app.Handler.GetShipByID) 
        ShipGroup.DELETE("/:shipID", app.Handler.DeleteShip) 
        ShipGroup.POST("/", app.Handler.CreateShip)
        ShipGroup.PUT("/:shipID", app.Handler.UpdateShip) 
        ShipGroup.POST("/:shipID/request", app.Handler.AddShipToRequest) 
        ShipGroup.DELETE("/:shipID/request", app.Handler.RemoveShipFromRequest)
        ShipGroup.POST("/:shipID/image",app.Handler.AddShipImage)
    }
    
    // Группа запросов для заявки
    RequestGroup := r.Group("/request")
    {
        RequestGroup.GET("/", app.Handler.GetRequests)
        RequestGroup.GET("/:requestID", app.Handler.GetRequestByID)
        RequestGroup.DELETE("/:requestID", app.Handler.DeleteRequest)
        RequestGroup.PUT("/:requestID/status/user", app.Handler.UpdateRequestStatusUser) 
        RequestGroup.PUT("/:requestID/status/moderator", app.Handler.UpdateRequestStatusModerator)  
    }

    addr := fmt.Sprintf("%s:%d", app.Config.ServiceHost, app.Config.ServicePort)
    r.Run(addr)
    log.Println("Server down")
}