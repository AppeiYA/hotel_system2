package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	guest_postgres "hotel_system2/internal/guest/adapters/postgres"
	guest_usecase "hotel_system2/internal/guest/use_case"
	"hotel_system2/internal/http"
	payment_external "hotel_system2/internal/payment/adapters/external"
	payment_http "hotel_system2/internal/payment/adapters/http"
	"hotel_system2/internal/payment/adapters/mock_gateway"
	payment_postgres "hotel_system2/internal/payment/adapters/postgres"
	payment_usecase "hotel_system2/internal/payment/use_case"
	reservation_external "hotel_system2/internal/reservation/adapters/external"
	reservation_http "hotel_system2/internal/reservation/adapters/http"
	reservation_postgres "hotel_system2/internal/reservation/adapters/postgres"
	reservation_usecase "hotel_system2/internal/reservation/use_case"
	room_http "hotel_system2/internal/room/adapters/http"
	room_postgres "hotel_system2/internal/room/adapters/postgres"
	room_usecase "hotel_system2/internal/room/use_case"
	"hotel_system2/internal/shared/config"
	database "hotel_system2/internal/shared/db"
	"hotel_system2/internal/shared/logger"
	"hotel_system2/internal/shared/session"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	_ "hotel_system2/docs"

	swagger "github.com/gofiber/swagger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.uber.org/zap"
)

// @title Hotel Reservation API
// @version 1.0
// @description Hotel Reservation System API
// @host localhost:3333
// @BasePath /api/v1
// @schemes http

func main() {
	// Load configuration
	cfg := config.SetupConfig()

	// Connect database
	db, err := database.ConnectDB(cfg)
	if err != nil {
		logger.Fatal("Error connecting db", zap.Error(err))
	}
	defer db.Close()

	txManager := database.NewTransactionManager(db)

	// Initialize session store
	session.InitSessionStore(cfg)

	app := fiber.New(fiber.Config{
		JSONEncoder: func(v interface{}) ([]byte, error) {
			buf := &bytes.Buffer{}
			encoder := json.NewEncoder(buf)
			encoder.SetEscapeHTML(false)

			err := encoder.Encode(v)
			return bytes.TrimRight(buf.Bytes(), "\n"), err
		},
		EnableTrustedProxyCheck: true,
		TrustedProxies:          []string{"0.0.0.0/0"},
		BodyLimit:               5 * 1024 * 1024,
	})

	// Middleware
	app.Use(cors.New(cors.Config{
		AllowOriginsFunc: func(origin string) bool {
			return true
		},
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS,PATCH",
		AllowCredentials: true,
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, session_id",
	}))

	app.Use(recover.New())

	app.Options("/*", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNoContent)
	})

	// ===========================
	// Repositories
	// ===========================

	roomRepo := room_postgres.NewRepository(db)
	guestRepo := guest_postgres.NewRepository(db)
	reservationRepo := reservation_postgres.NewRepository(db)
	paymentRepo := payment_postgres.NewRepository(db)
	mockgateway := mock_gateway.NewGateway()

	// ===========================
	// Use Cases
	// ===========================

	// room
	// _ = room_usecase.(roomRepo)
	listRooms := room_usecase.NewListRooms(roomRepo)
	// _ = room_usecase.NewUpdateRoomStatus(roomRepo)

	// guest
	_ = guest_usecase.NewCreateGuest(guestRepo)

	// reservationConfirmer
	reservationConfirmer := payment_external.NewReservationConfirmationAdapter(reservationRepo)

	// payment
	initializePayment := payment_usecase.NewInitializePayment(
		txManager,
		paymentRepo,
		reservationConfirmer,
		guestRepo,
		mockgateway,
	)
	completePayment := payment_usecase.NewCompletePayment(
		txManager,
		paymentRepo,
		reservationConfirmer,
		mockgateway,
	)

	// payment lookup adapter in reservation
	reservation_payment_lookup := reservation_external.NewPaymentLookupAdapter(paymentRepo)

	// reservation
	createReservation := reservation_usecase.NewCreateReservation(
		txManager,
		reservationRepo,
		roomRepo,
		guestRepo,
		reservation_payment_lookup,
	)
	listReservationByEmail := reservation_usecase.NewListReservationByEmail(reservationRepo)
	listReservations := reservation_usecase.NewListReservations(reservationRepo)
	checkIn := reservation_usecase.NewCheckIn(
		txManager,
		reservationRepo,
		roomRepo,
	)
	checkOut := reservation_usecase.NewCheckOut(
		txManager,
		reservationRepo,
		roomRepo,
	)


	// ===========================
	// Handlers
	// ===========================
	roomHandler := room_http.NewHandler(
		listRooms,
	)

	reservationHandler := reservation_http.NewHandler(
		*createReservation,
		*listReservationByEmail,
		*listReservations,
		*checkIn,
		*checkOut,
	)

	paymentHandler := payment_http.NewHandler(
		initializePayment,
		completePayment,
	)

	// ===========================
	// Routes
	// ===========================
	
	// Swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

	http.SetupAppRoutes(
		app,
		roomHandler,
		reservationHandler,
		paymentHandler,
	)

	port := strings.TrimSpace(cfg.PORT)

	// Graceful shutdown
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		<-c

		fmt.Println("Gracefully shutting down...")
		_ = app.Shutdown()
	}()

	log.Printf("Server running on :%s\n", port)

	if err := app.Listen(":" + port); err != nil {
		log.Fatal(err)
	}
}