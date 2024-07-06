package main

import (
	"testing"
)

func TestParadise(t *testing.T) {
	md := &MockDataStore{
		Users: map[int]User{
			2: {ID: 2, Frist: "Jenny"},
		},
	}
	s := &Service{
		ds: md,
	}
	u, err := s.GetUser(2)
	if err != nil {
		t.Errorf("Got error : %v", err)
	}
	if u.Frist != "Jenny" {
		t.Errorf("got : %s , wants : %s", u.Frist, "Jenny")
	}
}
