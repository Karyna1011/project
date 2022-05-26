package postgres

import (
	"database/sql"
	"gitlab.com/tokend/subgroup/project/internal/data"

	"github.com/fatih/structs"

	"github.com/Masterminds/squirrel"

	"gitlab.com/distributed_lab/kit/pgdb"
)

const tablePerson = "person"

type personQ struct {
	db  *pgdb.DB
	sql squirrel.SelectBuilder
}

func NewPersonQ(db *pgdb.DB) data.PersonQ {
	return &personQ{
		db:  db.Clone(),
		sql: squirrel.Select("*").From(tablePerson),
	}
}

func (d *personQ) New() data.PersonQ {
	return NewPersonQ(d.db)
}

func (d *personQ) Get() (*data.Person, error) {
	var result data.Person

	err := d.db.Get(&result, d.sql)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &result, err
}

func (d *personQ) Select(query pgdb.OffsetPageParams) ([]data.Person, error) {
	var result []data.Person

	err := d.db.Select(&result, query.ApplyTo(d.sql, "id"))
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return result, err
}

func (d *personQ) Insert(person data.Person) (data.Person, error) {
	clauses := structs.Map(person)

	query := squirrel.Insert(tablePerson).SetMap(clauses).Suffix("returning *")

	err := d.db.Get(&person, query)
	if err != nil {
		return data.Person{}, err
	}

	return person, err

}

func (d *personQ) Update(person data.Person) (data.Person, error) {
	clauses := structs.Map(person)

	query := squirrel.Update(tablePerson).Where(squirrel.Eq{"id": person.Id}).SetMap(clauses).Suffix("returning *")

	err := d.db.Get(&person, query)
	if err != nil {
		return data.Person{}, err
	}

	return person, err
}

func (d personQ) FilterById(Id int64) data.PersonQ {
	d.sql = d.sql.Where(squirrel.Eq{"id": Id})

	return &d
}

func (d personQ) FilterByName(name string) data.PersonQ {
	d.sql = d.sql.Where(squirrel.Eq{"name": name})

	return &d
}

func (d *personQ) Delete(name string) error {
	query := squirrel.Delete(tablePerson).Where(squirrel.Eq{"name": name})
	err := d.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
