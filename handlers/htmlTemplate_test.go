package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-shiori/dom"
	"golang.org/x/net/html"
)

func TestHandleTodoList(t *testing.T) {

	// expected values
	expectedStatus := http.StatusOK
	expectedElements := []string{
		"h1",
		"ul",
		"li",
	}

	// === Make the request ====
	req, err := http.NewRequest("GET", "/todo/MarcList", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(HandleTodoList)
	handler.ServeHTTP(rr, req)

	// === Check the response ===

	// check the status
	if status := rr.Code; status != expectedStatus {
		t.Errorf("handler returned the wrong status code: got %v wanted %v",
			status, expectedStatus)
	}

	// querying the DOM - https://ahmadrosid.com/blog/how-to-query-html-dom-in-golang
	document, err := html.Parse(rr.Body)

	if err != nil {
		t.Errorf("coudln't parse document body")
	}

	for _, selector := range expectedElements {
		if dom.QuerySelector(document, selector) == nil {
			t.Errorf("document doesn't contain element with selector %s", selector)
		}
	}

}
