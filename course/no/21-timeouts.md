# 21. Timeouts, deadlines og grenser

Et nettverksprogram skal ikke kunne vente ubegrenset på omverdenen.

Sett eksplisitte server- og klienttimeouts. Bruk context-deadlines for arbeid som skal kunne avbrytes. Begrens hvor mye input som leses når input kommer utenfra.

## Oppdrag
Legg timeouts på HTTP-serveren og klienten. Lag tester som simulerer en treg handler.
