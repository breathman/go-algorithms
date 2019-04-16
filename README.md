## package search
```
go test ./search -bench=. -benchtime=1000x
```

##### benchmarks

###### for slice with 100 elements

| Name                             | Times     | Speed       |
|----------------------------------|-----------|-------------|
| BenchmarkBinarySearchIterative-8 |   1000    |  14.5 ns/op |
| BenchmarkBinarySearchRecursive-8 |   1000    |  24.9 ns/op |
| BenchmarkLinearSearchUnsorted-8  |   1000    |  53.5 ns/op |
| BenchmarkLinearSearchSorted-8    |   1000    |  39.1 ns/op |


###### for slice with 10000 elements

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

| Name                             | Times     | Speed        |
|----------------------------------|-----------|--------------|
| BenchmarkBubbleSort-8            |  200000   |  6637 ns/op  |
| BenchmarkInsertionSort-8         |  500000   |  2995 ns/op  |
| BenchmarkQuickSort-8             |  100000   |  18201 ns/op |
| BenchmarkSelectionSort-8         |  300000   |  5433 ns/op  |


###### for slice with 10000 elements

| Name                             | Times     | Speed           |
|----------------------------------|-----------|-----------------|
| BenchmarkBubbleSort-8            |   20      | 52669694 ns/op  |
| BenchmarkInsertionSort-8         |   50      | 25950356 ns/op  |
| BenchmarkQuickSort-8             |   500     | 3171612 ns/op   |
| BenchmarkSelectionSort-8         |   10      | 108949321 ns/op |