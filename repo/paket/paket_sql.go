package paket

import (
	"github.com/jinzhu/gorm"
	"github.com/mohamadakbar/go-invitee/models"
	repo "github.com/mohamadakbar/go-invitee/repo"
	"net/http"
)

type sqlPaketRepo struct {
	Conn *gorm.DB
}

func NewSQLPaketRepo(conn *gorm.DB) repo.Paket {
	return &sqlPaketRepo{
		Conn: conn,
	}
}

func (p *sqlPaketRepo) InqGetById(id uint) (models.Pakets, int, error) {
	var paket models.Pakets
	query := p.Conn

	err := query.Where("id = ? ", id).Find(&paket).Error
	if err != nil {
		//log.Println("cannot fetch Slug : ",err)
		return models.Pakets{}, http.StatusNotFound, err
	}

	return paket, http.StatusOK, nil
}
