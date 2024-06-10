package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePostPayload_Validate(t *testing.T) {
	tests := []struct {
		name          string
		payload       CreatePostPayload
		expectedError string
	}{
		{
			name: "Valid Payload",
			payload: CreatePostPayload{
				Url:         "https://www.youtube.com/watch?v=asQjT7LSBLA",
				Description: "Justice Live",
			},
			expectedError: "",
		},
		{
			name: "Missing Required Fields",
			payload: CreatePostPayload{
				Url: "",
			},
			expectedError: "Url is required",
		},
		{
			name: "Bad Url",
			payload: CreatePostPayload{
				Url: "Thisisnotanurl",
			},
			expectedError: "Url is not a valid URL",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errs := tt.payload.Validate()

			if tt.expectedError != "" {
				assert.NotNil(t, errs)
				assert.Contains(t, errs[0].Error(), tt.expectedError)
			} else {
				assert.Nil(t, errs)
			}
		})
	}
}
