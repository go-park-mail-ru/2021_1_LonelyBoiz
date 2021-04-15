package old_tests

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"server/api"
	"testing"
)

func TestValidateCookie(t *testing.T) {
	testCases := []struct {
		name string
		in   string
		out  bool
	}{
		{
			name: "Success case",
			in:   "testcookie",
			out:  true,
		},
		{
			name: "Invalid cookie",
			in:   "invcookie",
			out:  false,
		},
	}

	a := api.App{Sessions: map[int][]http.Cookie{
		0: {
			{
				Name:  "token",
				Value: "testcookie",
			},
			{
				Name:  "token2",
				Value: "testcookie2",
			},
		},
	}}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.out, a.validateCookie(testCase.in))
		})
	}
}
