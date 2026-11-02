# Sim audit round 10: heal profile by band

Lane: wt/r10-healprofile (site, from origin/main bb619a188) and lane/r10-healprofile (fork, from 4d3129cae). Scope: owner ruling 1 (scale the heal profile below 60 by band), the sub-60 healer rows of data/curated/rank-choices.json, and the sub-60 healer talent builds. The healers lane's branches (wt/r10-healers 7e80b570c, lane/r10-healers 041bed648) are merged in at the integrator's request, and every rank row was confirmed again on the merged state.

## The rule

The rule is stated once, in `bands.rule` of data/curated/heal-profile.json. sim/request/healprofile_bands.go is the only code that applies it.

1. **Same fight at the band's level.** Below 60 each band fights the level-60 fight at its own level. Length (300 s), cadence (a 3.3 s tank swing, a 5 s pulse on 3 of 5 members), spread and members stay the same. Only the damage and the health change.
2. **A band's content.** It is the open dungeons whose last boss stands from the band's level to five levels above it. At 60 it is the open raid (Onyxia), because the level-60 profile is a raid fight.
3. **A band's damage.** Take the median of its last bosses' melee damage per second, times the content's client damage modifier. A boss's melee damage per second is its creature_template hit (attack power / 14 + mean weapon damage x damage multiplier, as in cmangos `Creature::UpdateDamagePhysical`) divided by its attack time. The tank hit and the raid pulse both scale by this, measured against the level-60 band. The tank-to-raid ratio stays the same.
4. **Health.** Tank and member health scale by the client's expected player health at the band's level, against level 60.
5. **Tuning and in-between levels.** The tuning scale (1.85, bounds 1.65 to 1.85) multiplies every band alike. A level between two bands fights the band at or below it, the same way rank choices run. A level below 20 fights band 20.
6. **Checked at load.** Each band row states its derived figures. Loading refuses any row whose figures do not follow from the rule.

Level 60 is unchanged. The top-level figures are the 60 row's, and the test pins them.

| band | content (last boss, level) | boss melee DPS, median | x mod | player health | tank health / hit | members / pulse |
|---|---|---|---|---|---|---|
| 20 | The Deadmines (Edwin VanCleef, 21), Wailing Caverns (Mutanus the Devourer, 22) | 47.0 | 1.2 | 503 | 1,836 / 124 | 966 / 32 |
| 30 | Razorfen Kraul (Charlga Razorflank, 33), Gnomeregan (Mekgineer Thermaplugg, 34), SM Graveyard (Bloodmage Thalnos, 34) | 47.1 | 1.2 | 878 | 3,204 / 124 | 1,687 / 32 |
| 40 | SM Armory (Herod, 40), Razorfen Downs (Amnennar the Coldbringer, 41), SM Cathedral (High Inquisitor Whitemane, 42) | 111.5 | 1.2 | 1,353 | 4,938 / 294 | 2,599 / 77 |
| 50 | Maraudon (Princess Theradras, 51), Sunken Temple (Shade of Eranikus, 55) | 256.5 | 1.2 | 1,928 | 7,036 / 677 | 3,703 / 177 |
| 60 | Onyxia's Lair (Onyxia, 63) | 967.3 | 1.0 | 2,603 | 9,500 / 2,128 | 5,000 / 555 |

Incoming damage per second: 57, 57, 135, 311, 978.

### Sources (checked 2026-11-02)

- **Boss levels and melee figures:** cmangos classic-db `creature_template` at the pinned commit ec4f596 (data/pipeline/classic_quest_levels.py). The figures are MinMeleeDmg, MaxMeleeDmg, DamageMultiplier, MeleeAttackPower and MeleeBaseAttackTime. Each one is stored in the row.
- **Hit formula:** cmangos mangos-classic `src/game/Entities/StatSystem.cpp`, `Creature::UpdateDamagePhysical`.
- **Second source for levels:** Wowhead Forever NPC tooltips (`nether.wowhead.com/forever/tooltip/npc/<id>`). They agree with creature_template on every last boss. Onyxia shows as a skull (63).
- **Client 1.60.1.70707, ExpectedStat** (Classic set, ExpansionID -2, row id = level): PlayerHealth 503 / 878 / 1,353 / 1,928 / 2,603 at 20 / 30 / 40 / 50 / 60.
- **Client 1.60.1.70707, damage modifier:**
  - ContentTuningXExpected rows 3400 to 3428 give every Classic dungeon's content tuning (5237 to 5256, named by LFGDungeons) ExpectedStatMod 489. That modifier sets CreatureAutoAttackDPSMod and CreatureSpellDamageMod to 1.2.
  - The raids' tuning (4509) has no modifier.
  - This rests on a single client table with no Blizzard post behind it. A search found only Blizzard's +20% dungeon kill XP (beta notes, 8 October), which is a different thing.
- **Why creature_template and not the client's creature curve:**
  - The client holds no boss damage. CreatureDifficulty lists 184 creatures and none of these bosses. ExpectedStat's creature columns are one generic curve with no line between a dungeon boss and a raid boss.
  - I first tried scaling the level-60 fight down that curve, and dropped it. At band 20 it asked for 306 damage a second. A bare Holy priest then covered 11% and ran dry at 27 s, which is the fault the ruling names.

## Findings

| # | finding | evidence and source | confirmed? | fix sha |
|---|---|---|---|---|
| P1 | **Every band below 60 fought the level-60 raid fight.** A level-20 tank had 9,500 health and took 2,128 hits. Bare healers covered 3 to 25% and most ran dry 25 to 60 s in, so every sub-60 healer number and rank choice measured healing per mana, not healing. | round 9 healers X3. `sim/request/healprofile.go` had no band input. Before table below. | CONFIRMED-BUG (owner ruling 1) | S e9bd57485 |
| P2 | **The rule now lives in one place.** `Attach` and `AttachWeights` read the healer's own level from the engine request. Every tool that runs a healer fights the same band with no change at its call site: /sim and the API (`applyRoleFight`), the ranker (`inproc.HealingRun`, `score_heal` weights and tuning bounds), spec-breakdown, heal-builds, dps-per-point and the ladder test. `ForLevel` can be repeated and composes with `AtScale`. Tests were written first and failed first: `sim/request/healprofile_bands_test.go` (9 tests). | `go test ./request -run 'HealProfile|HealBand'` | CONFIRMED-BUG (same) | S e9bd57485 |
| P3 | **The BiS page's "What this profile assumes" showed the level-60 tank on every healer band.** A sub-60 band now shows its own figures, plus a Content line naming its dungeons and last bosses with the rule as the reason. Level 60, and files published before the band table, are unchanged. Test written first: heal-view.test.ts. | `web/src/lib/bis/heal-view.ts` | CONFIRMED-BUG (page) | S 8ad647829 |
| P4 | **The healer rotations' notes said a band's rank "is chosen for its mana over a fight sized for level 60".** The five notes now say a band's rank is chosen on that band's own fight. The rank-choices notes now say healer rows below 60 are measured on their band's profile. The fork copies are synced, and apl-check is clean. | data/curated/apl/{druid-restoration, paladin-holy, priest-discipline, priest-holy, shaman-restoration}.json | Aligned (copy) | S 999f14377, E f59d100fd |

## Sub-60 healer rank choices

**Method:**
- Each band's own profile, bare preset.
- Greedy search over every learned rank of each ranked heal in the APL: start from the default rule, take the best single change while it gains 0.3%. 3,000 iterations, seed 1.
- Confirmed at 10,000 iterations on seeds 1 and 50001, alliance, then horde: default, current, the final set, each row alone, and the final set minus each row.
- A row stays only if it beats the default rule at both seeds and adds to the set.
- A downrank (below the default rule's rank) needs a fight that runs dry under the default rule. An uprank does not.
- No row turns a dry fight into one that ends with mana unspent.

Tools: scratch r10/healprofile/tools/rankband.py and decide.py, with a scratch spec-breakdown that reads the rank table from a file.

All rows were measured again after the healers lane merged in. They came out the same, plus one new row (Chain Heal).

| spec | levels | before | after | gain vs default rule, seed 1 / 50001 (alliance) | mana lasts after (s) |
|---|---|---|---|---|---|
| druid-restoration | 20-29 | Rejuvenation 1 | dropped: loses to the default (rank 3) by 4.7% | - | 272 |
| druid-restoration | 30-39 | Rejuvenation 2 | dropped: loses 2.7% | - | 412 |
| druid-restoration | 40-49 | Rejuvenation 7 | kept | +1.35% / +1.35% | 614 |
| druid-restoration | 50-59 | Rejuvenation 8 | dropped: rank 4 (default) is 3.2% better | - | 440 |
| paladin-holy | 40-49 | Holy Light 3 | dropped: no effect (Holy Light is not cast on this fight) | 0 | 389 |
| paladin-holy | 50-59 | Holy Light 3 | kept | +5.76% / +5.71% | 280 |
| priest-discipline | 30-39 | Shield 5, Renew 1 | Shield 5 only; Renew 1 loses | +0.38% / +0.19% | 626 |
| priest-discipline | 40-49 | Shield 2, Binding Heal 1 | Shield 1, Prayer of Healing 1, Penance 1 | set: +4.55% / +4.55% (horde +5.9%) | 259 |
| priest-discipline | 50-59 | Shield 8, Binding Heal 1 | + Heal 4, + Prayer of Healing 1 | set: +5.92% / +5.92% (horde +5.9%) | 159 |
| priest-holy | 40-49 | Flash Heal 4 | Flash Heal row dropped (no effect); + Prayer of Healing 1 | +0.39% / +0.39% | 283 |
| priest-holy | 50-59 | Flash Heal 6 | + Prayer of Healing 1, + Lightwell 1 | set: +8.98% / +8.75% (horde +5.5%) | 180 |
| shaman-restoration | 20-29 to 50-59 | Healing Wave 3 at every band | all four dropped: each loses to the default by 0.3 to 6% | - | 258 / 3,460 / 635 / 270 |
| shaman-restoration | 50-59 | - | Chain Heal 1 (merged state) | +0.32% / +0.28% (horde +0.2% / +0.3%) | 272 |

Notes on the table:
- **Horde, Holy priest at 50.** The set beats the default by 5.5% but trails Flash Heal 6 alone by 0.16%. The rows are faction-wide and alliance gains 3.9% from the extra two, so the set stays.
- **Bands 30 and 40 have nothing to choose.** Every healer covers 86 to 95% of the incoming damage, mostly with mana to spare, so most candidates tie.
- **Ranks are the engine's own.** Every adopted rank is the engine's existing spell (spec-breakdown mana rows: Penance 402174, Prayer of Healing 596, Power Word: Shield 17, Heal 6064, Lightwell 724).

## Swiftmend line (healers lane, at the integrator's request)

Merged engine and site. The new line (cast under 75%, keeps the heal over time) against the round-9 line from 03690eeab^, bare, alliance, 10,000 iterations:
- 20 and 30: equal (Swiftmend is not learned yet).
- 40: equal.
- 50: +0.8% at both seeds.
- 60: +1.6% at both seeds.

The line holds at every band.

## Sub-60 healer talent builds

**How the builds are made:**
- Below 60 the builds are not chosen. `leveling.LadderTalentString` walks the guide's single level-60 build top row down and stops when the band's points run out.
- I searched every one-point move from each published band build (one rank from any spent talent to any talent the move leaves legal) on the band's own profile. 2,000 iterations, then the top moves at 10,000 on seeds 1 and 50001, both factions, on the merged state.
- Bands 30 and 40: nothing moved more than 0.6%, because the fight is covered.

| # | finding | evidence (gain vs published, seed 1 / 50001) | confirmed? | fix sha |
|---|---|---|---|---|
| T1 | Restoration druid 20: one Naturalist point to Bloom. | +16.84% / +16.84% alliance, +19.67% / +19.67% horde | CONFIRMED (build loses beyond error) | - (no per-band build mechanism) |
| T2 | Holy paladin 20: Purifying Power to Illumination. | +5.15% / +5.15%, horde +4.36% | CONFIRMED | - |
| T3 | Discipline 20: Wand Specialization to Meditation. | +3.85% / +3.85%, horde +5.17% | CONFIRMED | - |
| T4 | Restoration shaman 20: Improved Healing Wave to the Water Shield talent. Mana then lasts 708 s. | +3.10% / +3.11%, horde +9.28% | CONFIRMED | - |
| T5 | Holy priest 50: Mental Strength to Spiritual Guidance. | +2.03% / +2.04%, horde +2.16% / +2.22% | CONFIRMED | - |
| T6 | Every other healer at 20 and 50: no move gains 1%. | talent search output | Aligned | - |

These cannot be fixed inside this lane. The walk is shared code for every spec, and a guide states one build. See "For other lanes".

## Coverage, and a question for the owner

Effective HPS divided by incoming damage, bare, alliance, after the change:

| band | druid | paladin | Discipline | Holy | shaman |
|---|---|---|---|---|---|
| 20 | 65% | 58% | 55% | 66% | 86% |
| 30 | 90% | 92% | 87% | 92% | 92% |
| 40 | 95% | 93% | 86% | 94% | 94% |
| 50 | 92% | 65% | 63% | 61% | 91% |
| 60 | 40% | 32% | 29% | 33% | 39% |

- **Bands 30 and 40 are too easy.** The content rule makes them fights a bare healer nearly heals in full. Healing power then buys almost nothing, so the ranker's sub-60 healer gear picks at 30 and 40 will tend to tie.
- **Band 30 equals band 20 in damage.** The band-30 last bosses are casters with light melee, and the rule takes melee only.
- **The level-60 profile has its own retune rule** (bare sets under half, mana a live constraint). Applying it per band would need a band multiplier on top of the content figure.
- **Sweep, 2,000 iterations, seed 1, strongest bare set:**

| band | x1.5 | x2 | x3 | x4 |
|---|---|---|---|---|
| 20 | 62% | 47% | 32% | 24% |
| 30 | 94% | 92% | 62% | 47% |
| 40 | 92% | 66% | 43% | 32% |
| 50 | 58% | 43% | 29% | 21% |

To match level 60's bare coverage of about 40%, the multiplier would be about x2 at 20 and 50, about x4 at 30 and about x3 at 40. This is not applied. The ruling asks for content-derived damage, so it is for the owner.

## Before/after

Before: $R/sb-before (origin/main bb619a188, fork 4d3129cae, level-60 profile at every band). After: this lane before the healers merge (S 999f14377, profile plus re-measured rows, same engine), so the change is this lane's alone. Published entries, bare, 10,000 iterations, effective HPS. On the merged state the Chain Heal row adds +0.3% at shaman 50. The large drops at 30 and 40 are the fight getting easier (healers cover 86 to 95%, so effective HPS falls to the incoming damage), not healers getting worse.

| spec | faction | band | before eff. HPS (seed 1) | after eff. HPS (seed 1) | change seed 1 / 50001 | mana lasts before to after (s, seed 1) | overheal after |
|---|---|---|---|---|---|---|---|
| druid-restoration | alliance | 20 bare | 41.4 | 38.6 | -6.76% / -6.76% | 260 to 272 | 1% |
| druid-restoration | alliance | 30 bare | 108.2 | 52.6 | -51.39% / -51.34% | 231 to 414 | 47% |
| druid-restoration | alliance | 40 bare | 173.4 | 128.0 | -26.18% / -26.14% | 215 to 617 | 24% |
| druid-restoration | alliance | 50 bare | 268.2 | 296.4 | +10.51% / +10.39% | 216 to 428 | 3% |
| druid-restoration | horde | 20 bare | 39.3 | 36.6 | -6.87% / -6.87% | 252 to 262 | 1% |
| druid-restoration | horde | 30 bare | 106.3 | 52.5 | -50.61% / -50.56% | 222 to 429 | 46% |
| druid-restoration | horde | 40 bare | 173.1 | 127.9 | -26.11% / -26.07% | 213 to 639 | 24% |
| druid-restoration | horde | 50 bare | 265.7 | 296.1 | +11.44% / +11.36% | 214 to 409 | 3% |
| paladin-holy | alliance | 20 bare | 30.8 | 33.0 | +7.14% / +7.14% | 40 to 124 | 0% |
| paladin-holy | alliance | 30 bare | 81.4 | 52.1 | -36.00% / -36.00% | 29 to 2604 | 0% |
| paladin-holy | alliance | 40 bare | 158.3 | 126.3 | -20.21% / -20.10% | 113 to 389 | 0% |
| paladin-holy | alliance | 50 bare | 216.1 | 203.7 | -5.74% / -5.78% | 188 to 280 | 0% |
| paladin-holy | horde | 20 bare | 29.5 | 32.1 | +8.81% / +8.81% | 43 to 127 | 0% |
| paladin-holy | horde | 30 bare | 78.4 | 52.1 | -33.55% / -33.46% | 28 to 2265 | 0% |
| paladin-holy | horde | 40 bare | 174.0 | 127.4 | -26.78% / -26.82% | 111 to 544 | 0% |
| paladin-holy | horde | 50 bare | 227.3 | 212.8 | -6.38% / -6.38% | 192 to 285 | 0% |
| priest-discipline | alliance | 20 bare | 32.4 | 31.2 | -3.70% / -3.41% | 30 to 175 | 0% |
| priest-discipline | alliance | 30 bare | 85.4 | 52.5 | -38.52% / -38.45% | 36 to 628 | 2% |
| priest-discipline | alliance | 40 bare | 149.5 | 119.4 | -20.13% / -20.13% | 42 to 258 | 8% |
| priest-discipline | alliance | 50 bare | 214.4 | 198.4 | -7.46% / -7.46% | 46 to 159 | 5% |
| priest-discipline | horde | 20 bare | 29.3 | 29.0 | -1.02% / -0.68% | 29 to 164 | 0% |
| priest-discipline | horde | 30 bare | 91.4 | 52.6 | -42.45% / -42.39% | 46 to 841 | 2% |
| priest-discipline | horde | 40 bare | 157.0 | 122.4 | -22.04% / -22.04% | 56 to 276 | 8% |
| priest-discipline | horde | 50 bare | 221.5 | 201.3 | -9.12% / -9.12% | 63 to 167 | 5% |
| priest-holy | alliance | 20 bare | 35.1 | 37.5 | +6.84% / +7.14% | 28 to 222 | 2% |
| priest-holy | alliance | 30 bare | 69.6 | 52.1 | -25.14% / -25.04% | 29 to 888 | 7% |
| priest-holy | alliance | 40 bare | 126.4 | 127.8 | +1.11% / +1.11% | 33 to 282 | 10% |
| priest-holy | alliance | 50 bare | 188.5 | 196.7 | +4.35% / +4.24% | 44 to 179 | 4% |
| priest-holy | horde | 20 bare | 30.1 | 33.7 | +11.96% / +11.96% | 24 to 188 | 1% |
| priest-holy | horde | 30 bare | 64.4 | 52.1 | -19.10% / -19.10% | 26 to 696 | 7% |
| priest-holy | horde | 40 bare | 120.2 | 127.9 | +6.41% / +6.49% | 30 to 283 | 10% |
| priest-holy | horde | 50 bare | 181.9 | 189.8 | +4.34% / +4.23% | 40 to 168 | 4% |
| shaman-restoration | alliance | 20 bare | 54.6 | 51.8 | -5.13% / -5.30% | 49 to 258 | 5% |
| shaman-restoration | alliance | 30 bare | 95.7 | 52.5 | -45.14% / -45.14% | 123 to 3460 | 10% |
| shaman-restoration | alliance | 40 bare | 171.5 | 128.4 | -25.13% / -25.09% | 252 to 667 | 8% |
| shaman-restoration | alliance | 50 bare | 246.1 | 283.6 | +15.24% / +15.29% | 63 to 270 | 2% |
| shaman-restoration | horde | 20 bare | 49.8 | 47.8 | -4.02% / -4.02% | 50 to 208 | 3% |
| shaman-restoration | horde | 30 bare | 95.1 | 52.6 | -44.69% / -44.79% | 120 to 3443 | 10% |
| shaman-restoration | horde | 40 bare | 169.8 | 128.4 | -24.38% / -24.38% | 244 to 676 | 8% |
| shaman-restoration | horde | 50 bare | 245.4 | 283.2 | +15.40% / +15.36% | 62 to 269 | 2% |

Level 60, bare and raid, both factions, all five healers: +0.00% / +0.00% (profile unchanged).

## For other lanes

- **Integrator:**
  - sim/request TestRotationLadder goldens for druid-restoration, paladin-holy, priest-discipline, priest-holy and shaman-restoration move below 60 (band profile and rows). I did not regenerate them.
  - My branch merges wt/r10-healers and lane/r10-healers.
  - Nothing here moves a non-healer spec.
- **Ranker lane (sim/cmd/leveling-bis, not touched):**
  - The heal profile now carries `bands`, and `heal_profile` in each healer BiS file will publish it automatically.
  - `score_heal`'s weights and tuning bounds get the band through `AttachWeights` and `Attach` with no change needed.
  - At bands 30 and 40, sub-60 healer gear will rank on a nearly covered fight (see the coverage question).
  - Per-band talent builds (T1 to T5) need a mechanism. Options: a curated per-band override, or a walk that takes the healing talents first. It belongs to whoever owns `leveling.LadderTalentString`.
- **Ranks lane:** I changed the notes sentence about healer rows in data/curated/rank-choices.json (shared file) and only the healer rows below 60. rotation_ranks.go is untouched.
- **Platform lane:** api/openapi.yaml's role_metrics text ("the curated heal profile's fake raid") is still true. The band is part of the curated profile. No regeneration needed.

## Still open

- Whether a per-band coverage multiplier should sit on top of the content rule (owner; sweep above).
- The dungeon x1.2 damage modifier rests on one client table (ContentTuningXExpected / ExpectedStatMod 489). If the server does not apply it, every band's damage drops by a sixth.
- Boss spell damage is not in the rule (no client or creature_template field gives it). This undercounts the caster last bosses of band 30 (Charlga Razorflank, Bloodmage Thalnos) and Whitemane.
- T1 to T5 (per-band talent builds) wait for a mechanism.
- Tests deselected: none. sim/request TestRotationLadder fails on the five healer goldens by design (integrator regenerates at the pin).

## Follow-up (integrator): boss spells in the rule, and per-band builds

This section supersedes the band figures, rank rows and coverage above.

### F1. The rule counts the last bosses' spells

The old rule counted boss melee only, so band 30, whose last bosses are casters, came out equal to band 20. A last boss's damage per second is now:
- **Melee**, less the share of the fight it spends casting or channelling.
- **Plus** each damaging spell its AI casts while tanked: mean damage per cast at its level, divided by the mean of its repeat timer.

Where each figure comes from:
- **Timers:** the spell list the template names, else its EventAI rows, else its ScriptDevAI script. Melee-range timers only. Onyxia counts her ground phase.
- **Amounts:** cmangos `spell_template`, using the 1.12 base points, dice and points per level. A periodic aura counts its ticks over its duration. A weapon-damage effect counts as one melee hit plus its bonus. Triggered spells count as they trigger.
- **Cast and channel times:** client SpellCastTimes and SpellDuration.
- **Why not the client's amounts:** the client stores these NPC spells only in a scaled form whose formula it does not state, and Wowhead's tooltips print no number for them.
- **Left out:** heals, buffs, summons, crowd control without damage, and one-off health-threshold casts.
- **Counting:** each spell counts once per cast, on one target.
- **Where it is stated:** still in one place, `bands.rule`. Each spell row in heal-profile.json names its source and derivation. Loading refuses a row with no source or no positive interval. Tests were written first: `healprofile_bands_test.go`.

Sources: classic-db `creature_spell_list` 454301 and `creature_ai_scripts` (commit ec4f596), and mangos-classic ScriptDevAI `boss_herod.cpp`, `boss_mograine_and_whitemane.cpp`, `boss_thermaplugg.cpp` and `boss_onyxia.cpp` (commit ff13624).

| boss | spells counted (damage per cast, interval) | damage per second, melee + spells |
|---|---|---|
| Edwin VanCleef (20) | none (Thrash passive, allies summon) | 46.8 |
| Mutanus the Devourer (20) | Thundercrack 104 / 13 s | 55.2 |
| Charlga Razorflank (30) | Chain Bolt 162 / 3.5 s (2.5 s cast) | 59.4 |
| Mekgineer Thermaplugg (30) | Knock Away 170 / 18.5 s (walking bombs are summons, left out) | 79.3 |
| Bloodmage Thalnos (30) | Flame Spike 214 / 12 s, Fire Nova 159 / 10.5 s, Flame Shock 190 / 38 s (Shadow Bolt is his ranged-mode main spell, left out) | 73.3 |
| Herod (40) | Cleave 319 / 12.5 s, Whirlwind 2,283 / 20 s (10 s channel) | 195.4 |
| Amnennar the Coldbringer (40) | Frost Nova 120 / 15 s, Frostbolt 196 / 3.9 s (3 s cast) | 88.7 |
| High Inquisitor Whitemane (40) | Holy Smite 150 / 5.75 s (melee mode, 2.5 s cast) | 67.1 |
| Princess Theradras (50) | Dust Field 1,200 / 22.5 s, Boulder 450 / 9 s | 261.6 |
| Shade of Eranikus (50) | War Stomp 402 / 20 s, Acid Breath 932 / 25 s | 366.9 |
| Onyxia (60) | Flame Breath 3,500 / 15 s, Cleave 2,250 / 7.5 s, Tail Sweep 800 / 17.5 s, Wing Buffet 750 / 22.5 s, Knock Away 1,965 / 32.5 s | 1,468.1 |

New band figures at the tuned scale (60 unchanged):

| band | median boss damage per second x mod | tank health / hit | members / pulse | incoming per second |
|---|---|---|---|---|
| 20 | 51.0 x 1.2 | 1,836 / 89 | 966 / 23 | 41 |
| 30 | 73.3 x 1.2 | 3,204 / 128 | 1,687 / 33 | 59 |
| 40 | 88.7 x 1.2 | 4,938 / 154 | 2,599 / 40 | 71 |
| 50 | 314.3 x 1.2 | 7,036 / 547 | 3,703 / 143 | 252 |
| 60 | 1,468.1 x 1.0 | 9,500 / 2,128 | 5,000 / 555 | 978 |

- **Band 30 is no longer equal to band 20.**
- **Bands 20 and 40 get lighter.** Onyxia's own spells raise the reference by half, and the dungeon bosses' spells add less than that.
- **Open item, pulls of several mobs.** A dungeon pull is often several mobs, and the rule counts one boss. A source for how many mobs a band's pulls hold would settle a pull factor. None is used, so no multiplier is added.

Commit: S 64a1c9717.

### F2. Per-band healer builds

**Mechanism** (S a09dd3b99, tests first):
- `data/curated/ladder-first.json` is embedded in sim/leveling (`make ladder-first-sync`).
- It names, per spec, ranks of the spec's own tree that `LadderTalentString` takes first. It takes one rank a step, each once the build can legally hold it, and never past the guide's rank for that talent.
- The ladder only adds points, and level 60 is the guide's build.
- A spec with no row walks byte for byte as before. Only the three healers below have rows, so DPS specs are unchanged.
- The behaviour is passed as `leveling.ForSpec`. Every production caller passes it: leveling-bis, addon-builds, dps-per-point, rotation-search, talent-search and the ladder test.
- Tests:
  - `ladder_first_test.go`: the walk unchanged without a row; reserved ranks taken one at a time; the guide's targets respected.
  - `ladder_first_curated_test.go`: every curated row is legal at every level from 10 to 60, never takes back a point, ends on the guide's level-60 build, and names real talents of the spec's own tree.

**Search:**
- Baseline: the current guide's plain walk. The published entries predate the healers lane's guide changes for the druid, Discipline and Holy.
- Hill climb over every legal one-point move that stays inside the guide's ranks, at 2,000 iterations, while a move gains at least 1%. Then 10,000 iterations, seeds 1 and 50001, both factions, every band.
- Adoption rule:
  - gain of at least 1% at both seeds
  - no talent dropped that a level-60 probe shows the engine ignores (removing it leaves the raid heal run exactly the same)
  - no talent dropped that grants a spell the rotation casts
  - no band loses

| spec | first ranks | band build change | gain, seed 1 / 50001 (alliance; horde) |
|---|---|---|---|
| Restoration druid | Heart of the Wild 5, Reflection 2, Bloom 1, Nature's Swiftness 1 | 20: Nature's Focus 5 to 3, Naturalist 1 to 0, Reflection 0 to 2, Bloom 0 to 1. 30: Bloom 1 and Nature's Swiftness kept. 40 to 60: unchanged. | 20: +23.76% / +23.76%; horde +24.58% / +24.58%. 30: +2.08% / +2.08%; horde +2.47% / +2.28%. 40: +0.00%. |
| Discipline | Meditation 1 | 20: Mental Agility 1 to 0, Meditation 0 to 1. Others unchanged. | +1.95% / +1.95%; horde +1.68% / +1.68% |
| Holy | Prayer of Mending 1 | 40: Spiritual Healing's last point to Prayer of Mending. The plain walk left the rotation's own capstone out at 40. Others unchanged. | +6.89% / +7.05%; horde +7.84% / +7.84% |

Not adopted:
- **Holy paladin 20, Healing Light to Illumination:** +0.91% / +1.21%, horde +0.31%, so it misses 1% at seed 1.
- **Restoration druid 50, Improved Tranquility to Genesis:** +1.8% at 2,000 iterations, but it moves a point out of the own tree, which a reservation cannot express. It is small.
- **No move at 1% or more:** shaman at any band; Discipline 30 to 50; Holy 20, 30 and 50; druid 40.
- **The earlier T1 to T5 moves** (Discipline Wand Specialization, shaman Water Shield, Holy 50 Spiritual Guidance and the rest) are void. They were measured from the stale published builds on the melee-only profile, and Spiritual Guidance 5 would go past the guide's level-60 rank.

The guides and planner state no band build (each healer guide's Leveling section names no talents), so no text changed. `data/builds/1.60.1.70707/addon-builds.json` follows for the three specs, as its committed-copy test requires.

### F3. Rank rows on the final bands and builds

Same method as before, now run with each band's final build (`-talents`), plus two checks: no row may turn a fight that runs dry into one with mana unspent, and no faction may lose. Commit: S 90d16b4f0.

| spec | levels | row | gain vs default, seed 1 / 50001 (horde) |
|---|---|---|---|
| Restoration druid | 50-59 | Rejuvenation 8 (added; an uprank) | +1.36% / +1.40% (+1.4%) |
| Discipline | 50-59 | Shield 8, Prayer of Healing 1, Binding Heal 1 (kept) | set +2.5% (+2.5%) |
| Holy | 40-49 | Renew 5 (added; an uprank) | +1.17% / +1.17% (+0.3% / +0.4%) |
| Holy | 50-59 | Flash Heal 6, Prayer of Healing 1 (kept) | set +10.4% (+10.6%) |

Dropped:
- **Rejuvenation 7 at 40:** loses 0.6%.
- **Holy Light 3 at 50:** gains 2% but turns a fight that runs dry at 274 s into one that ends with mana unspent.
- **Discipline Shield 5 at 30:** inside the printing resolution.
- **The Discipline band-40 set** (Shield 1, Prayer of Healing 1, Penance 1): loses 2%.
- **Heal 4 at 50.**
- **Holy Prayer of Healing 1 at 40, Lightwell 1 at 50.**
- **Chain Heal 1 at 50:** loses 0.4%.
- **Discipline Shield 2 at 20:** +1.3% alliance but -3.6% horde, so not adopted.

### Bare coverage per band (final state, alliance, seed 1, effective HPS / incoming)

| band | druid | paladin | Discipline | Holy | shaman |
|---|---|---|---|---|---|
| 20 | 92% | 81% | 78% | 85% | 94% |
| 30 | 92% | 92% | 93% | 92% | 92% |
| 40 | 90% | 89% | 93% | 98% | 92% |
| 50 | 95% | 72% | 79% | 87% | 98% |
| 60 | 40% | 32% | 29% | 33% | 39% |

- With the bosses' spells counted, every sub-60 band is covered at 72 to 98% by bare healers, against 29 to 40% at 60.
- **Owner question, still open:** level 60's own retune rule would add a per-band multiplier, roughly x2 to x4 by the earlier sweep. It is not applied.
- **Still open:** pull size (see F1).
