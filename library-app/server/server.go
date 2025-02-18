package server

import (
	"awesomeProject/library-app/global"
	"awesomeProject/library-app/handlers"
	"awesomeProject/library-app/managers"
	"awesomeProject/library-app/middlewares"
	"awesomeProject/library-app/security"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type Server struct {
	DB                 *gorm.DB
	Router             *mux.Router
	AuthClient         *security.AuthClient
	Config             *Config
	UserManager        *managers.UserManager
	ReservationManager *managers.ReservationManager
}

func (server *Server) Initialize(ctx context.Context) {
	server.InitializeAuthClient()
	server.InitializeDatabase()
	server.UserManager = managers.NewUserManager(server.DB)
	server.ReservationManager = managers.NewReservationManager(server.DB)
	server.InitializeRouter()
	server.StartReservationCleanupTicker(ctx)
}

func (server *Server) InitializeAuthClient() {
	server.AuthClient = &security.AuthClient{}
	server.AuthClient.Initialize(server.Config.AuthURL, server.Config.AuthRealm, server.Config.AuthClientID, server.Config.AuthClientSecret)
}

func (server *Server) InitializeRouter() {
	server.Router = mux.NewRouter()

	server.DefineRoutes()
}

func (server *Server) InitializeDatabase() {
	log.Info("[Server.InitializeDatabase] Connecting to database...")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		server.Config.POSTGRESHost, server.Config.POSTGRESPort, server.Config.POSTGRESUser,
		server.Config.POSTGRESPassword, server.Config.POSTGRESDb)

	const maxRetries = 5
	const retryDelay = 2 * time.Second

	for i := 0; i < maxRetries; i++ {
		db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
		if err != nil {
			log.Printf("[Server.InitializeDatabase] Failed to connect to database: %v", err)
			time.Sleep(retryDelay)
			continue
		}

		log.Info("[Server.InitializeDatabase] Connected to database")
		server.DB = db
		return
	}
	log.Fatal("[Server.InitializeDatabase] Failed to connect to database")
}

func (server *Server) StartWebServer() {
	log.Info("[Server.StartWebServer] Starting web server...")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})
	handler := c.Handler(server.Router)

	err := http.ListenAndServe(":"+server.Config.Port, handler)
	if err != nil {
		log.Fatal("[Server.StartWebServer] ListenAndServe: ", err)
	}
}

func setCORSHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4000")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}

func HandlePreflight(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		setCORSHeaders(w)
		log.Debug("setCORSHeaders successfully")
		if r.Method == http.MethodOptions {
			log.Debug("Setting CORS Headers")
			w.WriteHeader(http.StatusOK)
			return
		}

		nextHandler.ServeHTTP(w, r)
	})
}

func (server *Server) DefineRoutes() {
	log.Info("[Server.DefineRoutes] Defining routes...")
	server.Router.Use(HandlePreflight)
	server.Router.Use(middlewares.SetJSONMiddleware)

	userHandler := handlers.NewUserHandler(server.UserManager)
	authorHandler := handlers.NewAuthorHandler(managers.NewAuthorManager(server.DB))
	categoryHandler := handlers.NewCategoryHandler(managers.NewCategoryManager(server.DB))
	bookHandler := handlers.NewBookHandler(managers.NewBookManager(server.DB))
	loanHandler := handlers.NewLoanHandler(managers.NewLoanManager(server.DB))
	reservationHandler := handlers.NewReservationHandler(server.ReservationManager)
	ratingHandler := handlers.NewRatingHandler(managers.NewRatingManager(server.DB))

	server.Router.HandleFunc("/api/users", server.Protected(userHandler.GetAll, global.User, global.READ)).Methods(http.MethodGet)
	server.Router.HandleFunc("/api/users/{id:"+global.UuidRegex+"}", server.Protected(userHandler.Get, global.User, global.READ)).Methods(http.MethodGet)
	server.Router.HandleFunc("/api/users/{id:"+global.UuidRegex+"}", server.Protected(userHandler.Update, global.User, global.EDIT)).Methods(http.MethodPut)
	server.Router.HandleFunc("/api/users/{id:"+global.UuidRegex+"}", server.Protected(userHandler.Delete, global.User, global.DELETE)).Methods(http.MethodDelete)

	server.Router.HandleFunc("/api/authors", server.Protected(authorHandler.GetAll, global.Author, global.READ)).Methods(http.MethodGet)
	server.Router.HandleFunc("/api/authors", server.Protected(authorHandler.Create, global.Author, global.CREATE)).Methods(http.MethodPost)
	server.Router.HandleFunc("/api/authors/{id:"+global.UuidRegex+"}", server.Protected(authorHandler.Get, global.Author, global.READ)).Methods(http.MethodGet)
	server.Router.HandleFunc("/api/authors/{id:"+global.UuidRegex+"}", server.Protected(authorHandler.Update, global.Author, global.EDIT)).Methods(http.MethodPut)
	server.Router.HandleFunc("/api/authors/{id:"+global.UuidRegex+"}", server.Protected(authorHandler.Delete, global.Author, global.DELETE)).Methods(http.MethodDelete)

	server.Router.HandleFunc("/api/categories", server.Protected(categoryHandler.GetAll, global.Category, global.READ)).Methods(http.MethodGet)
	server.Router.HandleFunc("/api/categories", server.Protected(categoryHandler.Create, global.Category, global.CREATE)).Methods(http.MethodPost)
	server.Router.HandleFunc("/api/categories/{id:"+global.UuidRegex+"}", server.Protected(categoryHandler.Get, global.Category, global.READ)).Methods(http.MethodGet)
	server.Router.HandleFunc("/api/categories/{id:"+global.UuidRegex+"}", server.Protected(categoryHandler.Update, global.Category, global.EDIT)).Methods(http.MethodPut)
	server.Router.HandleFunc("/api/categories/{id:"+global.UuidRegex+"}", server.Protected(categoryHandler.Delete, global.Category, global.DELETE)).Methods(http.MethodDelete)

	server.Router.HandleFunc("/api/books", server.Protected(bookHandler.GetAll, global.Book, global.READ)).Methods(http.MethodGet)
	server.Router.HandleFunc("/api/books", server.Protected(bookHandler.Create, global.Book, global.CREATE)).Methods(http.MethodPost)
	server.Router.HandleFunc("/api/books/{id:"+global.UuidRegex+"}", server.Protected(bookHandler.Get, global.Book, global.READ)).Methods(http.MethodGet)
	server.Router.HandleFunc("/api/books/{id:"+global.UuidRegex+"}", server.Protected(bookHandler.Update, global.Book, global.EDIT)).Methods(http.MethodPut)
	server.Router.HandleFunc("/api/books/{id:"+global.UuidRegex+"}", server.Protected(bookHandler.Delete, global.Book, global.DELETE)).Methods(http.MethodDelete)

	server.Router.HandleFunc("/api/loans", server.Protected(loanHandler.GetAll, global.Loan, global.READ)).Methods(http.MethodGet)
	server.Router.HandleFunc("/api/loans", server.Protected(loanHandler.Create, global.Loan, global.CREATE)).Methods(http.MethodPost)
	server.Router.HandleFunc("/api/loans/{id:"+global.UuidRegex+"}", server.Protected(loanHandler.Get, global.Loan, global.READ)).Methods(http.MethodGet)
	server.Router.HandleFunc("/api/loans/{id:"+global.UuidRegex+"}", server.Protected(loanHandler.Update, global.Loan, global.EDIT)).Methods(http.MethodPut)
	server.Router.HandleFunc("/api/loans/{id:"+global.UuidRegex+"}", server.Protected(loanHandler.Delete, global.Loan, global.DELETE)).Methods(http.MethodDelete)

	server.Router.HandleFunc("/api/reservations", server.Protected(reservationHandler.GetAll, global.Reservation, global.READ)).Methods(http.MethodGet)
	server.Router.HandleFunc("/api/reservations", server.Protected(reservationHandler.Create, global.Reservation, global.CREATE)).Methods(http.MethodPost)
	server.Router.HandleFunc("/api/reservations/{id:"+global.UuidRegex+"}", server.Protected(reservationHandler.Get, global.Reservation, global.READ)).Methods(http.MethodGet)
	server.Router.HandleFunc("/api/reservations/{id:"+global.UuidRegex+"}", server.Protected(reservationHandler.Update, global.Reservation, global.EDIT)).Methods(http.MethodPut)
	server.Router.HandleFunc("/api/reservations/{id:"+global.UuidRegex+"}", server.Protected(reservationHandler.Delete, global.Reservation, global.DELETE)).Methods(http.MethodDelete)

	server.Router.HandleFunc("/api/ratings", server.Protected(ratingHandler.GetAll, global.Rating, global.READ)).Methods(http.MethodGet)
	server.Router.HandleFunc("/api/ratings", server.Protected(ratingHandler.Create, global.Rating, global.CREATE)).Methods(http.MethodPost)
	server.Router.HandleFunc("/api/ratings/{id:"+global.UuidRegex+"}", server.Protected(ratingHandler.Get, global.Rating, global.READ)).Methods(http.MethodGet)
	server.Router.HandleFunc("/api/ratings/{id:"+global.UuidRegex+"}", server.Protected(ratingHandler.Update, global.Rating, global.EDIT)).Methods(http.MethodPut)
	server.Router.HandleFunc("/api/ratings/{id:"+global.UuidRegex+"}", server.Protected(ratingHandler.Delete, global.Rating, global.DELETE)).Methods(http.MethodDelete)

	log.Info("[Server.DefineRoutes] Defined routes")
}

func (server *Server) StartReservationCleanupTicker(ctx context.Context) {
	ticker := time.NewTicker(global.ReservationCleanupInterval)

	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				server.ReservationManager.CleanupExpiredReservations()
			case <-ctx.Done():
				return
			}
		}
	}()
}

func (server *Server) RunReservationCleanupTicker(ticker *time.Ticker) {
	for range ticker.C {
		server.ReservationManager.CleanupExpiredReservations()
	}
}
