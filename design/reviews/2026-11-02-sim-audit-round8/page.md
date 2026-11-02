# Sim-accuracy round 8: page lane (2026-11-02)

Owner rulings 2 (two copy notes) and 4 (character page chunk headroom), plus the round-7 check that /sim's Raid-buffed list matches the curated preset. Site branch `wt/r8-page`. No engine commit. No per-spec audit in this lane.

## Owner rulings

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| P1 | **/tiers/healer did not say the healer order has no level-60 log behind it** (ruling 2). The first note above the list now ends: "No level-60 log is behind this order yet: it rests on the sim's heal model until launch logs exist." It joins the existing note about the stand-in fight, in the same quiet note style; still three notes, no pill, no number. It is healer-only. | `web/src/lib/tiers/tier-copy.ts` `healerNotes`. `tier-copy.test.ts` "says the healer order has no level-60 log behind it yet" failed first, then passed; it also holds the DPS and tank notes free of the line. Shots `shots/tiers-healer-{1440,390,2000}-{before,after}.png` (real-data build, Horde). | Copy (owner ruling) | S 3ede54cfd |
| P2 | **Survival had no line saying why its number fell** (ruling 2). The Survival guide's "In a raid" card gains a third bullet: "**Strength of Earth Totem:** Forever's totem gives less Strength than Era's, and the sim uses Forever's value. That is why Survival's Horde raid number fell: the totem, not the build." No figure is quoted. **Where, and why there:** Survival's number appears on four data-rendered surfaces (/tiers, the BiS page's Set DPS, the class landing's spec card, /sim). None of them has a per-spec prose slot: the BiS page's preset notes come from the nightly BiS file, and a per-row note on /tiers would be a new element. The guide is the spec's own hand-written page, the class landing's Set DPS card links to it, and the "In a raid" card is about raid conditions, which is where the totem comes from. It is on the guide's first screen at 1440 and 2000. It says **Horde** because the raid preset applies Strength of Earth to Horde characters only (`data/curated/presets.json` raid notes), so the Alliance number did not move for this reason. | `web/src/content/guides/hunter/survival.md`. `_sections.test.ts` "tells a Survival raider its number fell with the Forever Strength of Earth value, not the build" holds the totem's name, "not the build" and no digit in the card (fails without the bullet). Shots `shots/guide-hunter-survival-{1440,390,2000}-{before,after}.png`. | Copy (owner ruling) | S 3ede54cfd |
| P3 | **Character page chunk headroom** (ruling 4). `_astro/Character.*.js` read 18,427 bytes gzipped against 18,432 at origin/main 170b7b84c. The owner's coach history list (`CoachHistoryList.svelte`, with its coach copy and history-list helpers) was in the boot chunk for every viewer, but it only mounts for the owner on Premium, last on the page, and draws nothing until its first page is read. It is now its own chunk, loaded once /v1/me says the viewer is that owner (`createLazyComponent`, the same helper `HistoryMount` uses for the History section). The budget is unchanged. The paperdoll is not touched. | `scripts/check-island-size.mjs` on real-data builds (`npm run build`): before `Character.MXbxPYvM.js` 55,077 raw / **18,427** gzipped (5 bytes left); after `Character.BRu6xbCm.js` 51,324 raw / **17,263** gzipped (**1,169 bytes left**, -1,164). Full web vitest passes (710 files, 8,003 tests). Playwright: see Gates. | Done | S 722b0bd1d |
| P5 | **An unranked PvP source read "PvP rank 0"** (integrator, from the ranker lane). Forever's PvP rewards need no rank (owner ruling 2026-11-01), and the pipeline may emit pvp sources with none. `pvpSourceLabel` printed "PvP rank 0 · Horde" for rank 0; the item tooltip (`item-sources.ts`) and the BiS source cell (`source-cell.ts`) fell back to the bare source name or "unknown" when the rank was missing. All three now read "PvP reward · Horde", the ranker's own label (S 9947f008d on wt/r8-ranker). A ranked source still reads "PvP rank 11 · Knight-Lieutenant · Alliance"; a pvp source with no faction still falls back. The "Where to get it" footer already groups pvp by faction alone. | Vitest: `source-labels.test.ts`, `item-sources.test.ts` and `source-cell.test.ts` each gained a case that failed first (3 failed), then passed (441 tests in `lib/items` and `lib/bis`). | CONFIRMED-BUG (copy) | S 840f51d97 |
| P4 | **/sim Raid-buffed list against the curated preset** (round 7 follow-up). `RAID_BUFFS` is the synced `raid-preset.json` (14 buffs, 8 debuffs), and the per-spec raid rows carry the layered class kit (`preset-sets.ts`, `settings.ts presetBuffs`). | Go `TestRaidSetsCarryThePublishedBuffsAtEveryLevel` and `TestRaidConsumesJSONIsCommitted` pass at the pin cd4dfdcc0; `sync-raid-preset` regenerates no diff; `settings.test.ts` and `published-preset.test.ts` pass. No mismatch arrived from the auras lane. | Aligned | - |

## Gates

- Full web vitest at 722b0bd1d: 710 files, 8,003 tests pass. Scoped vitest after 840f51d97: `lib/items`, `lib/bis`, `lib/tiers` and the guide content tests pass. eslint and prettier are clean on the touched files.
- Playwright on a fixture build of 722b0bd1d (`FOREVER_DATA=fixture`, dead API base, `E2E_SKIP_BUILD`), desktop and mobile: every `character*`, `coach-history`, `paperdoll`, `tiers` and `guides*` spec plus every spec that opens a `/character/` page: 559 passed, 66 skipped, 1 failed. The failure, `character-identity.spec.ts:655` (the sync age ticking under a fake clock), passed 6 of 6 on a rerun, so it reads as machine load. 840f51d97 (the PvP label) landed after this build; it is covered by its unit tests.
- Go `sim/request` `TestRaidSetsCarryThePublishedBuffsAtEveryLevel` and `TestRaidConsumesJSONIsCommitted` pass with the engine at cd4dfdcc0.
- No server left running.

## Before/after

No sim figure moves: nothing here changes what the engine sims.

| measure | before (170b7b84c) | after |
|---|---|---|
| `_astro/Character.*.js` gzipped | 18,427 (budget 18,432) | 17,263 |
| `dist/sim-island.js` gzipped | 74,858 | unchanged |

## For other lanes

- **Physical (Survival APL):** the guide's Rotation card shows the Raptor Strike line note "gated at 20% mana so it does not starve Mongoose Bite below" (`data/curated/apl/hunter-survival.json:53`), but the line's condition is `currentManaPercent >= 0%` and the file's own notes say the gate was dropped on 2026-10-07. The guide prose already says "no mana gate at all". The line note should say the same. Visible in `shots/guide-hunter-survival-*`.
- **Physical (ruling 1):** the Survival guide's raid-build prose and `raidBuild` frontmatter are yours; P2's bullet is in the Overview's "In a raid" card, a different block, so the two merge cleanly.

## Still open

- The Character chunk now has about 1.1 KB. The next candidate, if more is needed, is `CharacterNotFound` (the 404 page a search can land on); lazy-loading it would add a second paint on that path, so it was left alone.
