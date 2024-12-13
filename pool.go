package objectpool

import (
	"math"
	"sync"
	"unsafe"
)

var op = objectPool{}

type PoolUintptr struct {
	uintptr
	*sync.Pool
}

type objectPool struct {
	m [math.MaxUint16][]*PoolUintptr
}

func (o *objectPool) get(p uintptr) *sync.Pool {
	p &= math.MaxUint16
	ss := o.m[p]
	for _, s := range ss {
		if s.uintptr == p {
			return s.Pool
		}
	}
	pu := &PoolUintptr{
		uintptr: p,
		Pool:    &sync.Pool{},
	}
	o.m[p] = append(o.m[p], pu)
	return pu.Pool
}

func Get[T any]() *T {
	var a interface{} = (*T)(nil)
	typPtr := *(*uintptr)(unsafe.Pointer(&a))
	p := op.get(typPtr)
	v := p.Get()
	if v != nil {
		return v.(*T)
	}
	return new(T)
}

func Put[T any](t *T) {
	var a interface{} = (*T)(nil)
	typPtr := *(*uintptr)(unsafe.Pointer(&a))
	p := op.get(typPtr)
	p.Put(t)
}

type Slice[T any] struct {
	Data []T
}

func GetSlice[T any]() *Slice[T] {
	return Get[Slice[T]]()
}

func PutSlice[T any](t *Slice[T]) {
	clear(t.Data)
	Put(t)
}

func GetMap[K comparable, V any]() map[K]V {
	var a interface{} = (map[K]V)(nil)
	typPtr := *(*uintptr)(unsafe.Pointer(&a))
	p := op.get(typPtr)
	v := p.Get()
	if v != nil {
		return v.(map[K]V)
	}
	return map[K]V{}
}

func PutMap[K comparable, V any](t map[K]V) {
	clear(t)
	var a interface{} = (map[K]V)(nil)
	typPtr := *(*uintptr)(unsafe.Pointer(&a))
	p := op.get(typPtr)
	p.Put(t)
}
