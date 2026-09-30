# 44. Release og distribusjon

En release er mer enn `go build`.

Definer støttede plattformer, bygg artefakter deterministisk så langt praktisk, lag checksums og test installasjonen fra et rent miljø.

## Oppdrag
Lag en lokal release-prosess for et CLI-program med flere GOOS/GOARCH-mål og SHA-256 checksums.
