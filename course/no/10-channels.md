# 10. Channels

Channels lar goroutines kommunisere ved å sende verdier.

## Læringsmål
Opprett channels, send og motta, forstå buffered/unbuffered channels, lukk en channel og bruk `range` på mottakersiden.

## Oppdrag
Lag en generator som sender kvadrattall på en channel. En annen goroutine skal lese og skrive dem ut.

## Prinsipp
Ikke bruk channels bare fordi de finnes. Velg dem når kommunikasjon mellom samtidige aktiviteter blir tydeligere.
