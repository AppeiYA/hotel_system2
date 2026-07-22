package reservation_http

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(router fiber.Router, handler *Handler) {
	reservations := router.Group("/reservations")
	reservations.Post("/", handler.CreateReservation)
	reservations.Get("/", handler.ListReservations)
	reservations.Get("/:email", handler.ListReservationsByEmail)
	reservations.Post("/:id/check-in", handler.CheckIn)
	reservations.Post("/:id/check-out", handler.CheckOut)
}