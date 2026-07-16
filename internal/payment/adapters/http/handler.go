package payment_http

import (
	payment_usecase "hotel_system2/internal/payment/use_case"
	"hotel_system2/internal/shared/logger"
	"hotel_system2/internal/shared/response"
	"hotel_system2/internal/shared/utils"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Handler struct {
	initializePayment *payment_usecase.InitializePayment
	completePayment     *payment_usecase.CompletePayment
}

func NewHandler(
	initializePayment *payment_usecase.InitializePayment,
	completePayment *payment_usecase.CompletePayment,
) *Handler {
	return &Handler{
		initializePayment: initializePayment,
		completePayment:     completePayment,
	}
}

func (h *Handler) Initialize(c *fiber.Ctx) error {

	var req initializePaymentRequest

	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid request body", nil)
	}

	if err := utils.Validate.Struct(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	resp, err := h.initializePayment.Execute(
		c.Context(),
		payment_usecase.InitializePaymentInput{
			ReservationID: req.ReservationID,
		},
	)

	if err != nil {
		logger.Error(
			"failed to initialize payment",
			zap.Error(err),
		)

		return response.Error(
			c,
			fiber.StatusInternalServerError,
			err.Error(),
			nil,
		)
	}

	return response.JSON(
		c,
		fiber.StatusCreated,
		"Payment initialized successfully",
		resp,
	)
}

func (h *Handler) Complete(c *fiber.Ctx) error {

	reference := c.Params("reference")

	if reference == "" {
		return response.Error(
			c,
			fiber.StatusBadRequest,
			"missing payment reference",
			nil,
		)
	}

	if err := h.completePayment.Execute(
		c.Context(),
		reference,
	); err != nil {

		logger.Error(
			"failed to complete payment",
			zap.Error(err),
			zap.String("reference", reference),
		)

		return response.Error(
			c,
			fiber.StatusInternalServerError,
			err.Error(),
			nil,
		)
	}

	return response.JSON(
		c,
		fiber.StatusOK,
		"Payment verified successfully",
		nil,
	)
}