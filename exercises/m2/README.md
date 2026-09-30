# M2 — Concurrency

M2 focuses on correctness before performance.

1. Goroutines and lifetime
2. Channels
3. select and context cancellation
4. Worker pools
5. Data races and synchronization
6. Race detector

## Completion gate

Run:

```sh
go test ./...
go test -race ./...
```

The final exercise is a cancellable worker pool. Students should be able to explain who owns each channel, who closes it, how workers terminate, and why output ordering is not guaranteed.
