package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	guest_postgres "hotel_system2/internal/guest/adapters/postgres"
	guest_usecase "hotel_system2/internal/guest/use_case"
	"hotel_system2/internal/http"
	ledger_external "hotel_system2/internal/ledger/adapters/external"
	ledger_postgres "hotel_system2/internal/ledger/adapters/postgres"
	ledger_usecase "hotel_system2/internal/ledger/use_case"
	payment_external "hotel_system2/internal/payment/adapters/external"
	payment_http "hotel_system2/internal/payment/adapters/http"
	"hotel_system2/internal/payment/adapters/mock_gateway"
	payment_postgres "hotel_system2/internal/payment/adapters/postgres"
	payment_usecase "hotel_system2/internal/payment/use_case"
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

	// mock
	mockgateway := mock_gateway.NewGateway(cfg.BaseUrl)

	// ledger
	ledgerAccountRepo := ledger_postgres.NewAccountRepository(db)
	ledgerTxRepo := ledger_postgres.NewTransactionRepository(db)

	postRoomCharge := ledger_usecase.NewPostRoomCharge(ledgerAccountRepo, ledgerTxRepo)
	postPaymentReceived := ledger_usecase.NewPostPaymentReceived(ledgerAccountRepo, ledgerTxRepo)
	getFolioBalance := ledger_usecase.NewGetFolioBalance(ledgerAccountRepo, ledgerTxRepo)

	reservationLedgerAdapter := ledger_external.NewReservationLedgerAdapter(postRoomCharge, getFolioBalance)
	paymentLedgerAdapter := ledger_external.NewPaymentLedgerAdapter(postPaymentReceived)

	// room
	roomRepo := room_postgres.NewRepository(db)

	// _ = room_usecase.(roomRepo)
	listRooms := room_usecase.NewListRooms(roomRepo)
	// _ = room_usecase.NewUpdateRoomStatus(roomRepo)

	roomHandler := room_http.NewHandler(
		listRooms,
	)

	// guest
	guestRepo := guest_postgres.NewRepository(db)

	_ = guest_usecase.NewCreateGuest(guestRepo)

	// reservation
	reservationRepo := reservation_postgres.NewRepository(db)

	createReservation := reservation_usecase.NewCreateReservation(
		txManager,
		reservationRepo,
		roomRepo,
		guestRepo,
		reservationLedgerAdapter,
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
		reservationLedgerAdapter,
	)

	reservationConfirmer := payment_external.NewReservationConfirmationAdapter(reservationRepo)


	reservationHandler := reservation_http.NewHandler(
		*createReservation,
		*listReservationByEmail,
		*listReservations,
		*checkIn,
		*checkOut,
	)
	//

	// payment
	paymentRepo := payment_postgres.NewRepository(db)

	initializePayment := payment_usecase.NewInitializePayment(
		paymentRepo,
		reservationRepo,
		guestRepo,
		mockgateway,
	)
	completePayment := payment_usecase.NewCompletePayment(
		txManager,
		paymentRepo,
		reservationConfirmer,
		mockgateway,
		paymentLedgerAdapter,
	)
	
	paymentHandler := payment_http.NewHandler(
		initializePayment,
		completePayment,
	)
	
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