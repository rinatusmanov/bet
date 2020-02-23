package task1

import "testing"

func BenchmarkLoadVariables(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LoadVariables()
	}
}
func BenchmarkToPostgresFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToPostgresFunc("user/1", "POST")
	}
}
func BenchmarkToPostgresFunc1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToPostgresFunc("user/1", "POST")
	}
}
func BenchmarkToPostgresFunc2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToPostgresFunc("user/1/kinder/2", "PUT")
	}
}
func BenchmarkToPostgresFunc3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToPostgresFunc("user/1/kinder/2/comment", "POST")
	}
}

var (
	benchmarkPostgresParamsMixin       []interface{}
	benchmarkPostgresParamsMixinSmall  = []interface{}{1, 2, 3}
	benchmarkPostgresParamsMixinMiddle = []interface{}{1, 2, 3, 12412, 1, 241, 24, 24, 24, 2, 4, 3, 2, 5, 235, 1, 2, 3, 12412, 1, 241, 24, 24, 24, 2, 4, 3, 2, 5, 235, 1, 2, 3, 12412, 1, 241, 24, 24, 24, 2, 4, 3, 2, 5, 235}
	benchmarkPostgresParamsMixinLarge  = []interface{}{1, 2, 3, 12412, 1, 241, 24, 24, 24, 2, 4, 3, 2, 5, 235, 1, 2, 3, 12412, 1, 241, 24, 24, 24, 2, 4, 3, 2, 5, 235, 1, 2, 3, 12412, 1, 241, 24, 24, 24, 2, 4, 3, 2, 5, 235, 1, 2, 3, 12412, 1, 241, 24, 24, 24, 2, 4, 3, 2, 5, 235, 1, 2, 3, 12412, 1, 241, 24, 24, 24, 2, 4, 3, 2, 5, 235}
)

func BenchmarkPostgresParamsMixin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PostgresParamsMixin(benchmarkPostgresParamsMixin)
	}
}

func BenchmarkPostgresParamsMixinSmall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PostgresParamsMixin(benchmarkPostgresParamsMixinSmall)
	}
}

func BenchmarkPostgresParamsMixinMiddle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PostgresParamsMixin(benchmarkPostgresParamsMixinMiddle)
	}
}

func BenchmarkPostgresParamsMixinLarge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PostgresParamsMixin(benchmarkPostgresParamsMixinLarge)
	}
}
