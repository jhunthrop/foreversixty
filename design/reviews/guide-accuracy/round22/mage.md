# Guide accuracy round 22: Mage

Audited on the round 22 audit tree (987d08a7c: main after guide round 21; no sim round 21 commit or `*sim-audit-round21*` report in the tree; engine pin d5c708834 in `sim/enginever/version.go`), against the ACTIVE client 1.60.1.70790 under `data/builds/1.60.1.70790/`.

Files read line by line: `web/src/content/guides/mage/{index,arcane,fire,frost}.md`; every row and note in `data/curated/apl/mage-{arcane,fire,frost}.json`; all six orders in `addon-builds.json` decoded level by level (index 0 = level 10) against `talents/mage.json` (every talent, rank 1 and top rank text); `trainables/mage.json` (trainer levels of Arcane Missiles, Fire Blast, Scorch, Flamestrike, Blizzard, Cone of Cold, Arcane Explosion, Frostfire Bolt, Blink, Ice Armor, Mage Armor, Conjure Mana Agate); `spelltips/mage.json` (every rank of Fireball, Fire Blast, Scorch, Pyroblast, Arcane Missiles, Frostbolt, Frostfire Bolt, Flamestrike, Blizzard, Arcane Explosion, Cone of Cold, Ice Lance, Arcane Wrath, plus Evocation and the Mana Ruby and Citrine conjures); `spellconst/mage.json` (base, points per level, max level, variance and coefficient of Frostbolt 116/205/837/7322/8406/10181/25304, Missiles 5143/7268, 5144/7269, 10212/10274, 25346, Fire Blast 2136/2137/2138/10199, Scorch 10207, Fireball 10151/25306, Pyroblast 18809, Frostfire Bolt 401502/1237312/1237313, Arcane Blast 400574/1239700); `levels.json` (mage 22 and 60); `bis/mage-fire.json` (60 raid slots, both factions, summed against `items/mage.json`).

Read first: round 21's mage report and the ledger's Round 21 section (and rounds 19 and 20 for what is routed). Nothing routed under "Sims routed", "Sim findings" or "Engine questions" in rounds 19 to 21 is re-filed.

## Changes since round 21's tree (e039e1490)

`git diff e039e1490..HEAD` touches mage in two notes only (a29d7eb32, regenerated in 93973f2f9 / 39050340f): the Arcane Missiles-first note (R21-1) and the Fire Blast note (R21-2). No page, order, race, bis or trainer file changed for mage. Both re-checked below; R21-1 holds, R21-2's new wording is false against a spell the Fire mage has.

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| MAGE-R22-1 | CONFIRMED | card note: superlative that another trained spell beats (round 21's own fix; the "your biggest hit" shape) | `mage-fire.json` priority line 5, the Fire Blast row, third sentence; bands 10 to 60 on both Fire orders, the qualifier puts it at 60 in raid gear (page card at 60, sim drawer, hover) | "In raid gear it is your most expensive single-target spell for its damage" | Every mage trains Frostfire Bolt at 40, 50 and 60 (`trainables/mage.json`, skill_line_ability; fire.md:84 itself discusses it). At 60 rank 3 (1237313) costs 370 for 292 base (mean 314 with its 0.153 variance) plus a 57 burn that has no spell power coefficient, and 0.814 of spell power; Fire Blast rank 7 (10199) costs 340 for 453 at 60 (438 + 3 a level to 59, mean 491) and 0.429. The page's own Fire raid set (`bis/mage-fire.json` band 60 raid: 299 spell power and 71 Fire power on the Alliance items, 289 and 71 Horde) plus the preset's consumables (Flask of Supreme Power, Greater Arcane Elixir, Brilliant Wizard Oil, Elixir of Fire Power) is about 535 spell power and about 645 for Fire. Bare per mana: Fire Blast (491 + 0.429 x 645) / 340 = 2.26; Frostfire Bolt (371 + 0.814 x 535) / 370 = 2.18 if it takes only plain spell power, 2.42 if it takes Fire power too. With the Fire build's own talents the gap widens to Fire Blast: Incineration's 10% crit is Fire Blast's alone, and at about 25% against 15% crit, Ignite (crit worth 2.1x) and Master of Elements' 30% refund put Fire Blast at about 3.1 damage a mana against Frostfire Bolt's 2.95 at best (both school treatments, Improved Scorch and Fire Power given to both). At bare 60 it is 1.44 against 1.00. So in the page's raid gear Frostfire Bolt is the most expensive single-target spell for its damage, not Fire Blast; Fire Blast only drops behind it at about 950 spell power. Against the spells the Fire card casts the claim holds: Scorch about 3.15, Fireball 3.08, Pyroblast 4.3, all above Fire Blast's 2.26 in raid gear. | "Fire Blast every time it is off cooldown: it is instant, and from Heating Up at 30 every hit is a stack. Both points of Improved Fire Blast, by 26, make it a 6 second cooldown. In raid gear it costs more for its damage than Scorch, Fireball or Pyroblast, but skipping it to save mana costs more damage than the mana buys." |

Counts: CONFIRMED 1, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0.

The fix needs `sim/request/apl/mage-fire.apl.json`, `addon-data.json` and Data.lua regenerated (note text only; a later sentence, so only the page card, drawer and hover change).

## Shoulds (wording, not findings)

New:
- **fire.md:80 "two of them ... eat a third of your bar"**: a raid Human Fire mage has about 5,200 mana (1,213 base at 60 plus Intellect from the raid set, Arcane Brilliance, Mark of the Wild and Cerebral Cortex), so 1,980 is about 38%, nearer two fifths. Approximate wording, not filed.
- **fire.md:84 "takes twice as long"**: on the leveling build at 60 Accelerant's point makes Frostfire Bolt 2.9 s against Scorch's 1.5 s. Exact on the raid build, which the header card loads. Same shape as round 21's Scorch-filler Should.

Carried from round 21 (still wording): the Missile Barrage "free" lines inside Arcane Power (arcane.md:30, :36, Frostbolt filler note); Fire Scorch filler "two Scorches in the time of one Fireball" at 22 and 23 on the leveling order; Fire Pyroblast row 20 to 29; Fire Scorch stack row 27 to 28 ("Filler Scorches refresh it anyway"); Frost Cone of Cold row 26 to 37; Arcane Missiles-first rows 10 to 19 (now conditional, "once you have it"); Frost Ice Lance first sentence 30 to 35; arcane.md:72 "before the fight starts"; arcane.md:102; fire.md:94 Greater Arcane Elixir; fire.md:116 Burning Soul's 23%; frost.md:63 Winter's Chill feeders; frost.md:71 and the Ice Barrier note's "the moment it breaks" against a 30 s cooldown; frost.md:95 Missile Barrage's Missiles; index.md:32 Accelerant; the leveling Frostbolt row's "stacks' spell damage" before Blast at 20; Arcane raid PoM on Frostbolt at 54 to 56.

## Checked and right this round

- **R21-1 fix (Arcane Missiles-first note)**, live L10 to 22, R10 to 23. "More damage for its mana than Arcane Blast once you have it": conditional on bands 10 to 19 (Blast at 20 on both orders); at 20 to 23 Missiles rank 2 is 141 for 140 against a fresh rank 1 Blast's 34 to 36 for 56 to 65 (15% of base mana; 415 at 22), and on a long fight the stacked cost (2x, 3x, 4x) pulls Blast far below, even counting rank 1 Missile Barrage's free Missiles at 22. "More damage a second than Frostbolt": Missiles rank 1 24.4 to 24.8 a second (3 x 7268 over 3 s) against Frostbolt rank 2 17.4 (205, 1.8 s) at 10 to 13 and rank 3 22.1 to 22.5 (837, 2.2 s) at 14 and 15; Missiles rank 2 34.0 to 35.2 (4 x 7269 over 4 s) against rank 3 22.5 to 23.0 at 16 to 19 and rank 4 28.3 to 29.3 (7322, 2.6 s) at 20 to 23. Neither order has a Frost talent before 45 (leveling) or 58 (raid), and Arcane Subtlety (R21, 22) cuts resistance for both. Spell power closes it only past about 200 at 23. Holds at every band. "The second at 23 on the leveling order and at 24 on the raid order": right.
- **Every superlative and comparison on the four pages, against every spell trained at that level** (best / most / biggest / longest / cheapest / cheaper / fastest / only / never / every / more than / less than / twice / half / a third / a quarter / five times): arcane.md:34 best burst; :41 most mana-hungry, five times a fresh Blast (+100% a stack, 4 stacks); :56 Pyroblast your longest cast (6 s; nothing else under PoM's 10 s is longer); :74 Frostbolt a third less (260 against 395, 34%); :80 390 mana drains a bar in under a minute; :116 Missiles per mana over Blast at 20 to 22; fire.md:34 two Scorches per Fireball at 60 (3.0 s against 3.5 s); :72 Fireball as a filler loses to Scorch (measured, routed); :76 pull more than anyone (threat, consensus); :84 Frostfire Bolt costs more than Scorch (205/285/370 against 100 to 150 at 40 to 60); :116 Frostbolt still the cheaper slow (160 to 260 against 205 to 370); frost.md:34 biggest damage buff (Ice Barrier 15% plus 5% against any racial's 10%, and against an on-use 175 spell power trinket's about 16% for 15 s); :37 safest; :85 Frostfire Bolt hits for less than Frostbolt and costs more (rank 3 371 with its burn against rank 10 413, 370 against 260 less Frost Channeling); :117 easiest, fastest experience, Ice Lance crits more often than not from 34 (Shatter 50% plus base crit); index.md:22 best toolkit (opinion), :42 most damage into Arcane Power windows. No "far less than your best", "biggest hit" or "hardest hit" anywhere on the mage pages or notes.
- **Every quoted cost, duration and cooldown at the ranks and talents the builds have**: Scorch 150 (rank 7, no Fire cost talent); Arcane Explosion 390 (rank 6, 54); Flamestrike 990 (rank 6, 56) and 3 s cast, 8 s burn; Blizzard rank 6 1,400 less Frost Channeling's 15% = 1,190, "about 1,200"; Mana Ruby restores 1,100 (Conjure Mana Ruby 10054), so "1,200 mana down" wastes none; Frostbolt rank 1 25 mana, "a few mana"; Ice Lance 8 s every rank; Arcane Wrath 3.5 s less Accelerant 2/2 = 3.3 s, 6 s cooldown; Evocation 8 s, 8 min; Ice Block 10 s, Ice Barrier 1 min buff and 30 s cooldown, Arcane Power 15 s; Fire Blast 8 s less Improved Fire Blast 2/2 = 6 s; Pyroblast 6 s and 1.5 s on 3 Heating Up stacks; Improved Frostbolt 2.5 s Frostbolt; Arcane Meditation 2/3 "a third" (33%), 3/3 "half" (50%); Burning Soul 23% and 47% at one and two ranks; Winter's Chill 2/5 40%.
- **Leveling paragraphs, spell and talent named before the player has it** (round 22 shape): Arcane 1 to 19 (Frostbolt and Conjure Water 4, Fire Blast 6, Missiles 8, Frost Nova 10, Arcane Explosion 14; Wand Specialization 10, Improved Channeling 11 to 14, Arcane Concentration 15 to 19); Arcane 20 to 60 (Blast 20, Arcane Resilience 21, Missile Barrage 22 and 23, Arcane Power 40, then Incineration 41 to 44, Improved Frostbolt 45 to 49, Precision 50 to 52, Ice Shards 53 to 57, Frost Channeling 58 to 60; Blink 20; Arcane Explosion rank 2 at 22; Mage Armor 34, Ice Armor 30); Fire 1 to 19 (Fireball 1/6/12/18, Frostbolt 4, Fire Blast 6, Frost Nova 10, Flamestrike 16; Incineration 10 to 14, Accelerant 15, Ignite 16 to 19); Fire 20 to 60 (Pyroblast 20, Scorch 22, Burning Soul 24 on the leveling order, Improved Scorch 27 to 29, Heating Up 30, Frostfire Bolt 40); Frost 1 to 19 (Frostbolt 4/8/14/20, Improved Frostbolt 10 to 14, Frost Warding 15 and 16, Precision 17 to 19); Frost 20 to 60 (Ice Shards 20 to 24, Frostbite 25, Permafrost 26, Frost Channeling 27 to 29, Ice Lance 30, Ice Block 31, Shatter 32 to 34, Fingers of Frost 36 and 37, Ice Barrier 40, Cone of Cold 26, Blizzard 20). Every line names the level or sits in a paragraph that starts after it.
- **Talent sections and point totals**: Arcane leveling 31/4/16, raid 32/16/3 (the raid's three extra Arcane points paid by one Improved Channeling rank, Arcane Geometry and one Frost point, as arcane.md:66 says); Fire 17/31/3 on both orders (31 = 28 damage and range points plus a two-point toll: Burning Soul 1 and Accelerant 1 leveling, Burning Soul 2 raid; Magic Absorption to a third Arcane Resilience); Frost 15/5/31; every talent value against `talents/mage.json`.
- **Races, professions, enchants**: no race, bis or item file changed since round 19; prose re-read against the race table and the Fire 60 raid set (Spellpower Goggles Xtreme Plus at 270 Engineering, Robe of the Archmage, Frozen Heart of the Mountain).
- **Voice**: no em or en dashes, no site or sim words in body prose, no nightly numbers.

## Decisions

- **D1. MAGE-R22-1 is CONFIRMED, not LIKELY.** Round 21 took R21-2 as LIKELY because the Scorch crossing depended on gear; its fix pinned the claim to "raid gear". At the page's own raid set the claim fails against Frostfire Bolt under both readings of how Frostfire Bolt takes school spell power, and with or without the Fire talents. It only comes true at about 950 spell power, beyond anything the page ships. The fix keeps the intended point by naming the three spells the card casts.
- **D2. R21-1's fix stands.** The per-second reason holds against every Frostbolt rank the card casts beside it, bare and with any spell power a level 10 to 23 mage carries.
- **D3. The 38% Flamestrike line and Frostfire Bolt's 2.9 s on the leveling build stay Shoulds**: approximate wording true of the build the header loads.

## Sims requested

1. Carried, never run (ledger round 21): Arcane leveling order, bare, seeds 1 and 50001: Missiles first (shipped) against Frostbolt first at 12, 16, 20 and 22. R21-1's per-second reason does not depend on it.
2. Carried (optional): Arcane raid order at 54 to 56, Presence of Mind on Pyroblast against on Frostbolt.

## Engine questions

- None new. (How Frostfire Bolt takes spell power, plain or the higher of Fire and Frost, does not change R22-1.)

## Sim findings

- None new. Component lane, carried: the addon rows carry no talent gate.

## Note sweep

One line per note: spec, row, bands it renders on, each named talent / spell / rank / value with the level the player has it (L = leveling order, R = raid order), verdict.

**mage-arcane**
1. Evocation (12051), bands 20 to 60: Evocation 20; 8 s, +1500%, 8 min; mana gem from 28 (Conjure Mana Agate 28); potion any level; Demonic Rune at 60; last 30 s. OK.
2. Arcane Power (12042), bands 10 to 60 (talent L40, R41): Presence of Mind L31, R32; 30% damage, 30% cost, 15 s; off the global; Sunder. OK.
3. Presence of Mind (12043), bands 10 to 60 (L31, R32): "with Arcane Power once you have it"; Pyroblast once Heating Up (R57); Frostbolt before that and on L; Blast 1.5 s. OK.
4. Pyroblast on PoM, gated Heating Up (R57 to 60): 6 s rank 8; most cast time PoM can save (longest cast under 10 s). OK.
5. Frostbolt on PoM, gated not Heating Up (L all, R to 56): Blast 1.5 s against the global. OK.
6. Missile Barrage Missiles, gated Missile Barrage (L22, R23): free or 30% inside Arcane Power (L40, R41); half channel, a missile every 0.5 s; 40%/20% at 2/2 (L23, R24), half at 1/2; Arcane Wrath (R31) and Fireball named as proc sources, conditional; rank 7 top trainer rank (56), rank 8 the Ahn'Qiraj book. OK.
7. Missiles first, gated not Missile Barrage 2/2 (live L10 to 22, R10 to 23; first sentence on the addon row at every band): second rank L23, R24; per mana over Blast "once you have it" (Blast 20); per second over Frostbolt ranks 2, 3, 4 at every live band (Checked, above). OK (R21-1 fix holds).
8. Pyroblast on 3 Heating Up, gated (R57): 1.5 s on 3 stacks. OK.
9. Fire Blast, gated Heating Up (R57): every hit a stack, a crit another; Improved Fire Blast 6 s (R55, 56); "on the way to 60" finisher. OK.
10. Arcane Wrath (1326082), bands 25 to 60 (talent R31): 3 Blast stacks (Blast 20); 5% of maximum mana. OK.
11. Arcane Blast (1239700), bands 20 to 60: 3 stacks; mana against fight share and 20%; 12% Intellect, +100% cost a stack; 18 s, no refresh; Clearcasting (Arcane Concentration 15 to 19 both orders) frees a stacked Blast. OK.
12. Fireball in a raid, gated Heating Up (R57): stacks' spell damage (Blast 20); Missile Barrage 20% (R24); Heating Up and Ignite (R49 to 53, R57). OK.
13. Frostbolt in a raid, gated Heating Up (R57): a quarter less than Fireball before 60 (260 against rank 10's 350, 26%); a third less than the top rank at 60 (395, 34%; book 410, 37%); Missile Barrage 20%; no Ignite or Heating Up. OK.
14. Frostbolt filler on the way to 60, gated not Heating Up (L all, R to 56): stacks' spell damage (Blast 20; Should before); a fraction of a stacked Blast (every band); 20% free Missiles once both ranks are in (L23, R24; Should inside Arcane Power); a cheap new set. OK.
15. Shoot (5019): no band renders it. Not rendered.

**mage-fire**
1. Evocation, bands 20 to 60: as Arcane 1. OK.
2. Combustion (11129), bands 40 to 60 (talent 40 both orders): Improved Scorch 5 stacks (3/3 by 29); 10% crit a hit until 3 non-periodic crits; Ignite (5/5 by L21, R19) and Heating Up (30). OK.
3. Scorch to 5 Improved Scorch stacks, gated Improved Scorch (27): 3% a stack, 5 stacks, 30 s, own damage only, a miss adds none; 1.5 s refresh; filler refresh (33%/67% at 27/28, Should). OK.
4. Pyroblast on 3 Heating Up, gated (30): Fire Blast hit plus Scorch (22), Fireball, Fire Blast crits; spends all; loses to Scorch at 6 s or 1 to 2 stacks. OK.
5. Fire Blast, bands 10 to 60: instant; Heating Up 30; Improved Fire Blast both points by 26 (25, 26 both orders), 6 s; "In raid gear ... your most expensive single-target spell for its damage" (Frostfire Bolt rank 3, trained at 60, costs more per damage in the page's raid set). MAGE-R22-1.
6. Scorch filler, bands 22 to 60: Scorch 22; two per Fireball (Should at L22, 23); Ignite; Improved Scorch 3/3 at 29; Heating Up 30. OK.
7. Fireball, bands 10 to 60: only before Scorch at 22. OK.
8. Shoot: no band renders it. Not rendered.

**mage-frost** (one order, leveling and raid identical)
1. Evocation, bands 20 to 60: as Arcane 1. OK.
2. Ice Barrier (13033), bands 40 to 60 (talent 40): 15% for 1 min plus 5% shielded; biggest damage buff (Checked); 30 s cooldown; "the moment it breaks" (Should). OK.
3. Ice Lance (1240047), gated Fingers of Frost (36): 8 s cooldown; four times its damage; Shatter 50% (3/3 at 34); two charges from 37. OK.
4. Missile Barrage Missiles, gated Missile Barrage (51): free (no Arcane Power), half channel, a missile every 0.5 s; 20% "with Missile Barrage 2/2" (52). OK.
5. Cone of Cold on 5 Winter's Chill, gated Winter's Chill (38): above half mana, within 10 yards (12 with Arctic Reach, conditional); Frostbolt and Fingers of Frost Ice Lances feed it (36); double damage (5 x 20%) and 35% of spell damage (5 x 7%). OK.
6. Frostbolt filler, bands 10 to 60: Fingers of Frost 36, Missile Barrage 51. OK.
7. Shoot: no band renders it. Not rendered.

## Sources

- data/builds/1.60.1.70790: talents/mage.json, trainables/mage.json, spelltips/mage.json, spellconst/mage.json (ids above), levels.json (mage 22 and 60), addon-builds.json (mage-arcane, mage-fire, mage-frost, both orders), bis/mage-fire.json (band 60 raid, both factions; preset consumables), items/mage.json.
- data/curated/apl/mage-{arcane,fire,frost}.json.
- design/reviews/guide-accuracy/round21/mage.md, round20/mage.md, round19/mage.md; ledger.md rounds 19 to 21; `git diff e039e1490..HEAD`.
- No web research was needed: the client rows and the page's own raid set decide R22-1.
