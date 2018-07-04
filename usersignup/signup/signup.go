package signup

import (
	"encoding/json"
	"fmt"
	"net/http"
	"shared"

	"github.com/couchbase/gocb"
)

var usersignup shared.UserSignup

//SIGNUP ENDPOINT

func SignUpEndpoint(w http.ResponseWriter, req *http.Request) {
	bucket := shared.Connect()
	var n1qlparams []interface{}

	err := json.NewDecoder(req.Body).Decode(&usersignup)

	if err != nil {
		fmt.Println("Json Error", err)
	}

	var newuser shared.UserSignup
	bucket.Get(usersignup.Email, &newuser)

	if newuser.Email == usersignup.Email {
		fmt.Println("Email Already Exists")
	} else {

		pwd := shared.RandStringBytes(10)
		usersignup.Password = pwd
		shared.Add(usersignup)
		//json.NewEncoder(w).Encode(usersignup)
		fmt.Println("JSON FORMAT", usersignup)
		fmt.Println(usersignup.Email)

		query := gocb.NewN1qlQuery("UPDATE `Login-Signup` SET `password` = $1 WHERE email = $2")
		n1qlparams = append(n1qlparams, usersignup.Password)
		n1qlparams = append(n1qlparams, usersignup.Email)

		_, err = bucket.ExecuteN1qlQuery(query, n1qlparams)
		if err != nil {
			fmt.Println("ERROR EXECUTING N1QL QUERY:", err)
		}

		fmt.Println("New User SignUp Data", newuser)

		var email = usersignup.Email

		if !shared.ValidateEmail(email) {
			fmt.Println("Invalid Email")

		}

		shared.CheckUser(usersignup.UserName, usersignup.Email)
		content := "Hello " + usersignup.UserName + ", Your Password is " + usersignup.Password
		shared.SendEmail(usersignup.UserName, usersignup.Email, usersignup.Password, content)

		var token = shared.TokenGenerator(10)
		fmt.Println(token)

		fmt.Println("User Added Successfully")

	}
}
