package tests

import (
	"server/api"
	"testing"
)

var TestCaseListUsers = []struct {
	users  map[int]User
	finder User
	res    int
}{
	{users: map[int]User{0: {Sex: "male", DatePreference: "male"}}, finder: User{Id: 1, Sex: "male", DatePreference: "male"}, res: 1},
	{users: map[int]User{0: {Sex: "male", DatePreference: "male"}}, finder: User{Id: 0, Sex: "male", DatePreference: "male"}, res: 0},
	{users: map[int]User{
		0: {Sex: "male", DatePreference: "male"},
		1: {Sex: "male", DatePreference: "male"},
		2: {Sex: "male", DatePreference: "male"},
		3: {Sex: "male", DatePreference: "male"},
		4: {Sex: "male", DatePreference: "male"},
		5: {Sex: "male", DatePreference: "male"}},
		finder: User{Id: 6, Sex: "male", DatePreference: "male"},
		res:    5},
}

func TestListUsers(t *testing.T) {
	var a api.App
	for _, v := range TestCaseListUsers {
		a.Users = v.users
		res := a.listUsers(v.finder)
		if len(res) != v.res {
			t.Error("List users error", "\nExpeted:", v.res, "\nGot:", len(res))
		}
	}
}
