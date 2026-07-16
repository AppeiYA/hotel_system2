package room_http

import (
	_ "hotel_system2/internal/shared/response"
)

// ListRooms godoc
//
// @Summary Get Rooms
// @Description Retrieve a list of all available rooms
// @Tags Rooms
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{data=[]roomResponse}
// @Failure 500 {object} response.ErrorResponse
// @Router /rooms [get]
func _ListRooms(){}