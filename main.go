package main

import (
	"fmt"
	"net/http"

	"github.com/DataInsightHub/Event-Service/generalrouter"
	"github.com/DataInsightHub/Event-Service/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			err := fmt.Errorf("%v", r)
			logrus.Errorln(err.Error())
		}
	}()

	err := godotenv.Load()
	if err != nil {
		logrus.Errorln("Could not load env")
		logrus.Errorln(err.Error())
		return 
	}

	logger, err := logger.BuildLogger()
	if err != nil {
		logrus.Errorln(err)
		return 
	}

	accessInfo, err := generalrouter.LoadAccessInfo()
	if err != nil {
		logger.Error(err.Error())
		return 
	}
	
	router := generalrouter.CreateDefaultRouter()
	routerPubGroup, routerSubGroup := generalrouter.CreateRouterGroups(router, accessInfo)

	// Middleware für Route /foo mit Basic Auth für Test1:pass1
	routerPubGroup.GET("/foo", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Authorized access to /foo"})
	})

	// Middleware für Route /bar mit Basic Auth für Test2:pass2
	routerSubGroup.GET("/bar", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Authorized access to /bar"})
	})

	logger.Info("Event service started at port 2345")
	router.Run("localhost:2345")
}