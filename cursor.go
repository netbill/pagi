package pagi

type SortField struct {
	Field  string
	Ascend bool // "asc" or "desc"
}

type Params struct {
	Cursor map[string]string
	Sort   *SortField
	Limit  int
}

type Page[T any] struct {
	Data       []T
	NextCursor map[string]string
	Total      int
}
