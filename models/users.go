package models

type Users struct {
	ID		uint 		`json:"id" gorm:"primary_key"`
	Slug	string		`json:"slug"`
	Name 	string		`json:"name"`
	IsNew	int 		`json:"is_new"`
	IDPaket int			`json:"id_paket"`
	Email	string		`json:"email"`
}
