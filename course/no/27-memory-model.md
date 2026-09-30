# 27. Go memory model

Concurrent kode må ha definerte synkroniseringspunkter.

## Læringsmål
Forstå happens-before på et praktisk nivå, hvorfor data races gjør resonnering ugyldig, og hvilke garantier channels, mutexer og atomics gir.

## Oppdrag
Ta en delt variabel uten synkronisering. Forklar hvorfor "den andre goroutinen ser sikkert verdien" ikke er en gyldig garanti. Rett designet på to forskjellige måter.

Les memory-model-dokumentasjonen når korrekthet avhenger av detaljene; ikke lær tilfeldige scheduler-observasjoner som språkregler.
