package models

import (
	"echo-user-app/db"
	"log"
	"net/http"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func FetchUsers() (Response, error) {
	var obj User
	var arrobj []User
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM users"

	rows, err := con.Query(sqlStatement)
	if err != nil {
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.ID, &obj.Name, &obj.Email, &obj.Age)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}

func FetchUser(id int) (Response, error) {
	var obj User
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM users where user_id=$1"

	err := con.QueryRow(sqlStatement, id).Scan(&obj.ID, &obj.Name, &obj.Email, &obj.Age)
	if err != nil {
		return res, err
	}

	// defer rows.Close()

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = obj

	return res, err
}

func StoreUser(name string, email string, age int) (Response, error) {
	var res Response
	var id int

	con := db.CreateCon()

	sqlStatement := "INSERT INTO users (name, email, age) VALUES ($1, $2, $3)RETURNING user_id"
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(name, email, age).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"last_inserted_id": int64(id),
	}

	return res, err
}

func UpdateUser(id int, name string, email string, age int) (Response, error) {
	var obj User
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE users SET name= $1, email = $2, age = $3 WHERE user_id = $4"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, email, age, id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	sqlStatement = "SELECT * FROM users where email=$1"

	err = con.QueryRow(sqlStatement, email).Scan(&obj.ID, &obj.Name, &obj.Email, &obj.Age)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]interface{}{
		"data":          obj,
		"rows_affected": rowsAffected,
	}
	// res.Data = obj

	return res, err
}

func DeleteUser(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM users where user_id=$1"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, err
}
