# Guide accuracy round 23: Shaman

Auditor: shaman (class index, Elemental, Enhancement, Restoration). Audit tree 24ee457d0 (main after guide round 22; sim round 21 not landed: `sim/enginever/version.go` is still d5c708834, the round 21 commits are on other refs only), engine source read-only at `gr23/engine`, ACTIVE client 1.60.1.70790 under `data/builds/1.60.1.70790/`.
Files: `web/src/content/guides/shaman/{index,elemental,enhancement,restoration}.md`, every note in `data/curated/apl/shaman-{elemental,enhancement,restoration}.json`, every rendered row of `addon-data.json` (`rotations/shaman-elemental` 26 bands, `shaman-enhancement` 19, `shaman-restoration` 23), `talents/shaman.json`, `spelltips/shaman.json`, `spellconst/shaman.json` (coefficients, `max_level`), `trainables/shaman.json`, `races.json`, `items/shaman.json` (bonding), `bis/shaman-*.json`, `addon-builds.json`, `data/curated/rank-choices.json`. Read first: round 22 `shaman.md`, ledger Rounds 19 to 22, briefs 1, 2 and 23, sim rounds 17 and 19 `shaman.md`.

## What moved since round 22

- 35241169d (round 22 shaman): `restoration.md:120` (Chain Heal costs "to about 300 and 385 for ranks 2 and 3"), Restoration Chain Heal note ("even after ranks 2 (46) and 3 (54) arrive"), `elemental.md:118` ("back to the top rank at 28 and 29 ... Each lower rank costs far less than your best bolt").
- `addon-data.json` and Data.lua regenerated (note text only). No talent, spelltip, trainable, race, BiS or point-order change for shaman.

## Findings

Each card fix needs `sim/request/apl/shaman-*.apl.json`, `addon-data.json` and Data.lua regenerated (notes only).

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| SHAMAN-R23-1 | CONFIRMED | page prose (rank reason backward on the client numbers) | restoration.md:120 (Levels 20 to 60) | "Your everyday Healing Wave in a dungeon is not your top rank: a cheaper one heals more before your mana runs out, because your gear adds little healing yet." | The reason is the wrong way round. On this client a higher Healing Wave rank has MORE base healing per mana: rank 2 1.58, 3 1.64, 4 1.65, 5 1.71, 6 1.83, 7 2.01, 8 2.21, 9 2.40, 10 2.60 (`spelltips`, heal / cost; Tidal Focus and Healing Way scale all alike). So with little healing on your gear the top rank heals more per mana at every level of the ladder. Gear is what turns it round: the bonus is the same per cast on every rank that still takes it, so it is worth more per mana on a cheap one. Healing power at which the listed rank passes the top rank per mana (`spellconst` coefficients 0.571 / 0.714 / 0.857, `max_level` + 12 rule): 11 at 20, 19 at 25, 18 at 30, 74 at 36, 101 at 37, 224 at 40, 270 at 41, 197 at 42, 440 at 46, 264 at 47, 341 at 50, 390 at 56, 446 at 58, 458 at 59. The rows themselves were measured on the frozen BiS sets, which carry a lot of healing: about 150 at band 20, 255 at 30, 375 at 40, 500 at 50 (`bis/shaman-restoration.json` slots, `healing` stat, plus spell power), against 570 at raid 60. Example at 48: rank 6 484 / 265 = 1.83, rank 8 972 / 440 = 2.21; with no gear a full mana bar of rank 8 heals about 20% more. The measured advice stands; the stated cause does not. Round 18's rule: every rank reason must hold per cast on the client numbers. | "Your everyday Healing Wave in a dungeon is not your top rank once you wear healing gear: that gear adds the same to every rank that still takes its bonus, so a cheaper one heals more before your mana runs out." |
| SHAMAN-R23-2 | CONFIRMED | card note and page prose (instruction contradicts the card and loses a tick) | `shaman-restoration.json` priorityList[6] (Riptide on the tank) note, first sentence, rendered with the talent from 50; restoration.md:71 | "Riptide on the tank whenever it is ready and the tank is hurt" | Riptide heals for 15 s (5 ticks, 3 s apart) on a 12 s cooldown (`spelltips` 408521 / 1239242 / 1239243: "over 15 sec", `cooldown_ms` 12000; `sim/shaman/riptide.go:21-22`). "Whenever it is ready" is every 12 s, and a recast replaces the running one (`riptide.go` `cancelRiptide`: "a new cast replaces whichever rank was there"), which throws away the fifth tick, 20% of what that cast paid for. The card does not do this: line 6's condition is `not dotIsActive(Riptide on the tank)` and tank under 90%, so it recasts only when the last one has run out, and the 25% Chain Heal bonus never lapses that way. The same sentence's "no single-target heal gives more healing for its mana" fails when clipped: 805 x 0.8 / 230 = 2.80 at 60 against Healing Wave rank 10 with Healing Way 3.26. Earlier rounds checked this sentence's comparison, not its timing. | Note: "Riptide on the tank when the tank is hurt and your last Riptide has run out: it is ready again at 12 seconds but heals for 15, so an early recast loses its last tick. No single-target heal gives more healing for its mana, and Chain Heal cast on a target with your Riptide heals it 25% more." restoration.md:71: "Riptide on the tank when the tank is hurt and your last Riptide has run out: it is ready again at 12 seconds but heals for 15, so an early recast loses its last tick. No single-target heal gives more healing for its mana, and it sets up the Chain Heal. It heals over its 15 seconds with nothing up front, so it never answers a big hit on its own." |
| SHAMAN-R23-3 | CONFIRMED | card note (quantity false at most bands) | `shaman-elemental.json` priorityList[5] (Chain Lightning) note, rendered 32 to 60 | "the jumps add most of another bolt on top" | Each jump is 30% less than the last (`spelltips` 421 to 10605), so the jumps are 0.7 and 0.49 of the first hit. Against the top Lightning Bolt the row renders beside (same Concussion, same crit talents), second target only / both jumps: 32 to 37 215 v 148: 1.02 / 1.73 bolts; 38 and 39 215 v 194: 0.78 / 1.32; 40 to 43 348 v 194: 1.26 / 2.13; 44 to 47 348 v 240: 1.02 / 1.73; 48 and 49 428 v 240: 1.25 / 2.12; 50 to 55 428 v 295: 1.02 / 1.73; 56 to 60 510 v 353: 1.01 / 1.72. "Most of another bolt" is true at one pair of levels on two targets only; elsewhere the jumps add at least a whole bolt and up to two. It undersells the spell, so it costs no one a wrong choice, but the quantity is false. | "Chain Lightning on two or more targets: the first hit is bigger than a Lightning Bolt and each jump still deals 70% of the hit before it, for the same cast time." |
| SHAMAN-R23-4 | LIKELY | card notes (rank reason needs its condition) | `shaman-restoration.json` priorityList[17] (Healing Wave) note, third sentence; priorityList[16] (Chain Heal) note, third sentence | [17] "While leveling from 20 a cheaper rank heals more before your mana runs out than your top rank"; [16] "cast rank 1 until 59 ... so it heals more before your mana runs out" | Same arithmetic as R23-1, stated without any gear condition. Healing Wave as in R23-1. Chain Heal base heal per mana: rank 1 251 / 260 = 0.97, rank 2 1.06, rank 3 1.21; rank 1 passes rank 2 only above about 205 healing and rank 3 above about 249 (346 at 59 at 90% share). Both notes are true on the sets they were measured on (sim round 19 S1: +0.3% to +1.2% on BiS) and false on a character without healing gear. `restoration.md:120`'s own Chain Heal clause ("the cheaper heal keeps you casting longer") is true at any gear and needs no change. Not CONFIRMED because the measurement stands for the geared player the card assumes. | [17] "While leveling from 20, once you wear healing gear a cheaper rank heals more before your mana runs out than your top rank:". [16] "... to about 300 for rank 2 and 385 for rank 3, so in healing gear it heals more before your mana runs out." |

Counts: CONFIRMED 3, LIKELY 1, NEEDS-SIM 0, NEEDS-OWNER 0.

## Shoulds (wording, not findings)

- **Enhancement Searing Totem note** "unless the fight ends within 20 seconds: a fresh totem would barely fire": the engine fires every 2.437 s (`fire_totems.go:67`), so 8 bolts with 20 s left, 2.2 damage per mana at 60 against Earth Shock's 2.1 with Shamanistic Focus. The guard is measured; the reason overstates it. Suggested "a fresh totem would get only a few bolts off". `enhancement.md:75` ("unless the boss is about to die") is fine.
- **elemental.md:118** "still keeps most of your spell damage, so it does more in total": the same effect as R23-1 in milder form (the lower rank passes the top per mana above 17 to 95 spell damage, `spellconst`), and the sentence already gives the true cause, so not filed. Suggested "once you have some spell damage on your gear" only if the Restoration line changes and the pages should match.
- Carried from round 22: elemental.md:90 "Ancestral Knowledge makes every point 10% bigger" (raid build only); Elemental Lightning Bolt filler "In a raid cast rank 10" ("in a raid at 60"); enhancement.md:52 "after a spell crits" (direct damage spells only). From round 21: Enhancement Stormstrike "hardest hit" near a tie at 40; Grace of Air note "rank 3 at 60"; elemental.md:74 "about twice". From round 20: elemental.md:66 "you never swing a weapon" against :114; Restoration Healing Wave note "17 levels past it" (one level early); enhancement.md:55 "the two Improved Ghost Wolf points" scoped to the raid build; index.md:28 "2.5 second casts before any talents" from rank 3; Elemental Flame Shock note "with Call of Flame ... 21 seconds" at bands 10 to 22; Elemental Chain Lightning "same cast time" outside Lightning Mastery.

## Checked and right this round

- **Round 22's fixes (new work).**
  - R22-1, restoration.md:120 "247 mana with Tidal Focus to about 300 and 385 for ranks 2 and 3": 260 x 0.95 = 247; 315 x 0.95 = 299.25; 405 x 0.95 = 384.75; Tidal Focus 5/5 from 22, before Chain Heal at 40, so no band sits on the wrong side. "about" now sits on both rounded numbers. OK.
  - R22-2, Chain Heal note: ranks 2 (46) and 3 (54) match `trainables`; the costs are scoped by "after they arrive"; rows cast rank 1 at 40 to 59. The later reason is SHAMAN-R23-4.
  - R22-3, elemental.md:118: "back to the top rank at 28 and 29" matches band 28 (943 on both rows); "Each lower rank" now excludes the top rank. Both sides at each band, Convection on both from 14: 45 v 85 at 26 and 27, 45 v 85 / 110 at 30 to 33, 60 v 110 / 135 at 34 to 38, 85 v 135 / 160 at 39 to 44, 110 v 160 / 190 at 45 to 50, 135 v 190 / 220 at 51 to 56, 160 v 220 at 57 to 59. Share kept: rank 3 at 33 90%; ranks 4 to 7 95% at their last level; rank 8 100%. OK. Totemic Focus and Thundering Strikes change neither side.
- **Side effects of a later talent on either side of a comparison.** Tidal Focus (every Restoration heal from 22), Healing Way (Healing Wave only from 41 to 43, checked in R21-7 and against the Nature's Swiftness note), Purification (all heals), Convection (Elemental from 14; Enhancement leveling 43 to 45 and no cost quoted), Shamanistic Focus (shocks, Enhancement from 24 or 27, no cost quoted), Totemic Focus (Elemental leveling 54 to 58, both sides of Searing against Magma), Trial by Fire (Magma 30 s at 29, 40 s from 30; Fire Nova, only named at its rotation value). No comparison flips.
- **Every superlative and comparison on the four pages** re-swept: index :26 to :36; elemental.md :34, :35, :37, :42, :52, :57, :74, :80, :82 ("cheapest bolt that still takes all of your spell damage at 60": rank 8 `max_level` 49, whole; rank 7 `max_level` 43, 75%), :104, :118; enhancement.md :52, :81, :83, :87, :89, :93, :117, :121; restoration.md :33, :37, :71 (the per-mana half; timing is R23-2), :76, :80 (rank 2 per mana past about 290 healing, rank 3 per cast, rank 1 never better at 60).
- **Every quoted cost, duration and cooldown**: Totemic Call 25% base mana, 30 s; Mana Spring 10 every 2 s for 5 min; Searing 55 s; Flame Shock 21 s; Magma 40 s and Fire Nova 6 s with Trial by Fire (Elemental), Fire Nova 10 s (Enhancement); Ghost Wolf 3 s; Lightning Mastery, Rage of the Farseer and Nature's Swiftness 180 s; Stormstrike 6 s; Maelstrom Weapon 30 s; Mana Tide 12 s, 300 s; Riptide 15 s on 12 s; Water Shield 10 min, 2%; Reincarnation 40 min with the talent; Lightning Bolt 198 / 171; Chain Heal 247 / about 300 / about 385; Shamanistic Focus 45%.
- **Talent bullets** on all three pages against every max-rank text and the decoded strings: Elemental 31 + 5 + 10 / 15 (leveling Tidal Focus 5, Mindfulness 3, Totemic Focus 5, Improved Reincarnation 2; raid Ancestral Knowledge 5 for Totemic Focus); Enhancement leveling 11 / 33 / 7, raid 12 / 31 / 8; Restoration 10 / 41. Prereqs: Lightning Mastery needs Call of Thunder, Elemental Fury needs Lightning Overload 3, Flurry needs Mental Dexterity 3, Riptide needs Healing Way 3, as the orders do.
- **Leveling prose against the shipped orders**: Elemental 10 to 14, 15 to 19, 40; Enhancement 10 to 14, 15 to 19, Mental Dexterity 20 to 22, Elemental Weapons 23 and 24, Stormstrike 25, 26, Shamanistic Focus 27, Flurry to 32, Improved Ghost Wolf 33 and 34; Restoration 10 to 14, Mindfulness 15 and 16, Tidal Focus from 17, Water Shield 20, Tidal Focus 22, Mindfulness 23, Mana Tide 25, Ancestral Healing and Healing Focus by 30, Riptide 50, last ten from 51. Every spell named against `trainables`.
- **Rank lists against every rendered row**: Elemental filler 26 to 60; Restoration Healing Wave 20 to 59; Mana Tide 38 / 48 / 58; Riptide 40 / 50 / 60.
- **Rotation prose against APL conditions**: Restoration tank thresholds 90 / 60 / 30 / 25, members 50 and 25 with mana above 50, Mana Tide 25; Enhancement Earth Shock 25 and 50 gates, Searing 20 s guard; Elemental top bolt on mana percent above fight percent, Fire Nova and Chain Lightning on two or more targets. R23-2 is the one mismatch.
- **Gear prose against both factions' raid rows**: Elemental Frozen Heart and Totem of Thunder the only bind-on-pickup crafts; Enhancement Frozen Heart (both), Molten Heart (Alliance only), leather winning some slots; Restoration Dormant Heart the only bind-on-pickup craft (Totem of Urgency binds when equipped), a staff in the main hand.
- **Races** against `races.json`; **voice**: no em or en dashes, no nightly numbers, no site or sim talk; no page or rendered note names 70755. Source links: frontmatter unchanged since round 22; not refetched.

## Decisions

- **D1. SHAMAN-R23-1 is CONFIRMED and the rows stay.** The rows are measured and stay. The "because" clause is a checkable statement about client numbers and it has the sign wrong. "The cheaper one heals more" remains true in the gear the rows were measured on. `grep` finds the "gear adds little" wording only on this page.
- **D2. SHAMAN-R23-2 is CONFIRMED.** Three hard facts: the client's 15 s against 12 s, the engine's replacing recast, and the card's own condition. The fix matches the card and costs no other line.
- **D3. SHAMAN-R23-3 is CONFIRMED despite understating.** The filing bar is false or misleading; this is false. It is the least harmful of the set.
- **D4. SHAMAN-R23-4 is LIKELY, taken if the integrator takes R23-1,** so the card and the page give the same reason.
- **D5. Not filed, per the brief and earlier rounds**: rows rendering before their talent (Enhancement Lightning Bolt and Fire Nova before Maelstrom Weapon and Elemental Focus; Stormstrike rows from 40; Restoration Mana Tide, Nature's Swiftness and its Healing Wave, Riptide 40 to 48); card actions casting lower ranks (Grace of Air rank 2, Strength of Earth rank 4, Healing Wave rank 9 for the "highest rank" row at 60, Chain Heal rank 3 at 60); the Elemental card's missing Fire totem below 26 on two targets; BiS race captions; the leveling tree's rendering.

## Engine questions

- None new. Coordinator holds Revenge rank 6, the below-level-20 penalty and BloomSpellCoeff; the second touches this class's sub-20 ranks as round 22 said.

## Sims requested

- Optional, wording only, does not block R23-1: Restoration Healing Wave and Chain Heal rows against the top rank at bare 30, 40 and 50 with the set's healing stat at 0 and at half, 10,000 iterations, seeds 1 and 50001, reporting overheal share. Decides whether "once you wear healing gear" should name a rough amount, and how much of the measured win is overhealing rather than mana.
- Carried, already routed: Stormstrike against an Earth Shock with Stormstrike's 20% at 40 and 44; Restoration healing per mana at 60; Enhancement half reserve against a quarter at 37 to 41; Elemental band 20 / 30 talent strings; Enhancement band 30 with Stormstrike; Restoration bands 40, 50 and bare 60 with the Chain Heal rows.

## Sim findings

- None new. Carried: the Elemental card drops no Fire totem below 26 on two or more targets; Stormstrike `RequiredLevel: 40` against the order's point at 25; Maelstrom Weapon proc rate; Lightning Overload rank spells, Healing Wave rank 10 and the book ranks.

## Sources

- Client data 1.60.1.70790 (`data/builds/1.60.1.70790/`): `spelltips/shaman.json`, `spellconst/shaman.json`, `talents/shaman.json`, `trainables/shaman.json`, `races.json`, `items/shaman.json`, `bis/shaman-*.json`, `addon-builds.json`, `addon-data.json`.
- Curated: `data/curated/apl/shaman-*.json`, `data/curated/rank-choices.json`.
- Engine d5c708834: `sim/core/spell.go` (`SpellLevelPenalty`, `LevelPenaltyFromMaxLevel`), `sim/shaman/riptide.go`, `sim/shaman/fire_totems.go:67`.
- Reviews: `design/reviews/2026-11-04-sim-audit-round17/shaman.md` R1, `design/reviews/2026-11-04-sim-audit-round19/shaman.md` S1 and S2, `guide-accuracy/round16/shaman.md` (Chain Heal per mana), `guide-accuracy/round22/shaman.md`, `ledger.md`.
- Low-rank rule, from round 1, not refetched: ForeverChanges, "WoW Forever downranking: the level penalty and which ranks to keep" https://foreverchanges.pro/downranking ; Blizzard, "WoW Forever Beta Development Notes" https://us.forums.blizzard.com/en/wow/t/forever-beta-development-notes/2360696 . No new web source was needed: every call this round is settled by client text.

## Note sweep (every row note, every band it renders)

Bands: Elemental 10 12 14 18 20 22 26 28 30 32 34 36 38 39 40 42 44 45 46 48 50 51 52 56 57 60; Enhancement 10 12 14 20 22 24 26 30 32 36 38 40 42 44 48 50 52 56 60; Restoration 10 12 18 20 24 26 28 30 32 36 37 38 40 42 44 47 48 50 52 56 58 59 60. L = leveling order, R = raid order, T = trainer. List indexes are `rotation.priorityList`.

**Elemental** (orders identical to 45)
- [0] Mana Spring, 26 to 60: T 26; 600 / 40 at rank 1 to 1500 / 100 at rank 4, "many times"; 5 min every rank. OK.
- [1] Lightning Mastery autocast (no band row): L and R 40; 15 s; Berserking after it. OK.
- [2] Flame Shock, 10 to 60: Call of Flame 23 to 25 L and R, conditional (Should); no Earth Shock row. OK.
- [3] Searing Totem, 10 to 60: 20 yards every rank; Magma named "from 26" (T 26); several times its mana at every band. OK.
- [4] Magma Totem, 26 to 60: 20 s every rank; 40 from 30 (Trial by Fire 29, 30). OK.
- [5] Chain Lightning, 32 to 60: first hit above the top bolt at every band; jumps **SHAMAN-R23-3**; same cast time (Should).
- [6] Fire Nova, 12 to 60: Searing T 10, Magma T 26; 10 yards; one-boss clause OK.
- [7] Lightning Bolt top rank, 10 to 60: "at 60" scoped; 198 / 171; +19.7%; Convection from 14. OK.
- [8] Lightning Bolt filler, 10 to 60: list matches every band (548, 943, 548, 915, 943, 6041, 10391, 10392); rank 10 at raid 60. OK.

**Enhancement** (orders as decoded)
- [0] autocast: no note.
- [1] Strength of Earth, 10 to 60: T 10; 30 yards every rank. OK.
- [2] Grace of Air, 42 to 60: T 42, rank 2 56; Windfury T 32, Rockbiter T 1. OK (rank 3 Should).
- [3] Stormstrike, 40 to 60: talent 25; 6 s; 20% (Should "hardest"). OK.
- [4] Lightning Bolt at five stacks, 10 to 60: conditional on all five points (42 L, 39 R); pre-talent rows D5. OK.
- [5] Searing Totem, 10 to 60: 20 s guard matches the APL; "barely fire" (Should). OK.
- [6] Fire Nova on Clearcasting, 12 to 60: Elemental Focus R 52, stated "in the raid build"; 10 yards. OK.
- [7] Earth Shock, 10 to 60: gates match the APL; Stormstrike "once you have" (25); Maelstrom L 37 to 41, R 35 to 38; no cost quoted. OK.

**Restoration** (both orders identical)
- [0] autocast: potions, runes; Nature's Swiftness 44. OK.
- [1] Mana Tide, 10 to 60: talent 25 (D5); 12 s; party; rank 4 T 58; Healing Stream before T 20 inside pre-talent rows (D5). OK.
- [2] Water Shield, 20 to 60: talent 20; three globes; 2%. OK.
- [3] Healing Stream, 20 to 60: every 2 s for 5 min every rank. OK.
- [4] Nature's Swiftness, 10 to 60: talent 44 (D5); text OK.
- [5] Healing Wave after Nature's Swiftness, 10 to 60: top Healing Wave beats a full Chain Heal at every band; rank 9 row at 60 routed. OK.
- [6] Riptide, 40 to 60: talent 50 (D5); per-mana clause holds on an unclipped cast; **SHAMAN-R23-2** ("whenever it is ready").
- [7] Lesser Healing Wave on the tank, 20 to 60: 1.5 s every rank. OK.
- [8] to [15] party rows, 20 to 60 and 40 to 60: no note.
- [16] Chain Heal, 40 to 60: three targets, half a jump, party rule; R22-2's scoping OK; 247 exact, about 300, about 385; "At 60 without raid buffs, cast rank 2" OK; reason **SHAMAN-R23-4**.
- [17] Healing Wave on the tank, 10 to 60: "To save mana at 60 drop to rank 8" (`max_level` 53, whole); ranks 1 to 4 none at 60; list matches every band; reason **SHAMAN-R23-4**; "17 levels past" (Should).

## Counts

CONFIRMED 3 (SHAMAN-R23-1 restoration.md:120 "because your gear adds little healing yet", backward on the client numbers; SHAMAN-R23-2 Riptide "whenever it is ready", a 12 s recast clips a 15 s heal and the card waits; SHAMAN-R23-3 Chain Lightning "most of another bolt", at least one and up to two), LIKELY 1 (SHAMAN-R23-4 Healing Wave and Chain Heal rank reasons with no gear condition), NEEDS-SIM 0, NEEDS-OWNER 0. Round 22's three fixes held.
