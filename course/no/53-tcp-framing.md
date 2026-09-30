# 53. TCP i dybden: framing

TCP leverer en ordnet byte-strøm, ikke applikasjonsmeldinger.

Studer delimiter-, length-prefix- og fixed-size-framing. Sett eksplisitte størrelsesgrenser.

## Oppdrag
Implementer en linjebasert PING/ECHO-protokoll over en `net.Conn` og test den med `net.Pipe`.
