# /setup and /connect fixtures

The reads the setup mocks (`design/mocks/ws_setup.py`, through `design/mocks/setup_facts.mjs`) draw, and the
files a `/setup` or `/connect` parity spec serves (spec `design/specs/2026-11-07-setup.md`). The account itself is
the home fixtures' `../home/me-34.json` (Felsnap, level 34 hunter, an addon build) and `../home/me-new.json` (no
character). The clock is pinned to 2026-11-06T10:00:00Z (`setup_facts.mjs` `NOW`).

- `devices.json`: `GET /v1/devices`, one companion (the device name `connect.spec.ts` uses).
- `connect-device.json`: `GET /v1/devices/connect/lookup` (as `connect.spec.ts`'s `DEVICE`).
- `pairing-code.json`: `POST /v1/devices/pair` (as `logs-companion.spec.ts`).
- `companion-latest.json`: `GET /dl/companion/latest.json`, the feed's companion manifest (0.2.37, Windows and Mac),
  for the download's direct link, version and size.
