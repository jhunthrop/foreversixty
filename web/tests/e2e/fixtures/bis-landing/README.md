# /bis landing fixtures

The accounts the /bis landing parity spec serves from `/v1/me` (`tests/e2e/support/bis-landing-page.ts`) and
`design/mocks/bis_landing_facts.mjs` reads, so both sides draw one account. The home fixtures supply Felsnap at 34
(`home/me-34.json`), Thandrel at 60 (`home/me-60.json`) and the account with no character (`home/me-new.json`).

- `me-24.json`: Felsnap at level 24 (the owner's report, 2026-11-07), wearing the 20 to 29 head and chest picks: `2 of 14 slots`.
- `me-nogear.json`: Felsnap at 34 with no build on file: the no-gear sentence and the addon link.
- `me-sim.json`: Felsnap at 34 on Premium with one finished server sim: the finished-sim line under the head.

Regenerate them with the node one-liner in the commit that added them; each derives from `home/me-34.json`.
