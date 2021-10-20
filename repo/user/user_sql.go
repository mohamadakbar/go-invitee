package user

import (
	"github.com/jinzhu/gorm"
	"github.com/mohamadakbar/go-invitee/models"
	repo "github.com/mohamadakbar/go-invitee/repo"
	"github.com/mohamadakbar/go-invitee/requests"
	"log"
	"net/http"
)

type sqlUserRepo struct {
	Conn *gorm.DB
}

func NewSQLUserRepo(conn *gorm.DB) repo.User {
	return &sqlUserRepo{
		Conn: conn,
	}
}

func (p *sqlUserRepo) InqGetBySlug(slug string) (models.Users, int, error) {
	var user models.Users
	query := p.Conn

	err := query.Where("slug = ? ", slug).Find(&user).Error
	if err != nil {
		//log.Println("cannot fetch Slug : ",err)
		return models.Users{}, http.StatusForbidden, err
	}

	return user, http.StatusOK, nil
}

func (p *sqlUserRepo) InqGetByEmail(email string) (models.Users, int, error) {
	var user models.Users
	query := p.Conn

	err := query.Where("email = ? ", email).Find(&user).Error
	if err != nil {
		//log.Println("cannot fetch email : ",err)
		return models.Users{}, http.StatusForbidden, err
	}

	return user, http.StatusOK, nil
}

func (p *sqlUserRepo) InqStore(req requests.StoreUser, pakets models.Pakets) (models.Users, int, error) {
	//var paket models.Pakets
	//fmt.Println(pakets)
	//os.Exit(1)
	//idpkt := pakets.ID
	query := p.Conn

	user := models.Users{
		Slug:  req.Slug,
		Name:  req.Name,
		IsNew: req.IsNew,
		Paket: pakets,
		Email: req.Email,
	}

	if err := query.Create(&user).Error; err != nil {
		log.Fatal("error : ", err.Error)
		return models.Users{}, http.StatusInternalServerError, nil
	}

	return user, http.StatusOK, nil
}
