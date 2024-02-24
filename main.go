package main

import (
	"fmt"
	"time"
)

func main() {
	rep := RepositoryStub{
		users: []User{
			{1, "User 1"},
			{2, "User 2"},
			{3, "User 3"},
		},
		days: []Day{
			{1, 1, time.Date(2023, 2, 4, 0, 0, 0, 0, time.Local)},
			{2, 1, time.Date(2023, 2, 12, 0, 0, 0, 0, time.Local)},
			{3, 2, time.Date(2023, 5, 27, 0, 0, 0, 0, time.Local)},
			{4, 1, time.Date(2023, 12, 31, 0, 0, 0, 0, time.Local)},
			{5, 2, time.Date(2023, 12, 31, 0, 0, 0, 0, time.Local)},
		},
	}

	u, _ := rep.CreateUser("new user")

	newUser, err := rep.ReadUser(u.id)
	if err != nil {
		panic(err)
	}

	fmt.Println(newUser.name)
}

type Repository interface {
	UserRepository
	DayRepository
}

type UserRepository interface {
	CreateUser(name string) (User, error)
	ReadUser(id uint64) (User, error)
	UpdateUser(user User) (User, error)
	DeleteUser(id uint64) error
}

type DayRepository interface {
	AddDay(Day) error
	LastDay(userId uint64) (Day, error)
	AllDays(userId uint64) ([]Day, error)
}

type User struct {
	id   uint64
	name string
}

type Day struct {
	id     uint64
	userId uint64
	date   time.Time
}

// ---------- repository stub ------------
type RepositoryStub struct {
	users []User
	days  []Day
}

func (r *RepositoryStub) CreateUser(name string) (User, error) {
	var user User
	if len(r.users) == 0 {
		user = User{1, name}
	} else {
		last := r.users[len(r.users)-1]
		user = User{last.id + 1, name}
	}
	r.users = append(r.users, user)
	return user, nil
}

func (r *RepositoryStub) ReadUser(id uint64) (User, error) {
	for _, u := range r.users {
		if u.id == id {
			return u, nil
		}
	}

	return User{}, fmt.Errorf("User %d not found", id)
}
