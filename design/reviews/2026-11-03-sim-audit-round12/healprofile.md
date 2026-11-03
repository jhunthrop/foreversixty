# Sim audit round 12: the heal profile's tank-and-group split

Lane: wt/r12-healprofile (site, from c6bdc855c). No engine change; lane/r12-healprofile (fork) has no commits of its own and sits at 2606bc05e.

Scope (owner ruling, round 12): derive the sub-60 heal profile's split between the tank and the group from encounter data and logs, with a stated rule; re-measure every healer at every band; re-check the sub-60 healer rank rows and band builds; recommend on line of sight in the pull rule.

Every run went one at a time under `nice -n 15`, GOMAXPROCS=9. No BiS ranking ran, no goldens or generated data were regenerated.

## Where the split was set, and what it did

- `data/curated/heal-profile.json` holds one tank hit (1,150 every 3.3 s) and one raid pulse (300 to three of five members every 5 s), both times the tuning scale 1.85. That is 348.5 a second on the tank against 180 on the group: **0.659 on the tank**.
- Below 60, `figuresFor` (sim/request/healprofile_bands.go) scaled the hit and the pulse by the same content ratio, so every band kept 0.659. The rule said so: "the tank-to-raid ratio stays what it is at 60".
- So every sub-60 healer number, rank row and band build was measured on a fight whose tank-versus-group balance was level 60's assumption, not any band's.

## The rule (bands.rule and bands.split_rule; applied only by healprofile_bands.go)

**For a Classic player:** below 60, the share of a band's damage that lands on the tank is the median, over the logged beta trash fights in that band's dungeons on foreverlogs.gg, of the share of the group's damage taken that the group's tank took. The tank is the warrior, paladin or druid in the group who took the most melee over the whole run, so a fight where a loose mob, a pet or a caster ate the hits still counts against the tank's share. Only player damage counts (pets are not healed by the profile). A band whose dungeons hold fewer than 30 logged trash fights uses every logged beta dungeon's trash fights instead. Trash fights are the ones counted because every sub-60 band fights a trash pull (round 11's pull multiplier is above 1 at all four). The band's total incoming damage is unchanged: the tank hit carries the share, the pulse the rest. Level 60 keeps the profile's own 0.659 until level-60 raid logs can be read.

A player can check it: open any foreverlogs.gg beta Deadmines or Wailing Caverns run, take a trash pull, and divide the tank's damage taken by the group's.

| band | logged trash fights in its dungeons | share fights | median tank share (IQR) | was | tank hit / pulse, was to now |
|---|---|---|---|---|---|
| 20 | 218 (Deadmines, Wailing Caverns) | 218 | **0.679** (0.44 to 0.89) | 0.659 | 143 / 37 to 147 / 35 |
| 30 | 182 (Razorfen Kraul, Gnomeregan, SM Graveyard) | 182 | **0.687** (0.50 to 0.89) | 0.659 | 263 / 69 to 274 / 63 |
| 40 | 6 (SM Armory) | 990 (every beta dungeon) | **0.677** (0.43 to 0.88) | 0.659 | 470 / 123 to 482 / 116 |
| 50 | 0 | 990 (every beta dungeon) | **0.677** | 0.659 | 1,010 / 263 to 1,037 / 250 |
| 60 | no public level-60 logs readable | - | 0.659 (the profile's own) | 0.659 | 2,128 / 555 unchanged |

Incoming damage a second is the same before and after at every band (65.5, 121, 216, 464, 978).

**Sensitivity** (median per fight, same tank rule):

| fights counted | band 20 | band 30 | every beta dungeon |
|---|---|---|---|
| trash only (the rule) | 0.679 (218) | 0.687 (182) | 0.677 (990) |
| boss kills only | 0.745 (179) | 0.557 (109) | 0.634 (849) |
| bosses and trash | 0.710 (397) | 0.656 (291) | 0.656 (1,839) |
| trash, pooled by damage, not per fight | 0.636 | 0.637 | 0.615 |
| trash, tank = the fight's top melee taker (tank-class fights only) | 0.706 | - | 0.734 |

Every reading sits between 0.56 and 0.75. The fixed 0.659 was close; the logged trash split moves 2 to 3 points of damage onto the tank.

**Why encounter data does not set the split:**

| source | what it says the tank takes |
|---|---|
| round 11's pulls (spawns and AI, every mob held by the tank) | 0.99 / 0.985 / 0.958 / 0.958 at bands 20 / 30 / 40 / 50 |
| last bosses from creature_template and AI (spells at a random target, area spells, chains counted as group) | VanCleef 1.00, Mutanus 0.86, Charlga 0.22, Thermaplugg 1.00, Thalnos 0.79, Herod 0.42, Amnennar 0.91, Whitemane 1.00, Theradras 0.61, Eranikus 0.95, Onyxia (ground phase) 0.97 |
| beta logs, same bosses (median per kill) | VanCleef 0.88 (14), Mutanus 0.72 (10), Charlga 0.40 (14), Thalnos 0.44 (15), Thermaplugg 0.67 (1) |
| beta logs, trash | 0.68 |

Encounter data assumes every mob stays on the tank all fight. Real groups lose mobs, pull with casters, and stand in area spells, so it overstates the tank's share of trash by about 30 points. Logs measure that directly, so they set the rule; encounter data is the cross-check.

**Sources (checked 2026-11-03):**
- foreverlogs.gg API: `reports/{id}` and `reports/{id}/encounters/{id}/timeline?event_types=damage` for the 119 public beta reports cached on 2026-10-31 (levels about 13 to 30). Every boss kill and a seeded sample of up to 12 trash fights of 10 s or more a report: 2,113 fights fetched, 1,863 five-man fights with a tank-class player kept. The API's `reports` list now asks for a token, so no newer report could be listed; report ids past 2811 return "Report not found".
- `statistics?phase=1&metric=avg_dtps&role=tank|dps` (foreverlogs, per dungeon): role averages only, no group split; used as a sanity read (Wailing Caverns tanks about 20 a second against dps about 4 to 7).
- `GET https://api.foreversixty.gg/v1/reports/recent` (2026-11-03): three reports, open-world trash only (no bosses, no five-man group); not usable.
- Level 60: Classic Era and SoD warcraftlogs need an API key (v1 returns 401, the site returns 403 to scripts), no browser extension was connected, and no WCL secret exists in the project's secret store. foreverlogs.gg holds no Molten Core or Onyxia fight (its Molten Core statistics are empty). So no level-60 split could be read.
- Encounter data: cmangos classic-db at ec4f596 and mangos-classic ScriptDevAI at ff13624 (round 10 and 11's rows), with each spell classed tank or group by its target type.
- Client 1.60.1.70753 (wago.tools, 2026-11-03): LFGDungeons states each Classic dungeon as 1 tank, 1 healer, 3 damage (the profile's five), and the dungeons' ContentTuningXExpected rows (29, ExpectedStatMod 489) and ExpectedStat PlayerHealth (503 / 878 / 1,353 / 1,928 / 2,603) are unchanged from 70707.
- Tools, in this folder: `healprofile/fetchdt.py` (fetch), `healprofile/split.py` (per-fight split, run tank, pull sizes), `healprofile/apply.py` (writes the rows), `healprofile/split-output.txt` (every figure above), `healprofile/figures_print_test.go.txt` (the scratch test apply.py uses to print each band's figures).

## Findings

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| H1 | **Every sub-60 band split its damage 0.659 tank / 0.341 group, level 60's assumed ratio.** Logged beta trash puts 0.68 on the tank at bands 20 and 30. Each band row now states its logged evidence (fights, dungeons, median, quartiles, source) and its share; loading refuses a share that is not the stated median, a split on fewer than 30 fights, a pooled split where the content has enough, and a reference band with logs or a share other than the profile's own. Tests first (they failed to build, then passed): `TestHealBandSplitMovesDamageNotTheTotal`, `TestHealBandReferenceKeepsTheProfilesSplit`, `TestCuratedHealBandsStateTheirSplit`, `TestHealProfileRefusesASplitThatIsNotItsLogs`; the band figures in the older band tests follow. | Rule and table above; sim/request/healprofile_bands.go `figuresFor`, `validateSplit` | CONFIRMED-BUG (owner ruling) | R10 S e9bd57485 (the "ratio stays as at 60" rule) | S c600673ed |
| H2 | **The BiS page's healer disclosure said nothing of the split.** A sub-60 band whose row states one now shows "Tank and group: The tank takes 68% of the damage, the median of 218 logged beta fights in these dungeons" (pooled bands say "logged beta dungeon fights (none at this band yet)"). Files published before this show no row. Vitest written first. | web/src/lib/bis/heal-view.ts `splitRows`, copy.ts, types.ts | CONFIRMED (page follows H1) | R10 S 8ad647829 | S c600673ed |
| H3 | **Client 70753 re-checked under the band rule.** ExpectedStat PlayerHealth, the 29 dungeon ContentTuningXExpected rows (mod 489, x1.2 melee and spell damage) and the raids' tuning 4509 (no modifier) are as 70707. LFGDungeons adds Forever's City of Dalaran (4512), Ruins of Lordaeron (5256), Hall of Thanes (7000) and Excavation Site: Wetlands (6874), all under mod 489; none has a creature_template boss, so none is in a band's content (as round 10 said). | wago.tools ExpectedStat, ExpectedStatMod, ContentTuningXExpected, LFGDungeons at 1.60.1.70753 | Aligned | - | - |
| H4 | **Discipline 20-29 and 30-39: the Power Word: Shield 2 and 5 rows no longer gain.** Renew 1 alone beats Renew 1 plus Shield 2 at 20 (alliance +0.77%, horde +1.1%; Shield 2 alone is -2.6% horde). At 30 Shield 5 adds 0.0% / +0.23% inside the set and loses alone (-0.25% / -0.36%). Both rows dropped. The same holds on the old split, so the split did not cause it. | 10,000 iterations, seeds 1 and 50001, both factions (table below) | CONFIRMED-BUG (row against its own bar) | landing b3690c448, f726a437b (Discipline's list changed after round 11 measured the rows) | S ac97e6245 |
| H5 | **Discipline 40-49 and 50-59 missed Binding Heal 1 and Shield 6.** At 40, Renew 1, Binding Heal 1, Shield 6 and Prayer of Healing 1 beat the old set (Prayer of Healing 1, Renew 1, Penance 1) by +0.69% / +0.61%; taking any one row out loses 0.5 to 0.9%, and adding Penance 1 back loses. At 50, Binding Heal 1 adds +0.55% / +0.37%. Both bands still run dry (about 160 s and 90 s). | same | CONFIRMED-BUG | landing f726a437b (party Binding Heals regrouped) | S ac97e6245 |
| H6 | **Holy paladin 40-49 missed a Flash of Light pair.** Flash of Light rank 1 on the group (19941) with rank 3 on the tank (19750), beside Holy Light 1: +2.77% / +3.07% over Holy Light 1 alone, both seeds equal. Either Flash of Light row alone, or the tank's rank 3 without the group's rank 1, loses (rank 3 alone -27%). The fight runs dry at 150 s instead of 275 s; both are dry, so the row's mana rule holds. The gain is there on the old split too (+3.0%). | same | CONFIRMED-BUG | old (round 11's greedy takes one row a step; this gain needs two together) | S ac97e6245 |
| H7 | **Restoration shaman band 20 left out Water Shield, which the rotation casts.** The ladder spent its eleventh point on Tidal Focus 4. Reserving Water Shield 1 after Mindfulness 2 gives 5320000010 at 20; bands 30 to 60 walk as before, and the reservation is sound at every level 10 to 60. +12.4% / +12.4% alliance, +15.9% / +15.9% horde (old split: +9.5%, +15.1%). Nothing modeled or cast is dropped. Test first: `TestRestorationShamanHoldsWaterShieldAtTwenty`. Shaman 20 rank rows re-checked on the new build: Healing Wave 1 still turns a dry fight into one with mana unspent (horde 246 s to 511 s), so no row. | data/curated/ladder-first.json; talent climb at 2,000 then 10,000 iterations | CONFIRMED-BUG (adoption rule met) | old (present before this round; round 10's T4 saw a smaller form, round 11's climb did not; blame is not decisive) | S d6ec64e69 |
| H8 | **Line of sight in the pull rule (round 11's open question). Recommendation: leave it out.** Line of sight can only split a pull into smaller ones. Logged beta trash fights are already bigger than the 10-yard rule's pulls: at band 20 the median logged trash fight holds 4 mobs (weighted by mobs: median 8, 75th percentile 15) against the rule's 2 and 3; band 30 holds 5 (8 and 13) against 2 and 3. Beta runs are speed runs by overlevelled groups, and a logged fight can join chain pulls, so these are upper bounds, but no reading points toward smaller pulls. | split-output.txt "logged trash pull sizes" | OPEN (recommend; not implemented) | - | - |

Nothing in this lane moves a non-healer spec. Effective DPS is equal to the last digit before and after for mage-fire 20, shaman-enhancement 40, shaman-elemental 30 and priest-shadow 50 (bare, horde, 3,000 iterations).

## Sub-60 healer rank rows (rank-choices.json, S ac97e6245)

**Method:** each band's current ladder build (`-talents`), the new profile. Greedy from the default rule over every learned rank of the rotation's ranked heals, 3,000 iterations. Then default, current, the greedy set and its subsets at 10,000 iterations, seeds 1 and 50001, both factions. A row is adopted where it gains beyond error at both seeds (the round-12 ruling; 0.3% is used as the bar, three times the run-to-run spread at this length). A row never turns a dry fight into one that ends with mana unspent.

| spec | levels | was | now | vs default, alliance s1 / s50001; horde s1 / s50001 | mana lasts (s) |
|---|---|---|---|---|---|
| priest-discipline | 20-29 | Shield 2, Renew 1 | Renew 1 | +7.71 / +7.71; +3.79 / +3.79% (old set +6.89; +2.62%) | 172 |
| priest-discipline | 30-39 | Shield 5, Renew 1 | Renew 1 | +5.01 / +5.01; +4.55 / +4.55% (old set +5.01; +4.78%, inside error) | 164 |
| priest-discipline | 40-49 | Prayer of Healing 1, Renew 1, Penance 1 | Renew 1, Binding Heal 1, Shield 6, Prayer of Healing 1 | +7.25 / +7.25; +7.21 / +7.21% (old set +6.52 / +6.45; +6.56 / +6.49%) | 162 |
| priest-discipline | 50-59 | Prayer of Healing 1, Renew 1 | + Binding Heal 1 | +5.69 / +5.64; +8.38 / +8.33% (old +5.11 / +5.06; +7.98 / +7.89%) | 90 |
| paladin-holy | 40-49 | Holy Light 1 | + Flash of Light 1 (group), Flash of Light 3 (tank) | +3.63 / +3.63; +3.86 / +3.80% (old +0.83; +0.77 / +0.71%) | 150 |

Rows that hold on the new profile (current set vs default, 10,000 iterations, alliance; horde within 1.5 points): druid 50-59 Rejuvenation 1 (+10.5%), paladin 20-29 Holy Light 1 (+15.5%), paladin 30-39 Holy Light 1 and Flash of Light 1 (+8.1%), Holy 20-29 and 30-39 Renew 1 (+2.6%, +8.0%), Holy 40-49 set (+13.7%), Holy 50-59 set (+18.2%), shaman 40-49 and 50-59 Healing Wave 2 (+34.5%, +27.1%).

Not adopted (the greedy's pick turns a dry fight into one with mana unspent): druid Rejuvenation 2 at 20 (943 s), 1 at 30 (540 s), 2 at 40 (540 s); paladin Holy Light 2 at 50 (1,363 s); shaman Healing Wave 1 at 20 (309 s) and 2 at 30 (630 s). These are the same calls round 11 made.

## Sub-60 healer band builds

Hill climb from each band's ladder build (one-point moves inside the guide's ranks, 2,000 iterations, 1% a step), unmodeled-talent probes taken fresh on this engine; every result confirmed at 10,000 iterations, seeds 1 and 50001, both factions, on both the old and the new split.

**The split moves no build across the 1% bar:** every move the climb found already gained beyond 1% on the old split, and the split shifts each by at most 3 points. One is adopted (H7). The rest:

| spec, band | climb's build | gain, new split (old split) | why not adopted |
|---|---|---|---|
| druid 30 | Nature's Focus 5 to 2, Improved Rejuvenation 0 to 3 | +4.2 to +4.5% (+3.3 to +4.0%) | band 20 holds Nature's Focus 2, so 30 cannot lose a point a reservation never gave; unsound (round 11's finding) |
| druid 40 | Nature's Focus and Heart of the Wild to Living Spirit 2 | +3.4% (+3.1%) | takes back Nature's Focus and Heart of the Wild points; unsound (round 11) |
| druid 50 | Improved Tranquility and Heart of the Wild to Genesis (Balance) | +3.3% (+3.2%) | moves points into another tree; a reservation cannot express it (round 11) |
| Holy 30 | Twilight Focus 3 to 0; Holy Binding, Litany of Light, Improved Healing +1 each | +14.0 to +16.1% (+13.6 to +15.7%) | band 20 holds Twilight Focus 3; unsound. Round 11 named this move as dropping Inspiration; the probe now finds Inspiration unmodeled but this move keeps it |
| Holy 40 | Twilight Focus and Divine Fury to Spiritual Guidance 4 | +4.8 to +5.1% (+4.7 to +5.0%) | a Spiritual Guidance 4 reservation leaves every band unchanged (the walk already takes Guidance's guide rank later); needs points taken back at 40 |
| shaman 30 | Improved Healing Wave, Healing Focus, Tidal Mastery to Healing Way, Restorative Totems, Mana Tide Totem | +18.5 to +18.8% (+16.7 to +17.5%) | takes back Improved Healing Wave points band 20 holds; unsound |
| shaman 40 | Healing Focus 3 to 1, Healing Way 0 to 2 | +3.3 to +3.5% (+2.8 to +3.0%) | the reservation that reaches it drops Mana Tide Totem at 40, which the rotation casts (round 11's finding, re-tested) |

No move at all: paladin 20 to 50; Discipline 20 to 50; Holy 20 and 50; druid 20; shaman 50.

## Before/after

Before: $R/sb-before from $R/site-before (c6bdc855c, the 0.659 split). After: this branch (S d6ec64e69): logged split, new rank rows, shaman band-20 build. Bare, current ladder talents, 10,000 iterations. Effective HPS at seed 1, relative change at seeds 1 / 50001, coverage (effective HPS over incoming) and mana lasts (s) at seed 1. Level 60 is byte-identical for all five healers, both factions, both seeds.

| spec | faction | band | eff. HPS | change s1 / s50001 | coverage | mana lasts |
|---|---|---|---|---|---|---|
| druid-restoration | alliance | 20 | 59.0 to 59.8 | +1.36% / +1.36% | 90% to 91% | 286 to 293 |
| druid-restoration | alliance | 30 | 97.9 to 97.8 | -0.10% / -0.10% | 81% to 81% | 282 to 277 |
| druid-restoration | alliance | 40 | 192.5 to 191.3 | -0.62% / -0.67% | 89% to 89% | 290 to 288 |
| druid-restoration | alliance | 50 | 332.3 to 329.3 | -0.90% / -0.93% | 72% to 71% | 266 to 264 |
| druid-restoration | horde | 20 | 57.1 to 57.8 | +1.23% / +1.05% | 87% to 88% | 283 to 284 |
| druid-restoration | horde | 30 | 96.0 to 95.7 | -0.31% / -0.31% | 79% to 79% | 277 to 272 |
| druid-restoration | horde | 40 | 192.7 to 191.3 | -0.73% / -0.83% | 89% to 89% | 291 to 288 |
| druid-restoration | horde | 50 | 327.0 to 324.1 | -0.89% / -0.95% | 70% to 70% | 267 to 262 |
| paladin-holy | alliance | 20 | 38.0 to 38.1 | +0.26% / +0.26% | 58% to 58% | 103 to 103 |
| paladin-holy | alliance | 30 | 108.0 to 108.6 | +0.56% / +0.46% | 89% to 90% | 259 to 263 |
| paladin-holy | alliance | 40 | 168.1 to 174.3 | +3.69% / +3.69% | 78% to 81% | 272 to 150 |
| paladin-holy | alliance | 50 | 218.5 to 218.5 | +0.00% / -0.09% | 47% to 47% | 135 to 128 |
| paladin-holy | horde | 20 | 37.8 to 37.9 | +0.26% / +0.00% | 58% to 58% | 110 to 109 |
| paladin-holy | horde | 30 | 104.5 to 105.2 | +0.67% / +0.67% | 86% to 87% | 231 to 234 |
| paladin-holy | horde | 40 | 168.2 to 174.8 | +3.92% / +3.86% | 78% to 81% | 264 to 147 |
| paladin-holy | horde | 50 | 214.2 to 214.1 | -0.05% / -0.19% | 46% to 46% | 134 to 126 |
| priest-discipline | alliance | 20 | 39.1 to 39.1 | +0.00% / +0.00% | 60% to 60% | 179 to 172 |
| priest-discipline | alliance | 30 | 84.4 to 83.8 | -0.71% / -0.71% | 70% to 69% | 169 to 164 |
| priest-discipline | alliance | 40 | 145.1 to 146.4 | +0.90% / +0.83% | 67% to 68% | 164 to 162 |
| priest-discipline | alliance | 50 | 218.5 to 219.3 | +0.37% / +0.37% | 47% to 47% | 92 to 90 |
| priest-discipline | horde | 20 | 35.4 to 35.6 | +0.56% / +0.28% | 54% to 54% | 155 to 152 |
| priest-discipline | horde | 30 | 88.4 to 87.4 | -1.13% / -1.13% | 73% to 72% | 174 to 163 |
| priest-discipline | horde | 40 | 147.3 to 148.7 | +0.95% / +0.95% | 68% to 69% | 170 to 166 |
| priest-discipline | horde | 50 | 218.9 to 219.9 | +0.46% / +0.41% | 47% to 47% | 99 to 97 |
| priest-holy | alliance | 20 | 39.4 to 39.2 | -0.51% / -0.51% | 60% to 60% | 182 to 178 |
| priest-holy | alliance | 30 | 63.9 to 63.5 | -0.63% / -0.47% | 53% to 53% | 143 to 133 |
| priest-holy | alliance | 40 | 139.6 to 140.7 | +0.79% / +0.79% | 65% to 65% | 173 to 174 |
| priest-holy | alliance | 50 | 198.0 to 197.4 | -0.30% / -0.40% | 43% to 43% | 102 to 100 |
| priest-holy | horde | 20 | 34.6 to 34.6 | +0.00% / +0.00% | 53% to 53% | 155 to 155 |
| priest-holy | horde | 30 | 60.5 to 59.8 | -1.16% / -0.99% | 50% to 50% | 124 to 118 |
| priest-holy | horde | 40 | 135.4 to 136.5 | +0.81% / +0.89% | 63% to 63% | 163 to 162 |
| priest-holy | horde | 50 | 191.8 to 190.9 | -0.47% / -0.52% | 41% to 41% | 97 to 95 |
| shaman-restoration | alliance | 20 | 54.8 to 60.7 | +10.77% / +10.77% | 84% to 93% | 211 to 358 |
| shaman-restoration | alliance | 30 | 89.7 to 88.6 | -1.23% / -1.45% | 74% to 73% | 208 to 184 |
| shaman-restoration | alliance | 40 | 194.4 to 194.9 | +0.26% / +0.21% | 90% to 90% | 283 to 281 |
| shaman-restoration | alliance | 50 | 329.4 to 323.6 | -1.76% / -1.70% | 71% to 70% | 211 to 190 |
| shaman-restoration | horde | 20 | 49.6 to 56.9 | +14.72% / +14.72% | 76% to 87% | 174 to 246 |
| shaman-restoration | horde | 30 | 89.0 to 87.9 | -1.24% / -1.35% | 73% to 73% | 202 to 179 |
| shaman-restoration | horde | 40 | 193.5 to 193.9 | +0.21% / +0.21% | 90% to 90% | 282 to 280 |
| shaman-restoration | horde | 50 | 328.5 to 322.8 | -1.74% / -1.70% | 71% to 70% | 209 to 188 |

The split alone moves every cell by -1.8% to +1.4% (heal-over-time and group-heal specs lose a little, the tank-heal-heavy paladin gains a little). The larger moves are H4 to H7.

## Files to regenerate

- `data/builds/1.60.1.70753/addon-builds.json` (`go run ./sim/cmd/addon-builds -repo-root .`): the Restoration shaman band-20 build moved (H7). `TestCommittedAddonBuildsFollowTheGuides` fails until it is regenerated.
- sim/request `TestRotationLadder` goldens for the five healers below 60 (split, rank rows, shaman band 20). Not regenerated.
- Each healer BiS file republishes `heal_profile` (with the band rows' new `tank_share` and `split`) at the next bis run; nothing to do by hand.

## Measure at refresh

- Every sub-60 healer pick, all five healers, all four bands (the tank hit and the pulse both moved).
- Restoration shaman 20 in particular: Water Shield's mana return lifts it to 93% / 87% coverage, so its healing-power picks may tie more.
- Holy paladin 40: the fight now runs dry at 150 s, so mana picks may matter more.
- Round 11's carried item: Restoration druid 20's shoulder alternative.

## For other lanes

- **Integrator:** regenerate addon-builds.json and the five healer ladder goldens (above). My fork branch has no commits. Nothing here moves a non-healer spec.
- **Ranks lane:** I changed only healer rows below 60 in data/curated/rank-choices.json and added one clause to its notes (the round-12 report). sim/rankchoices follows (make rank-choices-sync).
- **UX / wow-player:** the BiS page's new "Tank and group" disclosure row (S c600673ed) wants a screenshot review when BiS next publishes. H4 to H7 are written for a Classic player's read; Water Shield at 20 and Flash of Light rank 3 on the tank at 40 are the two worth a wow-player look.
- **Owner:** a warcraftlogs API key (v2 client) would let the level-60 split come from Era Molten Core and Onyxia logs by the same rule.

## Still open

- **Level 60's split** stays 0.659 (assumed). It needs readable level-60 raid logs: Era or SoD warcraftlogs for the unchanged Molten Core and Onyxia, or Forever's own after 9 December. No key or browser was available here. Applying the beta five-man figure (0.68) to a 40-man raid would be a guess, and the retune rule's 1.85 was tuned on 0.659, so it is left alone.
- **The beta logs are levels 13 to 30**, mostly speed runs. Bands 40 and 50 borrow every beta dungeon's split. Re-derive when level-40 and 50 dungeon logs exist (rerun split.py and apply.py after fetchdt.py).
- **Trash is sampled** (up to 12 fights a report, seeded). A full fetch would take about three hours at the API's rate limit; the medians moved by under 0.02 between the first 700 fights and all 2,113.
- **Line of sight:** recommended out (H8). Logged pulls are larger than the rule's, which points the other way; that gap is for the owner if pull size is revisited.
- **Unexpressible band moves** (table above): druid 30, 40, 50; Holy 30, 40; shaman 30, 40. They need a reservation that can drop a top-row point, or guide changes. Holy 30 (+14%) and shaman 30 (+18%) are the large ones.
