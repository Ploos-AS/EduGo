# 87. Health og readiness

Liveness svarer på om prosessen lever. Readiness svarer på om den bør motta trafikk. De er ikke samme spørsmål.

## Oppdrag
Lag separate `/healthz` og `/readyz`. La readiness reflektere en kontrollert dependency-tilstand uten å gjøre liveness avhengig av eksterne tjenester.
