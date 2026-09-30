# 13. Data races and synchronization

Run `go test -race ./...`.

## Mission
Intentionally create a shared counter with a data race. Confirm that the race detector reports it, then fix it.

Compare `sync.Mutex`, atomics, and a design where one goroutine owns the mutable state.

A program appearing to work is not evidence that concurrent code is correct.
