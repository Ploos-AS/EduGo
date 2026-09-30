# 56. HTTP transport under panseret

Studer forbindelsesgjenbruk, keep-alive, request bodies, response bodies og hvorfor klienter bør gjenbruke `http.Client`/Transport fremfor å lage ny transport for hvert kall.

## Oppdrag
Bruk `httptest.Server` og instrumentering til å undersøke flere requests mot samme lokale server.
