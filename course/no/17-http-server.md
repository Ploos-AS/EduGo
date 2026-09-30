# 17. HTTP-servere

## Læringsmål
Bruk `http.Handler`, `ServeMux`, request context og JSON-responser.

## Oppdrag
Lag en liten tjeneste med `GET /healthz` og `GET /api/v1/echo?text=...`.

Test handlerne med `httptest`; testene skal ikke trenge en ekte nettverksport.
