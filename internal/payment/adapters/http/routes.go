package payment_http

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(router fiber.Router, handler *Handler) {
	payments := router.Group("/payments")
	payments.Post("/initialize", handler.Initialize)
	payments.Post("/complete", handler.Complete)
}