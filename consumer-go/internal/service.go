package internal

type Service interface {
	GetAllParcels() []Parcel
}

type Parcel struct {
	ID         int64
	LastUpdate string
	Name       string
	State      string
	Tracking   string
	AccountId  int64
}
