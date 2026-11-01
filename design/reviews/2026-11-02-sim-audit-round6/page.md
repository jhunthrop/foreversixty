# Sim-accuracy round 6: page lane (2026-11-02)

The /sim landing page, `web/src/lib/sim/settings.ts`, /tiers/healer and the header e2e. Site branch `wt/r6-page`; no engine change. The per-spec audit was out of scope for this lane.

## Findings

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| P1 | **/sim below 60 defaulted to Raid-buffed; the published number is bare** (owner ruling 1). The ranker ranks every band bare and adds a raid entry only at 60 (`sim/cmd/leveling-bis/preset.go` `passesFor`), and the BiS page, tier list and home card headline the raid entry only where one exists (`web/src/lib/bis/presets.ts` `selectBand`). The page sent the raid buffs and that level's raid consumables at every level. At band 40 that read +71% (Fire) to +212% (Fury) over the published figure. Separately, the old "Solo" preset sent nothing at all, not even the class kit, so a Solo rogue had no poisons. **Fix:** `sim/request.PublishedPreset` / `RaidPresetBand` now state the rule once, and the ranker's `passesFor` reads it. `go run ./internal/genraidconsumes` also writes `sim/request/published-presets.json`: the rule as level rows, plus the bare set per spec and level (the class kit, `sim/leveling` `KitBuffs` and `KitConsumes`). `TestPublishedPresetsJSONIsCommitted` holds it current. `sync-raid-preset.mjs` carries the rule into the static `raid-preset.json` and the bare rows into a lazy `bare-preset.json` chunk (`preset-sets.ts`, renamed from `raid-consumes.ts`). Solo becomes **Bare**, the class kit, which is the published bare entry. `defaultSettings` picks the published preset for the character's level. A character swap re-defaults to the new band's preset until the player picks one; a Raid-buffed pick then holds across swaps. | `published-preset.test.ts` failed first (no module), then: the rule at every level against the generated rows; the rule against the headline entry of every published BiS file at bands 20 to 60; every spec's Bare defaults at bands 20 to 60 plus 33 and 59; a level-40 rogue's poisons and a Survival hunter's mark; swap and pick behaviour. Go `TestPublishedPresetFollowsTheRaidBand`, `TestBareSetsAreTheKitAtEveryLevel`. Agreement table below. | CONFIRMED-BUG (page) | S a1536dc3e |
| P2 | **Flaky header layout-shift e2e** (owner priority 5). `header-finished-sim` "no layout shift at 390px" failed about one run in eight on main. A layout-shift observer with `sources`, 40 runs at 390 px on both projects, found two causes and a baseline. (a) **0.0275:** the character header's answer cell. While the answer loads it holds the figures ghost (81 px). A sentence answer is 42 px: "no best in slot list", "all your gear matches" or "pick a spec". When the answer landed after the ready band's first paint, the button, the finished-sim line and the page below rose 39 px. Below 1024 px nothing held the band's height. The fixture's level-24 Holy priest has no band-20 list, so it always resolves to a sentence; in production any sentence answer does the same. (b) **0.20:** the home page's SkyBand ridgeline was an absolutely placed `<svg>` anchored to the band's bottom. The 320 KB home page can paint while the band is part-parsed, so the svg moved each time the band grew. The 0.0082 baseline in every run was the same svg following (a). **Fix:** a sentence or failed answer now shares its grid cell with the hidden figures it replaced (`figuresReserve` in `CharacterHeader.svelte`). The ridgeline is now the band's bottom background layer; a background is never a shift source, and the picture is unchanged (full width, 160 px, under the same gradient). The limit (0.02) and retries are unchanged. | Before: 5 of 40 runs over 0.02 (0.0275 twice, 0.2029 three times), every other run 0.0082. After: 60 of 60 runs at 0.0000, and the check itself passes 120 of 120 (`--repeat-each=30`, both viewports, desktop and mobile) | CONFIRMED-BUG (web) | S 414874898 |
| P3 | **Three real-engine e2e specs were stale** (wasm lane, via the integrator). `sim-talents-real-engine` and `sim-weights-real-engine` pasted talent strings written before 70707 ("not a legal build"). `sim-gear-real-engine` left `/v1/builds?mine=1` unstubbed. All four real-engine specs gated on `public/_sim/<engine sha>/`, while `make publish-wasm` writes `<sha>-<digest>/`, so they skipped unless the wasm was copied by hand. **Fix:** a support helper reads each spec's talent string from the active build's published BiS file (`publishedTalents`). The gate and asset URLs follow `PUBLIC_SIM_ARTIFACT`, the variable the page build reads (`E2E_ARTIFACT`). The gear spec drops its duplicated probe for the shared `hasWasmExports` and answers the builds call unreachable, as `sim-gear.spec.ts` does. | Built with `FOREVER_DATA=real PUBLIC_SIM_ENGINE=wasm PUBLIC_SIM_ARTIFACT=8baffe5fc-b3f1351a914c`: the four real-engine specs (smoke, Top Gear, weights, talents) pass on desktop. The first run, after the talent and gate fixes, failed only on the talents spec's unstubbed `/v1/builds?mine=1&page=1`. | CONFIRMED-BUG (test) | S aa280202e |
| P4 | **/tiers/healer after the nightly.** The nightly at 8baffe5fc (21d97955f) publishes `metrics.hps_error` and `metrics.hps_at_bounds` for every healer entry. At band 60 raid: Alliance Restoration druid 742.0, Restoration shaman 642.0, Holy priest 641.2, Holy paladin 553.1, Discipline 538.3 (errors 0.33 to 0.54). The shaman and priest swap at ×1.5 (618.5 against 658.8), so they tie. Holy paladin leads Discipline at both bounds by 3.1 and 14.8 against two standard errors of about 1.2. Horde: druid 739.8; shaman 639.8 and priest 623.4 swap at ×1.5, so they tie; Discipline 547.8 over Holy paladin 543.8 at ×1.65 and 552.0 over 532.9 at ×1.5. Expected 1, =2, =2, 4, 5 on both sides. | Real-data build at the pin, both factions at 1440 and 390 (`shots/tiers-healer-*`). The page reads Alliance 1 Restoration druid 742.0, =2 Restoration shaman 642.0, =2 Holy priest 641.2, 4 Holy paladin 553.1, 5 Discipline 538.3. Horde reads 1 Restoration druid 739.8, =2 Restoration shaman 639.8, =2 Holy priest 623.4, 4 Discipline 547.8, 5 Holy paladin 543.8. The tied rows carry "≈ tie". | Aligned | - |
| P6 | **The off hand under a two-hander looked like a broken slot** (wow-player, on the Fire band-40 capture at 390). The off-hand cell under Staff of Jordan drew the empty dashed socket of an unworn slot, next to "Two-hand in main hand". **Fix:** the paperdoll layout's off-hand cell now carries the covering two-hander (`coveredBy` in `lib/gear/paperdoll.ts`, was a boolean). The frame draws that item's icon greyed in a solid socket, as the game's character pane does. The styling keys off the cell's existing `data-state="two-hand"` in `paperdoll.css`, so the script only swaps the image. That keeps the character-page chunk at 18,424 of its 18,432-byte budget; a class swap in the template had put it 36 bytes over. | `paperdoll.test.ts` and `CharacterGearPanel.test.ts` failed first; re-shot `shots/sim-mage-fire-40-*` | CONFIRMED-BUG (page) | S 36ac04667 |
| P7 | **/sim's Damage tab showed one collapsed row named "Build"** (casters lane, same capture). /sim mounts the logs' per-actor table (`report/ActorTable.svelte`), which lists actors with each one's spells under a toggle, closed. A sim has one character, so the table read as a single 100% row. The row's name was the request's placeholder for a build with no character name (`NAMELESS_REQUEST_NAME`, a pasted export). The per-action data was intact: the fixture and the real wasm both carry every ability, and the headline sentence reads from them. **Fix:** /sim opens the first row on its spells (`openFirst`; the logs pages are unchanged), and the nameless player reads "You" (`namedSummary`). | `SimResults.test.ts` and `sentence.test.ts` failed first; re-shot every `shots/sim-*` | CONFIRMED-BUG (page) | S 36ac04667 |
| P5 | **Survival at band 60 against its published number.** The published horde raid set wears Teebu's Blazing Longsword and **Heartseeker** in the off hand, with Blackcrow (Survival 781.8 horde, 771.5 alliance). | A real wasm run on the published horde band-60 code at the page default (Raid-buffed, picked by the rule) reads 781 against the published 781.8 (`shots/sim-hunter-survival-60-*`). spec-breakdown on the same entry: 781.8 / 782.0 at seeds 1 / 50001. | Aligned | - |

## Page against published

"page" is spec-breakdown on the published horde entry under the preset the page now defaults to, which sends exactly the page's default ids. That the ids match is proven by `published-preset.test.ts` (bare) and round 5's tests (raid). "old" adds the raid buffs and that level's raid consumables to the bare entry: the page's default before P1. The published figure is the nightly's own run of the same entry. 10,000 iterations, seeds 1 / 50001.

| spec | band | preset | published | page | page vs published | old default vs page |
|---|---|---|---|---|---|---|
| Combat | 20 | bare | 40.5 | 40.7 / 40.7 | +0.53% / +0.48% | +199.0% |
| Combat | 40 | bare | 109.5 | 109.1 / 109.2 | -0.37% / -0.31% | +164.3% |
| Combat | 50 | bare | 158.7 | 158.2 / 158.3 | -0.33% / -0.26% | +167.9% |
| Combat | 60 | raid | 625.7 | 628.9 / 628.7 | +0.51% / +0.48% | unchanged |
| Fury | 20 | bare | 37.1 | 37.0 / 37.0 | -0.22% / -0.33% | +249.2% |
| Fury | 40 | bare | 125.3 | 126.0 / 126.0 | +0.58% / +0.55% | +212.1% |
| Fury | 50 | bare | 196.0 | 197.0 / 196.7 | +0.54% / +0.39% | +205.6% |
| Fury | 60 | raid | 781.7 | 782.2 / 782.9 | +0.07% / +0.16% | unchanged |
| Fire | 20 | bare | 34.5 | 34.6 / 34.7 | +0.29% / +0.62% | +90.3% |
| Fire | 40 | bare | 144.1 | 145.6 / 145.6 | +1.11% / +1.05% | +70.7% |
| Fire | 50 | bare | 184.5 | 186.2 / 185.9 | +0.93% / +0.76% | +152.6% |
| Fire | 60 | raid | 699.4 | 699.7 / 699.1 | +0.03% / -0.05% | unchanged |
| Survival | 20 | bare | 70.4 | 70.8 / 70.8 | +0.63% / +0.56% | +123.6% |
| Survival | 40 | bare | 157.2 | 157.4 / 157.2 | +0.12% / +0.03% | +131.1% |
| Survival | 50 | bare | 217.8 | 217.4 / 217.3 | -0.21% / -0.25% | +127.5% |
| Survival | 60 | raid | 781.8 | 781.8 / 782.0 | -0.01% / +0.01% | unchanged |
| Affliction | 20 | bare | 56.5 | 56.3 / 56.3 | -0.37% / -0.39% | +29.9% |
| Affliction | 40 | bare | 134.8 | 135.1 / 135.0 | +0.24% / +0.16% | +42.3% |
| Affliction | 50 | bare | 196.0 | 196.2 / 196.1 | +0.11% / +0.09% | +92.1% |
| Affliction | 60 | raid | 572.7 | 573.5 / 573.9 | +0.15% / +0.21% | unchanged |

The page and the published figure agree within about 1% at every band. The remaining gap is the published figure's own precision: `set_dps` is the ranker's verify run at 300 iterations (`sim/cmd/leveling-bis/verify.go:44`) on its fixed seed. Fire at 40 and 50 is the widest (+0.8% to +1.1%). Before P1 the page read +30% to +249% above the published figure below 60.

## Real wasm runs on the page

The local real-data build, with the wasm built at the pin (`make publish-wasm`, artifact `8baffe5fc-b3f1351a914c`). Each run pastes the published horde band entry's gear, race and talents with its level, as `FS1:...|level=N`, and runs at the page's default precision (3,000 iterations). Screenshots at 1440 and 390 are in `shots/`, for the wow-player review.

| spec | band | default preset | "What's in it" | page | published |
|---|---|---|---|---|---|
| Combat | 40 | Bare | Deadly Poison main hand, Instant Poison off hand | 109 ± 0.2 | 109.5 |
| Fury | 40 | Bare | nothing (warriors carry no kit) | 126 | 125.3 |
| Fire | 40 | Bare | Arcane Brilliance | 147 | 144.1 |
| Survival | 60 | Raid-buffed | the raid list, with Hunter's Mark and Heart of the Lion | 781 | 781.8 |

Fire's page run reads +2.0%. spec-breakdown on the same entry at 10,000 iterations reads 145.6 (+1.1%). The published figure is a 300-iteration verify run.

The captures were re-shot after P6 and P7, and the figures read the same.

## Gates

- Go: `sim/request` passes except `TestRotationLadder`, the integrator's goldens, which this lane does not regenerate and whose code this lane does not touch. With the owner's fork checkout instead of the pin, `TestLayerPutsThePresetStoneBesideARoguePoison` and `TestIDsMarkdownIsCommitted` also fail; both pass at the pin (go.work with an engine worktree at 8baffe5fc). `sim/cmd/leveling-bis` passes.
- Web: full vitest passes (709 files, 7,963 tests). `astro check` has 0 errors and `npm run lint` is clean.
- Full Playwright, fixture build, desktop and mobile, 3 workers: 3,373 passed, 0 failed, 415 skipped.
- Header check: `--repeat-each=30` on the layout-shift tests, 120 of 120 pass.
- lhci, all three push shards (7 URLs). Index (the page the header and sky fixes touch), guides, BiS and reports pass. Planner (0.85), /sim (0.93) and the saved sim (0.88) missed their performance minimums while the machine's load average was about 33 (other lanes and an Android emulator). An origin/main fixture build audited interleaved with the lane build scores the same, two rounds each: planner medians 0.875 / 0.875, /sim 0.94 / 0.94, saved sim 0.88 / 0.88. CI's quiet runners are the judge.
- Real-data wasm build: every island check passes. sim-island is 92,148 bytes against its 92,160 budget (round 5 left 285 bytes; P1 adds the published rule and the Bare half). `check-sim-engine-artifact` fails on the local 40 MB wasm, as in round 5.

## Still open

- Budgets: after P6 and P7 the real-data build's sim-island reads 87,366 bytes (budget 92,160), and the character-page chunk 18,424 (budget 18,432, 8 bytes to spare). The real-data build after P1 alone read sim-island at 92,148.
- **Unnamed spells in /sim's tables (found once P7 opened the rows; pipeline, not changed here).** Running every spec at bands 40 and 60 through spec-breakdown and resolving each action id against `simnames/<class>.json` and `_shared.json` leaves 36 ids unnamed. They read "An unnamed spell #id" in the Damage tab. Two causes:
  - The tables keep only a class's spellconst and talent ranks. That drops item procs and generic actions that `spells.json` does name: Thrash 21919 (25 specs), Chthonic Power 1249110 (14), Molten Fury 1249113 (6), Wound 18107 (5), the wand's Shoot 5019 (4), Defensive State 5302, Touch of the Grave 1260189/1260198/1260201, Haste 13494, Quick Strikes 469144, Strength of the Champion 16916, Lightning 18979/18980, Revitalize 23863, Searing Arrow 29638 and others.
  - Some engine ids have no row in the Forever client's SpellName table at all: Ignite 12654, 12577, 8514, 10607, 10611, 12867, 1311648, 20134, 20947, 20955, 20957, 15332, 15336, 16280 and 29178. They need a name source beside the client: the engine, or a curated list like `sync-data.mjs`'s RACIAL_SPELLS.
- The same Ignite and Shoot rows read "Physical". The engine's action metrics report no school for them (`sim/adapter/adapter.go` `combatLogSchool`: 0 maps to 0, and the page reads 0 as Physical). Ignite is Fire, and Jaina's Firestarter shoots Fire. This is an engine metrics question for the casters lane.
- Raid-buffed below 60 still sends the curated raid buffs without the class kit layered on (`RAID_BUFFS` for every spec). At 60 the only difference is Protection paladin's Devotion Aura (round 5). Below 60, Raid-buffed is a player's choice with no published number to match.
- The real-engine specs are not in any CI job. The `build-real` job already builds the real-data site, and with the wasm artifact the four take about 12 s on desktop. They belong there once the wasm can be served (it is over Cloudflare's 25 MiB static-asset limit locally; CI's own build is the judge).
