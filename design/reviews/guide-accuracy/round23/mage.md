# Guide accuracy round 23: Mage

Audited on the round 23 audit tree (24ee457d0: main after guide round 22; no sim round 21 commit or `*sim-audit-round21*` report in the tree; engine pin d5c708834 in `sim/enginever/version.go`), against the ACTIVE client 1.60.1.70790 under `data/builds/1.60.1.70790/`.

Files read line by line: `web/src/content/guides/mage/{index,arcane,fire,frost}.md`; every row and note in `data/curated/apl/mage-{arcane,fire,frost}.json`, with row order and conditions; the rendered rank rows per band in `addon-data.json` (`rotations.mage-*`); all six orders in `addon-builds.json` decoded level by level (index 0 = level 10) against `talents/mage.json` (every talent, first and top rank text); both FS1 strings of each spec; `trainables/mage.json` with `spelltips/mage.json` and `spellconst/mage.json` for 38 spells at every rank (cost, cast, cooldown, duration, trainer level, family mask); `spelltips/shared.json` racials; `items.json` (wands, Demonic Rune, Major Mana Potion, Robe, Frozen Heart, goggles, the Ahn'Qiraj tomes); `bis/mage-*.json` (Robe, Frozen Heart, goggles present); engine `sim/mage/arcane_blast.go`.

Read first: round 22's mage report and the ledger's Round 22 section (rounds 19 to 21 for what is routed), plus sim rounds 15 to 20's mage reports for the Eureka! mask, the gem cooldowns and any Fire Blast mana measurement. Nothing routed is re-filed.

## Changes since round 22's tree (987d08a7c)

`git diff 987d08a7c..HEAD` touches mage in one line: the third sentence of Fire's Fire Blast note (MAGE-R22-1). No page, order, race, BiS or trainer file changed for mage. Round 22's new sentence holds (see Checked); its last clause is the subject of MAGE-R23-4.

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| MAGE-R23-1 | CONFIRMED | card note: absolute claim false against spells the player has | `mage-fire.json` priority line 7, the Fireball row; the row is the card's filler on bands 10 to 21 (`addon-data.json`: Fire Blast and Fireball are the only rows at 10 to 18) | "Fireball only before you train Scorch at level 22, when it is the only filler you have." | Every mage trains Frostbolt at 4 and Arcane Missiles at 8 (`trainables/mage.json`), and fire.md:108 itself names Frostbolt at 4. Neither is only an option on paper: bare Missiles rank 2 (5144, level 16) is 34 damage a second for 136 damage per 140 mana (0.97), against Fireball rank 3 (145) at 22.8 a second with its burn, 57 for 65 mana (0.88), and rank 4 (3140, level 18) at 29.7 a second, 89 for 95 (0.94). Frostbolt rank 3 (837) is 46 for 50 mana (0.92). So Fireball is the only filler with a Fire school, which is the real reason a Fire mage casts it. Whether it also wins with Ignite is under Sims requested; the wording is false either way. | "Fireball only before you train Scorch at level 22, when it is the only Fire filler you have." |
| MAGE-R23-2 | CONFIRMED | race: a racial's reach stated wider than the client, on the page whose loop casts the excluded spell | `arcane.md:98`; `index.md:38` (Gnome, "pick for Arcane and Frost") | "Eureka! makes your damage spells hit 10% harder and cost 10% less for 8 seconds" | Client row Eureka! 1259817, effect 1 mask (551688947, 4096, 0, 8) against Arcane Wrath's family mask (0x40000000, 0x8000, 0, 0): no shared bit, so Arcane Wrath gets neither the damage nor the cost cut. Every other mage nuke matches: re-derived this round from `spellconst/mage.json` for Arcane Blast (1239700), Missiles, Frostbolt, Fireball, Fire Blast, Scorch, Pyroblast, Ice Lance, Cone of Cold, Arcane Explosion, Blizzard and Flamestrike. Sim round 15 G4 recorded this as engine-aligned but it never reached the page. arcane.md:74 tells the raid Gnome to cast Arcane Wrath every 6 seconds and arcane.md:98 to press Eureka! with Arcane Power, so a Wrath falls in every window. The periodic exclusions (Pyroblast's burn, Ignite) are covered by the word "hit", so Fire and Frost's lines stand. | arcane.md:98 and index.md:38: "Eureka! makes your damage spells, all but Arcane Wrath, hit 10% harder and cost 10% less for 8 seconds" (rest of each sentence unchanged). |
| MAGE-R23-3 | CONFIRMED | rotation: "whenever it is ready" against the card's own row order | `fire.md:70`; `mage-fire.json` priority line 5, first sentence, from 27 | fire.md:70 "Then Scorch to 5 stacks. Fire Blast whenever it is ready."; note "Fire Blast every time it is off cooldown" | `mage-fire.json` line 3, the stack Scorch (Improved Scorch known, and absent, under 5 stacks or at 1.5 s or less), sits above line 5, Fire Blast. Its engineering note says the gate "still runs ahead of Pyroblast and Fire Blast" and "matters for the opening stack-up". So on the pull, with Fire Blast ready from the first global, the card casts five Scorches (7.5 seconds) before it, and again after a target swap. Read in order, the opener tells the player to weave Fire Blast into the stack-up, which the card does not do. fire.md:72 ("Keep 5 Improved Scorch stacks rolling ... Fire Blast every time it is off cooldown") comes after the stacks and is fine. Whether the weave is better is under Sims requested; until then page and card should agree. | fire.md:70: "Then Scorch to 5 stacks, and from there Fire Blast whenever it is ready." Note, a new third sentence: "From Improved Scorch at 27, only the Scorches that build your 5 stacks go ahead of it." |
| MAGE-R23-4 | NEEDS-SIM | stated reason that may hold only while mana is not the limit | `mage-fire.json` priority line 5, last clause; `fire.md:78` | note "skipping it to save mana costs more damage than the mana buys"; fire.md:78 "Keep Fire Blast on cooldown even when you are short: ... the mana you would save is worth less than the Pyroblasts you would lose." | No committed sim measures this (sim rounds 15 to 20 test Fire Blast ranks, and Fire Blast before Heating Up at 20 and 28, never a mana gate at 60). On the client rows, at the page's raid set (about 645 Fire spell power): swapping one Fire Blast (340 mana, 767) for one Scorch (150 mana, 473) loses 294 damage and one Heating Up stack, a third of a Pyroblast's gain over Scorch (1,890 against 473, for 290 more mana). Fire Blast and its stack therefore return 767 damage for 287 mana, 2.67 a mana, while the Scorch that mana would buy returns 3.15. Master of Elements refunds both and does not close it. So once the fight is truly mana-bound, which is the case fire.md:78 names, the simple sum says the reason is backward; while mana is not the limit (a shorter fight, or the Scorch would have no global to go into) it is right. Only a run with the regen, gems, Evocation and Combustion in it can say where it lands. | None until measured. If the gate wins: "Keep Fire Blast on cooldown until you are short; then ...". If it loses: both lines stand as written. |

Counts: CONFIRMED 3, LIKELY 0, NEEDS-SIM 1, NEEDS-OWNER 0.

R23-1 and R23-3's note sentence need `sim/request/apl/mage-fire.apl.json`, `addon-data.json` and Data.lua regenerated (note text only). R23-1 changes a first sentence, so the addon row and page card change.

## Shoulds (wording, not findings)

New:
- **fire.md:78 "you land two of them in the time of one Fireball"** given as the reason Scorch stays the filler "when you are short". When short, the point is mana, not time: two Scorches cost 300 against Fireball's 395 and do about the same per mana (3.15 against 2.91 direct). True, but the better reason is available: "two of them cost less than one Fireball".
- **arcane.md:76 "hold it for a phase" on a longer fight** beside the note's "Arcane Power on cooldown". The note is the short rule and the prose the exception; not false.
- **arcane.md:64 "Improved Channeling stays at four ranks so mobs hitting you do not push back every Blast"**: the reason explains the four ranks, not why it stops there. True (56% at four).
- **frost.md:117 "chills anything that hits you"**: tooltip says "may". Minor.
- **arcane.md:74 and :54 "the next spell free"**: the tooltip says next damage spell. Minor.

Carried from round 22 (still wording): fire.md:80 "a third of your bar" (about 38%); fire.md:84 "takes twice as long" on the leveling build (2.9 s); the Missile Barrage "free" lines inside Arcane Power (arcane.md:30, :36, Frostbolt filler note); Fire Scorch filler "two Scorches in the time of one Fireball" at 22 and 23 on the leveling order; Fire Pyroblast row 20 to 29 ("loses to Scorch" at 20 and 21, gated to 30 so it never fires there); Fire Scorch stack row 27 to 28; Frost Cone of Cold row 26 to 37; Arcane Missiles-first rows 10 to 19; Frost Ice Lance first sentence 30 to 35; arcane.md:72 "before the fight starts"; arcane.md:102; fire.md:94 Greater Arcane Elixir; fire.md:116 Burning Soul's 23%; frost.md:63 Winter's Chill feeders; frost.md:71 and the Ice Barrier note's "the moment it breaks" against a 30 s cooldown; frost.md:95 Missile Barrage's Missiles; index.md:32 Accelerant; the leveling Frostbolt row's "stacks' spell damage" before Blast at 20; Arcane raid PoM on Frostbolt at 54 to 56.

## Checked and right this round

- **R22-1 fix (Fire Blast note, third sentence).** Damage a mana on the client rows at level 60, mean of the variance. Bare: Fire Blast 490.6 / 340 = 1.44; Scorch 196.4 / 150 = 1.31; Fireball rank 11 505.4 / 395 = 1.28 direct; Pyroblast 646.2 / 440 = 1.47. So "in raid gear" is a needed qualifier: at bare 60 Fire Blast beats Scorch and Fireball; it drops behind Scorch at about 86 spell power. Raid set (645): 2.26 against 3.15, 2.91 (more with the burn) and 2.93 direct (4.30 with the burn). Holds; Incineration's crit is on Scorch too and widens nothing. Frostfire Bolt no longer named, so R22-1's failure is gone. "In raid gear" pins the Scorch and Pyroblast names to 60, so nothing is named before its trainer level.
- **Both sides of every comparison, across all bands**: Frostbolt a quarter less than Fireball at 57 to 59 (260 against rank 10's 350, 26%) and a third at 60 (260 against 395, 34%), the row gated to Heating Up so 50 to 56 never render it; Frostbolt the cheaper slow than Frostfire Bolt at 40 (160/205), 44 (195/205), 50 (225/285), 56 (260/285), 60 (260/370; book rank 290); Frostfire Bolt against Frostbolt at 40, 50 and 60 for damage and cost, no talent reversing it; Missiles over Frostbolt per second at 10 to 23 and over Blast per mana at 20 to 22 (round 22's numbers re-derived from the rank rows; no talent later in either order touches the other side before 45 and 58); Frostfire Bolt against Scorch 205/285/370 against 100 to 150; Arcane Explosion 390 drains a bar in under a minute (260 a second).
- **Stated reasons, each read for direction**: Presence of Mind's "global already caps" (Blast 1.5 s, gcd 1.5 s); longest cast under 10 s (Pyroblast 6 s, Teleport 10 s excluded); a missed Blast still raising the next cost (`arcane_blast.go`: the stack is added in ApplyEffects regardless of the hit result, tooltip "each time you cast"); Arcane Mind and Ice Shards doubling the crit bonus; Arcane Subtlety and Frost Channeling threat; Arcane Meditation plus Mage Armor regeneration (50 + 50, 33 + 50); Precision's schools; Improved Scorch stacks and 30 s through a move; Burning Soul 23 and 47%; school spell power lines against the spells each spec casts; Frost "rarely needs Evocation" and "does not run out" (sim round 15 R4: this character never drops below 20% mana); gems, runes and Robe sharing (sim rounds 15 and 17: category 1153, 120 s, one minute lockout).
- **Card conditions against "on cooldown", "every time" and "whenever"**: Arcane Wrath, Ice Lance, Cone of Cold, Combustion, Pyroblast, Evocation and Arcane's Fire Blast all say what their condition says. Fire's Fire Blast is R23-3.
- **Ranks far below level**: the rank-1 Frostbolt line (25 mana, 21 with Frost Channeling, 1.5 s less Improved Frostbolt, slow 5 s) claims only the slow for a few mana, which holds at 60. The low Scorch, Fire Blast and Frostbolt rank rows at 30 to 49 are accepted card actions, and no note or line makes a claim those ranks break.
- **Every superlative, cost, duration and cooldown on the four pages**: re-derived against the rows listed above (Scorch 150; Flamestrike 990, 3 s, 8 s; Blizzard 1,400 less 15% = 1,190, "about 1,200"; Ruby 1,100, Citrine 850; Ice Lance 8 s every rank; Arcane Wrath 3.5 s less 0.2, 6 s; Evocation 8 s, 8 min; Ice Block 10 s, 5 min; Cold Snap; Ice Barrier 1 min, 30 s; Arcane Power, Presence of Mind and Combustion 3 min; Pyroblast 6 s, 1.5 s on 3 stacks; Improved Frostbolt 2.5 s; Fire Blast 6 s with both points). No error.
- **Leveling paragraphs, spell and talent before the player has it**: Arcane, Fire and Frost, 1 to 19 and 20 to 60, against the decoded orders and trainer rows; all right. Wand levels 5 and 13 (`items.json`). Point totals: Arcane 31/4/16 and 32/16/3, Fire 17/31/3, Frost 15/5/31; prose point counts in every bold talent line add up.
- **Races and professions**: values against `spelltips/shared.json` and `races.json`; Touch of the Grave stays 30% on the mage row (round 14 D2: the shared tip's 20% is another class row, engine `racial_touch_of_the_grave.go`). Tome ranks: Fireball XII, Frostbolt XI, Arcane Missiles VIII are the rare-quality books, so "rank 8 is a book" and "top trainer rank" hold.
- **Source links**: all six distinct frontmatter URLs return 200 to a plain fetch.
- **Voice**: no em or en dashes, no site or sim words in body prose, no nightly numbers.

## Decisions

- **D1. MAGE-R23-1 is CONFIRMED** though round 22 passed the note: the sentence is false on the card's live bands whichever filler wins; the fix does not depend on the sim.
- **D2. MAGE-R23-2 touches Arcane and the class page only.** The rule from round 17 lets a Eureka! line stand unless it contradicts the client row for the class. Arcane Wrath is that contradiction, on hard mask evidence. "Hit" already excludes periodic damage, so no "non-periodic" clause is added.
- **D3. MAGE-R23-3 follows the card**, as the card is what was measured. If the weave sim wins, flip the row and drop the fix.
- **D4. MAGE-R23-4 stays NEEDS-SIM.** The sum is a first-order estimate that ignores regeneration, gems and Combustion. It is strong enough to ask, not to change the page.
- No web research was needed: the client rows and row order decide every call.

## Sims requested

1. **R23-4.** Fire, raid 60, shipped raid build and BiS, 180 s and 300 s, 10,000 iterations, seeds 1 and 50001: Fire Blast on cooldown (shipped) against Fire Blast gated above 30% mana, and against gated above 15%.
2. **R23-3.** Same set, 180 s: Fire Blast above the stack Scorch line (weave during the stack-up) against the shipped order.
3. **R23-1, optional.** Fire leveling order, bare, levels 16, 18 and 20, seeds 1 and 50001: the shipped Fire Blast then Fireball against a Missiles filler. Decides only the card action, not the wording fix.
4. Carried (ledger round 21): Arcane leveling order, bare, seeds 1 and 50001: Missiles first (shipped) against Frostbolt first at 12, 16, 20 and 22.
5. Carried (optional): Arcane raid order at 54 to 56, Presence of Mind on Pyroblast against on Frostbolt.

## Engine questions

- None new.

## Sim findings

- None new. Component lane, carried: the addon rows carry no talent gate.

## Note sweep

One line per note: spec, row, bands it renders on (`addon-data.json` rank rows), each named talent / spell / rank / value with the level the player has it (L = leveling order, R = raid order), verdict.

**mage-arcane**
1. Evocation (12051), bands 20 to 60: Evocation 20; 8 s, +1500%, 8 min; mana gem from 28 (Agate 28); potion any level; Demonic Rune at 60; last 30 s. OK.
2. Arcane Power (12042), bands 10 to 60 (talent L40, R41): Presence of Mind L31, R32; 30% damage, 30% cost, 15 s; gcd 0; Sunder. Condition 3 s or later matches "a Sunder or two". OK.
3. Presence of Mind (12043), bands 10 to 60 (L31, R32): "with Arcane Power once you have it"; Pyroblast once Heating Up (R57); Frostbolt before that and on L; Blast 1.5 s. OK.
4. Pyroblast on PoM, gated Heating Up, bands 20 to 60 as a row, live R57 to 60: 6 s every rank; most cast time PoM can save. OK.
5. Frostbolt on PoM, gated not Heating Up (L all, R to 56), bands 10 to 60: Blast 1.5 s against the global. OK.
6. Missile Barrage Missiles, gated (L22, R23), bands 10 to 60: free or 30% inside Arcane Power (L40, R41); half channel, a missile every 0.5 s; 40%/20% at 2/2 (L23, R24), half at 1/2; Arcane Wrath (R31) and Fireball as conditional proc sources; rank 7 top trainer rank (56), rank 8 the Ahn'Qiraj book (Tome of Arcane Missiles VIII). OK.
7. Missiles first, gated not 2/2 (live L10 to 22, R10 to 23), bands 10 to 60: second rank L23, R24; per mana over Blast "once you have it" (Blast 20); per second over Frostbolt ranks 2, 3, 4 at every live band. OK.
8. Pyroblast on 3 Heating Up, gated (R57), bands 20 to 60: 1.5 s on 3 stacks. OK.
9. Fire Blast, gated Heating Up (R57), bands 10 to 60: every hit a stack, a crit another; Improved Fire Blast 6 s (R55, 56); "on the way to 60" finisher. OK.
10. Arcane Wrath (1326082), bands 25 to 60 (talent R31): 3 Blast stacks, condition matches; 5% of maximum mana. OK.
11. Arcane Blast (1239700), bands 20 to 60: 3 stacks; mana against fight share and 20%, condition matches; 12% Intellect, +100% cost a stack; 18 s, no refresh; Clearcasting (15 to 19 both orders) frees a stacked Blast. OK.
12. Fireball in a raid, gated Heating Up (R57), bands 10 to 60: stacks' spell damage (Blast 20); Missile Barrage 20% (R24); Heating Up and Ignite (R49 to 53, R57). OK.
13. Frostbolt in a raid, gated Heating Up (R57), bands 10 to 60 as a row: a quarter less at 57 to 59 (260/350), a third at 60 (260/395); Missile Barrage 20%; no Ignite or Heating Up. OK.
14. Frostbolt filler on the way to 60, gated not Heating Up, bands 10 to 60: stacks' spell damage (Blast 20; Should before); a fraction of a stacked Blast (row reached at 3 stacks; 4 x fresh against every Frostbolt rank); 20% free Missiles once both ranks are in (L23, R24; Should inside Arcane Power); a cheap new set. OK.
15. Shoot (5019): no band renders it. Not rendered.

**mage-fire**
1. Evocation, bands 20 to 60: as Arcane 1. OK.
2. Combustion (11129), bands 40 to 60 (talent 40 both orders): Improved Scorch 5 stacks (3/3 by 29); 10% crit a hit until 3 non-periodic crits; Ignite (5/5 by L21, R19) and Heating Up (30). Condition matches. OK.
3. Scorch to 5 stacks, gated Improved Scorch (27), bands 22 to 60: 3% a stack, 5 stacks, 30 s, own damage only, a miss adds none; 1.5 s refresh, condition matches; filler refresh (33%/67% at 27/28, Should). OK.
4. Pyroblast on 3 Heating Up, gated (30), bands 20 to 60: Fire Blast hit plus Scorch (22), Fireball, Fire Blast crits; spends all; loses to Scorch at 6 s or 1 to 2 stacks (Should at 20, 21). OK.
5. Fire Blast, bands 10 to 60: instant; Heating Up 30; Improved Fire Blast both points by 26 (25, 26 both orders), 6 s; "In raid gear ... Scorch, Fireball or Pyroblast" holds (Checked). First sentence against row 3 from 27: MAGE-R23-3. Last clause: MAGE-R23-4.
6. Scorch filler, bands 22 to 60: Scorch 22; two per Fireball (Should at L22, 23); Ignite; Improved Scorch 3/3 at 29; Heating Up 30. OK.
7. Fireball, bands 10 to 60: Scorch 22; "the only filler you have" false beside Frostbolt 4 and Missiles 8. MAGE-R23-1.
8. Shoot: no band renders it. Not rendered.

**mage-frost** (one order, leveling and raid identical)
1. Evocation, bands 20 to 60: as Arcane 1. OK.
2. Ice Barrier (13033), bands 40 to 60 (talent 40): 15% for 1 min plus 5% shielded; biggest damage buff; 30 s cooldown stated; "the moment it breaks" (Should). OK.
3. Ice Lance (1240047), gated Fingers of Frost (36), bands 20 to 60: 8 s every rank, condition matches; four times its damage; Shatter 50% (3/3 at 34); two charges from 37. OK.
4. Missile Barrage Missiles, gated (51), bands 10 to 60: free (no Arcane Power), half channel, a missile every 0.5 s; 20% "with Missile Barrage 2/2" (52). OK.
5. Cone of Cold on 5 Winter's Chill, gated (38), bands 26 to 60: above half mana, within 10 yards (12 with Arctic Reach, conditional), condition matches; Frostbolt and Fingers of Frost Ice Lances feed it (36); double damage and 35% of spell damage. OK.
6. Frostbolt filler, bands 10 to 60: Fingers of Frost 36, Missile Barrage 51. OK.
7. Shoot: no band renders it. Not rendered.

## Sources

- data/builds/1.60.1.70790: talents/mage.json, trainables/mage.json, spelltips/mage.json, spelltips/shared.json, spellconst/mage.json, addon-builds.json, addon-data.json, items.json, races.json, bis/mage-{arcane,fire,frost}.json.
- data/curated/apl/mage-{arcane,fire,frost}.json, including the engineering notes on the Fire rows.
- Engine d5c708834: sim/mage/arcane_blast.go, sim/core/eureka_auto_gen.go.
- design/reviews/2026-11-04-sim-audit-round15/mage.md (G4, R4, X1), round 17 (S2), round 18 (F2), round 19, round 20.
- design/reviews/guide-accuracy/round22/mage.md, round14/mage.md (D2); ledger.md rounds 19 to 22; `git diff 987d08a7c..HEAD`.
- No web research was needed.
