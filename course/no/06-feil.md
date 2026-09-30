# 6. Feilhåndtering

I Go er feil vanlige verdier. Funksjoner returnerer ofte både resultat og `error`.

## Læringsmål
Opprett og returner feil, sjekk `err != nil`, bruk `errors.Is` og pakk feil med `fmt.Errorf` og `%w`.

## Oppdrag
Lag `ParsePort(string) (int, error)`. Gyldige porter er 1–65535. Tom tekst, ikke-tall og verdier utenfor området skal gi meningsfulle feil.

## Tenk
Hvorfor bruker Go eksplisitte returverdier for feil i stedet for exceptions som hovedmodell?
