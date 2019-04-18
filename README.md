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

###### for slice with 100 elements with data range from 0 to 10000

| Name                             | Times     | Speed        | BytesPerOp   | AllocsPerOp     |
|----------------------------------|-----------|--------------|--------------|-----------------|
| BenchmarkBubbleSort-8            |  200000   |  6637  ns/op |    0 B/op    | 0 allocs/op     |
| BenchmarkBucketSort-8            |  100000   |  12365 ns/op | 81920 B/op   | 1 allocs/op     |
| BenchmarkInsertionSort-8         |  500000   |  2995  ns/op |    0 B/op    | 0 allocs/op     |
| BenchmarkMergeSort-8             |  500000   |  2616  ns/op | 896 B/op     | 1 allocs/op     |
| BenchmarkQuickSort-8             |  100000   |  18201 ns/op | 19208 B/op   | 371 allocs/op   |
| BenchmarkSelectionSort-8         |  300000   |  5433  ns/op |    0 B/op    | 0 allocs/op     |


###### for slice with 10000 elements with data range from 0 to 10000

| Name                             | Times     | Speed           | BytesPerOp   | AllocsPerOp     |
|----------------------------------|-----------|-----------------|--------------|-----------------|
| BenchmarkBubbleSort-8            |   20      | 52669694  ns/op |    0 B/op    | 0 allocs/op     |
| BenchmarkBucketSort-8            |   20000   | 92059     ns/op | 81920 B/op   | 1 allocs/op     |
| BenchmarkInsertionSort-8         |   50      | 25950356  ns/op |    0 B/op    | 0 allocs/op     |
| BenchmarkMergeSort-8             |   3000    | 412066    ns/op | 81920 B/op   | 1 allocs/op     |
| BenchmarkQuickSort-8             |   500     | 3171612   ns/op | 4907722 B/op | 32756 allocs/op |
| BenchmarkSelectionSort-8         |   10      | 108949321 ns/op |    0 B/op    | 0 allocs/op     |