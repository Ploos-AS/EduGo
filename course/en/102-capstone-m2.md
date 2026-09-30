# 102. Capstone M2 — Concurrency

Add a bounded queue and an explicit worker count.

## Acceptance
No unbounded goroutine growth, queue-full behavior is explicit, the race detector passes, and shutdown does not leave work in an ambiguous state.
