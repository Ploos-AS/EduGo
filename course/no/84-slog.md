# 84. Strukturert logging med slog

Logger skal hjelpe mennesker og verktøy å forstå hendelser uten å måtte parse fritekst.

Bruk stabile meldinger og strukturerte attributter. Ikke logg secrets eller store request bodies ukritisk.

## Oppdrag
Logg en request med metode, path, status, varighet og request-ID. Skill felter med lav og høy kardinalitet.
