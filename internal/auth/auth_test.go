package auth

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {

	cases := []struct {
		name    string
		header  string
		prefix  string
		apiKey  string
		wantErr bool
	}{
		{
			name:    "Can get ApiKey from Authorization",
			header:  "Authorization",
			prefix:  "ApiKey",
			apiKey:  "someapikey",
			wantErr: false,
		},
		{
			name:    "Empty ApiKey in Authorization",
			header:  "Authorization",
			prefix:  "ApiKey",
			apiKey:  "",
			wantErr: false,
		},
		{
			name:    "Wrong prefix in Authorization",
			header:  "Authorization",
			prefix:  "Bearer",
			apiKey:  "someapikey",
			wantErr: true,
		},
		{
			name:    "No Authorization",
			header:  "",
			prefix:  "",
			apiKey:  "",
			wantErr: true,
		},
	}

	for _, c := range cases {

		t.Run(c.name, func(t *testing.T) {

			header := http.Header{}
			if c.header != "" {
				header.Add(c.header, fmt.Sprintf("%s %s", c.prefix, c.apiKey))
			}
			apiKey, err := GetAPIKey(header)
			if (err != nil) != c.wantErr {
				t.Errorf("If I want err: %v. Got error: %v", c.wantErr, err)
			}
			if err == nil && apiKey != c.apiKey {
				t.Errorf("Expected ApiKey: %v, but got: %v", c.apiKey, apiKey)
			}
		})
	}
}
