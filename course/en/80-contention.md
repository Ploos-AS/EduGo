# 80. Contention and scaling

More goroutines do not automatically mean more throughput. Locks, shared state, scheduler overhead, and external bottlenecks can dominate.

## Mission
Benchmark a shared counter at multiple parallelism levels and identify where scaling stops.
