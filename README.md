# syncmap


[![Go Reference](https://pkg.go.dev/badge/github.com/graxinc/syncmap.svg)](https://pkg.go.dev/github.com/graxinc/syncmap)

## Purpose

`syncmap.Map` is [sync.Map](https://pkg.go.dev/sync#Map) using generics.

`syncmap.Map` has the same methods minus `CompareAndSwap` and `CompareAndDelete`, since they require comparable values. Let us know if you want support for them added in some way!

## Benchmark

Results for `sync.Map` benchmarks (found [here](https://cs.opensource.google/go/go/+/refs/tags/go1.23.2:src/sync/map_bench_test.go)):

```
BenchmarkLoadMostlyHits/std-8       	215163644	         6.043 ns/op
BenchmarkLoadMostlyHits/syncmap-8   	391649406	         3.078 ns/op
BenchmarkLoadMostlyMisses/std-8     	361615671	         3.417 ns/op
BenchmarkLoadMostlyMisses/syncmap-8 	434882744	         4.209 ns/op
BenchmarkLoadOrStoreBalanced/std-8  	 4030590	       406.5 ns/op
BenchmarkLoadOrStoreBalanced/syncmap-8         	 4442049	       283.4 ns/op
BenchmarkLoadOrStoreUnique/std-8               	 2159490	       652.7 ns/op
BenchmarkLoadOrStoreUnique/syncmap-8           	 2716921	       409.5 ns/op
BenchmarkLoadOrStoreCollision/std-8            	216356244	         5.754 ns/op
BenchmarkLoadOrStoreCollision/syncmap-8        	492943695	         2.385 ns/op
BenchmarkLoadAndDeleteBalanced/std-8           	206500206	         5.473 ns/op
BenchmarkLoadAndDeleteBalanced/syncmap-8       	344110657	         3.502 ns/op
BenchmarkLoadAndDeleteUnique/std-8             	470197647	         2.575 ns/op
BenchmarkLoadAndDeleteUnique/syncmap-8         	874849255	         1.350 ns/op
BenchmarkLoadAndDeleteCollision/std-8          	286711022	         5.346 ns/op
BenchmarkLoadAndDeleteCollision/syncmap-8      	393698605	         2.964 ns/op
BenchmarkRange/std-8                           	  300404	      3879 ns/op
BenchmarkRange/syncmap-8                       	  336896	      4539 ns/op
BenchmarkAdversarialAlloc/std-8                	 5690133	       218.6 ns/op
BenchmarkAdversarialAlloc/syncmap-8            	 7773882	       154.1 ns/op
BenchmarkAdversarialDelete/std-8               	20595362	        60.65 ns/op
BenchmarkAdversarialDelete/syncmap-8           	23633935	        50.75 ns/op
BenchmarkDeleteCollision/std-8                 	424353769	         2.827 ns/op
BenchmarkDeleteCollision/syncmap-8             	809968076	         1.695 ns/op
BenchmarkSwapCollision/std-8                   	 8332641	       155.7 ns/op
BenchmarkSwapCollision/syncmap-8               	 9771920	       115.6 ns/op
BenchmarkSwapMostlyHits/std-8                  	40996706	        39.66 ns/op
BenchmarkSwapMostlyHits/syncmap-8              	33398676	        33.78 ns/op
BenchmarkSwapMostlyMisses/std-8                	 2160728	       600.4 ns/op
BenchmarkSwapMostlyMisses/syncmap-8            	 2756742	       435.5 ns/op
BenchmarkClear/std-8                           	 3368053	       386.2 ns/op
BenchmarkClear/syncmap-8                       	 3294117	       398.6 ns/op
```
