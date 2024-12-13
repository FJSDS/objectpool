package objectpool

import (
	"sync"
	"testing"
	"unsafe"
)

type Struct struct {
	A unsafe.Pointer
}

func BenchmarkGetPut(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Put(Get[Struct]())
	}
}

func BenchmarkGetSlicePutSlice(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PutSlice(GetSlice[Struct]())
	}
}

func BenchmarkGetMapPutMap(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PutMap(GetMap[int, Struct]())
	}
}

var p = sync.Pool{
	New: func() interface{} {
		return new(Struct)
	},
}

func BenchmarkPool(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		p.Put(p.Get())
	}
}
