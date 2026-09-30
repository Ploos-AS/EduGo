# 45. Protokolldesign

En protokoll trenger tydelige meldingsgrenser, versjonering, størrelsesgrenser og definerte feil.

## Oppdrag
Design en liten tekstprotokoll med kommandoene `PING` og `ECHO`. Bestem maksimal linjelengde og hvordan ugyldige meldinger avvises.

Husk at TCP er en byte-strøm: én `Read` tilsvarer ikke nødvendigvis én melding.
