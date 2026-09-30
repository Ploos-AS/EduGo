# 62. Buffering og pipes

Buffering kan redusere systemkall, men påvirker latency og når data faktisk blir synlige. Pipes kobler produsent og konsument som en strøm.

## Oppdrag
Sammenlign direkte skriving med `bufio.Writer`, og bygg deretter en `io.Pipe` der producer og consumer kjører samtidig.
