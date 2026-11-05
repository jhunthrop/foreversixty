# Guide accuracy round 22: Rogue

Audited on the round 22 audit tree (987d08a7c: main after guide round 21; no sim round 21 commit and no `*sim-audit-round21*` report in the tree; engine pin d5c708834 per `sim/enginever/version.go`), against the active client 1.60.1.70790 (`data/builds/1.60.1.70790/`). Files: `web/src/content/guides/rogue/{index,assassination,combat,subtlety}.md`, every line; every note in `data/curated/apl/rogue-{assassination,combat,subtlety}.json`, every sentence, at every band it renders in `addon-data.json` (`rotations/rogue-assassination` 23 bands, `rogue-combat` 22, `rogue-subtlety` 22), against BOTH shipped orders in `addon-builds.json` and the trainer levels in `trainables/rogue.json`. Read first: round 21 rogue report and the ledger's Round 19, 20 and 21 sections.

## Changes since round 21's tree (e039e1490)

- Guides: assassination.md:121 only (1bc247702, the round 21 fix). No rogue card action or note moved; `addon-data.json` was regenerated (note text only, no rogue line changed).

## Shipped orders and trainer levels used

Unchanged from round 21 (see round21/rogue.md). Trainer levels added this round for the new shape checks: Kick 12, Expose Armor 14, Feint 16, Distract 22, Vanish 22 / 42, Detect Traps 24, Blind 34, Sprint 10 / 34 / 58, Kidney Shot 30 / 50 (5 s at five on rank 1, 6 s on rank 2), Cheap Shot 26 (60 energy, 4 s, 2 points).

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| ROGUE-R22-1 | CONFIRMED | card note (comparison names a spell on a band before it is trained) | `rogue-subtlety.json` Ghostly Strike note, band 28 (and 10 to 24 where the row renders early) | "it hits harder than Hemorrhage for a little more energy (40 against Hemorrhage's 35, or 30 with both points of Efficient Strikes)" | Shipped Subtlety order: Ghostly Strike 26, Hemorrhage 30 (`trainables/rogue.json` Hemorrhage 16511 level 30). The band 28 row casts Ghostly Strike (14278) with no Hemorrhage on the bar or in the rows; the player's builders at 28 are Backstab rank 4 (2591, 150% plus 46, 60 energy) and Sinister Strike rank 4 (1759, 100% plus 15, 45 energy). The note compares against a talent two levels away. What holds at every band: Ghostly Strike 125% (180% with a dagger) for 40 beats Sinister Strike (100% plus 3 to 68) for 45, and from 30 the Hemorrhage comparison is true (100 / 145% for 35). It does not beat Backstab at 28 (1.8W against 1.5W plus 46 for any leveling dagger), so do not name Backstab. | "Ghostly Strike on cooldown at four combo points or fewer: it hits harder than Sinister Strike for less energy, and from 30 harder than Hemorrhage for a little more (40 against Hemorrhage's 35, or 30 with both points of Efficient Strikes), and gives a combo point." |
| ROGUE-R22-2 | CONFIRMED | prose ("every" claim a talent covers only part of) | combat.md:54 | "Adrenaline Rush ... Every Backstab you fit into it is a free one." | Adrenaline Rush (13750): "Increases your Energy regeneration rate by 100% for 15 sec." Base regeneration is 10 energy a second, so the 15 seconds give 300 energy of which 150 is the talent's: about five Backstabs (60 each) fit into it and two and a half of them are paid by Adrenaline Rush, the rest by normal regeneration. Not every Backstab in it is free; half are. | "Adrenaline Rush (1 point): double energy regeneration for 15 seconds on a 5 minute cooldown: 150 extra energy, two and a half Backstabs' worth." |
| ROGUE-R22-3 | LIKELY | leveling (spell named in a 20 to 60 paragraph without its trainer level) | assassination.md:133; combat.md:132 | "Sprint, Evasion and Vanish get you out of the bad ones." / "Sprint, Evasion and Vanish get you out of the pulls that go wrong." | Both sentences sit in the Levels 20 to 60 section; Vanish (1856) trains at 22 (`trainables/rogue.json`), so at 20 and 21 it is not on the bar. Sprint (10) and Evasion (8) are. Every other spell in these paragraphs carries its level (Shank 24, Cheap Shot 26, Instant Poison 20, Deadly Poison 30, Mutilate 40). The same round 21 shape (Soul Harvest named without its level) was taken. | "Sprint, Evasion and, from 22, Vanish get you out of the bad ones." (Combat: same insertion.) |

Counts: CONFIRMED 2, LIKELY 1, NEEDS-SIM 0, NEEDS-OWNER 0.

## Round 21 fix, re-checked

- **assassination.md:121** "With a dagger in hand, Gouge, step behind and Backstab: it hits far harder than Sinister Strike." Levels 1 to 19: Backstab 53 (4, 150% plus 10) and 2589 (12, 150% plus 20) against Sinister Strike 1752 (1, plus 3), 1757 (6, plus 6), 1758 (14, plus 10). With any dagger the player has, 1.5W plus 10 to 20 against W plus 3 to 10 is about half again harder: "far harder" holds at every level 4 to 19, and Gouge (6) is named with its level in the sentence before. Holds. Ambush (18) is now not contradicted.

## Round 22 shapes, checked on every page and note

- **Spell named in a leveling paragraph before it trains:** every spell in the four Levels 1 to 19 paragraphs carries its trainer level (Backstab 4, Gouge 6, Evasion 8, Slice and Dice and Sprint 10, Kick 12, Garrote 14, Ambush 18, Poised 18). In Levels 20 to 60: Shank 24, Cheap Shot 26 (Combat, Subtlety; Assassination's sits after "Once Mutilate comes at 40"), Instant 20, Deadly 30, Venom 30, Mutilate 40, Relentless Strikes 39 to 44, Premeditation 20, Camouflage 24, Ghostly Strike 26, Hemorrhage 30, Boot Knife 35, Death Mark 40, Aggression by 30, Improved Gouge and Remorseless in the mid 40s (Combat 44 to 48). Rupture (20) and Gouge (6) are trained before the paragraph starts. The one gap is Vanish (R22-3).
- **Note comparing to ranks or spells before they are trained:** R22-1. Every other comparison in a note is to a spell the player has at that band (Backstab against Sinister Strike, every band; Eviscerate and Mutilate from 40; Death Mark from 40), or is conditional ("with Venom", "once you have Puncturing Wounds", "once you have them", "in the raid build ... from 54").
- **"Cheap" or "best" list that includes its own top item:** none. "Shank ... the cheapest energy you have" (assassination.md:84): 5 for 15 or 25, the only energy you buy with energy; Thistle Tea is a consumable named separately. "Efficient Strikes keeps Hemorrhage cheap" (subtlety.md:86): 30, the cheapest builder at 60 (Ghostly Strike 40, Sinister Strike 40, Backstab 60). "The best energy in the tree" (combat.md:57): Shank with five Flash of Steel nets 35 every 8 s (4.4 a second) against Adrenaline Rush's 150 every 5 minutes (0.5 a second). OK.
- **Rounded value quoted as exact:** costs all exact (Shank 5, Garrote 30 with Composure and with Dirty Deeds, Ambush 50 with Poised, Eviscerate 20 with Flawless Execution not quoted, Hemorrhage 30, Kidney Shot 25, Kick 25, Blade Flurry 25, Death Mark 25, Efficient Strikes 3 at one point exact on 13732 rank 1). One damage value is rounded: Mutilate "79" against the client's 79.2 (S1).
- **"Most damage for your mana" / "every heal":** no mana. "Every" claims checked: Seal Fate "every crit from a builder" (14186: abilities that add combo points), Venom "every Shank", Dual Wield Specialization "every off-hand swing" (includes Shank), Aggression "every Sinister Strike", Dirty Deeds "every Garrote after a Ghostly Strike", Boot Knife "every Shank" at 5/5, Remorseless Attacks' list at 13 (Mutilate and Ghostly Strike not yet trained), Lethality's list (no Ambush, no Shank), Cold Blood "Mutilate or Eviscerate" (a subset of its five), Opportunity's lists on each page. R22-2 is the one that fails. Subtlety :82 "stores every tick at twice the rate" reads as Rupture's ticks in context (S3).
- **"Your biggest hit" with raid buffs at 60:** no rogue page or note says biggest or hardest any more. Comparisons at 60: Ambush "hits harder on paper" than Garrote (assassination.md:72), Ghostly Strike over Hemorrhage (180 / 125% against 145 / 100%), Backstab over Sinister Strike (150% plus 140 against 100% plus 68), all true.
- **Round 20 and 21 shapes (raid-order credit, superlatives over a span, quoted values a build talent changes):** re-run; nothing new beyond the above. The carried shoulds stand.

## Shoulds (wording, not findings)

| id | file:line | note | suggested wording |
|---|---|---|---|
| ROGUE-R22-S1 | assassination.md:54 | Mutilate rank 3 (1241584) is "plus an additional 79.2 with each weapon"; the page says 79. Rounded, not wrong in play. | "plus 79.2 each at the top rank" (and "the flat 79.2 does not"). |
| ROGUE-R22-S2 | assassination.md:133 | "from then on most finishers hand their energy back": 25 back against Eviscerate's 35; Slice and Dice and Rupture (25) are full refunds. :84 already says "most of the finisher's cost". | "from then on most finishers hand most of their energy back." |
| ROGUE-R22-S3 | subtlety.md:82 | "Death Mark stores every tick at twice the rate": Deadly Poison ticks are stored at the normal rate; only Garrote and Rupture are doubled (1316709). The sentence is about Rupture. | "Death Mark stores every Rupture tick at twice the rate." |
| ROGUE-R22-S4 | Subtlety Garrote note, bands 14 to 24 | "Garrote after every Ghostly Strike" renders before Ghostly Strike (26), like the Ghostly Strike row itself (component lane, accepted since round 19). Nothing false at 28 and up. | Component lane: start the Ghostly Strike and Garrote-after-Ghostly-Strike rows at 26. |
| ROGUE-R22-S5 | carried | R21-S1 (index.md:27 "for Assassination until 60", pair starts at 30), R21-S2 (assassination.md:88, 30 to 39), R21-S3 (Assassination Eviscerate note at band 44 on the raid order), R21-S4 (Combat Shank 40 / every 8 at five points of Flash of Steel), R21-S5 ("with your racial" for Undead), R21-S6 (Subtlety bands 10 to 16 finishers, Sinister Strike 60-energy gate). | As round 21. |

## Prose, line by line (all four pages)

- **index.md** :22 to :37. Unchanged since round 21 and re-read: daggers line (Mutilate needs a dagger in each hand: engine `sim/rogue/mutilate.go`:122-125 matches the client dagger bit), poison hands (S5), Redirect (438040: 15 s, next non-player enemy), Poised (1328825), Shank (1329283 / 1329549: 5 energy, 15 / 25, 10 s, cannot miss), Hack and Slash (dagger / fist 5% crit). Races against `races.json` and round3/race-sims.md: unchanged, OK. Skyborne named only as the paid option.
- **assassination.md** :30 to :133. Strengths and weaknesses OK (Cold Blood 3 min, 14177; Mutilate needs both daggers). Builds and the trade at :52 exact. :54 S1. :55 Seal Fate 100% at 3/3. :56 Venom 20 + Improved Poisons 10 + Venom 20 = 50% at 60. :57 Relentless Strikes 20% a point, 25 energy, 80% at four. :58 Cold Blood. :59 Lethality 20% bonus on its six abilities (no Ambush). :60 to :64 Vile Poisons 20%, Improved Poisons, Coup de Grace 20%, Opportunity 20% (Mutilate and Garrote both on its list), Blood Splatter 6 / 9 s, Vigor 5 at one point (14983 rank 1), the five-point toll. :72 Ambush gets none of Lethality, OK. :78 to :84 OK (Thistle Tea 100 energy; engine `sim/core/consumes.go`:1397-1409). :88 S5. :92 Kidney Shot 25 / 20 s / 6 s at five on rank 2 (a 60 raid paragraph). :99, :103 OK. :111 races OK. :115 professions OK. :121 to :127 OK (round 21 fix holds; Improved Gouge 1.5 s at 3; Remorseless 40% at 2 on Sinister Strike, Backstab, Ambush, the three the player has at 13). :131 OK. :133 R22-3, S2.
- **combat.md** :30 to :132. :34 to :47 OK (Blade Flurry "strike an additional nearby opponent", 13877). :52 trade exact. :54 R22-2. :55 Blade Flurry 20%, 15 s. :56 Hack and Slash. :57 Flash of Steel 15 / 2 s at five (S5 for four). :58 Weapon Expertise 2%. :59 Dual Wield Specialization 50%. :60 Puncturing Wounds 30% / 45%. :63 Lethality 220%; "more than half your Backstabs crit" (Puncturing Wounds 30, Malice 5, Hack and Slash 4 at four points, plus raid gear crit); Aggression 15% on Sinister Strike and Eviscerate; Flawless Execution 15; Efficient Strikes 3 at one point; Composure 10 more, Garrote 50 less 20 = 30; row-four toll. :71 to :96 OK (Kick 25; Blade Flurry two per 3 minutes, Adrenaline Rush one). :100 to :116 OK. :122 to :128 OK (Lightning Reflexes 5% at 5/5 by 18, 13712). :132 R22-3.
- **subtlety.md** :29 to :133. Unchanged; re-read against 14278, 16511, 1316709 to 1325758, the talent rows and the one order. :33 no facing requirement holds (Dirty Deeds lets Garrote in from the front). :36 Elusiveness 90 s on Vanish and Blind. :53 to :62 OK (Death Mark 12% a point, double for Garrote and Rupture; Coup de Grace 20%; Initiative 100%; Dirty Deeds 20 at two points, 10 at one, frontal from one; Efficient Strikes 5 at two). :70 Cheap Shot 2 + Initiative 1 = 3, 4 s. :78, :80 Garrote 30 after Dirty Deeds. :82 S3. :90 Preparation 10 min, "finishes the cooldown on your other Rogue abilities" (14185): Premeditation 2 + Shank with Boot Knife 2 + Ghostly Strike 1 = 5. :94 Blind 5 min less 90 s. :123 to :133 OK (Ambush 60 less Poised 10 = 50; Master of Deception "as if you were 3 levels higher"; Opportunity 13 to 17; Dirty Deeds 18 and 19).
- **Voice:** no em or en dashes, no site, sim or band talk, no nightly numbers. No page names 70755.

## Decisions

- **D1.** R22-1 is CONFIRMED under round 22's added check ("a note comparing to a spell does not name that spell on bands before it is trained"): band 28 is a band where the Ghostly Strike row is legitimately owned (talent at 26) and Hemorrhage is not (30). The early rows 10 to 24 stay with the component lane; band 28 is not a rendering artefact. The fix compares to Sinister Strike, which the player has at every band and which Ghostly Strike beats on both damage and cost; Backstab is left out because it hits harder than Ghostly Strike at 28.
- **D2.** R22-2 is CONFIRMED under the "every attack a talent covers only some of" shape: Adrenaline Rush pays for half of the Backstabs in its window, not all of them. The fix quotes the client's fixed value (100% of a 10-a-second base for 15 s), not a nightly number.
- **D3.** R22-3 is LIKELY rather than CONFIRMED: the sentence names no level and sits at the end of a 40-level span, but a level 20 or 21 reader is told to use a button they do not have, and every other spell in the same paragraphs carries its level. Taken as a fix at the integrator's discretion, matching the round 21 warlock Soul Harvest call.
- **D4.** Mutilate's "79" is a SHOULD: the client's 79.2 rounds to it and nothing in play differs.
- **D5.** Nothing from the ledger's Round 19, 20 or 21 routed lists is re-filed; no sim round 21 has landed and no page states any of them falsely.

## Sims requested

- None new. Carried (already routed): Assassination poison hands at 34 and 38 on the leveling order; Subtlety Shank gate without Boot Knife at 24 to 32; Subtlety bands 10 to 16 finisher fallback (SIM-R11-1); Subtlety band 20 weapon (K4); rogue BiS and race re-runs at the re-rank; X2 re-check on the new BiS.

## Sim findings

- None new. Carried: K1 (Death Mark storage and the Mutilate / Shank off-hand damage, waiting on a level 40+ log).

## Engine questions

- None new. Carried (ledger Round 20): whether Vile Poisons, Venom and Malice scale the attack-power share of each Instant hit and Deadly tick. The pages state nothing about it.

## Sources

- Client 1.60.1.70790: `spelltips/rogue.json` (Adrenaline Rush 13750; Ambush 8676 to 11269; Backstab 53 to 25300; Blade Flurry 13877; Blind 2094; Cheap Shot 1833; Cold Blood 14177; Death Mark 1316709, 1325756 to 1325758; Evasion 5277; Garrote 703 to 11290; Ghostly Strike 14278; Hemorrhage 16511; Instant Poison 8679 to 11340; Deadly Poison 2823 to 25351; Kick 1766 to 1769; Kidney Shot 408, 8643; Mutilate 399956, 1241582, 1241584; Poised 1328825; Premeditation 14183; Preparation 14185; Redirect 438040; Shank 1329283, 1329549; Sinister Strike 1752 to 11294; Sprint 2983 to 11305; Vanish 1856, 1857; Venom 1329328); `talents/rogue.json` (every rank row of Remorseless Attacks, Malice, Ruthlessness, Improved Slice and Dice, Vigor, Lethality, Vile Poisons, Improved Poisons, Relentless Strikes, Seal Fate, Coup de Grace, Efficient Strikes, Lightning Reflexes, Puncturing Wounds, Composure, Aggression, Hack and Slash, Flash of Steel, Opportunity, Dirty Deeds, Camouflage, Initiative, Serrated Blades, Elusiveness, Boot Knife); `trainables/rogue.json`; `addon-builds.json`; `addon-data.json` (`rotations/rogue-*`, every band).
- Engine d5c708834: `sim/rogue/mutilate.go`:60-125, `sim/core/consumes.go`:1397-1409.
- `design/reviews/guide-accuracy/round21/rogue.md`; `ledger.md` Rounds 19 to 21.

## Appendix: Note sweep

One line per row note: spec, row, bands it renders on, each named talent, spell, rank or value with the level the player has it (L = leveling order, R = raid order), and the verdict.

**Assassination** (23 bands, 10 to 60; both orders)
1. Shank (24-60): rank 1 at 24 (5 for 15), rank 2 at 44 (5 for 25): pays back more than it costs at both. Cannot miss. Venom L30 / R30 behind "with Venom"; Vile Poisons in before Venom (L27, R25); 30%, 6 s and "land more often" (+20) from 1329328. OK.
2. Slice and Dice (10-60): rank 1 at 10, rank 2 at 42; four or more, under 3 s. OK.
3. Rupture (20-60): trained 20; Blood Splatter conditional (R22; never on L); +6 s; ticks crit (owner ruling). OK.
4. Cold Blood (25-60): talent L33 / R31; rows 25 to 32 render early (component lane); no cost, off the global (14177). OK.
5. Eviscerate (10-60): four or more; Mutilate from 40 (L40 / R40); Seal Fate from 45 (L45 / R43). OK; S5 (R21-S3 carried).
6. Mutilate (40-60): ranks 40 / 50 / 60; both daggers, two points, +20% poisoned at every rank. OK.
7. Backstab (10-60): rank 1 at 4 to rank 8 at 60; until Mutilate (40); "far harder than Sinister Strike": 150% plus 10 to 140 against 100% plus 6 to 68 at every band. OK.
8. Sinister Strike (10-60): before Mutilate; 60-energy gate (Backstab 60 at every rank). OK.

**Combat** (22 bands, 10 to 60; both orders)
9. Shank (24-60): 10 s at both ranks; every 8 with five Flash of Steel (R five by 43; L four, conditional clause). Pays back more than it costs. OK; S5 (R21-S4).
10. Eviscerate (10-60): at five. OK.
11. Slice and Dice (10-60): one or two points; five to Eviscerate. OK.
12. Backstab (10-60): Puncturing Wounds "once you have" (L19 / R17); "in the raid build Lethality from 54 and Opportunity from 56" (R54 / R56). OK.
13. Adrenaline Rush (40-60): talent L40 / R40. "With your racial" S5 (R21-S5). OK.
14. Sinister Strike (10-60): only when Backstab cannot be cast. OK.

**Subtlety** (22 bands, 10 to 60; one order)
15. Shank (24-60): three or fewer; cannot miss; pays back more than it costs at both ranks; Boot Knife from 35 "can add" two, every Shank at 5/5 from 39. OK.
16. Ghostly Strike (10-60): talent at 26 (rows 10 to 24 early, component lane). Names Hemorrhage (30) on band 28: **ROGUE-R22-1**. From 30: 40 against Hemorrhage 35, 30 with both Efficient Strikes (56-57, conditional); 125 / 180% against 100 / 145%. One combo point. OK from 30.
17. Garrote (14-60): trained 14 (8631 at 22, 8632 at 30, 8633 at 38, 11289 at 46, 11290 at 54). Ghostly Strike (26) lets it in without Stealth (S4 for 14 to 24, component lane). Opportunity from 13. Hemorrhage (30) and Death Mark (40) "once you have them". OK.
18. Rupture (20-60): trained 20; Serrated Blades from 21; Hemorrhage's debuff from 30. OK.
19. Death Mark (40-60): ranks 40 / 48 / 56 / 60; 6 s, Garrote and Rupture at twice the rate, 10 s cooldown at every rank. OK.
20. Eviscerate (10-60): three or more while Rupture ticks, only before Death Mark (40). OK; S5 (R21-S6) at 10 to 16.
21. Slice and Dice (10-60): three or more while Rupture ticks, under 3 s. OK; S5 at 10 to 16.
22. Hemorrhage (30-60): talent at 30; one point; 15% to Rupture and Garrote; hold rules (75 energy, under 1 s, after a Ghostly Strike). OK.
23. Backstab (10-60): until Hemorrhage (30); far harder than Sinister Strike at every band; Opportunity from 13 at 4% a point (14057). OK.
24. Sinister Strike (10-60): before Hemorrhage; 60-energy gate. OK; S5.

24 notes swept: 23 OK, 1 finding (R22-1). The other two items this round are page prose (R22-2, R22-3).
