package global

import "time"

const (
	DefaultConfigurationFilePath = "./.env"
	UuidRegex                    = "[a-f0-9]{8}(?:-[a-f0-9]{4}){3}-[a-f0-9]{12}"
	ReservationDuration          = 10 * time.Minute
	ReservationCleanupInterval   = 1 * time.Minute

	// Configuration
	DefaultPort        = "8080"
	DefaultLogFormat   = "text"
	DefaultLogSeverity = "debug"

	// Scopes
	GLOBAL = "global"
	READ   = "read"
	CREATE = "create"
	EDIT   = "edit"
	DELETE = "delete"

	// Resources
	Author      = "author"
	Category    = "category"
	Book        = "book"
	Loan        = "loan"
	Rating      = "rating"
	User        = "user"
	Reservation = "reservation"

	GLOBAL_SCOPE    = "global_scope"
	CURRENT_USER_ID = "current_user_id"
)
