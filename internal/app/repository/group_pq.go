package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"gitlab.com/mediasoft_internship/gotasks/task1/sentmalin/models"
)

type GroupPQ struct {
	db *sqlx.DB
}

func NewGroupPQ(db *sqlx.DB) *GroupPQ {
	return &GroupPQ{db: db}
}

//Create получает контекст и данные о группе, создает запись группы в бд
//и возвращает данные об операции
func (repository *GroupPQ) Create(ctx context.Context, data models.GroupDataCreate) error {

	_, err := repository.db.ExecContext(ctx, `insert into groups ("title", "parent_id") 
	values($1, $2)`, data.Name, data.ParentId)

	if err != nil {
		return err
	}
	return nil
}

//GetCountHuman получает контекст и строку с числом, возвращает
//количество людей в соответствующей группе
func (repository *GroupPQ) GetCountHuman(ctx context.Context, id string) (string, error) {
	res, err := repository.db.QueryContext(ctx, `select count(*) from humans 
	where group_id=$1`, id)
	if err != nil {
		return "", err
	}
	var strForCountQuery string
	//если результат запроса есть
	if res.Next() {
		//считываем значение в строку
		res.Scan(&strForCountQuery)
		return strForCountQuery, nil
	}
	return "", err
}

//GetCountHumanWithChildGroup получает контекст и строку с числом, возвращает
//количество людей в соответствующей группе и в дочерних группах при помощи рекурсивного запроса
func (repository *GroupPQ) GetCountHumanWithChildGroup(ctx context.Context, id string) (string, error) {
	res, err := repository.db.QueryContext(ctx, `with recursive searcher(id, parent_id) as (
		select id, parent_id 
		from groups where id=$1
		union all
		select g.id, g.parent_id 
		from groups g 
		join searcher s ON s.id=g.parent_id
	)
	select count(*) from humans 
	where group_id=any(select id from searcher)`, id)
	if err != nil {
		return "", err
	}
	var strForCountQuery string
	//если результат запроса есть
	if res.Next() {
		//считываем значение в строку
		res.Scan(&strForCountQuery)
		return strForCountQuery, nil
	}
	return "", err
}

//GetAll получает контекст и возвращает данные
//всех групп
func (repository *GroupPQ) GetAll(ctx context.Context) ([]models.GroupDataUpdate, error) {
	res, err := repository.db.QueryContext(ctx, `select * from groups`)
	if err != nil {
		return nil, err
	}
	var arrdata = []models.GroupDataUpdate{}
	//переменная для хранения parentid, который может быть nil
	var t *int
	//для всех строк в результате запроса
	for res.Next() {
		p := models.GroupDataUpdate{}
		//считываем id группы, название и parentid
		err = res.Scan(&p.Id, &p.Name, &t)
		//если у группы есть родитель, разыменовываем указатель, иначе оставляем nil
		if t != nil {
			p.ParentId = *t
		}
		if err != nil {
			continue
		}
		//добавляем считанные данные группы в итоговый массив
		arrdata = append(arrdata, p)
	}
	return arrdata, err
}

//Update получает контекст и даные группы, обновляет запись
//группы в базе данных и возвращает данные об операции
func (repository *GroupPQ) Update(ctx context.Context, data models.GroupDataUpdate) error {
	_, err := repository.db.ExecContext(ctx, `update groups 
	set title=$1, parent_id=$2 
	where id=$3`, data.Name, data.ParentId, data.Id)
	if err != nil {
		return err
	}
	return nil
}

//Delete получает контекст и число, удаляет соответствующую
//запись группы в базе данных и возвращает данные об операции
func (repository *GroupPQ) Delete(ctx context.Context, id int) error {
	//проверяем есть ли в бд группа с ид, подлежащим удалению
	res, err := repository.db.QueryContext(ctx, `select id from humans 
	where id=$1`, id)
	if err != nil {
		return err
	}
	//Если результат запроса содержит строку, удаляем запись
	if res.Next() {
		//очищаем удаляемую группу и дочерние группы от людей

		_, err = repository.db.ExecContext(ctx, `update humans 
		set group_id=$1 
		where group_id=any(select id from groups 
			where id=$2 or parent_id=$3)`, 1, id, id)

		if err != nil {
			return err
		}
		//удаляем группу и дочерние группы
		if _, err = repository.db.ExecContext(ctx, `delete from groups 
		where id=$1 OR parent_id=$2`, id, id); err != nil {
			return err
		}
		return nil
	}
	return err
}
