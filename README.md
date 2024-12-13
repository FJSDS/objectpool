# objectpool
Generic Object Pool,thread safe

```
goos: windows
goarch: amd64
pkg: github.com/FJSDS/objectpool
cpu: AMD Ryzen 5 5600 6-Core Processor              
BenchmarkGetPut
BenchmarkGetPut-12              	73732718	        14.90 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetSlicePutSlice
BenchmarkGetSlicePutSlice-12    	80940528	        15.18 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetMapPutMap
BenchmarkGetMapPutMap-12        	74947536	        16.18 ns/op	       0 B/op	       0 allocs/op
BenchmarkPool
BenchmarkPool-12                	100000000	        10.03 ns/op	       0 B/op	       0 allocs/op
PASS

Process finished with the exit code 0
```