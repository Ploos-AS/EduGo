# 57. TLS i praksis

Bruk `httptest.NewTLSServer` for lokal HTTPS-testing. Studer trust, hostname-verifisering og klientkonfigurasjon uten å slå av verifisering som snarvei.

## Oppdrag
La testserveren levere HTTPS og bruk dens testklient til å verifisere en request.
