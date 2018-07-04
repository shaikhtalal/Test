package updatepassword

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
	router.HandleFunc("/forgotpasswordreturn", ForgotPasswordReturnEndpoint).Methods("POST")

	return router
}

func TestSignUpEndpoint(t *testing.T) {

	passwordupdate := shared.UserSignup{

		Password: "0123456789abcd",
	}
	passwordupdatejson, _ := json.Marshal(passwordupdate)

	request, err := http.NewRequest("POST", "/forgotpasswordreturn", bytes.NewBuffer(passwordupdatejson))
	if err != nil {
		fmt.Println("Request Error", err)
	}

	response := httptest.NewRecorder()

	Router().ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code, "OK reponse is expected")

}
