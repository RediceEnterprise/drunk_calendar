package main

import (
	"net/http"
	"time"
)

type App struct {
	*RepositoryStub
	*http.ServeMux
}

func NewApp() *App {
	a := &App{&RepositoryStub{
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
	}, &http.ServeMux{}}
	return a
}
