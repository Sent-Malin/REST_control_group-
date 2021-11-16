package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gitlab.com/mediasoft_internship/gotasks/task1/sentmalin/models"
)

type HumanPQ struct {
	db *sqlx.DB
}

func NewHumanPQ(db *sqlx.DB) *HumanPQ {
	return &HumanPQ{db: db}
}

//Create получает контекст и данные о человеке,
//создает запись в бд и возвращает данные об операции
func (repository *HumanPQ) Create(ctx context.Context, data models.HumanDataCreate) error {
	//проверяем есть ли в бд группа, в которую стоит зачислить человека по id
	res, err := repository.db.QueryContext(ctx, `select id from groups 
	where id=$1`, data.GroupNumber)
	if err != nil {
		return err
	}
	//Если результат запроса содержит строку, добавляем запись
	if res.Next() {

		_, err = repository.db.ExecContext(ctx, `insert into humans ("name", "surname", "year_of_birth", "group_id") 
		values($1, $2, $3, $4)`, data.Name, data.Surname, data.YearOfBirth, data.GroupNumber)

		if err != nil {
			return err
		}
		return nil
	}
	return err
}

//GetHumansGroup получает контекст и строку с числом, возвращает
//данные всех людей из соответствующей числу группы
func (repository *HumanPQ) GetHumansGroup(ctx context.Context, id string) ([]models.HumanDataUpdate, error) {
	res, err := repository.db.QueryContext(ctx, `select * from humans 
	where group_id=$1`, id)
	if err != nil {
		return nil, err
	}
	var arrdata = []models.HumanDataUpdate{}
	//переменная для хранения года рождения, который может быть нулевым
	var t *string
	for res.Next() {
		p := models.HumanDataUpdate{}
		//считываем из каждой строки результата запроса id, имя, фамилию, год рождения и номер группы человека
		err = res.Scan(&p.Id, &p.Name, &p.Surname, &t, &p.GroupNumber)
		if t != nil {
			p.YearOfBirth = *t
		}
		if err != nil {
			continue
		}
		//добавляем человека в итоговый массив
		arrdata = append(arrdata, p)
	}
	return arrdata, err
}

//GetHumansGroupWithChildGroup получает контекст и строку с числом, возвращает
//данные всех людей из соответствующей числу группы и из дочерних групп
func (repository *HumanPQ) GetHumansGroupWithChildGroup(ctx context.Context, id string) ([]models.HumanDataUpdate, error) {
	res, err := repository.db.QueryContext(ctx, `select * from humans 
	where group_id=any(select id from groups 
		where id=$1 or parent_id=$2)`, id, id)
	if err != nil {
		return nil, err
	}
	var arrdata = []models.HumanDataUpdate{}
	//переменная для хранения года рождения, который может быть нулевым
	var t *string
	for res.Next() {
		p := models.HumanDataUpdate{}
		//считываем из каждой строки результата запроса id, имя, фамилию, год рождения и номер группы человека
		err = res.Scan(&p.Id, &p.Name, &p.Surname, &t, &p.GroupNumber)
		if t != nil {
			p.YearOfBirth = *t
		}
		if err != nil {
			continue
		}
		//добавляем человека в итоговый массив
		arrdata = append(arrdata, p)
	}
	return arrdata, err
}

//Update получает контекст и даные человека, обновляет
//запись человека в базе данных и возвращает данные об операции
func (repository *HumanPQ) Update(ctx context.Context, data models.HumanDataUpdate) error {
	//проверяем есть ли в бд человек с редактируемым ид
	res1, err := repository.db.QueryContext(ctx, `select id from humans 
	where id=$1`, data.Id)
	if err != nil {
		return err
	}
	//проверяем есть ли в бд группа, введенная в поле "номер группы"
	res2, err := repository.db.QueryContext(ctx, `select id from groups 
	where id=$1`, data.GroupNumber)
	if err != nil {
		return err
	}
	//Если человек существует, и редактирумый номер группы приемлемый, обновляем запись
	if res1.Next() && res2.Next() {

		_, err = repository.db.ExecContext(ctx, `update humans 
		set name=$1, surname=$2, year_of_birth=$3, group_id=$4 
		where id=$5`, data.Name, data.Surname, data.YearOfBirth, data.GroupNumber, data.Id)

		if err != nil {
			return err
		}
		return nil
	}
	return err
}

//Delete получает контекст и число, удаляет соответствующую
//запись человека в базе данных и возвращает данные об операции
func (repository *HumanPQ) Delete(ctx context.Context, id int) error {
	//проверяем есть ли в бд человек с таким ид
	res, err := repository.db.QueryContext(ctx, `select id from humans 
	where id=$1`, id)
	if err != nil {
		return err
	}
	//Если результат запроса содержит строку, удаляем запись
	if res.Next() {
		_, err = repository.db.ExecContext(ctx, `delete from humans 
		where id=$1`, id)

		if err != nil {
			return err
		}
		return nil
	}
	return err
}
