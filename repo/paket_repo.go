package repo

import (
	"github.com/mohamadakbar/go-invitee/models"
)

type Paket interface {
	InqGetById(id uint) (models.Pakets, int, error)
}
