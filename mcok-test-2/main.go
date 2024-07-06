package main

import (
	"fmt"
	"log"
)

// User represents a user with an id and first name
type User struct {
	ID    int
	Frist string
}

// MockDataStore is a temporary service that stores retrivable data
type MockDataStore struct {
	Users map[int]User
}

func (md MockDataStore) GetUser(id int) (User, error) {
	user, ok := md.Users[id]
	if !ok {
		return User{}, fmt.Errorf("User %v not found", id)
	}
	return user, nil
}

func (md MockDataStore) SaveUser(u User) error {
	_, ok := md.Users[u.ID]
	if ok {
		return fmt.Errorf("User %v already exists", u.ID)
	}
	md.Users[u.ID] = u
	return nil
}

type DataStore interface {
	GetUser(id int) (User, error)
	SaveUser(u User) error
}

// Servive reperesents a service that stores retrivable data
// We will atach a methods to service  so that we can use either of thers :
// -- MockDataStore
// -- *sql.DB
type Service struct {
	ds DataStore
}

func (s Service) GetUser(id int) (User, error) {
	return s.ds.GetUser(id)
}
func (s Service) SaveUser(u User) error {
	return s.ds.SaveUser(u)
}

func main() {
	db := MockDataStore{
		Users: make(map[int]User),
	}
	servs := Service{
		ds: db,
	}
	u1 := User{
		ID:    1,
		Frist: "Nahi",
	}
	err := servs.SaveUser(u1)
	if err != nil {
		log.Fatalf("Error %s", err)
	}
	u1Retuned, err := servs.GetUser(u1.ID)
	if err != nil {
		log.Fatalf("Error %s", err)
	}

	fmt.Println(u1)
	fmt.Println(u1Retuned)

}
