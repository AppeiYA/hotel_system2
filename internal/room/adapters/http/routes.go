package room_http

import "github.com/gofiber/fiber/v2"

func RegisterRoomRoutes(router fiber.Router, handler *Handler) {
	rooms := router.Group("/rooms")

	rooms.Get("/", handler.ListRooms)
}