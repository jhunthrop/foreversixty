# Sim-accuracy round 15: shaman lane

Specs: Elemental, Enhancement, Restoration. Client 1.60.1.70755. Site branch wt/r15-shaman (off int/r15 fd320b60e), fork branch lane/r15-shaman (off int/r15 d1d9f2775).

8 confirmed bugs, all fixed: Elemental 3, Enhancement 2, Restoration 3. By origin: R14 1, old 1 (also a guides handoff), new-data / 70755 6 (rank rows and rotation lines judged on 70753 numbers that 70755 changed under them). Two more rows go to the owner (O1, O2).

## Seed items

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| 1 | **Restoration: the Healing Wave cast after Nature's Swiftness was rank 2.** Nature's Swiftness only takes away the cast time, so the save should be your biggest Healing Wave. Round 14's rank-2 sweep moved this line to rank 2 along with the tank filler, because both lines named the same spell id. It now casts the top trainer rank at every band: rank 9 (10396) at 60, and rank 10 stays held as an AQ book rank (book-ranks.json, pending owner). | cf5fe63bb changed both 332 lines. Raid 60, effective HPS: 718.8 to 726.1 (horde, seed 1), 719.0 to 726.1 (horde, seed 50001), 719.8 to 727.0 / 727.2 (alliance): +1.0%. Bare 40 to 60 unchanged, because those fights are out of mana before Nature's Swiftness is spent. | CONFIRMED | R14 cf5fe63bb, guides-handoff | S cae4e3bb4, E f27227622 |
| 2 | Low-rank heal ranks per band. Every rank that the rank rewrite casts at bands 10 to 60 is learned at that level (trainables/shaman.json 70755: Healing Wave, Lesser Healing Wave, Chain Heal, Riptide, Healing Stream, Mana Tide). Healing Wave rank 2 as the raid and 40 to 59 filler is **pending owner** (the Healing Wave 2 tooltip, and the priest lane's new penalty curve). Riptide's bare-60 row is re-judged in R3 below. | ranks-by-band dump (RewriteRotationRanks at 10 to 60, both presets) | Aligned except R3 | - | - |
| 3 | Enhancement's Air totem. For the shaman's own damage, Grace of Air is the only Air totem that reaches you: Windfury Weapon turns off any Windfury Totem for you (buffs.go weaponTotemReaches). The raid preset already gives Windfury Totem from another shaman. Grace of Air vs Windfury Totem in the Air slot: 794.5 vs 770.0 (-3.1%). The raid-value question (your melee group's Windfury) cannot be simulated. See O1: the weapon imbue question this led to. | spec-breakdown raid 60, horde, seed 1, 10,000 iterations. Web: wowtbc.gg/warcraftforever Enhancement guide (undated, read 2026-11-03) says "either Windfury Totem or Grace of Air Totem" | Aligned | - | - |
| 4 | **Stormstrike at 25, not 40.** Stormstrike sits in the 15-point row (tier 3 since 70532). Its spell row says level 40, but the talent teaches it whenever the row is open. The engine's RequiredLevel 40 (stormstrike.go:41) never gates a cast: core uses RequiredLevel only for the spell-level penalty, and Stormstrike is physical. So the engine needs no fix. The bug was in the band build: the plain walk gave a level-30 shaman Flurry 4/5 and no Stormstrike. A ladder-first row now holds Stormstrike as soon as it fits, then Rage of the Farseer so band 40 keeps its point. | Beta logs (foreverlogs.gg) show Enhancement shamans of level 25 to 35 casting Stormstrike 17364, in reports 2712 (char 101264, 16 points, Earth Shock rank 4 / Healing Stream rank 1, so level 25 to 29), 2717, 2727, 2753, 2756, 2805, 2808 and 6 more. Band 30, -talents 055030310040 vs 055030310031: 77.4 to 97.2 (horde, seeds 1 / 50001), 77.5 to 96.7 (alliance): **+24.8 to +25.6%**. Bands 20 and 40 to 60 unchanged (walk compared at every 5 levels). | CONFIRMED | old (round 13's walk lane did not reach it), guides-handoff | S 03b4e3e5d |
| 5 | Guide round 1 sim rows: S1 is item 4. S2 (carried BiS files: talent strings, weights, Enhancement presets.raid `elemental_sharpening_stone` vs raid-consumes `windfury_weapon`) goes to Measure at refresh. S3 / SHAMAN-2 (Chain Heal jumps): the engine already jumps to the whole raid (chain_heal.go:80 mostInjuredMember walks Env.Raid.AllPlayerUnits). SHAMAN-11 (does rank 10 alone run dry?): no. At raid 60, rank 10 alone runs out of mana for 0.7 s of 180 and does **more** damage; that is E1. Restoration Grace of Air note: item 3. Healing Wave rank 2 filler: pending owner. Lightning Overload rank spells: pending owner. | see the rows named | see rows | - | - |

## Elemental

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| E1 | **The raid list dropped to Lightning Bolt rank 9 when mana fell behind the fight clock, and that now costs damage.** 70755 raised rank 10 to 353 base against rank 9's 295: 20% more damage for 16% more mana (it was 196 vs 178 on 70753). A raid-60 rank row moves the rank-9 line to rank 10. The bare bands keep their own rows: multiplier variants of the pacing condition lost 1.1% at bare 40. | spellconst 15207/15208 70753 vs 70755. Raid 60: 630.4 to 636.9 (horde, seed 1), 630.6 to 638.0 (seed 50001), 622.8 to 629.8 and 623.7 to 630.9 (alliance): +1.0 to +1.2%, out of mana 0.7 to 0.9 s. | CONFIRMED | new-data (70755) | S d214ab51c |
| E2 | **The bare-60 row held the mana-ahead bolt line at rank 7**, judged on 70753's numbers (rank 7 147 base, rank 10 196; now 194 and 353). Rank 10 with the rank-4 row wins, so the row is dropped. Bare 50 ties (182.2 vs 182.3) and keeps rank 7. | Bare 60: 234.1 to 234.9 / 234.2 to 235.1 (horde), 241.6 to 242.4 / 242.5 (alliance): +0.33 to +0.38%, beyond error (±0.2). | CONFIRMED | new-data (70755) | S 8b37ef55a |
| E3 | **Flame Shock rank rows (rank 1 at 30 to 49, rank 2 at 50 to 59) lose on 70755**, which raised every rank (rank 3 now 95 + 24 a tick, rank 6 292 + 80 a tick). The new rows: top rank learned at 30 and 50, rank 3 at 40. | Bare 30: 51.9 to 54.1 (+3.5 to +4.2%). Bare 40: 112.3 to 113.2 (+0.8 to +0.9%; rank 4 measured 113.0, below rank 3). Bare 50: 181.9 to 184.6 (+1.4 to +1.7%). Both factions, seeds 1 and 50001. The Lightning Bolt rows at 30 to 50 were re-swept on the new Flame Shock and hold (rank 3/5 and 4/6/8 variants all lower). | CONFIRMED | new-data (70755) | S e6c9cfeba |
| E4 | Single-target rotation on 70755's buffed shocks and Chain Lightning. Earth Shock while mana is ahead: -1.2%. Chain Lightning while mana is ahead: -1.9%. Both: -2.8%. Chain Lightning on cooldown: -9.7% (the raid fight is mana-bound). | raid 60, horde, seed 1, 3000 iterations | Aligned | - | - |
| E5 | Talents against 70755. Values checked: Convection 10%, Concussion 5%, Elemental Alacrity 0.5 s, Call of Flame, Elemental Focus, Call of Thunder 5% (talents.go:295, was 3), Lightning Overload 10/20/30% (talents.go:644), Elemental Fury, Lightning Mastery, Tidal Focus, Mindfulness, Totemic Focus. Every shaman trait-curve change from 70753 to 70755 (curvediff: Flurry, Lightning Overload, Mental Dexterity) is followed. The spellconst diff touches only Lightning Bolt, Chain Lightning, the shocks, Call of Thunder, Stormstrike, Elemental Reach, Riptide and Totemic Call, and each is in the engine. Elemental Devastation 1 in the build does nothing for a caster, but it is the toll to Elemental Fury, and no point in tiers 0 to 4 does anything in this sim (Reverberation, Eye of the Storm, Reach, Trial by Fire, Earthbound), so there is no legal replacement that "does something". | talents/shaman.json and spellconst 70753 vs 70755 | Aligned | - | - |

## Enhancement

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| N1 | Stormstrike at band 30: seed item 4. | | CONFIRMED | old | S 03b4e3e5d |
| N2 | **The leveling build's 3 points in Elemental Alacrity did nothing.** Alacrity shortens Lightning Bolt and Chain Lightning casts. The rotation casts Lightning Bolt only at 5 Maelstrom Weapon stacks, when it is already instant, and never casts Chain Lightning. Bare 60 is identical with the points removed (299.3 vs 299.3). They move to Elemental Devastation (9% melee crit for 10 s after a spell crit). Change is the frontmatter `build:` line only. | Bare 60: 298.9 to 302.4 / 298.7 to 302.3 (horde), 299.6 to 303.0 / 299.1 to 302.8 (alliance): +1.1 to +1.2%. Band 50 (walk 3200300 to 3230000): 213.3 to 215.4 / 213.2 to 215.6, 213.8 to 216.0 (alliance): +1.0 to +1.1%. | CONFIRMED | old (owner rule: a point that does nothing) | S 485f1fe9a |
| N3 | Rotation. Earth Shock without its 50% mana gate, or with a 25% gate: tie (0.0%). Lightning Bolt at 4 stacks: -0.9%. Lightning Bolt above Stormstrike: +0.06%, a tie. Stormstrike 150% of normalized weapon damage on a 6 s cooldown, Flurry 30%, Mental Dexterity 150%, Mental Quickness 30%, Maelstrom Weapon 20% a stack, Elemental Weapons: engine matches 70755. | raid 60, horde, seed 1 | Aligned | - | - |
| N4 | The raid talents (Concussion, Elemental Devastation, Reverberation, Elemental Focus, Tidal Focus, Mindfulness) all act in the rotation. | | Aligned | - | - |

## Restoration

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| R1 | Nature's Swiftness into rank 2: seed item 1. | | CONFIRMED | R14 | S cae4e3bb4, E f27227622 |
| R2 | **The tank's Riptide came after the tank's Lesser Healing Wave and Healing Wave lines.** That order dates from when Riptide was a 385-mana direct heal. On 70755 it is a 230-mana heal over time: about 4.6 healed per mana, against Lesser Healing Wave's 3.6. The tank's Riptide now comes right after the Nature's Swiftness save. | Raid 60: 726.1 to 731.8 / 732.3 (horde), 727.0 / 727.2 to 733.0 / 733.1 (alliance): +0.8%; mana lasts 139 to 160 s. Bare 50 +1.4%, bare 60 +0.8%, bare 40 unchanged (no Riptide in that build). | CONFIRMED | new-data (70755) | S d289877f9, E 65a4cacff |
| R3 | **Riptide's bare-60 row (rank 2) was judged while Riptide was a direct heal.** Rank 3 now wins, so the row is dropped. | Bare 60: 414.8 to 421.1 / 414.6 to 421.1 (horde), 420.0 to 426.3 / 420.0 to 426.2 (alliance): +1.5%. Rank 1: 410.2. | CONFIRMED | new-data (70755) | S 6e8fc04df |
| R4 | Chain Heal rank 3 vs ranks 1 and 2 at raid 60: rank 3 is best (rank 1 -1.9%, rank 2 -1.3%). Lesser Healing Wave rank 6 vs ranks 5, 4 and 3: rank 6 is best (rank 5 -4.8%). Riptide engine (riptide.go) matches 70755: dummy effect 0, ticks 89/115/161 at 0.1 coefficient, 12 s, 150/180/230 mana. Riptide unconditional on the tank, or member Riptides without the mana gate: ties. Chain Heal targets the whole raid. | raid 60, horde, seed 1 | Aligned | - | - |

## Cross-cutting

| # | finding | evidence and source | confirmed? | origin | fix sha |
|---|---|---|---|---|---|
| X1 | api/internal/coach TestWebHistory_PagesAreTheRoutesShape and TestWebFixtures_NoDrift fail on the baseline (site-before) too ("Battle Shout up 100%..."). Not from this lane. | go test ./internal/coach in $R/site-before | pre-existing | - | - |

## Before/after (raid preset, band 60, 10,000 iterations)

| Spec | Faction | Seed 1 before to after | % seed 1 / 50001 |
|---|---|---|---|
| Elemental | alliance | 622.8 to 630.2 | +1.19 / +1.22 |
| Elemental | horde | 630.4 to 637.3 | +1.09 / +1.21 |
| Enhancement (raid build) | both | 794.6 / 794.5, unchanged | 0.00 / 0.00 |
| Restoration (effective HPS) | alliance | 719.8 to 733.0 | +1.83 / +1.85 |
| Restoration (effective HPS) | horde | 718.8 to 731.8 | +1.81 / +1.85 |

Below 60: Enhancement band 30 +25%, band 50 +1.0 to +1.1%, bare 60 +1.1 to +1.2%. Elemental bare 30 +3.5 to +4.2%, 40 +0.8 to +0.9%, 50 +1.4 to +1.7%, 60 +0.3 to +0.4%. Restoration bare 50 +1.4%, bare 60 +2.3% (Riptide row plus order). No commit moves a spec outside the shaman class.

## Files to regenerate

- sim/request/testdata/ladder/shaman-elemental.golden.md, shaman-enhancement.golden.md, shaman-restoration.golden.md (TestRotationLadder fails on all three, explained by the rank rows, the Stormstrike reservation, the build line and the Riptide order).
- addon-data / addon-builds (rank rows, band builds), coach starters if they carry the shaman builds or ranks.
- Hand-synced, not regenerated: sim/request/apl/shaman-restoration.apl.json, the fork's ui/restoration_shaman/apls/forever_restoration.apl.json, sim/rankchoices/rank-choices.json and sim/leveling/ladder-first.json (copies of the curated files, held by their embed tests).

## Measure at refresh

- Enhancement band 30 (Stormstrike), band 50 and bare 60 (Elemental Devastation build); all Enhancement band strings (the published ones are illegal on 70755).
- Elemental bare 30 to 60 (Flame Shock and Lightning Bolt rows), raid 60 weights with rank 10.
- Restoration at every band with Nature's Swiftness rank 9 and Riptide first; bare 60 Riptide rank 3.
- Guide round 1 S2: Enhancement presets.raid lists elemental_sharpening_stone, while raid-consumes.json has windfury_weapon. Re-read all three statPriority lists after the refresh.

## For other lanes

- Priest lane (owns the new low-rank penalty curve for every healer): shaman rank edits this round, to avoid collision.
  - rank-choices.json: Elemental 15207 raid-60 row (rank 10) added; Elemental 15208 bare-60 row (rank 7) removed; Elemental Flame Shock rows now only 40 to 49 rank 3; Restoration Riptide bare-60 row removed.
  - Restoration APL: the Nature's Swiftness Healing Wave authored at 10396 (rank 9).
  - The Healing Wave 332 rows (12 to 39) are untouched.
  - The rank-2 filler is left for the tooltip ruling.
- Coach lane: X1.

## For the guide lane

Sentences are quoted from int/r15; land/guides-r1 may already have reworded some of them.

- enhancement.md:52. Was: "The leveling build puts its Elemental points into Convection, Concussion, Elemental Alacrity and Call of Flame for cheaper, faster, harder Lightning Bolts and a stronger Searing Totem." Replace with: "The leveling build puts its Elemental points into Convection, Concussion, Elemental Devastation and Call of Flame: cheaper shocks, harder bolts and shocks, 9% crit on your swings for 10 seconds after a spell crits, and a stronger Searing Totem."
- enhancement.md:113 (the Stormstrike leveling line; guides-r1 says "a level 40 spell, so take it at 40"). Replace with: "Take Stormstrike as soon as the fourth row opens at 25; it is your hardest hit from then on."
- elemental.md:80. "Even with raid buffs, a three minute fight of nothing but rank 10 runs you dry before the boss dies. So pace it: ..." is now wrong for a raid. Replace the pacing sentences with: "In a raid, cast rank 10 the whole fight: it does 20% more than rank 9 for 16% more mana, and raid buffs carry it to the end. Without raid buffs, pace it: cast rank 10 while your mana bar is fuller than the share of the fight left, and drop to rank 9 the moment it is not."
- APL notes, data/curated/apl/shaman-elemental.json:
  - Priority 7: "...drop to rank 9 as soon as you fall behind." Append: "In a raid you never need to."
  - Priority 8: "Lightning Bolt rank 9 for everything else: the cheaper rank, so a full raid fight never runs you dry." Replace with: "Lightning Bolt rank 9 when your mana is behind the fight, out of a raid: the cheaper rank, so a long fight never runs you dry. In a raid, rank 10 here too."
- APL note, shaman-restoration.json priority 5 (the Nature's Swiftness Healing Wave). Was: "Rank 2 is cheap, and it still carries your whole healing bonus." Replace with: "Cast your highest rank: Nature's Swiftness makes any rank instant, so this is your biggest heal at no cast time."
- restoration.md:77 already says Nature's Swiftness into Healing Wave; with the guides-r1 text ("your highest rank") it is right.

## Still open (owner)

- **O1. Enhancement weapon imbue.** In Forever, Rockbiter rank 7 raises melee attack power by 653 at 60 (client 16313: 554 + 16.5 a level; Wowhead shows 686, probably at level 62). Under the raid preset, which already gives Windfury Totem from another shaman, Rockbiter plus Grace of Air does 830.8 / 830.2 against Windfury Weapon's 794.5 / 796.7: **+4.5%**. If the Enhancement shaman drops the group's Windfury Totem himself and uses Rockbiter: 801.7 / 801.6, +0.9%, and the group gets Windfury but loses Grace of Air. Bare 60 (no totem): Windfury Weapon wins, 302.5 vs 298.7.
  - Web sources give no Forever consensus. zockify and gfinity (2026, undated) say Windfury at 60. wowtbc.gg's Forever guide says "either Windfury Totem or Grace of Air". Blizzard's Forever notes give imbues a 60-minute duration.
  - This is a raid-composition call: does the Enhancement shaman get another shaman's Windfury Totem? Not changed.
- **O2. Elemental Devastation 1 in the Elemental builds** does nothing for a caster, but it is the only legal toll and no alternative acts in the sim (E5). It is kept.
- Pending, not reopened: Healing Wave 2 tooltip (the filler), the low-rank penalty curve (priest lane), Lightning Overload rank spells, level-60 book ranks (Healing Wave 10, Grace of Air 3).

Sources: [wowtbc.gg Forever Enhancement guide](https://wowtbc.gg/warcraftforever/class-guides/enhancement-shaman/), [zockify Forever shaman](https://www.zockify.com/forever/shaman/), [gfinity Forever shaman guide](https://www.gfinityesports.com/article/world-of-warcraft-forever-shaman-class-guide-specs-builds), [Wowhead Forever Rockbiter Weapon 16316](https://nether.wowhead.com/forever/tooltip/spell/16316) (all read 2026-11-03); foreverlogs.gg reports 2701 to 2808 (beta).
