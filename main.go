package main

import (
	"context"
	"errors"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func main() {

}

func run() error {
}

func dbInit() error {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		return errors.New("you must set your 'MONGODB_URI' environmental variable")
	}
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}
	db = client.Database("myDB")
	return nil
}

func closeDB() error {
	return db.Client().Disconnect(context.Background())
}

func GetDBCollection(col string) *mongo.Collection {
	return db.Collection(col)
}

func LoadEnv() error {
	prod := os.Getenv("PROD")

	if prod != "true" {
		err := godotenv.Load()
		if err != nil {
			return err
		}
	}

	return nil
}

func runApp() error {
	err := LoadEnv()
	if err != nil {
		return err
	}
	err = dbInit()
	if err != nil {
		return err
	}
	defer closeDB()
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	var port string
	if port = os.Getenv("PORT"); port != "" {
		port = "8080"
	}
	app.Listen(":" + port)
	return nil
}

func addAccountGroup(app *fiber.App) {
	accountGroup := app.Group("/accounts")
	accountGroup.Get("/", getAccounts)
	accountGroup.Get("/:id", getAccount)
	accountGroup.Post("/", createAccount)
	accountGroup.Put("/:id", updateAccount)
	accountGroup.Delete("/:id", deleteAccount)
}

func getAccounts(c *fiber.Ctx) error {
	return nil
}

func getAccount(c *fiber.Ctx) error {
	return nil
}

func createAccount(c *fiber.Ctx) error {
	return nil
}

func updateAccount() {

}
func deleteAccount() {

}
