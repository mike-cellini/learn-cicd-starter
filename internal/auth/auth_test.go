package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	header := http.Header{}

	header.Set("Authorization", "ApiKey 12345")
	got, err := GetAPIKey(header)
	want := "12345"
	if err != nil {
		t.Fatalf("expected: %v, got: %v", want, err.Error())
	} else if want != got {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
