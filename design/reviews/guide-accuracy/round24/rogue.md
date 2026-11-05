# Guide accuracy round 24: Rogue

Audited on the round 24 audit tree (main ea9adea1c: guide round 23 landed with sim round 21 merged; engine pin 170a59c39 per `sim/enginever/version.go`, engine source gr24/engine), against the active client 1.60.1.70790 (`data/builds/1.60.1.70790/`). Files: `web/src/content/guides/rogue/{index,assassination,combat,subtlety}.md`, every line; every note in `data/curated/apl/rogue-{assassination,combat,subtlety}.json`, every sentence, at every band it renders in `addon-data.json` (`rotations/rogue-assassination` 23 bands, `rogue-combat` 22, `rogue-subtlety` 22), against both shipped orders in `addon-builds.json` (unchanged since round 23: `git diff 24ee457d0 HEAD` is empty for it) and the trainer levels in `trainables/rogue.json`. Read first: round 23 rogue report, the ledger's Rounds 19 to 23, `design/reviews/2026-11-05-sim-audit-round21/rogue.md` and `integration.md`, round3/race-sims.md.

## New work: sim round 21 (1039da44a, b0e54a7a6; regen 369342be0) and the round 23 merge

`git diff 24ee457d0 HEAD` on the rogue files: Subtlety card +115 lines (Shank condition and note, Eviscerate and Slice and Dice conditions and notes, Ghostly Strike note from round 23), Combat Adrenaline Rush note (round 23), and prose on all four pages.

| change | where | verdict |
|---|---|---|
| Shank held at four or five combo points only with 50 energy or more (`cp <= 3 or energy < 50`) | Subtlety card row 1 and note; subtlety.md:76; index.md:30 | Card note OK at every band 24 to 60 (Boot Knife named with its 35 and 5/5 by 39, both on the one order; 15 back at rank 1, 25 at rank 2: "pays back more than it costs" at both). subtlety.md:76 and index.md:30 match the condition. The leveling paragraph was not brought along: **ROGUE-R24-1**. Wording: S1, S2 |
| Eviscerate and Slice and Dice before Rupture (`dotIsActive(Rupture) or not spellIsKnown(Rupture)`; Eviscerate also needs Slice and Dice up) | Subtlety card rows 6 and 7 and notes; new subtlety.md:127 | Conditions match the notes at 10 to 19 (Rupture 1943 at 20; ranked `spellIsKnown` resolves to the learned rank, so the branch closes at 20) and from 20. The new 1-19 line matches the card; S3 for its missing "from 10" |
| Assassination poison hands "From 30 to 59" | assassination.md:88 | OK: Deadly trains at 30 (2823); K3's figures (+2.9% at 30 and 34, +3.5% at 38, +0.6% at 44, +0.5% at 48, tie at 54 and 58) give "at 30 to 39 Backstab stacks Deadly faster" and "the gap closes until ... even in the 50s" |
| "any damage racial" (round 23, R23-3) | combat.md:77, :87; assassination.md:78, :84; subtlety.md:76, :90; Combat Adrenaline Rush note | OK. Damage actives: Orc, Troll, Night Elf, Gnome (`races.json`); Undead, Human, Dwarf have none, and the lines no longer send them to a fear break |
| Ghostly Strike note dagger condition (R23-1) | Subtlety card row 2 | OK at every band. With a dagger: 1.8 x (W + 1.8 AP/14) against W + 1.7 AP/14 plus 15 to 68 (Sinister Strike ranks 4 to 8 from 26). "Harder than Hemorrhage with any weapon": the same weapon, 125% unnormalized against 100% normalized (2.4); every non-dagger main hand on the Subtlety table is 2.1 speed or slower (Diamond Hammer 2.5, Hammer of the Northern Wind 2.1, Teebu's Blazing Longsword 2.9), where it holds at any attack power. A main hand under 1.92 speed that is not a dagger can lose it at high attack power: S4 |
| Slow-sword wording (R23-2) | subtlety.md:105; index.md:26 | OK: the 60 table's main hand is Teebu's (sword, 2.9); R23-2's arithmetic stands |
| Garrote window reason (R23-4) | subtlety.md:80 | OK: Garrote 50 less Dirty Deeds 20 = 30 (Dirty Deeds 18 to 19) against Hemorrhage 35, 32 at 56, 30 from 57: "no more energy" at every level Hemorrhage exists; two points with Initiative 3/3 (from 32; a 60 paragraph) |
| "left sitting ready" (R23-5) | subtlety.md:86 | OK |

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| ROGUE-R24-1 | CONFIRMED | page prose ("on cooldown" against the card's own condition; a routed item sim round 21 resolved) | subtlety.md:133 | "Shank comes at 24 for every rogue: press it on cooldown." | Sim round 21 (K1) settled the routed "Shank gate without Boot Knife" item by keeping a hold at every level from 24: the Subtlety card casts Shank only at three combo points or fewer or under 50 energy (`rogue-subtlety.json` row 1), and the same page's :76 says "At four or five with 50 energy or more, finish first and Shank right after". The leveling paragraph runs to 60 and Boot Knife comes at 35 (5/5 at 39 on the one order; 1329286, 100% for two points at 5/5): from then on a Shank pressed at five combo points spills both Boot Knife points, the thing :76 and the card rule out. K1's leveling bands measure it: the card's rule beats no hold at all by about 0.3% at 50 (221.5 against about 220.8, from the old rule's 219.0 plus 0.8%) and 0.2 to 0.4% at 56 and 59 (250.2 against about 249.2 to 249.7). Below 35 the hold changes nothing (K1: 24 and 28 identical, 30 and 34 within 0.1%), so the fix keys it to Boot Knife. Round 23 D5 held this back only because the gate was routed. | "Shank comes at 24 for every rogue: press it on cooldown, and once Boot Knife comes at 35, finish first when you sit on four or five combo points with 50 energy or more." |

Counts: CONFIRMED 1, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0.

## Round 24 lenses with nothing to file

- **Class index teaching a rotation the cards no longer use:** index.md:26 to :31 against the three cards. Shank "on cooldown" with the Subtlety 50-energy exception (matches row 1 of each card); poison hands (Deadly main hand for Combat and Subtlety, matching the sim kit in `sim/request/raid_consumes_test.go`:139 to :141; Assassination's swap at 60, matching `assassination.md`:88; R21-S1 wording carried); daggers for Combat and Assassination (every Combat and Assassination band of `bis/` is daggers). Clean.
- **Rank advice against per-race rank rows:** no rogue spec has rows in `data/curated/rank-choices.json` or `sim/rankchoices/rank-choices.json`, and every rogue line in `addon-data.json` has empty `race_spell_ids`. The only rank statement is Mutilate's "79 each at the top rank" (1241584, 79.2; R22-S1 carried). Nothing to check against.
- **Reasons that hold only above or below a gear level:** Ghostly Strike against Sinister Strike and Hemorrhage (above; S4); slow sword against dagger (holds on the table's weapons at every attack power for Ghostly Strike, to about AP 5600 for Hemorrhage); "Below 50 energy, Shank anyway" (measured on bare 24 to 59, bare 60 and raid 60); Assassination's "Your off hand lands more than enough hits to hold Deadly Poison at five stacks" and Combat's "Two Instants lose to that pair" (60 statements, measured at 60 in sim round 20's `hands`); "Rupture is not worth a global on a boss" (Combat, raid build). None fails inside the range the page states.

## Shoulds (wording, not findings)

| id | file:line | note | suggested wording |
|---|---|---|---|
| ROGUE-R24-S1 | subtlety.md:76; Subtlety Shank note | "finish first and Shank right after" assumes a finisher is ready. Between marks (Death Mark's 10 second cooldown, Rupture ticking, Slice and Dice up) there is none at four or five; the card then presses Hemorrhage and lets Shank go once energy drops under 50, so the hold can be longer than "one global" (R23-S2 carried). | "At four or five with 50 energy or more, finish first if a finisher is ready and Shank right after; between marks, Shank as soon as you drop under 50." |
| ROGUE-R24-S2 | subtlety.md:76 "a point Boot Knife might spill"; Shank note "a spilled point" | At five combo points Boot Knife spills both points, at four one. | "the points Boot Knife might spill" / "worth more than spilled points". |
| ROGUE-R24-S3 | subtlety.md:127 (new, Levels 1 to 19) | Tells the player to keep Slice and Dice up inside the 1 to 19 section without its level; :123 gives "Slice and Dice ... at 10" two paragraphs up, so it is unspecific, not false. | "From 10, Rupture does not train until 20, so ..." |
| ROGUE-R24-S4 | Subtlety Ghostly Strike note "harder than Hemorrhage with any weapon" | Holds with a dagger and with any main hand of 1.92 speed or slower (every non-dagger main hand on the Subtlety table), at any attack power. With a faster non-dagger main hand, Hemorrhage's normalized 2.4 outgrows Ghostly Strike's 1.25 x speed per attack power (a 1.8-speed sword breaks even near W = 0.6 x AP/14). | "harder than Hemorrhage with any weapon you are likely to hold" or drop "with any weapon". |
| ROGUE-R24-S5 | carried | R23-S1 (Blade Flurry "at any energy"), R23-S2 (folded into S1), R23-S3 (assassination.md:133 "every few seconds"), R23-S4 (dagger-only percents at subtlety.md:53, :58), R22-S1, R22-S2, R22-S3, R22-S4 (Subtlety rows before their talent, component lane), R21-S1 (index.md:27 "until 60"; sim round 21 suggests "from 30 until 60"), R21-S3 (Assassination Eviscerate "from 45" against Seal Fate at 43 on the raid order), R21-S4, R19-S6 (the Sinister Strike 60-energy gate "with a dagger" on all three cards). R21-S2 and R21-S6 are closed by sim round 21 (assassination.md:88; the finishers before Rupture). | As filed. |

## Prose, line by line (all four pages)

- **index.md** :22 to :37. :26 OK (R23-2 fix). :27 S5 (R21-S1). :28 Redirect 438040, 15 s. :29 Poised 1328825 at 18. :30 Shank 1329283 at 24 (15), 1329549 at 44 (25), 10 s, cannot miss; the Subtlety clause matches the card; Venom, Flash of Steel, Boot Knife. :31, :35, :37 races against round3/race-sims.md and `races.json`: Undead first in every tree, Orc over Troll; Night Elf leads Assassination and Subtlety, Gnome Combat; Skyborne only as the paid option. OK.
- **assassination.md** :30 to :133. Strengths, weaknesses, the 36/10/5 trade (unchanged orders; round 23 sums stand). :54 Mutilate 90% plus 79.2 (S5). :55 Seal Fate once a cast. :56 Venom 30% and 20 points; 20 + 10 + 20 = 50% at 60. :57 Relentless 20% a point at 5/5 (client 4% a point a rank). :58 Cold Blood 3 min. :59 Lethality 220%. :64 toll. :72 Garrote opener. :78 and :80 match the card's order (Shank, Slice and Dice at four, Rupture with Blood Splatter, Cold Blood, Eviscerate at four, Mutilate). :84 OK. :88 new, OK. :92 to :95 Kidney Shot 25 / 20 s / 6 s at five on rank 2, Redirect, Feint, Vanish, Evasion. :99 stat reasons. :103 daggers. :111 races. :115 professions. :121 to :127 trainer levels exact (Backstab 4, Gouge 6, Evasion 8, Slice and Dice and Sprint 10, Garrote 14, Ambush and Poised 18); Improved Gouge 10 to 12, Remorseless 13 and 14, Malice 15 to 19 on the leveling order. :131 to :133 Shank 24, Mutilate 40, Instant 20, Deadly 30, Venom 30, Relentless 39 to 44, Vanish 22; S5.
- **combat.md** :30 to :132. :52 trade unchanged. :54 150 extra energy, two and a half 60-energy Backstabs. :55 to :63 as round 23. :77 matches the card (Shank, Eviscerate at five, Slice and Dice at one or two, Backstab, Sinister Strike without a dagger); "40 back with Flash of Steel" at 5/5 (raid, by 43). :81 OK (S5). :87 OK. :93 to :96 OK. :100, :104, :108 as round 23. :112 races against round3 (Gnome and Undead lead; Orc 756.6, Troll 756.0, Night Elf 755.9). :116 professions. :122 to :132 trainer and order exact (Coup de Grace 10 to 12, Efficient Strikes 13, Lightning Reflexes 14 to 18, Puncturing Wounds 19; Cheap Shot 26; Aggression 28 to 30; Improved Gouge and Remorseless 44 to 48; Vanish 22).
- **subtlety.md** :29 to :133. :51 12/8/31. :53 to :62 (S5 for the dagger percents). :70 opener: Premeditation 2 + Garrote 1 + Initiative 1, Ghostly Strike for the fifth; matches the prepull (Stealth -3 s, Premeditation -2 s, Garrote 0 s). :76 sim round 21 and round 23 text, OK against the card (S1, S2). :78, :80, :82 OK. :86 OK. :90 Preparation 2 + 2 + 1 = 5. :94 Blind 5 min less 90 s with Elusiveness. :101 to :117 OK. :123 trainer levels exact. :125 Poised, Ambush 60 less 10. :127 new, OK against the card (S3). :129 Master of Deception 10 to 12, Opportunity 13 to 17 (4% a point to Backstab, Garrote, Ambush, Mutilate), Dirty Deeds 18 and 19. :133 levels: Premeditation 20, Ghostly Strike 26, Hemorrhage 30, Boot Knife 35, Death Mark 40, Camouflage 24, Cheap Shot 26: exact; Shank line **ROGUE-R24-1**.
- **Voice:** no em or en dashes, no site, sim or band talk. No page names 70755 (the curated JSON top-level `notes` do, as history; not rendered on any page).

## Sweep of superlatives, comparisons and quoted values

- Re-run over every page and note after the changes above; round 23's table stands. New quoted values: Shank's 50 energy threshold (the card's condition), Boot Knife "from 35" and "5/5 by 39", Rupture "trains at 20", Death Mark "from 40": all exact. New comparisons: "the energy it hands back is worth more than a point Boot Knife might spill" (measured on every band from 30, S2 for "a point"); "pays back more energy than it costs" (5 against 15 or 25, both ranks); Ghostly Strike against Hemorrhage (S4). "Only" and "every": "your only damage cooldown" (Subtlety, no Cold Blood or Adrenaline Rush in 12/8/31); "every Shank at 5/5" (Boot Knife 100% at 5/5). Hold.

## Decisions

- **D1.** R24-1 is filed now, not routed: sim round 21 decided the Shank rule the routed item asked about, the card and :76 carry it, and :133 still says the reverse for 35 to 60. Below 35 the hold is measured as no change, so the fix names Boot Knife rather than 24.
- **D2.** The Subtlety Slice and Dice row sits under Eviscerate at 20 to 39 with a condition Eviscerate's covers. I checked whether that makes the Slice and Dice note unreachable: the engine skips a row it cannot pay for (`apl_actions_casting.go`:29-30, `spell.CanCast`), and Slice and Dice (25) becomes affordable before Eviscerate (35) as energy climbs after a builder, so the card does refresh it. Not filed.
- **D3.** "finish first ... a 10 second cooldown loses nothing to one global" (S1) is incomplete between marks, not false where a finisher is ready; under the filing bar it is a SHOULD.
- **D4.** Backstab rank 9 and Eviscerate rank 9 at 60 (Engine questions) are not guide findings: no page quotes either rank's value, and both cost what the pages say (60 and 35).
- **D5.** Nothing routed in the ledger's Rounds 19 to 23 is re-filed. Sim round 21 closed SIM-R11-1 (finishers before 20) and the Assassination hands item; the pages now state both correctly.

## Sims requested

- Optional: Subtlety at five combo points between marks. The card presses Hemorrhage at five (row 8 has no combo-point gate) while it holds Shank at 50 energy or more. Compare the current card against Hemorrhage gated at four or fewer (falling through to Shank or to waiting), raid 60 and bare 50, seeds 1 and 50001. Decides S1's wording if it moves.
- Carried (already routed): Subtlety band 20 weapon (K4 of round 20: the band 20 table's main hand is Diamond Hammer, a mace, while the 1 to 19 and card lines build on Backstab); rogue BiS and race re-runs at the re-rank; K1's 50-energy threshold on the re-ranked BiS.

## Sim findings

- None new. Carried: Death Mark storage and the Mutilate and Shank off-hand damage on a level 40+ log (sim round 21 "Still open").

## Engine questions

- **Rogue trainer rows at 60 beyond the cards' ranks** (same shape as WARRIOR-R23 W23-1 and the coordinator's held Revenge, Battle Shout and Heroic Strike rows): `trainables/rogue.json` lists Backstab rank 9 (25300, "plus 150") and Eviscerate rank 9 (31016, 108 base a point) at 60 with `source: skill_line_ability`, beside Feint rank 5 (25302) and Deadly Poison V (25351). The Combat, Assassination and Subtlety cards cast Backstab 11281 (plus 140) and Eviscerate 11300 at 60. Settle with the book-rank question for all classes.

## Sources

- Client 1.60.1.70790: `spelltips/rogue.json` (Shank 1329283 and 1329549, Ghostly Strike 14278, Hemorrhage 16511, Slice and Dice 5171 and 6774, Rupture 1943 to 11275, Garrote 703 to 11290, Death Mark 1316709 to 1325758, Sinister Strike 1752 to 11294, Backstab 53 to 25300, Eviscerate 2098 to 31016, Mutilate 399956 to 1241584, Premeditation 14183), `talents/rogue.json` (Boot Knife, Efficient Strikes, Initiative, Dirty Deeds, Relentless Strikes, Flash of Steel, Opportunity, Serrated Blades, Ruthlessness, Improved Slice and Dice, Blood Splatter), `trainables/rogue.json`, `addon-builds.json`, `addon-data.json` (every rogue band), `bis/rogue-subtlety.json` (main and off hands, every band and faction), `items/rogue.json`, `races.json`.
- Engine 170a59c39: `sim/core/apl_actions_casting.go`:29-30.
- `design/reviews/2026-11-05-sim-audit-round21/rogue.md` (K1 to K4), `integration.md`; `sim/request/testdata/ladder/rogue-subtlety.golden.md`; `sim/request/raid_consumes_test.go`:139-141; round23/rogue.md; round3/race-sims.md; ledger Rounds 19 to 23. Web research: none needed (every call decided on the client, the cards and the committed measurements).

## Appendix: Note sweep

One line per row note: spec, row, bands it renders on, each named talent, spell, rank or value with the level the player has it (L = leveling order, R = raid order), and the verdict.

**Assassination** (23 bands, 10 to 60; both orders; no note changed since round 23)
1. Shank (24-60): rank 1 at 24 (5 for 15), rank 2 at 44 (25); cannot miss; Venom L30 / R30 behind "with Venom"; Vile Poisons L27 / R25 before it; 30%, 6 s, 20 points (1329328). OK.
2. Slice and Dice (10-60): rank 1 at 10, rank 2 at 42; four or more, under 3 s; matches the condition. OK.
3. Rupture (20-60): trained 20; only with Blood Splatter (R22, never on L; `auraIsKnown` gate), +6 s; ticks crit (owner ruling); Slice and Dice up matches. OK.
4. Cold Blood (25-60): talent L33 / R31; rows 25 to 32 early (component lane); off the global, no cost (14177). OK.
5. Eviscerate (10-60): four or more; Mutilate from 40 (L40 / R40); Seal Fate "from 45" (L45; R43, S5 carried). OK.
6. Mutilate (40-60): ranks 40 / 50 / 60; both daggers, two points, +20% poisoned at every rank. OK.
7. Backstab (10-60): rank 1 at 4 to rank 8 at 60; until Mutilate; dagger stated; 150% plus 10 to 140 against 100% plus 3 to 68. OK.
8. Sinister Strike (10-60): before Mutilate; 60-energy gate (S5, R19-S6). OK.

**Combat** (22 bands, 10 to 60; both orders)
9. Shank (24-60): 10 s at both ranks; every 8 with five Flash of Steel (R 5/5 by 43; L 4 points, conditional clause). Pays back more than it costs. OK.
10. Eviscerate (10-60): at five; matches. OK.
11. Slice and Dice (10-60): one or two points when it drops; five to Eviscerate. OK.
12. Backstab (10-60): Puncturing Wounds once you have it (L19 / R17); raid Lethality from 54, Opportunity from 56 (R54 / R56). OK.
13. Adrenaline Rush (40-60): talent L40 / R40; "your damage racial if your race has one" (R23-3 fix): Gnome Eureka!, Orc Blood Fury, Troll Berserking, Night Elf Elune's Light have one; the others are told nothing false. OK.
14. Sinister Strike (10-60): only when Backstab cannot be cast; 60-energy gate (S5). OK.

**Subtlety** (22 bands, 10 to 60; one order)
15. Shank (24-60), **new**: "three combo points or fewer, or under 50 energy" matches the condition; cannot miss; pays back more than it costs at 24 (15) and from 44 (25); Boot Knife "from 35", "a chance" (20% a point), "every Shank at 5/5 by 39" (35 to 39): no early credit on 24 to 34; "at four or five with 50 energy or more you finish first" matches the hold (S1 for between marks); "the energy is worth more than a spilled point" (measured; S2). OK.
16. Ghostly Strike (10-60): talent at 26 (rows 10 to 24 early, component lane). With a dagger, harder than Sinister Strike, no more energy (40 against 45, 42 at 56, 40 from 57): OK every band. From 30 harder than Hemorrhage with any weapon: OK on the table's weapons (S4). 40 against 35, 30 from 57. One point. OK.
17. Garrote (14-60): ranks 14 / 22 / 30 / 38 / 46 / 54; Ghostly Strike (26) lets it in without Stealth; four or fewer, ahead of Rupture at four (row order); Opportunity from 13; Hemorrhage (30) and Death Mark (40) "once you have them". OK (R22-S4 for 14 to 24).
18. Rupture (20-60): trained 20; Serrated Blades from 21; Hemorrhage from 30. OK.
19. Death Mark (40-60): ranks 40 / 48 / 56 / 60; 6 s; Garrote and Rupture twice the rate; 10 s cooldown every rank; four or more while Rupture ticks matches. OK.
20. Eviscerate (10-60), **new**: three or more while Rupture ticks (20 to 39), or before Rupture trains at 20 while Slice and Dice is up (10 to 19), only before Death Mark (40): matches the condition at every band. OK.
21. Slice and Dice (10-60), **new**: three or more while Rupture ticks, or any time before 20, down or under 3 s: matches the condition (D2 on its place under Eviscerate). Rank 1 to 41, rank 2 from 42. OK.
22. Hemorrhage (30-60): talent 30; one point; 15% to Rupture and Garrote; holds at under 1 s / 75 energy and after a Ghostly Strike match the condition. OK.
23. Backstab (10-60): until Hemorrhage (30); dagger stated; far harder than Sinister Strike; Opportunity from 13 at 4% a point. OK.
24. Sinister Strike (10-60): before Hemorrhage, only when Backstab cannot be cast; 60-energy gate (S5). OK.

24 notes swept: 24 OK. The one finding this round is page prose (R24-1).
