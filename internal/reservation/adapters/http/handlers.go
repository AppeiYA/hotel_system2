package reservation_http

import (
	"errors"
	reservation_domain "hotel_system2/internal/reservation/domain"
	usecase "hotel_system2/internal/reservation/use_case"
	"hotel_system2/internal/shared/logger"
	"hotel_system2/internal/shared/response"
	"hotel_system2/internal/shared/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Handler struct {
	createReservation usecase.CreateReservation
	getReservation usecase.GetReservation
	listReservations usecase.ListReservations
	checkIn           usecase.CheckIn
	checkOut          usecase.CheckOut
}

func NewHandler(
	createReservation usecase.CreateReservation,
	getReservation usecase.GetReservation,
	listReservations usecase.ListReservations,
	checkIn usecase.CheckIn,
	checkOut usecase.CheckOut,
) *Handler {
	return &Handler{
		createReservation: createReservation, 
		getReservation: getReservation,
		listReservations: listReservations,
		checkIn:           checkIn,
		checkOut:          checkOut,
	}
}

func (h *Handler) CreateReservation(c *fiber.Ctx) error {
	var req createReservationRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	if err := utils.Validate.Struct(req); err != nil {
		errs := utils.ValidationErrors(err)
		return response.Error(c, fiber.StatusUnprocessableEntity, err.Error(), errs)
	}

	if err := req.Validate(); err != nil {
		logger.Error("Invalid check in and check out", zap.Error(err))
		return response.Error(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	input := usecase.CreateReservationInput{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.Phone,
		RoomID:    req.RoomID,
		CheckIn:   req.CheckIn,
		CheckOut:  req.CheckOut,
	}

	resp, err := h.createReservation.Execute(c.Context(), input)
	if err != nil {
		logger.Error("Error creating reservation at handlers.CreateReservation", zap.Error(err))
		return response.Error(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return response.JSON(c, fiber.StatusCreated, "Reservation created successfully", resp)
}

func (h *Handler) GetReservation(c *fiber.Ctx) error {
	id := c.Params("id")
	_, err := uuid.Parse(id)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid reservation id", err)
	}

	resp, err := h.getReservation.Execute(c.Context(), id)
	if err != nil {
		logger.Error("Error fetching reservation at handlers.GetReservation", zap.Error(err))
		return response.Error(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	return response.JSON(c, fiber.StatusOK, "Reservation fetched successfully", resp)
}

func (h *Handler) ListReservations(c *fiber.Ctx) error {
	resp, err := h.listReservations.Execute(c.Context())
	if err != nil {
		logger.Error("Error listing reservations at handlers.ListReservations", zap.Error(err))
		return response.Error(c, fiber.StatusInternalServerError, err.Error(), nil)
	}
	return response.JSON(c, fiber.StatusOK, "Reservations listed successfully", resp)
}

func (h *Handler) CheckIn(c *fiber.Ctx) error {
	reservationID := c.Params("id")
	_, err := uuid.Parse(reservationID)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid reservation id", nil)
	}

	if err := h.checkIn.Execute(
		c.Context(),
		reservationID,
	); err != nil {
		return response.Error(
			c,
			fiber.StatusBadRequest,
			err.Error(),
			nil,
		)
	}

	return response.JSON(
		c,
		fiber.StatusOK,
		"Guest checked in successfully",
		nil,
	)
}

func (h *Handler) CheckOut(c *fiber.Ctx) error {
	id := c.Params("id")

	if _, err := uuid.Parse(id); err != nil {
		return response.Error(
			c,
			fiber.StatusBadRequest,
			"invalid reservation id",
			nil,
		)
	}

	err := h.checkOut.Execute(c.Context(), id)
	if err != nil {

		logger.Error(
			"failed to check out reservation",
			zap.String("reservation_id", id),
			zap.Error(err),
		)

		switch {

		case errors.Is(err, reservation_domain.ErrReservationNotFound):
			return response.Error(c, fiber.StatusNotFound, err.Error(), nil)

		case errors.Is(err, reservation_domain.ErrReservationNotCheckedIn):
			return response.Error(c, fiber.StatusConflict, err.Error(), nil)

		case errors.Is(err, reservation_domain.ErrFolioBalanceOutstanding):
			return response.Error(c, fiber.StatusConflict, err.Error(), nil)

		default:
			return response.Error(
				c,
				fiber.StatusInternalServerError,
				"failed to check out reservation",
				nil,
			)
		}
	}

	return response.JSON(
		c,
		fiber.StatusOK,
		"Reservation checked out successfully",
		nil,
	)
}