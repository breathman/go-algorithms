## package search
```
go test ./search -bench=. -count=2 -short -v
```

##### benchmarks

| Name                             | Times     | Speed       |
|----------------------------------|-----------|-------------|
| BenchmarkBinarySearchIterative-8 |200000000  |  9.95 ns/op |
| BenchmarkBinarySearchIterative-8 |100000000  | 15.3 ns/op  |
| BenchmarkBinarySearchRecursive-8 |100000000  | 33.5 ns/op  |
| BenchmarkBinarySearchRecursive-8 |100000000  | 32.9 ns/op  |
| BenchmarkLinearSearchUnsorted-8  |300000000  | 10.1 ns/op  |
| BenchmarkLinearSearchUnsorted-8  |100000000  | 13.7 ns/op  |
| BenchmarkLinearSearchSorted-8    |50000000   | 24.2 ns/op  |
| BenchmarkLinearSearchSorted-8    |20000000   | 72.5 ns/op  |
