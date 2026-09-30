# 69. Secrets og sensitiv konfigurasjon

Secrets skal ikke hardkodes, committes eller skrives til logger.

Les dem ved prosessgrensen, hold levetiden enkel og unngå å spre dem gjennom unødvendige datastrukturer.

## Oppdrag
Lag en config-loader som krever en secret fra environment og returnerer en tydelig feil når den mangler, uten å inkludere secret-verdien i feilen.
