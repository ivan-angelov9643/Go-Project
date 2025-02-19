package db

import (
	"github.com/ivan-angelov9643/go-project/library-app/global"
	"gorm.io/gorm"
	"net/http"
)

type AccessScope struct {
	Global bool
	UserID string
}

func NewAccessScope(r *http.Request) *AccessScope {
	return &AccessScope{global.IsGlobal(r), global.GetOwnerID(r)}
}

func (as *AccessScope) Get() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if as.Global {
			return db
		} else {
			return db.Where("user_id = ?", as.UserID)
		}
	}
}
