# Guide accuracy round 24: Hunter

Audited on the round 24 audit tree (main after guide round 23, head ea9adea1c, which carries sim round 21; engine pin 170a59c39 in `sim/enginever/version.go`, engine source in gr24/engine at 170a59c39), against the active client 1.60.1.70790 (`data/builds/1.60.1.70790/`). Files: `web/src/content/guides/hunter/{index,beast-mastery,marksmanship,survival}.md`, the row notes and prepull in `data/curated/apl/hunter-{beast-mastery,marksmanship,survival}.json`, every rendered row of `addon-data.json` `rotations/hunter-*` (BM 18 bands, MM 19, SV 27 with the new 26 to 29 and 48 to 49 rows and their `race_spell_ids`), the six orders in `addon-builds.json` decoded level by level, `talents/hunter.json`, `trainables/hunter.json`, `spellconst/hunter.json`, `spelltips/hunter.json`, `data/curated/rank-choices.json`. Read first: round 23 hunter report, the ledger's Round 19 to 23 sections, `design/reviews/2026-11-05-sim-audit-round21/hunter.md` and `integration.md`, and `git diff 24ee457d0..HEAD` on the hunter pages, cards and coach files.

## New work: sim round 21 (commits 9045ac705, 2784c90b0, b584ebbd4, 369342be0)

What changed for hunter, and the verdict on each:

| change | where | verdict |
|---|---|---|
| Hunter's Mark first in the prepull (-20 s, ahead of the aspect); prose "mark ... first, so the mana is back by the pull" | all three cards' prepull; BM:87, MM:92, SV:86 | OK. Mark rank 4 costs 60 (14325); the aspect follows at -10 s. Survival adds the remark at two minutes: client duration 120000 ms. OK |
| Aimed Shot ungated ("whenever the cooldown it shares with Multi-Shot is up"; auto shots keep firing through the cast) | BM and MM Aimed notes; BM:91, MM:96 | OK on the engine change (A1, beta-log evidence). MM's new reason "with Barrage" is **HUNTER-R24-5** |
| Multi-Shot row: "on cooldown at 18 and 19 ... From 20 ... for packs and for moving" | BM and MM Multi-Shot notes | OK at every band (Aimed Shot sits above it with no gate from 20). First-sentence rendering at 20 to 60 is Should S2 |
| Sniper Shot ungated ("the moment it comes up") | MM Sniper note, MM:88, MM:96 | OK against the card. That the auto keeps firing through Sniper Shot is inferred from Aimed Shot (engine question E1) |
| Rapid Fire notes rewritten ("the shared cooldown is about to come back", "first Aimed Shot lands inside it") | BM and MM Rapid Fire notes | The level guards and the leveling-build cooldown were dropped: **HUNTER-R24-2** (MM), **HUNTER-R24-3** (BM) |
| Low-mana rule at every level (sting and Arcane Shot fire under a quarter of mana without waiting on the shared cooldown) | BM and MM sting and Arcane notes; the new "Multi-Shot comes at 18" paragraphs BM:145, MM:151; BM:99, MM:105 | OK against the APL conditions (`currentManaPercent <= 25%` OR branches). BM sting note's last sentence is **HUNTER-R24-4** |
| Survival Hunter's Mark row and note ("from 35 Expose Prey needs it") | SV card | OK: Expose Prey 35 and 36 on the one order; the row renders 10 to 60 and the Expose Prey clause is guarded |
| Immolation Trap note first sentence "on cooldown", "at your feet" moved | SV card | OK; cooldown 30 s (14305 category), 18 s with Survivalist's Discipline from 38, not quoted |
| Survival Raptor Strike rank rows at 26 to 29 (race-keyed) and 48 to 49 (Horde) with prose SV:144 and SV:148 | SV card, `rank-choices.json` | **HUNTER-R24-6** (26 to 28); 48 to 49 is Should S4 |

Sim round 21 left one line behind on the class page: **HUNTER-R24-1** (index:33).

## Round 23's fix

- **R23-1, BM:105 and MM:111:** "Explosive Trap under a pack the tank is holding hits everything within 10 yards, and at 60 it costs even more than Volley, so save it for the same big packs." Explosive Trap rank 3 (14317, trained 54) costs 520 and hits "all within 10 yards"; Volley rank 3 (14295, trained 58) costs 490. Neither build has Resourcefulness or Clever Traps; Efficiency does not name traps. Even if Efficiency reached Volley it would only widen the gap. True at 60 for both specs. OK.

## Findings

| id | status | kind | file | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| HUNTER-R24-1 | CONFIRMED | line left behind by a sim change | `index.md:33` | "Aimed Shot outhits Multi-Shot, so it takes that cooldown whenever the cast fits between two auto shots, and Multi-Shot takes it when an auto is about to fire." | Sim round 21 A1/A2 removed this weave. Both cards now cast Aimed Shot whenever the shared cooldown is up (no `autoTimeToNext` gate on 20904), and the spec pages say the opposite of this line: BM:91 "cast Aimed Shot every time the shared six second cooldown comes up. In Forever your auto shots keep firing while you cast it ... Multi-Shot is for packs and for when you have to move"; MM:96 the same. The class page teaches a rotation that the cards and both ranged spec pages now contradict, at every level from 20. | "Aimed Shot outhits Multi-Shot, and in Forever your auto shots keep firing while you cast it, so it takes that cooldown every time it comes up; Multi-Shot is for packs and for when you have to move." |
| HUNTER-R24-2 | CONFIRMED | quoted value a talent changes (leveling order) | `hunter-marksmanship.json` Rapid Fire note (bands 26 to 60) | "It is a five minute cooldown, so a boss fight gets one: put your trinkets and racial on the same press." | The Marksmanship leveling order takes Rapid Killing at 25 and 26 ("Reduces the cooldown on your Rapid Fire ability by 2 min"), so from 26, the level Rapid Fire trains (3045, category cooldown 300000), the leveling build's Rapid Fire is a three minute cooldown at every leveling band. The page itself says so: MM:83 "two minutes off Rapid Fire", MM:161 "Rapid Killing at 25 and 26 takes two minutes off Rapid Fire, which you train at 26, so you have it for most hard pulls". Round 22's note said exactly this ("On the leveling build, Rapid Killing has it back every three minutes, so press it on every hard pull"); sim round 21's rewrite dropped it. Five minutes holds only on the raid order (no Rapid Killing). | "... It is a five minute cooldown in the raid build, so a boss fight gets one: put your trinkets and racial on the same press. On the leveling build Rapid Killing has it back every three minutes, so press it on every hard pull." |
| HUNTER-R24-3 | CONFIRMED | talent named before the order gives it; quoted value a talent changes | `hunter-beast-mastery.json` Rapid Fire note (bands 26 to 60) | First sentence, which the card prints: "Rapid Fire on the pull with Bestial Wrath, as an auto shot leaves and the shared cooldown is about to come back." Later: "It is a five minute cooldown, so a boss fight gets one". | Bestial Wrath is the level 40 talent on both BM orders (decoded: `40:Bestial Wrath 1`); the Rapid Fire row renders from band 26, so at 26 to 38 the printed condition pairs Rapid Fire with a spell the player cannot have. On the BM leveling order Rapid Killing 1 comes at 57 ("Reduces the cooldown on your Rapid Fire ability by 1 min"), so at bands 58 and 60 on that build the cooldown is four minutes; BM:81 says so ("takes a minute off Rapid Fire"). Round 22's note carried both guards ("and from 40 with Bestial Wrath", "four once the level 60 build's point in Rapid Killing is in at 57"); sim round 21 dropped them. | "Rapid Fire on the pull as an auto shot leaves and the shared cooldown is about to come back, and from 40 with Bestial Wrath. ... It is a five minute cooldown, four once the level 60 build's point in Rapid Killing is in at 57, so a normal boss fight gets one: ..." |
| HUNTER-R24-4 | CONFIRMED | spell named before it is trained; reason false under the note's own condition | `hunter-beast-mastery.json` Serpent Sting note (bands 10 to 60), last sentence | "From 18, when both are ready, the shared cooldown goes first: a late Aimed Shot pushes it back, while a late sting costs only a moment of its damage." | At 18 and 19 the shared cooldown is Multi-Shot's alone: Aimed Shot trains at 20 (19434). The clause gives an Aimed Shot reason on two bands where there is no Aimed Shot. At those bands under a quarter of mana the sentence is also false: Multi-Shot is gated `currentManaPercent > 25%`, and the sting row passes on `currentManaPercent <= 25%`, so the sting goes and the shared cooldown waits, as the note's own first sentence ("unless you are under a quarter of your mana") and BM:145 say. Round 22's version named Multi-Shot ("a late Multi-Shot pushes back the shared cooldown"), true at 18. | "From 18, when both are ready and your mana is above a quarter, the shared cooldown goes first: a late Multi-Shot, or from 20 a late Aimed Shot, pushes it back, while a late sting costs only a moment of its damage." |
| HUNTER-R24-5 | CONFIRMED | talent credited before the order gives it; wrong reason | `hunter-marksmanship.json` Aimed Shot note (bands 20 to 60); `marksmanship.md:96` | Note: "and with Barrage it outhits Multi-Shot." MM:96: "and with Barrage Aimed Shot outhits Multi-Shot by a wide margin." | Barrage (19461): "Increases the damage done by your Multi-Shot, Aimed Shot, and Volley abilities by 15%." It multiplies both shots alike (the engine applies `barrageDamageMultiplier` to Multi-Shot too), so it is not why Aimed Shot outhits Multi-Shot; the flat bonus is (Aimed Shot rank 1 to 6 adds 40 to 350, Multi-Shot 2643 adds 0). Barrage arrives at 35 on the leveling order (one point, 5%) and 34 to 36 on the raid order, so at bands 20 to 33 the note credits a talent the player does not have for something that is true without it. Round 22's MM note said "It outhits Multi-Shot whenever the cast fits", with no talent. | Note: "... so there is no auto shot to wait for, and it outhits Multi-Shot." MM:96: "and Aimed Shot outhits Multi-Shot by a wide margin." |
| HUNTER-R24-6 | CONFIRMED | prose against shipped rank rows | `survival.md:144` | "From 26 to 29 keep Raptor Strike rank 3 on your bar instead of rank 4 ... If you still run dry over a chain of pulls, rank 2 is better again at 26 and 27." | The shipped rows (`addon-data.json` SV bands 26 to 28, `race_spell_ids`; `rank-choices.json` round 21 entries) give rank 2 by race, not by a mana condition: Orc at 26, Troll at 26 and 27, Tauren at 26, 27 and 28 (Windshaper Skyborne 26 and 27). A Tauren at 28 is shown rank 2 by the addon while the page limits rank 2 to 26 and 27; and at 26 every Horde race is shown rank 2 while the page makes rank 3 the default. The sim note says the Horde gain is measured on the band gear with no extra condition, and that rank 2 loses for the Alliance races, which the page's "if you still run dry" offers them too. | "From 26 to 29 keep Raptor Strike rank 3 on your bar instead of rank 4: ... On a Horde character rank 2 is better again: an Orc at 26, a Troll at 26 and 27, a Tauren at 26 to 28. Drag the lower rank onto your bar ..." |

Counts: CONFIRMED 6, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0.

## Shoulds (wording, not findings)

- **S1. Coach starters (outside this audit's file list, for the integrator).** `data/curated/coach/hunter-beast-mastery.md:24,27` and `hunter-marksmanship.md:24,29` still say "Aimed Shot only as the opener" and "Multi-Shot every time the cooldown it shares with Aimed Shot comes up", the reverse of the cards after sim round 21 A2. `hunter-marksmanship.md:17` repeats MM:48 (S3). They were regenerated in 369342be0 and read the same afterwards, so the generator is mapping the ungated Aimed Shot row to the prepull line. Not a guide-page finding; flag to the coach generator owner.
- **S2. Multi-Shot row on bands 20 to 60.** The card prints only the first sentence, "Multi-Shot on cooldown at 18 and 19, while you are above a quarter of your mana.", so from 20 the row's condition names levels the player has passed and says nothing about packs. True, and the drawer has the rest (cut-off notes are with the component lane since round 8). "Multi-Shot on packs and on the move, and on cooldown at 18 and 19 while above a quarter of your mana" would print the 20 to 60 use.
- **S3. MM:48, "Aimed Shot and Multi-Shot on their shared cooldown".** Since sim round 21 Multi-Shot does not spend the single-target cooldown from 20 (MM:96). Read as "the two shots on the shared cooldown" it holds; "Aimed Shot on the shared cooldown, Multi-Shot on packs" matches the rotation.
- **S4. SV:148, Horde rows at 48 and 49.** The rows keep rank 5 for an Orc, a Tauren and a Troll at 48 and 49 (Windshaper at 48) with no condition; the page gives rank 7 "unless your mana still runs out". Unlike R24-6 nothing on the page contradicts a row outright ("until 50" fits all three Horde base races), so it stays a Should; the R24-6 fix can carry it ("a Horde character keeps rank 5 through 49").
- **S4b. Web card at SV bands 26 to 28 (component lane).** Every race carries a `race_spell_ids` override at 26, 27 and 28, but the site has no code that reads them (only the addon's `Rotation.lua` does), so the guide page's card prints the base row, rank 4, at 26 to 28 beside SV:144's "keep rank 3 from 26 to 29". Not a prose error; the component should print the reader's race row or the Alliance default.
- **S5. BM:99 and MM:105, "under a quarter of your mana let Aimed Shot have the shared cooldown".** Reads as if Multi-Shot took the single-target cooldown above a quarter, which the new BM:91 and MM:96 rule out. Not false on a pack. "hold Multi-Shot first, even on packs" would fit.
- **S6. BM:61, "Mana runs out on long fights if you keep firing Multi-Shot on every cooldown".** Still true; the page no longer advises it on a boss.
- **Carried from round 23:** S1 (SV:142 "Throw Immolation Trap"), S2 (SV Wing Clip "cheap"), S3 (Falcon wording), S4 (Strider Kick "a third melee special"), S6 (instants and auto shot bullet, now beside "auto shots keep firing while you cast it"), S7 (index:28 Survival pet); round 23's S5 ("small extra") is gone with the rewrite. The older carried Shoulds in rounds 20 to 22 stand.

## Superlative, comparison and quoted value sweep (new and changed lines)

| line | claim | verdict |
|---|---|---|
| BM:91, BM Aimed note | Aimed Shot outhits Multi-Shot | OK: W+40 to W+350 against W; Barrage, where taken, on both |
| MM:96, MM Aimed note | "with Barrage" it outhits Multi-Shot | **HUNTER-R24-5** |
| MM Sniper note, MM:82, MM:88 | hardest hit; 4 s cast; 15 s cooldown; 590 at rank 3 | OK (1310687/785/786) |
| BM and MM Multi-Shot notes, BM:145, MM:151 | worst damage for its mana you fire | OK against every single-target shot at 18 (54 mana for W) and from 20; Volley carried Should |
| BM and MM sting notes | more per mana than Arcane Shot, and from 18 than Multi-Shot | OK (round 23 table unchanged) |
| BM:99, MM:105 | sting best per mana, Arcane next (raid, rank 6 Aimed Shot) | OK (sim round 21 G2 kept the raid paragraphs; round 21 D3) |
| BM and MM Rapid Fire notes | five minute cooldown | **HUNTER-R24-2**, **HUNTER-R24-3** |
| BM:89, MM:94, SV:88 | Rapid Fire five minutes in the raid build | OK: no Rapid Killing on the BM raid order, the MM raid order or the SV order |
| BM:87, MM:92, SV:86 | mark first so the mana is back by the pull | OK: 60 mana at -20 s (sim round 21 H3: level with no mark at all) |
| SV:86 | remark after two minutes | OK: 14325 duration 120000 ms |
| SV:144 | rank 3 nearly rank 4 for less mana | OK: +21/35 against +30/45 (14261, 14262); rank 2 +11/25 |
| SV:144 | Immolation Trap rank 2 at 26 runs mana short | OK: 50 mana rank 1 to 90 rank 2 (13795, 14302) |
| BM:105, MM:111 | Explosive Trap at 60 costs more than Volley | OK (520 against 490) |
| index:33 | 2 second cast, one 6 second cooldown shared with Multi-Shot | values OK (19434 cast 2000, category 6000); the weave sentence is **HUNTER-R24-1** |
| "every"/"whenever" claims: BM and MM Aimed notes "whenever the cooldown ... is up", Multi-Shot note "whenever it is up", SV Mark note | against the card order | OK: Aimed Shot has no condition and sits above Multi-Shot; rank 1 (75, 64 with Efficiency) is always affordable while the other rows still fire |

## Checked and right this round

- **Orders decoded:** BM leveling and raid identical to 50 (Bestial Wrath 40 on both; leveling Efficiency 4 at 48 to 51, Careful Aim 52 to 56, Rapid Killing 57, Trueshot Aura 58, Mortal Shots 59 and 60; raid Efficiency 3, Careful Aim 51 to 55, Trueshot Aura 56, Mortal Shots 57 to 60). MM leveling: Lethal Attacks 10 to 14, Efficiency 15 to 19, Careful Aim 20 to 24, Rapid Killing 25 and 26, Trueshot Aura 27, Mortal Shots 28 to 32, Rapid Recuperation 33 and 34, Barrage 35, Ranged Weapon Specialization 36 to 40, then Deadly Aspects, Endurance Training, Focused Fire, Unleashed Fury, Improved Tracking. MM raid: Improved Stings 15 and 16, Lone Wolf 27, Barrage 34 to 36, Sniper Shot 40, Survivalist 53 to 57, Surefooted 58 to 60. SV one order, unchanged from round 23. Every point count on the pages matches (MM:69 31/15/5 and 33/5/13; SV:69 35/11/5).
- **Rendered rows:** BM and MM Aimed Shot rank 1 at 20 to 58, MM rank 4 at 50, rank 6 at 60, matching BM:159 and MM:165. SV Raptor Strike rank 4 at 24 to 28 (base), rank 3 at 29, race rows as R24-6, rank 4 at 30, rank 5 at 32 to 38 and 46, rank 6 at 40 to 44, rank 7 at 48 to 54 (Horde rank 5 at 48 and 49), rank 8 from 56; SV:144 "put rank 4 back at 30" and SV:148 match the base rows.
- **APL against prose:** Rapid Fire as an auto leaves (`autoTimeToNext < 0.1`) with the shared cooldown under 0.5 s; the sting at under 1.5 s left, not while the shared cooldown is within 0.5 s unless at or under a quarter of mana; Arcane Shot only with the shared cooldown over 1.5 s away or at or under a quarter of mana; Multi-Shot above a quarter; Aimed Shot and Sniper Shot ungated; Survival order Mark, Raptor Strike, Counterattack, Mongoose Bite, Lacerate, Strider Kick, Immolation Trap, Wing Clip (SV:90 lists the same).
- **Leveling paragraphs new this round (BM:145, MM:151):** Multi-Shot 18 and Aimed Shot 20 by trainer; the "from 20" clause is labelled; the quarter-mana hold and the ungated sting and Arcane Shot match the rows at 18 and 19 and from 20. They read in order with BM:159 and MM:165.

## Decisions

- **D1.** R24-5 is CONFIRMED, not a Should: the note renders at 20 to 33 with a talent the player has not taken, and the client text shows Barrage cannot be the reason (it raises both shots by the same 15%).
- **D2.** R24-6 is filed on the rows the addon prints by race; the 48 to 49 half stays Should S4 because no sentence there contradicts a row outright.
- **D3.** The coach starters (S1) are outside this audit's file list and are generated; listed for the integrator, not counted.
- **D4.** Not re-filed: everything routed in the ledger's Round 19 to 23 sections, sim round 21's open items (BM Multi-Shot above Aimed Shot on the raid build, the hawk swing, Heart of the Lion on the pet, Immolation Trap's 2%, X1), and the Leveling strip and cooldown cards (component lane). Sim round 21 resolved the round 20 low-mana stall and the Survival Hunter's Mark opener; the pages now state both correctly.
- **D5.** No web research was needed: every call above is decided by the client rows, the shipped orders and the shipped rank rows.

## Sims requested

None new. Carried: Marksmanship raid 60 Auto Shot share after A1/A2 (round 21, still routed).

## Sim findings

None new.

## Engine questions

- **E1.** Sniper Shot shares Aimed Shot's hardcast path, so the engine fires auto shots through its 4 second cast (`RangedShootsWhileCasting`, sim round 21 A1), but the beta logs that support the rule cover Aimed Shot only (Sniper Shot is a level 40 talent, beyond the beta range). MM:88, MM:96 and the Sniper note state it as fact. A live log of a raiding Marksmanship hunter would settle it.
- Standing: Immolation Trap on melee or ranged attack power; sting and Arcane Shot share read at cast against per tick; Falcon's flat 5% against the tooltip; Explosive Trap and Volley attack power shares.

## Sources

**Client data, 1.60.1.70790:** `spellconst/hunter.json` and `spelltips/hunter.json` (3045 Rapid Fire; 19434 and 20900 to 20904 Aimed Shot; 2643 Multi-Shot; 1310687, 1310785, 1310786 Sniper Shot; 1130, 14323 to 14325 Hunter's Mark; 14260 to 14266 Raptor Strike; 13795, 14302 to 14305 Immolation Trap; 13813, 14316, 14317 Explosive Trap; 1510, 14294, 14295 Volley; 13549 to 13555 Serpent Sting; 3044, 14281 to 14287 Arcane Shot), `talents/hunter.json` (Barrage 19461, Rapid Killing, Bestial Wrath, Efficiency, Resourcefulness, Expose Prey), `trainables/hunter.json`, `spellranks.json`, `addon-builds.json`, `addon-data.json` rotations.

**Engine (gr24/engine, 170a59c39):** `sim/hunter/hunter.go:299` (`RangedShootsWhileCasting`), `sim/hunter/multi_shot.go` (Barrage multiplier on Multi-Shot, no flat bonus), `sim/core/attack.go:452`.

**Repository:** `design/reviews/2026-11-05-sim-audit-round21/hunter.md` (H1 to H4, A1, A2, C1, C2, G1, G2, S1, S2), `integration.md`, `design/reviews/guide-accuracy/round23/hunter.md`, `ledger.md` (Round 19 to 23), `data/curated/rank-choices.json`, `data/curated/coach/hunter-*.md`, `git diff 24ee457d0..HEAD`.

**Web:** none consulted this round (D5).

## Note sweep

One line per row note: spec, row, bands it renders, each named talent, spell, rank or value with the level the player has it (leveling and raid orders), verdict.

- BM Hunter's Mark (10 to 60): spell 6 (rank 1 to 21, rank 2 22, rank 3 40, rank 4 58); one mark per target. OK.
- BM Bestial Wrath (40 to 60): talent 40 on both orders; Rapid Fire 26; 50%, 18 s, fear, two minutes (19574). OK.
- BM Rapid Fire (26 to 60): spell 26; Bestial Wrath in the first sentence at 26 to 38 (talent 40); "shared cooldown" Multi-Shot 18, Aimed Shot 20; five minutes, four at 57 to 60 on the leveling order. **HUNTER-R24-3**.
- BM Aimed Shot (20 to 60): spell 20; rank 1 rows 20 to 58, rank 6 at 60; auto through the cast (engine A1); hits harder than Multi-Shot at every rank. OK.
- BM Multi-Shot (18 to 60): spell 18; quarter-mana gate; "From 20 Aimed Shot" labelled; worst per mana among single-target shots at every band. OK (S2 on rendering).
- BM Serpent Sting (10 to 60): ranks 2 to 8 by trainer; "from 18" gates; quarter-mana clause matches the APL; 5% RAP a tick; Multi-Shot comparison gated from 18; last sentence names Aimed Shot at 18 and 19 and fails under a quarter there. **HUNTER-R24-4**.
- BM Arcane Shot (10 to 60): ranks 1 to 8; "from 18" on both Multi-Shot clauses; quarter-mana clause matches; 15% RAP; instant. OK.
- MM Hunter's Mark (10 to 60): as BM. OK.
- MM Rapid Fire (26 to 60): spell 26; "shared cooldown" and "first Aimed Shot" from 26 hold; five minutes false on the leveling order from 26 (Rapid Killing 25 and 26). **HUNTER-R24-2**.
- MM Aimed Shot (20 to 60): spell 20; rank 1, rank 4 at 50, rank 6 at 60; Barrage named at 20 to 33 (leveling 35, raid 34) and as the reason. **HUNTER-R24-5**.
- MM Multi-Shot (18 to 60): as BM. OK (S2).
- MM Sniper Shot (40 to 60): raid talent 40 (leveling-order rows: component lane); ranks 40/48/58; 4 s cast, 15 s cooldown; auto through the cast (E1); hardest hit. OK.
- MM Serpent Sting (10 to 60): "from 18" gates; quarter-mana clause matches; 5% RAP; Improved Stings 15 and 16 raid order only, not named. OK.
- MM Arcane Shot (10 to 60): as BM. OK.
- SV Hunter's Mark (10 to 60): spell 6; Expose Prey "from 35" (35 and 36); "mark the boss before the pull" matches the -20 s prepull. OK.
- SV Raptor Strike (10 to 60): ranks 2 to 8, race rows 26 to 28 and 48 to 49; no global; Resourcefulness "at 33 and 34", 60% at rank 2. OK (the rank prose is R24-6).
- SV Counterattack (30 to 60): talent 32 (rows at 30 and 31: component lane); after a parry; cannot be dodged, blocked or parried. OK.
- SV Mongoose Bite (16 to 60): spell 16; 5 s dodge window; Expose Prey "at 35 and 36", 10 s, any hunter's mark. OK.
- SV Lacerate (30 to 60): spell 30; names Mongoose Bite (16). OK.
- SV Strider Kick (30 to 60): talent 30; 8 s; "the bleed" Lacerate 30. OK (carried Should on "sliver").
- SV Immolation Trap (16 to 60): ranks 16/26/36/46/56; "on cooldown"; Clever Traps 20 and 21, Survivalist's Discipline 37 and 38 as a "once ... are in" clause. OK.
- SV Wing Clip (12 to 60): ranks 12/38/60; Windfury Totem conditional; main hand hit with procs; little damage of its own. OK (carried Should on "cheap").
