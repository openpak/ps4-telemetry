# Protocol

PS4 system software sends usage logging to the `nsx` family
(`nsx.sec.np.dl.playstation.net` on PS4; the PS3 uses `nsx*.np.dl.playstation.net`).
The exact PS4 request shapes are unknown until a console capture exists.

This service is deliberately a sink: every request is answered `200` with an
empty body and only the method, host and path are logged. If a future capture
shows a host needs a shaped reply, that host gets its own route then.
