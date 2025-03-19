package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		want    string
		wantErr error
	}{
		{
			name:    "No authorization header",
			headers: http.Header{},
			want:    "",
			wantErr: ErrNoAuthHeaderIncluded,
		},
		{
			name: "Malformed authorization header (no ApiKey prefix)",
			headers: http.Header{
				"Authorization": []string{"Bearer 12345"},
			},
			want:    "",
			wantErr: errors.New("malformed authorization header"),
		},
		{
			name: "Malformed authorization header (too short)",
			headers: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			want:    "",
			wantErr: errors.New("malformed authorization header"),
		},
		{
			name: "Valid authorization header",
			headers: http.Header{
				"Authorization": []string{"ApiKey 12345"},
			},
			want:    "12345",
			wantErr: nil,
		},
		{
			name: "Authorization header with extra parts",
			headers: http.Header{
				"Authorization": []string{"ApiKey 12345 67890"},
			},
			want:    "12345",
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)
			// Check for error consistency
			if tt.wantErr != nil {
				if err == nil {
					t.Fatal("expected error but got nil")
				}
				if tt.wantErr.Error() != err.Error() {
					t.Fatalf("expected error %q, got %q", tt.wantErr, err)
				}
			} else {
				if err != nil {
					t.Fatalf("expected no error, got %q", err)
				}
			}
			// Check if the returned key matches the expected value
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}
