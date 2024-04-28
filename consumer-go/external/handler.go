package external

type Handler interface {
	PostParcel(data *Request) ([]byte, error)
}
