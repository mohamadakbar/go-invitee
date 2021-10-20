package repo

import (
	"github.com/mohamadakbar/go-invitee/models"
	"github.com/mohamadakbar/go-invitee/requests"
)

type User interface {
	InqGetBySlug(slug string) (models.Users, int, error)
	InqStore(user requests.StoreUser, paket models.Pakets) (models.Users, int, error)
	InqGetByEmail(email string) (models.Users, int, error)
}
