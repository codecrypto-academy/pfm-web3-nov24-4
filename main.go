package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/codecrypto-academy/pfm-web3-nov24-4/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
}

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "Puerto del servidor")
	flag.Parse()

	if flag.Arg(0) == "help" {
		flag.Usage()
		return
	}

	r := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}

	r.Use(cors.New(corsConfig))

	// Middleware de autenticación
	authMiddleware := func(c *gin.Context) {
		username, password, ok := c.Request.BasicAuth()
		if !ok || username != "admin" || password != "password" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}

	v1 := r.Group("/api/v1", authMiddleware)
	{
		v1.POST("/create-fabric-ca", handlers.CreateFabricCA)
		v1.POST("/register-fabric-ca", handlers.RegisterFabricCA)
		v1.POST("/create-fabric-peer", handlers.CreateFabricPeer)
		v1.POST("/create-fabric-orderer", handlers.CreateFabricOrderer)
		v1.POST("/enroll-fabric-user", handlers.EnrollFabricUser)
		v1.POST("/attach-fabric-user-connchain", handlers.AttachUserToConnection)
		v1.POST("/create-fabric-wallet", handlers.CreateFabricWallet)

		v1.GET("/check-fabric-ca-status", handlers.CheckFabricCAStatus)
		v1.GET("/check-fabric-peer-status", handlers.CheckFabricPeerStatus)
		v1.GET("/check-fabric-orderer-status", handlers.CheckFabricOrdererStatus)
		v1.GET("/get-fabric-connection-chain", handlers.GetFabricConnectionChain)

		v1.DELETE("/delete-fabric-orderer-org", handlers.DeleteOrdererOrganization)
		v1.DELETE("/delete-fabric-peer-org", handlers.DeletePeerOrganization)

	}

	r.Run(fmt.Sprintf(":%d", port))
}
