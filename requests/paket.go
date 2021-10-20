package requests

func (FindPaket) TableName() string {
	return "paket"
}

type FindPaket struct {
	ID uint `json:"id"`
}
