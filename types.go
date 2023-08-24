package main

import "math/rand"

type Account struct {
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
	ID        int    `json:"id"`
	Number    int64  `json:"number"`
	Balance   int64  `json:"balance"`
}

func NewAccount(firstName string, lastName string) *Account {
	return &Account{ID: rand.Intn(100000000), Firstname: firstName, Lastname: lastName, Number: rand.Int63n(10000000000)}
}
