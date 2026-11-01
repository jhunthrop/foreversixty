# Sim-accuracy round 6: integration (2026-11-02)

The owner's round-6 rulings and priorities, the owner's questions, the integration, the pin and the review. The lane reports beside this file have the per-spec rows: warlock.md, casters.md, physical.md, hybrids.md, healers.md, ranker.md, pipeline.md, wasm.md and page.md.

## Owner rulings, and where each landed

| # | ruling | outcome | where |
|---|---|---|---|
| R1 | /sim below 60: default buffs match the published number at the selected band | The ranker's rule (bare below 60, raid at 60) now lives once, in `sim/request.PublishedPreset`; the ranker reads it, and `sim/request/published-presets.json` carries it to the page with a staleness test. A new character starts on its band's preset. Raid-buffed stays a choice and holds across character swaps. "Solo" became "Bare", which sends the class kit (the old Solo sent no poisons). Page against published, five specs at bands 20, 40, 50 and 60: within about 1%. Before, the page read +30% to +249% above published below 60. | page.md, S a1536dc3e |
| R2 | Horde race picks within 1%: guides unchanged; the adoption rule's bar applies to race | Recorded, not reopened. | warlock.md R1; casters.md |
| R3 | Restoration druid's lead: left, blocked on launch logs | Recorded, not reopened. | hybrids.md; healers.md |

## Owner priorities

| # | priority | outcome | where |
|---|---|---|---|
| P1 | Protection paladin threat +50% | **It stands.** Righteous Fury alone is +50.7%: +33% on all threat and +90% on Holy threat, multiplied. Each input has two sources: the 70707 client and the hotfix push, plus wowforevertalents' independent hotfix-cache capture, and TrinityCore and cmangos for the multiplication. Holy Shield mana, the healing third and Champion of the Light move under 0.5% each, and block value moves nothing, because no paladin ability reads it. This round found two threat bugs that take 3.4% back off. Shield Specialization's mana made threat although its row is flagged no-threat. Mana-gain threat took Righteous Fury's multiplier. The net is +44.7 / +45.6% over round 5's before pin, 96 to 98% of a Protection warrior's threat. The all-threat amount has only the server's data behind it; Blizzard published no note. A launch threat log stays the settling source. | hybrids.md Item 1 |
| P2 | Torch of Light in every warlock set | **Dropped from the raid sets.** Its client rows are right and it is an Enchanting craft open at launch, but it sims last of the rival wands in every warlock spec. The cause was the ranker, which credited a wand's whole tooltip DPS whenever the rotation had a Shoot line. Raid warlocks Shoot 0% of the fight. The ranker now credits the wand by the band's own Shoot time (S 413808063). Its verify step also needed a flat 1% gain and published a +0.33% runner-up as "0, verified" (S 2c379af92). Separately, wand shots were Physical, so armor took about 40% of each hit: items now carry their wand's school (E ee0393589, S e395edd50). Torch stays where mages Shoot 22 to 46% of a bare fight, and the sim agrees there. | warlock.md T3, T4; ranker.md W1, W2 |
| P3 | Spell generators read the hotfix push automatically, with a test | `pipeline fetch` layers wago.tools' push into every fetched raw table except ItemSparse/Item (S f9155ec20, `data/pipeline/hotfix_push.py`). The newest push wins, removals are honoured, and a failed fetch fails the run. Every generator reads the layered `raw/`, so none needed its own change: simconst, trainables, spellranks, normalize, simdb, the set rows, and the fork's relicmods, periodiccrit and eureka. There are 32 tests through the real generators: Ice Barrier 5, Frigid Power 15, Touch of the Grave 15 s. Round 5's hand dump had missed 78 SpellEffect, 312 SpellMisc and 148 SpellName rows and every Trait and Curve row, because wago's pager drops rows. client.yml now re-fetches the active build when the push moves (S c251c7ef0), and the engine follow hashes the push record. | pipeline.md |
| P4 | sim.wasm size | **At risk, fixed.** The deployed raw module was 24.7 MB, 94% of Cloudflare's 25 MiB per-file limit, and the pin grows it about 100 KB a round. The 40.4 MB local build came from the laptop's go1.27; CI's go1.25.11 builds 24.8 MB from the same tree. publish-wasm now writes `sim.wasm.gz` (4.46 MB), and the page inflates it with DecompressionStream. A 20 MiB per-file budget (configuration, with tests) fails the postbuild over it. The Makefile pins the toolchain to sim/go.mod's go line, and `make artifacts` warns when the engine isn't the pin. | wasm.md, S 16af343cd |
| P5 | header-finished-sim "no layout shift at 390px" | **Fixed at the cause.** There were two causes: a sentence answer was 39 px shorter than its reserved ghost, and the SkyBand ridgeline moved while the large page parsed. Answers now keep the ghost's height, and the ridgeline is a background layer. 60 of 60 diagnostic runs record zero shift, and `--repeat-each=30` passes 120 of 120. The limit and retries are unchanged. | page.md, S 414874898 |
| P6 | Fresh re-audit of every spec | Done by every audit lane; rows in each report. | all |

## The owner's questions

- **Survival's Improved Concussive Shot.** The two points in Improved Concussive Shot did nothing. They were filler in Marksmanship's first row, and our build search wrongly protected them as an untested talent. They now go to Deadly Aspects: about +4% on the leveling build and about +8% on a new raid build. Efficiency is modelled correctly but worth nothing for Survival, because Resourcefulness already cuts melee mana costs by 60% and the hunter never runs out of mana in a 180 s fight. Beast Mastery's raid build also improves (+1.3 / +1.5%); Marksmanship's winner (+0.9%) is held under the adoption rule. (physical.md; S 34afb4634)
- **Restoration druids and Wild Growth.** Wild Growth is in Forever: a 1-point Restoration capstone with trainer ranks at 50 and 60, and the engine matches the client. The build skips it because our heal fight runs a druid out of mana before it runs out of time, and in that fight Rejuvenation with Bloom heals more per point of mana. We found one real flaw on the way: raid-wide damage never hit the healer, so a fifth of every group heal landed on a full-health druid. With that fixed, Wild Growth comes out level with the guide (-0.3%, inside error), so the build stands, and the guide now says to take it for heavy bursts of group damage. The same audit of every healer fixed Discipline (it never cast Greater Heal, and Binding Heal was missing), Holy paladin (it precast Blessing of Light, which Forever removed) and Restoration shaman (Chain Heal pacing, and Mana Tide rank 4). (healers.md)
- **Bear tanks.** The bear was in Dire Bear Form the whole fight, with the right form stats, armor, 2.5 s swings and +50% form threat. It looked undertuned because Maul made 1x its damage in threat where Classic gives 1.75x, and Swipe used a Season of Discovery 2x. Both are 1.75x now: bear threat +29.5%, about 85% of a Protection warrior's. Lacerate's "high threat" factor is unsourced and stays open until a launch log. (hybrids.md)

## Integrator findings

| # | finding | evidence | confirmed? | fix sha |
|---|---|---|---|---|
| I1 | **web.yml red on main since f95ff8e77, so nothing deployed after bc1f08b6f.** That included round 5's pin. The cause was prettier on two mobile-codegen files (a keyword array reflow). | the wasm lane's W6; run on 21d97955f | CONFIRMED (CI) | S bef2cb599, pushed alone; web deployed |
| I2 | **Two lanes wrote the same ranged attack power dedup.** Merged, it ran twice and cut a ranged-only excess (10/25 gave 5, not 15). Only the merge produced it; one implementation stays. | test_simdb_equip | CONFIRMED (merge) | S 0ee44a9fd |
| I3 | **Hybrids' Illumination called UpdateManaRegenRates, which physical's per-tick regen removed.** The calls are unneeded now. | int/r6 build | CONFIRMED (merge) | E 661293644 |
| I4 | Stale test expectations: The Immovable Object's block value is Wowhead's shield total (41 + 25 = 66) since round 5, SimItem has damage_school, and the smoke test's Mana Tide id moved to rank 4. | sim and data tests | test-only | S 0ee44a9fd |
| I5 | `make engine-pin` refused an engine worktree (`.git` is a file there). | Makefile | tooling | S 0ee44a9fd |

## Integration

- **Fork `int/r6`** merges lane/r6-casters, -pipeline, -warlock, -dots, -hybrids, -physical and -healers. Conflicts, all resolved:
  - `sim/core/dot.go`: Dynamic (DoTs) and DynamicSnapshot (HoTs), both kept.
  - Regenerated goldens: TestP1Mage, TestEnhancement, and the paladin and priest conformance goldens with SUMMARY (E 1ba9f4c2d).
  - The fork suite passes with `-tags=with_db`, except `sim/web`, which also fails setup at 8baffe5fc.
- **The pin is b57765fb9.** data/proto is re-vendored at it, for damage_school and pulses_reach_healer. The rotation ladder goldens are regenerated at it with `make simdb` and a scratch go.work.
- **The site** merges wt/r6-wasm, -pipeline, -warlock, -casters, -physical, -hybrids, -healers, -ranker and -page. The only conflicts were page and wasm in the real-engine e2e support, resolved to both: page's artifact gate plus wasm's `sim.wasm.gz`.
- **Gates, all exit 0:**
  - `go vet` and `go test ./...` in sim
  - the coach and sims packages in api
  - `make apl-check` against the pin
  - `make styles-check`
  - the full data pytest
  - the full web vitest
  - astro check
  - `npm run lint`
  - the sim, tiers, planner-share-sim and header-finished-sim e2e specs, 78 of 78
  - The page lane also ran the full Playwright fixture suite (3,373 passed) for its shared-header change.

## Before/after

Pin 8baffe5fc with origin/main, against pin b57765fb9 with this branch. Both run on today's published data: the BiS sets and builds, and the simdb built from the committed inputs. Raid preset, band 60, 10,000 iterations. Healers are measured in effective HPS on the heal profile, tanks in threat per second.

| Spec | Alliance, seed 1 | Alliance % (1 / 50001) | Horde, seed 1 | Horde % (1 / 50001) |
|---|---|---|---|---|
| druid-balance | 537.2 to 565.7 | +5.3 / +5.2 | 540.3 to 567.2 | +5.0 / +5.1 |
| druid-feral | 615.6 to 624.1 | +1.4 / +1.4 | 616.5 to 624.7 | +1.3 / +1.3 |
| druid-feral-bear | 677.2 to 877.0 | +29.5 / +29.5 | 682.8 to 885.3 | +29.7 / +29.6 |
| druid-restoration | 741.5 to 743.3 | +0.2 / +0.2 | 739.2 to 741.8 | +0.4 / +0.3 |
| hunter-beast-mastery | 659.0 to 669.3 | +1.6 / +1.7 | 674.2 to 685.6 | +1.7 / +1.6 |
| hunter-marksmanship | 661.6 to 661.6 | +0.0 / +0.0 | 679.7 to 679.7 | +0.0 / +0.0 |
| hunter-survival | 771.5 to 775.0 | +0.5 / +0.5 | 781.8 to 785.4 | +0.5 / +0.4 |
| mage-arcane | 585.3 to 586.8 | +0.3 / +0.2 | 588.0 to 589.5 | +0.3 / +0.2 |
| mage-fire | 687.8 to 648.3 | -5.7 / -5.8 | 699.7 to 656.7 | -6.1 / -5.9 |
| mage-frost | 542.0 to 584.3 | +7.8 / +7.7 | 546.9 to 590.1 | +7.9 / +7.8 |
| paladin-holy | 553.8 to 578.5 | +4.5 / +4.5 | 544.2 to 564.6 | +3.7 / +3.9 |
| paladin-protection | 1027.9 to 992.7 | -3.4 / -3.4 | 1043.2 to 1008.2 | -3.4 / -3.4 |
| paladin-retribution | 652.6 to 666.2 | +2.1 / +2.1 | 658.8 to 669.3 | +1.6 / +1.7 |
| priest-discipline | 538.4 to 597.9 | +11.1 / +11.0 | 548.0 to 608.1 | +11.0 / +10.9 |
| priest-holy | 641.4 to 680.7 | +6.1 / +6.0 | 623.2 to 662.6 | +6.3 / +6.3 |
| priest-shadow | 707.5 to 707.8 | +0.0 / +0.0 | 713.0 to 713.9 | +0.1 / +0.2 |
| rogue-assassination | 621.9 to 621.9 | +0.0 / +0.0 | 625.0 to 625.0 | +0.0 / +0.0 |
| rogue-combat | 622.7 to 622.7 | +0.0 / +0.0 | 628.9 to 628.9 | +0.0 / +0.0 |
| rogue-subtlety | 558.4 to 558.4 | +0.0 / +0.0 | 564.7 to 564.7 | +0.0 / +0.0 |
| shaman-elemental | 485.9 to 459.4 | -5.5 / -5.5 | 494.6 to 467.6 | -5.5 / -5.5 |
| shaman-enhancement | 630.0 to 608.0 | -3.5 / -3.5 | 634.3 to 612.5 | -3.4 / -3.5 |
| shaman-restoration | 641.9 to 690.9 | +7.6 / +7.6 | 640.2 to 687.7 | +7.4 / +7.3 |
| warlock-affliction | 572.4 to 563.0 | -1.6 / -1.6 | 573.5 to 564.3 | -1.6 / -1.6 |
| warlock-demonology | 644.3 to 625.8 | -2.9 / -2.9 | 646.8 to 628.5 | -2.8 / -2.8 |
| warlock-destruction | 670.2 to 617.8 | -7.8 / -7.8 | 673.4 to 620.6 | -7.8 / -7.8 |
| warrior-arms | 728.3 to 728.3 | +0.0 / +0.0 | 746.8 to 746.8 | +0.0 / +0.0 |
| warrior-fury | 772.2 to 771.3 | -0.1 / -0.1 | 782.2 to 782.2 | +0.0 / +0.0 |
| warrior-protection | 1036.0 to 1067.7 | +3.1 / +3.1 | 1024.6 to 1055.9 | +3.1 / +3.0 |

Where the big moves come from:
- **Improved Scorch is out of the raid preset.** Blizzard's mage deep dive says it now benefits only the casting mage, and the client text agrees. This drives most of Destruction, Elemental and Demonology, and part of Fire and Enhancement.
- **Fire:** Ignite now rolls crits into one burn.
- **Frost:** Ice Lance's x4 now covers spell power.
- **Balance:** Moonfire and Wrath were registered twice, and the rotation cast the copies without talents.
- **Bear:** Maul and Swipe threat.
- **Healers:** the healer now takes raid damage. Discipline casts Greater Heal and Binding Heal. Restoration shaman gets Chain Heal pacing and Mana Tide rank 4. Holy paladin gains Illumination's casting regen (+13%) and loses Blessing of Light (-10%).
- **Enhancement:** Windfury no longer doubles Elemental Weapons, Stormstrike is one charge, and Flurry lasts 15 s.
- **Protection warrior:** Defensive Stance and Defiance multiply.

What this table does not hold, because each needs data.yml's regeneration and the bis.yml re-rank:
- **Item data:**
  - wand damage school, which lifts bare casters (Frost about +3%)
  - classic-db single-school spell damage (180 rows)
  - doubled ranged attack power on three items (horde hunters about -1.2%)
  - Rune of the Guard Captain at 14 attack power (horde melee about -0.8%)
  - classic-db shield block
  - half-up stat rounding
- **Re-ranks:**
  - Torch of Light leaves the raid sets.
  - Empty trinket slots fill.
  - The new Survival and Beast Mastery builds are published.

## wow-player review

Verdicts on the lane reports, the before/after table and the page shots:
- **SHIP:** physical, pipeline, wasm and page.
- **SHIP WITH FIXES:** warlock and ranker, casters, hybrids and healers. Each was settled before the push:
  - **Holy paladin's split.** Blessing of Light removed is -10.1%, Benediction on Holy Shock +0.3%, and Illumination's 15 s casting regen +13%. Illumination rests on the 70707 client rows alone: SpellEffect 1237075, aura 134 at 50, for 15 s. cmangos sums aura 134, caps it and applies it to spirit only, as the engine does. The mana math is consistent: mana lasts 205 s without it and 259 s with it, on a mana-bound profile. It stands, single-source, until a launch log (hybrids.md).
  - **Discipline: +11% integrated against the healers lane's +13.7%.** The casters lane's Mental Agility and Spiritual Healing mask fixes take about 2% back off.
  - **Enhancement: -3.5% against hybrids' -3.0%.** Improved Scorch out of the preset is the other -0.5%.
  - **Fire band 40's "Scorch and Fire Blast are 62% of your damage" is right** (casters.md F6). At 40, Scorch hits for about half a Fireball, in less than half the cast time and for half the mana. Only Scorch and Fire Blast get Incineration's crit, which feeds Heating Up. With Fireball as the filler the mage loses 1.7%.
  - **Fire band 40's off-hand slot** read "Two-hand in main hand" beside an empty icon, and the Damage tab showed a single "Build" row. Both are fixed: a two-hander now covers the off hand greyed, as the character pane does, and /sim opens on the player's spells (page.md P6, P7; S 36ac04667). Opening the rows showed 36 unnamed action ids, and no school on Ignite and Shoot. Both are display issues: the engine deals Ignite as Fire. They are listed in the ledger.
  - **/tiers/healer:** the order flips at the nightly, because Discipline passes Holy paladin. The page renders from published metrics, so it follows the nightly. Re-check it after the run.
- **Copy suggestions, not sim bugs, left for the page owners:**
  - a line on BiS saying why Torch loses in raid
  - checking that Enhancement's page shows its shield
- **The owner's sanity questions:**
  - Protection paladin at 96 to 98% of a warrior and bear at about 85% are plausible. The bear figure is provisional on Lacerate's factor.
  - The healer moves are plausible in direction.
  - The caster moves are what a player expects when Improved Scorch stops covering the raid.
