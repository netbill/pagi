package pagi

import (
	"net/http"
	"strconv"
	"strings"
)

func GetLimit(r *http.Request) int {
	limit := 20

	if s := r.URL.Query().Get("limit"); s != "" {
		n, err := strconv.Atoi(s)
		if err == nil {
			limit = n
		}
	}

	if limit <= 0 {
		limit = 20
	}
	if limit > 100 {
		limit = 100
	}

	return limit
}

func GetSort(r *http.Request) *SortField {
	sortStr := strings.TrimSpace(r.URL.Query().Get("sort"))
	if sortStr == "" {
		return nil
	}

	sort := &SortField{
		Ascend: true,
	}

	if strings.HasPrefix(sortStr, "-") {
		sort.Ascend = false
		sort.Field = strings.TrimPrefix(sortStr, "-")
	} else {
		sort.Field = sortStr
	}

	sort.Field = strings.TrimSpace(sort.Field)
	if sort.Field == "" {
		return nil
	}

	return sort
}

func GetCursor(r *http.Request) *Cursor {
	s := strings.TrimSpace(r.URL.Query().Get("cursor"))
	if s == "" {
		return nil
	}
	return &Cursor{Value: s}
}

func GetParams(r *http.Request) Params {
	return Params{
		Cursor: GetCursor(r),
		Sort:   GetSort(r),
		Limit:  GetLimit(r),
	}
}
