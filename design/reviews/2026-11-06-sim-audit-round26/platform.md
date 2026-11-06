# Sim-accuracy round 26: platform lane

Site branch `wt/r26-platform` (from int/r26). Fork branch `lane/r26-platform` (from int/r26 f871a65ae). Client 1.60.1.70790. BUILD MODE: touched tests only, no regens. Every sim went through `$R/sim.sh`. Every CONFIRMED-BUG has a test that was red before its fix.

**Confirmed and fixed: 5.** P6 `old`; P7 `R21 559e39867` (with round 25's P1/P2); P8 `R25 f5d1cd239`; V1 `old`; I1 `old`.
**Cross-class DPS movers:** none in published numbers. P7 and P8 change only `spec-breakdown -level 10` to `19` (a measuring tool). P6 changes the TestRotationLadder goldens (listed below, not regenerated). V1 changes no test output.

## Live logs

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| L1 | Still no new live Forever logs. `GET https://api.foreversixty.gg/v1/reports/recent` lists only `z2opt2x2gs5u` (20 Oct, no kills). Nothing in this lane needs a log. | last looked 21:37 UTC 6 Nov | OPEN (no data) | - | - |

## Seed 1: the ladder's own gear pickers (cross-cutting)

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P6 | **The ladder handed characters the other faction's items.** `ladderGear` (sim/request/ladder.go) ran `pickGearItem`, `pickShieldItem` and `pickWandItem` over the whole class table. Round 25's P5 filtered only the below-band refill. Across both factions, 31 picks in 21 spec/level pairs were items only the other faction can earn: for example Alliance casters at 20 held Crescent Staff 6505 (Horde only, Leaders of the Fang), an Alliance Holy paladin at 10 held Mug of Muddled Memories 277247 (a Horde quest's), Horde shamans at 38 and 40 held Salbac Shield 4652 (Alliance only) and Horde warriors, rogues and hunters at 38 Sword of Serenity 6829 (Alliance only), and Alliance Assassination held the Horde-only Scout's Blades 19545/19544. `ladderGear` now takes the character's other-faction set (`loadOtherFactionItems`, the same one `LevelGear` uses) and drops those items before any pick. The ladder reads its faction off the smoke race (`LoadRaceFaction`, races.json). | `sim/request/ladder.go` (`ladderGear`), `sim/request/level_gear.go` (`withoutOtherFaction`, `LoadRaceFaction`), `sim/request/ladder_test.go`; test `TestTheLadderGearIsOneTheFactionCanEarn` (every spec, both factions, every ladder level; red: 31 violations) | CONFIRMED-BUG (tool) | old (the ladder's pickers predate faction data; round 25 noted it as still open) | S e0a7c346a |
| P8 | **A quest of unknown faction counted as the other faction's, for both factions.** loot.json carries 41 quest rows whose faction it could not read (`"faction": "unknown"`). Round 25's filter treated any quest not named `both`, `""` or the character's own faction as closed, so an unknown quest's reward was closed to Horde and Alliance alike. The BiS ranker reads unknown as open to both (`questFactionSide`, sim/cmd/leveling-bis/data.go:891). Example: Scholarly Pendant 277203 and Erudite's Amulet 277204, rewards of Friend of the Library, which both factions take. The filter now closes a quest only when it names the other faction. | test `TestAQuestOfUnknownFactionIsOpenToBoth` (red for both factions); the armor guard below also flagged both necks | CONFIRMED-BUG (tool) | R25 f5d1cd239 | S e0a7c346a |

## Seed 2: armor slots below 20

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| P7 | **Below 20 an armor slot the level could not hold stayed empty** when no runner-up fit. `LevelGear.replacement` refilled weapons, shields and wands from the class table, but returned nothing for armor. Across 28 specs, both factions, levels 10 to 19 there were 3,537 empty armor, jewellery and trinket slots that band 20 fills. `pickArmorItem` now refills them from the class table: same table slot (both rings are `finger`, both trinkets `trinket`), an item the level holds (`EffectiveRequiredLevel` with quest floors), the class's armor proficiency (mail for shamans and hunters and plate for warriors and paladins only from 40, owner rule 30 Sep), the faction, and the build's sim database. It scores the sum of the spec's weight stats on the item, then item level, then the lower id. A relic stays a relic of its kind. The proficiency table moved from `cmd/leveling-bis/eligible.go` to `leveling.ArmorProficiencyLevel` so the ranker and the tool share one source (no behaviour change in the ranker; `TestEligible` and `TestPublishedBISFilesPassSanityChecks` pass). After the fix the only empty slots are neck and shoulder at 10 to 16: no class can earn a neck or shoulder before 17 (first necks: Erudite's Amulet 277204 and Roadwatcher's Confidence 281265, quest floor 17; first shoulders: Double-Stitched Woolen Shoulders 4314, required level 17). | `sim/request/level_gear.go` (`pickArmorItem`, `tableSlot`), `sim/leveling/armor_proficiency.go`; test `TestNoSpecBelowTheLowestBandHasAnEmptyArmorSlot` (every spec, both factions, 10 to 19: every armor slot band 20 fills is filled from the level where the slot can be earned, no armor the class cannot wear yet, nothing only the other faction earns). Red: 3,537 empty slots and 48 other-faction necks (P8). | CONFIRMED-BUG (tool) | R21 559e39867 (`-level` below 20), left open by round 25 | S e0a7c346a |

## Seed 3: fork go vet

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| V1 | **`go vet ./sim/core` failed** at `sim/core/test_generators.go:535`: `testOptions := *StatWeightsDefaultSimTestOptions` copies a proto message that holds a mutex. It is now `googleProto.Clone(StatWeightsDefaultSimTestOptions).(*proto.SimOptions)`. SimOptions holds only scalars, so the cloned request has the same field values and no generated test output changes. Note: the class golden suites in the fork (for example `sim/paladin/protection`, `sim/warrior/dps_warrior`) panic before and after this change on `No item with id: 12620` / `12640` (their Era test gear is not in the Forever item database); that is pre-existing and unrelated. | `go vet ./sim/core` red, then clean | CONFIRMED-BUG (build hygiene) | old | E ac387de1b |

## Seed 4: the Deep Wounds icon

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| I0 | **The icon itself is right.** The client gives Deep Wounds 12834 (the talent) and 12162 (the bleed it triggers, SpellEffect 12834 e0 aura 42 -> 12162) icon file 132090, which is `ability_backstab` (the same file as Backstab 53). Wowhead Forever (12834, 12162) and Wowhead Classic (12834, 12721) show `ability_backstab`. `spellicons.json` maps 12867 to `ability_backstab`; the talent tree draws `ability_backstab` (talents/warrior.json node 105950); `icons/ability_backstab.webp` is the dagger-and-blood Backstab art. The addon draws game textures and names Deep Wounds only as a talent; the coach and guide cards carry no Deep Wounds id. | `$SP/b70790/70790/SpellMisc.csv`, SpellEffect.csv, SpellName.csv; https://nether.wowhead.com/forever/tooltip/spell/12834, /12162; https://nether.wowhead.com/classic/tooltip/spell/12834, /12721 (read 6 Nov) | Aligned | - | - |
| I1 | **Deep Wounds at one or two talent points had no icon.** The engine reports the bleed under one id per point (`deepWoundsSpellIDs = {12834, 12849, 12867}`, fork sim/warrior/talents.go:163). Only 12867 had an icon source (round 4 Nov, `icon_rows` 12867 from 12162), so a leveling warrior's result page drew no icon for its Deep Wounds row while the build held one or two points. `extra-spell-icons.json` now lists 12834 (a client row) in `spell_ids` and maps 12849 to 12162's icon in `icon_rows`. | `data/curated/extra-spell-icons.json`; test `test_every_deep_wounds_rank_the_engine_reports_has_an_icon_source` (red: 12834 and 12849 missing) | CONFIRMED-BUG (page) | old (10cb0b576, 4 Nov) | S 052925e67 |

## Before/after

Raid and band-floor numbers do not move: the refill and the faction filter run only below the band, and the ladder pickers do not feed spec-breakdown. Spot check: Arms raid 60, horde, seed 1, frozen BiS, guide build: 846.8 before and after (`$R/before.tsv`).

Below the band (bare band 20, horde, 10,000 iterations; seeds 1 and 50001 agree within 0.1 DPS; before = `$R/sb-before`, after = `$R/platform/sb`; `$R/platform/below20.tsv`):

| spec | level | before | after | change |
|---|---|---|---|---|
| Combat | 12 | 15.4 | 18.4 | +19.5% |
| Combat | 16 | 26.6 | 27.4 | +3.0% |
| Fire | 12 | 11.4 | 13.0 | +14.0% |
| Fire | 16 | 24.6 | 25.0 | +1.6% |
| Arms | 14 | 19.7 | 21.0 | +6.6% |
| Arms | 18 | 28.5 | 28.7 | +0.7% |
| Beast Mastery | 12 | 42.3 | 46.1 | +9.0% |
| Beast Mastery | 16 | 64.1 | 65.9 | +2.8% |

## Re-rank needed

None. Below-20 gear is not in the published BiS bands, the ladder pickers do not feed the ranker, and the ranker's armor gate reads the same table with the same values.

## Files to regenerate (integrator, at the pin)

- `TestRotationLadder` goldens (sim/request), from P6. The ladder smokes Alliance for every class but shaman (Orc) and priest (Undead), so these rungs change their weapon, shield or wand: druid-balance 20, druid-feral-bear 10, mage-arcane 20, mage-fire 20, mage-frost 20, warlock-affliction 20, warlock-demonology 20, warlock-destruction 20, paladin-holy 10, rogue-assassination 30 and 38 (off hand), shaman-elemental 10, 38 and 40, shaman-enhancement 38 and 40, shaman-restoration 30.
- `data/builds/1.60.1.70790/spellicons.json` through `python -m pipeline icons` (needs raw/), from I1: adds 12834 and 12849 (`ability_backstab`).

## Measure at refresh

- Nothing for the BiS.

## For other lanes

- **ranks:** your re-judging of the 20 caster and healer rows at 10 to 19 should run on this branch (e0a7c346a or later). Below-band characters now wear armor in every slot they can earn (P7) and keep neutral quest rewards (P8). Mana-bound rows can move: Fire at 12 gains 14% from armor alone.
- **Any lane, the result page:** 54 more engine-only ids in `data/curated/engine-spell-names.json` have a name but no icon source (Windfury Totem 8514/10607/10611, Clearcasting 12577, Ignite 12654, Shadow Weaving 15331-15334, Spirit Tap 15335-15338, Flurry 16277-16280, Furor 17058-17061, Vengeance 20059, Redoubt 20134, Seal of Command 20944-20947, Holy Shield 20955-20957 and 27179, Master Demonologist 23825, Elemental Devastation 29177/29178, Seal of Fury 1311648, Sudden Insight 450527/450576, Eureka! 1000101, and the world buffs). Each needs an `icon_rows` entry from its client row, checked against SpellMisc the way I1 was.

## For the guide lane

- Nothing: no guide, APL or card text changed.

## Still open

- L1 (logs).
- The fork's class golden suites panic on Era test items the Forever item database lacks (12620, 12640). Pre-existing; not this lane's.
