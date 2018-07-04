package main

//var usersignup shared.UserSignup

//USER SIGNUP ENDPOINT

// func SignUpEndpoint(w http.ResponseWriter, req *http.Request) {
// 	bucket := shared.Connect()
// 	var n1qlparams []interface{}
//
// 	err := json.NewDecoder(req.Body).Decode(&usersignup)
//
// 	if err != nil {
// 		fmt.Println("Json Error", err)
// 	}
//
// 	var newuser shared.UserSignup
// 	bucket.Get(usersignup.Email, &newuser)
//
// 	if newuser.Email == usersignup.Email {
// 		fmt.Println("Email Already Exists")
// 	} else {
//
// 		pwd := shared.RandStringBytes(10)
// 		usersignup.Password = pwd
// 		shared.Add(usersignup)
// 		//json.NewEncoder(w).Encode(usersignup)
// 		fmt.Println("JSON FORMAT", usersignup)
// 		fmt.Println(usersignup.Email)
//
// 		query := gocb.NewN1qlQuery("UPDATE `Login-Signup` SET `password` = $1 WHERE email = $2")
// 		n1qlparams = append(n1qlparams, usersignup.Password)
// 		n1qlparams = append(n1qlparams, usersignup.Email)
//
// 		_, err = bucket.ExecuteN1qlQuery(query, n1qlparams)
// 		if err != nil {
// 			fmt.Println("ERROR EXECUTING N1QL QUERY:", err)
// 		}
//
// 		fmt.Println("New User SignUp Data", newuser)
//
// 		var email = usersignup.Email
//
// 		if !shared.ValidateEmail(email) {
// 			fmt.Println("Invalid Email")
//
// 		}
//
// 		shared.CheckUser(usersignup.UserName, usersignup.Email)
// 		content := "Hello " + usersignup.UserName + ", Your Password is " + usersignup.Password
// 		shared.SendEmail(usersignup.UserName, usersignup.Email, usersignup.Password, content)
//
// 		var token = shared.TokenGenerator(10)
// 		fmt.Println(token)
// 	}
// }

//LOGIN ENDPOINT

// func LoginEndpoint(w http.ResponseWriter, req *http.Request) {
// 	var login shared.UserLogin
//
// 	bucket := shared.Connect()
// 	var n1qlparams []interface{}
// 	err := json.NewDecoder(req.Body).Decode(&login)
//
// 	if err != nil {
// 		fmt.Println("Login Json Error", err)
// 	}
//
// 	query := gocb.NewN1qlQuery("SELECT * FROM `Login-Signup` WHERE email = $1")
// 	fmt.Println(login.Email)
// 	n1qlparams = append(n1qlparams, login.Email)
//
// 	rows, err := bucket.ExecuteN1qlQuery(query, n1qlparams)
// 	if err != nil {
// 		fmt.Println("ERROR EXECUTING N1QL QUERY:", err)
// 	}
//
// 	var row interface{}
// 	for rows.Next(&row) {
// 		fmt.Printf("Login: %+v\n", row)
// 	}
//
// 	json.NewEncoder(w).Encode(login)
// 	//bucket.Get(usersignup.Email, &login)
// 	fmt.Println("---", login)
// 	//
// 	var loginmap shared.UserSignup
//
// 	bucket.Get(login.Email, &loginmap)
//
// 	fmt.Println(loginmap)
//
// 	if loginmap.Email == login.Email && loginmap.Password == login.Password {
// 		fmt.Println("Login Sucessfully")
// 	} else {
// 		fmt.Println("Login Failed")
// 	}
//
// 	// err := Decode(row, &checklogin)
//
// }

//FORGOT PASSWORD ENDPOINT

// func ForgotPasswordEndpoint(w http.ResponseWriter, req *http.Request) {
// 	bucket := shared.Connect()
// 	var n1qlparams []interface{}
//
// 	err := json.NewDecoder(req.Body).Decode(&usersignup)
//
// 	if err != nil {
// 		fmt.Println("Json Error", err)
// 	}
//
// 	var forgetpassword shared.UserSignup
// 	bucket.Get(usersignup.Email, &forgetpassword)
// 	fmt.Println("Email checker", forgetpassword.Email)
// 	if usersignup.Email != forgetpassword.Email {
// 		fmt.Println("Create account first")
// 	} else {
//
// 		token := shared.TokenGenerator(20)
// 		usersignup.Token = token
// 		shared.AddToken(usersignup)
// 		//json.NewEncoder(w).Encode(usersignup)
// 		fmt.Println("JSON FORMAT", usersignup)
// 		fmt.Println(usersignup.Email)
//
// 		query := gocb.NewN1qlQuery("UPDATE `Login-Signup` SET token = $1 WHERE email = $2")
// 		n1qlparams = append(n1qlparams, usersignup.Token)
// 		n1qlparams = append(n1qlparams, usersignup.Email)
//
// 		_, err = bucket.ExecuteN1qlQuery(query, n1qlparams)
// 		if err != nil {
// 			fmt.Println("ERROR EXECUTING N1QL QUERY:", err)
// 		}
//
// 		if usersignup.Token == forgetpassword.Token {
// 			fmt.Println("Token Authenticated")
// 		}
//
// 		fmt.Println("Forgot Password", forgetpassword)
// 		//var url = "<a href='http://localhost:8080/confirm-email%27%3E</a>" + token
// 		url := "http://localhost:9090/forgotpasswordreturn?token=" + usersignup.Token + "&email=" + usersignup.Email
// 		content := "To Reset password click on this link " + "<a href=" + url + "> Reset Password</a></strong>"
//
// 		shared.SendEmail(usersignup.UserName, usersignup.Email, usersignup.Password, content)
// 	}
//
// }

//FORGOT PASSWORD ENDPOINT USING WITH TOKEN AND EMAIL

// func ForgotPasswordReturnEndpoint(w http.ResponseWriter, req *http.Request) {
// 	bucket := shared.Connect()
// 	var n1qlparams []interface{}
//
// 	var pwd shared.UserSignup
//
// 	err := json.NewDecoder(req.Body).Decode(&pwd)
//
// 	if err != nil {
// 		fmt.Println("Json Error", err)
// 	}
//
// 	map1 := req.URL.Query()
// 	var token = map1.Get("token")
//
// 	map2 := req.URL.Query()
// 	//var email = map2.Get("email")
//
// 	fmt.Println("URL EMAIL", map2.Get("email"))
// 	fmt.Println("URL TOKEN", token)
//
// 	var checktoken shared.UserSignup
// 	//	usersignup.Email = email
//
// 	// usersignup.Token = token
// 	bucket.Get(usersignup.Email, &usersignup)
// 	checktoken.Token = token
// 	fmt.Println("This your url token: ", checktoken.Token)
// 	fmt.Println("This is Couchbase token: ", usersignup.Token)
// 	fmt.Print("Check Total Struct", usersignup)
// 	if usersignup.Token != checktoken.Token {
// 		fmt.Println("Invalid Token")
//
// 	} else {
// 		//bucket.Get(usersignup.Email, &checktoken)
// 		fmt.Println("Valid Token")
// 		fmt.Println("check token struct password", pwd.Password)
// 		fmt.Println("-------------", checktoken)
// 		//fmt.Println(usersignup)
// 		shared.UpdatePassword(pwd)
// 		query := gocb.NewN1qlQuery("UPDATE `Login-Signup` SET `password` = $1 WHERE email = $2")
// 		n1qlparams = append(n1qlparams, pwd.Password)
// 		n1qlparams = append(n1qlparams, usersignup.Email)
//
// 		_, err = bucket.ExecuteN1qlQuery(query, n1qlparams)
// 		if err != nil {
// 			fmt.Println("ERROR EXECUTING N1QL QUERY:", err)
//
// 		}
//
// 		var newpassword shared.UserSignup
// 		bucket.Get(usersignup.Email, &newpassword)
// 		fmt.Println(newpassword)
//
// 	}
//
// }

// func main() {

//shared.Connect()
// router := mux.NewRouter()
// routes.Router()
// log.Fatal(http.ListenAndServe(":9090", router))

// router.HandleFunc("/signup", signup.SignUpEndpoint).Methods("POST")
// router.HandleFunc("/login", login.LoginEndpoint).Methods("POST")
// router.HandleFunc("/forgotpassword", forgotpassword.ForgotPasswordEndpoint).Methods("POST")
// router.HandleFunc("/forgotpasswordreturn", updatepassword.ForgotPasswordReturnEndpoint).Methods("POST")

//
