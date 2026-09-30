# 100. Capstone: Go Service

The final project combines EduGo into one small production-oriented service: **Job Service**.

It accepts jobs over HTTP, processes them through a bounded worker pool, and exposes status.

## Minimum
- `POST /api/v1/jobs`
- `GET /api/v1/jobs/{id}`
- `/healthz` and `/readyz`
- bounded request body and validation
- concurrency-safe store
- bounded queue and worker pool
- context, timeouts, and graceful shutdown
- structured logging
- tests without external services

Build for correctness first. Measure before optimizing.
