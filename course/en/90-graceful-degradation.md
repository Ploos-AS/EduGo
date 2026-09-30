# 90. Graceful degradation

Not every dependency failure needs to make the whole service unavailable.

## Mission
Classify dependencies as required or optional. Fail an optional dependency in a controlled way and expose degraded behavior through logs, metrics, or readiness where appropriate.
