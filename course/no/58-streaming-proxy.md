# 58. Streaming og proxying

Streaming handler om å flytte data gradvis fremfor å buffre alt først. Proxyer må håndtere cancellation, backpressure, headers og feil på begge sider.

## Oppdrag
Lag en lokal streaming-handler med `io.Copy` og test at klient-cancellation avslutter arbeidet.
