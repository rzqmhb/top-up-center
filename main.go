package main

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rzqmhb/top-up-center/database"
	"github.com/rzqmhb/top-up-center/handler/api"
	"github.com/rzqmhb/top-up-center/middleware"
	"github.com/rzqmhb/top-up-center/repository"
	"github.com/rzqmhb/top-up-center/service"
)

func init()  {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		
		router := gin.Default()

		postgresDB, err := database.InitDB()
		if err != nil {
			panic(err)
		}

		router = RunServer(router, postgresDB)

		fmt.Println("Server is running on port 9000")
		err = router.Run(":9000")
		if err != nil {
			panic(err)
		}
	}()

	wg.Wait()
}

func RunServer(ginRouter *gin.Engine, postgresDB *database.PostgresDB) *gin.Engine {
	var (
		userRepo = repository.NewUserRepository(postgresDB)
		sessionRepo = repository.NewSessionRepository(postgresDB)
		itemRepository = repository.NewItemRepository(postgresDB)
		orderRepository = repository.NewOrderRepository(postgresDB)
	)

	var (
		userService = service.NewUserService(userRepo, sessionRepo)
		itemService = service.NewItemService(itemRepository)
		orderService = service.NewOrderService(orderRepository)
	)

	var (
		userAPI = api.NewUserAPI(userService)
		itemAPI = api.NewItemAPI(itemService)
		orderAPI = api.NewOrderAPI(orderService, userService)
	)

	api := ginRouter.Group("/api") 
	{
		api.POST("/user/login", userAPI.Login)
		api.POST("/user/register", userAPI.Register)

		api.Use(middleware.Auth())
		api.GET("/home", itemAPI.GetAll)
		api.GET("/items/:gameId", itemAPI.GetByGame)
		api.GET("/search", itemAPI.GetByKeywords)
		api.GET("/user/orders", orderAPI.GetUserOrder)
		api.GET("/user/logout", userAPI.Logout)
	}
	return ginRouter
}
