package main

import (
	_ "embed"
	"log"

	"github.com/gofiber/fiber/v3"
	"gopkg.in/yaml.v3"
)

type BarcodeConfig struct {
	Codes map[string]string `yaml:"codes"`
}

//go:embed data.yml
var barcodeConfig string
var barCodeMap map[string]string

func init() {
	var config BarcodeConfig
	if err := yaml.Unmarshal([]byte(barcodeConfig), &config); err != nil {
		panic(err)
	}
	barCodeMap = config.Codes
}

func main() {
	app := fiber.New()

	app.Get("/code/:barcode", func(c fiber.Ctx) error {
		if url, ok := barCodeMap[c.Params("barcode")]; ok {
			return c.Redirect().To(url)
		}

		return c.SendStatus(fiber.StatusNotFound)
	})

	log.Fatal(app.Listen(":3500"))
}
