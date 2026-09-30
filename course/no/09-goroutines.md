# 9. Goroutines

En goroutine er en funksjon som kjører samtidig med annet arbeid i programmet.

## Læringsmål
Start goroutines med `go`, forstå at programslutt avslutter gjenværende goroutines, og bruk `sync.WaitGroup` når du må vente på arbeid.

## Oppdrag
Start fem arbeidere som utfører hver sin beregning. Vent eksplisitt på alle før programmet avsluttes.

## Tenk
En goroutine er ikke det samme som en OS-tråd. Hvorfor kan denne forskjellen være nyttig?
