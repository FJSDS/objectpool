package objectpool

import (
	"math"
	"sync"
	"unsafe"
)

const (
	maxIndex = math.MaxUint16 - 1
)

var op = objectPool{}

type poolUintptr struct {
	uintptr
	*sync.Pool
}

type objectPool struct {
	m  [math.MaxUint16][]*poolUintptr
	ml [math.MaxUint16]sync.Mutex
}

func (o *objectPool) get(p uintptr) *sync.Pool {
	index := (p >> 6) & maxIndex

	ss := o.m[index]
	for _, s := range ss {
		if s.uintptr == p {
			return s.Pool
		}
	}

	// lock for index conflict,
	o.ml[index].Lock()
	for _, s := range o.m[index] {
		if s.uintptr == p {
			o.ml[index].Unlock()
			return s.Pool
		}
	}
	pu := &poolUintptr{
		uintptr: p,
		Pool:    &sync.Pool{},
	}
	o.m[index] = append(o.m[index], pu)
	o.ml[index].Unlock()
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
