package global

import (
	"net/http"
	"strconv"
)

func IntPtr(i int) *int {
	return &i
}

func IsGlobal(request *http.Request) bool {
	return request.Context().Value(GLOBAL_SCOPE) != nil
}

func GetOwnerID(request *http.Request) string {
	return request.Context().Value(CURRENT_USER_ID).(string)
}

func GetPageSize(r *http.Request) int {
	query := r.URL.Query()
	pageSize, _ := strconv.Atoi(query.Get("page_size")) // Error is ignored because wrong or missing parameters are handled as 0
	switch {
	case pageSize > 500:
		pageSize = 500
	case pageSize <= 0:
		pageSize = 10
	}
	return pageSize
}

func GetPage(r *http.Request) int {
	query := r.URL.Query()
	page, _ := strconv.Atoi(query.Get("page")) // Error is ignored because wrong or missing parameters are handled as 0
	if page <= 0 {
		page = 1
	}
	return page
}

func GetOffset(r *http.Request) int {
	return (GetPage(r) - 1) * GetPageSize(r)
}
