package storage

import (
	"context"
	"errors"
)

type Storage[K comparable, V any] interface {
	Get(ctx context.Context, k K) ([]V, error)

	Set(ctx context.Context, k K, v []V) error

	Delete(ctx context.Context, k K) error

	Append(ctx context.Context, values map[K][]V) error

	GetAll(ctx context.Context) (map[K][]V, error)
}

var (
	ErrKeyNotFound = errors.New("key not found")
)
