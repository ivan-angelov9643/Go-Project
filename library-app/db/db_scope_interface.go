package db

import "gorm.io/gorm"

type DBScope interface {
	Get() func(db *gorm.DB) *gorm.DB
}

func ApplyScopes(db *gorm.DB, scopes []DBScope) *gorm.DB {
	for _, scope := range scopes {
		db = db.Scopes(scope.Get())
	}
	return db
}
