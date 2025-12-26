package pagi

type Params struct {
	Cursor map[string]string
	Limit  uint
}

type Page[T any] struct {
	Data   T
	Cursor map[string]string
	Total  uint
}
