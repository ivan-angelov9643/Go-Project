package global

const (
	DefaultConfigurationFilePath = "./.env"
	ConfigFileName               = ".env"
	UuidRegex                    = "[a-f0-9]{8}(?:-[a-f0-9]{4}){3}-[a-f0-9]{12}"

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
	Review      = "review"
	User        = "user"
	Reservation = "reservation"

	GLOBAL_SCOPE    = "global-scope"
	CURRENT_USER_ID = "current-user-id"
)
