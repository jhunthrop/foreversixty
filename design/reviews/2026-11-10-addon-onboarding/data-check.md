# Round-3 data items, checked against the shipped data (build 1.60.1.70790)

## Talent names on the last page

The tour shows the spec lines from `L.specPlaystyle` on main.

| Name | Where it is | Verdict |
|---|---|---|
| Sniper Shot | Marksmanship talent, tier 7, column 2, node 104997 (addon/ForeverSixty/Data.lua, hunter Marksmanship tab). Also in data/builds/1.60.1.70790/talents/hunter.json and the Marksmanship rotation rows (spell 1310687 / 1310785). | Real. Kept. |
| Barrage | Marksmanship tier 5, node 105001. | Real. Kept. |
| Lacerating Strikes | Survival talent, tier 7, column 2, node 104984 (Data.lua, hunter Survival tab). Also in talents/hunter.json. | Real. Kept. |

## Impaling Harpoon as the Beast Mastery band-20 main hand

The shipped BiS lives in `addon/ForeverSixty/Data.lua`. ForeverSixtyData carries only the nightly ratings and guild standings, with no BiS.

| Check | Source | Result |
|---|---|---|
| Pick | Data.lua `bis["hunter-beast-mastery"][20].alliance.main_hand` | `5200` |
| Pick, BiS run | data/builds/1.60.1.70790/bis/hunter-beast-mastery.json, band 20 bare alliance | main_hand item 5200, `verified: true` |
| Slot | data/builds/1.60.1.70790/items/hunter.json | main_hand, two-hand polearm (inventory type 17), rare (quality 3), item level 22, binds on pickup |
| Required level | items/hunter.json (client source) | 17 |
| Source | loot.json, `dungeon:the-deadmines`, boss Captain Greenskin (npc 647) | 30% chance |

The tour reads "Best in slot at level 20" for a level 6 character. Band 20 is the first BiS band, and the item can be equipped from 17, so the line is correct.

## No hand-typed item in the tour

`TourExample` reads its item from the shipped Data at runtime. It takes the character's first panel upgrade, or failing that the Data BiS pick for the spec it follows (or its first path spec) at its band or the band ahead. The nightly refresh therefore redraws it.

There is deliberately no hard-coded fallback item. When Data has no pick, the row shows no example. A fixed item would quote a nightly figure (the drop chance) and go stale, which feedback-prose-never-quotes-nightly-numbers.md rules out.

The mock now reads its example from the BiS json as well (`example_card` in design/mocks/addon_ux.py). The page 2 capture renders byte-identical to the reviewed round-3 board, so the visible text is unchanged.

A new spec, `onboarding_spec.lua` ("reads a named main hand from the shipped data"), pins only the shape of that read: an item id, its name, and "upcoming". It never pins the item.
