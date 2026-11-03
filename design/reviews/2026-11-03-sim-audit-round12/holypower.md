# Sim audit round 12: Elixir of Holy Power (holypower lane)

One item, routed from itemscrit K4. Engine `lane/r12-holypower` (from 2606bc05e), site `wt/r12-holypower` (from c6bdc855c). Client 1.60.1.70753 with the hotfix push (the `.hf.csv` copies in `$R/itemscrit/wago/` agree row for row with the base tables for every row below).

## Elixir of Holy Power

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| H1 | **The engine had no Elixir of Holy Power.** Forever turned Era's Elixir of Greater Firepower (item 21546) into "Elixir of Holy Power": 40 Holy spell damage for 30 minutes, level 40. Nothing else in Classic gives Holy spell damage, and Classic elixirs of different schools stack (no Battle/Guardian split before 2.1), so it takes no other elixir's place. Retribution's Seal of Command, its Judgement, Consecration and Hammer of Wrath are Holy; a Protection paladin's threat is Holy damage. The engine now has `Consumes.holy_power_buff` (`HolyPowerBuff.ElixirOfHolyPower`, +40 Holy power), the site vocabulary gains `elixir_of_holy_power` (level gate 40 from items.json), and both paladin raid presets carry it. | ItemSparse 21546 "Elixir of Holy Power", RequiredLevel 40, Item class 0/2; ItemXItemEffect 140014 → ItemEffect 103511 (spell 1310077, category 2593, 3 s category cooldown); SpellEffect 1339308: effect 6, aura 13, base 40, misc 2 (Holy); SpellMisc: duration index 30 (1,800,000 ms), school mask 2; SpellName 1310077 still reads "Greater Firepower", Spell description "Increases spell Holy damage by up to $s1 for $d." Wowhead Forever tooltips item/21546 and spell/1310077 (2026-11-03): "Increases spell Holy damage by up to 40 for 30 min." Present as "Elixir of Holy Power" in every Forever build's items.json since 1.60.1.69893 (forever-prebeta still read Greater Firepower). | CONFIRMED-BUG (missing consumable) | old (since 69893; no round before 12 looked; itemscrit K4 found it) | E 46be83bcd, S 408d35cde |
| H2 | **Obtainable at launch.** Recipe 21547 "Recipe: Elixir of Greater Firepower" (Alchemy 250) teaches spell 26277, whose SpellEffect 703706 creates item 21546 and whose SpellName is "Elixir of Holy Power". The recipe drops from Dark Iron Taskmasters (npc 5846, levels 47-48, 1.0%) and Dark Iron Slavers (1.1%) in Searing Gorge; reagents Firebloom, Fire Oil, Crystal Vial. | ItemXItemEffect 21547 → ItemEffect 223217 (trigger 6, spell 26277); SpellEffect 703706 (effect 24, item 21546); data/builds/1.60.1.70753/loot.json and raw/classicdb/sources.json + crafted-recipes.json; Wowhead Forever item/21547 (2026-11-03). The drop source is classic-db only (single source); the recipe and craft spell are client rows. | Aligned (obtainable) | - | - |
| H3 | **Elixir of Greater Firepower gave a fire mage +40 fire that no Forever item gives.** The engine's `FirePowerBuff.ElixirOfGreaterFirepower` was Era's item 21546. In Forever that item is the Holy elixir and the client carries no 40-fire elixir, so a /sim custom pick of `elixir_of_greater_firepower` added 40 fire spell damage out of nothing. The value is now deprecated and applies nothing (as `LesserArcaneElixir` already does). No preset or kit carries it (fire specs carry Elixir of Fire Power), so no published number moves. | As H1; the client's only fire elixir is Elixir of Fire Power 6373 (10). `sim/core/consumes.go` applied FirePower 40 for the value. | CONFIRMED-BUG (custom only) | old (upstream; Forever since 69893) | E 46be83bcd |

Tests: `sim/core/consumes_holy_power_test.go` (`TestElixirOfHolyPowerMatchesTheClient`, `TestElixirOfGreaterFirepowerIsGone`) failed first (no field; then 40 fire), then passed. Site: `TestRaidPresetConsumesFollowTheRole` gained the two paladin rows and failed against the old presets.json, then passed; `TestItemLevelsReadTheClientsRequiredLevel` gained `elixir_of_holy_power` at 40; `gear-extras.test.ts` failed four cases without the new `CONSUMABLE_GIVES` line, then passed. `TestRaidConsumesJSONIsCommitted` and `TestIDsMarkdownIsCommitted` pass with the regenerated files. The fork's mage, shaman and warlock character tests that name `ElixirOfGreaterFirepower` behave the same before and after (the mage and shaman ones fail at base 2606bc05e already, `No item with id: 14332`; warlock and warden pass both ways), so no fork golden moved.

Proto: `proto/common.proto` gains `enum HolyPowerBuff` and `Consumes.holy_power_buff = 38`; `sim/core/proto/common.pb.go` regenerated with the fork's `make sim/core/proto/api.pb.go` (protoc 36.2 per `.protoc-version`, protoc-gen-go v1.36.6 per go.mod); only that file changed. The fork's TS proto is not committed (ui/core/proto/.gitignore).

## Before/after

Raid 60, 10,000 iterations. Before is `$R/before.tsv`. After is engine 46be83bcd with site 408d35cde. Retribution is DPS, Protection threat per second. I checked that the lane engine without the preset row reproduces the baseline exactly (Retribution horde s1 690.1, Protection horde s1 1101.0).

| spec | alliance s1 | alliance s50001 | horde s1 | horde s50001 | absolute (alliance s1) |
|---|---|---|---|---|---|
| paladin-retribution | +1.63% | +1.63% | +1.64% | +1.62% | 693.0 → 704.3 |
| paladin-protection | +2.28% | +2.26% | +2.24% | +2.15% | 1088.3 → 1113.1 |

Error is about ±0.06% at 10,000 iterations, so both gains are far beyond it at both seeds. The paladins have no published raid band below 60; the /sim page's raid set (raid-consumes.json) carries the elixir from level 40.

## Files to regenerate

- data/proto/common.proto, data/proto/ENGINE_SHA, sim/enginever/version.go: follow the engine at the pin (integrator).
- data/pipeline/simproto/common_pb2.py (`python -m pipeline simproto --engine <clone at the pin>`).
- data/builds/1.60.1.70753/simbuffs.json: gains an `elixir_of_holy_power` entry (name and icon inv_potion_60 / 134842) so the guide's consumable table shows the icon. Nightly-owned.
- web/src/data/generated/sim-ids.json and raid-consumes.json are build-time syncs, not committed.
- sim/request TestRotationLadder goldens for paladin-retribution and paladin-protection if the ladder runs the raid preset.
- Committed in S 408d35cde already: sim/request/raid-consumes.json (genraidconsumes) and sim/request/IDS.md (genids). Both need the engine at 46be83bcd or later; at the old pin `TestIDsMarkdownIsCommitted` and the preset's vocabulary check would fail.
- No addon, coach or Data.lua file names raid consumables (grep for the consumable ids finds none).

## Measure at refresh

- Retribution and Protection BiS re-rank, stat weights and the tier list under the new raid preset (the ranker runs the preset).
- Published DPS / threat for both paladin specs.

## For other lanes

- **Platform** (web/src/lib/sim/buffs.ts): add `holy_power_buff: 'battle-elixir'` to `BY_CONSUMES_FIELD`; without it the panel files the elixir with the potions. Also hide `elixir_of_greater_firepower` (now inert, H3) and `lesser_arcane_elixir` (already inert) from the panel, the way `INERT_BUFFS` hides buffs.
- **Casters**: fork tests in sim/mage/mage_test.go, sim/shaman/elemental/elemental_test.go, sim/shaman/enhancement/enhancement_test.go, sim/shaman/warden/warden_test.go and sim/warlock/dps/dps_warlock_test.go name `ElixirOfGreaterFirepower`; it now gives nothing. Switch them to `ElixirOfFirepower` when those goldens are next regenerated.
- **Integrator**: E 46be83bcd changes no class code; it moves no published number by itself. S 408d35cde moves only the two paladin specs (table above).

## Still open

- None for this item. The recipe's drop rate is classic-db's Era figure (single source); Forever may have changed it.
