package http

import (
	reservation_http "hotel_system2/internal/reservation/adapters/http"
	room_http "hotel_system2/internal/room/adapters/http"
	"hotel_system2/internal/shared/response"

	"github.com/gofiber/fiber/v2"
)

func SetupAppRoutes(
	app *fiber.App, 
	rh *room_http.Handler,
	reh *reservation_http.Handler,
	) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Hotel System API")
	})

	v1 := app.Group("/api/v1")

	v1.Get("/health", func(c *fiber.Ctx) error {
		return response.JSON(c, fiber.StatusOK, "API is Healthy", nil)
	})

	room_http.RegisterRoomRoutes(v1, rh)
	reservation_http.RegisterRoutes(v1, reh)

	app.Use(func(c *fiber.Ctx) error {
		return response.Error(c, fiber.StatusNotFound, "route not found", nil)
	})
}