# 46. Retries og backoff

Retries kan gjøre et midlertidig problem mindre synlig, men kan også forsterke overlast.

Bruk begrenset antall forsøk, context/deadline og backoff. Retry bare operasjoner der semantikken tillater det.

## Oppdrag
Lag en testbar retry-funksjon der venting kan injiseres slik at testene ikke trenger reell sleep.
