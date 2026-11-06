# Character page: experience spec

Status: built on `lane/character-build` (mocks, build and parity gate; section 5 records the lane's calls). Author: home lane, 2026-11-06. System: `2026-11-05-web-layout.md` (type §3.3, primitives §3A, the shared Paperdoll §4, BiS §5); pattern: `2026-11-06-home.md` (mock every state, build from `ui/` primitives, parity at 1440).

Route: `/character/<region>/<ruleset>/<slug>` (`pages/character/[...path].astro` and the `/character` shell, island `components/Character.svelte`). Out of scope: `/logs`, `/reports`, `/rankings`, `/ratings` (the redesign lane); this page only links to them.

## 0. What the page is today (survey 2026-11-06)
One scroll, no tabs (`#gear` and `#history` anchors only): identity header with a body render at lg, Forever Score card, rating card, owner Premium history charts, the gear panel (`CharacterGearPanel` → `GearDoll mode="worn"` on a Plate, no inspector), In game, Best per encounter, Every ranked fight, Builds seen, owner Premium coach history. States in `Character.svelte`: `missing`, `loading`, `failed`, `unlogged` (owner: `CharacterUnlogged`; visitor: `CharacterNotFound` with the other-realm buttons, sign in, lookup, addon line), known-but-empty, ready (visitor / owner; gear-first when no score; export not read). Data: `GET /v1/characters/{region}/{ruleset}/{slug}` (`CharacterPage`, `lib/rankings/api.ts`), `.../rating`, `.../history`, `.../coach/history`, `/v1/characters/search`, `/v1/me`, the build's item and BiS files. Fixture: `tests/e2e/support/character-fixture.ts` (`CHARACTER_PATH = /character/us/hardcore/elyra-duskvale`, Elyra Duskvale, level 60 Priest), `src/fixtures/character-gear.ts`, `character-history.ts`.

## 1. Persona rulings (wow-player, 2026-11-06) and the lane's calls
C1. **The Plate is the worn paperdoll** (the inspect frame a player expects). One shared `Paperdoll`, `worn` mode, the same component BiS uses.
C2. **Head answer: the upgrade DPS available**, `+312 DPS available`, in the sim's unit, with `Open in simulator ›` and `Talents in planner ›`. Forever Score is a small figure beside it; item level is not a headline. Lane call: the API's `gear_gap_total` is in score points, so the build converts per slot with the shared gain rule (`lib/home/upgrades.ts` `slotGainFor` against the band's BiS file), the same DPS the home page shows; a slot with no verified number adds nothing (never a fabricated figure).
C3. **Head facts:** spec, level, race, realm and ruleset, guild, and the talent split as a planner link (`Arms 31/20/0`, from the newest `builds_seen`). Owner only: the sync line. The body render image is cut (one paperdoll, no second figure of the character).
C4. **Paperdoll cells:** every slot an icon in its quality border, the item name, its item level; an enchant mark when the slot should carry one (missing in red); owner only: the `+DPS` gain to the band's pick. Selecting a slot opens the Inspector (side from 1760, in place below): the worn item, the band's pick with its source and drop chance, two runners-up, one `Open in best in slot ›`. Never the whole BiS list.
C5. **Score and rating merge into one open Section** (`Forever Score`): score, tier and rating as Figures, the one `Most to gain` line; parts and trend open in place. No second card.
C6. **Fights:** `Best per encounter`, five Rows (boss, parse percentile in its parse colour, DPS, date), then `All ranked fights (N) ›`. The `Every ranked fight` table is removed (logs own fights).
C7. **Builds seen:** three Rows (split, spec, first seen), each to the planner.
C8. **Owner extras:** Premium history (score over time, then parse per boss) collapsed below the fights, a one-line lock for free owners; coach history last, collapsed, Premium only.
C9. **In game:** cut to nothing; its score and planner link are already in the head and C5.
C10. **Not found:** one line, `No character named <Name> on <Region> <Ruleset>.`, then the lookup box filled with the name, the same name elsewhere as Rows (level, class, spec), `Sign in to see your characters`, and `Export with the addon and the companion to add this character.` When another realm matches, those Rows come first.
C11. **Owner with no logs (a leveler):** the rating collapses to one line, `No logs yet. Upload a log to get a rating.`; fights are hidden; a `Next` line (planner or sim) sits under the head, as on home.
C12. **Phone first screen:** name in class colour, spec, level and realm, the DPS figure, the paperdoll with tappable icons; the owner's Next line. Sync and the two links drop to a row under the doll.
C13. **Never:** a number without its unit, a slot name with no icon, a marketing line, item level or score as large as the DPS gap.

## 2. Layout (primitives)
`Page split`: PageHeader (crest, name in the class colour, the fact line, answer C2); summary (the owner's Next line, C11); main: Plate with the Paperdoll (C1, C4) and its Inspector, Section `Forever Score` (C5), Section `Best per encounter` (C6), Section `Builds seen` (C7), owner Sections (C8); rail from 1024: the Inspector at 1760 and up, else empty. Not found: `Page ledger`, PageHeader with the name, the C10 Sections.

## 3. States to mock (parity at 1440)
`visitor-60` (Elyra Duskvale, the fixture), `owner-34` (a leveler with no logs), `owner-60-premium` (history open), `no-score` (gear, no logs), `unread` (export not read), `not-found`, `elsewhere` (the same name on another realm), `loading`, `error`.

## 4. Data gaps
- DPS per slot when the worn item is not a listed runner-up of a weapon slot (C2: shown as no number).
- Drop chance for the inspector's pick comes from the BiS page's `loot.json` facts; the client does not load them today.

## 5. Mocks, data and the gate (character lane, 2026-11-06)

| What | Path |
|---|---|
| Generator | `design/mocks/ws_character.py` (only `ws_prims` primitives; `assert_no_surface_styles` holds; `ws_prims` gains `Disclosure`, the twin of the new `ui/Disclosure.svelte`) |
| Mocks, one per state | `design/mocks/web/web-character-<state>.html` for the nine states of section 3 |
| States the gate checks | `design/mocks/web/character-parity.json` (1440 x 1000) |
| Shared data | `web/tests/e2e/fixtures/character/*.json` (`python3 design/mocks/character_fixtures.py`) plus the home fixtures for the level 34 hunter. Every worn item is the band's pick or a listed runner-up, so every gain is the ranker's own number. The fixture build prerenders `/character/us/normal/felsnap` and `/character/us/hardcore/thalrin` beside Elyra (`lib/report/shell-paths.ts`) |
| Gate | `web/tests/e2e/parity-character.spec.ts` through `tests/e2e/support/character-page.ts`; side-by-sides in `web/test-results/parity/character-<state>.png`; the chrome captures in `design/mocks/data/chrome/web-character-*.json` |

Build: `components/Character.svelte` (states), `character/CharacterGear.svelte` (Section `Gear`, the Plate, the shared Paperdoll in `worn` mode and its Inspector), `CharacterScore.svelte` (C5), `CharacterOwnerSections.svelte` (C8), `CharacterMissing.svelte` (C10); the words in `lib/character/page-view.ts`, the cells in `lib/character/doll.ts`, the reads in `lib/character/gear-context.ts`. Cut: `CharacterGearPanel`, `ProfileDescriptor`, `CharacterInGame`, `CharacterNotFound`, `CharacterUnlogged`, `HistoryMount`, `HistoryReserve`, `UpgradeRow`, `lib/gear/panel.ts`, `lib/character-prompt.ts`, `lib/character-layout.ts`, `lib/character/history-mount.ts`, the body render and the `Every ranked fight` table.

Lane calls:
- **Layout `ledger`, not `split`.** Below 1760 the rail would be an empty 340 px column (the Inspector lives in the Plate, as on BiS); one column keeps the doll's rows wide. The Inspector sits at the side from 1760, in place below, as the shared Paperdoll does everywhere.
- **The Plate sits in a `Gear` Section** (meta `Level 60 raid` or `Levels 30 to 39`, action `Best in slot`), as home's upgrades do; the unread state titles it `Best in slot for level 30 to 39` and draws the band's picks with their sources.
- **Head answer:** `Upgrades available`, `+128.9 HPS` (the band's unit: HPS for a healer), with the Forever Score as a small Figure beside it and `Open in simulator ›`, `Talents in planner ›`. No answer when nothing is priced or the export was not read.
- **Spec fallback:** with no talents, builds or fights, the spec is the one the gear is priced for (`gear_spec`), so a leveler still reads `Beast Mastery Hunter` and gets a BiS link.
- **C5:** the parts open in place under `Score parts` (rows of points against each ceiling); the trend chart stays in the owner's History.
- **C8:** History and Coach history are Disclosures; their charts and list are read only when opened. A free owner gets one row to the plans.
- **C10:** the lookup is the name field and `Look up` (to `/characters?q=<name>`); region and ruleset are answered by the `Same name elsewhere` rows, which come first. `Sign in to see your characters` shows only to a signed-out viewer; the addon row is `ctaCopy.getAddonAndCompanion`.
- **Supporter (coordinator, 2026-11-06):** a Premium character's head meta carries the `Supporter` Tag (`premiumBlockCopy.supporterChip`); the level 60 fixture is a supporter, so visitor-60 and owner-60-premium check it.
- **Section action on touch:** `ui/Section`'s action link keeps its 14 px look and gets a 44 px hit area under `pointer: coarse` (12 px of padding, the same negative margin), so no head row grows.
- **`All ranked fights (N) ›`** goes to `/logs` until logs can filter by character (the redesign lane).

Data gaps (named, not built):
- **Enchants (C4):** the API's `CharacterGearSlot` carries no worn enchant, so the cell has no enchant mark and no red "missing" mark. It needs `worn.enchant_id` from the export (a migration and an API field).
- **Drop chance** comes from the API's `pick.chance`, so the inspector shows it without `loot.json`; the BiS runners-up carry no chance.
- **C12 phone order:** the sync line and the two head links stay in the head on a phone; moving them under the doll needs a PageHeader slot below the Plate.
