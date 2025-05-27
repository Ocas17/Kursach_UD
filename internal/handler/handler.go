package handler

import (
	"github.com/Ocas17/Kursach_UD/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// Add CORS middleware
	router.Use(h.corsMiddleware)
	router.Static("/front", "./front")

	api := router.Group("/api")
	{
		clients := api.Group("clients")
		{
			clients.POST("/", h.createClient)      // создание клиента
			clients.GET("/", h.getAllClients)      // получение всех клиентов
			clients.GET("/:id", h.getClientById)   // получение клиента по id
			clients.PUT("/:id", h.updateClient)    // обновление клиента
			clients.DELETE("/:id", h.deleteClient) // удаление клиента

			policies := clients.Group("/:id/policies")
			{
				policies.POST("/", h.createPolicy)  // создание полиса для клиента
				policies.GET("/", h.getAllPolicies) // получение всех полисов клиента
			}
		}

		policies := api.Group("policies")
		{
			policies.GET("/:id", h.getPolicyById)   // получение полиса по id
			policies.PUT("/:id", h.updatePolicy)    // обновление полиса
			policies.DELETE("/:id", h.deletePolicy) // удаление полиса

			claims := policies.Group("/:id/claims")
			{
				claims.POST("/", h.createClaim) // создание страхового случая
				claims.GET("/", h.getAllClaims) // получение всех страховых случаев
			}
		}

		claims := api.Group("claims")
		{
			claims.GET("/:id", h.getClaimById)   // получение страхового случая по id
			claims.PUT("/:id", h.updateClaim)    // обновление страхового случая
			claims.DELETE("/:id", h.deleteClaim) // удаление страхового случая
		}
	}

	return router
}

// CORS middleware
func (h *Handler) corsMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
		return
	}

	c.Next()
}
