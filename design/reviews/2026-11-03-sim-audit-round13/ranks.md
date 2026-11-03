# Sim-accuracy round 13: ranks lane

Scope: the per-band rank-choice mechanism (`sim/request/rank_choices.go`, `rotation_ranks.go`, `sim/rankchoices`), owner ruling 3 (rank rows are keyed by race), and every non-healer row of `data/curated/rank-choices.json`. The three bare-60 healer rows round 12 keyed by faction were re-judged too, because the faction key no longer exists; the other healer rows were not touched.

How it was measured:
- Site `wt/r13-ranks` from int/r13 (93c5d2510), engine int/r13 (076a5a517). This lane made no engine change. Client 1.60.1.70753.
- spec-breakdown on each band's published gear and talents for the race's faction, 10,000 iterations, seeds 1 and 50001.
- The new `-race` flag puts each race on that gear. A scratch override, not committed, swaps the table (`RC_FILE`, as in round 12).
- Scratch scripts and raw output are in scratchpad `r13/ranks/` (`tools/sweep.py`, `judge.py`, `explore.py`, `confirm.py`, `ba.py`; `out/sweep.jsonl`, `out/judge.txt`, `out/explore-*.jsonl`).

The error bar is as in rounds 11 and 12. A DPS row passes when its gain beats twice the combined standard error at each seed. Healer and tank fights print no error, so their rows use a 0.2% bar with both seeds agreeing.

**How a row is keyed now.** Every row was measured on every race that can play its class (combos.json):
- with the row removed (the default rule);
- at the row's rank;
- at one rank on either side.

A race takes the row's rank unless one of two things holds at both seeds:
- it loses with the row beyond error; or
- another rank beats the row's rank for it beyond error.

A race that ties goes with the row: a tie costs it nothing, and fewer keys keep the table plain. A row names races only when some race is left out. A race left out casts the default rule, the highest rank it has learned, as the game itself does. A race given another rank gets its own row.

**The default for a race the player picks in the app.** Every race of every row's class is measured, so no race the app offers is unjudged. A race no row names casts the default rule (the highest learned rank, never above the authored one). That is the client's own behaviour and the rotation as written, so it is always a legal cast. It is also what each left-out race measured as best.

## The race key (mechanism)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| K1 | **Rank rows are keyed by race, and the faction key is deleted.** `rankchoices.Row.Faction` is now `Races []string` (race slugs); a row with none applies to every race. The loader decodes strictly: a stale `"faction"` key, or a misspelt `"races"`, now fails instead of applying a row to every race. It refuses an unknown race, an empty list, a race named twice, and two rows that meet one race. `request.Build`, `RewriteRotationRanks` and the ladder take the character's race slug, and `sim/request/faction.go` (`FactionOfRace`) is deleted. spec-breakdown follows the entry's race and gains `-race`. rotation-search follows its band's race. The coach's `ChosenRanks` reads any race's row (a log's rank is no mistake where the site casts it for some race). Tests: `TestRankChoicesFollowTheRace`, `TestRaceRankChoicesNameRacesOfTheClass`, `TestBuildCastsTheRacesRankChoice`, `TestBuildRequestTakesTheRace`, `TestLoadRotationFollowsTheRace` (the spec-breakdown tests failed first, on the missing flag). | `sim/request/rank_choices.go`, `rotation_ranks.go:76`, `request.go:222`, `sim/cmd/spec-breakdown/main.go` | owner ruling 3 | - | S 33d1a01e4 |
| K2 | **The addon export and reader follow the race.** `faction_spell_ids` is now `race_spell_ids`. Data.lua writes `raceSpellId = { ["undead"] = .. }`, quoted because `night-elf` is no bare Lua name. It names only the races whose id differs from the line's every-race `spellId`. `Rotation.lua` shows the player's own id through `Compat.playerRaceSlug`. That function, with the client-token map (`Scourge` → `undead`, Skyborne by faction), moved out of `Export.lua`, so the export string and the rotation read one map. Tests: `test_a_race_rank_choice_below_sixty_names_each_races_rank`, `test_a_race_row_that_casts_the_every_race_id_names_no_race`, `test_the_lua_names_each_races_id_only_where_they_differ`, `rotation_spec.lua` "Rotation with a race's rank" and "Compat.playerRaceSlug". All failed first. The full addon busted suite (2,243) and luacheck pass. | `data/pipeline/addonrotation.py`, `addonlua.py`, `models.py`, `addon/ForeverSixty/{Compat,Export,Rotation}.lua` | owner ruling 3 | - | S 33d1a01e4 |

## The eight faction rows, re-judged per race

Gains are over the default rule, at seed 1 / seed 50001. Source: `out/judge.txt`.

| # | row | per race (the row's rank unless named) | verdict | origin | fix sha |
|---|---|---|---|---|---|
| R1 | Shadow Mind Blast at 50 (was: rank 2 Alliance, rank 3 Horde) | rank 2: Dwarf +2.54/+2.62, Night Elf +2.62/+2.65, Gnome +2.44/+2.50, Troll +2.31/+2.39. Rank 3: Human +2.01/+2.05 (rank 2 +1.3), Undead +0.40/+0.37 (rank 2 ties, -0.07). Error 0.21. | **rank 2 for Dwarf, Night Elf, Gnome, Troll; rank 3 for Human, Undead.** Round 12 gave the Troll rank 3 (as Horde) and the Human rank 2 (as Alliance), and never measured the Night Elf. | owner ruling 3 | S 2fbc44364 |
| R2 | Shadow Devouring Plague at bare 60 (was: rank 2 Alliance) | rank 2: Human +0.62, Dwarf +1.73, Night Elf +1.69, Gnome +0.77, Troll +1.82; Undead **-4.61/-4.67** (Dark Sacrifice pays for the top rank). Error 0.13 to 0.17. | **every race but the Undead.** The Troll gains 1.8%, which the faction key withheld. | owner ruling 3 | S 2fbc44364 |
| R3 | Fire Blast rank 3 at 30 (was: Horde) | Human +0.50/+0.70, Orc +0.76/+0.77, Undead +0.69/+0.73, Gnome **+0.59/+0.60** (tied in round 12), Troll +0.62/+0.64, High Order Skyborne +0.50/+0.67. Error 0.27. | **every race**, no key. The gain fell from round 12's +1.7% to about +0.7%, but every race now clears error. | owner ruling 3 | S 2fbc44364 |
| R4 | Destruction Corruption rank 3 at 40 (was: Horde) | Troll +0.16/+0.16, Human +0.14/+0.14 (error 0.13 to 0.14). Orc -0.09/+0.03, Undead +0.07/+0.01, Gnome -0.08/-0.06: all within error. | **every race**: no race loses beyond error. | owner ruling 3 | S 2fbc44364 |
| R5 | Discipline Binding Heal rank 1 at bare 60 (healer; was: Alliance) | Human +0.54/+0.51; Dwarf +0.17, Night Elf +0.20, Gnome +0.23/+0.13, Troll +0.17/+0.07; Undead -0.12/-0.09. All but the Human are within the 0.2% bar. | **every race.** Recheck at integration (healers lane: the new level-60 split moves this fight). | owner ruling 3 | S 2fbc44364 |
| R6 | Holy priest Prayer of Healing rank 1 at bare 60 (healer; was: Horde) | Dwarf +0.22/+0.22; Human +0.15/+0.18, Night Elf +0.22/+0.19, Undead +0.17/+0.12, Troll +0.22/+0.16, Gnome -0.18/-0.15: within the bar. | **every race.** Recheck at integration (healers lane). | owner ruling 3 | S 2fbc44364 |
| R7 | Holy priest Lightwell rank 1 at bare 60 (healer; was: Horde) | Human **-0.50/-0.47**, Undead **-0.26/-0.32**, Troll **-0.22/-0.22**; Dwarf, Night Elf, Gnome -0.13 to -0.19 (ties). Round 12 measured +1.3% for the Troll. | **row removed**: no race gains with it. A CONFIRMED regression of the row, for the healers lane to explain (probably round 12's sub-60 heal-profile change, which landed after round 12's ranks measurement). | R12 (healprofile ac97e6245, not isolated) | S 2fbc44364 |

R8 is the eighth row: Mind Blast at 50 had two faction rows, and R1 covers both.

## Regression over round 12's rows (round-13 base)

Every non-healer row still beats error on at least one published race at both seeds (`out/sweep.jsonl`). The six sub-1% rows hold:
- the three Aimed Shot rows;
- Fire Blast rank 4 at 40;
- Shadow Bolt rank 4 at 30;
- Prayer of Healing (see R6).

| # | row | result | verdict |
|---|---|---|---|
| G1 | Elemental Lightning Bolt 15207 rank 4 at 30/40/50/bare 60, and the paced 15208 rank 7 at 50/bare 60 | rank 4: +7.8 to +36.3% (Orc and Dwarf); rank 7: +0.40 to +0.52% (error 0.17) | holds, every race |
| G2 | Shadow Mind Blast and Devouring Plague by race | see R1, R2, and R9 to R11 below | re-keyed |
| G3 | Fire Blast rank 3 at 30 | R3 | holds, every race |
| G4 | BM Arcane Shot rank 3 at 40 | Orc +0.15/+0.20, Tauren +0.16/+0.13, Troll (published) +0.11/+0.11 (error 0.10). The Dwarf (published) and the other Alliance races tie: -0.03/+0.04. | holds for the Horde races; the others tie, so it stays every race. **Recheck at integration** (within 2x error). |
| G5 | BM Arcane Shot rank 1 at 50 and bare 60 | +0.23 to +0.38% at 50; +0.11 to +0.38% at bare 60 (Night Elf +0.11/+0.12, error 0.11) | holds, every race |
| G6 | Elemental Flame Shock rank 1 at 30 | Orc +1.37/+1.36, Dwarf +2.59/+2.59 | holds |
| G7 | Destruction Corruption rank 3 at 40 | R4 | holds at the edge. **Recheck at integration.** |
| G8 | Sub-1% rows | BM Aimed Shot 40 +0.18 to +0.27 (error 0.10); BM bare 60 +0.30 to +0.53; MM bare 60 +0.30 to +0.36; Fire Blast 40 Gnome +0.44/+0.47 (error 0.29); Shadow Bolt 30 +0.20 to +0.61 (error 0.12); Arcane Frostbolt 30 Gnome +0.37/+0.49 (error 0.25) | hold. Recheck at integration: Fire Blast 40 Gnome, Frostbolt 30 Gnome, Shadow Bolt 30 Human (+0.27/+0.20). |

All the other rows hold for every race, beyond error at both seeds. Druid-balance's three rows hold on all four druid races, including both Skyborne.

## New race rows and missed rows

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R9 | **A Dwarf or Night Elf Shadow priest at 40 to 49 cast Mind Blast rank 4 where rank 3 is better.** Rank 3: Dwarf +1.01/+1.01, Night Elf +1.39/+1.34. Rank 4: +0.44/+0.41 and +0.76/+0.75. Rank 2 is below rank 3 (+0.41 and +0.97). The others keep rank 4: Human +1.8%, Undead +1.4%, Gnome +1.4%, Troll +1.0%. Round 12 measured only the published Gnome and Undead. Test: `TestRound13RaceRankChoicesCastTheWinningRank` (failed first). | `out/judge.txt`, `confirm.py` | CONFIRMED-BUG (non-published races) | old (round 12's race check covered only the faction rows) | S 2fbc44364 |
| R10 | **A Night Elf Shadow priest at 50 to 59 cast Devouring Plague rank 1 where rank 2 is better.** Night Elf: rank 2 +2.17/+2.23 against rank 1 +0.52/+0.48. Rank 3 is +1.0. Every other race gains +2.9 to +4.2% with rank 1. Test: same. | same | CONFIRMED-BUG (non-published race) | old | S 2fbc44364 |
| R11 | **A Human Shadow priest at 50 to 59 lost 0.4% with Shadow Word: Death rank 2.** Human -0.36/-0.35, error 0.20. Every other race gains +0.2 to +1.3%. The Human now casts the default, rank 3. Test: same. | same | CONFIRMED-BUG (non-published race) | old | S 2fbc44364 |
| R12 | **Protection paladins at 40 to 49 cast Holy Strike rank 1.** Threat per second: Human +0.73/+0.78, Dwarf +0.79/+0.63, Undead +0.63/+0.89 over rank 5. Rank 2 ties (-0.05) and rank 3 loses (-0.5 to -1.0). The 5-mana rank leaves mana for more Judgements, Seals of Fury and Hammers of the Righteous: Hammer casts go from 10.9 to 12.8 a fight, Judgements from 5.0 to 5.9. At 30, 50 and 60, rank 1 loses 3% to 5%. Round 12's search skipped the tank bands, whose report prints no out-of-mana time. Test: same (failed first). | `explore-*.jsonl`, `confirm.py` | CONFIRMED-BUG | old (round 12's search looked only at bands out of mana > 3 s; the tank report never says) | S 34583ae09 |
| X1 | The rest of the missed-row search found nothing to adopt. It covered every mana-using non-healer spec (15) at 20/30/40/50/bare 60/raid 60, on both factions' published races, at 3,000 iterations: 2,091 runs, each ranked spell cast at least twice a fight at ranks cur-1, cur-2, cur-3, rank 1 and (where a row exists) cur+1. Above +0.15%: <br>- Ice Barrier, Divine Protection, Divine Shield: defensives, never rank choices. <br>- Balance Moonfire rank 4 at 30, +0.17 (error 0.29). <br>- Elemental's 15208 at 20, a line that barely fires. <br>- Destruction Immolate rank 4 at 40, +0.16 (error 0.25). <br>Rage and energy specs are out: in the client every rank of their ranked spells costs the same (spellconst cost per rank). | `explore-*.jsonl` | Aligned | - | - |

## Is each row's spell still there, still cast, still the client's?

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| L1 | All 80 rows name a trainer rank (`trainables/<class>.json`, 70753), learned by the row's first level (`tools/learn.py`). | learn.py | Aligned | - | - |
| L2 | Every row's spell is cast: no row's sweep run left the fight unchanged on any race (265 race-rows). `TestCuratedRankChoicesAreCastableDownranks` holds every row to an id its rotation names, and to a rank other than the default. | sweep.jsonl | Aligned | - | - |
| L3 | spellranks and spellconst for every rank of every row's spell are unchanged from 70707 to 70753. The committed 70753 hotfix push (fetched 2026-11-02 23:15 UTC) carries no SpellEffect, SpellMisc or SpellPower change (`tools/datacheck.py`). But see C1. | datacheck.py | Aligned with the committed data | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| C1 | **wago.tools now serves a hunter hotfix on 70753 that the committed push does not carry.** SpellEffect read with `useHotfixes=1` against without (fetched 2026-11-03 09:44 CST, `https://wago.tools/db2/SpellEffect/csv?build=1.60.1.70753&useHotfixes=1`): <br>- Aimed Shot: 20 → 40, 34 → 70, 55 → 115, 89 → 185, 125 → 260, 166 → 350 (19434, 20900 to 20904). <br>- Arcane Shot: rank 1 20 → 25 ... rank 8 217 → 280 (3044, 14281 to 14287). <br>- Serpent Sting: 111 → 135 a tick at rank 9 (25295); every rank about +20%. <br>- Aspect of the Hawk: 15 → 20 ... 55 → 110 (13165, 14318 to 14322). <br>- Trueshot Aura: 30 → 50, 40 → 75, 50 → 100, plus the Forever copies. <br>- Sniper Shot: doubled (1310687, 1310785, 1310786). <br>- Lacerating Strikes (1310536): a SpellMisc attribute change. <br>The other changes are NPC spells. Wowhead's Forever tooltips the same morning still show 166 / 217 / 555 (`https://nether.wowhead.com/forever/tooltip/spell/20904`, `14287`, `25295`). The engine uses the old values (`sim/hunter/client_damage.go:52`, `ArcaneShotBaseDamage`). Two sources disagree: wago's live CDN hotfix state against Wowhead's tooltip cache. To settle it: re-fetch the push (`pipeline fetch-wago-hotfixes`) and read the in-game tooltip; that is the same check ruling 5 already asks of the owner for Arcane Shot. If it is real, every hunter number moves (Aimed Shot about 2.1x), and the BM and MM Aimed Shot and Arcane Shot rows need re-measuring. | `r13/ranks/wago/changed.txt` | OPEN (new-data; hunter lane, data workflow) | new-data | - |
| C2 | **Era book ranks at 60 (guides lane's handoff: Immolate 25309, Heroic Strike 25286 and 31 more).** Forever's SkillLineAbility (70753, hotfixed) gives the 33 ids `curated/book-ranks.json` marks `book: true` the same row as their trainer ranks: AcquireMethod 0, MinSkillLineRank 1, no flags. Examples: 25309 and 11668, 25286 and 11567, 25307, 25311, 25306, 25304, 25345, 25292, 25357, 25315, 25295, 25289, 25288, 25300, 31016, 31018. Era's client looks the same, because books are taught server-side, so the client table cannot settle it. `book-ranks.json`'s own sources are the engine's IncludeAQ, the vanilla list and the owner. No rank row names a book rank. | wago SkillLineAbility csv | OPEN (trainables lane: a level-60 trainer, or Wowhead's "Taught by", settles it) | - | - |
| C3 | The tank report (`spec-breakdown -tank`) prints no error and no out-of-mana time. Threat at 40 came out identical for seeds 1 to 6 to one decimal (190.8). So a tank row's bar is the 0.2% convention, not a measured error. | `confirm.py` runs | OPEN (platform: print the threat's standard error and the mana-dry time) | old | - |

## Before/after

Spec-breakdown, bare preset, 10,000 iterations, published race. "Before" is `$R/sb-before` from `$R/site-before/sim`; "after" is this branch. Raid 60 does not move (no row this lane touched applies to the raid preset), so `before.tsv` is unchanged.

| spec, band, faction (race) | metric | before, seed 1 | after, seed 1 | change seed 1 / 50001 |
|---|---|---|---|---|
| mage-fire 30 Alliance (Gnome) | DPS | 73.3 | 73.7 | +0.59% / +0.60% |
| warlock-destruction 40 Alliance (Gnome) | DPS | 147.8 | 147.6 | -0.08% / -0.06% (within error) |
| paladin-protection 40 Alliance (Dwarf) | threat/s | 190.8 | 192.3 | +0.79% / +0.63% |
| paladin-protection 40 Horde (Undead) | threat/s | 191.2 | 192.4 | +0.63% / +0.89% |
| priest-discipline bare 60 Horde (Undead) | effective HPS | 320.7 | 320.3 | -0.12% / -0.09% (within the bar) |
| priest-holy bare 60 Horde (Troll) | effective HPS | 317.8 | 318.5 | +0.22% / +0.22% |
| priest-holy bare 60 Alliance (Gnome) | effective HPS | 330.1 | 329.5 | -0.18% / -0.15% (within the bar) |
| priest-shadow 40, 50, bare 60 (Gnome, Undead) | DPS | - | - | 0.00% (their races' ranks are unchanged) |

Non-published races, the gain over the table before this round: Shadow 40 Dwarf +0.6%, Night Elf +0.6% (R9). Shadow 50 Night Elf: Devouring Plague +1.7% (R10); Mind Blast unchanged (rank 2 as before). Shadow 50 Human: Shadow Word: Death +0.4% (R11), Mind Blast +0.7% (now rank 3). Shadow bare 60 Troll +1.8% (R2; the Night Elf, as Alliance, already had rank 2). Fire 30 Gnome +0.6%, High Order Skyborne +0.6% (R3).

## Files to regenerate

- `sim/request/testdata/ladder/*.golden.md` (TestRotationLadder): the ladder's races are the lowest race id per class, priests Undead. These move: mage-fire (30, Human now casts Fire Blast rank 3), warlock-destruction (40, Human now casts Corruption rank 3), paladin-protection (40, Holy Strike rank 1).
- addon-data: `data/builds/1.60.1.70753/addon-data.json` and `addon/ForeverSixtyData/Data.lua` (`pipeline addon-data`). Lines change at Shadow 40 and 50, Fire 30, Destruction 40 and Protection 40, and `raceSpellId` appears on the Shadow 40 and 50 lines (Mind Blast, Devouring Plague, Shadow Word: Death). Data.lua never carried round 12's `factionSpellId`, so no shipped addon reads the old field.
- `sim/rankchoices/rank-choices.json` is synced in each commit. Coach starters: none (the coach reads the embedded table).

## Measure at refresh

- Bare-band BiS picks where a published race's rotation changed: Fire 30 Alliance, Destruction 40 Alliance, Protection 40 (both), Discipline and Holy priest bare 60 (both).

## For other lanes

- **Healers:** R5, R6 and R7 re-keyed three bare-60 healer rows because the faction key is gone. Binding Heal and Prayer of Healing now apply to every race; Lightwell is removed (it now loses for Human, Undead and Troll). All three sit at the 0.2% bar. Re-judge them on the new level-60 split; key a race only where it loses or another rank wins beyond the bar.
- **Hunter / data:** C1, the live hunter hotfix (Aimed Shot about 2.1x, Arcane Shot about +29%, Serpent Sting +20%, Aspect of the Hawk, Trueshot Aura, Sniper Shot).
- **Trainables:** C2, the 33 Era book ranks the Forever client lists like trainer ranks.
- **Platform:** C3, the tank report's missing error and mana-dry time. Also, `Rotation.lua` now reads `raceSpellId`, and `Compat.playerRaceSlug` is the one race-slug reader (Export uses it).

## Still open

- C1, C2, C3.
- Recheck at integration (within 2x error, and the core lane's random regen tick phase may move them): BM Arcane Shot rank 3 at 40, Corruption rank 3 at 40, Shadow Mind Blast rank 3 at 50 for the Undead (+0.40/+0.37, error 0.22), Fire Blast rank 4 at 40 Gnome, Arcane Frostbolt rank 4 at 30 Gnome, Shadow Bolt rank 4 at 30 Human, and the three healer rows R5 to R7.
