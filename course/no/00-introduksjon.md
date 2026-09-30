# 0. Introduksjon

Velkommen til EduGo.

Målet med kurset er at du skal gå fra å være ny i Go til å kunne lage nyttige kommandolinjeprogrammer, nettverksverktøy og tjenester på egen hånd.

## Hva du trenger

Du kan bruke enten:

- en lokal Go-installasjon, eller
- studentmiljøet i `student-oci/`.

Kurset er laget slik at du ikke trenger intern Ploos-infrastruktur.

## Arbeidsmåte

Hvert kapittel følger omtrent samme rytme:

1. Les konseptet.
2. Kjør eksemplene.
3. Endre dem.
4. Løs oppgavene.
5. Bygg en liten del av et større prosjekt.

## Første program

```go
package main

import "fmt"

func main() {
    fmt.Println("Hei, Go!")
}
```

Kjør programmet med:

```sh
go run .
```

I neste kapittel ser vi nærmere på Go-verktøyene, `go mod`, `go fmt`, `go test` og hvordan et Go-prosjekt er bygget opp.
