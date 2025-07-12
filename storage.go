package main

type Storage[T any] struct {
	Filename string
}

func NewStorage[T any](filename string) *Storage[T] {
	return &Storage[T]{Filename:filename}
}