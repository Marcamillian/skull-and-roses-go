package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleTodoList(t *testing.T) {

	// expected values
	expectedStatus := http.StatusOK

	// make the request
	req, err := http.NewRequest("GET", "/todo/MarcList", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(HandleTodoList)
	handler.ServeHTTP(rr, req)

	// check the status
	if status := rr.Code; status != expectedStatus {
		t.Errorf("handler returned the wrong status code: got %v wanted %v",
			status, expectedStatus)
	}

}
