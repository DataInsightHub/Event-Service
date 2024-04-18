package generalrouter

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CreateDefaultRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST, OPTIONS, GET, PUT", "DELETE"},
		AllowHeaders:     []string{"Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"},
		AllowCredentials: true,
	}))

	router.SetTrustedProxies(nil)

	return router
}

func CreateRouterGroups(router *gin.Engine, accessInfo AccessInfo) (*gin.RouterGroup, *gin.RouterGroup) {
	authorizedRouterPub := router.Group("/pub", gin.BasicAuth(gin.Accounts{
		accessInfo.PublisherUsername: accessInfo.PablisherPassword,
	}))

	authorizedRouterSub := router.Group("/sub", gin.BasicAuth(gin.Accounts{
		accessInfo.SubscriberUsername: accessInfo.SubscriberPassword,
	}))

	return authorizedRouterPub, authorizedRouterSub
}
