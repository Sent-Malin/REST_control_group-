package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"gitlab.com/mediasoft_internship/gotasks/task1/sentmalin/internal/app/repository"
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

type Handler struct {
	Group
	Human
}

func NewHandler(repo *repository.Repository) *Handler {
	var handler Handler = Handler{
		Group: repo.Group,
		Human: repo.Human,
	}
	http.HandleFunc("/createGroup", handler.createGroup)
	http.HandleFunc("/createHuman", handler.createHuman)
	http.HandleFunc("/updateHuman", handler.updateHuman)
	http.HandleFunc("/updateGroup", handler.updateGroup)
	http.HandleFunc("/deleteHuman", handler.deleteHuman) //требуется параметр id в url
	http.HandleFunc("/deleteGroup", handler.deleteGroup) //требуется параметр id в url

	http.HandleFunc("/getGroups", handler.getGroups)
	http.HandleFunc("/getCountInGroup", handler.getCountInGroup)                   //требуется параметр id в url
	http.HandleFunc("/getCountInGroupWithChild", handler.getCountInGroupWithChild) //требуется параметр id в url

	http.HandleFunc("/getHumansGroup", handler.getHumansGroup)                        //требуется параметр id в url
	http.HandleFunc("/getHumansGroupWithChild", handler.getHumansGroupWithChildGroup) //требуется параметр id в url
	return &handler
}

func (s *Handler) ServerRun(port string) error {
	server := &http.Server{
		Addr: port,
	}
	err := server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func (s *Handler) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
