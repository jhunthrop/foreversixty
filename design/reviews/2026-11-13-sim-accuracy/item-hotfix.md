# The hotfix layer on 1.60.1.70938 (issue #27)

Data lane. Scope: the item hotfix layer that lost push 16778086 on 70938 (Fiery War Axe 870 back to 2.9 s), widened mid-lane to every table the push layer covers, because on 70938 the spell side layered nothing either.

**Result: the spell-side and item hotfix layers now read every locale's pushed values from wago.tools' hotfix page. The 70938 data run 43468132768 (273623f0f) regenerated the build. Of the 100 items whose 70938 values differed from 70790 (the 66 stat changes plus names, binding and effect text), 25 were dropped hotfixes and are restored, 75 are real 70938 values, and 12 more items the old layer had wrong are corrected too. On the spell side the 70938 push now reaches 20 rows (Stormstrike, Maelstrom Weapon, Mental Dexterity, Gloves of Holy Might). A data test pins the server's values for seven items, and sim.yml runs the sim's unit tests in every mode.**

## 1. What broke

Two things changed at wago.tools between the 70790 run (2026-11-12) and the 70938 run (2026-11-13 06:53 UTC).

- **The hotfixed CSV export** (`db2/ItemSparse/csv?build=<b>&useHotfixes=1`) stopped applying push 16778086 on every build (70532 to 70938 all return ItemDelay 2900 for 870). The items depended on it alone (`pipeline.wago_hotfixes`).
- **The hotfix page** keeps every push, but every enUS entry of 70938's push carries `data: None`. The values sit on the other locales' copies (870: deDE, frFR, koKR and zhTW carry ItemDelay 3600). `pipeline.hotfix_push` read enUS only, so 70938's committed `hotfix-push.json` showed 0 changed rows in all 36 spell-side tables, every pushed record withheld. That is the coordinator's widening: push 114008 (Gloves of Holy Might) reached the engine only as a hand row.

## 2. The fix

Shas: 7e7d37d26 and 9133b3727 (the item values crawl), 8dfa31a97 (the spell side), f693ffb2c (the saved stores), 60d6782da (the guard and ladder goldens), f5baea7cb and 334b807a6 (sim.yml).

- **Values crawl** (`pipeline.hotfix_push_store.fetch_values`). An entry that carries values reads `<locale> <push> <record> <build> <table> [<id>,...`, so one search per table, `"<Table> ["`, reads only those, in every locale. The split crawl and the saved, incremental store are the spell side's own (`crawl_terms`, `fetch_saved`). Each store keeps one deciding entry per record: every push comes in up to eleven locales with the same values.
- **Spell side.** `layer_hotfix_push` reads the push twice: the enUS crawl as before, for the statuses (a removal or invalidation carries no values), and the values crawl into `raw/wago-push/values.json`. Both go through the one layer rule.
- **Items** (`pipeline.item_push`). The values crawl for ItemSparse and Item into `raw/wago-push/items.json`. `hotfix_merge.merge_hotfix_table` layers it last, over the play-session dump and the CSV export. The export can drop a push; the page keeps it.
- **The rule** (`newest_per_record`): per record, the newest push wins, but a valid entry with no values never outranks an older entry that has some. At 70938 no spell record changes under this rule alone (checked against the committed enUS entries). A row read from another locale keeps the localised text (`*_lang`) of the row under it, so Feurige Kriegsaxt stays Fiery War Axe.
- **Cost.** About 3 s a search on 13 Nov. The first 70938 crawls took 45 minutes and 8,592 searches for ItemSparse (6,771 records) and 6,352 searches for the spell side. Both stores are committed with the build, so the next run splits only the searches whose totals moved. A table new to a values store is crawled alone; it no longer forces the full crawl.
- **Tests**: `tests/test_item_push.py` (the real 70938 page entries and play-session rows for 870 and 2801: 870 swings at 3.6 s with its English name, 2801 takes the pushed 2.0 s and drops its five stats, the push beats the CSV export's row, another build's store is refused), plus `test_the_layer_takes_the_values_another_locale_carries` and the store tests in `tests/test_hotfix_push_store.py`.

Not covered: a removal or invalidation of an item record (no values, so the values crawl cannot see it; the CSV export still carries those).

## 3. The 66 items (100 with names, bonding and effect text)

Method: the regenerated 70938 items (273623f0f) against the 70938 commit before it (0238fc498) and against 70790, with the 70938 push's deciding row for each record and Wowhead Forever's tooltip for the doubtful ones.

**Dropped hotfix, restored (25).** 70790's value came from a push the CSV export no longer applies; the 70938 page still lists it, and the regenerated data matches 70790 again:
- Weapons: Fiery War Axe 870 (3.6 s, 116 to 175), Glowing Brightwood Staff 812 (94 to 142), Emberstone Staff 5201, Coldspire Staff 271215, Corpse Chopper 271217, Nail Spitter 17042, Dancing Flame 6806, Unstable Power Core 279847, Dull Sawblade 285292, Wisesight Wand 286750, Ukta's Conduit 286754, Serrated Knife 287063.
- Armor and rings: Tail Spike 6448, Slick Deviate Leggings 6480, Monkey Ring 6748, Heartwood Girdle 7000, Mantle of Woe 7750, Ghastly Trousers 15449, Abomination Bones 271205, Rotmender's Leggings 271207.
- Binding: Al'Aketh Bracers, Footwraps, Chain, Armguards and Shoes (277063, 277073, 277059, 277069, 277079) are Binds when equipped again. The last three keep 70938's "Al'aketh" spelling, which is a real client rename.
- Wowhead agrees on every one checked (870, 812, 5201, 6480, 271215, 277059, 286750, 287063).

**Real 70938 value (75).** The regenerated data keeps the 70938 value:
- 48: the 70938 push itself states it, for example Blade of Hanna 2801 (push 16780017: 2.0 s, no stats; 70790's 2.1 s with five stats came from the old play-session dump), Snake Hoop 6750, Spire of the Stoneshaper 12532 (64 spell power), Icemail Jerkin 1981, Energized Stone Circle 9522, Geomancer Headdress 274155, Furlfeather Ring 286745, Footman's Pocketknife 287059. Wowhead agrees on each of these.
- 18: the client row changed: the 14 Premier mail and scale pieces (272484 to 274217; Wowhead still shows 70790's stats for 272484, a stale scrape) and the four effect-text rows 289294 to 289297.
- 9: no push and no client change; the value comes from another source (classic-db, the required-level fill or a name): Quel'Serrar 18348 and 289302, 271939, 274718 to 274720, 277202, 281636, 281637. Wowhead agrees on 18348 and 274718.

**Also corrected (12)**, outside the 100 because 70790 had them wrong too: Prelacy Cape 7004, Shilly Mitts 9609, Hand of Edward the Odd 2243 (33 to 62), Boots of Avoidance 14549 (6 Strength, 11 Stamina, 1.5% dodge), 18755, 270023, 279864, 280031, 286748, 287004, 287036, 287070. Wowhead agrees on 2243, 7004, 9609 and 14549.

## 4. Spell and item rows the regenerated data changes, by class

The regenerated build's layer record (`hotfix-push.json`) now changes rows in five tables where it changed none. Every row, from `raw/` before and after the layer:

| Class | Table, row | Spell | Before (client 70938) | After (push) |
|---|---|---|---|---|
| shaman | SpellEffect 1368308 | Stormstrike 17364 | 150% weapon damage | 100% |
| shaman | SpellCooldowns 55241 | Stormstrike 17364 | 6 s | 8 s |
| shaman | SpellEffect 1137418 | Maelstrom Weapon 408498 | 50 | 2 |
| shaman | CurvePoint 251361 to 251364 (curve 85306) | Maelstrom Weapon, ranks 1 to 4 | -4, -8, -12, -16 | -20 each |
| shaman | TraitDefinitionEffectPoints 25235, CurvePoint 372521 to 372525 (new curve 128299) | Maelstrom Weapon effect 1 | none | 6, 7, 8, 9, 10 by rank |
| shaman | CurvePoint 251264 to 251266 (curve 85272) | Mental Dexterity 415140 | 50, 100, 150 | 33, 66, 100 |
| shared (items) | SpellAuraOptions 249242 | Gloves of Holy Might 1330855 (push 114008) | ProcTypeMask 69636, 32 | 69972, 0 |
| none | SpellEffect 1129612, 1129613 | [DNT] Wendel Mathers Escape 437589 | Effect 169 | Effect 3 |

Wowhead Forever shows Stormstrike at 100% and 8 s and Maelstrom Weapon's proc rate on weapon speed, as the push does. Published data that moves: `spellconst/shaman.json` (Stormstrike's cooldown and amount, Maelstrom Weapon's amount), `trainables/shaman.json` and the Stormstrike tooltip. No other class's spellconst, trainables or tooltips change.

For the sim lanes: MR46 kept Stormstrike, Mental Dexterity and Maelstrom Weapon as engine hand values "until a client or push row states them"; the push now states all three, and the data carries them. Gloves of Holy Might's trigger (MR106, a hand row) is now in the data as well. The engine's next follow should read them and drop the hand rows; whether its hand values already match is the shaman and paladin lanes' check.

## 5. The guard

`data/tests/test_server_pushed_items.py` reads the active build's `simdb.bin` and `items/*.json` and pins the server's values, each checked against Wowhead Forever's tooltip on 13 Nov:

| Item | Server value |
|---|---|
| Fiery War Axe 870 | 3.6 s, 116 to 175, +15 Strength |
| Blade of Hanna 2801 | 2.0 s, 96 to 145, no stats |
| Emberstone Staff 5201 | 3.0 s, 34 to 51, Stamina 8, Fire Resistance 8, 24 spell power |
| Coldspire Staff 271215 | 3.6 s, 43 to 66, Frost Resistance 13, 32 Frost power |
| Wisesight Wand 286750 | 1.7 s, 10 to 20, +1 Intellect |
| Serrated Knife 287063 | 2.0 s, 10 to 20, +1 Agility, +1 Spirit |
| Slick Deviate Leggings 6480 | uncommon, 68 armor, +4 Agility, +4 Stamina |

The weapons are checked in `simdb.bin` (what the sim swings) and in the published items, the leggings in the published items. data.yml runs it in the always-run list and again straight after `simdb` in the data run, so a run that loses a push fails before it commits. On the old data it fails on 870; on 273623f0f it passes (13 tests).

It runs in every mode: data.yml's always-run list carries it beside the addon-data contract tests.

## 6. sim.yml

`vars.LAUNCH_MODE` made sim.yml skip every `go test`, so `TestWeaponsSwingAtTheirForeverSpeed` (the test that caught 870) never ran on CI. The test job now runs `go test` on every package but `cmd/spec-breakdown` in every mode: adapter, api, bulk, cmd/addon-builds, cmd/dps-per-point, cmd/forever-measure, cmd/forever-sim, cmd/heal-builds, cmd/leveling-bis, cmd/rotation-search, cmd/talent-search, cmd/wasm, combine, enginever, internal/enginetalents, internal/inproc, internal/simdb, internal/simdrain, internal/spellranks, internal/strcase, leveling, measure, rankchoices, request, runner, score, specs and talents. On the box that is 59 s of tests (request 57 s is the longest; spec-breakdown alone took 318 s and stays out). On CI (run 43491093975) the step takes 2 min 26 s and the whole test job 4 min. The `-race` run with coverage over everything, spec-breakdown included, still runs only when LAUNCH_MODE is off. The runbook's launch-mode list says so.

Turning the tests on found that the rotation ladder goldens had moved with this lane's data: 9 ladders change one row each, all on the restored items (Wisesight Wand is the level 10 casters' wand again, Fiery War Axe swings at 3.6 s for Retribution and Arms at 38, Arms at 50 leaves Glowing Brightwood Staff for 2915). Regenerated on the box at the pinned engine (5603877e4).

## 7. BiS picks affected

Published picks on 70938 that wear an item this data run changes (ranked before it):

| Item | Change | Bands |
|---|---|---|
| Fiery War Axe 870 | 2.9 s back to 3.6 s | Retribution, Enhancement and Arms bare 40, both factions |
| Glowing Brightwood Staff 812 | 127 to 191 back to 94 to 142 (stats unchanged) | Restoration druid bare 50 and raid 60; Discipline bare 50, Horde bare 60 and raid 60; Holy priest bare 50 and raid 60; Restoration shaman bare 50, bare 60 and raid 60 |
| Hand of Edward the Odd 2243 | 50 to 94 to 33 to 62 (stats unchanged) | Balance and Shadow bare 60 |
| Boots of Avoidance 14549 | 14 Agility and 24 dodge to 6 Strength, 11 Stamina, 18 dodge | Protection paladin bare 40, 50 and raid 60; Protection warrior bare 40 to 60 and raid 60; Arms bare 40 |
| Prelacy Cape 7004 | rare to uncommon, fewer stats | Restoration druid and Discipline bare 30 (Alliance) |
| Shilly Mitts 9609 | rare to uncommon, fewer stats | Balance and Destruction bare 30 (Alliance) |

Blade of Hanna 2801 (Retribution bare 60) does not change: 70938's 2.0 s with no stats is the server's value, so MR105's "2.1 s again" was wrong. Stormstrike and Maelstrom Weapon move every Enhancement band. The shaman lane should check the engine against them before the re-rank.

## 8. Re-rank

MR107 in the ledger (supersedes MR105). Rankings 43445589374 and 43462171108 ran on the data before this run; none was cancelled.
