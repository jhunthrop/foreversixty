# Guide accuracy round 20: Mage

Audited on the round 20 audit tree (4234643f5: main after guide round 19 and sim round 20; engine pin d5c708834), against the ACTIVE client 1.60.1.70790 under `data/builds/1.60.1.70790/`.

Files read line by line: `web/src/content/guides/mage/{index,arcane,fire,frost}.md`; every note, action and gate in `data/curated/apl/mage-{arcane,fire,frost}.json`; every rendered row of `addon-data.json` `rotations` (mage-arcane bands 10, 12, 14, 16, 18, 20, 22, 24, 25, 26, 30 to 50 by twos, 54, 56, 60; mage-fire 10, 12, 14, 18, 20, 22, 24, 28, 30, 36, 38, 40, 42, 46, 48, 50, 52, 54, 58, 60; mage-frost 10, 14, 16, 20, 24, 26, 28, 30, 32, 34, 38 to 52 by twos, 56, 58, 60), matched note by note with the rank id each band casts; all six orders in `addon-builds.json` decoded level by level (index 0 = level 10) and checked against the decoded frontmatter strings; `talents/mage.json` (every talent's rank 1 and top-rank text); `trainables/mage.json`; `spelltips/mage.json` (Evocation, Arcane Power, Presence of Mind, Arcane Blast, Arcane Wrath, Arcane Missiles 1 and 7, Fireball 4 to 11, Frostbolt 4, 5, 8 to 11, Scorch 1 to 4 and 7, Fire Blast 1, 3, 4, 7, Pyroblast 1 and 8, Frostfire Bolt 1 to 3, Ice Lance 1 to 6, Ice Barrier 1 and 4, Ice Block, Cold Snap, Cone of Cold 1 and 5, Frost Nova 1 and 4, Blizzard 6, Arcane Explosion 6, Flamestrike 6, Mage Armor 1 and 3, Combustion); `spellconst/mage.json` (400589 Missile Barrage, 12042 Arcane Power, 12536 Clearcasting, 400573 Arcane Blast stack); `simconsumes.json` and `items.json` (Major Mana Potion 13444, Mana Agate 5514, Mana Jade 5513, Mana Citrine 8007, Mana Ruby 8008, Demonic Rune 12662).

Read first: round 19's mage report, the sim round 20 mage report and integration.md, the ledger's round 19 section, rounds 1 to 18 as summarised there. The card machinery is as round 19 describes it: the page card prints the first clause of each note at 60 for the described build (`guide-rotation.ts`, `buildRunsGates`); the addon row carries the first sentence at every band with no talent gate. The Shoot rows are in the APLs but in no band of `addon-data.json`, so they render nowhere.

## Changes since round 19's tree (6beb919c2)

- Round 19 fixes (b84e48a2d, then 085a6691b's review wording): the three Evocation notes, Arcane Presence of Mind and its Frostbolt row, Arcane Missile Barrage Missiles, Arcane Missiles-first, Arcane leveling Frostbolt filler, Fire Fire Blast, Fire Scorch filler, Frost Ice Lance. Each re-checked band by band below; one is wrong (MAGE-R20-1).
- Sim round 20 (fd5a0aa9f): the Arcane raid build drops Arcane Geometry for Arcane Meditation 3/3 (`131205203100311531/525000012001/003`), with the Meditation bullet (arcane.md:55), the raid paragraph (:66) and the Human race line (:98) rewritten, and frost.md:99's race words. The regenerated raid order (0598dea66) matches the string: Arcane Meditation 28 to 30, Arcane Wrath 31, Presence of Mind 32, Arcane Power 41, Missile Barrage 23 and 24, Pyroblast 54, Improved Fire Blast 55 and 56, Heating Up 57. The leveling order is unchanged.
- All six frontmatter strings name 1.60.1.70790; no page or note names 70755.

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| MAGE-R20-1 | CONFIRMED | card note (rank) | `mage-fire.json`:257, the Scorch filler row (round 19's own fix), rendered at bands 24 (levels 24 to 27) and 28 (28 and 29) on both Fire orders | "from Improved Scorch at 27 every cast keeps your stacks rolling" | Improved Scorch rank 1: "Your Scorch spell has a 33% chance to cause your target to be vulnerable to Fire damage"; rank 3 is 100% (`talents/mage.json`, so rank 2 is 67%). Both Fire orders take it at 27, 28 and 29. At 27 one Scorch in three adds or refreshes a stack and at 28 two in three; "every cast" is true only from 29. The row is the live filler on those bands (no gate). Same shape as MAGE-R19-3, R19-6 and R19-7: the top rank credited before the order gives it. | "Scorch is the filler, not Fireball. You land two Scorches in the time of one Fireball, so you roll twice the crits for Ignite; once all three ranks of Improved Scorch are in at 29 every cast keeps your stacks rolling, and from 30 every crit counts for Heating Up too." |
| MAGE-R20-2 | CONFIRMED | card note (rank value) | `mage-arcane.json`:432, the raid Frostbolt row, gated on Heating Up, so live at 57 to 60 on the raid order (band 56 covers 57 to 59) | "It costs a third less" | Band 56 casts Fireball rank 10 (10150, 350 mana; rank 11 trains at 60) and Frostbolt rank 10 (10181, 260 mana): 26% less, a quarter. A third holds only at 60 against Fireball rank 11 (10151, 395 mana: 34% less). A level 60 figure quoted on a band where a lower rank renders, the case the round 20 lens names. arcane.md:74 ("costs a third less") describes 60 and stays. | "Frostbolt in a raid only when you are too low on mana for a Fireball. It costs a quarter less than the Fireball you have before 60 and a third less than the top rank, and still procs Missile Barrage 20% of the time, but it feeds no Ignite or Heating Up, so it only stretches the last of your mana." |
| MAGE-R20-3 | LIKELY | card note and prose (cost) | `mage-arcane.json`:209, the Missile Barrage Missiles row (live from 22 leveling, 23 raid; the claim is wrong inside Arcane Power from 40 and 41); arcane.md:53 | row: "The proc makes it free"; :53 "The proc makes your next Arcane Missiles free" | Missile Barrage 400589 effect 2 is aura 108 misc 14 at -100; Arcane Power 12042 effect 1 the same aura at +30; Clearcasting 12536 is -1000, the only -1000 cost mod in the mage file, where the druid's, shaman's and priest's are -100. Sim round 20 S2 reads that as an additive server (a -100 mod would already make anything free if mods multiplied) and the engine prices a Barrage Missiles under Arcane Power at 30% of its cost (`core/spell_mod.go`, `missile_barrage.go`). No tooltip states the combined cost, so LIKELY; Arcane casts every Arcane Power window into Barrage procs, so the page should not promise free Missiles there. The Frost page has no Arcane Power and its "free" stands; "makes a big share of your damage free" (:36) stays true outside the window. Take it. | Row: "Arcane Missiles the moment Missile Barrage lights up. The proc makes it free, or 30% of its cost inside Arcane Power, halves the channel and fires a missile every half second. ..." (rest unchanged). arcane.md:53: "The proc makes your next Arcane Missiles free (30% of its cost inside Arcane Power), halves its channel and fires a missile every half second." |

Counts: CONFIRMED 2, LIKELY 1, NEEDS-SIM 0, NEEDS-OWNER 0.

All three are card notes (R20-3 also one prose line). Every other note sentence holds at every band it renders on (Note sweep below), and every prose line on the four pages holds against the 70790 files, the shipped orders and the trainer. Fixes need `sim/request/apl/mage-*.apl.json`, `addon-data.json` and Data.lua regenerated (notes only; R20-1 and R20-2 are later sentences, so the addon rows do not change, only the page and hover text).

## Shoulds (wording, not findings)

New:
- **`mage-fire.json` Scorch stack row (bands 24 and 28, live from 27)**: "Filler Scorches refresh it anyway" is one cast in three at 27 and two in three at 28. The refresh row is a Scorch too, so the priority it explains still holds; not filed. With R20-1's fix the filler row now names 29, which covers it.
- **`mage-arcane.json` leveling Frostbolt row (bands 10 to 18)**: "The stacks' spell damage still applies to it" before Arcane Blast at 20 is idle (no stacks exist), the round 18 D1 case.
- **Arcane raid order 54 to 56**: the card spends Presence of Mind on Frostbolt while Pyroblast is already known (gate on Heating Up, 57). The note states the card truthfully; the action is a ladder question (Sims requested 1).

Carried from round 19, still wording: Fire Blast "most expensive spell for its damage" (D5 of round 19); Fire Pyroblast row 20 to 29; Fire Scorch stack row 22 to 26; Frost Cone of Cold row 26 to 37; Arcane Missiles-first rows 10 to 19 (round 18 Sims requested 3); Frost Ice Lance first sentence 30 to 35; arcane.md:72 "before the fight starts"; arcane.md:116 "from 23"; arcane.md:102 professions reasons; fire.md:94 Greater Arcane Elixir; fire.md:116 Burning Soul's 23%; frost.md:63 Winter's Chill feeders; frost.md:71 and the Ice Barrier note's "the moment it breaks" against a 30 s cooldown; frost.md:95 Missile Barrage's Missiles; index.md:32 Accelerant.

## Checked and right this round

- **Round 19's fixes**: Evocation ("mana gem from 28": Conjure Mana Agate 759 trains at 28; a mana potion has no level gate that matters; Demonic Rune 12662 has required level 0, and "at 60" is advice, not a gate); Presence of Mind and its Frostbolt row (raid PoM 32, Pyroblast 54, Heating Up 57: the card spends it on Frostbolt to 56 and Pyroblast from 57, as the note says); Missile Barrage Missiles ("half that with one rank": 20% and 10% at rank 1, 22 leveling and 23 raid); Missiles-first ("23 on the leveling order and 24 on the raid order", unchanged by the new raid order); leveling Frostbolt filler ("once both ranks of Missile Barrage are in"); Fire Blast ("by 26": both Fire orders 25 and 26); Frost Ice Lance ("From 37": Fingers of Frost 36 and 37, rank 1 "next 1 spell", rank 2 "next 2"). Fire Scorch filler: wrong (R20-1).
- **Sim round 20's prose**: arcane.md:7 decodes to 32/16/3 with Arcane Meditation 3, no Arcane Geometry; :55 (Meditation 17/33/50%, Mage Armor 50% on all three ranks: 6117 and 22783 spelltips, so all of it at 3/3); :66 (against the leveling build: Improved Channeling 4 to 3, Arcane Geometry 1 to 0 and one more Arcane point pay for Arcane Subtlety 2 and Arcane Wrath, 32 in Arcane; Fire 5+2+5+1+2+1 = 16; Precision 3); :98 ("all of that Spirit's regeneration": Meditation 3/3 on both builds plus Mage Armor); frost.md:99 (sim round 20 S5: Skyborne +0.7% over Gnome, Human -0.2%, Undead then Troll then Orc within 0.5%).
- **Sim round 20's engine changes against the pages**: a Presence of Mind Pyroblast no longer spends Heating Up (A1) and Clearcasting is not spent on a Barrage Missiles (S3); no line on any page says otherwise. arcane.md:54 "that includes a 4-stack Arcane Blast" holds under the additive rule (Clearcasting -1000 against +400 and +30).
- **Every other note at every band**: see the Note sweep.
- **Prose, all four pages**, against the 70790 files: every trainer level named (Frostbolt 4/8/14/20, Conjure Water 4, Fire Blast 6, Arcane Missiles 8, Frost Nova 10, Arcane Explosion 14, Fireball 1/6/12/18, Flamestrike 16, Scorch 22, Cone of Cold 26, Mage Armor 34, Evocation 20, Frostfire Bolt 40/50/60, Teleport: Dalaran 50); every talent level named against the orders (arcane.md:110 and :116, fire.md:110 and :116, frost.md:111 and :117); every talent value (Improved Fire Blast 2 s and 15 yards, Burning Soul 23/47%, 10/20%, Heating Up feeders, Fire Power 3% to 12%, Incineration's four spells, Arcane Mind's crit bonus, Ice Lance 300% increased, Shatter 50%, Fingers of Frost 15% and 2 spells, Winter's Chill 40% at 2 and its two feeders, Ice Barrier 15% and 5%, Arctic Reach 20% of 10 yards); spell values at 60 (Scorch 150 mana 1.5 s, Arcane Explosion 390 and 10 yards, Flamestrike 990, 3 s, 8 s, Blizzard 1400 less 15%, Frostfire Bolt 370 mana and 3 s against Frostbolt 290 and Scorch 1.5 s, Arcane Wrath 3.5 s less 0.2 and 6 s, Evocation 8 s, 1500%, 8 min, Arcane Power and Presence of Mind 3 min, Ice Block 5 min, Ice Barrier 30 s); point counts (Fire 31 with a 2-point toll; 17 Arcane and 3 Frost; Frost 31/15/5; Arcane leveling 31/4/16); stat prose explains and never ranks; race and gear lines as settled in rounds 14 to 19 (no change to bis or races since round 19).
- **Voice**: no em or en dashes, no site or sim words in body prose, no nightly numbers.
- **Source links**: all six frontmatter URLs answer HTTP 200 and say what their labels say (unchanged since round 18).

## Decisions

- **D1. Rank credits on a talent's own text decide it.** R20-1 follows round 19 D1: Improved Scorch's 33/67/100% is a per-rank value, so "every cast" before rank 3 is a finding, as Missile Barrage's and Improved Fire Blast's were.
- **D2. A cost ratio is a rank-dependent value.** R20-2: "a third less" is a level 60 figure on a band where Fireball rank 10 renders; filed under the round 20 lens, smallest finding.
- **D3. Missile Barrage under Arcane Power is LIKELY, recommended.** The client rows plus the sim lane's closed S2 make 30% the best reading; no tooltip states the sum, so not CONFIRMED. The fix costs nothing if the server turns out to multiply.
- **D4. Shoot notes render nowhere** (no Shoot row in any mage band of `addon-data.json`, and the page card reads the 60 band), so they are listed in the sweep and not judged by band.

## Sims requested

1. Arcane raid order at 54 to 56 (optional): Presence of Mind on Pyroblast (gate `spellIsKnown` Pyroblast) against on Frostbolt (gate on Heating Up, shipped), bare band 54 and 56, seeds 1 and 50001. If Pyroblast wins, the PoM rows' gates and notes move to "once you have Pyroblast".
2. Carried, already routed (ledger round 19): Fire Scorch against Fireball filler at 22 to 30 (MAGE-R19-9, fire.md:116). Not re-filed.

## Engine questions

- None. The engine already prices Missile Barrage under Arcane Power at 30% (sim round 20 S2) and Improved Scorch at 33/67/100% (sim round 20 F1).

## Sim findings

- None new. Component lane, carried: the addon rows carry no talent gate.

## Note sweep

One line per note: spec, row, bands it renders on, each named talent / spell / rank / value with the level the player has it (L = leveling order, R = raid order), verdict.

**mage-arcane**
1. Evocation (12051), bands 20 to 60: Evocation 20; 8 s, +1500%, 8 min (spelltip); mana gem from 28 (Conjure Mana Agate 28); mana potion (any level); Demonic Rune at 60 (no level gate); last 30 s. OK.
2. Arcane Power, bands 10 to 60 (talent L40, R41; earlier bands are the component's level filter): Presence of Mind L31, R32, both before Arcane Power; 30% damage, 30% cost, 15 s; off the global; Sunder. OK.
3. Presence of Mind, bands 10 to 60 (talent L31, R32): "with Arcane Power once you have it"; Pyroblast "once the raid build has Heating Up" (R57); Frostbolt before that and on L; Arcane Blast 1.5 s cast (400574 spelltip) and the 1.5 s global. OK.
4. Pyroblast on Presence of Mind, bands 20 to 60, gated Heating Up (R57 only): Pyroblast 6 s (11366 to 18809 spelltips), instant under PoM. OK.
5. Frostbolt on Presence of Mind, bands 10 to 60, gated not Heating Up (L all, R to 56): leveling build, raid before Heating Up; Arcane Blast 1.5 s. OK.
6. Missile Barrage Missiles, bands 10 to 60, gated Missile Barrage (L22, R23): free (MAGE-R20-3 inside Arcane Power from L40, R41); half the channel, a missile every 0.5 s (talent text, every Missiles rank); 40/20% at 2 ranks (L23, R24), half that at one; rank 7 top trainer rank (10212 at 56); rank 8 book. MAGE-R20-3.
7. Missiles first, bands 10 to 60, gated not Missile Barrage rank 2 (L to 22, R to 23): second rank L23, R24; more damage per mana than Blast or Frostbolt (bands 10 to 19 before Blast: carried Should); Blast takes over once both are in. OK.
8. Pyroblast on 3 Heating Up stacks, bands 20 to 60, gated Heating Up (R57): 3 stacks of 25% make 6 s into 1.5 s. OK.
9. Fire Blast, bands 10 to 60, gated Heating Up (R57): every hit a stack, a crit another (Heating Up text); Improved Fire Blast 6 s (R55 and 56, both before 57); "on the way to 60" finishes a mob. OK.
10. Arcane Wrath, bands 25 to 60 (talent R31; earlier bands the level filter): 3 Blast stacks (Blast 20); 5% of maximum mana (talent text). OK.
11. Arcane Blast, bands 20 to 60: Blast 20; 3 stacks, then the mana rule and 20%; 12% of Intellect, +100% cost a stack; 18 s from the first, no refresh (spelltip); Clearcasting (Arcane Concentration L15 to 19, R15 to 19) makes a stacked Blast free (-1000). OK.
12. Fireball in a raid, bands 10 to 60, gated Heating Up (R57): stacks' spell damage (Blast 20); Missile Barrage 20% (2/2 from R24); crits feed Heating Up (R57) and Ignite (5/5 R53). OK.
13. Frostbolt in a raid, bands 10 to 60, gated Heating Up (R57): "a third less" (band 56: Frostbolt 10181 260 against Fireball 10150 350, 26%; at 60 against 10151 395, 34%); Missile Barrage 20% (R24); no Ignite or Heating Up (Frost). MAGE-R20-2.
14. Frostbolt filler on the way to 60, bands 10 to 60, gated not Heating Up (L all, R to 56): stacks' spell damage (Blast 20; idle before, Should); a fraction of a stacked Blast; 20% once both Missile Barrage ranks are in (L23, R24); the next Blast starts a new set. OK.
15. Shoot (5019): no band renders it. Gem, rune or Evocation. Not rendered.

**mage-fire**
1. Evocation, bands 20 to 60: as Arcane 1. OK.
2. Combustion (11129), bands 40 to 60 (talent 40 on both orders): Improved Scorch 5 stacks (3/3 by 29); 10% Fire crit a hit until 3 non-periodic crits (talent text); each Scorch or Fire Blast crit an Ignite (5/5 by L21, R19) and a Heating Up stack (30); trinkets and racial; threat. OK.
3. Scorch to 5 Improved Scorch stacks, bands 22 to 60, gated Improved Scorch (27): 3% a stack, 5 stacks, 30 s, own Fire damage only, a miss adds none (talent text); 1.5 s refresh window; filler refreshes it anyway (33% at 27, 67% at 28: Should). OK.
4. Pyroblast on 3 Heating Up stacks, bands 20 to 60, gated Heating Up (30): feeders Fire Blast hit plus Scorch, Fireball, Fire Blast crits (talent text; Scorch from 22); the cast spends them (no Presence of Mind on Fire); 6 s or 1 to 2 stacks lose to Scorch (from 30). OK.
5. Fire Blast, bands 10 to 60: instant; Heating Up at 30; both Improved Fire Blast points by 26 (25 and 26, both orders), 8 s less 2; most expensive for its damage (carried Should). OK.
6. Scorch filler, bands 22 to 60: Scorch 22; two Scorches per Fireball (1.5 s against Fireball rank 4 3.0 s at 22 and 23, 3.5 s from rank 5); Ignite (5/5 by 21); "from Improved Scorch at 27 every cast keeps your stacks rolling" (rank 1 33% at 27, rank 2 67% at 28, rank 3 at 29); Heating Up from 30. MAGE-R20-1.
7. Fireball, bands 10 to 60: only before Scorch at 22. OK.
8. Shoot: no band renders it. Not rendered.

**mage-frost**
1. Evocation, bands 20 to 60: as Arcane 1. OK.
2. Ice Barrier, bands 40 to 60 (talent 40): 15% Frost damage for 1 min and 5% while shielded (every rank's spelltip, 11426 to 13033); 30 s cooldown (every rank); "the moment it breaks" (carried Should). OK.
3. Ice Lance, bands 20 to 60, gated Fingers of Frost (36): 8 s cooldown (every rank 1312002 to 1240047); frozen through the charge; four times its damage (300% increased); Shatter 50% (3/3 at 34); two charges from 37 (rank 2). OK.
4. Missile Barrage Missiles, bands 10 to 60, gated Missile Barrage (51): free (no Arcane Power in the build), half the channel, a missile every 0.5 s; Frostbolt 20% "with Missile Barrage 2/2" (52). OK.
5. Cone of Cold on 5 Winter's Chill stacks, bands 26 to 60, gated Winter's Chill (38): above half mana, within 10 yards; Frostbolt and Fingers of Frost Ice Lances build the stacks (talent text; Fingers of Frost 36); double damage and 35% more spell damage (20% and 7% a stack); 12 yards with Arctic Reach (not in the build, conditional: 20% at 2/2). OK.
6. Frostbolt filler, bands 10 to 60: Fingers of Frost 36, Missile Barrage 51. OK.
7. Shoot: no band renders it. Not rendered.

## Sources

- data/builds/1.60.1.70790: talents/mage.json (Improved Scorch, Missile Barrage, Fingers of Frost, Improved Fire Blast, Heating Up, Winter's Chill, Arctic Reach, Ice Lance, Shatter, Burning Soul, Arcane Meditation, Arcane Geometry and the rest quoted above), trainables/mage.json, spelltips/mage.json (ids listed in the header), spellconst/mage.json (400589, 12042, 12536, 400573), simconsumes.json (13444, 8007, 8008), items.json (5513, 5514, 12662), addon-builds.json, addon-data.json (rotations.mage-*).
- data/curated/apl/mage-{arcane,fire,frost}.json.
- web/src/lib/guides/guide-rotation.ts (`guideRotation`, `buildRunsGates`).
- design/reviews/2026-11-04-sim-audit-round20/mage.md (S2, S3, S5, A1, A2, F1) and integration.md; design/reviews/guide-accuracy/round19/mage.md; ledger.md round 19; rounds 1, 16 and 17 for the gem and rune cooldown sentence.
- `git show b84e48a2d 085a6691b fd5a0aa9f` over the mage pages and APLs.
- Frontmatter URLs fetched 2026-11-04, all HTTP 200: Blizzard Deep Dive panel recap (https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap), Wowhead Skyborne First Look, Power Up Gaming talent changes, Talents Forever data.json, foreversixty.gg/sim/weights and /sim/drops.
