package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleJsonDecode(t *testing.T) {

	expectedFirstname := "Joe"
	expectedLastname := "Bloggs"
	expectedAge := "35"

	expectedStatus := http.StatusOK

	expectedResponse := fmt.Sprintf("%v %v is %v years old!",
		expectedFirstname, expectedLastname, expectedAge)

	requestBody := []byte(`{
		"firstname": "Joe",
		"lastname": "Bloggs",
		"age":35
	}`)

	req, err := http.NewRequest("POST", "/json/decode", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(HandleJsonDecode)
	handler.ServeHTTP(rr, req)

	// check the status
	if status := rr.Code; status != expectedStatus {
		t.Errorf("handler returned the wrong status code: got %v want %v",
			status, expectedStatus)
	}

	// check the returned value
	if rr.Body.String() != expectedResponse {
		t.Errorf("handler returned unexpected body: got %v wanted %v",
			rr.Body.String(), expectedResponse)
	}
}

func TestHandleJsonEncode(t *testing.T) {
	expectedStatus := http.StatusOK

	expectedFirstname := "John"
	expectedLastname := "Doe"
	expectedAge := 25

	// expectedResponse := []byte(`{
	// 	"firstname": "John",
	// 	"lastname": "Doe",
	// 	"age": 25
	// }`)

	// make the request
	req, err := http.NewRequest("GET", "/json/encode", nil)
	if err != nil {
		t.Fatal(err)
	}
	// make the response reader
	rr := httptest.NewRecorder()

	// make the request
	handler := http.HandlerFunc(HandleJsonEncode)
	handler.ServeHTTP(rr, req)

	// check the status
	if status := rr.Code; status != expectedStatus {
		t.Errorf("handler returned the wrong status code: got %v, wanted %v",
			status, expectedStatus)
	}

	// TODO:Marc - finish the test for the retuned JSON
	var user User
	json.NewDecoder(rr.Body).Decode(&user)

	if user.Firstname != expectedFirstname {
		t.Errorf("Handler returned the wrong body: Got firstname: %v, wanted %v",
			user.Firstname, expectedFirstname)
	}

	if user.Lastname != expectedLastname {
		t.Errorf("Handler returned the wrong body: Got lastname: %v, wanted %v",
			user.Lastname, expectedLastname)
	}

	if user.Age != expectedAge {
		t.Errorf("Handler returned the wrong body: Got age: %v, wanted %v",
			user.Age, expectedAge)
	}

}
