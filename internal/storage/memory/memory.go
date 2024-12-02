package memory

import (
	"context"
	"fmt"
	"maps"
	"sync"

	"local/transaction/internal/storage"
)

type Storage[K comparable, V any] struct {
	sync.RWMutex
	objects map[K][]V
}

func (c *Storage[K, V]) Get(_ context.Context, k K) ([]V, error) {
	c.RLock()
	defer c.RUnlock()

	value, found := c.objects[k]
	if !found {
		return value, fmt.Errorf("%w: %v", storage.ErrKeyNotFound, k)
	}

	return value, nil
}

func (c *Storage[K, V]) Set(_ context.Context, k K, v []V) error {
	c.Lock()
	defer c.Unlock()

	c.objects[k] = v
	return nil
}

func (c *Storage[K, V]) Delete(_ context.Context, k K) error {
	c.Lock()
	defer c.Unlock()
	delete(c.objects, k)
	return nil
}

func (c *Storage[K, V]) Append(_ context.Context, values map[K][]V) error {
	c.Lock()
	defer c.Unlock()

	for k, v := range values {
		c.objects[k] = append(c.objects[k], v...)
	}

	return nil
}

func (c *Storage[K, V]) GetAll(_ context.Context) (map[K][]V, error) {
	c.Lock()
	defer c.Unlock()

	cloned := maps.Clone(c.objects)

	return cloned, nil
}

func New[K comparable, V any]() storage.Storage[K, V] {
	return &Storage[K, V]{objects: make(map[K][]V)}
}
