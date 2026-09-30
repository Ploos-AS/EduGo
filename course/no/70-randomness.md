# 70. Sikker randomness

Bruk `crypto/rand` når verdier må være uforutsigbare. `math/rand` er nyttig for simulering og testing, men er ikke en erstatning for kryptografisk randomness.

## Oppdrag
Generer et tilfeldig token med fast antall bytes og kod det med URL-safe base64.
