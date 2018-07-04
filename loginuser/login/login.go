package login

import (
	"encoding/json"
	"fmt"
	"net/http"
	"shared"

	"github.com/couchbase/gocb"
)

func LoginEndpoint(w http.ResponseWriter, req *http.Request) {
	var login shared.UserLogin

	bucket := shared.Connect()
	var n1qlparams []interface{}
	err := json.NewDecoder(req.Body).Decode(&login)

	if err != nil {
		fmt.Println("Login Json Error", err)
	}

	query := gocb.NewN1qlQuery("SELECT * FROM `Login-Signup` WHERE email = $1")
	fmt.Println(login.Email)
	n1qlparams = append(n1qlparams, login.Email)

	rows, err := bucket.ExecuteN1qlQuery(query, n1qlparams)
	if err != nil {
		fmt.Println("ERROR EXECUTING N1QL QUERY:", err)
	}

	var row interface{}
	for rows.Next(&row) {
		fmt.Printf("Login: %+v\n", row)
	}

	json.NewEncoder(w).Encode(login)
	//bucket.Get(usersignup.Email, &login)
	fmt.Println("---", login)
	//
	var loginmap shared.UserSignup

	bucket.Get(login.Email, &loginmap)

	fmt.Println(loginmap)

	if loginmap.Email == login.Email && loginmap.Password == login.Password {
		fmt.Println("Login Sucessfully")
	} else {
		fmt.Println("Login Failed")
	}

	// err := Decode(row, &checklogin)

}
