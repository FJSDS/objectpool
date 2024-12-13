# objectpool
Generic Object Pool

```
goos: windows
goarch: amd64
pkg: github.com/FJSDS/objectpool
cpu: AMD Ryzen 5 5600 6-Core Processor              
BenchmarkGetPut
BenchmarkGetPut-12              	86345842	        13.84 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetSlicePutSlice
BenchmarkGetSlicePutSlice-12    	85620714	        14.19 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetMapPutMap
BenchmarkGetMapPutMap-12        	68151614	        17.33 ns/op	       0 B/op	       0 allocs/op
BenchmarkPool
BenchmarkPool-12                	100000000	        10.21 ns/op	       0 B/op	       0 allocs/op
PASS

Process finished with the exit code 0
```