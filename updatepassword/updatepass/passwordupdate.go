package updatepassword

import (
	"encoding/json"
	"fmt"
	"net/http"
	"shared"

	"github.com/couchbase/gocb"
)

var usersignup shared.UserSignup

func ForgotPasswordReturnEndpoint(w http.ResponseWriter, req *http.Request) {
	bucket := shared.Connect()
	var n1qlparams []interface{}

	var pwd shared.UserSignup

	err := json.NewDecoder(req.Body).Decode(&pwd)

	if err != nil {
		fmt.Println("Json Error", err)
	}

	map1 := req.URL.Query()
	var token = map1.Get("token")

	map2 := req.URL.Query()
	//var email = map2.Get("email")

	fmt.Println("URL EMAIL", map2.Get("email"))
	fmt.Println("URL TOKEN", token)

	var checktoken shared.UserSignup
	//	usersignup.Email = email

	// usersignup.Token = token
	bucket.Get(usersignup.Email, &usersignup)

	checktoken.Token = token

	fmt.Println("This your url token: ", checktoken.Token)
	fmt.Println("This is Couchbase token: ", usersignup.Token)
	fmt.Print("Check Total Struct", usersignup)

	if usersignup.Token != checktoken.Token {
		fmt.Println("Invalid Token")

	} else {
		//bucket.Get(usersignup.Email, &checktoken)
		fmt.Println("Valid Token")
		fmt.Println("check token struct password", pwd.Password)
		fmt.Println("-------------", checktoken)
		//fmt.Println(usersignup)
		shared.UpdatePassword(pwd)
		query := gocb.NewN1qlQuery("UPDATE `Login-Signup` SET `password` = $1 WHERE email = $2")
		n1qlparams = append(n1qlparams, pwd.Password)
		n1qlparams = append(n1qlparams, usersignup.Email)

		_, err = bucket.ExecuteN1qlQuery(query, n1qlparams)
		if err != nil {
			fmt.Println("ERROR EXECUTING N1QL QUERY:", err)

		}

		var newpassword shared.UserSignup
		bucket.Get(usersignup.Email, &newpassword)
		fmt.Println(newpassword)

	}

}
