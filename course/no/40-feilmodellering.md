# 40. Feilmodellering

Feil er en del av API-et.

Bruk wrapping med `%w`, `errors.Is` og `errors.As` når kalleren trenger å forstå feilkategorien. Skill forventede domene-feil fra uventede interne feil.

## Oppdrag
Gi noteservice maskinlesbare domene-feil uten å lekke interne detaljer gjennom HTTP.
