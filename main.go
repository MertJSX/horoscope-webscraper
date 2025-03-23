package main

import (
	"log"

	"github.com/MertJSX/horoscope-webscrapper/utils"
	"github.com/gofiber/fiber/v2"
)

type HoroscopeResponse struct {
	Message string `json:"message"`
}

func main() {
	app := fiber.New()

	app.Get("/horoscope", func(c *fiber.Ctx) error {
		horoscope := c.Query("horoscope")
		time := c.Query("time")
		message := utils.GetHoroscopeComment(horoscope, time)

		log.Println(message)

		return c.JSON(HoroscopeResponse{Message: message})
	})

	app.Listen(":3000")
}
