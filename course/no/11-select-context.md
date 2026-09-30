# 11. select og context

`select` venter på flere channel-operasjoner. `context.Context` brukes blant annet til cancellation og deadlines.

## Oppdrag
Lag en arbeidsgoroutine som sender periodiske resultater, men stopper når context blir kansellert.

## Viktig
Send `context.Context` som parameter. Ikke lagre den unødvendig i structs, og sørg for at goroutines faktisk kan avsluttes.
