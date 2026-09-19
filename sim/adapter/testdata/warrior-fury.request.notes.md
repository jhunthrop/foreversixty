# warrior-fury.request.json: where the gear came from

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
| head | 12640 Lionheart Helm | kept | this id has a row in Forever's build |
| neck | 18404 Onyxia Tooth Pendant | kept | this id has a row in Forever's build |
| shoulder | 12927 Truestrike Shoulders | **15055 Volcanic Shoulders** | absent from Forever's build; nearest by stat vector in this slot (distance 38.5) |
| back | 18541 Puissant Cape | **20068 Deathguard's Cloak** | absent from Forever's build; nearest by stat vector in this slot (distance 41.0) |
| chest | 11726 Savage Gladiator Chain | **13090 Breastplate of the Chosen** | absent from Forever's build; nearest by stat vector in this slot (distance 20.0) |
| wrist | 19146 Wristguards of Stability | **275739 Wild Wristguards** | absent from Forever's build; nearest by stat vector in this slot (distance 16.5) |
| hands | 14551 Edgemaster's Handguards | **281280 Mountain Climbers** | absent from Forever's build; nearest by stat vector in this slot (distance 18.0) |
| waist | 19137 Onslaught Girdle | **20252 90 Green Warrior Waistband** | absent from Forever's build; nearest by stat vector in this slot (distance 26.7) |
| legs | 14554 Cloudkeeper Legplates | **22000 Legplates of Heroism** | absent from Forever's build; nearest by stat vector in this slot (distance 20.1) |
| feet | 14616 Bloodmail Boots | **20050 Highlander's Chain Greaves** | absent from Forever's build; nearest by stat vector in this slot (distance 14.9) |
| finger1 | 17063 Band of Accuria | **20307 90 Green Rogue Ring** | absent from Forever's build; nearest by stat vector in this slot (distance 3.0) |
| finger2 | 18821 Quick Strike Ring | kept | this id has a row in Forever's build |
| trinket1 | 11815 Hand of Justice | **276337 Thaelemaches' Talisman** | absent from Forever's build; nearest by stat vector in this slot (distance 21.8) |
| trinket2 | 13965 Blackhand's Breadth | **2820 Nifty Stopwatch** | absent from Forever's build; nearest by stat vector in this slot (distance 2.0) |
| main_hand | 17075 Vis'kag the Bloodletter | **21521 Runesword of the Red** | absent from Forever's build; nearest by stat vector in this slot (distance 29.2) |
| off_hand | 18832 Brutality Blade | **19168 Blackguard** | absent from Forever's build; nearest by stat vector in this slot (distance 23.1) |
| ranged | 17069 Striker's Mark | **272594 Premier High Warlord's Recurve** | absent from Forever's build; nearest by stat vector in this slot (distance 36.4) |

## What changed when the build's hand types were fixed

Build `1.60.1.69893` first emitted every one-handed weapon as
`HandTypeMainHand`, so the only rows that could go in an off hand were six
`HandTypeOffHand` fist weapons and this table had to take one of them:
`272598 Premier High Warlord's Left Claw`, at a stat distance of 53.0. The build
was regenerated with InventoryType 13 mapped to `HandTypeOneHand`, so the off
hand can now hold a real one-hander and the nearest sword wins instead:
**19168 Blackguard**, distance 23.1. Nothing else in this table moved.


## What the item-ratings fix did, and did not, change

The database this fixture resolves against was regenerated when the data lane
found that an item's hit, crit, dodge, parry and block were being written as
the client's rating POINTS where the engine reads percentages (main `8aaeddc`;
hit/10, crit/14, dodge/12, parry/15, block/5, the factors read from
`gametables/combatratings.txt`). 702 item rows changed value.

No id in the table above moved: the gear set is the same set of items, so no
slot was re-pointed and nothing here needed rewriting. The DISTANCES quoted in
the table are the ones computed at the time each slot was chosen, against the
pre-fix stat vectors, and they have not been recomputed - the nearest row to a
missing item could in principle be a different item under the corrected
values. Re-point these at a curated Forever gear set when the data lane
publishes one, and the distances go with it.
