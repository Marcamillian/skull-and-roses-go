package handlers

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestHandleBookRequest(t *testing.T) {

	expectedBook := "Catch22"
	expectedPage := "35"

	expectedResponse := fmt.Sprintf("You've requested the book: %s on page %s\n", expectedBook, expectedPage)

	// create the request (input) to the function
	req, err := http.NewRequest("GET", "/books/{title}/page/{page}", nil)
	if err != nil {
		t.Fatal(err)
	}

	// add the url parameters to the request
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("title", expectedBook)
	rctx.URLParams.Add("page", expectedPage)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	// make the response recorder
	rr := httptest.NewRecorder()

	// apply the handler to the response recorder
	handler := http.HandlerFunc(HandleBookRequest)
	handler.ServeHTTP(rr, req)

	// check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned the wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	fmt.Println("Something is wrong")

	// check the body of the code
	if rr.Body.String() != expectedResponse {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expectedResponse)
	}

}
