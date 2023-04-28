package main

import (
	"github.com/jpdel518/go-graphql-gateway/user/infrastructure/file"
	"github.com/jpdel518/go-graphql-gateway/user/infrastructure/rdb"
	"github.com/jpdel518/go-graphql-gateway/user/infrastructure/rdb/mysql"
	"github.com/jpdel518/go-graphql-gateway/user/presentation/handler"
	"github.com/jpdel518/go-graphql-gateway/user/usecase"
	"github.com/jpdel518/go-graphql-gateway/user/utils"
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
	userRepository := rdb.NewUserRepository(client)
	fileRepository := file.NewUserFileRepository()
	userUsecase := usecase.NewUserUsecase(userRepository, fileRepository, 30*time.Second)
	handler.NewHandler(userUsecase)
}
