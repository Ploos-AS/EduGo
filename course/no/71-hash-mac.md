# 71. Hashing og MAC

En hash gir ikke autentisitet. HMAC kombinerer en hemmelig nøkkel med en hash for å autentisere data.

## Oppdrag
Signer en melding med HMAC-SHA-256 og verifiser den med konstant-tid-sammenligning gjennom standardbibliotekets API.

Ikke bruk vanlig hash som passordlagring; passordlagring krever en dedikert password-hashing/KDF-løsning.
