package main

import "github.com/sirupsen/logrus"

type Repository interface {
	Book(string) (string, error)
}
type BookRepository struct {
}

func (b *BookRepository) Book(id string) (string, error) {
	logrus.Info("received book req")
	return "success", nil
}
