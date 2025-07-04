package auth

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

func TestGetApi(t *testing.T) {
		type test struct {
				name string
				input http.Header
				result string
				error error
		}

		invalidHeader := http.Header{};
		invalidHeader.Add("Authorization", "InvalidAuthHeader");

		tests := []test{
				{ name: "No header", input: http.Header{}, result: "", error: errors.New("no authorization header included") },
				{ name: "Invalid header", input: invalidHeader, result: "", error: errors.New("malformed authorization header") },
		}

		for _, tc := range tests {
				response, error := GetAPIKey(tc.input);

				if !reflect.DeepEqual(response, tc.result) {
						t.Fatalf("%s: expected value: %v, got %v", tc.name, tc.result, response);
				}

				if !reflect.DeepEqual(error.Error(), tc.error.Error()) {
						t.Fatalf("%s: expected error: %v, got %v", tc.name, tc.error.Error(), error.Error());
				}
		}
}
