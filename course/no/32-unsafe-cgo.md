# 32. unsafe og cgo

Dette kapitlet handler om verktøy som bevisst går utenfor den vanlige Go-sikkerhetsmodellen.

## unsafe
Bruk `unsafe` bare når du kan forklare invariantene, dokumentere hvorfor vanlig Go ikke er tilstrekkelig og teste på relevante arkitekturer.

## cgo
cgo kobler Go til C og er viktig for enkelte systembiblioteker, men påvirker bygging, portabilitet og runtime-grenser.

## Lab
Lag en minimal, valgfri cgo-demo. Resten av kurset skal fortsatt kunne bygges med `CGO_ENABLED=0`.

Ikke bruk unsafe eller cgo som snarvei i vanlige øvelser.
