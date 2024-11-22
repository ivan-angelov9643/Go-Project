package main

import (
	"awesomeProject/todo-app/configuration"
	"awesomeProject/todo-app/global"
	"awesomeProject/todo-app/handlers"
	"awesomeProject/todo-app/managers/implementations"
	"awesomeProject/todo-app/middlewares"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

func StartWebServer(config *configuration.Config, db *gorm.DB) {
	router := mux.NewRouter()

	DefineRoutes(router, db)

	log.Info("[StartWebServer] Starting web server...")
	err := http.ListenAndServe(":"+config.Port, router)
	if err != nil {
		log.Fatal("[StartWebServer] ListenAndServe: ", err)
	}
}

func DefineRoutes(router *mux.Router, db *gorm.DB) {
	log.Info("[DefineRoutes] Defining routes...")
	router.Use(middlewares.SetJSONMiddleware)

	authorHandler := handlers.NewAuthorHandler(implementations.NewAuthorManager(db))
	categoryHandler := handlers.NewCategoryHandler(implementations.NewCategoryManager(db))
	bookHandler := handlers.NewBookHandler(implementations.NewBookManager(db))
	loanHandler := handlers.NewLoanHandler(implementations.NewLoanManager(db))
	reservationHandler := handlers.NewReservationHandler(implementations.NewReservationManager(db))
	reviewHandler := handlers.NewReviewHandler(implementations.NewReviewManager(db))

	router.HandleFunc("/api/authors", authorHandler.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/api/authors", authorHandler.Create).Methods(http.MethodPost)
	router.HandleFunc("/api/authors/{id:"+global.UuidRegex+"}", authorHandler.Get).Methods(http.MethodGet)
	router.HandleFunc("/api/authors/{id:"+global.UuidRegex+"}", authorHandler.Update).Methods(http.MethodPut)
	router.HandleFunc("/api/authors/{id:"+global.UuidRegex+"}", authorHandler.Delete).Methods(http.MethodDelete)

	router.HandleFunc("/api/categories", categoryHandler.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/api/categories", categoryHandler.Create).Methods(http.MethodPost)
	router.HandleFunc("/api/categories/{id:"+global.UuidRegex+"}", categoryHandler.Get).Methods(http.MethodGet)
	router.HandleFunc("/api/categories/{id:"+global.UuidRegex+"}", categoryHandler.Update).Methods(http.MethodPut)
	router.HandleFunc("/api/categories/{id:"+global.UuidRegex+"}", categoryHandler.Delete).Methods(http.MethodDelete)

	router.HandleFunc("/api/books", bookHandler.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/api/books", bookHandler.Create).Methods(http.MethodPost)
	router.HandleFunc("/api/books/{id:"+global.UuidRegex+"}", bookHandler.Get).Methods(http.MethodGet)
	router.HandleFunc("/api/books/{id:"+global.UuidRegex+"}", bookHandler.Update).Methods(http.MethodPut)
	router.HandleFunc("/api/books/{id:"+global.UuidRegex+"}", bookHandler.Delete).Methods(http.MethodDelete)

	router.HandleFunc("/api/loans", loanHandler.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/api/loans", loanHandler.Create).Methods(http.MethodPost)
	router.HandleFunc("/api/loans/{id:"+global.UuidRegex+"}", loanHandler.Get).Methods(http.MethodGet)
	router.HandleFunc("/api/loans/{id:"+global.UuidRegex+"}", loanHandler.Update).Methods(http.MethodPut)
	router.HandleFunc("/api/loans/{id:"+global.UuidRegex+"}", loanHandler.Delete).Methods(http.MethodDelete)

	router.HandleFunc("/api/reservations", reservationHandler.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/api/reservations", reservationHandler.Create).Methods(http.MethodPost)
	router.HandleFunc("/api/reservations/{id:"+global.UuidRegex+"}", reservationHandler.Get).Methods(http.MethodGet)
	router.HandleFunc("/api/reservations/{id:"+global.UuidRegex+"}", reservationHandler.Update).Methods(http.MethodPut)
	router.HandleFunc("/api/reservations/{id:"+global.UuidRegex+"}", reservationHandler.Delete).Methods(http.MethodDelete)

	router.HandleFunc("/api/reviews", reviewHandler.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/api/reviews", reviewHandler.Create).Methods(http.MethodPost)
	router.HandleFunc("/api/reviews/{id:"+global.UuidRegex+"}", reviewHandler.Get).Methods(http.MethodGet)
	router.HandleFunc("/api/reviews/{id:"+global.UuidRegex+"}", reviewHandler.Update).Methods(http.MethodPut)
	router.HandleFunc("/api/reviews/{id:"+global.UuidRegex+"}", reviewHandler.Delete).Methods(http.MethodDelete)

	log.Info("[DefineRoutes] Defined routes")
}

func ConnectToDatabase(config *configuration.Config) *gorm.DB {
	log.Info("[ConnectToDatabase] Connecting to database...")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	log.Info("[ConnectToDatabase] Connected to database")
	return db
}

func main() {
	log.Info("[main] Starting app...")

	config, err := configuration.LoadConfig(".")
	if err != nil {
		log.Fatal("[main] Cannot load configuration: ", err)
	}

	db := ConnectToDatabase(config)

	StartWebServer(config, db)

	//err = db.AutoMigrate(&models.Author{})
	//if err != nil {
	//	log.Errorf("[main] Cannot migrate authors: %v", err)
	//}
	//newAuthor := models.Author{
	//	FirstName:   "Jane",
	//	LastName:    "Austen",
	//	Nationality: "British",
	//}
	//db.Create(&newAuthor)
}
