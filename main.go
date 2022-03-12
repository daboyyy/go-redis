package main

import (
	"go-redis/repositories"
	"go-redis/services"

	"github.com/gofiber/fiber/v2"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func main() {
	db := initDatabase()
	redisClient := initRedis()

	productRepo := repositories.NewProductRepositoryDB(db)
	productService := services.NewCatalogServiceRedis(productRepo, redisClient)
	_ = productService

	app := fiber.New()
	app.Get("/hello", func(c *fiber.Ctx) error {
		// time.Sleep(time.Millisecond * 10)
		return c.SendString("Hello World")
	})
	app.Listen(":8000")
}

func initDatabase() *gorm.DB {
	dial := mysql.Open("root:P@ssw0rd@tcp(localhost:3306)/infinitas")
	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func initRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
