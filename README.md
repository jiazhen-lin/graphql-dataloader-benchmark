# graphql-dataloader-benchmark

Framework
 - GraphQL: [graph-gophers/graphql-go](!https://github.com/graph-gophers/graphql-go)
 - Dataloader: [dataloader](!https://github.com/graph-gophers/dataloader)

Benchmark Results

```
go test -bench=. -benchmem -benchtime=30s

BenchmarkTestWithLoader-8      	     100	 587755146 ns/op	13771534 B/op	  317360 allocs/op
BenchmarkTestWithoutLoader-8   	     100	 602792219 ns/op	14162485 B/op	  325612 allocs/op
PASS
ok  	github.com/jiazhen-lin/graphql-dataloader-benchmark	123.330s
```
