package server

import (
	"awesomeProject/library-app/global"
	"awesomeProject/library-app/handlers"
	"awesomeProject/library-app/managers"
	"awesomeProject/library-app/middlewares"
	"awesomeProject/library-app/security"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type Server struct {
	DB         *gorm.DB
	Router     *mux.Router
	AuthClient *security.AuthClient
	Config     *Config
}

func (server *Server) Initialize() {
	server.InitializeAuthClient()
	server.InitializeDatabase()
	server.InitializeRouter()
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
	err := http.ListenAndServe(":"+server.Config.Port, server.Router)
	if err != nil {
		log.Fatal("[Server.StartWebServer] ListenAndServe: ", err)
	}
}

func (server *Server) DefineRoutes() {
	log.Info("[Server.DefineRoutes] Defining routes...")
	server.Router.Use(middlewares.SetJSONMiddleware)

	userHandler := handlers.NewUserHandler(managers.NewUserManager(server.DB))
	authorHandler := handlers.NewAuthorHandler(managers.NewAuthorManager(server.DB))
	categoryHandler := handlers.NewCategoryHandler(managers.NewCategoryManager(server.DB))
	bookHandler := handlers.NewBookHandler(managers.NewBookManager(server.DB))
	loanHandler := handlers.NewLoanHandler(managers.NewLoanManager(server.DB))
	reservationHandler := handlers.NewReservationHandler(managers.NewReservationManager(server.DB))
	reviewHandler := handlers.NewReviewHandler(managers.NewReviewManager(server.DB))

	server.Router.HandleFunc("/api/users", server.Protected(userHandler.GetAll, global.User, global.READ)).Methods(http.MethodGet)
	server.Router.HandleFunc("/api/users/{id:"+global.UuidRegex+"}", server.Protected(userHandler.Get, global.User, global.READ)).Methods(http.MethodGet)
	server.Router.HandleFunc("/api/users/{id:"+global.UuidRegex+"}", server.Protected(userHandler.Update, global.User, global.WRITE)).Methods(http.MethodPut)
	server.Router.HandleFunc("/api/users/{id:"+global.UuidRegex+"}", server.Protected(userHandler.Delete, global.User, global.WRITE)).Methods(http.MethodDelete)

	server.Router.HandleFunc("/api/authors", server.Protected(authorHandler.GetAll, global.Author, global.READ)).Methods(http.MethodGet)
	server.Router.HandleFunc("/api/authors", server.Protected(authorHandler.Create, global.Author, global.WRITE)).Methods(http.MethodPost)
	server.Router.HandleFunc("/api/authors/{id:"+global.UuidRegex+"}", server.Protected(authorHandler.Get, global.Author, global.READ)).Methods(http.MethodGet)
	server.Router.HandleFunc("/api/authors/{id:"+global.UuidRegex+"}", server.Protected(authorHandler.Update, global.Author, global.WRITE)).Methods(http.MethodPut)
	server.Router.HandleFunc("/api/authors/{id:"+global.UuidRegex+"}", server.Protected(authorHandler.Delete, global.Author, global.WRITE)).Methods(http.MethodDelete)

	server.Router.HandleFunc("/api/categories", server.Protected(categoryHandler.GetAll, global.Category, global.READ)).Methods(http.MethodGet)
	server.Router.HandleFunc("/api/categories", server.Protected(categoryHandler.Create, global.Category, global.WRITE)).Methods(http.MethodPost)
	server.Router.HandleFunc("/api/categories/{id:"+global.UuidRegex+"}", server.Protected(categoryHandler.Get, global.Category, global.READ)).Methods(http.MethodGet)
	server.Router.HandleFunc("/api/categories/{id:"+global.UuidRegex+"}", server.Protected(categoryHandler.Update, global.Category, global.WRITE)).Methods(http.MethodPut)
	server.Router.HandleFunc("/api/categories/{id:"+global.UuidRegex+"}", server.Protected(categoryHandler.Delete, global.Category, global.WRITE)).Methods(http.MethodDelete)

	server.Router.HandleFunc("/api/books", server.Protected(bookHandler.GetAll, global.Book, global.READ)).Methods(http.MethodGet)
	server.Router.HandleFunc("/api/books", server.Protected(bookHandler.Create, global.Book, global.WRITE)).Methods(http.MethodPost)
	server.Router.HandleFunc("/api/books/{id:"+global.UuidRegex+"}", server.Protected(bookHandler.Get, global.Book, global.READ)).Methods(http.MethodGet)
	server.Router.HandleFunc("/api/books/{id:"+global.UuidRegex+"}", server.Protected(bookHandler.Update, global.Book, global.WRITE)).Methods(http.MethodPut)
	server.Router.HandleFunc("/api/books/{id:"+global.UuidRegex+"}", server.Protected(bookHandler.Delete, global.Book, global.WRITE)).Methods(http.MethodDelete)

	server.Router.HandleFunc("/api/loans", server.Protected(loanHandler.GetAll, global.Loan, global.READ)).Methods(http.MethodGet)
	server.Router.HandleFunc("/api/loans", server.Protected(loanHandler.Create, global.Loan, global.WRITE)).Methods(http.MethodPost)
	server.Router.HandleFunc("/api/loans/{id:"+global.UuidRegex+"}", server.Protected(loanHandler.Get, global.Loan, global.READ)).Methods(http.MethodGet)
	server.Router.HandleFunc("/api/loans/{id:"+global.UuidRegex+"}", server.Protected(loanHandler.Update, global.Loan, global.WRITE)).Methods(http.MethodPut)
	server.Router.HandleFunc("/api/loans/{id:"+global.UuidRegex+"}", server.Protected(loanHandler.Delete, global.Loan, global.WRITE)).Methods(http.MethodDelete)

	server.Router.HandleFunc("/api/reservations", server.Protected(reservationHandler.GetAll, global.Reservation, global.READ)).Methods(http.MethodGet)
	server.Router.HandleFunc("/api/reservations", server.Protected(reservationHandler.Create, global.Reservation, global.WRITE)).Methods(http.MethodPost)
	server.Router.HandleFunc("/api/reservations/{id:"+global.UuidRegex+"}", server.Protected(reservationHandler.Get, global.Reservation, global.READ)).Methods(http.MethodGet)
	server.Router.HandleFunc("/api/reservations/{id:"+global.UuidRegex+"}", server.Protected(reservationHandler.Update, global.Reservation, global.WRITE)).Methods(http.MethodPut)
	server.Router.HandleFunc("/api/reservations/{id:"+global.UuidRegex+"}", server.Protected(reservationHandler.Delete, global.Reservation, global.WRITE)).Methods(http.MethodDelete)

	server.Router.HandleFunc("/api/reviews", server.Protected(reviewHandler.GetAll, global.Review, global.READ)).Methods(http.MethodGet)
	server.Router.HandleFunc("/api/reviews", server.Protected(reviewHandler.Create, global.Review, global.WRITE)).Methods(http.MethodPost)
	server.Router.HandleFunc("/api/reviews/{id:"+global.UuidRegex+"}", server.Protected(reviewHandler.Get, global.Review, global.READ)).Methods(http.MethodGet)
	server.Router.HandleFunc("/api/reviews/{id:"+global.UuidRegex+"}", server.Protected(reviewHandler.Update, global.Review, global.WRITE)).Methods(http.MethodPut)
	server.Router.HandleFunc("/api/reviews/{id:"+global.UuidRegex+"}", server.Protected(reviewHandler.Delete, global.Review, global.WRITE)).Methods(http.MethodDelete)

	log.Info("[Server.DefineRoutes] Defined routes")
}
