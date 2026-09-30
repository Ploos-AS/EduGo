# M8 — Protocols and distributed systems

M8 introduces protocol framing, bounded retries, idempotency, deadline propagation, rate limiting, backpressure, and partial failure.

The first exercise implements retry mechanics with injected waiting so tests stay fast and deterministic. Retry policy remains the caller's responsibility: the helper does not guess whether an operation is safe to repeat.
