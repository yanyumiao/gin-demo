package model

import db "gin-demo/database"

type Person struct {
	Id        int    `json:"id" form:"id"`
	FirstName string `json:"firstname" form:"firstname"`
	LastName  string `json:"lastname" form:"lastname"`
}

func (p *Person) Add() (id int64, err error) {
	rs, err := db.My.Exec("INSERT INTO person(firstname, lastname) VALUES (?, ?)", p.FirstName, p.LastName)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}

func (p *Person) Get() (person Person, err error) {
	// http://go-database-sql.org/retrieving.html
	row := db.My.QueryRow("SELECT id, firstname, lastname FROM person WHERE id = ?", p.Id)
	err = row.Scan(&person.Id, &person.FirstName, &person.LastName)
	if err != nil {
		return
	}
	return
}

func (p *Person) GetAll() (persons []Person, err error) {
	persons = make([]Person, 0)
	rows, err := db.My.Query("SELECT id, firstname, lastname FROM person")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var person Person
		rows.Scan(&person.Id, &person.FirstName, &person.LastName)
		persons = append(persons, person)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}

func (p *Person) Del() (ra int64, err error) {
	rs, err := db.My.Exec("DELETE FROM person WHERE id=?", p.Id)
	if err != nil {
		return
	}
	ra, err = rs.RowsAffected()
	return
}

func (p *Person) Update() (ra int64, err error) {
	rs, err := db.My.Exec("UPDATE person SET firstname=?, lastname=? WHERE id=?", p.FirstName, p.LastName, p.Id)
	if err != nil {
		return
	}
	ra, err = rs.RowsAffected()
	return
}
