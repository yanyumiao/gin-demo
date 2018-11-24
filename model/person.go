package model

import (
	db "gin/database"
	//"log"
)

type Person struct {
	Id        int    `json:"id" form:"id"`
	FirstName string `json:"firstname" form:"firstname"`
	LastName  string `json:"lastname" form:"lastname"`
}

// author: YanYuMiao
func (p *Person) Get() (person Person) {
	// TODO 翻译总结 go-database-sql.org
	// http://go-database-sql.org/retrieving.html
	row := db.SqlDB.QueryRow("SELECT id, firstname, lastname FROM person WHERE id = ?", p.Id)
	row.Scan(&person.Id, &person.FirstName, &person.LastName)
	return
}

// int64 ? even use int
func (p *Person) Add() (id int64, err error) {
	rs, err := db.SqlDB.Exec("INSERT INTO person(firstname, lastname) VALUES (?, ?)", p.FirstName, p.LastName)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}

func (p *Person) GetAll() (persons []Person, err error) {
	persons = make([]Person, 0)
	rows, err := db.SqlDB.Query("SELECT id, firstname, lastname FROM person")
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
