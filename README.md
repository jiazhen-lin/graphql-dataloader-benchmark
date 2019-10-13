# graphql-dataloader-benchmark

Framework
 - GraphQL: [graph-gophers/graphql-go](https://github.com/graph-gophers/graphql-go)
 - Dataloader: [dataloader](https://github.com/graph-gophers/dataloader)

Execute resolvers per request in parallel by 50 goroutines, here is benchmark result

```
go test -bench=. -benchmem -benchtime=30s
BenchmarkTestWithLoader-8      	     100	 523596938 ns/op	 1867980 B/op	   44451 allocs/op
BenchmarkTestWithoutLoader-8   	     100	 561794184 ns/op	 2422942 B/op	   53185 allocs/op
PASS
ok  	github.com/jiazhen-lin/graphql-dataloader-benchmark	123.330s
```

Use loader to reduce DB query count efficiently

```
BenchmarkTestWithLoader-8      	     Query count: 1010, Cost time: 32243951000 (nano), Ave time: 31924703.960396 (nano)
BenchmarkTestWithoutLoader-8   	     Query count: 50406, Cost time: 2498154653000 (nano), Ave time: 49560660.496766 (nano)
```
