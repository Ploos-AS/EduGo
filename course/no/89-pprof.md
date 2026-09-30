# 89. pprof i drift

`net/http/pprof` kan gi svært nyttig diagnostikk, men diagnostikk-endepunkter skal behandles som administrative flater.

## Oppdrag
Aktiver pprof kun på en separat lokal/admin listener i laben. Dokumenter hvorfor den ikke bør eksponeres ukritisk på tjenestens offentlige listener.
