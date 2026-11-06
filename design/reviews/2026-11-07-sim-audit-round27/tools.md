# Sim-accuracy round 27: tools lane

Site branch `wt/r27-tools` (from int/r27 af9721524). Fork branch `lane/r27-tools` (from int/r27 227d95bd1). Client 1.60.1.70790. Ship mode: touched tests only, no regens, no sims needed (no fix changes a sim number). Every CONFIRMED-BUG has a test that was red before its fix.

**Confirmed and fixed: 3.** T1 `old` (S c7ec83f10), T2 `old` (E ec0b9f1c0; the round-26 diagnosis was wrong, see below), T3 `old` (S 6a2c38dd5).
**Cross-class DPS movers:** none. T1 changes only what spec-breakdown accepts; T2 changes a panic message; T3 changes icons on the result page.

**Lane note:** the disk was full when the lane started (2.6 GB free, falling to under 120 MB, from writes outside this lane). A full site worktree (3.8 GB) did not fit, so `wt/r27-tools` is checked out **sparse**: only `data/builds/1.60.1.70790`, no `apps/`, and under `design/` only round 26, `sim-accuracy`, `talent-search` and this round. Commits are normal; the integrator merges the branch as usual. No Go test binary with `--tags=with_db` was built (see T2).

## Live logs

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| L1 | Still no new live Forever logs. `GET https://api.foreversixty.gg/v1/reports/recent` lists only `z2opt2x2gs5u` (20 Oct, 42 fights, no kills). Nothing in this lane needs a log. | last looked 23:48 UTC 6 Nov | OPEN (no data) | - | - |

## Seed 1: `spec-breakdown -race`

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| T1 | **spec-breakdown measured characters no player can make.** `-race` replaced the published entry's race without a check (`sim/cmd/spec-breakdown/main.go:454`), so a Skyborne priest ran without a word. The client's race-class pairs (`data/builds/1.60.1.70790/combos.json`) give priests human, dwarf, night elf, undead, gnome and troll only. A race of the other faction ran too, wearing this faction's published gear and buffs. `legalRequest` (the guard every run goes through: main, heal, tank) now refuses a race the class cannot be in the run's `-faction` and names the legal ones: `-race "windshaper-skyborne": illegal race: a horde priest is one of undead, troll`. The legal-race reader moved from the BiS ranker's race sweep (`sim/cmd/leveling-bis/race_sweep.go`, round 25 S b49b84ffe) to `request.LegalRaces` (`sim/request/races.go`), so the ranker and the tool read one source. `TestEveryPublishedBuildIsLegal` now also holds every published entry's race to it. | tests `TestARaceTheClassCannotBeIsRefused` (red: all five accepted: Skyborne Shadow and Holy priests, Orc Retribution, Human Arms on Horde, an unknown slug), `TestEveryLegalRaceIsAccepted`, `TestEveryPublishedBuildIsLegal` (sim/cmd/spec-breakdown/legal_test.go); `TestLegalRacesReadTheClientCombinationsPerFaction` moved to sim/request/races_test.go; leveling-bis race-sweep tests pass; `go vet` clean on sim/request, the two commands and api | CONFIRMED-BUG (tool) | old (`-race` predates the race sweep's legal-race source) | S c7ec83f10 |

## Seed 2: the fork's class suites and items 12620 / 12640

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| T2a | **Round 26's diagnosis was wrong: the Forever item database has both items.** Enchanted Thorium Helm 12620 and Lionheart Helm 12640 are in the fork's `assets/database/db.json` (7,573 items). Every item of every `ui/*/gear_sets/*.gear.json` is in it. 12620 is the first (head) item of `ui/tank_warrior/gear_sets/forever_l60` and `ui/protection_paladin/gear_sets/forever_l60`; 12640 is the head of `ui/warrior/gear_sets/phase_1` (dps_warrior's set). The panic comes from running the suites without `--tags=with_db`: then `sim/core/database_load.go` is left out, `ItemsByID` is empty, and `NewItem` panics on the gear set's first item. `make test` runs `go test --tags=with_db ./sim/...` (fork makefile:250). No test gear needs replacing and no test needs a skip. | fork `assets/database/db.json`; scan of every gear set file against it (none missing); `sim/core/database.go:259-263`; `sim/core/database_load.go:1` | Aligned (no missing items) | - | - |
| T2 | **The panic named an item, not the missing tag,** which is how round 26 misread it. `NewItem` now says `No item with id: 12620 (the item database is not compiled in; test suites need --tags=with_db)` when `WITH_DB` is false. With the tag the message is unchanged. Production is unaffected: the site sends its own items in the request database, and an id missing there still panics, now with the hint. | test `TestAnUnknownItemWithoutTheDatabaseNamesTheBuildTag` (sim/core/item_lookup_test.go; red on the old message, skips under the tag) | CONFIRMED-BUG (test hygiene) | old | E ec0b9f1c0 |

Not run: the protection paladin, tank warrior and dps warrior suites with `--tags=with_db`. The disk had under 200 MB free, and a with_db test binary would not fit. No gear set or golden changed, so CI's `make test` decides whether those suites pass.

## Seed 3: icons for engine-only spell ids

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| T3 | **54 ids the engine reports had a name but no icon on the result page.** Each id in `data/curated/engine-spell-names.json` was missing from `spellicons.json` and from `extra-spell-icons.json`. Checked against the 70790 client (SpellName and SpellMisc, the full 32,088 rows) and Wowhead Classic or Forever tooltips (read 6 Nov): **17 are client rows** and join `spell_ids`: the world buffs 15366, 16609, 22817, 22818, 22820, 22888, 23735 to 23738, 23766 to 23769 and 24425, plus the set procs 450527 and 450576 (Sudden Insight). **35 are Era ids the Forever client does not carry.** Each takes its family's client row through `icon_rows`, and each source row's icon file matches Wowhead Classic's icon for the engine id: Windfury Totem's buff 8514, 10607 and 10611 from 10610 (`spell_nature_windfury`), Clearcasting 12577 from 12536 (`spell_shadow_manaburn`), the Ignite burn 12654 from the talent 11119 (file 135818, `spell_fire_incinerate`), Shadow Weaving 15331 to 15334 from 15258 (`spell_shadow_blackplague`), Spirit Tap 15335 to 15338 from 15271 (`spell_shadow_requiem`), the shaman's Flurry 16277 to 16280 from 16257 (`ability_ghoulfrenzy`), Furor 17058 to 17061 from 17056 (`spell_holy_blessingofstamina`), Vengeance 20059 from 20050 (`ability_racial_avatar`), Redoubt 20134 from 20128 (`ability_defend`), the Seal of Command proc 20944 to 20947 from 20424 (`ability_warrior_innerrage`), Holy Shield's damage 20955 to 20957 from 20925 and 27179 from rank 4, 1239570 (`spell_holy_blessingofprotection`), Master Demonologist 23825 from 23759 (`spell_shadow_shadowpact`), Elemental Devastation 29177 and 29178 from 30165 (`spell_fire_elementaldevastation`), and the Seal of Fury shield 1311648 from rank 1's proc 1311647 (`spell_holy_retributionaura`). The engine chose 1311648 itself, and Wowhead Forever has no entity for it. **2 were stale.** Elune's Light 58984 and Eureka! 1000101 have not been reported since the fork moved the racials to client ids on 7 Oct (fork d1af3529f). The engine now reports 1259799 (Elune's Light, `spell_nature_moonglow`) and 1259812, 1259813, 1259817, 1259821 and 1259823 (Eureka!, `inv_gnometoy`), and none of those had an icon either. The two stale names are deleted, as the file's own notes ask, and the six client ids join `spell_ids`. The `notes` field lists every row and its source. | `data/curated/extra-spell-icons.json`, `data/curated/engine-spell-names.json`; client `$SP/b70790/70790/SpellName.csv`, `SpellMisc.csv`; https://nether.wowhead.com/classic/tooltip/spell/<id> for each engine id and source row, https://nether.wowhead.com/forever/tooltip/spell/1259799, /1259812, /1311647, /1311648; test `test_every_engine_only_spell_id_has_an_icon_source` (data/tests/test_addonrotation.py; red: 54 missing) | CONFIRMED-BUG (page) | old (each entry since it was added to engine-spell-names.json; round 26 counted them) | S 6a2c38dd5 |

Could not source: none. 27179 has no Wowhead Classic entry (it is a TBC id). It takes the engine's own named client row, rank 4 1239570.

## Before/after

No sim number moves. T1 refuses inputs before any sim runs (`-race windshaper-skyborne -spec priest-shadow` now exits with the error above; a legal race builds the same request as before). T2 changes a panic string. T3 is page data. No sims were run.

## Files to regenerate (integrator, at the pin)

- `data/builds/1.60.1.70790/spellicons.json` through `python -m pipeline icons` (needs raw/), from T3. It adds the 23 `spell_ids` and 35 `icon_rows` above, and downloads any icon art not yet present: `spell_shadow_blackplague`, `spell_fire_incinerate`, `spell_shadow_shadowpact`, `spell_fire_elementaldevastation`, `spell_nature_moonglow`, `inv_gnometoy` and the world-buff icons may be new.
- The web's `simnames/_shared.json` (web/scripts/sync-data.mjs) at the next sync drops 58984 and 1000101.

## Re-rank needed

None.

## Measure at refresh

- Nothing.

## For other lanes

- **Integrator / CI:** the fork's class suites (`sim/paladin/protection`, `sim/warrior/tank_warrior`, `sim/warrior/dps_warrior`) need `--tags=with_db`. Run them that way before reading any panic as missing data.
- **Any lane, the result page:** the passive racial rows the engine registers (5227 Underwater Breathing, 20552 Cultivation, 1259691 Attuned, 1259707 Elemental Insight, 1259710 Wind Blessed, 1259719 Artisan's Precision, 1259721 Big Game Hunter, 1259802 and 1259803 Expansive Mind, 1259918 Plainsrunning) are client rows with no icon in `spellicons.json`. They are not in engine-spell-names.json, and I did not confirm that a result ever lists them as actions. If one shows, add it to `spell_ids`.

## For the guide lane

- Nothing: no guide, APL or card text changed.

## Still open

- L1 (logs).
- The three class suites under `--tags=with_db` were not run here because of the disk. CI decides.
