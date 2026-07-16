package room_http

import (
	usecase "hotel_system2/internal/room/use_case"
	"hotel_system2/internal/shared/logger"
	"hotel_system2/internal/shared/response"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Handler struct{
	listRooms *usecase.ListRooms
}

func NewHandler(listRooms *usecase.ListRooms) *Handler {
	return &Handler{
		listRooms: listRooms,
	}
}

func (h *Handler) ListRooms(c *fiber.Ctx) error {
	rooms, err := h.listRooms.Execute(c.Context())
	if err != nil {
		logger.Error("Error fetching rooms", zap.Error(err))
		return response.Error(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	resp := make([]roomResponse, len(rooms))
	for i, room := range rooms {
		resp[i] = toRoomResponse(&room)
	}

	return response.JSON(c, fiber.StatusOK, "Rooms fetched successfully", rooms)
}