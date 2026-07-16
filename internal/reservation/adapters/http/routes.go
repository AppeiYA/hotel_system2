package reservation_http

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(router fiber.Router, handler *Handler) {
	reservations := router.Group("/reservations")
	reservations.Post("/", handler.CreateReservation)
	reservations.Get("/", handler.ListReservations)
	reservations.Get("/:id", handler.GetReservation)
	reservations.Post("/check-in", handler.CheckIn)
	reservations.Post("/check-out", handler.CheckOut)
}