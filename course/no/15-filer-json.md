# 15. Filer, I/O og JSON

## Læringsmål
Bruk `io.Reader` og `io.Writer`, arbeid med filer via `os`, og kod/dekod JSON med `encoding/json`.

## Oppdrag
Lag en adressebok som kan lese JSON fra fil og skrive resultatet tilbake.

## Design
La kjernelogikken ta `io.Reader`/`io.Writer` når det passer. Da blir den enklere å teste enn kode som alltid åpner bestemte filer selv.
