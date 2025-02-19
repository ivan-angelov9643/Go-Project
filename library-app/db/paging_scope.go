package db

import (
	"github.com/ivan-angelov9643/go-project/library-app/global"
	"gorm.io/gorm"
	"net/http"
)

type PagingScope struct {
	PageSize int
	Page     int
	Offset   int
}

func NewPagingScope(r *http.Request) *PagingScope {
	return &PagingScope{
		PageSize: global.GetPageSize(r),
		Page:     global.GetPage(r),
		Offset:   global.GetOffset(r),
	}
}

func (s *PagingScope) Get() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(s.Offset).Limit(s.PageSize)
	}
}
