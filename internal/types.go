package internal

type Data[T any] struct {
	Data T
}

type IController interface {
	GetList(result string, limit int, offset int) (any, error)
	GetSpecific(result string, search string) (any, error)
}
