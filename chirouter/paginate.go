package chirouter

import (
	"context"
	"net/http"
	"strconv"
)

type Pagination struct {
	Page   int
	Limit  int
	Offset int
	Total  int
}

func Paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pageStr := r.URL.Query().Get("page")
		page, err := strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			page = 1
		}

		limitStr := r.URL.Query().Get("limit")
		limit, err := strconv.Atoi(limitStr)
		if err != nil || limit < 1 {
			limit = 20
		}
		if limit > 100 {
			limit = 100
		}

		offset := (page - 1) * limit

		pagination := &Pagination{
			Page:   page,
			Limit:  limit,
			Offset: offset,
		}

		ctx := context.WithValue(r.Context(), "pagination", pagination)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetPagination(r *http.Request) *Pagination {
	if p, ok := r.Context().Value("pagination").(*Pagination); ok {
		return p
	}
	return nil
}
