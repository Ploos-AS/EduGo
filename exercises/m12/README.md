# M12 — Performance engineering

M12 teaches measurement-driven optimization: benchmarks, allocations and escape, buffering/batching, `sync.Pool`, contention, atomics, profiling, and evidence-based optimization decisions.

`linebuild` supplies two intentionally different implementations with identical output. Run:

```sh
go test -bench=. -benchmem
```

The lesson is not that one technique always wins; students must measure their workload and justify added complexity.
