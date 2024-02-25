package main

import (
	"fmt"
	"time"
)

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
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

type Day struct {
	Id     uint64    `json:"id"`
	UserId uint64    `json:"user_id"`
	Date   time.Time `json:"date"`
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
		user = User{last.Id + 1, name}
	}
	r.users = append(r.users, user)
	return user, nil
}

func (r *RepositoryStub) ReadUser(id uint64) (User, error) {
	for _, u := range r.users {
		if u.Id == id {
			return u, nil
		}
	}

	return User{}, fmt.Errorf("User %d not found", id)
}
