# 13. Data races og synkronisering

Samtidighet gjør delte mutable data farlige.

Kjør:

```sh
go test -race ./...
```

## Oppdrag
Lag med vilje en usikker teller som flere goroutines oppdaterer. Bekreft at race detector finner problemet. Rett den deretter med passende synkronisering.

Sammenlign `sync.Mutex`, atomiske operasjoner og en løsning der én goroutine eier tilstanden.

## Regel
Et program som "ser ut til å virke" er ikke bevis på at concurrent kode er korrekt.
