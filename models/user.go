package models

type User struct {
	ID         uint64
	Email      string
	Name       string
	Group      uint64
	Department string
}
