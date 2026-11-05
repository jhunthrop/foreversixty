# Guide accuracy round 21: Hunter

Audited on the round 21 audit tree (main after guide round 20, head e039e1490; no sim round 21 commit in the log, so there is no sim round 21 report to read; engine pin d5c708834 per `sim/enginever/version.go`, source in gr21/engine), against the active client 1.60.1.70790 (`data/builds/1.60.1.70790/`). Files: `web/src/content/guides/hunter/{index,beast-mastery,marksmanship,survival}.md`, every row note in `data/curated/apl/hunter-{beast-mastery,marksmanship,survival}.json`, every rendered row of `addon-data.json` `rotations/hunter-*` (BM 18 bands, MM 19, SV 20), the six orders in `addon-builds.json`, `talents/hunter.json`, `trainables/hunter.json`, `spellconst/hunter.json`, `spelltips/hunter.json`, `gametables/basemp.txt`, `races.json`, `items.json`. Read first: round 20 hunter report, the ledger's Round 19 and Round 20 sections, sim round 20 `hunter.md`, and the briefs back to round 1.

## Changes since round 20's tree (4234643f5)

`git diff 4234643f5 HEAD` touches only five hunter notes, all round 20's fixes: BM Aimed Shot (R20-2), BM and MM Serpent Sting (R20-3 wording), BM and MM Arcane Shot (R20-1). No hunter prose, order, rank row or APL condition moved. Each fix is re-checked below.

## Round 21 lens: every superlative and quoted cost

Superlatives and comparisons found on the four pages and the cards (grep for best, biggest, most, hardest, worst, only, more than any, far more, than anyone, plus a line read), each checked at every band it renders and against the other specs:

| line | claim | verdict |
|---|---|---|
| index:33 | Aimed Shot outhits Multi-Shot | OK: both normalized weapon damage, Aimed adds 40 to 350 (19434-20904), Multi-Shot adds 0 (2643); Barrage scales both |
| BM:75 | Focused Fire the only flat damage bonus to your own shots in the BM tree | OK: no other BM talent raises the hunter's damage |
| BM:93, MM:98 | Auto Shot does more damage than any button you press | OK: sim round 20 B1 38% (next button 13%), M1 46.5% (next 14%) |
| BM:99, MM:105 | Multi-Shot the worst damage for its mana; sting best, Arcane next | OK in the raid context of the section (top rank Aimed): B3 BM 6.4/4.5/3.8/3.1, MM 8.8/6.5/3.5/3.4; Sniper Shot (365 flat, W+590) sits between Arcane Shot and Multi-Shot |
| BM:113, MM:119 | a missed auto shot is the most expensive thing that can happen to you | **HUNTER-R21-1** |
| BM:153 | Beast Mastery is the best way to level a hunter | OK: Warcraft Tavern's Forever leveling guide (a page source) recommends BM as the smoothest; the sim is level at bare 40/50 but cannot model the pet tanking |
| BM:157 | rank 1 still hits for well over half of what the newest rank does | OK, narrowly: about 56 to 63% at 52 to 60 on the bands' gear (W about 300 to 400 against +260/+350); Should below |
| BM:159, MM:165, SV:144 | mail at 40 the biggest armor upgrade | OK: leather to mail is the largest armor-class step a hunter gets |
| MM:60 | output mostly auto shot, so movement hurts you more than anyone | **HUNTER-R21-2** |
| MM:82, MM:88, MM Sniper note | Sniper Shot the hardest hit / biggest shot | OK at every band it renders: r1 +320 (40), r2 +450 (48), r3 +590 (58) against the rendered Aimed rank 1 (+40, Barrage 5%), rank 4 at 50 (+185) and rank 6 at 60 (+350 x 1.15; Sniper leads while normalized weapon damage is under 1250) |
| MM:88 | Sniper Shot the capstone | OK: tier 6 of the Marksmanship tree |
| SV:72 | Expose Prey the only way Mongoose Bite opens on a boss not attacking you | OK: 1495 "only after you dodge" |
| SV:86 | from 30 to 59 Beast is the only melee aspect | OK: Beast 30, Falcon 60, Monkey gives dodge |
| SV:90 | the only mistake is letting Raptor Strike sit unqueued | OK (rhetorical, true of the button) |
| BM Multi-Shot and MM Multi-Shot notes | the worst damage for its mana you fire | OK at 18 to 60: 13.9% of base mana (54 at 18 to 239 at 60, `basemp.txt`), normalized weapon damage only |
| BM and MM Arcane Shot notes (R20-1 fix) | far more for its mana than Multi-Shot | OK at 18 to 60: about 2x at 18 to 20 (r2/r3 35/50 mana, arcane, plus 15% RAP, against 54 to 63 mana of armored weapon damage), about 1.5x at 40 and raid 60 |
| BM and MM Serpent Sting notes (R20-3 fix) | more for its mana than Arcane Shot or Multi-Shot | OK at every band: per cast the sting carries 25% RAP to Arcane Shot's 15% and has the higher base per mana at every rank pair (r1/r1 15/15 against 25/25, r3 70/50 against r3 50/50, r8 500/230 against r8 280/190) |
| BM Aimed Shot note (R20-2 fix) | a small extra over Multi-Shot; Multi-Shot fine above a quarter mana | OK: matches the APL gate and BM:91 |

Quoted costs: the pages quote no numeric cost. Cost comparisons checked: BM:157 "far less mana" (75 against 115 to 310), MM:163 "a fraction of the mana", SV:146 (Immolation Trap 135 to 190 at 46; Raptor Strike rank 5 55 against rank 6 70, +35 against +40), SV:82 (Efficiency on top of Resourcefulness), BM:105/MM:111 Volley "a lot of mana" (350/420/490). All OK. Strider Kick's "sliver of mana" is a Should.

Levels 1 to 19 (all three pages) re-read against `trainables/hunter.json` and the orders: Heart of the Lion 1; Serpent Sting and Aspect of the Monkey 4; Arcane Shot and Hunter's Mark 6; Concussive Shot 8; Aspect of the Hawk 10; Wing Clip and Mend Pet 12; Mongoose Bite 16; Multi-Shot 18; Raptor Strike ranks 8 and 16; Track Beasts 1, Humanoids 10, Undead 18; Deadly Aspects 10-14 then Focused Fire 15-16 then Endurance Training 17 (BM); Lethal Attacks then Efficiency from 15, 3% a point (MM); Improved Tracking, Savage Strikes 15-16 (2% a point), Survivalist from 17 (2% a point) (SV). No superlative fails there.

## Findings

| id | status | kind | file | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| HUNTER-R21-1 | CONFIRMED | stats (superlative) | `beast-mastery.md:113`, `marksmanship.md:119` | "because a missed auto shot is the most expensive thing that can happen to you" | Your shots roll the same ranged hit table as Auto Shot, and a single missed special costs more. Aimed Shot rank 6 is normalized weapon damage + 350 for 310 mana (20904; `aimed_shot.go` FlatCost 310), and Sniper Shot is normalized weapon damage + 590 for 365 mana on a 4 second cast that, per MM:88, holds your auto shot twice as long as Aimed Shot (1310786; `sniper_shot.go`). An auto shot is plain weapon damage for no mana. So a missed Aimed Shot or Sniper Shot loses more damage and more mana than a missed auto shot, and the superlative fails on both ranged pages. What is true, and what the page already says at BM:93/MM:98, is that auto shot is the biggest share of your damage. | Both pages: "Hit matters up to 9%, the cap against a raid boss, because every miss throws away a whole shot, auto shots included, and auto shot is more of your damage than any button you press; past the cap it adds nothing." (MM keeps its following Surefooted sentence; drop BM's trailing "and past the cap it adds nothing" since the fix carries it.) |
| HUNTER-R21-2 | LIKELY | voice/superlative | `marksmanship.md:60` | "Your output is mostly auto shot, so movement hurts you more than anyone" | Sim round 20 M1: raid 60 Marksmanship after the attack power shares, Auto Shot 46.5% of damage; after B3 (Arcane Shot fills every gap) the share is lower still, so auto shot is the largest part, not most of it. "More than anyone" also fails against casters, who lose every cast when they move, while MM:103 itself tells you to fire Arcane Shot and Multi-Shot on the run. LIKELY because "mostly" can be read as "mainly"; the fix is true either way. | "Auto shot is the biggest share of your damage and stops whenever you move, so a fight full of movement costs you more than it costs a Beast Mastery hunter, whose pet keeps fighting" |

Counts: CONFIRMED 1, LIKELY 1, NEEDS-SIM 0, NEEDS-OWNER 0.

## Shoulds (wording, not findings)

- **BM:157 "well over half":** at 58 to 60 rank 1 is about 56 to 63% of the newest rank; "more than half" is safe at every band.
- **Notes at bands 10 to 17:** the BM and MM Serpent Sting and Arcane Shot notes compare with Multi-Shot ("does more for its mana than ... Multi-Shot", "far more for its mana than Multi-Shot", BM's "When both are ready, Multi-Shot goes first") in the sim drawer at 10 and 12, before Multi-Shot is trained at 18. The comparisons are true of the spells and gate nothing, so not filed; "from 18" in front of the Multi-Shot clause would close it.
- **SV Strider Kick note, "for a sliver of mana":** 5.81% of base mana is 45 at 30 and 100 at 60, the same as the Raptor Strike rank of that band (45 at 24, 100 at 56). "for little mana once Resourcefulness is in" is exact.
- **Carried from round 20:** MM Serpent Sting and Arcane Shot "the shared cooldown" at 18 and 19; BM Serpent Sting "Multi-Shot goes first" under a quarter mana; SV Mongoose Bite chance wording (5% at 35, 10% at 36); SV Immolation Trap "ranged attack power"; the 3 November Blizzard notes as a page source; and every older carried Should listed in round 20.

## Checked and right this round

- **Round 20's three fixes:** the BM Aimed Shot note now matches the quarter-mana gate (APL row 5 `currentManaPercent > 25%`) and BM:91/99; the Arcane Shot and Serpent Sting notes hold at every band (table above).
- **Talent claims on all pages** re-read against `talents/hunter.json` max-rank text and the decoded builds (BM 31/20/0 leveling and raid; MM 15/31/5 leveling, 5/33/13 raid; SV 5/11/35): Mortal Shots 6% of the bonus a point (212%, 224%, 230%), Barrage 5% a point, Improved Stings 13% at 2/3, Rapid Killing a minute and 10% a point, Rapid Recuperation, Lone Wolf 20%, Surefooted 3%, Predator's Edge 230% and 50% off hand, Resourcefulness 60%, Lightning Reflexes 15%, Expose Prey 10%, Lacerating Strikes 40% over 21 s, Trueshot Aura 100 at 60 (20906), Bestial Wrath 50% 18 s 2 min, Rapid Fire 40% 5 min.
- **BM raid move (BM:83):** leveling MM side Efficiency 4, Rapid Killing 1, Mortal Shots 2; raid Efficiency 3, Rapid Killing 0, Mortal Shots 4. Exact.
- **Races:** Blood Fury 10% AP/RAP/SP 15 s 2 min, Berserking 10% 15 s 3 min, Elune's Light 10% crit 15 s 3 min, Command 5% pet damage, Ruthless 1% crit, Artisan's Precision 1% hit, Big Game Hunter (`races.json`).
- **Note sweep:** every row below.

## Decisions

- **D1.** R21-1 is CONFIRMED on client values: one missed Aimed Shot or Sniper Shot costs more than one missed auto shot, so "the most expensive thing" is false; the replacement uses the page's own "more of your damage than any button you press".
- **D2.** R21-2 is LIKELY: 46.5% is a plurality, and "more than anyone" is unprovable; the fix compares only with Beast Mastery, which the page's own pet lines support.
- **D3.** BM:99 and MM:105 sit in the raid rotation (Major Mana Potion, Sniper Shot, top-rank Aimed Shot), where the sting-then-Arcane order is the sim's measured order; the leveling bands' rank 1 Aimed Shot is handled by the leveling paragraphs and round 20's note fixes, so the prose is not filed.
- **D4.** BM:153 rests on Warcraft Tavern's Forever leveling guide ("Beast Mastery is our recommended choice", "the smoothest levelling experience"); kept.
- **D5.** Not re-filed: the Survival Hunter's Mark opener (sim lane), and every Sims routed, Sim findings and Engine questions item in the ledger's Round 19 and Round 20 sections, none of which sim round 21 has resolved in this tree.

## Sims requested

1. Marksmanship raid 60 spec breakdown on the shipped list after B3 (guide raid build, frozen BiS, both factions): Auto Shot's share of damage. Settles R21-2's "mostly".
2. Carried from round 20: damage per mana by spell for rank 1 Aimed Shot, Serpent Sting and Arcane Shot (MM 40, 50, 52, 58, bare 60; BM 40, 50, 58); BM and MM at 18 to 20 below a quarter mana with the Multi-Shot readiness check skipped while Multi-Shot is held; the earlier carried race and gear items.

## Sim findings

None new. Round 20's (the 18 and 19 low-mana stall) and the carried items stay open with the sim lane.

## Engine questions

None new. Round 20's two (Immolation Trap on melee or ranged attack power; sting and Arcane Shot share read at cast against per tick) stand.

## Sources

**Client data, 1.60.1.70790:** `spellconst/hunter.json` (19434-20904, 2643, 3044-14287, 1978-25295, 1310687-1310786, 14260-14266, 13795-14305, 1317257, 24118-1299332, 1495-14271, 19306-20910, 2974-14268, 3045, 19574, 1299346-20906, 13165-25296, 13161-1299447, 781-14273, 1510-14295), `spelltips/hunter.json` (same), `talents/hunter.json`, `trainables/hunter.json`, `addon-builds.json`, `addon-data.json` rotations, `gametables/basemp.txt`, `races.json`, `items.json`.

**Engine d5c708834 (gr21/engine):** `sim/hunter/aimed_shot.go`, `sniper_shot.go`, `multi_shot.go`.

**Repository:** `design/reviews/guide-accuracy/round20/hunter.md`, `ledger.md` (Rounds 19 and 20), `design/reviews/2026-11-04-sim-audit-round20/hunter.md` (B1, B3, M1).

**Web (fetched 2026-11-04):** Warcraft Tavern, "Forever Hunter Leveling Guide", https://www.warcrafttavern.com/forever/guides/hunter-leveling/ (recommends Beast Mastery as the smoothest leveling spec; any spec levels).

## Note sweep

One line per row note: spec, row, bands it renders, each named talent, spell, rank or value with the level the player has it, verdict. The raid orders (BM raid, MM raid; SV ships one order) were read beside the leveling orders for every talent a note names.

- BM Hunter's Mark (10-60): spell 6; one mark per target - OK.
- BM Bestial Wrath (40-60): talent 40 on both orders; 50%, 18 s, fear, two minutes (19574) - OK.
- BM Rapid Fire (26-60): spell 26; Multi-Shot 18; Bestial Wrath "from 40" (both orders); five minutes (3045), four from Rapid Killing at 57 (leveling order only; raid order has none, five minutes) - OK.
- BM Aimed Shot (20-60): spell 20; rank 1 at 20-58, rank 6 at 60; small extra over Multi-Shot (+40 to +350 per cast); quarter-mana gate matches the APL - OK (R20-2 fix holds).
- BM Multi-Shot (18-60): spell 18; Aimed Shot 20 in the parenthetical; quarter-mana gate; worst damage per mana at every band - OK.
- BM Serpent Sting (10-60): ranks 2 to 8 by trainer; "from 18" Multi-Shot; 5% RAP a tick; more per mana than Arcane Shot and Multi-Shot at every band - OK; Multi-Shot clauses at 10-12 and "Multi-Shot goes first" under a quarter mana - Shoulds.
- BM Arcane Shot (10-60): ranks 1 to 8; "from 18"; 15% RAP; far more per mana than Multi-Shot at 18-60 - OK; comparison at 10-12 - Should.
- MM Hunter's Mark (10-60): as BM - OK.
- MM Rapid Fire (26-60): spell 26; Multi-Shot 18; Rapid Killing 25-26 on the leveling order, three minutes on every leveling band; raid order has no Rapid Killing, five minutes - OK.
- MM Aimed Shot (20-60): spell 20; rank 1, rank 4 at 50, rank 6 at 60; outhits Multi-Shot at each (Barrage on both: 1 point 35 leveling, 34-36 raid) - OK.
- MM Multi-Shot (18-60): as BM - OK.
- MM Sniper Shot (40-60): raid talent at 40 (leveling-band rows render it: component lane); ranks 40/48/58; 4 s cast, 15 s cooldown (1310687-1310786); hardest hit at every band against the rendered Aimed ranks - OK.
- MM Serpent Sting (10-60): "the shared cooldown" at 18-19 (carried Should); 5% RAP; more per mana than Arcane Shot and Multi-Shot (Improved Stings 15-16 on the raid order only; true without it) - OK.
- MM Arcane Shot (10-60): "the shared cooldown" (carried Should); 15% RAP; far more per mana than Multi-Shot with Barrage and Ranged Weapon Specialization on Multi-Shot - OK.
- SV Raptor Strike (10-60): ranks 2 to 8 (46-47 rank 5 row); no global; Resourcefulness 33-34, 60% at rank 2 - OK (46-47 wording carried Should).
- SV Counterattack (30-60): talent 32 (row at 30: component lane); cannot be dodged, blocked or parried (20910 text) - OK.
- SV Mongoose Bite (16-60): spell 16; 5 s window after a dodge; Expose Prey 35-36, 10 s, any hunter's mark; one press per window - OK (chance wording Should).
- SV Lacerate (30-60): spell 30 (ranks 30/40/50/60); Mongoose Bite 16 - OK.
- SV Strider Kick (30-60): talent 30; 8 s cooldown, 100% weapon damage (1317257); "sliver of mana" - Should.
- SV Immolation Trap (16-60): ranks 16/26/36/46/56; Clever Traps 20-21 and Survivalist's Discipline 37-38 as a "once in" clause; attack power share (Blizzard 16 October) - OK.
- SV Wing Clip (12-60): ranks 12/38/60; Windfury Totem conditional; 5/25/50 damage; main hand hit with procs - OK.
