# 87. Health and readiness

Liveness asks whether the process is alive. Readiness asks whether it should receive traffic. They are different questions.

## Mission
Create separate `/healthz` and `/readyz` endpoints and let readiness reflect controlled dependency state without making liveness depend on external services.
