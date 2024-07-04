package main

import "fmt"

type User struct {
	ID    int
	First string
}

// MockDatastore is atemporary service that sotres retrivable data.
type MockDatastore struct {
	Users map[int]User
}

func (md MockDatastore) GetUser(id int) (User, error) {
	user, ok := md.Users[id]
	if !ok {
		return User{}, fmt.Errorf("User %v not found", id)
	}
	return user, nil
}

func (md MockDatastore) SaveUser(u User) error {
	_, ok := md.Users[u.ID]
	if ok {
		return fmt.Errorf("user %v already exists", u.ID)
	}
	md.Users[u.ID] = u
	return nil
}
