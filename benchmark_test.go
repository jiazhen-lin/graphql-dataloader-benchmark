package main

import (
	"context"
	"testing"

	graphql "github.com/graph-gophers/graphql-go"
	gql "github.com/jiazhen-lin/graphql-dataloader-benchmark/gql"
)

func BenchmarkTestWithoutLoader(b *testing.B) {
	gql.CreateTestData()
	b.ResetTimer()

	s := graphql.MustParseSchema(gql.Schema, &gql.Resolver{}, graphql.MaxParallelism(50))
	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		ctx = gql.Attach(ctx, false)
		s.Exec(ctx, `query(){
			users{
				name
				posts{
					text
				}
			}
		}`, "", nil)
	}
}

func BenchmarkTestWithLoader(b *testing.B) {
	gql.CreateTestData()
	b.ResetTimer()

	s := graphql.MustParseSchema(gql.Schema, &gql.Resolver{}, graphql.MaxParallelism(50))
	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		ctx = gql.Attach(ctx, true)
		s.Exec(ctx, `query(){
			users{
				name
				posts{
					text
				}
			}
		}`, "", nil)
	}
}
