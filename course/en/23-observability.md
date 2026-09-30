# 23. Observability and slog

Use `log/slog` for structured logs and request context for useful fields.

## Mission
Log method, path, status, duration, and request ID. Do not indiscriminately log secrets or request bodies.

Logs, metrics, and traces are distinct signals; this lesson concentrates on logging and local measurement.
