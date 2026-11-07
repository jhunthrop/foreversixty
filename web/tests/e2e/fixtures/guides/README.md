# /guides fixtures

The signed-in accounts the guide mocks draw (`design/mocks/ws_guides.py`, through `design/mocks/guides_facts.mjs`)
and a guides parity spec serves (`../../support/guides-page.ts`; spec `design/specs/2026-11-08-guides.md`, gap G6).
No real fixture is a Survival hunter or a Fire mage, so these are DERIVED from the home fixtures, changing only the
fields listed (every other field, the name, gear and guild included, stays the source's):

- `me-survival-34.json`: `../home/me-34.json` (Felsnap) with the character's `spec` `Survival` and
  `build.talents.points` `[0, 0, 23]`.
- `me-survival-60.json`: `../home/me-60.json` (Thandrel) with `class` `Hunter`, `spec` `Survival`, `race`
  `Night Elf` and points `[5, 11, 35]`.
- `me-fire-60.json`: `../home/me-60.json` (Thandrel) with `class` `Mage`, `spec` `Fire`, `race` `Human` and
  points `[14, 34, 3]`.

After a change to a home fixture, apply the same edits here; `guides_facts.mjs` and the parity spec read these files.
