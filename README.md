## package search
```
go test ./search -bench=.
```

##### benchmarks

###### for slice with 100 elements

| Name                             | Times     | Speed       |
|----------------------------------|-----------|-------------|
| BenchmarkBinarySearchIterative-8 |   1000    |  14.5 ns/op |
| BenchmarkBinarySearchRecursive-8 |   1000    |  24.9 ns/op |
| BenchmarkLinearSearchUnsorted-8  |   1000    |  53.5 ns/op |
| BenchmarkLinearSearchSorted-8    |   1000    |  39.1 ns/op |


###### for slice with 1000 elements

| Name                             | Times     | Speed       |
|----------------------------------|-----------|-------------|
| BenchmarkBinarySearchIterative-8 |   1000    |  21.1 ns/op |
| BenchmarkBinarySearchRecursive-8 |   1000    |  62.4 ns/op |
| BenchmarkLinearSearchUnsorted-8  |   1000    |  287 ns/op  |
| BenchmarkLinearSearchSorted-8    |   1000    |  199 ns/op  |
        


## package sort         
```
go test ./sort -bench=.
```

###### for slice with 100 elements

| Name                             | Times     | Speed       |
|----------------------------------|-----------|-------------|
| BenchmarkBubbleSort-8            |   1000    |  6899 ns/op |
| BenchmarkQuickSort-8             |   1000    | 14013 ns/op |


###### for slice with 1000 elements

| Name                             | Times     | Speed        |
|----------------------------------|-----------|--------------|
| BenchmarkBubbleSort-8            |   1000    | 483946 ns/op |
| BenchmarkQuickSort-8             |   1000    | 134578 ns/op |