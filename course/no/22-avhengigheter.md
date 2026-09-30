# 22. Avhengigheter og testbar design

Go trenger ikke et stort dependency-injection-rammeverk for å være testbart.

## Læringsmål
Send eksplisitte avhengigheter via konstruktører og små interfaces. Skill program-wiring i `main` fra domenelogikk.

## Oppdrag
La miniservice få logger og en liten lagringsabstraksjon uten globale mutable variabler.

## Prinsipp
Lag interfaces der konsumenten trenger abstraksjonen, ikke automatisk rundt alle konkrete typer.
