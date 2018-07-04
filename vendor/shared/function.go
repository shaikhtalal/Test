package shared

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/bradfitz/latlong"
	gocb "github.com/couchbase/gocb"
	"github.com/glendc/go-external-ip"
	"github.com/joho/godotenv"
	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func Connect() *gocb.Bucket {
	err := godotenv.Load("../../../.env")
	if err != nil {
		fmt.Println("Environment Error", err)
	}

	cluster, err := gocb.Connect("couchbase://localhost")
	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: os.Getenv("ADMINISTRATOR"),
		Password: os.Getenv("PASSWORD"),
	})

	if err != nil {
		fmt.Println("Cluster Connection", err)
	}

	fmt.Println(os.Getenv("ADMINISTRATOR"))
	fmt.Println(os.Getenv("PASSWORD"))

	bucketname := os.Getenv("BUCKETNAME")
	fmt.Println(os.Getenv("BUCKETNAME"))
	bucket, err := cluster.OpenBucket(bucketname, "")
	bucket.Manager("", "").CreatePrimaryIndex("", true, false)
	if err != nil {
		fmt.Println("Bucket Opening", err)
	}

	return bucket
}

func Add(data UserSignup) bool {
	bucket := Connect()
	bucket.Insert(data.Email, data, 0)

	return true

}

func AddToken(data UserSignup) bool {
	bucket := Connect()
	bucket.Insert(data.Email, data.Token, 0)

	return true

}
func UpdatePassword(data UserSignup) bool {
	bucket := Connect()
	bucket.Insert(data.Token, data.Password, 0)
	return true
}

func RandStringBytes(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func SendEmail(username string, email string, password string, content string) bool {

	from := mail.NewEmail("Login", "no-reply@Signup-Login.io")
	subject := "Login-Signup"
	to := mail.NewEmail(email, email)
	plainTextContent := "Login-Signup"
	htmlContent := content //"Hello " + username + ", Your Password is:" + password
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient("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")
	response, err := client.Send(message)

	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
	return true
}

func ValidateEmail(email string) bool {

	Regix := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

	return Regix.MatchString(email)
}

func CheckUser(username string, password string) bool {

	if username == "" || password == "" {
		fmt.Println("Fields cannot be empty")
		return false
	}

	return true
}

func TokenGenerator(n int) string {
	token := make([]byte, n)
	rand.Read(token)
	return string(token)
}

func ExtractTimeZone(latitude float32, longitute float32) time.Time {
	lat := float64(latitude)
	lon := float64(longitute)
	timezone := latlong.LookupZoneName(lat, lon)
	loc, _ := time.LoadLocation(timezone)
	now := time.Now().In(loc)
	fmt.Println(timezone)
	return now
}

func LogsInformation(req *http.Request) {

	consensus := externalip.DefaultConsensus(nil, nil)
	ip, err := consensus.ExternalIP()
	externalip := ip.String()
	fmt.Println(ip, externalip)
	fmt.Printf("/%T\n", externalip)
	if err == nil {
		fmt.Println("External IP", ip.String()) // print IPv4/IPv6 in string format
	}

	ipurl := "http://api.ipstack.com/" + externalip + "?access_key=5dd814cb89a53b26899a5060c8de5d42"

	response, err := http.Get(ipurl)

	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {

		defer response.Body.Close()
		logsdata, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}

		var geoip GeoIP

		err1 := json.Unmarshal([]byte(logsdata), &geoip)
		if err1 != nil {
			fmt.Println(err1)
		}
		fmt.Println(geoip)
		fmt.Println(req.UserAgent())

		timezone := ExtractTimeZone(geoip.Lat, geoip.Lon)

		fmt.Println("Current Time according to time zone", timezone)

	}
}
