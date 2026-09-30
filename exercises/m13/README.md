# M13 — Observability and operations

M13 covers structured logging, metrics, tracing concepts, liveness/readiness, runtime diagnostics, pprof, graceful degradation, and an operations lab.

The `health` exercise makes one important operational distinction executable: a process may be alive while deliberately not ready to receive traffic. Its state uses `atomic.Bool`, so the exercise also reconnects to earlier concurrency material.
