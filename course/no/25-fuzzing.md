# 25. Fuzzing

Go har innebygd fuzzing i testverktøyet.

## Oppdrag
Lag en fuzz-test for `ParsePort`. Programmet skal aldri panikke for vilkårlig tekst, og enhver akseptert port skal ligge mellom 1 og 65535.

Start med gode seed-cases og la fuzzer utforske resten.
