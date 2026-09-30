# 79. Pooling

`sync.Pool` can reduce pressure from temporary objects, but it is not a general object cache and entries may disappear at any time.

## Mission
Benchmark a temporary buffer with and without a pool. Keep pooling only when measurements show a relevant benefit.
