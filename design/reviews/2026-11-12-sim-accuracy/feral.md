# Sim accuracy, 2026-11-12: Feral druid (cat and bear)

Lane brief: `design/reviews/sim-accuracy/brief.md`. Site from origin/main a38608c67, fork from master 15e2c6bae (the pin). Client on the site: 1.60.1.70790; live client 1.60.1.70888 with its hotfix push (wago.tools, `useHotfixes=1`). Scope: Feral cat (DPS) and Feral bear (tank), levels 10 to 60, bare bands and raid 60, against the client tables, Blizzard's notes to 12 November and every Feral druid in the public logs. The Feral specs had the most churn since their last full check: the 11 November physical-logs lane read the cat against logs (0 fixed), the tanks lanes moved the bear's fight length, Maul gate and boss scaling, and the 12 November melee lane's rage refund (MR65) moved the bear's goldens.

**Totals: 5 confirmed bugs, all fixed, all old: Rake scales with the paw and with attack power (Blizzard's 3 November line and five logged cats; the engine had the flat amounts alone, so every earlier measurement that dropped Rake ran on the wrong Rake), Lacerate's hit counts the stack it adds, Rend and Tear reaches a bleed put on a bleeding target, Pounce Bleed and Thick Hide rank 1 on the 70790 rows. The cat keeps Rake up again: +19 to +20% at bare 30 to 60 and +10.6% at raid 60. Bear threat +10 to +12% at 25 to 35, +7 to +9% bare 40 to 60, +5.4% raid 60; chance of death unchanged. The bear still holds after MR65 (chance of death 1.8 to 3.7% at 30 to 35, where tanks-below-45 left it). Engine pin 15e2c6bae to 4f8bb748d.**

## Method

1. **Client.** Every spell, talent, form and proc the two cards and builds touch, read from the full wago.tools DB2 dumps for 1.60.1.70790 and 1.60.1.70888 (both with the hotfix push layered): SpellEffect, SpellMisc, SpellPower, SpellCooldowns, SpellLevels, SpellCategories, SpellAuraOptions, SpellClassOptions, SpellDuration, Talent, TraitDefinition, TraitDefinitionEffectPoints and CurvePoint; the site's `talents/druid.json` and `trainables/druid.json` for the rendered rank texts and the learn rows. Every druid row (730 spells: SpellClassSet 7, the druid talent ranks and two hops of their triggers) was diffed between the two builds (`feral/client_diff.py`).
2. **Blizzard.** The Forever class-changes thread (t/2360696) to its 3 November post, and the hotfix thread to 12 November.
3. **Logs.** foreverlogs.gg reports 2815 to 2849 (5 to 11 November, Ragefire Chasm to Zul'Farrak): every damage, cast, avoidance, resource and buff event from all 185 druid rows, with the roster snapshots (level from the talent points, stats, gear with each weapon's DPS). 24 of the rows are Feral Combat. Our own API holds nothing newer than the 8 November advanced report, whose group has no druid. The logs carry no energy stream, no threat and no debuffs on creatures, so a Lacerate's stacks are inferred from the cast timing (each cast refreshes 15 s) and a Rip's combo points from the combo-point gains since the last finisher.
4. **Sims** on the Windows box (`tools/remote-test.sh --pin-engine` or `--engine-sha`, `feral/jobs.sh`, one spec-breakdown at a time, `nice -n 15`, GOMAXPROCS=5, 10,000 iterations, seed 1 unless stated, both factions' published races): the cat's bare bands 20 to 60 and raid 60, the bear's tank fight at 20, 25, 30, 32, 35, 40, 50, 60 and raid 60, at the pin and on the fixed fork, then the rotation variants.

## Findings

| # | finding | source | origin | fix |
|---|---|---|---|---|
| C1 | **Rake scales with the paw and with attack power; the engine had the client's flat amounts alone.** Blizzard, 3 November: "Rake: Base points and attack power scaling increased by about 25%" (the base half is the 70755 push the engine already carries; round 22 C1 kept the attack power half out because the tooltip does not render it). Five logged cats at 24 to 30 (Valdo, Losemid, Insomnia, Lez; Boofaholic a Balance druid in Cat Form), each on a 20.3 DPS staff with attack power 179 to 296 off their stats: the non-crit Rake hit less the rank's 27, divided by the armor factor their white hits give, is 0.93 to 1.01 of the paw's weapon damage (the hit is the paw plus the flat, as Claw is at 140%); the non-crit tick less the rank's 20 is 0.070 to 0.089 of their attack power on the four Feral cats (0.052 on the Balance druid). The alternative, a flat attack power share on the hit, would make the hit's paw ratio fall from 1.07 to 0.81 across that attack power range; it does not. | Blizzard t/2360696 (3 Nov); `feral/logs3.py`; tooltip 9904 | old | F 99db16987, `rake.go`: the hit is weapon damage plus the roll, each tick adds 8% of attack power (`rakeTickAttackPowerShare`) |
| L1 | **Lacerate's hit counts the stack it adds.** The engine landed 10% weapon damage for each stack already on the target, so the first application hit for nothing. A level 43 bear (Jam, 2842; white 124.9 after armor) hits for 13.2, 25.7, 38.3, 55, 70.2 and 77.3 with 0 to 5 stacks already up: 0.106, 0.206, 0.307, 0.44, 0.56 and 0.62 of the paw, 10% a stack with the new one counted, 60% on a full stack. A level 31 bear (Moon, 2833) reads 0.092, 0.18 and 0.25 at 0 to 2. | `feral/logs4.py`; 414647 (effect 31, 20%) and the live text | old | F 99db16987, `lacerateHitDamage` |
| L2 | **Rend and Tear reaches a bleed applied to a target that already bleeds.** The engine put it on direct damage only. The same bear's Lacerate ticks (rank 6 a tick, the level 34 rank it still used at 43): 6.1 at one stack (the first application, no bleed on the target as it landed), then 14, 28 and 34.6 at two, four and five stacks against 12, 24 and 30, 1.15 to 1.17 times the rank, the talent's 5/5. A bear at 31 without the talent ticks 20 at five stacks of the 4 rank, exactly the rank. | `feral/logs4.py` | old | F 99db16987: Lacerate, Rake, Rip and Pounce snapshot the multiplier read before the bleed lands (`bleedRendAndTear`), so a bleed never counts itself; Rake, Rip and Pounce follow the rule Lacerate measured |
| P1 | **Pounce Bleed ticks 19, 25 and 31 in 1.60.1.70790** (9007, 9824, 9826); the engine had 70009's 15, 20 and 25. Blizzard's 3 November Pounce line is the Rake line's twin; its attack power half is unmeasured (two logged Pounce ticks) and stays out. No fight casts Pounce. | SpellEffect 70790h, 70888h | old | F 99db16987, `PounceTickDamage` |
| T1 | **Thick Hide's armor per level is the same at every rank.** The engine gave 1, 2 and 3 a level by rank. The talent spell's effect 1 is 1/2/3 by rank (curve 111990) plus 3 a level above level 1 (EffectRealPointsPerLevel 3 on the row), 180 at 60 on every rank as the rendered text says ("180 additional base Armor" at ranks 1 and 3 alike). The top rank is unchanged (3 + 177 = 180 = 3 x 60); rank 1 at level 16 is 46 base armor, not 16. The bear walk has ranks 1 and 2 at 16 and 17 only; no band or published number moves. | 16929 effect 1; TraitDefinitionEffectPoints 25097, CurvePoint 111990 | old | F 99db16987, `thickHideBaseArmor` |

Natural Instinct's mask in 70888 adds the Lacerate hit spell (414647, word 3 bit 3) and the Pounce Bleed spells (word 3 bit 7, which 70888 gives them); the engine already reaches both through the spell's crit multiplier, which its bleed ticks share (round 10 C11). Aligned, no change.

## Checked, nothing to change (client)

Every row below was read in 70790h and 70888h; the two builds differ on no number the cards use (the druid diff is Nature's Swiftness's mask, Natural Instinct's mask above, the Poised server dummies 1330763 and 1330765 that the rogue lane recorded as MR60, Pounce and Feral Charge (Cat) attribute bits, and an item set's lost class bit).

- **Forms.** Cat Form 768 (55% of base mana, 1.5 s global, learned 20) and its passive 3025 (attack power 12 plus 2 a level from level 6, the engine's 2 a level; threat -29%; Faerie Fire free, +6 s cooldown, -0.5 s global); Bear Form 5487 / 1178 (armor from items +180%, health 20 + 18 a level to 40, attack power 30 + 3 a level); Dire Bear Form 9634 / 9635 (+360%, 600 + 32 a level, 120 + 3 a level, learned 40); Bear Form (Passive2) 21178 threat +115%. Tiger's Fury, Natural Weapons and Primal Fury are not in Forever: the trainables table has no Tiger's Fury row (its only SkillLineAbility row is a Season of Discovery rune, round 1), and the 70790 trees carry neither talent; Blood Frenzy (16958, whose rage proc the logs name "Primal Fury") is the rage-on-crit and combo-point-on-crit talent.
- **Energy and rage.** Energy regenerates at 10 a second (the rogue's advanced log, physical-logs L2; the cat has the same bar). The bear's paw is a 2.5 s one-hander for normalized rage (8.65 a landed swing, 17.3 on a crit); the public logs carry no white-swing rage gains, so the swing's amount rests on the warrior's advanced log (L1 there). Furor pays 10 rage on every logged Bear Form shift that paid (68 of 68, "Furor" gains). Blood Frenzy pays 5 rage (488 of 492 "Primal Fury" gains, the other four partial bars), after Swipe, white and Maul crits.
- **Costs.** Maul 150 (15 rage), Swipe 200, Lacerate 150, Primal Bite 200 with its 6 s category cooldown, Demoralizing Roar 100, Shred 60, Claw 45, Rake 40, Rip 30, Ferocious Bite 35, Ravage 60, Pounce 50, Shifting Power 55% of base mana with 40 energy on an 8 s cooldown and a 1 s global; Ferocity's two effects (-10 tenths on Maul, Swipe and Primal Bite; -1 on Claw and Rake) and Shredding Attacks' (-6/-12/-18 on Shred, -10/-20/-30 tenths on Lacerate, curves 85350 and 85349: 1, 2 and 3 rage).
- **Amounts.** Maul +18/27/37/49/71/101/128; Swipe 18/25/36/60/83 on up to three targets, the engine's 3% of attack power (Blizzard, 1 October) checked on twelve logged bears against their white hits (Jam at 43: 70.9 logged against 74 predicted with Savage Fury and Feral Instinct 3; Moon at 31: 31.6 against 34.8); Lacerate 4/6/10/12/15 a tick over 15 s, five stacks; Primal Bite +42/62/77 at 100% weapon damage; Demoralizing Roar -46 to -193 and its threat 8 to 42 plus 0.8 a level; Faerie Fire 175/285/395/505 armor; Shred +22 to +80 at 200%; Claw +27 to +115 at 140%; Ravage +42 to +98 at 435%; Rake 27/39/56/76 and 20/26/32/42; Rip 4 to 19 a tick plus 5.5 to 31.9 a point over 12 s; Ferocious Bite's base, per-point and per-energy columns on every rank; Enrage (10 rage at once, 2 a second for 10 s, 60 s cooldown, no global); Frenzied Regeneration (1% a rage every second for 10 s, 3 minutes); Barkskin (-20% physical for 15 s, 60 s).
- **Talents.** Savage Fury's two effects (damage on Maul, Swipe, Claw, Rake and Shred; periodic damage on Rake's bleed), Feral Instinct's Swipe bit, Sharpened Claws, Predatory Strikes, Heart of the Wild (10% Intellect, 20% Stamina in bear, 10% Strength in cat at 5), Naturalist, Nature's Majesty, Leader of the Pack 3%, Genesis's mask (Rake, Rip, Pounce, Lacerate, Moonfire, Insect Swarm and the heals), Natural Reaction, Rend and Tear 3% a rank, Omen of Clarity (10 s cooldown, the 2 a minute fit the logs confirmed on the cat; bears proc it on 4.5% of their hits with Swipe's three targets each counted), Poised's -10 on the cat's families.
- **Maul in the logs.** Thirteen bears' non-crit Maul over their non-crit white hit: 17 to 26 at rank 2 (flat 27), 35 to 37 at rank 3 (37), 70 at rank 5 (71), the flat under the target's armor. Maul is the paw plus the flat, nothing more.
- **Rake in the logs**, beyond C1: Rake's hit over Claw's on the same cat is 0.70 to 0.72 at the same rank (predicted 0.71 with both on the paw), 0.59 for Insomnia on Claw rank 2 (predicted 0.59).

## Not settled by these logs

- **Rip's and Ferocious Bite's attack power shares (round 22's 1.25% and 3.75% a point).** The combo points at a finisher can only be counted from the builder gains since the last finisher, and a dungeon cat changes targets, so the count runs high: a cat at 24 reads 36.8 a tick at a counted five points where 31.5 plus the share predicts 43, and 37.5 at four (37.6 predicted). The one cat with a clean run (Losemid, 8 Rips) reads 47.5 against 45 predicted. Open until a raid log holds one target.
- **Faerie Fire (Feral)'s threat.** The engine gives the feral cast 108 flat threat, Classic's server value; every client rank's threat effect (63) reads 0, and Blizzard's 16 October note says only that it "now generates threat" in bear forms. Logs carry no threat. Open, worth under 1% of a bear's threat.
- **Pounce's attack power half** (P1).

## Before and after

All on the box, 10,000 iterations (the variants' sweep at 5,000), seed 1, the published entries and rotations; `feral/base_pin.txt` (pin 15e2c6bae) against `feral/after_fix.txt` (fork 99db16987, the same rotations), then the Rake line.

**Cat, the published Shred-only rotation on the fixed engine** (Rip's ticks take Rend and Tear over the raid's bleed; nothing else the list casts moved): bare 20 0.0%, bare 30 0.0%, bare 40 +0.2 to +0.3%, bare 50 +0.4%, bare 60 +0.9%, raid 60 +3.4 to +3.5% (Alliance / Horde).

**Cat, the Rake line over the published rotation**, both on the fixed engine, seeds 1 (5,000) and 50001 (10,000), Alliance / Horde:

| band | seed 1 | seed 50001 |
|---|---|---|
| bare 20 (level 20, before Rake) | 0.0 / 0.0% | - |
| bare 30 | +19.0 / +19.1% | +19.0 / +19.2% |
| bare 40 | +20.2 / +20.1% | +20.4 / +20.1% |
| bare 50 | +20.0 / +20.3% | +20.0 / +20.4% |
| bare 60 | +18.2 / +19.2% | +18.1 / +19.1% |
| raid 60 | +10.6 / +10.8% | +10.6 / +10.8% |

Against the old pin, the whole landing is therefore bare 30 +19%, bare 40 to 50 +20 to +21%, bare 60 +19 to +20%, raid 60 +14.4 to +14.7%.

**Bear, the tank fight**, damage taken a second / TMI / chance of death / threat a second, pin to fixed, Alliance (Horde in brackets):

| level | before | after | threat |
|---|---|---|---|
| 20 | 54.6 / 22.6 / 0.0% / 88.9 | 54.7 / 22.6 / 0.0% / 88.6 | -0.3% (-0.2%) |
| 25 | 61.7 / 22.9 / 0.0% / 120.3 | 61.7 / 22.9 / 0.0% / 132.6 | +10.2% (+10.3%) |
| 30 | 106.4 / 29.6 / 3.7% / 217.7 | 106.5 / 29.6 / 3.7% / 243.6 | +11.9% (+12.0%) |
| 32 | 107.9 / 28.7 / 3.3% / 225.2 | 108.0 / 28.7 / 3.0% / 252.2 | +12.0% (+12.0%) |
| 35 | 110.3 / 27.5 / 3.2% / 244.3 | 110.5 / 27.6 / 3.2% / 272.5 | +11.6% (+11.7%) |
| 40 | 148.0 / 29.2 / 0.0% / 317.9 | 148.4 / 29.3 / 0.0% / 343.8 | +8.2% (+8.3%) |
| 50 | 233.5 / 31.6 / 0.0% / 491.6 | 234.2 / 31.8 / 0.0% / 534.3 | +8.7% (+8.9%) |
| 60 bare | 392.0 / 47.1 / 2.3% / 798.2 | 392.4 / 47.1 / 2.3% / 857.2 | +7.4% (+7.5%) |
| 60 raid | 350.7 / 28.5 / 0.0% / 1727.5 | 350.9 / 28.6 / 0.0% / 1821.8 | +5.5% (+5.4%) |

The threat comes from the first Lacerate landing weapon damage (and every later one landing one stack more) and from the bleed under Rend and Tear from 36; at 20 the walk has no Lacerate yet and the difference is Thick Hide's rounding. Chance of death at 30 to 35 is 1.8 to 3.7% on both engines (2 to 4% in tanks-below-45), so the bear holds after MR65 and after this landing.

## Rotation: Rake

On the fixed engine, three lines were tried at every bare band from 20 and at raid 60, both factions (`feral/after.txt`, 5,000 iterations, seed 1; the winner again at seed 50001, 10,000):

- **v1, Rake when it is not on the target, after the finishers and before Shred:** the table above.
- **v2, Rake ahead of the finishers:** the same as v1 within 0.1 DPS at every band (the Rake line and the finishers never compete for the same global in practice).
- **v3, Rake only while Rip is up:** +9.8 to +16.2% bare, +9.9 to +10.0% raid: 4 to 9 points behind v1 at bare 30 to 60, 0.7 behind at raid 60.

The published list takes v1 (`data/curated/apl/druid-feral.json`, synced to the fork's `ui/feral_druid/apls/forever_feral.apl.json` and `sim/request/apl/druid-feral.apl.json`): Ravage, Faerie Fire, Shifting Power at 60 energy or less, Rip at 5 points under 2 s, Ferocious Bite at 5 points while Rip has 5+ s, Rake when it is off the target, Shred, Claw at 20 or in front. It drops nothing the old list cast, gains beyond the error at every level from 24 (Rake's trainer level) and at raid 60 on both seeds, and the band 20 entry at level 20 is unchanged. Red first: `TestFeralCatKeepsRakeUp` (S df5e3397e, red on the box at the old list; green on the landing commit). The guide's rotation, Strengths, talents, stat and leveling sections, the card notes and the coach starter follow; the wow-player HELD once on the Rake note's "40 energy" (Ferocity makes it 35) and the bear note's "sixth stack", then SHIPped on the wording in this landing. The builds do not move here: Ferocity 5 is already in both cat builds (the tree's first row needs the points), and whether the raid build now values Genesis, Ferocity or Savage Fury differently is a talent-search question for the refresh (MR69).

## Measure at refresh

| what | specs / bands | why |
|---|---|---|
| MR69: the cat on Rake | druid-feral bare bands 20 to 60 (levels 24 to 60) and raid 60: picks, weights (attack power and Strength now reach Rake), the raid talent search (Genesis, Ferocity, Savage Fury against the rest), the Rake and Claw rank rows | C1 and the Rake line, +19 to +20% bare, +10.6% raid |
| MR70: the bear on Lacerate's first stack and the bleeds' Rend and Tear | druid-feral-bear bare bands 20 to 60 and raid 60: threat, the tank weights (Strength, attack power and weapon DPS reach every Lacerate now), the published metrics; the Maul gate was tuned on the old Lacerate (round 30 B1, tanks-below-45 B2) and should be re-swept | L1, L2: threat +10 to +12% at 25 to 35, +7 to +9% at 40 to 60, +5.4% raid |
| MR71: the Pounce attack power half | druid-feral, the opener (Pounce against Ravage from Prowl) once a log with Pounce ticks at a known attack power exists | P1, Blizzard's 3 November Pounce line |

## Gates

On the box, one run at a time:
- **Fork, red then green:** at the test sha the `druid` package does not build (`rakeHitDamage`, `rakeTickDamage` and `thickHideBaseArmor` do not exist yet), `TestRendAndTearReachesRipAndRakeOnABleedingTarget` reads 1.0000 for 1.15 and a 42.00 Rake tick for 148.64, and `TestRendAndTearReachesLaceratesBleedFromTheSecondStack` reads 30 for 34.5. At the fix every druid package passes with `-tags with_db`; the bear suite golden (`TestFeralTank.results`: average DPS 685 to 717, TPS 2711 to 2850, and the stat weights) was regenerated in a writable copy on the box and committed; `TestP1Feral` is unchanged; `core` and `conformance` pass.
- **Site, red then green:** `TestFeralCatKeepsRakeUp` reads "Rake 0.0 casts in a 180 s fight" on the old list and passes on the landing commit. `go vet ./...` clean; `go test` for request (the ladder goldens regenerated there with `FOREVER_UPDATE_GOLDEN=1`, then held), leveling, score, rankchoices, specs, cmd/leveling-bis, cmd/addon-builds and cmd/spec-breakdown green at the pin.
- **Ladder goldens:** druid-feral rungs 30 to 60 move +12 to +20% (Rake 1822 to 9904 among the top casts from 30), druid-feral-bear rungs 30 to 60 move +6 to +9%; the level 10 and 20 rungs and the violations lists are unchanged.
- **Web:** `npm test` on the sim, guides, rotation-lists, copy, leveling-builds, apl-gates and rotation-icons suites: 102 files, 2,158 tests green (the first run held the cat guide to four Strengths bullets and the rotation's Rake ids to an icon row, both fixed); `npm run lint` clean; `astro check` 0 errors, 0 warnings. No Playwright surface changed.
- **Data:** `uv run pytest` on test_card_note_levels, test_addonrotation, test_genproto and test_addondata: 129 passed; the three failures are the standing ones every lane sees until the nightly (the allowlist lines of other specs' reworded notes, the vendored pin against the artifacts built at 2b10bae6c, the Data.lua idle budget), none a Feral card.
- `make apl-check`: the fork and site copies of both Feral rotations match the curated files.

## Sim changes in this landing

- **Fork `lane/sim-feral`:** 1e89716d5 (tests, red), 99db16987 (Rake, Lacerate, Rend and Tear on bleeds, Pounce, Thick Hide), f7db82b4a (the bear suite golden and the Thick Hide rank 1 test), a54d67ca6 (the cat rotation synced from the site), measured on master 15e2c6bae; rebased onto master fc2ed8cd3 (the warlock lane, warlock files only; the Arcane lane pinned that same sha) as 08aaa1c7d, 3a876eded, 204f24a51 and 4f8bb748d. Master fast-forwarded to 4f8bb748d.
- **Site pin:** 4f8bb748d (`sim/enginever/version.go`, `web/src/lib/sim/version.ts`, the sim fixtures).
- **Site:** `data/curated/apl/druid-feral.json` (the Rake line, its panel line, the file's notes) and the Lacerate note in `druid-feral-bear.json`, their `sim/request/apl` copies, both guides, the cat coach starter, the two ladder goldens regenerated on the box, `feral_rake_test.go`, the four Rake rank rows added by hand to `spellicons.json` (the icons step's own output once a rotation casts Rake; the art is already in the build), this report and its scripts and tables, the ledger row.
- **No regen:** no BiS rank, data run, simdb rebuild or Data.lua render; `addon-data.json` keeps the old card lines until the nightly (the new Rake line and the reworded notes reach the addon then).

## Sources

- Blizzard: t/2360696 (the class-changes thread, posts to 3 November 2026), t/2403040 and t/2404092 (hotfixes to 12 November).
- wago.tools DB2 CSVs for 1.60.1.70790 and 1.60.1.70888, `useHotfixes=1`; TraitDefinitionEffectPoints and CurvePoint for 1.60.1.70790.
- nether.wowhead.com/forever/tooltip/spell/9904 (read 12 November).
- foreverlogs.gg reports 2815 to 2849; foreversixty API `/v1/reports/recent`.
