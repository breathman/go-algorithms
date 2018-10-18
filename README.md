## package search
```
go test ./search -bench=. -count=2
```

##### benchmarks

###### for slice with 100 elements

| Name                             | Times     | Speed       |
|----------------------------------|-----------|-------------|
| BenchmarkBinarySearchIterative-8 |   100     |  37.7 ns/op |
| BenchmarkBinarySearchIterative-8 |   100     |  39.9 ns/op |
| BenchmarkBinarySearchRecursive-8 |   100     |  58.5 ns/op |
| BenchmarkBinarySearchRecursive-8 |   100     |  57.5 ns/op |
| BenchmarkLinearSearchUnsorted-8  |   100     |  35.7 ns/op |
| BenchmarkLinearSearchUnsorted-8  |   100     |  36.5 ns/op |
| BenchmarkLinearSearchSorted-8    |   100     |  44.3 ns/op |
| BenchmarkLinearSearchSorted-8    |   100     |  45.2 ns/op |

###### for slice with 10000 elements

| Name                             | Times     | Speed       |
|----------------------------------|-----------|-------------|
| BenchmarkBinarySearchIterative-8 |  10000    | 52.4 ns/op  |
| BenchmarkBinarySearchIterative-8 |  10000    | 52.6 ns/op  |
| BenchmarkBinarySearchRecursive-8 |  10000    | 89.4 ns/op  |
| BenchmarkBinarySearchRecursive-8 |  10000    | 88.7 ns/op  |
| BenchmarkLinearSearchUnsorted-8  |  10000    | 1468 ns/op  |
| BenchmarkLinearSearchUnsorted-8  |  10000    | 1448 ns/op  |
| BenchmarkLinearSearchSorted-8    |  10000    | 2380 ns/op  |
| BenchmarkLinearSearchSorted-8    |  10000    | 2365 ns/op  |
