# graphql-dataloader-benchmark

Framework
 - GraphQL: [graph-gophers/graphql-go](https://github.com/graph-gophers/graphql-go)
 - Dataloader: [dataloader](https://github.com/graph-gophers/dataloader)

Execute resolvers per request in parallel by 50 goroutines, here is benchmark result

```
go test -bench=. -benchmem -benchtime=30s
BenchmarkTestWithLoader-8      	     100	 587755146 ns/op	13771534 B/op	  317360 allocs/op
BenchmarkTestWithoutLoader-8   	     100	 602792219 ns/op	14162485 B/op	  325612 allocs/op
PASS
ok  	github.com/jiazhen-lin/graphql-dataloader-benchmark	123.330s
```

Use loader to reduce DB query count efficiently

```
BenchmarkTestWithLoader-8      	     Query count: 1047, Cost time: 37725346000 (nano), Ave time: 36031849.092646 (nano)
BenchmarkTestWithoutLoader-8   	     Query count: 50221, Cost time: 2698396740000 (nano), Ave time: 53730446.227674 (nano)
```
