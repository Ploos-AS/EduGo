# 80. Contention og skalering

Flere goroutines betyr ikke automatisk mer throughput. Locks, shared state, scheduler overhead og eksterne flaskehalser kan dominere.

## Oppdrag
Benchmark en delt teller ved ulike parallelism-nivåer og identifiser hvor skaleringen stopper.
