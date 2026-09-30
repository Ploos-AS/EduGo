# 20. Prosesslivssyklus og graceful shutdown

Produksjonsprogrammer må kunne stoppe kontrollert.

## Læringsmål
Håndter OS-signaler med `signal.NotifyContext`, stopp nytt arbeid, la pågående arbeid avsluttes innen en frist og frigjør ressurser.

## Oppdrag
Utvid miniservice slik at SIGINT/SIGTERM starter graceful shutdown med en eksplisitt timeout.

## Tenk
Hva kan skje med klienter og data dersom prosessen bare termineres umiddelbart?
