package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/trongdth/mroom_backend/api"
	"github.com/trongdth/mroom_backend/config"
	"github.com/trongdth/mroom_backend/daos"
	"github.com/trongdth/mroom_backend/services"
	"go.uber.org/zap"
)

func main() {
	// load config
	conf := config.GetConfig()

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("failed to create zap logger: %v", err)
	}
	defer logger.Sync()

	// init daos
	if err := daos.Init(conf); err != nil {
		panic(err)
	}

	if err := daos.AutoMigrate(); err != nil {
		logger.Fatal("failed to auto migrate", zap.Error(err))
	}

	var (
		userDAO = daos.NewUser()
		userSvc = services.NewUserService(userDAO, conf)
	)

	r := gin.New()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://*", "https://*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		MaxAge:           12 * time.Hour,
	}))

	svr := api.NewServer(r, userSvc, conf)
	authMw := api.AuthMiddleware(string(conf.TokenSecretKey), svr.Authenticate)
	svr.Routes(authMw)

	if err := r.Run(fmt.Sprintf(":%d", conf.Port)); err != nil {
		logger.Fatal("router.Run", zap.Error(err))
	}
}
