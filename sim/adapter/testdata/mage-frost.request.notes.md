# mage-frost.request.json: where the gear came from

Generated, then committed. The request beside this file is a plain
api.SimRequest with no extra keys - the envelope crosses lane boundaries and
carries no field that means nothing to the product - so the provenance of each
slot lives here instead.

The starting point is the engine's own phase-one preset for this class, which is
what sim/internal/genfixture used before sim/cmd/forever-sim replaced it. Forever
re-itemises, so most of those ids have no row in the active build
(web/src/data/active-build.json, embedded by `make simdb`). Each one that does
not is replaced by the nearest item in the same slot by Euclidean distance over
the stat vector, with weapon DPS folded in, restricted to items the class may
wear and - for a weapon - to hands the item can actually go in, because
core.EquipItem routes a weapon by its HandType rather than by the slot index.

Re-point these at a curated Forever gear set when the data lane publishes one.

| slot | engine preset | Forever | why |
|---|---|---|---|
| head | 16795 Arcanist Crown | **23263 Champion's Silk Cowl** | absent from Forever's build; nearest by stat vector in this slot (distance 20.2) |
| neck | 18814 Choker of the Fire Lord | kept | this id has a row in Forever's build |
| shoulder | 11782 Boreal Mantle | **20357 63 Green Frost Mantle** | absent from Forever's build; nearest by stat vector in this slot (distance 30.0) |
| back | 13386 Archivist Cape | **10249 Master's Cloak** | absent from Forever's build; nearest by stat vector in this slot (distance 5.7) |
| chest | 14152 Robe of the Archmage | kept | this id has a row in Forever's build |
| wrist | 16799 Arcanist Bindings | **22063 Sorcerer's Bindings** | absent from Forever's build; nearest by stat vector in this slot (distance 7.1) |
| hands | 13253 Hands of Power | **20716 Sandworm Skin Gloves** | absent from Forever's build; nearest by stat vector in this slot (distance 10.5) |
| waist | 19136 Mana Igniting Cord | **18405 Belt of the Archmage** | absent from Forever's build; nearest by stat vector in this slot (distance 15.7) |
| legs | 16915 Netherwind Pants | **21346 Enigma Leggings** | absent from Forever's build; nearest by stat vector in this slot (distance 21.9) |
| feet | 16800 Arcanist Boots | **22084 Virtuous Sandals** | absent from Forever's build; nearest by stat vector in this slot (distance 13.8) |
| finger1 | 19147 Ring of Spell Power | kept | this id has a row in Forever's build |
| finger2 | 19147 Ring of Spell Power | kept | this id has a row in Forever's build |
| trinket1 | 18820 Talisman of Ephemeral Power | kept | this id has a row in Forever's build |
| trinket2 | 12930 Briarwood Reed | **272438 Weakness Analyzer** | absent from Forever's build; nearest by stat vector in this slot (distance 7.0) |
| main_hand | 17103 Azuresong Mageblade | **17015 Dark Iron Reaver** | absent from Forever's build; nearest by stat vector in this slot (distance 42.4) |
| off_hand | 10796 Drakestone | **281720 Restored Chopper** | absent from Forever's build; nearest by stat vector in this slot (distance 13.3) |
| ranged | 15283 Lunar Wand | **20363 63 Green Frost Wand** | absent from Forever's build; nearest by stat vector in this slot (distance 11.6) |

## What changed when the build's hand types were fixed

Build `1.60.1.69893` first emitted every one-handed weapon as
`HandTypeMainHand`, so the off hand could only hold one of six
`HandTypeOffHand` fist weapons and this table took `11863 White Bone Shredder`
at a stat distance of 310.9. With InventoryType 13 mapped to
`HandTypeOneHand` the nearest row is **281720 Restored Chopper**, distance 13.3.
A frost mage swings neither, and the fixture's numbers did not move; only the
order of the zero-damage rows the engine reports changed with it. Nothing else
in this table moved.
