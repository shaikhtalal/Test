package login

import (
	"bytes"
	"encoding/json"
	"fmt"

	"net/http"
	"net/http/httptest"
	"shared"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/login", LoginEndpoint).Methods("POST")

	return router
}

func TestSignUpEndpoint(t *testing.T) {

	signupdata := shared.UserSignup{

		Email:    "shaikh.talal.hussain@gmail.com",
		Password: "0123456789abcd",
	}
	signupjson, _ := json.Marshal(signupdata)

	request, err := http.NewRequest("POST", "/login", bytes.NewBuffer(signupjson))
	if err != nil {
		fmt.Println("Request Error", err)
	}

	response := httptest.NewRecorder()

	Router().ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code, "OK reponse is expected")

}
