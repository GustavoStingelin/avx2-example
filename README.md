# AVX2 Example

This repository contains simple examples showing how to use Go assembly (via the
[Avo](https://github.com/mmcloughlin/avo) generator) and AVX2 instructions. It
includes implementations of basic addition functions and a reference SHA256
implementation used for benchmarks.

## Layout

- `pkg/asm`: Functions written in assembly. The `add` and `addslice` examples use
  Avo to generate the `.s` files. Generated stubs are committed so tests can run
  without regeneration.
- `pkg/gostd`: Pure Go equivalents of the assembly routines plus a minimal
  SHA256 implementation.
- Benchmarks live in `*_bench_test.go` at the repository root.

## Running

```bash
go test ./...
```

To run benchmarks:

```bash
go test -bench=.
```

If you want to regenerate the assembly output (requires Go and AVX2 capable
hardware), run:

```bash
go generate ./...
```

## Benchmark Example

Running the benchmarks on an AMD64 system produced output similar to:

```
goos: linux
goarch: amd64
pkg: avx2-example
cpu: AMD Ryzen 7 5700X 8-Core Processor             
BenchmarkAdd/asm-16              913769707                1.298 ns/op
BenchmarkAdd/gostd-16           1000000000               0.2164 ns/op
BenchmarkAddSlice/asm-16          16372576                71.77 ns/op
BenchmarkAddSlice/gostd-16         2684764                612.3 ns/op
BenchmarkSHA256/asm-16             2347438                511.0 ns/op
BenchmarkSHA256/stdgo-16            289142                 4034 ns/op
PASS
ok      avx2-example    7.837s
```

These numbers compare the assembly versions with their pure Go counterparts.

