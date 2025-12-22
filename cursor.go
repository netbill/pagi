package pagi

import (
	"strconv"

	"github.com/google/uuid"
)

type Cursor struct {
	Value string
}

func (c *Cursor) ParseToUUID() (uuid.UUID, error) {
	return uuid.Parse(c.Value)
}

func (c *Cursor) ParseToString() string {
	return c.Value
}

func (c *Cursor) ParseToInt() (int, error) {
	return strconv.Atoi(c.Value)
}

func (c *Cursor) ParseToUint() (uint, error) {
	parsed, err := strconv.ParseUint(c.Value, 10, 64)
	if err != nil {
		return 0, err
	}

	return uint(parsed), nil
}

type SortField struct {
	Field  string
	Ascend bool // "asc" or "desc"
}

type Params struct {
	Cursor *Cursor
	Sort   *SortField
	Limit  int
}

type Page[T any] struct {
	Data       []T
	NextCursor *Cursor
	Total      int
}
