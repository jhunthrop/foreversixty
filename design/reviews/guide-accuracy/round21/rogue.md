# Guide accuracy round 21: Rogue

Audited on the round 21 audit tree (e039e1490: main after guide round 20; no sim round 21 commit in the tree, no `*sim-audit-round21*` report; engine pin d5c708834 per `sim/enginever/version.go`), against the active client 1.60.1.70790 (`data/builds/1.60.1.70790/`). Files: `web/src/content/guides/rogue/{index,assassination,combat,subtlety}.md`, every line; every note in `data/curated/apl/rogue-{assassination,combat,subtlety}.json`, every sentence, at every band it renders in `addon-data.json` (`rotations/rogue-assassination` 23 bands, `rogue-combat` 22, `rogue-subtlety` 22), read against BOTH shipped orders in `addon-builds.json` and the trainer levels in `trainables/rogue.json`. Read first: round 20 rogue report, ledger Round 19 and Round 20, sim round 20 `rogue.md`.

## Changes since round 20's tree (4234643f5)

- Guides (db7aaa9dc, 85721a49e): index.md:27 (poison hands by tree), assassination.md:56 and :78 (Venom's Instant Poison on the main hand at 60).
- APL: the Combat Backstab note (Lethality from 54, Opportunity from 56). No other rogue card action or note moved; `addon-data.json` regenerated (note text only).
- All four round 20 fixes re-checked below as new work.

## Shipped orders used (index 0 = level 10; re-decoded this round, unchanged)

- Assassination leveling: Improved Gouge 10-12, Remorseless Attacks 13-14, Malice 15-19, Ruthlessness 20-22, Improved Slice and Dice 23-24, Lethality 25-26, Vile Poisons 27-29 and 31-32, Venom 30, Cold Blood 33, Improved Poisons 34-38, Relentless Strikes 39 and 41-44, Mutilate 40, Seal Fate 45-47, Coup de Grace 48-50, Lightning Reflexes 51-52, Puncturing Wounds 53-55, Precision 56-58, Master of Deception 59, Opportunity 60.
- Assassination raid: Malice 10-14, Ruthlessness 15-17, Improved Slice and Dice 18-20, Vigor 21, Blood Splatter 22, Lethality 23-24, Vile Poisons 25-29, Venom 30, Cold Blood 31, Improved Poisons 32-36, Relentless Strikes 37-39 and 41-42, Mutilate 40, Seal Fate 43-45, Coup de Grace 46-48, Lightning Reflexes 49-50, Puncturing Wounds 51-53, Precision 54-55, Opportunity 56-60.
- Combat leveling: Coup de Grace 10-12, Efficient Strikes 13, Lightning Reflexes 14-18, Puncturing Wounds 19-21, Precision 22-24, Flawless Execution 25, Weapon Expertise 26-27, Aggression 28-30, Dual Wield Specialization 31-32, Blade Flurry 33, Hack and Slash 34-38, Flash of Steel 39 and 41-43 (4), Adrenaline Rush 40, Improved Gouge 44-46, Remorseless Attacks 47-48, Malice 49-53, Ruthlessness 54-56, Improved Slice and Dice 57, Master of Deception 58-60.
- Combat raid: Coup de Grace 10-12, Lightning Reflexes 13-16, Puncturing Wounds 17-19, Precision 20-22, Composure 23-24, Flawless Execution 25, Weapon Expertise 26-27, Aggression 28-30, Dual Wield Specialization 31-32, Blade Flurry 33, Hack and Slash 34-37, Flash of Steel 38-39 and 41-43 (5), Adrenaline Rush 40, Malice 44-48, Ruthlessness 49-50, Improved Slice and Dice 51-53, Lethality 54-55, Opportunity 56-60.
- Subtlety (one order): Master of Deception 10-12, Opportunity 13-17, Dirty Deeds 18-19, Premeditation 20, Serrated Blades 21-23, Camouflage 24 and 27-28, Elusiveness 25 and 33, Ghostly Strike 26, Initiative 29 and 31-32, Hemorrhage 30, Preparation 34, Boot Knife 35-39, Death Mark 40, Malice 41-45, Ruthlessness 46-48, Improved Slice and Dice 49-50, Blood Splatter 51, Lethality 52, Coup de Grace 53-55, Efficient Strikes 56-57, Precision 58-60.
- Trainer: Backstab 4 / 12 (rank 2: 150% plus 20), Ambush 18 (rank 1: 250% plus 70), Instant Poison 20, Shank 24 / 44, Cheap Shot 26, Deadly Poison 30, Kidney Shot 30 / 50, Mutilate 40 / 50 / 60, Death Mark 40 / 48 / 56 / 60.

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| ROGUE-R21-1 | CONFIRMED | leveling (superlative that fails on a band where another spell is trained) | assassination.md:121 | "With a dagger in hand, Gouge, step behind and Backstab: it is the hardest hit you have this early." | The paragraph covers levels 1 to 19 and trains Ambush at 18 in the sentence before. `spelltips/rogue.json`: Ambush rank 1 (8676, trained at 18) "250% weapon damage plus 70", against Backstab rank 2 (2589, trained at 12, the top rank until 20) "150% weapon damage plus 20". At 18 and 19 Ambush is the harder hit, and the page's own next paragraph tells you to open with it. What holds at every level 4 to 19 is the comparison the Backstab card makes: it beats Sinister Strike (100% weapon damage plus 3 to 15, ranks 1752 to 1758). | "With a dagger in hand, Gouge, step behind and Backstab: it hits far harder than Sinister Strike." |

Counts: CONFIRMED 1, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0.

## Round 20 fixes, re-checked

- **index.md:27** "Deadly Poison on the main hand and Instant Poison on the off hand is right for Combat and Subtlety, and for Assassination until 60, where it swaps them." Holds: sim round 20 X2 kept Deadly MH for Combat (798.2 against 792.2 mirrored) and Subtlety (ties two Deadlys, beats the mirror), and Assassination swaps at 60 (+1.5 to +1.7%) and ties at bare 40 and 50. Deadly trains only at 30, so "until 60" really means 30 to 59 (Should S1). Agrees with assassination.md:88, :107, :133, combat.md:108, subtlety.md:109.
- **assassination.md:56** Venom: Instant Poison 20% (8679 to 11340, every rank), Improved Poisons 5/5 +10 points (14113), Venom +20 points (1329328): 50% on main-hand hits, the hand that carries Instant at 60. The engine adds the chances the same way (`sim/rogue/poisons.go`:37-50: 0.2 + improvedPoisons + additivePoisonBonusChance). Holds.
- **assassination.md:78** same arithmetic, "at 60" scopes the hand. "another 30% on top of Vile Poisons" reads as additive, which is what the client and engine do (`client_values_test.go`:119-136: 1 + 0.30 + 0.20). Holds.
- **Combat Backstab note** "in the raid build Lethality from 54 and Opportunity from 56 land on it too": raid order Lethality 54-55, Opportunity 56-60; Lethality (14128) and Opportunity (14057) both name Backstab; the leveling order takes neither and the "in the raid build" clause covers it. Holds.

## Every poison-hand line on the four pages

index.md:27 (above); assassination.md:56, :78 (above); :88 (Instant MH / Deadly OH at 60, Deadly MH below 60; S2 carried for 30 to 39); :103 (main hand carries Instant at 60); :107 (60 pair); :133 (Instant 20, Deadly 30, Deadly MH until 60); combat.md:104 (off hand carries Instant), :108 (Deadly MH, X2), :132 (Instant 20, Deadly 30); subtlety.md:105 (off hand carries Instant), :109 (Deadly MH, X2). All consistent with each other and with X2. No page puts Instant on the off hand at 60 for Assassination any more.

## Shoulds (wording, not findings)

| id | file:line | note | suggested wording |
|---|---|---|---|
| ROGUE-R21-S1 | index.md:27 | "for Assassination until 60": from 20 to 29 there is no Deadly Poison (trains at 30), so the pair starts at 30. Nothing false for a reader who has both. | "... is right for Combat and Subtlety, and for Assassination from 30 until 60, where it swaps them." |
| ROGUE-R21-S2 | assassination.md:88 | Carried R20-S2 ("Below 60 ... at 30 to 39": X2 measured bare 30 only). | As R20-S2. |
| ROGUE-R21-S3 | `rogue-assassination.json` Eviscerate note, band 44 | Carried R20-S1 (raid order has Seal Fate from 43; "from 45" under-credits). | As R20-S1. |
| ROGUE-R21-S4 | `rogue-combat.json` Shank note; combat.md:57, :77 | Carried R19-S1 / R19-S2 (leveling build's four points of Flash of Steel give 37 back every 8.4 s; "40" and "every 8" are the raid build's five). Both lines name five points or sit in the raid rotation. | As R19. |
| ROGUE-R21-S5 | `rogue-combat.json` Adrenaline Rush note; combat.md:87; assassination.md:84 | Carried R19-S3 ("with your racial": Undead has no cooldown racial to line up). | As R19-S3. |
| ROGUE-R21-S6 | Subtlety Eviscerate and Slice and Dice notes at bands 10 to 16; Sinister Strike 60-energy gate | Carried R19-S4 to S6. | As R19. |

## Prose, line by line (all four pages)

- **index.md** :22 to :37. :27 above. Redirect level 1, 15 s storage (438040); Poised 18, 10 energy (1328825); Shank 24 / 44, 5 energy, 15 / 25 back, 10 s, cannot miss (1329283 / 1329549); Subtlety Shank at three or fewer with Boot Knife's two points (four would overflow five). Races: Touch of the Grave, Blood Fury 10% / 15 s / 2 min, Berserking 10% / 15 s / 3 min, Elune's Light 10% / 15 s / 3 min, Ruthless 1%, Artisan's Precision 1% hit, Eureka! (non-periodic damaging, class rows), Expansive Mind 5% energy, Will of the Forsaken 2 min, Wind Blessed 2% (`races.json`); order matches `round3/race-sims.md` rogue rows (Night Elf / Human / Dwarf trailing in Subtlety; Gnome for Combat; Undead then Orc then Troll). Skyborne only as the paid option. OK.
- **assassination.md** :30 to :133. Builds: leveling 38/11/2, raid 36/10/5; the trade at :52 is exact (raid +Vigor 1, Blood Splatter 1, Improved Slice and Dice 1, Opportunity 4 = 7 against leveling +Improved Gouge 3, Remorseless 2, Precision 1, Master of Deception 1 = 7). Mutilate 90% plus 79.2 at rank 3, 60 energy, two points, below Venom (prereq 105712). Seal Fate 100% at 3/3. Relentless Strikes 20% a point, 25 energy. Cold Blood 0 cost, 3 min (14177). Lethality 220% (bonus rule). Vile Poisons 20%, Improved Poisons, Coup de Grace 20%. Opportunity 20% to Mutilate and Garrote; Blood Splatter 6 / 9 s; Vigor 5; Coup de Grace 3 plus Lightning Reflexes 2 = the five points to reach row two. Ambush gets none of Lethality (not on its list). Ruthlessness 3 points in both builds. Kidney Shot 25 / 20 s / 6 s at five (rank 2, 8643). Thistle Tea item 7676 exists. Levels 1 to 19 against the leveling order: :121 is R21-1; :123 Poised 10; :125 Improved Gouge 1.5 s, Remorseless 40%, Malice 5% hold. Levels 20 to 60: Shank 24, Mutilate 40, Cheap Shot two points, Instant 20, Deadly 30, Venom 30, Relentless Strikes 39 to 44 on the leveling order. OK apart from R21-1.
- **combat.md** :30 to :132. Builds: leveling 14/34/3, raid 12/34/5; the trade at :52 is exact (leveling +Improved Gouge 3, Remorseless 2, Ruthlessness 1, Efficient Strikes 1, Hack and Slash 1, Lightning Reflexes 1, Master of Deception 3; raid +Lethality 2, Improved Slice and Dice 2, Flash of Steel 1, Composure 2, Opportunity 5). Adrenaline Rush 100% / 15 s / 5 min; Blade Flurry 20%, 25 energy, 2 min (two in a 3 minute kill); Flash of Steel 15 / 2 s at five; Weapon Expertise 2%; Dual Wield Specialization 50%; Puncturing Wounds 30% / 45%; Aggression 15% to Sinister Strike and Eviscerate only; Flawless Execution 15; Efficient Strikes 3 at one point; Composure makes Poised 20, Garrote 50 to 30; Comp 2 + Lightning Reflexes 4 + Coup de Grace 3 + Puncturing Wounds 3 + Precision 3 = 15 for row four. Levels 1 to 19 (Kick 12, Garrote 14, Ambush 18; Coup de Grace 10-12, Efficient Strikes 13, Lightning Reflexes 14-18, Puncturing Wounds 19) and 20 to 60 (Cheap Shot 26, Shank 24, Aggression by 30, Improved Gouge and Remorseless 44-48) match the leveling order. OK.
- **subtlety.md** :29 to :133. 12/8/31 exact (Malice 5, Ruthlessness 3, Improved Slice and Dice 2, Blood Splatter 1, Lethality 1; Coup de Grace 3, Efficient Strikes 2, Precision 3). Hemorrhage 145% / 15% / 35 energy (16511); Serrated Blades 9% / 30%; Boot Knife 100% at 5/5; Death Mark 25 energy / 10 s / 6 s / twice for bleeds at every rank (1316709, 1325756 to 1325758); Ghostly Strike 125 / 180%, 40 energy, 7 s, 20 s (14278); Premeditation 2 points / 20 s / 2 min, no Stealth in the text; Initiative 100%; Dirty Deeds 20 and frontal Garrote (Garrote 30 after Dirty Deeds); Cheap Shot 2 points plus Initiative = 3, 4 s stun; Preparation 10 min; Blind 5 min less Elusiveness 90 s = 3.5 min. Levels: Ambush 60 less Poised 10 = 50; Master of Deception three levels; Opportunity 13-17; Dirty Deeds 18-19; Premeditation 20, Ghostly Strike 26, Hemorrhage 30, Boot Knife 35, Death Mark 40, Camouflage 24. OK.
- **Shapes from round 20 checked on every line:** raid-order early credit (only the Assassination and Combat pages ship a raidBuild; every level stated in their prose is in the Leveling section and matches the leveling order, and the talent sections describe 60); superlatives and comparisons (the grep of "hardest / best / only / cheapest / most / beats / harder" over all four pages turned up R21-1 only; Combat's "more than half your Backstabs crit" is a raid 60 claim and holds with Puncturing Wounds 30, Malice 5, Hack and Slash 4 on top of gear crit); quoted values a talent changes (Garrote 30 with Composure and with Dirty Deeds, Blind 3.5 min with Elusiveness, Shank 40 with five Flash of Steel, Hemorrhage 30 with both Efficient Strikes all stated with their talent; nothing in the builds changes Kidney Shot, Cold Blood, Adrenaline Rush, Blade Flurry, Premeditation or Preparation).
- **Build names:** no rogue page names 70755; only the APL root comments (provenance, not rendered) do.
- **Voice:** no em or en dashes, no site, sim or band talk, no nightly numbers.
- **Source links:** unchanged since round 4's check (two Blizzard articles, the Wowhead Skyborne post, the site's own pages); not re-fetched.

## Decisions

- **D1.** R21-1 is CONFIRMED under round 20's shape (2): a superlative over a span (levels 1 to 19) that fails at 18 and 19, where the same paragraph trains Ambush and the next one has you open with it. The fix uses the comparison that holds at every level and that the Backstab card already makes.
- **D2.** index.md:27 is not a finding. "For Assassination until 60" describes the pair a player can hold, and from 20 to 29 there is only one poison to place. Wording only (S1).
- **D3.** Round 20's four fixes all hold against the client, the engine's additive poison chance and both orders.
- **D4.** Nothing carried from the ledger's Round 19 and Round 20 sim and engine lists is re-filed: no sim round 21 has landed, and no page now states any of them falsely.

## Sims requested

- None new. Carried (already routed): Assassination poison hands at 34 and 38 on the leveling order (decides S2); Subtlety Shank gate without Boot Knife at 24 to 32; Subtlety bands 10 to 16 finisher fallback (SIM-R11-1); Subtlety band 20 weapon (K4); rogue BiS and race re-runs at the re-rank; X2 re-check on the new BiS.

## Sim findings

- None new. Carried: K1 (Death Mark storage and the Mutilate / Shank off-hand damage, waiting on a level 40+ log).

## Engine questions

- None new. Carried (ledger Round 20): whether Vile Poisons, Venom and Malice scale the new attack-power share of each Instant hit and Deadly tick. The pages state nothing about it.

## Sources

- Client 1.60.1.70790: `spelltips/rogue.json` (Backstab 53, 2589, 2590; Ambush 8676; Sinister Strike 1752 to 1759; Instant Poison 8679 to 11340; Deadly Poison 2823 to 25351; Shank 1329283, 1329549; Venom 1329328; Mutilate 399956, 1241582, 1241584; Death Mark 1316709, 1325756 to 1325758; Ghostly Strike 14278; Hemorrhage 16511; Kidney Shot 408, 8643; Blade Flurry 13877; Adrenaline Rush 13750; Cold Blood 14177; Premeditation 14183; Preparation 14185; Blind 2094; Cheap Shot 1833; Garrote 703 to 11290; Poised 1328825; Redirect 438040); `talents/rogue.json`; `trainables/rogue.json`; `races.json`; `items.json` (Thistle Tea 7676); `addon-builds.json` (both orders, three specs); `addon-data.json` (`rotations/rogue-*`, every band).
- Engine d5c708834: `sim/rogue/poisons.go`:37-50, `sim/rogue/client_values_test.go`:119-136.
- `design/reviews/2026-11-04-sim-audit-round20/rogue.md` (X1, X2); `design/reviews/guide-accuracy/round20/rogue.md`; `round3/race-sims.md`; `ledger.md` Rounds 19 and 20.
- Blizzard, "WoW Forever Beta Development Notes" (Kaivax, 3 November 2026), Rogue poison attack-power coefficients: https://us.forums.blizzard.com/en/wow/t/wow-forever-beta-development-notes-%E2%80%93-updated-november-2/2360696 (cited by round 20; the pages quote no coefficient).

## Appendix: Note sweep

One line per row note: spec, row, bands it renders on, each named talent, spell, rank or value with the level the player has it (L = leveling order, R = raid order), and the verdict.

**Assassination** (23 bands, 10 to 60; both orders)
1. Shank (24-60): rank 1 at 24 (5 for 15), rank 2 at 44 (5 for 25): "pays back more than it costs" at both. Cannot miss (client text). Venom L30 / R30, conditional ("with Venom"). Vile Poisons in before Venom on both orders (L27, R25). 30% / 6 s / "land more often" (+20) from 1329328. OK.
2. Slice and Dice (10-60): rank 1 at 10, rank 2 at 42; four or more, under 3 s. OK.
3. Rupture (20-60): trained at 20. Blood Splatter conditional (R22; never on L). +6 s (1329256). Ticks crit (owner ruling). OK.
4. Cold Blood (25-60): talent L33 / R31; rows 25 to 32 render early (component lane). No cost, off the global (14177). OK.
5. Eviscerate (10-60): four or more. Mutilate from 40 (L40 / R40). Seal Fate from 45 (L45 / R43). OK; S3 (R under-credited at 44).
6. Mutilate (40-60): ranks at 40 / 50 / 60; both weapons, two points, +20% poisoned at every rank. OK.
7. Backstab (10-60): rank 1 at 4 to rank 8 at 60. "Until Mutilate" (40). "Far harder than Sinister Strike": 150% plus 10 to 140 against 100% plus 3 to 68 at every band. OK.
8. Sinister Strike (10-60): "before Mutilate"; waits for 60 energy (Backstab 60 at every rank). OK.

**Combat** (22 bands, 10 to 60; both orders)
9. Shank (24-60): 10 s base at both ranks; "every 8 with five points of Flash of Steel" (R five at 43; L stops at four, conditional). Pays back more than it costs. OK; S4.
10. Eviscerate (10-60): at five. OK.
11. Slice and Dice (10-60): one or two points; five to Eviscerate. OK.
12. Backstab (10-60): Puncturing Wounds "once you have" (L19 / R17); "in the raid build Lethality from 54 and Opportunity from 56" (R54 / R56; never on L). OK (round 20 fix holds).
13. Adrenaline Rush (40-60): talent L40 / R40. "With your racial" S5. OK.
14. Sinister Strike (10-60): only when Backstab cannot be cast. OK.

**Subtlety** (22 bands, 10 to 60; one order)
15. Shank (24-60): at three or fewer; cannot miss; pays back more than it costs at both ranks; Boot Knife from 35 "can add" two, every Shank at 5/5 from 39. OK.
16. Ghostly Strike (10-60): talent at 26 (rows 10 to 24 early, component lane). 40 energy (14278) against Hemorrhage 35 (16511), 30 with both Efficient Strikes (56-57). 125 / 180% against 100 / 145%. One point. OK.
17. Garrote (14-60): trained at 14. Ghostly Strike (26) lets it in without Stealth. Opportunity from 13. Hemorrhage (30) and Death Mark (40) "once you have them". OK.
18. Rupture (20-60): trained at 20. Serrated Blades from 21. Hemorrhage's debuff from 30. OK.
19. Death Mark (40-60): ranks at 40 / 48 / 56 / 60; 6 s, twice for Garrote and Rupture, 10 s cooldown at every rank. OK.
20. Eviscerate (10-60): three or more while Rupture ticks, before Death Mark (40). OK; S6 at 10 to 16.
21. Slice and Dice (10-60): three or more while Rupture ticks. OK; S6 at 10 to 16.
22. Hemorrhage (30-60): talent at 30. One point; 15% to Rupture and Garrote; hold rules match the card condition (75 energy, under 1 s). OK.
23. Backstab (10-60): until Hemorrhage (30). Opportunity from 13 at 4% a point (14057). OK.
24. Sinister Strike (10-60): before Hemorrhage; 60-energy gate. OK; S6.

24 notes swept: 24 OK, 0 findings. The one finding this round is page prose (R21-1).
