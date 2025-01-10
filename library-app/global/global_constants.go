package global

import "time"

const (
	DefaultConfigurationFilePath = "./.env"
	ConfigFileName               = ".env"
	UuidRegex                    = "[a-f0-9]{8}(?:-[a-f0-9]{4}){3}-[a-f0-9]{12}"
	ReservationDuration          = 10 * time.Minute
	ReservationCleanupInterval   = 1 * time.Minute

	// Configuration
	DefaultPort        = "8080"
	DefaultLogFormat   = "text"
	DefaultLogSeverity = "debug"

	// Scopes
	PUBLIC = "public"
	READ   = "read"
	WRITE  = "write"
	GLOBAL = "global"

	// Resources
	Author      = "author"
	Category    = "category"
	Book        = "book"
	Loan        = "loan"
	Rating      = "rating"
	User        = "user"
	Reservation = "reservation"

	GLOBAL_SCOPE    = "global-scope"
	CURRENT_USER_ID = "current-user-id"
)
