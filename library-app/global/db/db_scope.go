package db

import "gorm.io/gorm"

type DBScope struct {
	Global bool
	UserID string
}

func NewDBScope(global bool, userID string) *DBScope {
	return &DBScope{global, userID}
}

func (dbs *DBScope) Exec(db *gorm.DB) *gorm.DB {
	if dbs.Global {
		return db
	} else {
		return db.Where("user_id = ?", dbs.UserID)
	}
}
