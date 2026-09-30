# 60. io.Reader og io.Writer i dybden

`io.Reader` og `io.Writer` er blant Gos viktigste abstraksjoner. De kobler filer, nettverk, komprimering, hashing og prosesser uten at komponentene trenger å kjenne hverandre.

## Oppdrag
Bygg en pipeline med `io.Copy`, `io.LimitReader` og `io.TeeReader`. Forklar hvor data kopieres og hvor størrelsesgrenser håndheves.
