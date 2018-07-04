package forgotpassword

import (
	"encoding/json"
	"fmt"
	"net/http"
	"shared"

	"github.com/couchbase/gocb"
)

var usersignup shared.UserSignup

func ForgotPasswordEndpoint(w http.ResponseWriter, req *http.Request) {
	bucket := shared.Connect()
	var n1qlparams []interface{}

	err := json.NewDecoder(req.Body).Decode(&usersignup)

	if err != nil {
		fmt.Println("Json Error", err)
	}

	var forgetpassword shared.UserSignup
	bucket.Get(usersignup.Email, &forgetpassword)
	fmt.Println("Email checker", forgetpassword.Email)
	if usersignup.Email != forgetpassword.Email {
		fmt.Println("Create account first")
	} else {

		token := shared.TokenGenerator(20)
		usersignup.Token = token
		shared.AddToken(usersignup)
		//json.NewEncoder(w).Encode(usersignup)
		fmt.Println("JSON FORMAT", usersignup)
		fmt.Println(usersignup.Email)

		query := gocb.NewN1qlQuery("UPDATE `Login-Signup` SET token = $1 WHERE email = $2")
		n1qlparams = append(n1qlparams, usersignup.Token)
		n1qlparams = append(n1qlparams, usersignup.Email)

		_, err = bucket.ExecuteN1qlQuery(query, n1qlparams)
		if err != nil {
			fmt.Println("ERROR EXECUTING N1QL QUERY:", err)
		}

		if usersignup.Token == forgetpassword.Token {
			fmt.Println("Token Authenticated")
		}

		fmt.Println("Forgot Password", forgetpassword)
		//var url = "<a href='http://localhost:8080/confirm-email%27%3E</a>" + token
		url := "http://localhost:9090/forgotpasswordreturn?token=" + usersignup.Token + "&email=" + usersignup.Email
		content := "To Reset password click on this link " + "<a href=" + url + "> Reset Password</a></strong>"

		shared.SendEmail(usersignup.UserName, usersignup.Email, usersignup.Password, content)
	}

}
