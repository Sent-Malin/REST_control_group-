package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"gitlab.com/mediasoft_internship/gotasks/task1/sentmalin/models"
)

//createGroup обрабатывает запрос создание записи группы
//в бд и выводит пользователю сообщение о результате
func (handler *Handler) createGroup(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		handler.respond(w, r, http.StatusInternalServerError, map[string]interface{}{
			"error-message": err.Error(),
		})
		return
	}

	var dataCreateGroup = models.GroupDataCreate{}

	err = json.Unmarshal(body, &dataCreateGroup)

	if err != nil {
		handler.respond(w, r, http.StatusInternalServerError, map[string]interface{}{
			"error-message": err.Error(),
		})
		return
	}

	err = handler.Group.Create(r.Context(), dataCreateGroup)

	if err != nil {
		handler.respond(w, r, http.StatusInternalServerError, map[string]interface{}{
			"error-message": err.Error(),
		})
		return
	}
	handler.respond(w, r, http.StatusOK, map[string]interface{}{
		"message": "Group record added",
	})
}

//getGroups обрабатывает запрос на вывод информации
//о всех группах и отправляет пользователю сообщение с результатом
func (handler *Handler) getGroups(w http.ResponseWriter, r *http.Request) {
	//вызов функции обращения к бд
	result, err := handler.Group.GetAll(r.Context())
	if err != nil {
		handler.respond(w, r, http.StatusInternalServerError, map[string]interface{}{
			"error-message": err.Error(),
		})
		return
	}
	handler.respond(w, r, http.StatusOK, map[string]interface{}{
		"All group:": result,
	})
}

//getCountInGroup обрабатывает запрос на вывод информации
//о количестве людей в конкретной группе и отправляет пользователю сообщение с результатом
func (handler *Handler) getCountInGroup(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		handler.respond(w, r, http.StatusBadRequest, map[string]interface{}{
			"error-message": "invalid param id",
		})
		return
	}
	//вызов функции обращения к бд
	result, err := handler.Group.GetCountHuman(r.Context(), strconv.Itoa(id))
	fmt.Println("fff")
	if err != nil {
		handler.respond(w, r, http.StatusInternalServerError, map[string]interface{}{
			"error-message": err.Error(),
		})
		return
	}
	if result != "" {
		handler.respond(w, r, http.StatusOK, map[string]interface{}{
			"Count humans in group": result,
		})
	} else {
		handler.respond(w, r, http.StatusInternalServerError, map[string]interface{}{
			"message": "not query result",
		})
	}
}

//getCountInGroupWithChild обрабатывает запрос на вывод информации
//о количестве людей в конкретной группе и в дочерних группах
//и отправляет пользователю сообщение с результатом
func (handler *Handler) getCountInGroupWithChild(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		handler.respond(w, r, http.StatusBadRequest, map[string]interface{}{
			"error-message": "invalid param id",
		})
		return
	}
	result, err := handler.Group.GetCountHumanWithChildGroup(r.Context(), strconv.Itoa(id))
	if err != nil {
		handler.respond(w, r, http.StatusInternalServerError, map[string]interface{}{
			"error-message": err.Error(),
		})
		return
	}
	if result != "" {
		handler.respond(w, r, http.StatusOK, map[string]interface{}{
			"Count humans in group with child groups": result,
		})
	} else {
		handler.respond(w, r, http.StatusInternalServerError, map[string]interface{}{
			"message": "not query result",
		})
	}
}

//updateGroup обрабатывает запрос обновление записи группы
//в бд и выводит пользователю сообщение о результате
func (handler *Handler) updateGroup(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		handler.respond(w, r, http.StatusInternalServerError, map[string]interface{}{
			"error-message": err.Error(),
		})
		return
	}

	var dataForUpdateGroup = models.GroupDataUpdate{}

	err = json.Unmarshal(body, &dataForUpdateGroup)

	if err != nil {
		handler.respond(w, r, http.StatusInternalServerError, map[string]interface{}{
			"error-message": err.Error(),
		})
		return
	}
	//вызов функции обращения к бд
	err = handler.Group.Update(r.Context(), dataForUpdateGroup)
	if err != nil {
		handler.respond(w, r, http.StatusInternalServerError, map[string]interface{}{
			"error-message": err.Error(),
		})
		return
	}
	handler.respond(w, r, http.StatusOK, map[string]interface{}{
		"message": "Group record update",
	})
}

//deleteGroup обрабатывает запрос на удаление записи
//группы и отправляет пользователю сообщение о результате
func (handler *Handler) deleteGroup(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		handler.respond(w, r, http.StatusBadRequest, map[string]interface{}{
			"error-message": "invalid param id",
		})
		return
	}
	//вызов функции обращения к базе данных
	err = handler.Group.Delete(r.Context(), id)
	if err != nil {
		handler.respond(w, r, http.StatusInternalServerError, map[string]interface{}{
			"error-message": err.Error(),
		})
		return
	}
	handler.respond(w, r, http.StatusOK, map[string]interface{}{
		"message": "Group record deleted",
	})
}
