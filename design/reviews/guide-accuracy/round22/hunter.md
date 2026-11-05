# Guide accuracy round 22: Hunter

Audited on the round 22 audit tree (main after guide round 21, head 987d08a7c; no sim round 21 commit in the log, so there is no sim round 21 report to read; engine source in gr22/engine at the pin in `sim/enginever/version.go`), against the active client 1.60.1.70790 (`data/builds/1.60.1.70790/`). Files: `web/src/content/guides/hunter/{index,beast-mastery,marksmanship,survival}.md`, every row note in `data/curated/apl/hunter-{beast-mastery,marksmanship,survival}.json`, every rendered row of `addon-data.json` `rotations/hunter-*` (BM 18 bands, MM 19, SV 20), the six orders in `addon-builds.json` (decoded level by level), `talents/hunter.json`, `trainables/hunter.json`, `spellconst/hunter.json`, `spelltips/hunter.json`, `gametables/basemp.txt`, `races.json`, `bis/hunter-*.md`. Read first: round 21 hunter report, the ledger's Round 19 to 21 sections, and the briefs back to round 1.

## Changes since round 21's tree (e039e1490)

`git diff e039e1490 HEAD` touches three hunter lines, all round 21's fixes: the BM and MM hit paragraphs (R21-1) and Marksmanship's movement weakness (R21-2). No hunter note, order, rank row or APL condition moved (the addon-data.json regeneration changed no hunter row).

- **BM:113, MM:119 (R21-1 fix):** "every miss throws away a whole shot, auto shots included, and auto shot is more of your damage than any button you press". True: every shot rolls the same ranged hit table, and sim round 20 B1/M1 put Auto Shot at 38% (BM) and 46.5% (MM) against 13% and 14% for the next button. OK on both pages; MM keeps its Surefooted sentence and BM ends "past the cap it adds nothing" once.
- **MM:60 (R21-2 fix):** "Auto shot is the biggest share of your damage and stops whenever you move, so a fight full of movement costs you more than it costs a Beast Mastery hunter, whose pet keeps fighting". True: plurality share (M1), auto shot cannot fire on the move, and BM loses a smaller share with the pet still attacking. OK.

## Round 22 shapes, checked on every page and note

1. **A spell named in a "Levels 20 to 60" paragraph without its trainer level.** Every spell in the three 20 to 60 blocks checked against `trainables/hunter.json`: BM (Hunter's Mark 6, Serpent Sting 4, Arcane Shot 6, Wing Clip 12, Feign Death 30, Aspect of the Cheetah 20, Aimed Shot 20 and rank 2 at 28, Multi-Shot 18, Bestial Wrath 40), MM (Rapid Killing 25 and 26, Rapid Fire 26, Rapid Recuperation 33 and 34, Aimed Shot ranks 28 and 52 stated, the rest as BM), SV (Lone Wolf 60, Raptor Strike, Mongoose Bite 16, Wing Clip, Aspect of the Monkey to 30, Aspect of the Beast 30, Feign Death 30, Cheetah 20, Raptor Strike ranks 40 and 48, Immolation Trap rank at 46). **One fails: Lacerate in SV:142 (HUNTER-R22-1).** Immolation Trap (16) and MM's Aimed Shot rank 4 (44) carry no level but are trained before the lines that name them apply; Shoulds.
2. **A note comparing to a rank or spell on bands before it is trained.** No hunter note names a rank. Four notes compare with Multi-Shot on bands 10 and 12, before Multi-Shot trains at 18 (**HUNTER-R22-2**). The round 22 brief makes this a check; round 21 had it as a Should.
3. **A list where "far less than your best" includes the best.** No such list on any hunter page or note.
4. **A rounded cost quoted as exact.** No page or note quotes a numeric mana cost. Numbers quoted (Heart of the Lion 280 at 60, Trueshot Aura 100 at 60, Sniper Shot 590 at rank 3, 4 second cast, 15 second cooldown, Bestial Wrath 50% 18 s 2 min, Rapid Fire five minutes, Mortal Shots 212/224/230%, Barrage 5%/15%, Improved Stings 13%) all match the client exactly.

## Superlative and comparison sweep

Every line from round 21's table re-read on the current text; all hold as round 21 found, with one new angle (Volley) recorded as a Should below.

| line | claim | verdict |
|---|---|---|
| index:33, BM:91, MM:96, MM Aimed note | Aimed Shot outhits Multi-Shot | OK: W+40 to W+350 against W |
| BM:75 | Focused Fire the only flat damage bonus to your shots in the BM tree | OK |
| BM:93, MM:98, BM:113, MM:119 | auto shot more damage than any button | OK (B1, M1) |
| BM:99, MM:105, BM and MM Multi-Shot notes | Multi-Shot the worst damage for its mana you fire | OK against every single-target shot at 18 to 60 (54 to 239 mana, `basemp.txt` x 13.9%); against Volley see Should S3 |
| BM:99, MM:105 | sting best per mana, Arcane next | OK in the raid section (round 21 D3); Immolation Trap at 60 (690 + 2% RAP a tick, 245 mana, no Efficiency) sits below Arcane Shot once Arcane Shot's crit and Mortal Shots count |
| BM:153 | BM the best way to level | OK (Warcraft Tavern, round 21 D4) |
| BM:157 | rank 1 well over half of the newest rank, far less mana | OK; "well over half" carried Should |
| MM:163 | rank 1 most of the newer ranks, a fraction of the mana | OK: 75 against 115 to 260 (Efficiency on both) |
| MM:82, MM:88, MM Sniper note | Sniper Shot hardest/biggest hit | OK at 40 to 60 against Aimed ranks 3 to 6 with Barrage (Sniper leads while W is under 715 at 40 to 47 and under 1250 at 60) |
| SV:72 | Expose Prey the only way Mongoose Bite opens on a boss not attacking you | OK |
| SV:86 | 30 to 59 Beast the only melee aspect; Falcon equals Beast at 60, 10 more with Hawk rank 7 | OK: Falcon = highest Hawk (110, 120 book), Beast rank 4 110 |
| SV:146 | rank 5 Raptor Strike nearly rank 6 for less mana | OK: +35 for 55 against +40 for 70 |
| BM:159, MM:165, SV:144 | mail at 40 the biggest armor upgrade | OK |
| SV Strider Kick note | a sliver of mana | carried Should |

## Findings

| id | status | kind | file | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| HUNTER-R22-1 | CONFIRMED | leveling (trainer level) | `survival.md:142` | "Press Mongoose Bite whenever it lights up and keep Lacerate on anything that will live long enough to bleed." in the "Levels 20 to 60" block | Lacerate is trained at 30 (`trainables/hunter.json` Lacerate ranks 24118 at 30, 24119 at 40, 24120 at 50, 1299332 at 60); the SV card shows no Lacerate row before band 30. The paragraph applies from 20 and names no level for Lacerate, though it gives one for Feign Death, Aspect of the Beast and Cheetah in the same paragraph, so a level 20 to 29 Survival hunter is told to keep up a bleed he cannot train. | "Open with Hunter's Mark, send the pet, then walk in and queue Raptor Strike. Press Mongoose Bite whenever it lights up, and from 30 keep Lacerate on anything that will live long enough to bleed." (rest of the paragraph unchanged) |
| HUNTER-R22-2 | CONFIRMED | card note (spell before trainer) | `hunter-beast-mastery.json` Serpent Sting and Arcane Shot notes; `hunter-marksmanship.json` Serpent Sting and Arcane Shot notes | BM Serpent Sting: "so it does more for its mana than Arcane Shot or Multi-Shot ... When both are ready, Multi-Shot goes first"; BM and MM Arcane Shot: "so it does far more for its mana than Multi-Shot"; MM Serpent Sting: "so it does more for its mana than Arcane Shot or Multi-Shot" | All four rows render at bands 10 and 12 (`addon-data.json` rotations: BM and MM band 10 Hunter's Mark, Serpent Sting 13549, Arcane Shot 3044; band 12 Arcane Shot 14281). Multi-Shot (2643) trains at 18. Each note gates its first sentence "from 18" but then compares with Multi-Shot, and BM's tells you Multi-Shot goes first, on two bands where the player has no Multi-Shot. The round 22 brief rule: a note comparing to a spell must not name that spell on bands before it is trained. | BM Serpent Sting: "Serpent Sting when it has under 1.5 seconds left, and from 18 never while Multi-Shot is about to come back. Refreshing any earlier throws away a tick. Every tick carries 5% of your ranged attack power on top of its own damage, so it does more for its mana than Arcane Shot, and from 18 more than Multi-Shot too: keep it rolling for the whole fight. From 18, when both are ready, Multi-Shot goes first: a late Multi-Shot pushes back the shared cooldown, while a late sting costs only a moment of its damage." MM Serpent Sting: same third sentence ("so it does more for its mana than Arcane Shot, and from 18 more than Multi-Shot too: keep it rolling for the whole fight."). BM and MM Arcane Shot: "... and every hit carries 15% of your ranged attack power, so from 18 it does far more for its mana than Multi-Shot: keep it going even when the mana runs low." |

Counts: CONFIRMED 2, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0.

## Shoulds (wording, not findings)

- **S1. SV:130, Immolation Trap's level.** The card casts Immolation Trap from band 16 and SV:142 tells you to throw it on every pull, but no line gives its level. Not false (trained at 16, before the 20 to 60 block), but the 1 to 19 list already names what arrives at 16: "Wing Clip at 12, and Mongoose Bite and Immolation Trap at 16" closes it.
- **S2. MM:163, Aimed Shot rank 4.** "At 50 and 51, rank 4 and rank 1 come out even" names rank 4 without its level; rank 4 trains at 44 (20902), and "keep rank 1 on your bar through 49" already covers 44 to 49, so nothing is false. "Rank 4 arrives at 44; keep rank 1 over it through 49" would stop a reader assuming it comes at 50.
- **S3. "the worst damage for its mana you fire" against Volley.** On one target Volley is worse per mana than Multi-Shot from about 40: at 60, 672 Arcane damage for 490 mana (no attack power share, `sim/hunter/volley.go` BonusCoefficient 0) against Multi-Shot's normalized weapon damage, roughly 550 in raid gear, for 239 mana (Efficiency and Barrage scale both). Not filed: the claim sits in the boss priority and on single-target card rows, where nobody fires Volley, the same context reading round 21 D3 used for "best for its mana". "the worst damage for its mana of the shots you fire at a boss" would make it literal.
- **Carried from round 21:** BM:157 "well over half" (about 52 to 63% at 52 to 60); SV Strider Kick "a sliver of mana" (5.81% of base mana, the same as the Raptor Strike rank of that band before Resourcefulness); MM Serpent Sting and Arcane Shot "the shared cooldown" at 18 and 19; SV Mongoose Bite chance wording; SV Immolation Trap "ranged attack power"; the 3 November Blizzard notes as a page source; and the older carried Shoulds in rounds 20 and 21.

## Checked and right this round

- **Decoded builds and orders:** BM 31/20/0 leveling (Efficiency 4, Rapid Killing 1 at 57, Mortal Shots 2) and raid (Efficiency 3, Mortal Shots 4); MM 15/31/5 leveling (Rapid Killing 25 to 26, Rapid Recuperation 33 to 34, Barrage 1 at 35) and 5/33/13 raid (Improved Stings 15 to 16, Lone Wolf 27, Barrage 34 to 36, Sniper Shot 40, Surefooted 58 to 60); SV 5/11/35 one order (Resourcefulness 33 to 34, Expose Prey 35 to 36, Survivalist's Discipline 37 to 38, Lacerating Strikes 40, Lone Wolf 60). Every point-count and talent sentence on the three spec pages matches (BM:70 to 83, MM:69 to 88, SV:69 to 82).
- **Levels 1 to 19, all pages:** Heart of the Lion 1, Serpent Sting and Monkey 4, Arcane Shot and Hunter's Mark 6, Concussive Shot 8, Hawk 10, Wing Clip and Mend Pet 12, Mongoose Bite 16, Multi-Shot 18, Raptor Strike ranks 8 and 16, Track Beasts 1, Humanoids 10, Undead 18; Deadly Aspects 10 to 14, Focused Fire 15 to 16, Endurance Training 17 (BM); Lethal Attacks then Efficiency from 15 at 3% a point (MM); Improved Tracking, Savage Strikes 15 to 16 at 2%, Survivalist from 17 at 2% (SV).
- **Levels 20 to 60 rank lines:** BM Aimed Shot rank 1 kept to 60 against rank 2 at 28; MM rank 1 through 49, rank 4 at band 50 (card 20902), rank 1 again at 52 (rank 5 trains at 52), top rank in a raid at 60 (card 20904); SV Raptor Strike rank 6 at 40, rank 5 at 46 and 47 (card 14263), rank 7 at 48, Immolation Trap rank 4 at 46 costing 190 against 135.
- **Races, consumables, professions** against `races.json` and the unchanged BiS files (generated at the 70790 activation): Elune's Light, Ruthless, Artisan's Precision 1% hit, Big Game Hunter, Command 5%, Blood Fury 10% 15 s 2 min, Berserking 10% 15 s 3 min, Wind Blessed 2%; Frozen and Molten Heart of the Mountain (Enchanting, in the 60 raid rows), Legionite Glaive (Blacksmithing main hand, BM and MM), Hyper Deluxe Sniper Rifle (Engineering), Heartseeker (SV raid off hand).
- **Note sweep:** every row below.

## Decisions

- **D1.** R22-1 is CONFIRMED: the paragraph covers 20 to 60, gives levels for its other late spells, and Lacerate trains at 30.
- **D2.** R22-2 is CONFIRMED under the round 22 rule (a note comparing to a spell names it on bands before it is trained); round 21 recorded it as a Should before that rule existed. The fixes gate the Multi-Shot clauses "from 18", the wording the notes already use in their first sentences.
- **D3.** S3 (Volley) stays a Should, consistent with round 21 D3: the superlative is read inside the single-target boss priority it sits in.
- **D4.** Not re-filed: the Survival Hunter's Mark opener (sim lane), and every Sims routed, Sim findings and Engine questions item in the ledger's Round 19, 20 and 21 sections; sim round 21 has not landed in this tree, so none is resolved.

## Sims requested

None new. Carried: Marksmanship raid 60 Auto Shot share on the shipped list after B3 (round 21); damage per mana by spell for rank 1 Aimed Shot, Serpent Sting and Arcane Shot (MM 40, 50, 52, 58, bare 60; BM 40, 50, 58); BM and MM at 18 to 20 below a quarter mana with the Multi-Shot readiness check skipped while Multi-Shot is held.

## Sim findings

None new. Round 20's 18 and 19 low-mana stall and the carried items stay with the sim lane.

## Engine questions

None new. Round 20's two (Immolation Trap on melee or ranged attack power; sting and Arcane Shot share read at cast against per tick) stand.

## Sources

**Client data, 1.60.1.70790:** `trainables/hunter.json` (Lacerate 24118 at 30; Multi-Shot 2643 at 18; Immolation Trap 13795 at 16; Aimed Shot 19434 to 20904 at 20/28/36/44/52/60; Raptor Strike 2973 to 14266), `spelltips/hunter.json` and `spellconst/hunter.json` (1510 to 14295 Volley, 469145 Falcon, 1299447 Beast, 25296 and 14322 Hawk, 409580 Heart of the Lion, 20906, 1310687 to 1310786, 3045, 2643, 3044 to 14287, 1978 to 25295, 13795 to 14305, 1317257, 781, 14317), `talents/hunter.json`, `addon-builds.json`, `addon-data.json` rotations, `gametables/basemp.txt`, `races.json`, `bis/hunter-*.md`.

**Engine (gr22/engine):** `sim/hunter/volley.go` (flat 350/420/490, BonusCoefficient 0), `multi_shot.go` (13.9% base mana), `immolation_trap.go` and `attack_power_scaling.go` (0.02 a tick).

**Repository:** `design/reviews/guide-accuracy/round21/hunter.md`, `ledger.md` (Rounds 19 to 21), `design/reviews/2026-11-04-sim-audit-round20/hunter.md` (B1, B3, M1).

**Web:** none needed this round; the client data settles every call. Warcraft Tavern, "Forever Hunter Leveling Guide", https://www.warcrafttavern.com/forever/guides/hunter-leveling/ still backs BM:153 (round 21 D4).

## Note sweep

One line per row note: spec, row, bands it renders, each named talent, spell, rank or value with the level the player has it (leveling and raid orders), verdict.

- BM Hunter's Mark (10 to 60): spell 6; one mark per target. OK.
- BM Bestial Wrath (40 to 60): talent 40 on both orders; 50%, 18 s, fear, two minutes (19574). OK.
- BM Rapid Fire (26 to 60): spell 26; Multi-Shot 18; Bestial Wrath "from 40" (both orders); five minutes (3045), four with the level 60 build's Rapid Killing at 57 (leveling order only; raid order none). OK.
- BM Aimed Shot (20 to 60): spell 20; rank 1 at 20 to 58, rank 6 at 60; small extra over Multi-Shot; quarter-mana gate matches the APL. OK.
- BM Multi-Shot (18 to 60): spell 18; "(on cooldown before Aimed Shot arrives at 20)"; quarter-mana gate; worst per mana among single-target shots at every band (Volley: S3). OK.
- BM Serpent Sting (10 to 60): ranks 2 to 8 by trainer; "from 18" gate; 5% RAP a tick; beats Arcane Shot per mana at every band; Multi-Shot comparison and "Multi-Shot goes first" at 10 and 12, before 18: **HUNTER-R22-2**.
- BM Arcane Shot (10 to 60): ranks 1 to 8; "from 18" gate; 15% RAP; far more per mana than Multi-Shot at 18 to 60; that comparison at 10 and 12: **HUNTER-R22-2**.
- MM Hunter's Mark (10 to 60): as BM. OK.
- MM Rapid Fire (26 to 60): spell 26; Multi-Shot 18; Rapid Killing 25 and 26 on the leveling order (three minutes from 26); raid order none, five minutes, stated as the raid build. OK.
- MM Aimed Shot (20 to 60): spell 20; rank 1, rank 4 at band 50, rank 6 at 60; outhits Multi-Shot at each (Barrage 1 at 35 leveling, 3 at 34 to 36 raid, on both shots). OK.
- MM Multi-Shot (18 to 60): as BM. OK.
- MM Sniper Shot (40 to 60): raid talent at 40 (leveling-order rows: component lane); ranks 40/48/58; 4 s cast, 15 s cooldown; hardest hit at every band. OK.
- MM Serpent Sting (10 to 60): "from 18" gate; "the shared cooldown" (carried Should); 5% RAP; Improved Stings 15 to 16 raid order only, true without it; Multi-Shot comparison at 10 and 12: **HUNTER-R22-2**.
- MM Arcane Shot (10 to 60): "from 18" gate; 15% RAP; far more per mana than Multi-Shot at 18 to 60; that comparison at 10 and 12: **HUNTER-R22-2**.
- SV Raptor Strike (10 to 60): ranks 2 to 8 (rank 5 row at 46); no global; Resourcefulness "at 33 and 34", 60% at rank 2. OK.
- SV Counterattack (30 to 60): talent 32 (rows at 30 and 31: component lane); cannot be dodged, blocked or parried. OK.
- SV Mongoose Bite (16 to 60): spell 16; 5 s dodge window; Expose Prey "at 35 and 36", 10 s, any hunter's mark; one press per window. OK (chance wording carried Should).
- SV Lacerate (30 to 60): spell 30; names Mongoose Bite (16). OK.
- SV Strider Kick (30 to 60): talent 30; 8 s, 100% weapon damage; "sliver of mana" carried Should. OK.
- SV Immolation Trap (16 to 60): ranks 16/26/36/46/56; Clever Traps 20 to 21 and Survivalist's Discipline 37 to 38 as a "once in" clause; attack power share. OK.
- SV Wing Clip (12 to 60): ranks 12/38/60; Windfury Totem conditional; main hand hit with procs; little damage of its own (5/25/50). OK.
