package auth

import (
	"net/http"
	"testing"
)

func TestAuth(t *testing.T) {
	type test struct {
		name       string
		headers    http.Header
		want       string
		shouldPass bool
	}

	tests := []test{
		{"should pass",
			http.Header{
				"Accept":        []string{"text/html"},
				"Authorization": []string{"ApiKey 12345"}},
			"12345",
			true,
		},
		{"no auth header", http.Header{
			"Accept": []string{"text/html"}},
			"",
			false,
		},
		{"malformed", http.Header{
			"Accept":        []string{"text/html"},
			"Authorization": []string{"ApiKii"}},
			"",
			false,
		},
	}
	for _, tc := range tests {
		got, err := GetAPIKey(tc.headers)
		if err != nil && tc.shouldPass {
			t.Fatalf("want: %v, got: %v", tc.want, got)
		}
		if got != tc.want {
			t.Errorf("want: %v, got: %v", tc.want, got)
		}
	}

}
