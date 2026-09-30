# 24. Benchmarks og profilering

Optimaliser etter måling.

Bruk `go test -bench=.`, `-benchmem` og profileringsverktøy som `pprof`.

## Oppdrag
Benchmark en funksjon før og etter en bevisst endring. Rapporter både tid og allokeringer.

## Regel
En mikrobenchmark er ikke automatisk representativ for et helt system. Dokumenter hva du faktisk målte.
