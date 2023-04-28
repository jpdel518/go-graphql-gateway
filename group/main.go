package main

import (
	"github.com/jpdel518/go-graphql-gateway/group/infrastructure/rdb"
	"github.com/jpdel518/go-graphql-gateway/group/infrastructure/rdb/mysql"
	"github.com/jpdel518/go-graphql-gateway/group/presentation/handler"
	"github.com/jpdel518/go-graphql-gateway/group/usecase"
	"github.com/jpdel518/go-graphql-gateway/group/utils"
	"os"
	"time"
)

func init() {
	utils.LoggingSettings(os.Getenv("LOG_FILE"))
	mysql.InitDatabase()
}

func main() {
	// Dependency Injection
	client := mysql.NewClient()
	siteRepository := rdb.NewSiteRepository(client)
	siteUsecase := usecase.NewGroupUsecase(siteRepository, 30*time.Second)
	handler.NewHandler(siteUsecase)
}
