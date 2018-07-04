package main

import (
	"log"
	"login/userlogin/loginuser/login"
	updatepassword "login/userlogin/updatepassword/updatepass"
	"login/userlogin/userforget/forgotpassword"
	"login/userlogin/usersignup/signup"
	"net/http"
	"shared"

	"github.com/gorilla/mux"
)

func IndexPage(w http.ResponseWriter, r *http.Request) {

	shared.LogsInformation(r)
	// consensus := externalip.DefaultConsensus(nil, nil)
	// ip, err := consensus.ExternalIP()
	// externalip := ip.String()
	// fmt.Println(ip, externalip)
	// fmt.Printf("/%T\n", externalip)
	// if err == nil {
	// 	fmt.Println("External IP", ip.String()) // print IPv4/IPv6 in string format
	// }
	//
	// result := "http://api.ipstack.com/" + externalip + "?access_key=5dd814cb89a53b26899a5060c8de5d42"
	// response, err := http.Get(result)
	// if err != nil {
	// 	fmt.Printf("%s", err)
	// 	os.Exit(1)
	// } else {
	// 	defer response.Body.Close()
	// 	contents, err := ioutil.ReadAll(response.Body)
	// 	if err != nil {
	// 		fmt.Printf("%s", err)
	// 		os.Exit(1)
	// 	}
	//
	// 	var info shared.GeoIP
	// 	err1 := json.Unmarshal([]byte(contents), &info)
	// 	if err1 != nil {
	// 		fmt.Println(err1)
	// 	}
	// 	fmt.Println(info)
	// 	fmt.Println(r.UserAgent())
	//
	// }
}

func main() {
	shared.Connect()
	router := mux.NewRouter()
	router.HandleFunc("/signup", signup.SignUpEndpoint).Methods("POST")
	router.HandleFunc("/login", login.LoginEndpoint).Methods("POST")
	router.HandleFunc("/forgotpassword", forgotpassword.ForgotPasswordEndpoint).Methods("POST")
	router.HandleFunc("/forgotpasswordreturn", updatepassword.ForgotPasswordReturnEndpoint).Methods("POST")
	router.HandleFunc("/", IndexPage).Methods("GET")

	log.Fatal(http.ListenAndServe(":9090", router))

}
