package web

import (
	notesController "notesTask/controllers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/api/v1/") // Create a new router group with the /api/v1 prefix

	apiv1.GET("/", notesController.Home)
	apiv1.POST("/note", notesController.CreateNote)
	apiv1.GET("/notes", notesController.GetNotes)
	apiv1.PUT("/note/:id", notesController.UpdateNote)
	apiv1.DELETE("/notes/", notesController.DeleteNotes)

	apiv1.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{"PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:8080"
		},
		MaxAge: 12 * time.Hour,
	}))

	return r
}
