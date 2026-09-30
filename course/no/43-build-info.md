# 43. Build metadata og reproducerbarhet

Programmer bør kunne fortelle hvilken versjon de er bygget fra.

Studer `runtime/debug.ReadBuildInfo`, linker-injiserte verdier og hvordan VCS-informasjon kan inngå i diagnostikk.

## Oppdrag
Lag en `version`-kommando som rapporterer versjon og build-informasjon uten å påvirke kjernelogikken.
