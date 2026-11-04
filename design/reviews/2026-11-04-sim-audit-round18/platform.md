# Sim-accuracy round 18: platform lane

Site branch `wt/r18-platform` (from int/r18 4600b7d82). No engine changes and no fork branch. No sims were run. Nothing here changes DPS for any spec.

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P1 | Each guide's Leveling strip listed every range at its FIRST level. "Level 30 to 39" showed the level 30 build, so a level 36 reader never saw the points taken at 31 to 36. | `web/src/lib/guides/leveling-builds.ts` `guideLevelingBuilds` cut the ladder at `band`, and `LevelingBandStrip.astro` decoded that string. | CONFIRMED-BUG | web | S d36888614, reverted (superseded by land/leveling-tree) |
| P2 | The iOS and Android item tooltips dropped expertise rating. Round 17 added the line on the web only (f81ca7bdc). | Both `ItemVocabulary.statOrder` tables had no `expertise` key, so `statLines` never printed it. | CONFIRMED-BUG | apps | S 09ada75b0 |
| P3 | The Arcane opener card said "Frostbolt 3 s before the pull, without Heating Up". The prose and the APL both say 2.5 s. | See below. | CONFIRMED-BUG (card) | web | S 3df28dc3e |
| P4 | The Rebirth and Revive tooltips still show raw `$q` because spelltips were not regenerated after the `$q` fill. | See below. | OPEN (data job) | data | - |

### P1: per-band talent lists

**Reverted.** land/leveling-tree (47b55e8b6, owner request) replaces the chip strip with a talent tree and a 10-60 level slider, so this change was reverted on the branch. The notes below record what was done.

- **Change.** Each range now lists the build held at its LAST level: 29, 39, 49, 59 and 60. The build comes from the same monotone ladder the sim uses.
- **Level stamps.** Each talent chip names the levels its points come at inside the range, for example "Lv 32-36" or "Lv 23-24, 26". There is one point a level from 10, so the first range also stamps levels 10 to 19.
- **Order.** A row reads in taking order: the talents carried from earlier ranges first, then the ones it adds, sorted by the level they come at.
- **Planner link.** "Load in planner" loads the same end-of-range build.
- **Legend.** It now reads: "Each range lists what you hold by its last level. Lv: the character level each point comes at, one point a level. Outlined: points added in this range."
- **Tests** (written first and seen failing), all in `web/src/lib/guides/`:
  - `leveling-talents.test.ts`, for every guide:
    - every point from 10 to 60 is stamped exactly once;
    - each chip's rank equals the ladder's count at the range's last level;
    - rows are in taking order.
  - `leveling-builds.test.ts`: the last levels are 29/39/49/59/60, holding 20/30/40/50/51 points.
  - `pages/guides/_spec.test.ts`: frost's 30 to 39 row renders 30 points, stamped 30 to 39.
- **Not run.** `tests/e2e/guides-sections.spec.ts` was updated to the new legend and the stamp, but Playwright was not run (standing rule).
- **wow-player: SHIP on all ten captures.**
  - Pages: mage arcane at 1440 and 390, warrior arms, rogue combat, hunter marksmanship, warlock affliction, priest shadow, druid feral, paladin retribution, shaman enhancement.
  - First pass: the reviewer asked for taking order inside a row, a larger and brighter stamp, and a legend line saying what Lv means. All three were done.
  - The first pass also called the paladin page NO-SHIP for an unstamped chip. That was a misread, which the reviewer retracted.

### P2: app tooltips, expertise

- **Change.** `expertise` joins `statOrder` after hit in both apps. The label comes from the new copy key `items.stat.expertise` = "Expertise". The copy was regenerated with `web/scripts/gen-mobile-copy.mjs`; that diff is only the new key in `Copy.swift`, `Localizable.xcstrings` and `strings.xml`.
- **Tests:**
  - Android `ItemPageTest`: failed first, then passed. `HomePresenterTest` also passes.
  - iOS `ItemPageLogicTests` plus `ItemTooltipTests` (xcodebuild, iPhone 17 Pro simulator, 20 tests): passed.
  - `web/src/lib/mobile` codegen check: 34 passed.
  - The iOS test was not seen failing on its own, because the copy key does not exist before the change. The Android twin carries the same assertion and was seen red.
- **When it shows.** As on the web, the line appears only once the items JSON carries `expertise`. No committed items file has it yet (1.60.1.70707 and 70755: zero matches).

### P3: Arcane opener card

- **APL.** `data/curated/apl/mage-arcane.json` prepull has a Frostbolt (10181 rank 10) at -2.5 s, gated `not auraIsKnown(400625 Heating Up)`.
- **Who casts it.** The raid build takes Heating Up, so only the level 60 build (`build:` in `mage/arcane.md`) casts this Frostbolt. That build takes Improved Frostbolt 5/5 (frost digits `05350003`).
- **Cast time.** The client's Frostbolt 10181 has `cast_time_ms` 3000 (spellconst). Improved Frostbolt rank 5 reads "Reduces the casting time of your Frostbolt spell by 0.5 sec." That makes a 2.5 s cast, so the APL's -2.5 s lands on the pull.
- **Verdict.** The prose (line 72, "2.5 seconds out") and the APL are right. The card was wrong.
- **Cause.** `guideRotation` timed every pre-pull cast by the described (raid) build's cast-time cuts.
- **Fix.** A pre-pull cast whose talent gates rule out the described build and admit the other build now reads the other build's cuts. Both the web page and the app guide file pass the other build. Prose and APL are unchanged.
- **Tests.** `guide-rotation.test.ts` checks the Frostbolt reads 2.5 s (red then green), and that Pyroblast stays at 6 s.
- **Review.** wow-player SHIP on the rendered card.

### P4: Rebirth and Revive tooltips

- **The fill.** The `$q` fill landed in round 17 as 4f1fdf041 (2026-11-04 13:59 UTC). It touched `data/pipeline/spelltext.py` only.
- **No data run since.** The last `data.yml` run was at 06:34 UTC (run 41786387659, land/guides-r5; main's run 41786351663 was at 06:33). Both predate the fix. Committed spelltips were last written by f83b247a1 (2026-11-04 04:25 UTC).
- **Still raw.** The committed `data/builds/1.60.1.70755/spelltips/` files still print raw `$q`, and not only for Rebirth and Revive:
  - druid: Rebirth and Revive, 10 strings;
  - paladin: Redemption, 5;
  - priest: Resurrection, 5;
  - shaman: Ancestral Spirit, 5.
- **Not regenerated here.** My local `raw/` is the committed tree, so nothing was regenerated from stale raw.
- **Note for the coordinator.** The data job's spelltips step (`data.yml:309`, `pipeline spelltips`) needs one run on the re-layered 70755 raw. Then check that no `$q` is left in spelltips or talents.

## Before/after

No sim-facing change, so no DPS moved. Nothing outside the guide pages, the app tooltips and the app guide file changes.

## Files to regenerate

- `data/builds/1.60.1.70755/spelltips/*.json` and talents text: the data job (P4).
- `data/builds/1.60.1.70755/items/*.json` with the `expertise` key: the data job, carried over from round 17. The P2 line shows after that.

## Measure at refresh

None.

## For other lanes

None. `guideRotation` gained an optional `otherTrees` argument. Its callers are the guide page and `lib/app-data/guide-file.ts`, and both now pass it.

## For the guide lane

None. The Arcane prose is right as written.

## Still open

- P4 waits on the data job run.
- The e2e spec `tests/e2e/guides-sections.spec.ts` was updated but not run (integrator's gate).
