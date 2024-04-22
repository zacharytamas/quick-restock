package main

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"gopkg.in/yaml.v3"
)

// BarcodeMatch contains information about a product associated with a bar code.
type BarcodeMatch struct {
	Barcode string `yaml:"barcode"`
	Name    string `yaml:"name"`
	Kind    string `yaml:"kind"`
	Id      string `yaml:"id"`
}

// BarcodeConfig is the configuration for mapping bar codes to URLs.
type BarcodeConfig map[string]BarcodeMatch

//go:embed data.yml
var barcodeConfig string
var barCodeMap BarcodeConfig

func init() {
	var config BarcodeConfig
	if err := yaml.Unmarshal([]byte(barcodeConfig), &config); err != nil {
		panic(err)
	}
	barCodeMap = config
}

func newApp() *fiber.App {
	app := fiber.New()
	app.Use(logger.New())

	app.Get("/code/:barcode<len(13)>", func(c fiber.Ctx) error {
		if match, ok := barCodeMap[c.Params("barcode")]; ok {
			log.Infof("Successful barcode match. %s -> %s", c.Params("barcode"), match)
			switch match.Kind {
			case "walmart":
				return c.Redirect().To(fmt.Sprintf("https://walmart.com/ip/%s", match.Id))
			default:
				c.Response().SetStatusCode(fiber.StatusUnprocessableEntity)
				return c.SendString("Unknown kind: " + match.Kind)
			}
		}

		log.Warnf("No barcode match for %s", c.Params("barcode"))
		return c.SendStatus(fiber.StatusNotFound)
	})

	return app
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3500"
	}

	app := newApp()

	if err := app.Listen(":" + port); err != nil {
		log.Fatal(err)
	}
}
