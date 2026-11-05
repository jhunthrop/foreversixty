# Guide accuracy round 24: Shaman

Auditor: shaman (class index, Elemental, Enhancement, Restoration). Audit tree ea9adea1c (main after guide round 23, which merged sim round 21; `sim/enginever/version.go` 170a59c39), engine source read-only at `gr24/engine`, active client 1.60.1.70790 (`data/builds/1.60.1.70790/`).
Files: `web/src/content/guides/shaman/{index,elemental,enhancement,restoration}.md`; every note in `data/curated/apl/shaman-{elemental,enhancement,restoration}.json`; every rendered row of `addon-data.json` (`rotations/shaman-elemental` 26 bands, `shaman-enhancement` 19, `shaman-restoration` 23; first sentence is what the addon prints); `web/src/lib/guides/rotation-lists.ts` (`clauseOf`, the level 60 lists' clause); `talents/shaman.json`, `spelltips/shaman.json`, `spellconst/shaman.json`, `trainables/shaman.json`, `addon-builds.json` (both orders decoded, index 0 = level 10). Read first: sim round 21 `2026-11-05-sim-audit-round21/shaman.md` and `integration.md`, round 23 `shaman.md`, ledger Rounds 19 to 23, the round 2 and round 24 briefs.

## What moved since round 22 (new work)

Sim round 21 (5d59b485a, 0a192ab69, 47e610012), merged in 617d4e9c4:
- Elemental card line 3 (Searing Totem): now `(targets == 1 or not known Magma Totem) and Searing not up`; note rewritten, then merged with round 22's wording on main.
- Elemental card line 6 (Fire Nova) note: "(Searing Totem until 26, Magma Totem from then on)".
- elemental.md:118: "On packs keep dropping Searing Totem for Fire Nova until Magma Totem takes its place at 26".
- Enhancement card line 4 (Lightning Bolt): gated on `auraIsKnown(408498 tag 5)`; note adds "once you have all five points" and the held-swing reason.
- Enhancement card line 7 (Earth Shock): 25% gate only; note loses the half reserve.
- enhancement.md:74 (loop 3), :77 (loop 6), :85 (Packs), :87 (Mana, half-reserve sentence removed), :121 (Maelstrom Weapon 37 to 42).

Round 23 (139e28daa, 9577f4bd3): Restoration Riptide note and restoration.md:71; Healing Wave and Chain Heal notes' gear condition; restoration.md:120 gear reason (reviewer's coefficient wording); Elemental Chain Lightning note's 70% jumps.

`addon-data.json` was regenerated after both (369342be0, 632d030e8, b1235b9b1): every new note's first sentence is on the rendered rows; no half-reserve text survives anywhere (grep "half" on enhancement.md hits only Mental Dexterity's "one and a half").

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| SHAMAN-R24-1 | CONFIRMED | card note ("only" claim false on two levels it renders; new in sim round 21) | `shaman-elemental.json` priorityList[3] (Searing Totem), first sentence, the one the addon prints; rendered 10 to 60, the clause bites on band 22 (levels 22 to 25) | "Searing Totem on a single target, and in the pack until Magma Totem at 26, since until then it is the only Fire totem Fire Nova can burst from." | Fire Nova "Instantly inflicts 33 fire damage to enemies within 10 yd of your active Fire totem" (`spelltips` 408341 to 408345): any Fire totem. Frost Resistance Totem is trained at 24 (`trainables`, 8181) and is a Fire totem: client `school_mask` 4 like Searing, Magma and Flametongue (Mana Spring is 16), and its reagent is the Fire Totem (Wowhead Classic, Frost Resistance Totem, "Tools: Fire Totem"). So at 24 and 25 Searing is not the only Fire totem Fire Nova can burst from; it is the only one that does damage. The card's action is right (a Frost Resistance Totem in a pack adds nothing); the stated reason is false on two levels the row shows. Least harmful of the shapes: no one is told to do the wrong thing. | "Searing Totem on a single target, and in the pack until Magma Totem at 26, since until then it is the only Fire totem that does damage for Fire Nova to burst from." (`sim/request/apl/shaman-elemental.apl.json`, `addon-data.json`, Data.lua regenerated, note only.) |

Counts: CONFIRMED 1, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0.

## Shoulds (wording, not findings)

- **Elemental Lightning Bolt filler at band 60** (brief's "level 60 rule missing from the first sentence"): the addon prints "Lightning Bolt for everything else." beside a rank 10 icon; the 60 rule (rank 8 without raid buffs) is in the third sentence. Not false: the first sentence is the row's own default condition and the band 60 row is the raid row, whose rank 10 the page also teaches (:82 "In a raid, cast rank 10 the whole fight"). Suggested first sentence: "Lightning Bolt for everything else: rank 10 in a raid, a cheaper rank without raid buffs." Note `clauseOf` would then print "For everything else" on the guide list, unchanged.
- **restoration.md:120** "that gear's bonus is worth far more for its mana on a cheaper rank": bonus per mana, listed rank against the top rank, coefficient by base cast time (0.571 / 0.714 / 0.857) and the 17-level rule: 2.3x at 20, about 2x at 26 to 36, 1.5x at 41, 1.3x at 47, 1.6x at 50, 1.3x at 53, **1.16x at 55** (rank 6, 23 levels past, 0.857 x 0.7 / 265 against rank 8 0.857 / 440), 1.6x at 58, 1.27x at 59. True in sign everywhere; "far" overstates at 53 to 55. "worth more for its mana" would hold at every level.
- **enhancement.md:74** "it is instant and free then, so it costs you nothing": it still costs a global (1.5 s) that Stormstrike or Earth Shock could have used. "costs you no mana and no swing" is exact.
- **enhancement.md:87** "Earth Shock never needs to wait": settled by sim round 20 N7 on the old 50% Maelstrom chance. The doubled chance only adds free instant bolts, which spend no mana and take globals from Earth Shock, so it should still hold; a recheck is listed under Sims requested, optional.
- Carried, unchanged: Enhancement Searing "would barely fire" (8 bolts in 20 s); elemental.md:90 "Ancestral Knowledge" (raid build only); Elemental filler "In a raid cast rank 10" ("in a raid at 60"); enhancement.md:52 "after a spell crits"; Stormstrike "hardest hit" near a tie at 40; Grace of Air "rank 3 at 60"; elemental.md:74 "about twice"; elemental.md:66 against :114; Healing Wave note "17 levels past" (one level early); enhancement.md:55 scoped to the raid build; index.md:28 "2.5 second casts"; Elemental Flame Shock "with Call of Flame" at 10 to 22; Chain Lightning "same cast time" outside Lightning Mastery; sim round 21's own SHOULD (the addon prints "until Magma Totem at 26" at 60 too, harmless).

## Checked and right this round

- **Sim round 21, Elemental.**
  - Searing condition against the note: the `or` resolves as the report says (Magma Totem rank 1 at 26, `trainables` 8190); on one target nothing changed; on packs 10 to 25 Searing goes down, from 26 Magma takes the pack. The note's "from 26, Magma Totem costs several times its mana and only burns within 8 yards" holds both sides at every band (Magma 230 to 650 against Searing at the band's rank; 8190 / 10587 "within 8 yards"); "bolts reach 20 yards" every rank. The "only" is R24-1.
  - Fire Nova note: trained 12 (408341), 10 yd of the active Fire totem; "Searing Totem until 26, Magma Totem from then on" matches the two rows. "Lightning Bolt is the better cast" on one boss: unchanged, single-target card never casts Fire Nova.
  - elemental.md:118 new clause: Searing 10, Fire Nova 12, Magma 26, all `trainables`; the paragraph still reads in order after the Levels 1 to 19 Fire Nova sentence (:114) and before Chain Lightning 32 and Lightning Mastery 40 (leveling order 40). No pack mana claim is made.
  - Searing at 10 to 25 lasts 30 to 50 s by rank, so "keep dropping" is right; Trial by Fire (29, 30 on both orders) touches only Magma and Fire Nova.
- **Sim round 21, Enhancement.**
  - Maelstrom Weapon text (`talents` 104741): "by 4% / 8% / 12% / 16% / 20% ... Stacks up to 5 times. Lasts 30 sec." Five stacks at 2/5 = 40% off (1.5 s cast), 3/5 60% (1.0 s), 4/5 80% (0.5 s and 20% of the mana): "even five stacks leave the bolt a cast" true at 1 to 4 points; "instant and free" only at 5/5.
  - The reason: the engine stops melee for a cast (`sim/core/cast.go:226,311` `StopMeleeUntil`), so "a cast that holds up your next swing" is the mechanism the engine measured; "the swing is worth more" is the measured result (N2, +1.9% to +3.7% at 38 to 46). The half of the gain that came from the reserve is described by loop 6 and the Mana paragraph, not credited to the swing.
  - Orders: leveling Maelstrom 37, 38, 39, (Rage of the Farseer 40), 41, 42, so "starts at 37 and is full at 42" and "from 42" are exact; raid 35 to 39. The card's `tag 5` gate matches "once you have all five points" on both orders; rows 10 to 36 render before the talent (D5 earlier rounds).
  - Earth Shock: condition `mana > 25%` only; note, loop 6 and the Mana paragraph all say a quarter; Stormstrike "once you have" it (25 on both orders); Stormstrike 8% of base mana (`spelltips` 17364). No half-reserve sentence anywhere on the page, the card or the rendered rows.
  - Packs: "once you have all five points of Maelstrom Weapon keep spending five stacks" matches the card; Fire Nova 10 s (no Trial by Fire in either Enhancement build).
- **Round 23 fixes.**
  - Riptide note and restoration.md:71: 15 s heal, 12 s cooldown on every rank (408521, 1239242, 1239243); the card recasts only when `not dotIsActive` on the tank under 90%; "an early recast loses its last tick" true. Per-mana half unchanged.
  - Healing Wave and Chain Heal notes: "once you wear healing gear" / "in healing gear" scope the measured rows as R23-4 asked. Chain Heal 247 / about 300 / about 385 unchanged.
  - restoration.md:120: the sign is now right on the client numbers (see the Should for "far").
  - Chain Lightning note: "each jump still deals 70% of the hit before it" (`spelltips` 421 to 10605, "Each jump reduces the damage by 30%"); first hit above the top bolt at every band (round 23 table).
- **Brief's shapes, swept on all four pages.**
  - Class index teaching a rotation: index.md teaches none; its Fire Nova (:29), Maelstrom (:31) and Totemic Call 22 / 30 s (:27) lines match the client and the cards.
  - Reasons against numbers: Maelstrom (above); Restoration gear reason (above); Elemental bolt ranks (:82, :118, round 22 tables) unchanged; Enhancement Mana "Stormstrike, which hits harder than the shock you skip" (carried Should).
  - Level 60 rule in the first sentence: every row's first sentence read at band 60 and as `clauseOf` prints it on the level 60 lists. Only the Elemental filler leaves its 60 rank rule to a later sentence (Should); Restoration Chain Heal's first sentence carries "while your mana is ahead of the fight", the rule :80 teaches; Healing Wave's rank 9 at 60 is a card action (routed earlier).
  - "never waits", "every time", "every pull": Elemental :114 "Searing Totem goes down on every pull from 10" (card: one target, or any pack before 26); Enhancement :117 "on every pull that lasts" (the 20 s guard); :87 (Should, sim requested); Restoration :73 "never below rank 8" against a rank 9 row.
  - Leveling "on cooldown" against held rows: Enhancement Stormstrike "press it the moment you learn it" (the card has no hold); Earth Shock "on cooldown" carries its quarter gate wherever it is said; Elemental names no "on cooldown" leveling spell.
- **Superlatives, comparisons, quoted costs, durations and cooldowns**: round 23's sweep re-run against the merged notes; nothing in sim round 21 changed a cost, duration or cooldown the pages quote (Fire Nova 10 s and 6 s with Trial by Fire, Magma 20 / 40 s, Searing 55 s at 60, Earth Shock and Stormstrike 6 s, Maelstrom 30 s, Riptide 15 / 12 s, Totemic Call 30 s, Lightning Bolt 198 / 171, Chain Heal 247 / about 300 / about 385).
- **Races, gear, professions, voice, sources**: unchanged since round 23 (no shaman BiS, race or build commit since 24ee457d0); no em or en dashes; no nightly numbers; no page names 70755.

## Decisions

- **D1. SHAMAN-R24-1 is CONFIRMED, low severity.** The sentence gives a reason, the reason is a checkable client fact, and it is false at 24 and 25. The fix keeps the action and states the true reason (the only damaging Fire totem).
- **D2. The Elemental filler first sentence is a Should, not a finding.** Unlike the warlock Amplify Curse clause, "for everything else" is true of the row at 60 and the rank is a card action; the raid band teaches the raid rank the page itself teaches.
- **D3. Not filed, per the brief and earlier rounds**: rows rendering before their talent (Enhancement Lightning Bolt and Fire Nova, Stormstrike from 40, Restoration Mana Tide, Nature's Swiftness, Riptide 40 to 48); card actions casting lower ranks (Grace of Air 2, Strength of Earth 4, Healing Wave 9 at 60, Chain Heal 3 at 60); BiS race captions; leveling tree rendering; the below-20 healing penalty (routed round 22).

## Engine questions

- None new. Coordinator holds Revenge rank 6, the below-level-20 penalty and BloomSpellCoeff.

## Sims requested

- Optional, wording only (enhancement.md:87 "Earth Shock never needs to wait"): sim round 20 N7 re-run on the doubled Maelstrom chance: Enhancement raid 60, guide raid build, frozen BiS, Horde and Alliance, 10,000 iterations, seeds 1 and 50001, the 25% Earth Shock gate against no gate, plus time under 25% mana. Identical keeps the clause; otherwise drop "so Earth Shock never needs to wait".
- Carried, already routed: Restoration rank rows at healing 0 and half; Stormstrike against an Earth Shock with Stormstrike's 20% at 40 and 44; Restoration healing per mana at 60; Elemental band 20 / 30 talent strings; Enhancement band 30 with Stormstrike; Restoration bands 40, 50 and bare 60 Chain Heal rows.

## Sim findings

- None new. Carried from sim round 21 "Still open": Maelstrom Weapon flat 100% against 20 a minute on a fast weapon (Manual Crowd Pummeler, bands 30 to 39); Lightning Overload rank spells, Healing Wave 2 tooltip, level 60 book ranks.

## Sources

- Client data 1.60.1.70790: `spelltips/shaman.json` (408341 to 408345, 8181, 8190, 10587, 3599, 10438, 421, 10605, 408521, 1239243, 17364), `spellconst/shaman.json` (`school_mask` of 8181, 3599, 8190, 8227, 5675), `talents/shaman.json` (104741), `trainables/shaman.json`, `addon-builds.json`, `addon-data.json`.
- Curated: `data/curated/apl/shaman-*.json`; renderer `web/src/lib/guides/rotation-lists.ts`.
- Engine 170a59c39: `sim/core/cast.go:226,311`, `sim/core/attack.go:864` (`StopMeleeUntil`).
- Reviews: `design/reviews/2026-11-05-sim-audit-round21/shaman.md` (S1, N1, N2), `2026-11-04-sim-audit-round20/shaman.md` N7, `guide-accuracy/round23/shaman.md`, `ledger.md`.
- Web: Wowhead Classic, "Frost Resistance Totem" (Tools: Fire Totem), https://www.wowhead.com/classic/spell=8181/frost-resistance-totem (read 5 Nov 2026). Forever kept the totem (client 8181, trained 24, Fire school).

## Note sweep (every row note, every band it renders)

Bands: Elemental 10 12 14 18 20 22 26 28 30 32 34 36 38 39 40 42 44 45 46 48 50 51 52 56 57 60; Enhancement 10 12 14 20 22 24 26 30 32 36 38 40 42 44 48 50 52 56 60; Restoration 10 12 18 20 24 26 28 30 32 36 37 38 40 42 44 47 48 50 52 56 58 59 60. L = leveling order, R = raid order, T = trainer. Indexes are `rotation.priorityList`.

**Elemental** (orders identical to 45)
- [0] Mana Spring, 26 to 60: T 26; 4 to 10 mana every 2 s for 5 min, cost 40 to 100, "many times" every rank; no-stack clause. OK.
- [1] Lightning Mastery autocast (no band row): L and R 40; 15 s; Berserking after it. OK.
- [2] Flame Shock, 10 to 60: Call of Flame 23 to 25 L and R, conditional (carried Should). OK.
- [3] Searing Totem, 10 to 60: single target / pack until Magma T 26 matches the condition; **SHAMAN-R24-1** ("only Fire totem", Frost Resistance Totem T 24, band 22); 20 yards every rank; Magma "from 26", several times the mana, 8 yards. Rest OK.
- [4] Magma Totem, 26 to 60: 20 s every rank; 40 from 30 (Trial by Fire 29, 30). OK.
- [5] Chain Lightning, 32 to 60: first hit above the top bolt every band; 70% a jump (client). OK.
- [6] Fire Nova, 12 to 60: T 12; Searing T 10 until 26, Magma from 26; 10 yd; one-boss clause. OK.
- [7] Lightning Bolt top rank, 10 to 60: "at 60" scoped; 198 / 171; about 20% more; Convection from 14. OK.
- [8] Lightning Bolt filler, 10 to 60: rank list matches every band; rank 10 in a raid (first-sentence Should at 60). OK.

**Enhancement** (orders as decoded)
- [0] autocast: no note.
- [1] Strength of Earth, 10 to 60: T 10; 30 yards every rank. OK.
- [2] Grace of Air, 42 to 60: T 42, rank 2 56; Windfury Totem T 32, Rockbiter T 1. OK.
- [3] Stormstrike, 40 to 60: talent 25 L and R; 6 s; 20% on the next Lightning Bolt, Chain Lightning or Earth Shock (client 17364). OK.
- [4] Lightning Bolt at five stacks, 10 to 60: gate `tag 5` = 42 L, 39 R; "instant and free" only at 5/5 (client 20% a stack); fewer-points sentence true at 1 to 4 points; Stormstrike's 20% "when you can" (25). Rows before the talent D3. OK.
- [5] Searing Totem, 10 to 60: 20 s guard matches the APL; "barely fire" (carried Should). OK.
- [6] Fire Nova on Clearcasting, 12 to 60: Elemental Focus R 52 only, stated "in the raid build"; 10 yards. OK.
- [7] Earth Shock, 10 to 60: 25% gate matches the APL at every band; Stormstrike "once you have" (25); no cost quoted. OK.

**Restoration** (both orders identical)
- [0] autocast: potions, runes; Nature's Swiftness 44. OK.
- [1] Mana Tide, 10 to 60: talent 25 (D3); 12 s; party; rank 4 T 58. OK.
- [2] Water Shield, 20 to 60: talent 20; three globes; 2%. OK.
- [3] Healing Stream, 20 to 60: every 2 s for 5 min every rank. OK.
- [4] Nature's Swiftness, 10 to 60: talent 44 (D3). OK.
- [5] Healing Wave after Nature's Swiftness, 10 to 60: rank 9 row at 60 routed. OK.
- [6] Riptide, 40 to 60: talent 50 (D3); 12 s / 15 s every rank; recast only when run out (matches `not dotIsActive`); per-mana half on an unclipped cast; 25% Chain Heal. OK.
- [7] Lesser Healing Wave on the tank, 20 to 60: 1.5 s every rank. OK.
- [8] to [15] party rows, 20 to 60 and 40 to 60: no note.
- [16] Chain Heal, 40 to 60: three targets, half a jump; party rule; ranks 2 (46) and 3 (54) T; 247 / about 300 / about 385; "in healing gear" scoped; rank 2 at 60 without raid buffs. OK.
- [17] Healing Wave on the tank, 10 to 60: rank 8 at 60 (`max_level` 53, whole); ranks 1 to 4 none at 60; rank list matches every band; "once you wear healing gear" scoped; "17 levels past" (carried Should). OK.

## Counts

CONFIRMED 1 (SHAMAN-R24-1, Elemental Searing Totem note: "the only Fire totem Fire Nova can burst from" is false at 24 and 25, where Frost Resistance Totem is a Fire totem), LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0. Sim round 21's Elemental and Enhancement card, note and prose changes held apart from that clause; round 23's four fixes held.
