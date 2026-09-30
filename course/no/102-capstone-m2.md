# 102. Capstone M2 — Concurrency

Legg til bounded queue og et eksplisitt antall workers.

## Akseptanse
Ingen ubegrenset goroutine-vekst, queue-full håndteres eksplisitt, race detector er grønn, og shutdown etterlater ikke arbeid i en uklar tilstand.
