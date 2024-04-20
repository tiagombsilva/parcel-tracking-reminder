package handler

type Handler interface {
	PostParcel(data *Request) ([]byte, error)
}
