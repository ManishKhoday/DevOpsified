package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandlers(t *testing.T) {
	tests := []struct {
		name                string
		route               string
		expectedStatus      int
		expectedContentType string
		expectedBodyContent string // Optional: for more specific content checks
	}{
		{
			name:                "HomePage",
			route:               "/home",
			expectedStatus:      http.StatusOK,
			expectedContentType: "text/html; charset=utf-8",
			expectedBodyContent: "<title>Home</title>", // Example
		},
		{
			name:                "ResourcePage",
			route:               "/resource",
			expectedStatus:      http.StatusOK,
			expectedContentType: "text/html; charset=utf-8",
			expectedBodyContent: "<title>Resource</title>", // Example
		},
		{
			name:                "AboutPage",
			route:               "/about",
			expectedStatus:      http.StatusOK,
			expectedContentType: "text/html; charset=utf-8",
			expectedBodyContent: "<title>About</title>", // Example
		},
		{
			name:                "ContactPage",
			route:               "/contact",
			expectedStatus:      http.StatusOK,
			expectedContentType: "text/html; charset=utf-8",
			expectedBodyContent: "<title>Contact</title>", // Example
		},
		// Add more test cases here!
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) { // Important: use t.Run for subtests
			req, err := http.NewRequest("GET", tt.route, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(getHandlerForRoute(tt.route)) // Get the correct handler

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler for %s returned wrong status code: got %v want %v",
					tt.route, status, tt.expectedStatus)
			}

			if contentType := rr.Header().Get("Content-Type"); contentType != tt.expectedContentType {
				t.Errorf("handler for %s returned unexpected content type: got %v want %v",
					tt.route, contentType, tt.expectedContentType)
			}
			if tt.expectedBodyContent != "" {
				body := rr.Body.String()
				if !strings.Contains(body, tt.expectedBodyContent) {
					t.Errorf("handler for %s returned unexpected body content: did not find '%s'", tt.route, tt.expectedBodyContent)
				}
			}
		})
	}
}
func getHandlerForRoute(route string) http.HandlerFunc {
	switch route {
	case "/home":
		return homePage
	case "/resource":
		return resourcePage
	case "/about":
		return aboutPage
	case "/contact":
		return contactPage
	default:
		return func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "Not Found", http.StatusNotFound)
		}
	}
}
