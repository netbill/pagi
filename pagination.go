package pagi

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetPagination(r *http.Request) (page, size uint) {
	pageStr := chi.URLParam(r, "page")
	if pageStr != "" {
		n, err := strconv.Atoi(pageStr)
		if err != nil {
			page = 1
		}
		page = uint(n)
	}

	sizeStr := chi.URLParam(r, "size")
	if sizeStr != "" {
		n, err := strconv.Atoi(sizeStr)
		if err != nil {
			size = 20
		}
		size = uint(n)
	}

	return page, size
}

func PagConvert(page, size uint) (limit, offset uint) {
	if page == 0 {
		page = 1
	}
	if size > 100 {
		size = 100
	}
	if size == 0 {
		size = 20
	}

	limit = uint(size)
	offset = uint((page - 1) * size)
	return
}
