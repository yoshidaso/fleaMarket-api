package main

import (
	"gin-freemarket/controllers"
	"gin-freemarket/infra"
	//"gin-freemarket/models"
	"gin-freemarket/repositories"
	"gin-freemarket/services"
	"github.com/gin-gonic/gin"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()
	//items := []models.Item{
	//	{ID: 1, Name: "商品1", Price: 1000, Description: "説明1", SoldOut: false},
	//	{ID: 2, Name: "商品2", Price: 2000, Description: "説明2", SoldOut: true},
	//	{ID: 3, Name: "商品3", Price: 3000, Description: "説明3", SoldOut: false},
	//}

	// 各インスタンスの作成
	//itemRepository := repositories.NewItemMemoryRepository(items)
	itemRepository := repositories.NewItemRepository(db)
	itemService := services.NewItemServices(itemRepository)
	itemController := controllers.NewItemController(itemService)

	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)

	r := gin.Default()

	// ルーティングの設定
	itemRouter := r.Group("/items")
	authrouter := r.Group("/auth")

	itemRouter.GET("", itemController.FindAll)
	// idはパスパラメーターなので「:」をつける
	itemRouter.GET("/:id", itemController.FindById)
	itemRouter.POST("", itemController.Create)
	itemRouter.PUT("/:id", itemController.Update)
	itemRouter.DELETE("/:id", itemController.Delete)

	authrouter.POST("/signup", authController.Signup)
	r.Run("localhost:8080")
}
