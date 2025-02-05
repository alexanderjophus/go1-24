# Go 1.24 Bytesize Updates

## Generic Type Aliases

// todo

## Tool command

```bash
go get -tool golang.org/x/tools/cmd/stringer
go tool stringer -type=Pill
```

stringer is used as an example, it could just as easily be swagger spec generation, protofile creation, etc.

## Benchmarking

A slightly better way to write benchmark tests.

```diff
func BenchmarkHello(b *testing.B) {
	// set up
+	for b.Loop() {
+		sum(1, 2, 3, 4, 5)
+	}
-	for i := 0; i < b.N; i++ {
-		sum(1, 2, 3, 4, 5)
-	}
	// tear down
}
```

```bash
go test -bench=.
```
