package main

import (
	//	"fmt"
	//	"net/http"
	//	"net/http/httptest"
	"testing"
	// "github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {
	//	t.Run("Unable to get IP", func(t *testing.T) {
	//		DefaultHTTPGetAddress = "http://127.0.0.1:12345"
	//
	//		_, err := handler(events.APIGatewayProxyRequest{})
	//		if err == nil {
	//			t.Fatal("Error failed to trigger with an invalid request")
	//		}
	//	})

	//	t.Run("Non 200 Response", func(t *testing.T) {
	//		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//			w.WriteHeader(500)
	//		}))
	//		defer ts.Close()
	//
	//		DefaultHTTPGetAddress = ts.URL
	//
	//		_, err := handler(events.APIGatewayProxyRequest{})
	//		if err != nil && err.Error() != ErrNon200Response.Error() {
	//			t.Fatalf("Error failed to trigger with an invalid HTTP response: %v", err)
	//		}
	//	})

}
