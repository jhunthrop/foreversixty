# Guide accuracy round 20: Rogue

Audited on the round 20 audit tree (4234643f5: main after guide round 19 and sim round 20; engine pin d5c708834), against the active client 1.60.1.70790 (`data/builds/1.60.1.70790/`). Files: `web/src/content/guides/rogue/{index,assassination,combat,subtlety}.md`, every line; every note in `data/curated/apl/rogue-{assassination,combat,subtlety}.json`, every sentence, at every band it renders in `addon-data.json` (`rotations/rogue-assassination` 23 bands, `rogue-combat` 22, `rogue-subtlety` 22), read against BOTH shipped orders in `addon-builds.json` (leveling and raid) and the trainer levels in `trainables/rogue.json`. Read first: round 19 rogue report, sim round 20 `rogue.md` and `integration.md`, ledger Round 19.

## Changes since round 19's tree (6beb919c2)

- APL: round 19's two fixes only (Assassination Eviscerate note, Subtlety Shank note). Sim round 20 changed no rogue card action or note.
- Guides (sim round 20, X1/X2): poisons now scale with attack power (stat priority and consumables lines on all three specs), and Assassination swaps to Instant Poison on the main hand and Deadly on the off hand at 60 (Poisons, Best in slot, Enchants and consumables, Levels 20 to 60).
- The attack-power claim is sourced: Blizzard's development notes post (3 November) reads "Deadly Poison Attack Power coefficient increased to 0.009 (Was 0.0045)" and "Instant Poison Attack Power coefficient increased to 0.01 (Was 0.005)" (read this round, Sources). The client rows show 0, so the post is the evidence. The prose quotes no number. Holds.

## Shipped orders used (index 0 = level 10)

- Assassination leveling: Improved Gouge 10-12, Remorseless Attacks 13-14, Malice 15-19, Ruthlessness 20-22, Improved Slice and Dice 23-24, Lethality 25-26, Vile Poisons 27-29 and 31-32, Venom 30, Cold Blood 33, Improved Poisons 34-38, Relentless Strikes 39 and 41-44, Mutilate 40, Seal Fate 45-47, Coup de Grace 48-50, Lightning Reflexes 51-52, Puncturing Wounds 53-55, Precision 56-58, Master of Deception 59, Opportunity 60.
- Assassination raid: Malice 10-14, Ruthlessness 15-17, Improved Slice and Dice 18-20, Vigor 21, Blood Splatter 22, Lethality 23-24, Vile Poisons 25-29, Venom 30, Cold Blood 31, Improved Poisons 32-36, Relentless Strikes 37-39 and 41-42, Mutilate 40, Seal Fate 43-45, Coup de Grace 46-48, Lightning Reflexes 49-50, Puncturing Wounds 51-53, Precision 54-55, Opportunity 56-60.
- Combat leveling: Coup de Grace 10-12, Efficient Strikes 13, Lightning Reflexes 14-18, Puncturing Wounds 19-21, Precision 22-24, Flawless Execution 25, Weapon Expertise 26-27, Aggression 28-30, Dual Wield Specialization 31-32, Blade Flurry 33, Hack and Slash 34-38, Flash of Steel 39 and 41-43 (4), Adrenaline Rush 40, Improved Gouge 44-46, Remorseless Attacks 47-48, Malice 49-53, Ruthlessness 54-56, Improved Slice and Dice 57, Master of Deception 58-60.
- Combat raid: Coup de Grace 10-12, Lightning Reflexes 13-16, Puncturing Wounds 17-19, Precision 20-22, Composure 23-24, Flawless Execution 25, Weapon Expertise 26-27, Aggression 28-30, Dual Wield Specialization 31-32, Blade Flurry 33, Hack and Slash 34-37, Flash of Steel 38-39 and 41-43 (5), Adrenaline Rush 40, Malice 44-48, Ruthlessness 49-50, Improved Slice and Dice 51-53, Lethality 54-55, Opportunity 56-60.
- Subtlety (leveling and raid orders identical): Master of Deception 10-12, Opportunity 13-17, Dirty Deeds 18-19, Premeditation 20, Serrated Blades 21-23, Camouflage 24 and 27-28, Elusiveness 25 and 33, Ghostly Strike 26, Initiative 29 and 31-32, Hemorrhage 30, Preparation 34, Boot Knife 35-39, Death Mark 40, Malice 41-45, Ruthlessness 46-48, Improved Slice and Dice 49-50, Blood Splatter 51, Lethality 52, Coup de Grace 53-55, Efficient Strikes 56-57, Precision 58-60.
- Trainer: Instant Poison 20, Shank 24 / 44, Cheap Shot 26, Deadly Poison 30, Mutilate 40 / 50 / 60, Death Mark 40 / 48 / 56 / 60.

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| ROGUE-R20-1 | CONFIRMED | consumable (page contradicts a spec page) | index.md:27 | "Deadly Poison on the main hand and Instant Poison on the off hand is right for all three trees." | Sim round 20 X2 (Instant MH / Deadly OH +1.5 to +1.7% at raid 60, both seeds and factions) moved Assassination at 60 to Instant on the main hand, and assassination.md:88, :103 and :107 now say so. The class page still says the old pair is right for all three trees. | "**Poisons and stones stack.** A sharpening stone sits on a weapon beside a poison. Deadly Poison on the main hand and Instant Poison on the off hand is right for Combat and Subtlety, and for Assassination until 60, where it swaps them." |
| ROGUE-R20-2 | CONFIRMED | talent / consumable (new work made a line false) | assassination.md:56 | Venom "... 20 percentage points to each poison's chance to land for 6 seconds, so Instant Poison procs on half your off-hand hits." | The Talents section describes the 60 raid build, and at 60 the same page puts Instant Poison on the main hand and Deadly on the off hand (:88, :107). The off hand no longer carries Instant. The arithmetic stays right (Instant Poison 20%, 11338; Improved Poisons 5/5 +10%, 14113; Venom +20%, 1329328: 50%), but on main-hand hits. | "... for 6 seconds, so Instant Poison procs on half your main-hand hits at 60." |
| ROGUE-R20-3 | CONFIRMED | rotation (same cause) | assassination.md:78 | "with Venom, for the next 6 seconds your poisons get another 30% on top of Vile Poisons and Instant Poison lands on half your off-hand hits." | As R20-2. The priority section is the 60 raid rotation, where Instant is on the main hand. | "... and Instant Poison lands on half your main-hand hits." |
| ROGUE-R20-4 | CONFIRMED | card note (talent credited early on the raid order) | `rogue-combat.json` Backstab note (priority line 4), on the Combat Backstab row at bands 10 to 60 | "... and in the raid build Opportunity and Lethality land on it too." | The page ships a raid build (combat.md:7) whose order a player can level with: Lethality 54-55, Opportunity 56-60. On bands 10 to 52 a player on that order has neither (Lethality 14128, Opportunity 14057 both name Backstab, so the talents are right; the timing is not). The leveling order never takes either, which the "in the raid build" qualifier already covers. This is the druid round 20 pattern: an error only the raid order exposes. The first clause is untouched, so the addon row does not move. | "Backstab from behind with a dagger in the main hand: once you have Puncturing Wounds it crits more often and can give a second combo point, and in the raid build Lethality from 54 and Opportunity from 56 land on it too." |

Counts: CONFIRMED 4, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0.

## Shoulds (wording, not findings)

| id | file:line | note | suggested wording |
|---|---|---|---|
| ROGUE-R20-S1 | `rogue-assassination.json` Eviscerate note, band 44 | "from 45 a Seal Fate crit makes it three" holds on the leveling order (Seal Fate 45-47). On the raid order Seal Fate starts at 43, so at band 44 a raid-order player already has 2/3. The note under-credits; it states nothing false. | "From 40 Mutilate gives two points a cast, and once Seal Fate is in (from 43 to 45) a crit can make it three, ..." |
| ROGUE-R20-S2 | assassination.md:88 | "Below 60, keep Deadly on the main hand ... at 30 to 39 Backstab stacks Deadly faster": from 20 to 29 there is no Deadly Poison (trains at 30). Sim X2 measured only bare 30 for the "30 to 39" range (73.2 against 75.4). Nothing false. | "From 30 to 59, keep Deadly on the main hand and Instant on the off hand: ..." (and see Sims requested). |
| ROGUE-R20-S3 | `rogue-combat.json` Shank note | Carried R19-S1 (Flash of Steel 4 points on the leveling order). The raid order reaches five points at 43, so "every 8 with five points" is true there. | As R19-S1. |
| ROGUE-R20-S4 | combat.md:77 | Carried R19-S2 ("40 back with Flash of Steel"). | As R19-S2. |
| ROGUE-R20-S5 | `rogue-combat.json` Adrenaline Rush note; combat.md:87; assassination.md:84 | Carried R19-S3 ("with your racial": the Undead racial has no cooldown to line up). | As R19-S3. |
| ROGUE-R20-S6 | Subtlety Eviscerate and Slice and Dice notes at bands 10 to 16; Combat Sinister Strike 60-energy gate; R18-S3 items | Carried R19-S4 to S6. | As R19. |

## Prose, line by line (all four pages)

- index.md: :22 to :37 re-read. Only :27 (R20-1) is wrong. Shank 24 / 44, 25 at rank 2, 10 s (1329283 / 1329549); Redirect 15 s; Poised 18; races as measured.
- assassination.md: the new :88, :99, :103, :107 and :133 lines hold apart from S2 (Deadly at 30, Instant at 20, swap at 60, AP scaling sourced). The older lines hold against the client: Mutilate 90% plus 79.2 at rank 3 (1241584), Seal Fate 100% at 3/3, Relentless Strikes, Cold Blood 14177 (no cost, 3 min; the list includes Mutilate and Eviscerate), Lethality 220%, Opportunity 20% on Mutilate and Garrote, Blood Splatter 6 / 9 s, Coup de Grace 20%, levels 1 to 19 against the leveling order. The exceptions are :56 and :78 (R20-2, R20-3).
- combat.md: the new :100 and :108 lines hold (Combat keeps Deadly on the main hand, sim X2: 798.2 against 792.2 mirrored, 767.1 two Instants). :104 "the off hand carries Instant Poison" holds. Talent counts (Hack and Slash 4 in raid, Flash of Steel 4 in the level 60 build, Composure 20 off Poised), levels 1 to 19 and 20 to 60 (Aggression by 30, Improved Gouge and Remorseless 44 to 48, Cheap Shot 26) all hold.
- subtlety.md: the new :101 and :109 lines hold (Deadly on the main hand ties two Deadlys and beats the mirror, X2). :105 "your Instant Poison" on the off hand holds. Hemorrhage 145% / 15%, Ghostly Strike 180% / 7 s / 20 s, Boot Knife at 5 points, Death Mark 25 energy / 10 s / 6 s / twice for Garrote and Rupture at every rank (1316709, 1325756 to 1325758), Serrated Blades 9% / 30%, Efficient Strikes 5 at two points, Dirty Deeds 20, Initiative 100%, Premeditation 20 s: all hold, and every level 1 to 60 line matches the order.
- Voice: no em or en dashes, no site, sim or band talk, no nightly numbers in the new prose.

## Decisions

- **D1.** R20-1 to R20-3 are CONFIRMED and not routed to the sim lane. Sim round 20 measured and shipped the Assassination hand swap and rewrote the Assassination page around it, but left the Venom lines (:56, :78) and the class page (:27) describing the old hands. A page that contradicts itself or its spec page counts under the round 14 rule.
- **D2.** R20-4 is CONFIRMED under this round's "both orders" lens. "In the raid build" names the build, not the level, and the raid order gives Lethality only at 54 and Opportunity only at 56, while the row renders from 10.
- **D3.** Assassination Eviscerate "from 45" is a Should (S1). It names the later of the two orders' levels, so it never credits Seal Fate early.
- **D4.** The poison attack-power prose is accepted on the Blizzard post. The guides quote no coefficient, which keeps them clear of nightly numbers.

## Sims requested

- **Assassination poison hands at 34 and 38 on the leveling order** (optional, decides S2's "30 to 39"): Deadly MH / Instant OH against Instant MH / Deadly OH, seeds 1 and 50001. Sim X2 measured bare 30 only.
- Carried: Subtlety Shank gate before Boot Knife at 24 to 32 (R19); Subtlety bands 10 to 16 finisher fallback (SIM-R11-1); Subtlety band 20 weapon (K4); rogue BiS and race re-runs at the re-rank; X2 re-check on the new BiS.

## Sim findings

- None new. Carried: K1 (Death Mark storage and the Mutilate / Shank off hand, waiting on a level 40+ log); the doubled poison coefficients on a live log.

## Engine questions

- Do Vile Poisons, Venom's 30% and the poison crit from Malice ("all attacks and Poisons") multiply the new attack-power share of each Instant hit and Deadly tick, as they do the client base? The guide claims nothing about it, so this is for the sim round 21 integrator only.

## Sources

- Blizzard, "WoW Forever Beta Development Notes" (Kaivax, post of 3 November 2026), Rogue: "Deadly Poison Attack Power coefficient increased to 0.009 (Was 0.0045)", "Instant Poison Attack Power coefficient increased to 0.01 (Was 0.005)": https://us.forums.blizzard.com/en/wow/t/wow-forever-beta-development-notes-%E2%80%93-updated-november-2/2360696 (read 2026-11-05).
- Client 1.60.1.70790: `talents/rogue.json` (Venom 1329328, Improved Poisons 14113, Lethality 14128, Opportunity 14057, Seal Fate 14186, Flash of Steel 1329284, Boot Knife 1329286, Efficient Strikes 13732, Hemorrhage 16511, Ghostly Strike 14278, Death Mark 1316709); `spelltips/rogue.json` (8679, 11338, 2823, 25351, 1329283, 1329549, 399956, 1241582, 1241584, 1316709, 1325756 to 1325758, 14177, 13750, 11281, 11294, 11290, 5171, 6774); `trainables/rogue.json`; `addon-builds.json` (both orders, all three specs); `addon-data.json` (`rotations/rogue-*`, every band).
- `design/reviews/2026-11-04-sim-audit-round20/rogue.md` (X1, X2), `integration.md`; `design/reviews/guide-accuracy/round19/rogue.md`; `ledger.md` Round 19.

## Appendix: Note sweep

One line per row note: spec, row, bands it renders on, each named talent, spell, rank or value with the level the player has it (L = leveling order, R = raid order), and the verdict.

**Assassination** (rows 10 to 60; raid order as well as leveling)
1. Shank (24-60): cannot miss and pays back more than it costs (rank 1 at 24, 5 for 15; rank 2 at 44, 5 for 25). Venom (L30 / R30) is conditional. "On top of Vile Poisons": Vile Poisons is in before Venom on both orders (L27, R25). 30% and 6 s from 1329328. "Land more often" is Venom's +20%. OK.
2. Slice and Dice (10-60): four or more points, under 3 s. Trained at 10. OK.
3. Rupture (20-60): Rupture is trained at 20. "Only with Blood Splatter (the raid build)" is conditional (R22; never on L). +6 s from 1329256. Ticks crit (owner ruling). OK.
4. Cold Blood (25-60): talent at L33 / R31. Rows 25 to 32 render early (component lane, R19 D5). No cost and off the global (14177; engine). "Right before the Eviscerate". OK.
5. Eviscerate (10-60): four or more. Mutilate "from 40" (L40 / R40). Seal Fate "from 45" (L45 / R43). OK; S1 (raid order under-credited at 44).
6. Mutilate (40-60): rank 1 at 40, 2 at 50, 3 at 60. Both weapons, two points, +20% poisoned at every rank. OK.
7. Backstab (10-60): rank 1 at 10 (trained at 4). "Until Mutilate" (40). Harder than Sinister Strike at every rank (150% plus flat, against 100% plus flat). OK.
8. Sinister Strike (10-60): "before Mutilate" plus the 60-energy sentence (Backstab costs 60). OK.

**Combat** (rows 10 to 60; raid order as well as leveling)
9. Shank (24-60): 10 s base. "Every 8 with five points of Flash of Steel" holds on R (five points at 43); L stops at four. OK; S3.
10. Eviscerate (10-60): at five. OK.
11. Slice and Dice (10-60): one or two points, and five always to Eviscerate. OK.
12. Backstab (10-60): Puncturing Wounds "once you have" (L19 / R17). "In the raid build Opportunity and Lethality" (R56 / R54; never on L). **ROGUE-R20-4.**
13. Adrenaline Rush (40-60): talent at L40 / R40. "With your racial" is S5. OK.
14. Sinister Strike (10-60): "only when Backstab cannot be cast". OK.

**Subtlety** (rows 10 to 60; one order)
15. Shank (24-60): at three or fewer points; cannot miss; pays back more than it costs. Boot Knife from 35 "can add" two points, every Shank once 5/5 at 39 (1329286 at 20% a rank). OK (round 19 fix holds).
16. Ghostly Strike (10-60): talent at 26 (rows 10 to 24 early, component lane). 40 energy (14278). Hemorrhage 35 (16511). "30 with both points of Efficient Strikes" (56-57, rank 2 is 5). Harder than Hemorrhage (125 / 180% against 100 / 145%). One combo point. OK.
17. Garrote (14-60): trained at 14. "Ghostly Strike lets it in without Stealth" (26). Opportunity from 13. Hemorrhage (30) and Death Mark (40) "once you have them". OK.
18. Rupture (20-60): trained at 20. Serrated Blades "from 21" (21-23). Hemorrhage's debuff "from 30". OK.
19. Death Mark (40-60): ranks at 40, 48, 56 and 60. 6 s, twice for Garrote and Rupture, 10 s cooldown at every rank (1316709, 1325756 to 1325758). OK.
20. Eviscerate (10-60): three or more while Rupture ticks, "only before you have Death Mark: from 40". OK.
21. Slice and Dice (10-60): three or more while Rupture ticks. OK; carried S6 for bands 10 to 16.
22. Hemorrhage (30-60): talent at 30. One point. 15% to Rupture and Garrote. Hold at Ghostly Strike under 1 s unless 75 energy or more, and hold after Ghostly Strike for the Garrote (matches the card condition). OK.
23. Backstab (10-60): "until Hemorrhage" (30). Opportunity "from 13" at 4% a point (14057). OK.
24. Sinister Strike (10-60): "before Hemorrhage" plus the 60-energy sentence. OK.

24 notes swept: 23 OK, 1 finding (R20-4).
