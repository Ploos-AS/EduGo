# 26. Bygg og distribusjon

## Læringsmål
Forstå `GOOS`, `GOARCH`, `CGO_ENABLED`, build metadata og forskjellen mellom ren Go og programmer som trenger cgo.

## Oppdrag
Bygg samme CLI for Linux amd64, Linux arm64, Windows amd64 og macOS arm64 der kildekoden tillater det.

Undersøk binærfilene med `file` på Linux. Ikke lov "statisk binær" uten å verifisere egenskapene til den faktiske bygningen.
