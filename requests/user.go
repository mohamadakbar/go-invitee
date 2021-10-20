package requests

func (StoreUser) TableName() string {
	return "users"
}

type StoreUser struct {
	Slug    string `json:"slug"`
	Name    string `json:"name"`
	IsNew   uint   `json:"is_new"`
	IDPaket uint   `json:"id_paket"`
	Email   string `json:"email"`
}
