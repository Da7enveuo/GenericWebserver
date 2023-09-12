package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// the difference in external and internal is that external testing makes requests out to the specified address. Internal just tests the function
func TestAllEndpoints(t *testing.T) {

}

// test the login function is serving properly internally
func TestLogin(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/login", nil)
	w := httptest.NewRecorder()
	serveLogin(w, r)
	res := w.Result()
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if string(data) != string("someFile") {
		t.Errorf("Exepected different result")
	}

}

//

func TestDashboard(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/dashboard", nil)
	w := httptest.NewRecorder()
	serveLogin(w, r)
	res := w.Result()
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if string(data) != string("someFile") {
		t.Errorf("Exepected different result")
	}

}

func TestTransactions(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/transactions", nil)
	w := httptest.NewRecorder()
	serveLogin(w, r)
	res := w.Result()
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if string(data) != string("someFile") {
		t.Errorf("Exepected different result")
	}

}
