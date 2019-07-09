package main

import "github.com/go-log/log"

type Repository interface {
	Book(string) (string, error)
}
type BookRepository struct {
}

func (b *BookRepository) Book(id string) (string, error) {
	log.Logf("received book req")
	return "success", nil
}
