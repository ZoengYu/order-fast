package api

import (
	db "github.com/ZoengYu/order-fast-project/db/sqlc"
	util "github.com/ZoengYu/order-fast-project/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	db_service	db.DBService
	router		*gin.Engine
	config		util.Config
}

func NewServer(config util.Config, db_service db.DBService) (*Server, error) {
	server := &Server{
		db_service:	db_service,
		config:		config,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	v1 := router.Group("/v1")
	// allow all origins request
	router.Use(cors.Default())

	v1.POST("/store", server.createStore)
	v1.GET("/store/:id", server.getStore)
	v1.GET("/store/name", server.getStoreByName)
	v1.PUT("/store", server.updateStore)
	v1.DELETE("/store/:id", server.delStore)

	v1.POST("/store/menu", server.createStoreMenu)
	v1.GET("/store/menu", server.getStoreMenu)
	v1.PUT("/store/menu", server.updateStoreMenu)
	v1.DELETE("/store/menu/:id", server.deleteStoreMenu)
	v1.GET("/store/menu_list", server.listStoreMenu)

	v1.POST("store/menu/food", server.createMenuFood)
	v1.DELETE("store/menu/food/:id", server.deleteMenuFood)
	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
