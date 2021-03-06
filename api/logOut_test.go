package api

import (
	"net/http"
	"testing"
)

func TestDeleteCookie(t *testing.T) {
	key := "key"
	a := App{Sessions: map[int][]http.Cookie{
		0: {
			{
				Name:  "token",
				Value: key,
			},
		},
	}}

	deleteCookie(key, &a.Sessions)
	if len(a.Sessions[0]) != 0 {
		t.Error("Cookie didn't delete")
	}
}
