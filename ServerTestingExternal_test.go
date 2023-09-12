package main

import (
	"net/http"
	"testing"
)

// the difference in external and internal is that external testing makes requests out to the specified address. Internal just tests the function
func (ws *WebServer) TestExternalLogin(t *testing.T) {
	resp, err := http.Get("https://" + ws.Context.DomainName + "/login")

	if err != nil {
		t.Errorf("Error: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		t.Errorf("Error, status code: %v", resp.StatusCode)
	}
}

func (ws *WebServer) TestExternalDashboard(t *testing.T) {
	resp, err := http.Get("https://" + ws.Context.DomainName + "/dashboard")

	if err != nil {
		t.Errorf("Error: %v", err)
	}
	defer resp.Body.Close()
	// so here is where we do an unauthenticated external request, so we should be looking to ensure we get a redirect instead of a 200 code.
	if resp.StatusCode != 200 {
		t.Errorf("Error, status code: %v", resp.StatusCode)
	}
}

func (ws *WebServer) TestPing(t *testing.T) {
	resp, err := http.Get("https://" + ws.Context.DomainName + "/ping")

	if err != nil {
		t.Errorf("Error: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		t.Errorf("Error, status code: %v", resp.StatusCode)
	}
}
