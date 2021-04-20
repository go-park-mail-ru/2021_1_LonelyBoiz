package tests

import (
	"net/http"
	"server/api"
	"testing"
)

func TestDeleteCookie(t *testing.T) {
	key := "key"
	a := api.App{Sessions: map[int][]http.Cookie{
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
