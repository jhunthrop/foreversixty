# Guide accuracy round 17: Mage

Audited on the round 17 audit tree, main 919f7e8a7 (sixteen guide rounds plus sim round 18). Client 1.60.1.70755 (the live 1.60.1.70790 carries identical mage data), engine pin 6ea675594 (sim/enginever/version.go). Sim round 19 (int/r19, not on main) read for its Ice Lance ruling and its mage-arcane.json diff.

Files read line by line:
- web/src/content/guides/mage/index.md, arcane.md, fire.md and frost.md.
- Every card note, action and gate in data/curated/apl/mage-{arcane,fire,frost}.json, priorityList and prepullActions, and their rendered rows at levels 20, 22, 24, 30, 40 and 60 in addon-data.json.
- The mage rows of data/curated/rank-choices.json.
- addon-builds.json (all six mage orders), talents/mage.json (every rank text), trainables/mage.json, spelltips/mage.json, races.json and the Gnome row.
- Engine at the pin: sim/mage/ice_lance.go, fingers_of_frost.go, heating_up.go, constants_auto_gen.go.
- Sim round 18 mage report, sim round 19 mage report and integration.md, and `git -C r19/int diff 919f7e8a7 HEAD -- web/src/content/guides data/curated/apl`.

Read first: the mage reports for rounds 15 and 16 in full, the round 16 section of ledger.md, and rounds 1 to 14 skimmed.

## What changed since round 16

Main moved from 2c3135a2e to 919f7e8a7. The mage changes, all from sim round 18:
- **mage-arcane.json:** the Presence of Mind Frostbolt line (6) now checks `not auraIsKnown(400625)` so it belongs to the leveling build only (A2); the Missiles-first line (8) is gated on `not auraIsKnown(400588 tag 2)` and its note now reads "at 20 to 22 ... From 23" (A1); the Frostbolt filler note (14) now leads with the raid case (A3).
- **arcane.md:74:** "When you are too low on mana for a Fireball, cast a Frostbolt, which costs a third less; it feeds no Ignite or Heating Up, so it only stretches the last of your mana."
- **arcane.md:116:** "From 20 to 22, until both ranks of Missile Barrage are in, Arcane Missiles still does more damage for its mana than Blast, so lead with Missiles; from 23 you level on Blast and free Missiles."
- **rank-choices.json:** the mage rows are unchanged in content (Arcane Frostbolt 4 at 30 to 33; Fire Blast 3 at 30 to 39 and 4 at 40 to 49; Scorch 1/2/3 at 30 to 39, 40 to 45, 46 to 49; Frost Frostbolt 4 at 26 to 29 and 5 at 32 to 33).
- **Engine pin:** b2c4a2d05 to 6ea675594. The only mage engine change is the rank-tagged Missile Barrage trigger aura.

No change in addon-builds.json, talents/mage.json, trainables/mage.json, spelltips/mage.json, races.json, the mage BiS files, fire.md, frost.md or index.md.

**Sim round 19 (not on main)** sets Ice Lance's spell power coefficient to 0.107 on every rank (client column 0, nine beta players' logs) and splits the Arcane Frostbolt line into a raid line (gated on Heating Up known) and a leveling line. It changes no mage guide prose.

## Builds decoded

Unchanged from round 16 (talents/mage.json, tier/column order), all five legal at 51 points with every tier gate and prerequisite met at the level the shipped order takes it:
- **Arcane leveling 31/4/16:** Wand Specialization 1, Improved Channeling 4, Arcane Resilience 1, Arcane Concentration 5, Missile Barrage 2, Arcane Geometry 1, Arcane Impact 3, Arcane Blast 1, Arcane Meditation 3, Presence of Mind 1, Arcane Mind 5, Arcane Instability 3, Arcane Power 1 / Incineration 4 / Improved Frostbolt 5, Precision 3, Ice Shards 5, Frost Channeling 3.
- **Arcane raid 32/16/3:** the leveling Arcane points with Improved Channeling 3 and Arcane Meditation 2, plus Arcane Subtlety 2 and Arcane Wrath 1 / Incineration 5, Accelerant 2, Ignite 5, Pyroblast 1, Improved Fire Blast 2, Heating Up 1 / Precision 3.
- **Fire leveling 17/31/3:** Wand Specialization 2, Improved Channeling 5, Arcane Resilience 2, Magic Absorption 1, Arcane Concentration 5, Arcane Meditation 2 / Incineration 5, Accelerant 1, Ignite 5, Flame Throwing 2, Burning Soul 1, Pyroblast 1, Improved Fire Blast 2, Improved Scorch 3, Heating Up 1, Master of Elements 3, Critical Mass 3, Fire Power 3, Combustion 1 / Precision 3.
- **Fire raid:** as leveling with Arcane Resilience 3 for Magic Absorption and Burning Soul 2 for Accelerant 1.
- **Frost 15/5/31, both strings:** Wand Specialization 2, Improved Channeling 3, Arcane Concentration 5, Missile Barrage 2, Arcane Impact 3 / Incineration 5 / Frost Warding 2, Improved Frostbolt 5, Precision 3, Ice Shards 5, Frostbite 1, Permafrost 1, Frost Channeling 3, Ice Lance 1, Ice Block 1, Shatter 3, Cold Snap 1, Fingers of Frost 2, Winter's Chill 2, Ice Barrier 1.

Point counts the prose quotes, re-added this round: Fire's damage talents plus Flame Throwing 2 = 28, plus the 2-point toll and Combustion = 31 (fire.md:62); Fire's 17 Arcane reach Arcane Meditation at 15 (fire.md:64); Frost's 31 with Winter's Chill 2 as the toll to 30 above Ice Barrier (frost.md:63); Arcane Impact at 10 Arcane points (frost.md:63); Arcane raid 32 = 31 minus a rank each of Meditation and Improved Channeling plus Subtlety 2, Wrath 1 and one more (arcane.md:66); Arcane's 16 Frost = Improved Frostbolt 5, Precision 3, Ice Shards 5, Frost Channeling 3 (arcane.md:42).

## Leveling orders against the prose (addon-builds.json, index 0 = level 10)

Every level the prose names matches the shipped order, and the Leveling strip (now the 10 to 60 talent tree) shows the same points:
- **Arcane:** Wand Specialization 10, Improved Channeling 11 to 14, Arcane Concentration 15 to 19; Arcane Blast 20, Arcane Resilience 21, Missile Barrage 22 and 23; every point Arcane to Arcane Power at 40; Incineration 41 to 44, Improved Frostbolt 45 to 49, Precision 50 to 52, Ice Shards 53 to 57 ("by 57"), Frost Channeling 58 to 60.
- **Fire:** Incineration 10 to 14, Accelerant 15, Ignite 16 to 19 (fifth rank at 21); Pyroblast 20, Burning Soul 24, Improved Scorch 27 to 29, Heating Up 30 ("from 30"), Combustion 40.
- **Frost:** Improved Frostbolt 10 to 14, Frost Warding 15 and 16, Precision 17 to 19; Ice Shards 20 to 24, Frostbite 25, Permafrost 26, Frost Channeling 27 to 29, Ice Lance 30, Ice Block 31, Shatter 32 to 34 ("by 34"), Cold Snap 35, Fingers of Frost 36 and 37, Winter's Chill 38 and 39, Ice Barrier 40.

Trainer levels (trainables/mage.json) as the pages state: Frostbolt 4/8/14/20, Fireball 1/6/12/18, Conjure Water 4, Fire Blast 6, Arcane Missiles 8, Frost Nova 10, Arcane Explosion 14, Flamestrike 16, Remove Lesser Curse 18, Blink 20, Scorch 22, Counterspell 24, Cone of Cold 26, Ice Armor 30, Mage Armor 34, Frostfire Bolt 40/50/60, Teleport: Dalaran 50, Arcane Brilliance 56.

## Findings

None.

Counts: CONFIRMED 0, LIKELY 0, NEEDS-SIM 0 new, NEEDS-OWNER 0.

MAGE-R15-1 (the Missiles to Blast boundary) is closed: sim round 18 A1 measured level 22 (Missiles first 49.0 against 41.4) and 23 (Blast first 51.0 against 50.2), and the page and the card note now say "20 to 22" and "from 23". Sim round 19 S1 re-measured 22, 23 and 24 and it holds.

## Covered by sim round 19

Nothing to file. The round 19 mage-arcane.json diff splits the Frostbolt line into:
- a raid line gated on `auraIsKnown(400625)`: "Frostbolt in a raid only when you are too low on mana for a Fireball. It costs a third less and still procs Missile Barrage 20% of the time, but it feeds no Ignite or Heating Up, so it only stretches the last of your mana."
- a leveling line gated on `not auraIsKnown(400625)`: the old note's leveling sentences.

Both notes are true: Frostbolt rank 10 is 260 mana against Fireball rank 11 at 395 (34% less); Heating Up feeds on Fireball, Frostfire Bolt, Fire Blast and Scorch only; Ignite is Fire crits only; Missile Barrage 2/2 is 20% on Frostbolt. The r19 wording matches arcane.md:74 on main.

## New text this round, checked

- **arcane.md:74, the Frostbolt sentence.** Frostbolt r10 (10181) 260 mana, Fireball r11 (10151) 395 mana: a third less. Frostbolt is not in Heating Up's feeder list (talent 400624: Fireball, Frostfire Bolt, Fire Blast, Scorch) and is Frost, so no Ignite. The card casts those two ranks (the rank 11 Frostbolt and rank 12 Fireball at 290 and 410 are the level-60 book ranks the cards do not use).
- **arcane.md:116 and card line 8, "20 to 22 ... from 23".** Missile Barrage is taken at 22 and 23 in the shipped order, so 22 has one rank (20% on Blast, 10% on Frostbolt). Measured at 22 and 23 by round 18 and again by round 19. The sentence does not contradict Levels 1 to 19 ("From 8, open with Frostbolt ... channel Missiles as it closes") or the "On the way to 60" loop at :74, which describes play once you blast.
- **Card line 6 gate (leveling only):** true to the page. arcane.md:56 and :72 give the Presence of Mind Frostbolt to the leveling build and the Pyroblast to the raid build, and only the raid build has Heating Up.
- **Card line 14 note (raid first):** the raid case is what the page's raid loop says; the leveling half matches "Blast to 3 stacks, Frostbolt as the filler" and "The stacks do not refresh, so the next Blast starts a fresh, cheap set".

## Ice Lance, every line, against round 19 and the client

Round 19 measured Ice Lance's spell power coefficient at 0.107 (client SpellEffect column 0 on every rank; the engine's 0.4286 at the pin was the generator's instant-spell default). The Fingers of Frost Ice Lance line still gains 2.5% over dropping it (R2), so the rotation stands. Every Frost line about Ice Lance was read against that:
- **frost.md:54** "four times its damage on a frozen target ... 8 second cooldown in Forever": talent 1312002 and every spelltips rank say "300% increased damage to Frozen targets", cooldown_ms 8000 on 1240047. The engine multiplies the whole hit by 4 (ice_lance.go, iceLanceFrozenMultiplier) and the logs' frozen-to-unfrozen ratio is 4.0. No spell power claim.
- **frost.md:35, :71, :83, :117 and the Ice Lance card note:** "hits for four times its damage with Shatter's crit", "Spend a Fingers of Frost charge on Ice Lance on the run", "After a Frost Nova or a Frostbite freeze, Ice Lance hits a mob for four times its damage": all base-damage statements, none names spell power.
- **frost.md:71** "Fingers of Frost gives you two charges and Ice Lance takes one, so your next Frostbolt spends the other with the same crit bonus": talent 400647 rank 2 is 2 spells; the engine spends one charge per direct-damage cast as it resolves (fingers_of_frost.go) and adds Shatter's 50% to each.
- **frost.md:89 and :95** "Spell power goes into everything you cast" and "Frost spell power on gear and an Elixir of Frost Power are close to full value": Frost power counts the same as general spell power on Ice Lance, whatever the coefficient, and Frostbolt (0.814) carries most of the damage. Neither line quotes a coefficient or ranks Ice Lance's scaling. Not findings.
- **frost.md:63** "five ranks of Incineration, which adds 10% crit to Ice Lance": talent 18459 rank 5 names Ice Lance.
- **frost.md:99 (Skyborne by a hair over Gnome):** measured in round 3 on the old coefficient. Haste matters more as the instant's share falls, so the direction does not move against the page; a re-run is requested below, not filed.

## Re-checked this round (all hold)

- **Talent text (talents/mage.json), every value the pages quote:** Arcane Blast 12% Intellect a stack, +100% cost, 4 stacks, 18 s; Missile Barrage 40/20% at 2 ranks; Arcane Concentration 10%; Arcane Meditation 17/33/50%; Presence of Mind under 10 s; Arcane Power 30%/30%, 15 s; Arcane Mind 10% Intellect and +100% Arcane crit bonus; Arcane Instability 3%/3%; Arcane Subtlety 15 resistance and 30% Arcane threat at 2; Improved Channeling 80% Missiles and 56% Blast at 4; Arcane Geometry 3 yards; Arcane Impact 6%; Arcane Wrath 5% of maximum mana; Wand Specialization 13%; Arcane Resilience armor from Intellect; Incineration 10% on Fire Blast, Ice Lance, Arcane Blast and Scorch; Accelerant 0.1 s a rank; Ignite 40% over 4 s; Flame Throwing 6 yards; Burning Soul 23%/10% and 47%/20%; Improved Fire Blast 2 s and 15 yards; Improved Scorch 3% a stack, 5 stacks, 30 s, 100% on a landed Scorch at 3 ranks; Heating Up 25% a stack, 3 stacks, Fire Blast hits add a stack; Master of Elements 30%; Critical Mass 6%; Fire Power 3% an effect to 12%, mixed schools excluded; Combustion 10% a hit until 3 non-periodic crits; Impact stun; Frost Warding 30%; Improved Frostbolt 0.5 s; Precision 6%; Ice Shards +100% bonus; Frostbite freeze; Frost Channeling 15%/30%; Ice Block 10 s; Shatter 50% at 3 ranks; Cold Snap finishes Frost cooldowns; Fingers of Frost 15%, 2 spells, 15 s; Winter's Chill 40% at 2 ranks from Frostbolt and empowered Ice Lance, 20% Cone of Cold damage and 7% spell damage a stack, 5 stacks; Ice Barrier 15% for 1 min, 5% while shielded, no pushback; Improved Blizzard and Improved Cone of Cold exist.
- **Heating Up at three stacks is a 1.5 second Pyroblast:** the engine subtracts 0.25 of the cast per stack (heating_up.go), so three stacks leave a quarter of 6 s.
- **Spell values (spelltips/mage.json):** Presence of Mind, Arcane Power and Combustion 180 s; Evocation 480 s, 8 s, +1500%; Ice Block 300 s; Cold Snap 600 s; Ice Barrier r4 30 s; Fire Blast r7 340 mana, 8 s; Ice Lance 8 s every rank; Scorch r7 150 mana, 1.5 s; Flamestrike r6 990 mana, 3 s cast, 8 s burn; Arcane Explosion r6 390 mana, 10 yards; Blizzard r6 1,400 mana (1,190 with Frost Channeling, "about 1,200"), 8 s; Frostfire Bolt r3 370 mana, 3 s, 292 plus 57 against Frostbolt r11's 475; Frostfire Bolt r1 205 against Frostbolt r7 160; Arcane Blast 1.5 s at 15% of base mana; Arcane Wrath r4 3.5 s, 6 s cooldown; Mana Ruby 1,100; Mage Armor 50% regeneration at every rank; Cone of Cold r5 slows 40% for 6 s; Pyroblast r8 6 s.
- **Rank rows (rank-choices.json):** every rank is trainable at its band (Fire Blast 3 at 22, 4 at 30; Scorch 2 at 28, 3 at 34; Frostbolt 4 at 20, 5 at 26), and no prose names a rank below the top one except the rank-1 Frostbolt slow (frost.md:75), which is true at any level.
- **Card notes against their actions and gates (all three APL files):** each note describes its own condition. Arcane: Evocation under 5% with more than 30 s left; Arcane Power and Presence of Mind from 3 s ("Blast twice ... then press"); Pyroblast on Presence of Mind only with Heating Up, Frostbolt only without; Missiles on the Barrage aura; Missiles-first below Missile Barrage 2; Pyroblast on 3 Heating Up stacks; Fire Blast with Heating Up known; Arcane Wrath at 3 stacks; Arcane Blast below 3 or while mana percent beats the fight's remaining percent and exceeds 20%; Fireball with Heating Up known; Frostbolt last. Fire: Combustion at 5 Improved Scorch stacks; Scorch to stack or refresh at 1.5 s left; Pyroblast on 3 stacks; Fire Blast on cooldown; Scorch filler; Fireball before 22. Frost: Ice Barrier when down (pre-pull at -4 s); Ice Lance on a known and active Fingers of Frost aura; Missiles on Barrage; Cone of Cold at 5 Winter's Chill stacks above 50% mana; Frostbolt filler (pre-pull at -2.5 s). The pre-pull Pyroblast at -6 s is gated on Heating Up and the Frostbolt at -2.5 s on its absence, as arcane.md:72 says.
- **Rendered rows (addon-data.json):** every row's condition text is the first sentence of its curated note, for all three specs at every level checked. Rows shown at levels before their build takes the talent (Pyroblast on the Arcane card at 20, Heating Up on the Fire card at 20) are the component's level filter, not a note error.
- **Opener and loop wording against real play:** Arcane's two Blasts before Arcane Power take 3 s, the line's gate; Fire's Combustion waits for 5 stacks; Frost's Ice Barrier is up before the Frostbolt that lands at the pull.
- **Race lines:** every racial value matches races.json: Eureka! (the Gnome row reads "damaging and healing abilities cost 10% less and deal 10% more for 8 seconds on a 2 minute cooldown", so "damage spells hit 10% harder and cost 10% less" is the mage's share of it), Expansive Mind 5%, Ruthless 1% with all spells and attacks, The Human Spirit 10%, Touch of the Grave up to 5% of maximum health once per 15 s (30% on the mage row, rounds 14 and 15), Will of the Forsaken breaks fears, Blood Fury 10% spell power 15 s every 2 min, Berserking 10% casting speed 15 s every 3 min, Wind Blessed 2% haste, Read Leyline 50% health and mana regeneration 15 s every 2 min. Every order and closeness word matches round3/race-sims.md, the only measured mage race run. Skyborne is named only as the paid option. Mage races in combos.json: Human, Gnome, Undead, Orc, Troll and High Order Skyborne.
- **Stat priority prose:** explains and never ranks on all three pages; each reason holds (hit to the cap with Precision's 6%; Intellect as spell damage through Arcane Blast; Arcane Mind doubling the Arcane crit bonus; Ice Shards doubling the Frost crit bonus; Shatter's 50% on a Fingers of Frost charge; Fire's spell power and Fire power the same stat on a boss; Master of Elements on crits).
- **Gear and professions prose:** names only bind-on-pickup crafts (Robe of the Archmage, Tailoring; Frozen Heart of the Mountain, Enchanting) and the goggles' Engineering 270, all in the live Fire and Frost raid rows (round 15), and no Arcane row carries an Engineering item.
- **The gem, rune and Robe sentence (arcane.md:78, fire.md:78, frost.md:75):** the era spell rows for 10052, 10057, 10058 (gems), 16666 (Demonic Rune), 27869 (Dark Rune) and 18385 (the Robe) each carry a 60 s category 30 recovery and no own recovery, which is "using any of the three locks the other two out for a minute". The gems' and runes' own 2 minute cooldown and the Robe's 5 minute one are item cooldowns, which items.json does not carry; they rest on the round 1 research and sim round 18 S1. Not re-filed.
- **index.md:** Ice Block tier 3 (15 points) with Cold Snap tier 4; no Molten Armor or Invisibility in trainables; Mage Armor 50%; Precision all schools; the new talents and the Frostfire Bolt and Teleport: Dalaran levels.
- **Source links:** all seven fetched 2026-11-04 and returned HTTP 200; labels follow the round 4 rule.
- **Voice:** no em or en dashes; no site, sim, APL, band, nightly or pipeline words in body prose; no quoted nightly numbers.

## SHOULDs (wording, not findings)

Carried from rounds 14 to 16, still wording:
- **arcane.md:72:** "it can roll Missile Barrage before the fight starts". The Frostbolt lands as the fight starts.
- **arcane.md:74:** "On the way to 60 ... Blast to 3 stacks" could say "from 23", now that :116 leads with Missiles to 22.
- **arcane.md:102:** the professions reasons; the real ones are the Robe and the Brilliant Wand.
- **fire.md:94:** "Greater Arcane Elixir is your elixir" beside a table that also lists Elixir of Fire Power.
- **fire.md:116:** Burning Soul's 23% "keeps a mob's hits from pushing back every cast".
- **frost.md:63:** Winter's Chill also procs from a Fingers of Frost Ice Lance (the card note at :79 says so; the talent paragraph does not).
- **frost.md:71:** "Recast Ice Barrier the moment it breaks" against its 30 s cooldown.
- **index.md:32:** Accelerant is missing from the list of new talents.
- **New, frost.md:95:** "Frost casts Frostbolt and Ice Lance for nearly all its damage" could name Missile Barrage's Missiles as the Arcane share that Frost power misses, which :89 already does.

## Decisions

- **D1.** Sim round 18's three Arcane edits (the Frostbolt sentence at :74, "20 to 22 / from 23" at :116, the card gates and notes) are true on the 70755 client values and measured at 22, 23 and 24 on two seeds. Nothing to file.
- **D2.** Ice Lance's 0.107 coefficient changes no mage prose: no line states a spell power claim for Ice Lance, and round 19 R2 keeps its rotation line (+2.5% over dropping it). The frozen multiplier is ×4 on the whole hit, which the talent text, the engine and the logs agree on.
- **D3.** The round 19 mage-arcane.json split is right in both notes; nothing goes in the "Covered by sim round 19" table.
- **D4.** Arcane Missiles rank 8 (25345), Fireball rank 12 (25306), Frostbolt rank 11 (25304) and Conjure Water rank 8 (468766) sit in trainables/mage.json as skill-line rows at 56 and 60. The card note calls rank 8 an Ahn'Qiraj book, per sim round 4 and the round 18 review, and the level-60 book ranks are on the owner's list (round 19). No page states a rank for these, so there is nothing to file; the note follows the ruling.
- **D5.** The gem and rune 2 minute cooldown is not in the data this round reads (item cooldowns are not exported); the era spell rows support the one-minute lockout sentence. Round 18 S1 kept the wording on consensus. Not a finding.
- **D6.** Round 16 D2 and D3 stand (Fire card 5's "most expensive spell for its damage" judged at 60; Boughs in the Wind stays out of the level 10 quest sentence).

## Sims requested

- **Frost races at 60 on the shipped 15/5/31 build with the 0.107 Ice Lance coefficient** (seeds 1 and 50001, both factions): Gnome, Human, Skyborne; Undead, Orc, Troll. frost.md:99 rests on round 3's run at the old coefficient; the direction should hold (haste gains as the instant's share falls), but the page's "by a hair" and "close behind" words want the new numbers.
- **Carried, data lane:** the Frost BiS and weights re-rank after the Ice Lance change (round 19 "Measure at refresh"); the Frost band 60 rows ranked on the round 2 string; the round 14/15 BiS items (three-profession Fire and Frost sets, Frozen Heart for Arcane and Fire, the Arcane raid rows on 32/3/16).
- **Carried, sim lane:** Arcane and Frost race re-runs on the shipped builds; the Frost raid talent search after the coefficient change (Arcane Impact against the Shatter and Fingers of Frost points, round 19); the Arcane raid pre-pull on two seeds; the Presence of Mind target; Arcane Wrath at 3 against 4 stacks; S1, S2, SIM-R2-1, SIM-R2-3, SIM-5, SIM-6, SIM-2; the cost-mod stacking question (Arcane Power on a Barrage Missiles and on stacked Blast, round 19 S5, settled in game at 40+).
- **Closed:** MAGE-R15-1 (round 18 A1, round 19 S1).

## Sim findings

- None new. SIM-R7-1 is closed by round 19 S4: Arcane Subtlety's 15 resistance cut has no effect on a 0-resistance target whose level-based part is not penetrable, so the engine and the page ("helps only your own spells") agree.
- Recorded, not a guide finding: Ice Lance's 0.107 comes from beta logs at levels 24 to 35 (round 19 S3); a level-60 launch log should re-check it.

## Sources

- data/builds/1.60.1.70755:
  - talents/mage.json: every rank text, tier, column and prerequisite.
  - trainables/mage.json: every spell the pages name, plus the four book ranks.
  - spelltips/mage.json: 1312002, 400640, 1240044, 1240047, 7322, 8408, 10181, 25304, 10151, 25306, 5144, 10212, 25345, 10207, 10199, 18809, 12051, 12043, 12042, 11129, 12472, 13033, 10216, 10187, 10202, 1237313, 401502, 10161, 10054, 10053, 6117, 22782, 22783, 400574, 1239700, 1326082, 11958, 1953, 6131, 475, 2139.
  - races.json (the Gnome row and the per-race racial texts), addon-builds.json (all six mage orders), addon-data.json (rotations for the three specs), items.json (8007, 8008, 12662, 20520, 14152), bis/mage-*.json as round 15 read them.
- data/curated/apl/mage-{arcane,fire,frost}.json and data/curated/rank-choices.json (mage-arcane, mage-fire, mage-frost rows).
- Engine 6ea675594: sim/mage/ice_lance.go (frozen multiplier, 8 s cooldown, IceLanceSpellCoeff 0.4286 at the pin), fingers_of_frost.go (2 charges, one per direct-damage cast, Shatter 17/33/50), heating_up.go (0.25 a stack), constants_auto_gen.go (Frostbolt and Fireball mana by rank).
- Era client tables in the main checkout (/Users/jh/code/forever/SpellCooldowns_era.csv): 10052, 10057, 10058, 16666, 18385, 27869, 10199, 400640.
- design/reviews/2026-11-04-sim-audit-round18/mage.md (A1, A2, A3, A4, F1, F2, R1, R2); the r19 tree's design/reviews/2026-11-04-sim-audit-round19/mage.md (S1 to S8, A1, A2, F1, R1, R2, X1, X2) and integration.md; `git -C r19/int diff 919f7e8a7 HEAD`.
- design/reviews/guide-accuracy/round1 to round16/mage.md, round3/race-sims.md and ledger.md.
- Source links fetched 2026-11-04, all HTTP 200:
  - Blizzard, Deep Dive panel recap: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
  - Wowhead, Skyborne First Look: https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829
  - Power Up Gaming, Forever talent changes vs Classic: https://powerupgaming.co.uk/2026/09/14/new-wow-forever-talents-changes-vs-classic-guide/
  - Talents Forever data export: https://talentsforever.com/data.json
  - https://foreversixty.gg/sim/weights, https://foreversixty.gg/sim/drops and https://foreversixty.gg/planner?class=mage

## Counts

CONFIRMED 0, LIKELY 0, NEEDS-SIM 0 new (1 closed), NEEDS-OWNER 0. Covered by sim round 19: 0. SHOULDs: 9 (8 carried, 1 new).

## Fix log (round 17)

- No finding. The Frost race re-run on the shipped build with the new Ice Lance coefficient (Sims requested 1) goes to the sim lane.
