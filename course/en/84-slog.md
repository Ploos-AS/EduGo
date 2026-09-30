# 84. Structured logging with slog

Logs should help people and tools understand events without parsing free-form prose. Use stable messages and structured attributes, and never casually log secrets or large request bodies.

## Mission
Log an HTTP request with method, path, status, duration, and request ID. Distinguish low- from high-cardinality fields.
