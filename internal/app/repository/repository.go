package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"gitlab.com/mediasoft_internship/gotasks/task1/sentmalin/models"
)

type Group interface {
	Create(ctx context.Context, data models.GroupDataCreate) error
	GetCountHuman(ctx context.Context, id string) (string, error)
	GetCountHumanWithChildGroup(ctx context.Context, id string) (string, error)
	GetAll(ctx context.Context) ([]models.GroupDataUpdate, error)
	Update(ctx context.Context, data models.GroupDataUpdate) error
	Delete(ctx context.Context, id int) error
}

type Human interface {
	Create(ctx context.Context, data models.HumanDataCreate) error
	GetHumansGroup(ctx context.Context, id string) ([]models.HumanDataUpdate, error)
	GetHumansGroupWithChildGroup(ctx context.Context, id string) ([]models.HumanDataUpdate, error)
	Update(ctx context.Context, data models.HumanDataUpdate) error
	Delete(ctx context.Context, id int) error
}

type Repository struct {
	Group
	Human
}

func NewRepos(config *models.Config) (*Repository, error) {
	db, er := newConnect(config)
	if er != nil {
		return nil, er
	}
	return &Repository{
		Group: NewGroupPQ(db),
		Human: NewHumanPQ(db),
	}, nil
}

//newConnect открывает соединение с базой данных
func newConnect(config *models.Config) (*sqlx.DB, error) {
	psqlconn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", config.User, config.Password, config.Dbname)
	db, err := sqlx.Open("postgres", psqlconn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
