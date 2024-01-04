package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "amonzark"
	password = "nuansa12"
	dbname   = "postgres"
)

// connection to DB
func connect() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to DB")
	return db, nil
}

// listing table of db
func ListStudentsHandler() []Students {
	db, err := connect()

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM tb_student")

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil

	}
	student := []Students{}
	for results.Next() {
		var std Students
		err = results.Scan(&std.ID, &std.Name, &std.Age, &std.Grade)

		if err != nil {
			panic(err.Error())
		}
		student = append(student, std)
	}
	return student
}

// create new student in table
func CreateStudentsHandler(student Students) {
	db, err := connect()
	if err != nil {
		fmt.Println("Err", err.Error())
	}

	defer db.Close()
	insert, err := db.Query(
		"insert into tb_student (id, name, age, grade) values ($1, $2, $3, $4)",
		student.ID, student.Name, student.Age, student.Grade)

	if err != nil {
		fmt.Println("Err", err.Error())
	}
	defer insert.Close()
}

// get student details from db by id
func GetStudentsbyID(id string) *Students {
	db, err := connect()
	std := &Students{}
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}
	defer db.Close()

	results, err := db.Query("SELECT * FROM tb_student where id=$1", id)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	if results.Next() {
		err = results.Scan(&std.ID, &std.Name, &std.Age, &std.Grade)
		if err != nil {
			return nil
		}
	} else {
		return nil
	}
	return std
}

// get student details by age
func GetStudentsbyAge(age string) *Students {
	db, err := connect()
	std := &Students{}
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}
	defer db.Close()

	results, err := db.Query("SELECT * FROM tb_student where age=$1", age)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	if results.Next() {
		err = results.Scan(&std.ID, &std.Name, &std.Age, &std.Grade)
		if err != nil {
			return nil
		}
	} else {
		return nil
	}
	return std
}

// update student age by id
func UpdateStudentsbyID(student Students) {
	db, err := connect()
	if err != nil {
		fmt.Println("Err", err.Error())
	}

	defer db.Close()
	insert, err := db.Query(
		"update tb_student set age = $1 where id = $2",
		student.Age, student.ID)

	if err != nil {
		fmt.Println("Err", err.Error())
	}
	defer insert.Close()
}

// delete student by id
func DeleteStudentsbyID(id string) *Students {
	db, err := connect()
	std := &Students{}
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}
	defer db.Close()

	results, err := db.Query("delete from tb_student where id = $1", id)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	if results.Next() {
		err = results.Scan(&std.ID, &std.Name, &std.Age, &std.Grade)
		if err != nil {
			return nil
		}
	} else {
		return nil
	}
	return std
}
