# Guide accuracy round 23: Hunter

Audited on the round 23 audit tree (main after guide round 22, head 24ee457d0; no sim round 21 commit in the log, so no sim round 21 report exists to read; engine source in gr23/engine at pin d5c708834), against the active client 1.60.1.70790 (`data/builds/1.60.1.70790/`). Files: `web/src/content/guides/hunter/{index,beast-mastery,marksmanship,survival}.md`, the 21 row notes in `data/curated/apl/hunter-{beast-mastery,marksmanship,survival}.json`, every rendered row of `addon-data.json` `rotations/hunter-*` (BM 18 bands, MM 19, SV 20), the six orders in `addon-builds.json` (decoded level by level), `talents/hunter.json`, `trainables/hunter.json`, `spellconst/hunter.json`, `spelltips/hunter.json`, `gametables/basemp.txt`, `races.json`, `enchants.json`, `items.json`, `bis/hunter-*.md`. Read first: round 22 hunter report, the ledger's Round 19 to 22 sections, the briefs back to round 1.

**Not re-checked this round, carried from round 22 and earlier:** the frontmatter source links (the lists did not change, no fetches made), quiver speed percentages (`items.json` carries no effect text; item levels were checked, the 10 to 15% figures were not re-derived), and pet training claims.

## Changes since round 22's tree

The only hunter commit after round 21 is 58eb16c2e (round 22's fixes); nothing has touched the hunter pages, notes or orders since it.

- **R22-1 fix, SV:142:** "and from 30 keep Lacerate on anything that will live long enough to bleed". Lacerate trains at 30 (24118); the first card row with it is band 30. OK. The paragraph still reads in order with "Wear Aspect of the Monkey until 30".
- **R22-2 fix, four notes:** BM and MM Serpent Sting "more for its mana than Arcane Shot, and from 18 more than Multi-Shot too"; BM and MM Arcane Shot "so from 18 it does far more for its mana than Multi-Shot"; BM Serpent Sting "From 18, when both are ready, Multi-Shot goes first". No clause now names Multi-Shot unguarded at bands 10 and 12. Arcane Shot (6) is trained on both early bands, so the ungated half holds there. OK.
- **Both sides across bands (this round's side-effect check):** Serpent Sting against Arcane Shot at every rank pair the rows render:

| band | sting rank, damage/mana | Arcane rank, damage/mana |
|---|---|---|
| 10 | r2 35/30 | r1 25/25 |
| 12 | r2 35/30 | r2 35/35 |
| 18 | r3 70/50 | r2 35/35 |
| 20 | r3 70/50 | r3 50/50 |
| 26 | r4 130/80 | r3 50/50 |
| 28 | r4 130/80 | r4 85/80 |
| 34 | r5 200/115 | r4 85/80 |
| 36 | r5 200/115 | r5 120/105 |
| 42 | r6 290/150 | r5 120/105 |
| 44 | r6 290/150 | r6 175/135 |
| 50 | r7 375/190 | r6 175/135 |
| 52 | r7 375/190 | r7 220/160 |
| 58 | r8 500/230 | r7 220/160 |
| 60 | r8 500/230 | r8 280/190 |

The sting leads on base per mana at every pair and carries the larger attack power share (25% over the duration against 15%). No talent in either order moves one side only in a way that reverses this: Efficiency covers both, and Improved Stings only helps the sting. OK.

## Findings

| id | status | kind | file | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| HUNTER-R23-1 | LIKELY | comparison (trash mana) | `beast-mastery.md:105`, `marksmanship.md:111` | BM: "Volley is a channel, costs a lot of mana, and is worth it only on a big pack that will live through the whole channel. ... Explosive Trap under a pack the tank is holding is good damage for its mana." MM: "Volley is a mana-heavy channel, worth it only on a big pack that lives through it. ... Explosive Trap under a pack the tank is holding is good damage for its mana." | The pair reads as Volley the expensive button and the trap the efficient one; the client numbers run the other way. Explosive Trap rank 3 (14317) costs 520 for 229 + 330 = 559 a target, 1.07 per mana per target, and Efficiency does not cover it ("Shots, Stings, and melee abilities"; neither build has Resourcefulness). Volley rank 3 (14295) costs 490 for 112 x 6 = 672 a target, 1.37, and Marksmanship's Barrage adds 15% to it. Lower ranks: trap 0.96 and 1.02 (275 mana for 265, 395 for 403); Volley 1.20 and 1.30 (350 for 420, 420 for 546). The trap costs more than Volley at 44 to 49 and from 54; at no level with both trained is its damage per mana higher. Multi-Shot, the page's main pack button, costs 239 at 60 for three targets. Status is LIKELY, not CONFIRMED: both spells also scale with attack power on the server by coefficients no source gives (sim round 20 S4; a web search this round found none for Forever), and the trap does beat single-target shots per mana on a pack of about five or more, so "good" is not false on its own. The costs and the contrast are hard data. | BM: "Explosive Trap under a pack the tank is holding hits everything within 10 yards, and at 60 it costs even more than Volley, so save it for the same big packs." MM: same sentence. Survival's line (SV:96) makes no mana claim and has Resourcefulness; leave it. |

Counts: CONFIRMED 0, LIKELY 1, NEEDS-SIM 0, NEEDS-OWNER 0.

## Shoulds (wording, not findings)

- **S1. SV:142, "Throw Immolation Trap under anything".** New this round. Traps go at your own feet (round 5 D4; the Immolation Trap note says "at your feet"). The verb reads as idiom and a melee hunter stands on the target, so nothing misleads in play, but "Drop Immolation Trap at your feet on anything that will live long enough to burn" matches the note.
- **S2. SV Wing Clip note, "a cheap melee special".** New this round. At 12 to 37 rank 1 costs 40, above Raptor Strike rank 2 (25) at 12, rank 3 (35) at 16 to 23 and Mongoose Bite rank 1 (30), but the line applies only under a Windfury Totem, which a group brings well after that. With Resourcefulness at 34 it is 16 mana. Not filed; "a melee special that can proc Windfury" drops the word with no loss.
- **S3. SV:86, "Deadly Aspects procs from your swings under either one".** Sim round 15 W2 reads the client rows as Falcon's own flat 5% that the talent cannot raise; Falcon's tooltip says the talent "can now proc from and benefit melee attacks". Round 10 D1 kept the line and the engine item is already carried, so it is not re-filed. "Your swings can proc the same 30% attack speed under either one" is true under both readings and keeps "not a clear upgrade".
- **S4. SV:80, "Strider Kick is a third melee special".** True against Raptor Strike and Mongoose Bite; Lacerate, Counterattack and Wing Clip are specials too. "another melee special" is exact.
- **S5. BM Aimed Shot note and BM:91, "a small extra" at band 60.** The rank 1 rows add 40; the band 60 row is rank 6, W+350, not small per cast. The claim holds as a share of the fight (sim round 20 B1, Aimed Shot 5%) once the 2 second cast's cost to auto shots counts. Leave unless the note is touched; "a small share of your damage" would be exact.
- **S6. Instants and auto shot.** The bullet "Fire your instants right after an auto shot leaves, so they never push the next one back" (BM:96, MM:101) sits beside the notes' "It is instant, so it never delays your next shot". Same advice, different implication; no client data settles the wind-up behaviour.
- **S7. index:28, "Survival fights in melee with no pet".** The page's Marksmanship clause says "raids"; Survival levels with a pet (SV:140). "and Survival fights in melee, with no pet in a raid" lines them up.
- **Carried from round 22:** Immolation Trap's level absent from the 1 to 19 list (S1 there); MM:163 rank 4's level (S2 there); "the worst damage for its mana you fire" against Volley on one target (S3 there); BM:157 "well over half"; SV Strider Kick "a sliver of mana"; MM Serpent Sting and Arcane Shot "the shared cooldown" at 18 and 19; SV Mongoose Bite chance wording; SV Immolation Trap "ranged attack power"; the 3 November Blizzard notes as a page source; the older carried Shoulds in rounds 20 and 21.

## Superlative, comparison and quoted value sweep

| line | claim | verdict |
|---|---|---|
| index:33, BM:91, MM:96, MM Aimed note | Aimed Shot outhits Multi-Shot | OK: W+40 to W+350 against W; Barrage is on both |
| BM:75 | Focused Fire the only flat damage bonus to your shots in the BM tree | OK |
| BM:93, MM:98, BM:113, MM:119, MM:60 | auto shot more damage than any button, biggest share | OK (sim round 20 B1, M1) |
| BM:99, MM:105, Multi-Shot notes | Multi-Shot the worst damage for its mana you fire | OK against every single-target shot (54 mana at 18, 239 at 60); Volley carried Should |
| BM:99, MM:105 | sting best per mana, Arcane next | OK in the raid section (round 21 D3; B3 BM 6.4/4.5/3.8/3.1, MM 8.8/6.5/3.5/3.4) |
| Arcane notes | far more for its mana than Multi-Shot from 18 | OK: about 1.5x at raid 60 (BM), more on MM |
| BM:105, MM:111 | Volley heavy, Explosive Trap good for its mana | **HUNTER-R23-1** |
| BM:101 | Mend Pet cheaper than Revive Pet | OK: rank 7 is 480; Revive Pet is 100% of base mana (982), 1720 at 60 |
| BM:153 | BM the best way to level | OK (round 21 D4) |
| BM:157, MM:163 | rank 1 Aimed Shot against newer ranks | OK; 75 mana against 115 to 310, Efficiency on both |
| MM:82, MM:88, MM Sniper note | Sniper Shot hardest/biggest hit | OK: W+320 at 40, W+450 at 48, W+590 at 58, ahead of Barrage-boosted Aimed Shot while W is under about 713 at 44 and 1250 at 60 |
| SV:72 | Expose Prey the only way Mongoose Bite opens on a boss not attacking you | OK |
| SV:86 | 30 to 59 Beast the only melee aspect; Falcon equals Beast at 60, 10 more with Hawk rank 7 | OK: Hawk r6 110, r7 120, Beast r4 110 |
| SV:146 | rank 5 Raptor Strike nearly rank 6 for less mana; trap rank costs more at 46 | OK: +35/55 against +40/70; 190 against 135 |
| BM:159, MM:165, SV:144 | mail at 40 the biggest armor upgrade | OK |
| "every" claims: BM:72, BM:73, BM:74, MM:73, MM:75, MM:76, MM:79, SV:51, SV:77, SV:80, SV:102 | a talent covers all of it | OK against the talent text: Unleashed Fury, Ferocity plus Frenzy 100%, Lethal Attacks "all attacks", Efficiency "Shots, Stings", Lone Wolf "all attacks", Deadly Aspects "all melee auto attacks", Resourcefulness 100% at 2 points. MM:73 "everything your bow or gun does": the engine excludes only Serpent Sting and the hawk, neither a weapon hit (talent on the sting untested, sim round 20 H4) |
| quoted values | Heart of the Lion 280 at 60; Trueshot Aura 100; Sniper Shot 590, 4 s, 15 s; Bestial Wrath 50%, 18 s, 2 min; Rapid Fire 5 min, 4 and 3 with Rapid Killing; Mortal Shots 212/224/230%; Predator's Edge 230%; Barrage 5/15%; Improved Stings 13%; Lacerating Strikes 40% over 21 s; Strider Kick 8 s, 100%; Expose Prey 10%, 10 s; Resourcefulness 60% | OK against client rows. No page or note quotes a numeric mana cost, so no rounded cost is stated as exact |

## Checked and right this round

- **Decoded builds and orders:** BM 31/20/0 leveling and raid (Efficiency 4 to 3, Rapid Killing 1 at 57 leveling only, Mortal Shots 2 to 4); MM 15/31/5 leveling and 5/33/13 raid (Sniper Shot at 40 on the raid order, 30 tree points reached at 39); SV 5/11/35 one order (Clever Traps 20 to 21, Surefooted 22 to 24, Deterrence 25, Strider Kick 30, Counterattack 32, Resourcefulness 33 to 34, Expose Prey 35 to 36, Survivalist's Discipline 37 to 38, Lacerating Strikes 40, Lone Wolf 60). Every point count on the spec pages matches. Prerequisites read: Bestial Swiftness to Intimidation to Bestial Wrath; Ferocity to Frenzy; Careful Aim to Mortal Shots; Rapid Killing to Rapid Recuperation; Trueshot Aura to Sniper Shot; Deterrence to Counterattack; Expose Prey to Lacerating Strikes.
- **Levels named in leveling paragraphs:** every spell against `trainables/hunter.json` (Heart of the Lion and Auto Shot 1; Monkey and Serpent Sting 4; Arcane Shot and Hunter's Mark 6; Concussive Shot 8; Hawk 10; Wing Clip and Mend Pet 12; Mongoose Bite 16; Multi-Shot 18; Aimed Shot and Cheetah 20; Rapid Fire 26; Beast, Feign Death and Lacerate 30; Bestial Wrath 40; Track Beasts 1, Humanoids 10, Undead 18). Dual Wield at 20 rests on round 5 D3.
- **Rank lines against the rows:** BM rank 1 Aimed Shot to 58, rank 6 at 60; MM rank 1 to 48, rank 4 at 50, rank 1 at 52 and 58, rank 6 at 60; SV Raptor Strike rank 6 at 40 to 44, rank 5 at 46, rank 7 at 48.
- **Prose against the APL:** quarter-mana gate on Multi-Shot, sting under 1.5 s and not while Multi-Shot is within 0.5 s, Arcane Shot only when Multi-Shot is over 1.5 s away, Aimed and Sniper only with over 1.2 s to the next auto, Rapid Fire as an auto leaves with Multi-Shot under 0.5 s away, Survival's listed order, Wing Clip only under Windfury.
- **Races:** Elune's Light 10% crit 15 s; Ruthless 1% crit; Artisan's Precision 1% hit; Big Game Hunter 5% on Beasts; Command 5% pet damage; Blood Fury 10% 15 s 2 min; Berserking 10% 15 s 3 min (`races.json`). Wind Blessed 2% carried from round 22.
- **Gear and consumables:** Frozen and Molten Heart of the Mountain, Legionite Glaive, Hyper Deluxe Sniper Rifle and Heartseeker appear in the band 60 rows as the pages say; Jagged Arrow 40, Ice Threaded Arrow 51, Thorium Headed Arrow and Thorium Shells 52, Quickdraw Quiver 40, Harpy Hide Quiver and Gnoll Skin Bandolier 55; Sniper Scope and Biznicks 247x128 Accurascope in `enchants.json`; only a Minor Agility bracer enchant exists, so "no bracer enchant carries real Agility" holds; Major Frenzy Potion for Survival and Ground Scorpok Assay for hunters per `presets.json`.

## Decisions

- **D1.** R23-1 is LIKELY rather than a Should: the two sentences together tell the reader the wrong way round which AoE button is costly, which is misleading at 60. It is not CONFIRMED because the server's attack power shares on both spells are unknown and "good" alone holds on a large enough pack.
- **D2.** Falcon (S3) stays a Should, following round 10 D1, and the engine item stays where it is already carried.
- **D3.** Not re-filed: the Survival Hunter's Mark opener, and every item routed in the ledger's Round 19 to 22 sections. Sim round 21 has not landed, so none is resolved.

## Sims requested

None new. Carried: Marksmanship raid 60 Auto Shot share after B3; damage per mana by spell for rank 1 Aimed Shot, Serpent Sting and Arcane Shot (MM 40, 50, 52, 58, bare 60; BM 40, 50, 58); BM and MM at 18 to 20 below a quarter mana with the Multi-Shot readiness check skipped while Multi-Shot is held.

## Sim findings

None new. Round 20's 18 and 19 low-mana stall and the carried items stay with the sim lane.

## Engine questions

None new. Standing: Immolation Trap on melee or ranged attack power; sting and Arcane Shot share read at cast against per tick; Falcon's flat 5% against the tooltip's talent wording. R23-1 would firm up with a live reading of Explosive Trap's and Volley's attack power shares (sim round 20 S4 left both flat and unmodeled).

## Sources

**Client data, 1.60.1.70790:** `spellconst/hunter.json` and `spelltips/hunter.json` (13813, 14316, 14317 Explosive Trap; 1510, 14294, 14295 Volley; 2643 Multi-Shot; 982 and 13544 pet spells; 469145 Falcon; every rank in the tables above), `talents/hunter.json`, `trainables/hunter.json`, `addon-builds.json`, `addon-data.json` rotations, `gametables/basemp.txt` (hunter 250 at 12, 391 at 18, 910 at 34, 1720 at 60), `races.json`, `enchants.json`, `items.json`, `bis/hunter-*.md`.

**Engine (gr23/engine):** `sim/hunter/talents.go` (Ranged Weapon Specialization exclusions), `aspect_of_the_falcon.go`, `attack_power_scaling.go`.

**Repository:** `design/reviews/guide-accuracy/round22/hunter.md`, `round21/hunter.md`, `round10/hunter.md` (D1), `ledger.md`, `design/reviews/2026-11-04-sim-audit-round20/hunter.md` (B1, B3, M1, H4, S4), `design/reviews/2026-11-04-sim-audit-round15/hunter.md` (W2), `data/curated/presets.json`.

**Web:** one search for Forever attack power coefficients on Explosive Trap and Volley ("WoW Forever hunter wiki Attack Formulas Explosive Trap Volley attack power coefficient") returned only older-expansion pages and nothing for Forever, so none is cited as evidence. Warcraft Tavern, "Forever Hunter Leveling Guide", https://www.warcrafttavern.com/forever/guides/hunter-leveling/ still backs BM:153 per round 21 D4 (not re-fetched).

## Note sweep

One line per row note: spec, row, bands it renders, each named talent, spell, rank or value with the level the player has it (leveling and raid orders), verdict.

- BM Hunter's Mark (10 to 60): spell 6; one mark per target. OK.
- BM Bestial Wrath (40 to 60): talent 40 on both orders; Rapid Fire 26; 50%, 18 s, fear, two minutes (19574). OK.
- BM Rapid Fire (26 to 60): spell 26; Multi-Shot 18; Bestial Wrath "from 40"; five minutes (3045), four with Rapid Killing at 57 (leveling order only, stated as the level 60 build). OK.
- BM Aimed Shot (20 to 60): spell 20; rank 1 at 20 to 58, rank 6 at 60; Multi-Shot 18; quarter-mana line matches the APL; "small extra" at 60 is Should S5. OK.
- BM Multi-Shot (18 to 60): spell 18; "(on cooldown before Aimed Shot arrives at 20)"; quarter-mana gate; instant; worst per mana among single-target shots at every band. OK.
- BM Serpent Sting (10 to 60): ranks 2 to 8 by trainer; every Multi-Shot clause gated "from 18"; 5% RAP a tick; beats Arcane Shot per mana at every band (table above). OK.
- BM Arcane Shot (10 to 60): ranks 1 to 8; "from 18" on both Multi-Shot clauses; 15% RAP; instant. OK.
- MM Hunter's Mark (10 to 60): as BM. OK.
- MM Rapid Fire (26 to 60): spell 26; Multi-Shot 18; Rapid Killing 25 and 26 on the leveling order (three minutes from 26); raid order none, five minutes, stated as the raid build. OK.
- MM Aimed Shot (20 to 60): spell 20; rank 1, rank 4 at band 50, rank 6 at 60; outhits Multi-Shot at each (Barrage 1 at 35 leveling, 3 at 34 to 36 raid, on both shots). OK.
- MM Multi-Shot (18 to 60): as BM. OK.
- MM Sniper Shot (40 to 60): raid talent at 40 (leveling-order rows: component lane); ranks 40/48/58; 4 s cast, 15 s cooldown; hardest hit at every band. OK.
- MM Serpent Sting (10 to 60): "from 18" gates; "the shared cooldown" at 18 and 19 (carried Should); 5% RAP; Improved Stings 15 to 16 raid order only, true without it. OK.
- MM Arcane Shot (10 to 60): "from 18" gates; 15% RAP; far more per mana than Multi-Shot at 18 to 60 with Barrage and Ranged Weapon Specialization on Multi-Shot. OK.
- SV Raptor Strike (10 to 60): ranks 2 to 8 (rank 5 row at 46); no global (gcd 0); Resourcefulness "at 33 and 34", 60% at rank 2. OK.
- SV Counterattack (30 to 60): talent 32 (rows at 30 and 31: component lane); after a parry; cannot be dodged, blocked or parried (client text). OK.
- SV Mongoose Bite (16 to 60): spell 16; 5 s dodge window; Expose Prey "at 35 and 36", 10 s, any hunter's mark; one press per window. OK (chance wording carried Should).
- SV Lacerate (30 to 60): spell 30; names Mongoose Bite (16). OK.
- SV Strider Kick (30 to 60): talent 30; 8 s, 100% weapon damage; "the bleed" Lacerate 30; "sliver of mana" carried Should. OK.
- SV Immolation Trap (16 to 60): ranks 16/26/36/46/56; Clever Traps 20 to 21 and Survivalist's Discipline 37 to 38 as a "once in" clause; attack power share (engine question standing). OK.
- SV Wing Clip (12 to 60): ranks 12/38/60; Windfury Totem conditional; main hand hit with procs; little damage of its own (5/25/50); "cheap" is Should S2. OK.
