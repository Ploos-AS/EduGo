# 89. pprof in operations

`net/http/pprof` provides powerful diagnostics, but diagnostic endpoints are administrative surfaces.

## Mission
Enable pprof only on a separate local/admin listener in the lab and document why it should not be casually exposed on the public service listener.
