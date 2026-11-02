# Sim-accuracy round 8: integration (2026-11-02)

The owner's round-8 rulings, the integration, the pin and the review. The lane reports beside this file have the per-spec rows: warlock.md, casters.md, physical.md, hybrids.md, healers.md, auras.md, items.md, ranker.md, pipeline.md, page.md and wasm.md. Auras and items are new cross-cutting lanes this round, for the areas recent rounds showed to run deep.

## Owner rulings, and where each landed

| # | ruling | outcome | where |
|---|---|---|---|
| R1 | Survival's raid build is +0.7% over its leveling build, under the 1% bar: one build | **Done.** The guide's `raidBuild:` line is gone. The ranker's raid pass, talent-search, dps-per-point and addon-builds fall back to the leveling build, and a vitest holds that Survival carries no raid build. The ruling alone is -0.6% alliance and -0.75% horde at band 60 raid. | physical.md; S 768c32817 |
| R2a | /tiers/healer says the order has no level-60 log behind it | **Done.** The first healer note ends "No level-60 log is behind this order yet: it rests on the sim's heal model until launch logs exist." No pill, no number. | page.md; S 3ede54cfd |
| R2b | Survival says why its number fell, without numbers | **Done, corrected at integration.** The page lane wrote "Horde", because the preset note said the totems were Horde-only. The physical lane measured that both factions get Strength of Earth, and that the totem explains all of round 7's drop on each. This round's Curse of Recklessness fix lowers the number again. The guide's raid card now reads: "Forever's Strength of Earth Totem gives less Strength than Era's, and Forever's Curse of Recklessness strips less armor. The sim uses Forever's values for both, on either faction. The raid buffs moved the number, not the build." A test holds both names, no digits and no "Horde". The published number is still the pre-round-7 one until the refresh, so both causes apply to it. | page.md; S 3ede54cfd, d4c2619b5; shots `guide-hunter-survival-*-after-int.png` |
| R3 | One loot generator, stable output, with a test | **Done.** `loot` writes `raw/client/loot-classification.json` (Map.csv instance types and ItemSparse PvP ranks), and the full run and `loot-merge` build through one function from it. The merge fails fast without the file and never reads loot.json back. On 70707 the two write byte-identical loot.json. The fixture test runs full, merge, full, merge for both switch values. | pipeline.md; S 52b3f1eb3 |
| R3+ | Mid-round: battlegrounds and PvP ranks exist at launch. Keep the AV/AB sources every run; keep rank 5/10 rewards labelled the way Forever sells them; add Darkspear Islands if the client has it; PvP gear is eligible only where the guide rules allow it | **Done.** `BATTLEGROUND_SPAWNS_ARE_SOURCES = True` is held by a test. Every battleground source matches the committed loot.json. The rank 5/10 tabards keep "rank N", because Forever's server text names Rank 1 and Rank 6. **467 Forever PvP listings whose rank came only from their Era title** are now unranked "PvP reward (Alliance/Horde)" sources: the client's ItemExtendedCost has honor costs but no item link, so how they are sold is not in the client. Each carries `title_rank`, which the ranker reads for its rank-10 cap, so default BiS eligibility is unchanged (97 listings at or under the cap before and after, the same set). The 20 Horde "Stone Guard's" pieces lost a spurious rank 10 and source from their vendor, like their Alliance twins; a local re-rank keeps every pick to the decimal. Darkspear Islands (map 2997) has no spawns or loot in the client; its vendor is already a source, gated "later". | pipeline.md; S e1b402d61, 970ec9e90; ranker.md; S 19f27af30, 9947f008d; page.md; S 840f51d97 |
| R4a | Character page chunk: trim code, don't raise the budget | **Done.** 18,427 to 17,260 bytes gzipped against the unchanged 18,432 budget (1,172 bytes left, measured again on the integration build). The owner-only coach history list loads lazily. The paperdoll is untouched. | page.md; S 722b0bd1d |
| R4b | sim.wasm.gz: cheap wins, record what's left | **Done.** 4,481,935 to 4,276,658 bytes, 95.0% to 90.6% of 4.5 MiB, with the module byte-identical (56 node runs and 4 real-engine specs agree). pigz's zopfli mode writes a standard gzip stream. The budget now gates the file that ships, not a second gzip of the module. What is left is listed in wasm.md: brotli via a Worker route (about 1.2 MB) is the one big lever; no-inlining is a speed trade. | wasm.md; S 2c0943d1c, a96af40c1 |

How the ranker treats PvP and battleground items today (recorded, no policy change):
- PvP rank rewards are open at launch (content-phases "honor-system"). Rank 11 and up is never a default pick, only a labelled alternative (`pvpRankCap = 10`). A vendor row inherits its matching PvP source's rank.
- 69 published slot picks are PvP-sourced, all rank 7 to 10, mostly tank and melee bands 30 to 60.
- **No battleground item is eligible at any band.** Battleground reputation (WSG, AB, AV, the AV quartermasters and quests) sits behind content-phases' "battleground-reputation" gate, and Darkspear Islands behind "darkspear-islands". Both gates are "later".
- For the owner: should both gates open at launch? And should the `title_rank` cap stay now that Forever's PvP gear is not rank-locked?

## Integrator findings

| # | finding | evidence | confirmed? | fix sha |
|---|---|---|---|---|
| I1 | **`tools/wasm-gzip.sh` was committed without its exec bit.** `make artifacts` and deploy call `./tools/wasm-gzip.sh`, so they would stop on "Permission denied". | The gates' `test_wasm_gzip.sh`: 12 of 19 checks failed | CONFIRMED | S a96af40c1 |
| I2 | The Survival copy said "Horde" (R2b). | physical lane's toggle of Strength of Earth back to 61 on cd4dfdcc0: both factions return to round 7's before | CONFIRMED (copy) | S d4c2619b5 |
| I3 | Three bugs were fixed by two lanes each: Innervate at 40 (casters, healers), the shaman's 17 + Spirit/5 regen (casters, healers), Curse of Recklessness (physical, auras). The healers' Innervate and regen stay. The auras lane's Curse of Recklessness stays, because it also ranks the debuff by level, and the physical lane's single-rank constant and its test go. | int/r8 merges | merge | E 0c43b4e2f, 590900fed |
| I4 | The items lane cherry-picked the physical and warlock proc commits so its procs share one helper. The merge takes the items side of `weapon_proc_rolls.go` and `item_effects.go`. | int/r8 merge | merge | E 7fdbdced2 |
| I5 | Six fork goldens and the druid conformance golden move under several lanes together. | fork suite | test-only | E 40021374a |
| I6 | settings.ts's comment still described faction-gated blessings and totems; the fork removed the gates in 86d33a100. | auras lane | copy | S d4c2619b5 |

Follow-ups the integrator sent back to lanes, all settled this round:
- **Kindling Stave's Ignition was left open at first.** A client proc modelled at zero is wrong, and Necropile Drain, with the same row shape, fired on every cast. Both now use the engine's one-a-minute rule for an unstated "chance on" proc. Two Forever beta-log fits support it: Dream Venom fired 10 times in 358 swings, and Rotmending 10 times in about 258 heals. The items lane put every other unmodelled item proc on the same helper.
- **The Survival copy's cause.** The physical lane bisected it (I2).
- **The stale Raptor Strike note.** The curated APL is fixed (S 3f8d91d77). The Rotation card reads it from the nightly-owned addon-data.json, so the page shows the new note only after the refresh.
- **Innervate and the other learn-level gaps in healer APLs.** Innervate was the only one.
- **Mutilate rank 1 at 30**, not 40 (S d4f52df2e). No published number moves.
- **PvP labels and `title_rank`**, across the pipeline, ranker and page.
- **The ranker's K2 path losses.** Geared weights left some band rows below what origin/main's ranker picks: Shadow 60, Arms raid 60, Balance 50 Horde, Elemental 60, and the tanks. A fix that picks worse than the code it replaces isn't done. Each band whose geared sweep replaced the ladder sweep is now ranked start to finish under both weight sets, and each faction publishes the run whose final set measures higher. The incumbent guard still runs after that.
  - **Re-rank:** 336 rows against origin/main; 38 of 300 DPS and heal rows up by more than 0.5%, 3 down by more than 0.5%.
  - **Re-sim check:** every row more than 0.3% below origin/main was re-simmed at 3,000 iterations over two seeds. Every one came out level; Arcane 20 Horde and Arcane 30 Alliance are within error.
  - **Tanks:** level or better on tank score.
  - **Run time:** about 2x (1,021 to 2,008 s for 28 specs).
  - S a08601096, f3697b72f.

## Integration

- **Fork `int/r8`** merges lane/r8-healers, -casters, -physical, -hybrids, -warlock, -auras and -items. Pipeline, page, ranker and wasm made no engine change. The suite passes with `-tags=with_db`, except `sim/web`, which also fails setup at the old pin.
- **The pin is 40021374a.** data/proto/ENGINE_SHA follows it, and the proto is unchanged. The rotation ladder goldens are regenerated at it with `make simdb` and a scratch go.work: all 27 DPS and healer ladders and warrior-protection move.
- **The site** merges wt/r8-pipeline, -page, -warlock, -casters, -physical, -hybrids, -healers, -auras, -items, -wasm and -ranker.
- **Gates, all exit 0:**
  - `go vet` and `go test ./...` in sim
  - the coach and sims packages in api
  - `make apl-check` against the pin
  - `make styles-check`
  - tools' `test_wasm_gzip.sh`
  - the full data pytest
  - web sync, the full vitest, astro check, lint and format:check
  - a real-data build: the character chunk is 17,260 of 18,432 bytes
- **Fixed during the gates:** I1.

## Before/after

Pin cd4dfdcc0 with origin/main (170b7b84c), against pin 40021374a with this branch. Both run on today's published data: the BiS sets and builds, and the simdb built from the committed inputs. Raid preset, band 60, 10,000 iterations. Healers are in effective HPS on the heal profile, tanks in threat per second.

| Spec | Alliance, seed 1 | Alliance % (1 / 50001) | Horde, seed 1 | Horde % (1 / 50001) |
|---|---|---|---|---|
| druid-balance | 565.9 to 578.1 | +2.2 / +2.2 | 567.3 to 579.5 | +2.2 / +2.2 |
| druid-feral-bear | 864.6 to 851.9 | -1.5 / -1.4 | 868.6 to 856.3 | -1.4 / -1.4 |
| druid-feral | 620.9 to 613.8 | -1.1 / -1.2 | 616.3 to 608.9 | -1.2 / -1.2 |
| druid-restoration | 772.7 to 779.6 | +0.9 / +0.9 | 770.3 to 777.3 | +0.9 / +0.9 |
| hunter-beast-mastery | 670.9 to 658.1 | -1.9 / -1.9 | 674.8 to 661.8 | -1.9 / -1.9 |
| hunter-marksmanship | 662.7 to 649.2 | -2.0 / -2.0 | 666.7 to 653.1 | -2.0 / -2.0 |
| hunter-survival | 769.3 to 756.3 | -1.7 / -1.7 | 772.6 to 759.5 | -1.7 / -1.7 |
| mage-arcane | 589.5 to 597.6 | +1.4 / +1.3 | 592.5 to 601.4 | +1.5 / +1.5 |
| mage-fire | 653.2 to 663.0 | +1.5 / +1.6 | 660.0 to 670.8 | +1.6 / +1.7 |
| mage-frost | 583.1 to 594.7 | +2.0 / +2.0 | 589.8 to 601.5 | +2.0 / +2.0 |
| paladin-holy | 656.9 to 662.7 | +0.9 / +0.9 | 642.0 to 648.2 | +1.0 / +1.0 |
| paladin-protection | 995.3 to 1001.2 | +0.6 / +0.5 | 1011.5 to 1018.2 | +0.7 / +0.5 |
| paladin-retribution | 662.4 to 661.0 | -0.2 / -0.2 | 659.6 to 658.8 | -0.1 / -0.1 |
| priest-discipline | 610.0 to 619.6 | +1.6 / +1.5 | 614.2 to 620.0 | +0.9 / +1.0 |
| priest-holy | 682.8 to 687.0 | +0.6 / +0.7 | 670.8 to 674.1 | +0.5 / +0.6 |
| priest-shadow | 707.8 to 717.5 | +1.4 / +1.4 | 714.1 to 724.2 | +1.4 / +1.4 |
| rogue-assassination | 623.5 to 611.7 | -1.9 / -1.9 | 620.6 to 609.3 | -1.8 / -1.8 |
| rogue-combat | 618.3 to 605.8 | -2.0 / -2.0 | 618.3 to 606.0 | -2.0 / -2.0 |
| rogue-subtlety | 542.2 to 532.6 | -1.8 / -1.8 | 542.0 to 532.6 | -1.7 / -1.7 |
| shaman-elemental | 465.1 to 468.2 | +0.7 / +0.7 | 473.0 to 476.8 | +0.8 / +0.8 |
| shaman-enhancement | 600.2 to 594.7 | -0.9 / -0.9 | 599.7 to 594.0 | -1.0 / -1.0 |
| shaman-restoration | 823.3 to 841.5 | +2.2 / +2.2 | 820.7 to 838.9 | +2.2 / +2.2 |
| warlock-affliction | 573.2 to 583.7 | +1.8 / +1.8 | 574.5 to 585.2 | +1.9 / +1.8 |
| warlock-demonology | 639.2 to 649.2 | +1.6 / +1.7 | 642.5 to 654.9 | +1.9 / +1.7 |
| warlock-destruction | 636.4 to 647.0 | +1.7 / +1.7 | 639.1 to 649.7 | +1.7 / +1.6 |
| warrior-arms | 718.0 to 708.9 | -1.3 / -1.2 | 729.9 to 720.8 | -1.2 / -1.2 |
| warrior-fury | 764.1 to 754.4 | -1.3 / -1.3 | 767.8 to 757.8 | -1.3 / -1.3 |
| warrior-protection | 1055.5 to 1049.0 | -0.6 / -0.7 | 1043.7 to 1036.9 | -0.7 / -0.7 |

Where the moves come from:
- **Curse of Recklessness** is Forever's 505 armor with no attack power for the boss (Era's was 640 plus 90). This is about -1.9% for every physical attacker.
- **Raid foods** at Forever's values, with re-picked preset foods:
  - casters get Nightfin Soup
  - strength users and tanks get Smoked Desert Dumplings
  - rogues and hunters keep Grilled Squid, now 1% crit

  This lifts every damage caster by 1.3 to 2.3%. It also explains why Retribution barely moves (-0.1 / -0.2): Dumplings offset the curse.
- **Kindling Stave's Ignition** adds 0.1 to 0.5% where the staff is worn.
- **Healers:**
  - the "damage and healing" set bonuses now heal (Vestments of the Devout 3-piece on every healer)
  - the shaman's 17 + Spirit/5 regen (Restoration shaman +2.2%)
  - the Holy paladin kit's Wisdom in place of Might
  - Lightwell's single renew, Divine Grace and Contingency Plan
- **Tanks:**
  - Protection paladin +0.6% (foods and debuff ranks)
  - Protection warrior -0.7%
  - bear -1.4% (the curse)

The bare bands move far more than band 60. Each lane report has its tables. The largest moves:
- Balance bare 40 and 50, +23% and +26%. The Starfire-only filler left a bare druid out of mana for half the fight.
- Restoration druid bare 20 -24% and 30 -16%: Innervate at 40, plus the level penalty.
- Holy paladin bare 20 to 50, +11 to +20%: Wisdom in its kit.
- Priests at band 20, -6 to -10%: Fortitude and Divine Spirit by rank.
- Protection warrior bare 20, -16%: Sunder by rank.

What this table does not hold. Each needs data.yml's regeneration and the bis.yml re-rank, both held by the owner until a zero round.
- **Item data:**
  - stat 45 as healing too (the published healer sets were picked with those items at zero healing)
  - Wowhead percentages kept to the decimal
  - stat 50 as bonus armor
  - round 7's held items, feral attack power and supplement wand schools
  - paladin axes in `items/paladin.json`
- **Item effects at the leveling bands** (items.md):
  - Manual Crowd Pummeler: Arms 30 -23%, Enhancement 30 -19%, Retribution 30 and 40 about -18%, Feral 30 to 50 -12 to -15%
  - Aegis of the Scarlet Commander: Protection paladin 40 +3.8%
  - Fiery War Axe: Enhancement 40 +2.5%
  - Gut Ripper and Gloves of Holy Might: Survival 40 +2.3%
- **Re-ranks:** the ranker's geared weights, paladin axes, the stat trinket, the new Survival single build, the healer sets with healing on stat 45, and the heal-profile coverage check after them.
- **Loot:** the relabelled PvP rewards and the Stone Guard's vendor source (3,213 to 3,227 sources).

## wow-player review

Verdicts on the lane reports, the before/after table and the shots:
- **SHIP:** warlock, items, hybrids, pipeline, wasm and page.
- **SHIP WITH FIXES:** casters, healers, auras, physical and ranker. Each fix, and how it was settled:
  - **Balance B2 is a rotation change, not data.** It stays a confirmed bug: the published bare rotation ran dry for half the fight, which no player would do. The guide's raid rotation is unchanged, because a raid-buffed druid keeps both DoTs and casts Starfire; the mana-pace Wrath fires only below the raid's regen.
  - **The sub-20 level penalty rests on server cores only.** It is marked unverified in healers.md.
  - **The published healer sets were picked with stat 45 worth no healing.** This is listed above; the re-rank fixes it.
  - **Armor-debuff stacking (D5) is open.** The engine keeps Era's rule. Listed as open below.
  - **settings.ts's faction-gate comment:** fixed (I6).
  - **The Survival Raptor Strike note:** fixed in the curated APL; the page follows at the refresh.
  - **Pummeler charges:** the engine models the 3 min cooldown, not the 3 charges. A fight would need to run past 6 minutes to spend all three, and every published fight is 180 s, so the charges never bind (E febc3c25a).
  - **Ranked PvP rows keep their rank label.** Only rank 0 or missing reads "PvP reward" (S 840f51d97).
  - **Paladin axe picks** wait for the regenerated items/paladin.json.
  - **The Survival line.** It was re-shot with the corrected wording (`guide-hunter-survival-raid-card-{1440,390}-after-int.png`). Strength of Earth stays in it: the published Survival number predates round 7's totem fix, so both the totem and the curse lower it at the refresh. The line is time-bound; take it out after launch, or once the refreshed number has been live a while.
- **The /tiers/healer shot shows today's published HPS, not this table.** That is expected under the hold: the page reads the nightly-owned tiers. After the refresh the order changes (shaman ahead of druid on these numbers), and the healer caveat matters then.
- **Numbers players will question first:**
  - Balance bare +23 / +26%
  - Restoration druid bare 20 -24%
  - Curse of Recklessness at 505 with no attack power
  - Strength of Earth at 42
  - Kindling Stave's one-a-minute Ignition
  - the healer order after the refresh

None is a bug. Each row's evidence is in its lane report.

## Still open

- **Owner decisions:**
  - Open content-phases' "battleground-reputation" and "darkspear-islands" gates at launch? And keep the `title_rank` rank-10 cap now that Forever's PvP gear is not rank-locked?
  - Rotation policy: published rotations always cast the top learned rank. Elemental runs dry at every bare band; downranking Lightning Bolt measured +14% at band 50. Below 60, the rank rewrite lifts the deliberate band-60 Renew and Power Word: Shield downranks.
  - Demonology bands 30 to 50 cast a sacrifice that the next summon cancels. Dropping it gains 0.0 to 0.3%; a build without Demonic Sacrifice is the clean fix.
  - wasm: brotli through a Worker route (about 1.2 MB), or no-inlining (a speed trade).
  - The Survival guide's time-bound line, after launch.
- **Held under the adoption rule:** Marksmanship +0.9%, Affliction's Improved Bane of Agony +0.7%, Destruction's Immolate below Conflagrate, Fire's Improved Scorch upkeep +0.4%, Assassination leveling +3.6% (drops Murder and Venom), Discipline's Inspiration.
- **Blocked on logs** (launch is 2026-11-04):
  - the one-a-minute rate for every "chance on" proc with no client rate (Ignition, Necropile, Black Menace, Verigan's Fist, Toxic Revenger, Gloves of Holy Might, Blood Siphon, and the 26 weapon procs' rates)
  - armor-debuff stacking under Forever's aura 674
  - the sub-20 level penalty
  - Swiftmend's amount
  - Faerie Fire's form threat
  - every round-6 and round-7 row still listed
- **Single source or open:**
  - tank weight sweeps disagree between seeds at any iteration count, and some fail ("stamina weighs nothing"); the dual ranking keeps the tank sets at or above origin/main
  - one crit stat for melee and spell (a sharpening stone also raises spell crit)
  - Mana Tide's per-tick mana
  - Retribution Aura, Thorns and Stoneskin ranks (not in any preset)
  - the Satchels' thrown damage
  - Red Dragonscale Leggings' half point of Intellect
  - Concentration Aura, Healing Wave rank 10, Chain Heal's party-only jumps, Power in Light, forever-bugs #419
- **Page and pipeline:**
  - `real-engine` does not gate deploy yet
  - `allow_shrink` stays while data.yml passes it to `loot`
  - the coach catalogue still names the old foods
  - the custom /sim panel still offers the four caster-only debuff flags, which now do nothing
  - the Discord bot skips an unranked PvP line
