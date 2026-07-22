package reservation_http

import (
	usecase "hotel_system2/internal/reservation/use_case"
	"hotel_system2/internal/shared/logger"
	"hotel_system2/internal/shared/response"
	"hotel_system2/internal/shared/utils"
	shared_http "hotel_system2/internal/shared/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Handler struct {
	createReservation usecase.CreateReservation
	listReservationByEmail usecase.ListReservationByEmail
	listReservations usecase.ListReservations
	checkIn           usecase.CheckIn
	checkOut          usecase.CheckOut
}

func NewHandler(
	createReservation usecase.CreateReservation,
	listReservationByEmail usecase.ListReservationByEmail,
	listReservations usecase.ListReservations,
	checkIn usecase.CheckIn,
	checkOut usecase.CheckOut,
) *Handler {
	return &Handler{
		createReservation: createReservation, 
		listReservationByEmail: listReservationByEmail,
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

	input, err := req.toInput()
	if err != nil {
		logger.Error("Error converting request to input", zap.Error(err))
		return response.Error(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	resp, err := h.createReservation.Execute(c.Context(), input)
	if err != nil {
		logger.Error("Error creating reservation at handlers.CreateReservation", zap.Error(err))
		return response.Error(c, fiber.StatusInternalServerError, err.Error(), nil)
	}
	output := toReservationDetailsOnly(resp)

	return response.JSON(c, fiber.StatusCreated, "Reservation created successfully", output)
}

func (h *Handler) ListReservationsByEmail(c *fiber.Ctx) error {
	email := c.Params("email")
	if email == "" {
		return response.Error(c, fiber.StatusBadRequest, "invalid email", nil)
	}

	resp, err := h.listReservationByEmail.Execute(c.Context(), email)
	if err != nil {
		logger.Error("Error fetching reservation at handlers.ListReservationsByEmail", zap.Error(err))
		return response.Error(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	result := []reservationDetailsOnly{}
	for _, reservation := range resp {
		output := toReservationDetailsOnly(reservation)
		result = append(result, output)
	}

	return response.JSON(c, fiber.StatusOK, "Reservation fetched successfully", result)
}

func (h *Handler) ListReservations(c *fiber.Ctx) error {
	resp, err := h.listReservations.Execute(c.Context())
	if err != nil {
		logger.Error("Error listing reservations at handlers.ListReservations", zap.Error(err))
		return response.Error(c, fiber.StatusInternalServerError, err.Error(), nil)
	}

	result := []reservationDetailsOnly{}
	for _, reservation := range resp {
		output := toReservationDetailsOnly(reservation)
		result = append(result, output)
	}

	return response.JSON(c, fiber.StatusOK, "Reservations listed successfully", result)
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
		return response.Error(c, shared_http.StatusFor(err), err.Error(), nil)
	}

	return response.JSON(
		c,
		fiber.StatusOK,
		"Reservation checked out successfully",
		nil,
	)
}