# 100. Capstone: Go Service

Sluttprosjektet samler hele EduGo i én liten produksjonsorientert tjeneste: **Job Service**.

Tjenesten tar imot jobber over HTTP, behandler dem med en begrenset worker pool og eksponerer status.

## Minimum
- `POST /api/v1/jobs`
- `GET /api/v1/jobs/{id}`
- `/healthz` og `/readyz`
- bounded request body og validering
- concurrency-safe store
- bounded queue og worker pool
- context, timeouts og graceful shutdown
- strukturert logging
- tester uten eksterne tjenester

Bygg først korrekthet. Mål før du optimaliserer.
