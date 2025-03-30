package auth

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	tests := []struct {
		input http.Header
		want  string
		err   error
	}{
		{
			input: http.Header{
				"Authorization": []string{""},
			},
			want: "",
			err:  ErrNoAuthHeaderIncluded,
		},
		{
			input: http.Header{
				"Authorization": []string{"ApiKey 123"},
			},
			want: "123",
			err:  nil,
		},
		{
			input: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			want: "",
			err:  ErrMalformedAuthHeader,
		},
	}

	for _, tc := range tests {
		got, err := GetAPIKey(tc.input)
		if !errors.Is(err, tc.err) {
			t.Fatalf("expected error to be: %v, got: %v", tc.err, err)
		}

		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
