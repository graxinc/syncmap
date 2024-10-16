package syncmap_test

// mapInterface below is from:
// https://cs.opensource.google/go/go/+/refs/tags/go1.23.2:src/sync/map_reference_test.go
// but with generic params.

import (
	"sync"
)

// This file contains reference map implementations for unit-tests.

// mapInterface is the interface Map implements.
type mapInterface[K comparable] interface {
	Load(key K) (value any, ok bool)
	Store(key K, value any)
	LoadOrStore(key K, value any) (actual any, loaded bool)
	LoadAndDelete(key K) (value any, loaded bool)
	Delete(K)
	Swap(key K, value any) (previous any, loaded bool)
	Range(func(key K, value any) (shouldContinue bool))
	Clear()
}

type StdMap[K comparable, V any] struct {
	m sync.Map
}

func (m *StdMap[K, V]) Load(key K) (value V, ok bool) {
	v, ok := m.m.Load(key)
	vv, _ := v.(V)
	return vv, ok
}

func (m *StdMap[K, V]) Store(key K, value V) {
	m.m.Store(key, value)
}

func (m *StdMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	a, loaded := m.m.LoadOrStore(key, value)
	aa, _ := a.(V)
	return aa, loaded
}

func (m *StdMap[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	v, loaded := m.m.LoadAndDelete(key)
	vv, _ := v.(V)
	return vv, loaded
}

func (m *StdMap[K, V]) Delete(key K) {
	m.m.Delete(key)
}

func (m *StdMap[K, V]) Range(f func(key K, value V) bool) {
	m.m.Range(func(k, v any) bool {
		kk := k.(K)
		vv, _ := v.(V)
		return f(kk, vv)
	})
}

func (m *StdMap[K, V]) Swap(k K, v V) (previous V, loaded bool) {
	p, loaded := m.m.Swap(k, v)
	pp, _ := p.(V)
	return pp, loaded
}

func (m *StdMap[K, V]) Clear() {
	m.m.Clear()
}
