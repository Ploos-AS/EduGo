# 8. Testing

Go leveres med testing i standardverktøyene.

## Læringsmål
Skriv `TestXxx`, bruk table-driven tests, subtests, benchmarks og test coverage.

Kjør:

```sh
go test ./...
go test -cover ./...
go test -bench=. ./...
```

## Oppdrag
Skriv tabelltester for `ParsePort`: minste og største gyldige port, null, 65536, negativ verdi, tekst og tom input.

## Prinsipp
Tester er ikke pynt til slutt. Fra nå av skal kursprosjekter ha tester for logikk som kan testes automatisk.
