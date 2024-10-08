package handler

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"todo-app/pkg/service"

	_ "todo-app/docs"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

const ItemIdPath = "/:item_id"
const IdPath = "/:id"

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		Item := api.Group("/lists")
		{
			Item.POST("/", h.createList)
			Item.GET("/", h.getAllLists)
			Item.GET(IdPath, h.getListById)
			Item.PUT(IdPath, h.updateList)
			Item.DELETE(IdPath, h.deleteList)

			items := Item.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
				items.GET(ItemIdPath, h.getItemById)
				items.PUT(ItemIdPath, h.updateItem)
				items.DELETE(ItemIdPath, h.deleteItem)
			}
		}
	}

	return router
}
