package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"itmx_test/config"
	"itmx_test/middleware"
	"itmx_test/service/entity"
	"itmx_test/service/customer/delivery"
	"itmx_test/service/customer/repository"
	"itmx_test/service/customer/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"
)

func init() {
	configFile := "./config.dev.yml"

	env := os.Getenv("APP_ENV")
	if env == "production" {
		log.Println(env)
		configFile = "./config.prod.yml"
	}

	viper.SetConfigFile(configFile)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
}

func main() {
	dbConn := config.InitDB()

	f := fiber.New(fiber.Config{
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		Prefork:      false,
		ServerHeader: "Fiber",
		// BodyLimit:   30 * 1024 * 1024, // 30 MB
	})

	corsAllowList := viper.GetStringSlice(`header.cors`)
	middL := middleware.CORSMiddleware(corsAllowList)
	f.Use(middL)

	loggerMiddleware := logger.New(logger.Config{
		TimeFormat: "2006-01-02 15:04:05",
		Format:     "${time} | ${status} | ${latency} | ${ips} | ${method} | ${path}\n",
	})
	f.Use(loggerMiddleware)

	customerRepo := repository.NewCustomerRepository(dbConn)

	customerUsecase := usecase.NewCustomerUsecase(customerRepo)

	delivery.NewCustomerHandler(f, customerUsecase)

	// initially create some customer
	customers := []*entity.Customer{
		{Name: "John Doe", Age: 23},
		{Name: "Jane Smith", Age: 44},
	}

	for _, customer := range customers {
		customerUsecase.CreateCustomer(customer)
	}

	f.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "pong",
		})
	})

	log.Fatal(f.Listen(fmt.Sprintf("%s:%s", viper.GetString(`server.host`), viper.GetString("server.port"))))
}
