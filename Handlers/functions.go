package Handlers

import (
	"encoding/json"
	"fmt"
	"github.com/my_todo/Database"
	"github.com/my_todo/Models"
	"log"
	"net/http"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func Test(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("testing")
	return nil
}

func Register(w http.ResponseWriter, r *http.Request) error {
	var u Models.Users
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	sqlStatement := `INSERT INTO users (Name, Email, Password) VALUES ($1, $2, $3)`
	_, err = Database.DB.Exec(sqlStatement, u.Name, u.Password, u.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	fmt.Println("check")
	w.WriteHeader(http.StatusOK)
	return nil
}

func Show(w http.ResponseWriter, r *http.Request) error {

	rows, err := Database.DB.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}

	var u []Models.Users

	for rows.Next() {
		var newU Models.Users
		rows.Scan(&newU.Name, &newU.Email, &newU.Password)
		u = append(u, newU)
	}

	peopleBytes, _ := json.MarshalIndent(u, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(peopleBytes)

	return nil
}

/*
func Login(w http.ResponseWriter, r *http.Request) error {
	q := r.URL.Query().Get("err")
	if q != "" {
		return errors.New(q)
	}
	w.Write([]byte("foo"))
	return nil
}

func Signup(w http.ResponseWriter, r *http.Request) error {
	q := r.URL.Query().Get("err")
	if q != "" {
		return errors.New(q)
	}
	w.Write([]byte("foo"))
	return nil
}
*/

// CreateUser create a user in the postgres db
func CreateUser(w http.ResponseWriter, r *http.Request) error {
	// set the header to content type x-www-form-urlencoded
	// Allow all origin to handle cors issue
	fmt.Println("i'm at create")
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// create an empty user of type models.User
	var user Models.Users

	// decode the json request to user
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	// call insert user function and pass the user
	insertID := insertUser(user)

	// format a response object
	res := response{
		ID:      insertID,
		Message: "User created successfully",
	}

	// send the response
	json.NewEncoder(w).Encode(res)

	return nil
}

// insert one user in the DB
func insertUser(user Models.Users) int64 {

	// create the postgres db connection

	// close the db connection

	// create the insert sql query
	// returning userid will return the id of the inserted user
	sqlStatement := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`

	// the inserted id will store in this id
	var id int64

	// execute the sql statement
	// Scan function will save the insert id in the id
	err := Database.DB.QueryRow(sqlStatement, user.Name, user.Email, user.Password).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	// return the inserted id
	return id
}
