# 81. Atomics

`sync/atomic` is useful for small, well-defined concurrency problems but does not simplify complex invariants.

## Mission
Compare a simple mutex counter with an atomic counter and explain why the result does not justify replacing every mutex.
