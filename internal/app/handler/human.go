package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"gitlab.com/mediasoft_internship/gotasks/task1/sentmalin/models"
)

//createHuman обрабатывает запрос создание записи человека
//в бд и выводит пользователю сообщение о результате
func (handler *Handler) createHuman(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		handler.respond(w, r, http.StatusInternalServerError, map[string]interface{}{
			"error-message": err.Error(),
		})
		return
	}

	var dataCreateHuman = models.HumanDataCreate{}

	err = json.Unmarshal(body, &dataCreateHuman)

	if err != nil {
		handler.respond(w, r, http.StatusInternalServerError, map[string]interface{}{
			"error-message": err.Error(),
		})
		return
	}
	//вызов функции обащения к базе данных
	err = handler.Human.Create(r.Context(), dataCreateHuman)
	if err != nil {
		handler.respond(w, r, http.StatusInternalServerError, map[string]interface{}{
			"error-message": err.Error(),
		})
		return
	}
	handler.respond(w, r, http.StatusInternalServerError, map[string]interface{}{
		"message": "Person record added",
	})
}

//getHumansGroup обрабатывает запрос на вывод информации
//о людях в конкретной группе и отправляет пользователю сообщение с результатом
func (handler *Handler) getHumansGroup(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		handler.respond(w, r, http.StatusBadRequest, map[string]interface{}{
			"error-message": "invalid param id",
		})
		return
	}
	result, err := handler.Human.GetHumansGroup(r.Context(), strconv.Itoa(id))
	if err != nil {
		handler.respond(w, r, http.StatusInternalServerError, map[string]interface{}{
			"error-message": err.Error(),
		})
		return
	}
	handler.respond(w, r, http.StatusOK, map[string]interface{}{
		"All human:": result,
	})
}

//getHumansGroupWithChildGroup обрабатывает запрос на вывод информации
//о людях в конкретной группе и дочерних группах и отправляет пользователю сообщение с результатом
func (handler *Handler) getHumansGroupWithChildGroup(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		handler.respond(w, r, http.StatusBadRequest, map[string]interface{}{
			"error-message": "invalid param id",
		})
		return
	}
	result, err := handler.Human.GetHumansGroupWithChildGroup(r.Context(), strconv.Itoa(id))
	if err != nil {
		handler.respond(w, r, http.StatusInternalServerError, map[string]interface{}{
			"error-message": err.Error(),
		})
		return
	}
	handler.respond(w, r, http.StatusOK, map[string]interface{}{
		"All human:": result,
	})
}

//updateHuman обрабатывает запрос обновление записи человекка
//в бд и выводит пользователю сообщение о результате
func (handler *Handler) updateHuman(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		handler.respond(w, r, http.StatusInternalServerError, map[string]interface{}{
			"error-message": err.Error(),
		})
		return
	}

	var dataUpdateHuman = models.HumanDataUpdate{}

	err = json.Unmarshal(body, &dataUpdateHuman)

	if err != nil {
		handler.respond(w, r, http.StatusInternalServerError, map[string]interface{}{
			"error-message": err.Error(),
		})
		return
	}
	//вызов функции обращения к бд
	err = handler.Human.Update(r.Context(), dataUpdateHuman)
	if err != nil {
		handler.respond(w, r, http.StatusInternalServerError, map[string]interface{}{
			"error-message": err.Error(),
		})
		return
	}
	handler.respond(w, r, http.StatusOK, map[string]interface{}{
		"message": "Person record update",
	})
}

//deleteHuman обрабатывает запрос на удаление записи
//человека и отправляет пользователю сообщение о результате
func (handler *Handler) deleteHuman(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		handler.respond(w, r, http.StatusBadRequest, map[string]interface{}{
			"error-message": "invalid param id",
		})
		return
	}
	//вызов функции обащения к базе данных
	err = handler.Human.Delete(r.Context(), id)
	if err != nil {
		handler.respond(w, r, http.StatusInternalServerError, map[string]interface{}{
			"error-message": err.Error(),
		})
		return
	}
	handler.respond(w, r, http.StatusOK, map[string]interface{}{
		"message": "Person record deleted",
	})
}
