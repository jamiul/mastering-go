# Benchmarking
Writing benchmarks in Go is another first-class feature of the language and it is very similar to writing tests.

```go
func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a")
	}
}
```

```bash
go test -bench=.
```

To run the benchmarks do go test -bench=. (or if you're in Windows Powershell go test -bench=".")

```bash
go test -bench=. -benchm
```

We can use BenchmarkRepeat to confirm that strings.Builder significantly improves performance.
Run go test -bench=. -benchmem