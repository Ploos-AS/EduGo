# 18. TCP og UDP

HTTP skjuler mye. Nå går vi et nivå ned.

## Læringsmål
Forstå forbindelsesorientert TCP mot datagram-basert UDP, adresser, `net.Conn`, deadlines og enkel framing.

## Oppdrag
Lag først en lokal TCP echo-server og klient. Lag deretter et lite UDP-eksempel og observer forskjellene.

## Tenk
TCP er en byte stream, ikke en meldingsprotokoll. Hvordan vet mottakeren hvor én melding slutter?
