# Best in slot landing (/bis) in the flat system: experience spec

Author: LANDINGS lane, 2026-11-07, ship mode. The owner found /bis still in the old style after the rollout (an eyebrow over an all-caps title, the gold "Your character" card with caps labels, a caps tag and a caps button, a 60 selected for a 20 to 29 character, an "ANOTHER CLASS OR SPEC" caps heading, the sim banner inside the card). This replaces the layout of `2026-10-10-bis-landing.md`; its jobs stand (answer "what should I wear right now?", reach your own spec at your own range in one click, no figure per spec: the tier list owns rankings). Built only from `ui/` primitives with the parity gate's region names (`design/parity.md`), consistent with the BiS spec page (`2026-11-05-web-layout.md` section 5), the guides index and home B.

## 0. Mocks, data and the gate

| What | Path |
|---|---|
| Facts | `design/mocks/bis_landing_facts.mjs` runs `lib/bis/landing-*.ts` on the active build and the account fixtures |
| Generator | `design/mocks/ws_bis_landing.py` (only `ws_prims` primitives) |
| Fixtures | `web/tests/e2e/fixtures/bis-landing/` (README there) plus the home fixtures |
| Mocks | `design/mocks/web/web-bis-landing*.html`: `out`, `out-30` to `out-60`, `24`, `34`, `60`, `new`, `nogear`, `sim`, `loading`, `error` |
| Gate | `web/tests/e2e/parity-bis-landing.spec.ts` (1440 for every state, 390 for `out` and `24`) |

## 1. Layout (L1 Ledger)

1. `PageHeader`: the Title `Best in slot` in Marcellus, once; the meta line `Levels 20 to 60` and `Updated <d Mon>`; the Lead `The best gear you can wear at your level, for every spec. Open yours.`; signed in, the answer at the right (section 3).
2. The finished server sim line, when one waits (`FinishedSimLine`, the shipped component and its dismiss; held before paint by Base.astro's `data-finished` hint so nothing moves).
3. `ControlBar`: one segmented group, `Level range`: `20 to 29`, `30 to 39`, `40 to 49`, `50 to 59`, `60`; at its end `Which spec is strongest? Tier list ›`. Phone: the one line with the range and `Change`.
4. The summary line (region `summary`): the selected range's `Top sources: ...` caption (computed at build time from the BiS files), `Each list is the best you can wear inside one level range; gear changes about every ten levels.`, at 60 `60: full <raid preset> gear.`; while the character's own range is not the selected one, `Your range: 30 to 39 ›`.
5. The one `Plate`: every class in game order as a group (crest 40 and the name in the class colour), its specs as slim link Rows (spec icon 28, the name, a muted `Tank` or `Healer` Tag, the gold `Yours` Tag and the selected ground on the character's spec). Three columns from 1024, two from 768, one below. No heading over it.

## 2. The default range (owner 2026-11-07)

`?band=` wins (the bar writes it), then the character's own range, then the viewer's last pick in storage, then 20 to 29. A stored 60 never beats a level 24 character.

## 3. The answer (ui/CharacterAnswer with its figure)

`Your character`, then the crest at the outer edge, the name in the class colour, `Level 24 · Dwarf Beast Mastery Hunter`; then `2 of 14 slots` as a Figure with the caption `Best-in-slot pieces you wear` (`Pieces you wore at the last sync` when the addon sync is stale) over a Meter, and the note naming the slots the range has no pick for; then one sentence-case primary Button, `Beast Mastery gear for 20 to 29`. In place of the figure: `Best in slot starts at 20.` (under 20), the no-gear sentence and `Get the addon and companion ›`, or `Your gear count did not load.` with `Try again`. No character: `No character on this account yet.` and `Set up your character ›`. A failed account read: `We could not load your character.` and `Try again`. Signed out: no answer column. Reserved at 380 x 224 from 900 while the account loads.

## 4. Behaviour

The page is prerendered; `lib/bis/landing-client.ts` follows the bar (the URL is the truth, the pick also kept in storage) and the header selector (the answer, the default range, `Yours`, the faction in every spec link) in place. Only the newest account read paints.

## 5. Kept from the old landing, restyled

Band captions, the tier list pointer, role words, Yours, the character's own range, the count with its stale wording and excluded-slot note, the no-gear and under-20 sentences, the one action, and the finished server sim line with its dismiss. Dropped: the eyebrow, the job paragraph's caps title, the gold card, the Supporter tag in the card (the header chip carries the account), the caps grid heading.

## 6. Parity states

As in section 0; the account states come from `tests/e2e/support/bis-landing-page.ts`, which also serves every spec's list from `data/builds/<active build>/bis` so the count reads the same picks on both sides.

## 7. /app/auth

The page a phone-app sign-in lands on when the app did not catch the link takes the flat message page (as the build-unavailable page): PageHeader `Open Forever Sixty` with `Your Battle.net sign-in is done. Open the app to finish.` as its meta line, the form layout, the one primary `Open Forever Sixty` (with a code) or the start-again line, then the not-you line. Mocks `web-app-auth`, `web-app-auth-code` from `design/mocks/ws_app_auth.py`; gate `parity-app-auth.spec.ts`.
