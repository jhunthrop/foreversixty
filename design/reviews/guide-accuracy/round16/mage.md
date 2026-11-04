# Guide accuracy round 16: Mage

Audited on the round 16 audit tree, land/guides-r15 2c3135a2e (fifteen guide rounds, launch state). Client 1.60.1.70755, engine pin b2c4a2d05 (sim/enginever/version.go).

Files read line by line:
- web/src/content/guides/mage/index.md, arcane.md, fire.md and frost.md.
- Every card note and action in data/curated/apl/mage-{arcane,fire,frost}.json, priorityList and prepullActions.
- bis/mage-frost.json band 60 raid row and presets (consumables), as a spot check on the stat and consumables prose.

Read first: the mage reports for rounds 1 to 15 (round 15 in full, with its decisions and sims requested), round3/race-sims.md (the only measured mage race run), and the round 15 section of ledger.md.

## What changed since round 15

Nothing for mage. Main moved from 4600b7d82 to 2c3135a2e with two commits:
- b2aee8a38, the round 15 fixes. Its APL edits touch rogue-combat, rogue-subtlety and shaman-restoration only.
- a61b0ff9f, addon 2.0.25 (CHANGELOG and TOC).

There are no diffs since 4600b7d82 in any of these:
- the mage guides and mage APL files
- addon-builds.json, talents/mage.json, trainables/mage.json, spelltips/mage.json and bis/mage-*.json
- the engine pin

So this round is a fresh full audit of the same text, not a review of new work.

## Builds decoded

These are unchanged from round 15 (talents/mage.json, tier/column order), and all five are legal at 51 points:
- **Arcane leveling 31/4/16:**
  - Arcane: Wand Specialization 1, Improved Channeling 4, Arcane Resilience 1, Arcane Concentration 5, Missile Barrage 2, Arcane Geometry 1, Arcane Impact 3, Arcane Blast 1, Arcane Meditation 3, Presence of Mind 1, Arcane Mind 5, Arcane Instability 3, Arcane Power 1.
  - Fire: Incineration 4.
  - Frost: Improved Frostbolt 5, Precision 3, Ice Shards 5, Frost Channeling 3.
- **Arcane raid 32/16/3:** the leveling Arcane points with Improved Channeling 3 and Arcane Meditation 2, plus Arcane Subtlety 2 and Arcane Wrath 1. Fire: Incineration 5, Accelerant 2, Ignite 5, Pyroblast 1, Improved Fire Blast 2, Heating Up 1. Frost: Precision 3.
- **Fire leveling 17/31/3:**
  - Arcane: Wand Specialization 2, Improved Channeling 5, Arcane Resilience 2, Magic Absorption 1, Arcane Concentration 5, Arcane Meditation 2.
  - Fire: Incineration 5, Accelerant 1, Ignite 5, Flame Throwing 2, Burning Soul 1, Pyroblast 1, Improved Fire Blast 2, Improved Scorch 3, Heating Up 1, Master of Elements 3, Critical Mass 3, Fire Power 3, Combustion 1.
  - Frost: Precision 3.
- **Fire raid:** as leveling, with Arcane Resilience 3 in place of Magic Absorption, and Burning Soul 2 in place of Accelerant 1.
- **Frost 15/5/31, both strings:**
  - Arcane: Wand Specialization 2, Improved Channeling 3, Arcane Concentration 5, Missile Barrage 2, Arcane Impact 3.
  - Fire: Incineration 5.
  - Frost: Frost Warding 2, Improved Frostbolt 5, Precision 3, Ice Shards 5, Frostbite 1, Permafrost 1, Frost Channeling 3, Ice Lance 1, Ice Block 1, Shatter 3, Cold Snap 1, Fingers of Frost 2, Winter's Chill 2, Ice Barrier 1.

Tier gates re-added, with the point counts the prose quotes:
- **Arcane Impact for Frost:** reached at 10 Arcane points (Wand Specialization 2, Improved Channeling 3, Arcane Concentration 5).
- **Arcane Meditation for Fire:** reached at 15 points, plus its Arcane Concentration prerequisite.
- **Presence of Mind:** out of Fire's reach at 17 Arcane points (fire.md:42).
- **Combustion:** 28 points of damage talents plus Flame Throwing, then the 2-point toll (fire.md:62).
- **Arcane raid against leveling:** +3 points (Arcane Subtlety 2 and Arcane Wrath 1), paid by -1 Arcane Meditation, -1 Improved Channeling and one more Arcane point (arcane.md:66).
- **Arcane leveling, 16 Frost points:** Improved Frostbolt 5, Precision 3, Ice Shards 5 and Frost Channeling 3 (arcane.md:42).

## Leveling orders against the prose (addon-builds.json, index 0 = level 10)

Every level the prose names matches the shipped order:
- **Arcane:**
  - Wand Specialization at 10, Improved Channeling 11 to 14, Arcane Concentration 15 to 19.
  - Arcane Blast 20, Arcane Resilience 21, Missile Barrage 22 and 23, Arcane Power 40.
  - Incineration 41 to 44, Improved Frostbolt 45 to 49, Precision 50 to 52, Ice Shards 53 to 57 ("by 57"), Frost Channeling 58 to 60.
- **Fire:**
  - Incineration 10 to 14, Accelerant 15, Ignite 16 to 19, Pyroblast 20 (so "open with Pyroblast" in Levels 20 to 60 holds).
  - Burning Soul 24, Improved Scorch 27 to 29, Heating Up 30, Combustion 40.
- **Frost:**
  - Improved Frostbolt 10 to 14, Frost Warding 15 and 16, Precision 17 to 19.
  - Ice Shards 20 to 24, Frostbite 25, Permafrost 26, Frost Channeling 27 to 29.
  - Ice Lance 30, Ice Block 31, Shatter 32 to 34, Cold Snap 35, Fingers of Frost 36 and 37, Winter's Chill 38 and 39, Ice Barrier 40.

Trainer levels are as round 15 listed them. The kit named in index.md:22 is all trainable (Remove Lesser Curse, Counterspell, Polymorph, Blink, Arcane Brilliance, portals, Teleport: Dalaran).

## Findings

None.

Counts: CONFIRMED 0, LIKELY 0, NEEDS-SIM 0 new, NEEDS-OWNER 0.

MAGE-R15-1 (Arcane "from 22" Missiles) is routed to sim round 18 and is not re-filed. The page does not state it any more falsely than it did in round 15.

## Re-checked this round (all hold)

- **Talent text (talents/mage.json), every rank the builds take:**
  - Improved Channeling 4/5: 80% for Missiles, 56% for Blast and Wrath. So "do not push back every Blast" and "most hits no longer push back your Missiles" hold.
  - Arcane Geometry 1/2: 3 yards.
  - Arcane Impact 3/3: 6% Arcane crit.
  - Arcane Meditation: 17, 33 and 50%. With Mage Armor's 50% (spells 6117, 22782 and 22783), that is "all of it at three ranks, most of it at two".
  - Burning Soul: 23% and 10% at 1 rank, 47% and 20% at 2.
  - Accelerant: 0.1 s a rank.
  - Frost Warding 2/2: 30% armor from Frost Armor and Ice Armor.
  - Fingers of Frost 2/2: 15% from Chill effects, 2 spells, 15 s.
  - Winter's Chill 2/5: 40% chance from Frostbolt and empowered Ice Lance. Each stack adds 20% Cone of Cold damage and 7% spell damage, so 5 stacks double it and add 35%.
  - Ice Barrier: 15% plus 5%, no pushback while shielded.
  - Missile Barrage: 40% from Blast, 20% from Wrath, Fireball, Frostbolt and Frostfire Bolt.
  - Heating Up: from Fireball, Frostfire Bolt, Fire Blast and Scorch crits, plus every Fire Blast hit.
  - Master of Elements: 30% on Fire and Frost crits.
  - Presence of Mind: under 10 s.
  - Arcane Mind: +100% Arcane crit bonus, which "doubles the bonus".
  - Ice Shards: +100% Frost crit bonus.
  - Improved Blizzard and Improved Cone of Cold exist on 70755 (frost.md:65 and :117).
- **Ice Block and Cold Snap tiers:** Ice Block is tier 3 (15 points) and Cold Snap tier 4, so "15 points instead of 20, with Cold Snap below it" holds (index.md:29).
- **Spell values (spelltips/mage.json):**
  - Cooldowns: Presence of Mind 180 s, Arcane Power 180 s, Combustion 180 s, Evocation 480 s, Ice Block 300 s, Cold Snap 600 s, Ice Barrier 30 s, Ice Lance 8 s, Fire Blast 8 s.
  - Costs: Arcane Explosion r6 390 mana, Flamestrike r6 990 mana with a 3 s cast and 8 s burn, Blizzard r6 1,400 mana, Scorch r7 150 mana and 1.5 s, Frostfire Bolt r3 370 mana and 3 s.
  - Damage: Frostfire Bolt r3 hits for 292 plus 57 against Frostbolt r11's 475, so "hits for less than Frostbolt".
  - Arcane Blast is 1.5 s at 15% of base mana. Mana Ruby restores 1,100. Arcane Missiles rank 8 (25345) is the Ahn'Qiraj book rank.
- **Shared cooldown claim:** in the era client tables (SpellCategories and SpellCooldowns), the mana gems (10052, 10057, 10058), Demonic Rune 16666, Dark Rune 27869 and the Robe's 18385 all sit in category 30 with a 60 s category lockout. That matches "using any of the three locks the other two out for a minute".
- **Race lines:**
  - Every racial value matches races.json: Eureka!, Expansive Mind 5% mana, Ruthless 1%, The Human Spirit 10%, Touch of the Grave (5% of max health, once per 15 s), Blood Fury 10% spell power for 15 s every 2 min, Berserking 10% for 15 s every 3 min, Wind Blessed 2% haste, Read Leyline.
  - Every order and closeness word matches round3/race-sims.md: Arcane Gnome > Human > Skyborne, Undead > Orc > Troll; Fire Human ≈ Gnome > Skyborne, Undead > Orc > Troll; Frost Skyborne > Gnome > Human, Undead ≈ Orc ≈ Troll.
  - Mage races in combos.json are Human, Orc (new), Undead, Gnome, Troll and High Order Skyborne (new), matching index.md:36-38.
- **Fire card 5, "your most expensive spell for its damage":** this holds at 60 with raid spell power.
  - Fire Blast r7 is 340 mana for 438 damage, Scorch r7 150 for 178. Both have the 1.5/3.5 coefficient, against 1.0 for Fireball and Pyroblast.
  - With about 600 spell power, that is about 2.0 damage per mana for Fire Blast, against about 2.9 for Scorch, 2.8 for Fireball and 3.2 for Pyroblast.
- **Frost stat and consumables prose:**
  - The band 60 raid row puts frost_power at 0.85 of spell_power. That fits "nearly everything ... only Missiles and the odd Fire Blast miss out".
  - Elixir of Frost Power is in the raid consumables.
- **Source links:** all seven fetched again on 2026-11-04, and all returned HTTP 200.
  - Wowhead Skyborne First Look says Alliance Skyborne can be Mages.
  - Every label follows the round 4 source-link rule: Talents Forever is labelled as beta client data, and /sim/weights and /sim/drops are on damage specs only.
- **Voice:** no em or en dashes, no site, sim, APL, band or pipeline words in body prose, no quoted nightly numbers.

## SHOULDs (wording, not findings)

These are carried from round 15, are still wording, and are not re-filed:
- **arcane.md:72:** "it can roll Missile Barrage before the fight starts". The Frostbolt lands as the fight starts.
- **frost.md:71:** "Recast Ice Barrier the moment it breaks" against its 30 s cooldown.
- **fire.md:94:** "Greater Arcane Elixir is your elixir" beside a table that also lists Elixir of Fire Power.
- **arcane.md:102:** the professions reasons. The real ones are the Robe and Brilliant Wand.
- **fire.md:116:** Burning Soul's 23% "keeps a mob's hits from pushing back every cast".
- **index.md:32:** Accelerant is missing from the list of new talents.
- **arcane.md:74:** add "from 22" to the "On the way to 60" sentence.
- **frost.md:63:** Winter's Chill also procs from empowered Ice Lance.

## Decisions

- **D1.** No new text since round 15. A second full pass found no statement that is false on 70755 or that contradicts the page's own tables, cards, race data or point order.
- **D2.** Fire card 5's "most expensive spell for its damage" is judged at 60 with raid spell power, the level the page means. At 0 spell power Scorch costs more per point of damage, but that is not the context of a raid card. Not a finding.
- **D3.** Boughs in the Wind (Windbough Wand) stays out of the level 10 quest sentence, per round 4 D2. loot.json still does not tag it for the mage class.

## Sims requested

None new. Carried, and already routed:
- **Sim round 18:** MAGE-R15-1, Arcane at level 22 and 23 with Missile Barrage below 2 ranks, the line 8 extension against the shipped APL.
- **Data lane:** the Frost band 60 BiS rows re-ranked on the shipped 15/5/31 string, and the round 14/15 BiS items (three-profession Fire and Frost sets, Frozen Heart for Arcane and Fire, the Arcane raid rows on 32/3/16).
- **Sim lane:** race re-runs on the shipped builds, and the round 15 carried list (S1, S2, SIM-R2-1, SIM-R2-3, SIM-5, SIM-6, SIM-2, the Arcane raid pre-pull, the Presence of Mind target, Arcane Wrath at 3 against 4 stacks, gem timing).

## Sim findings

None new. SIM-R7-1 is carried: Arcane Subtlety's 15 resistance cut is not modeled.

## Sources

- data/builds/1.60.1.70755:
  - talents/mage.json: every rank text, tier and prerequisite.
  - spelltips/mage.json: 400574, 1239700, 1326082, 12043, 12042, 11129, 12051, 11958, 12472, 13033, 1240047, 10199, 10207, 10216, 10187, 10202, 1237313, 25304, 25306, 18809, 10054, 25345 and Mage Armor 6117, 22782 and 22783.
  - trainables/mage.json, races.json, combos.json and loot.json (quest 93797).
  - addon-builds.json: mage-arcane, mage-fire and mage-frost, leveling and raid.
  - bis/mage-frost.json: band 60 raid row and presets.
- Era client tables in the main checkout (/Users/jh/code/forever/SpellCategories_era.csv and SpellCooldowns_era.csv): spells 10052, 10057, 10058, 16666, 27869, 18385, 12043, 12042, 12051 and 13033.
- data/curated/apl/mage-{arcane,fire,frost}.json.
- design/reviews/guide-accuracy/round1 to round15/mage.md, round4/mage-levels.md, round3/race-sims.md and ledger.md.
- Source links fetched 2026-11-04, all HTTP 200:
  - Blizzard, Deep Dive panel recap: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
  - Wowhead, Skyborne First Look: https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829
  - Power Up Gaming, Forever talent changes vs Classic: https://powerupgaming.co.uk/2026/09/14/new-wow-forever-talents-changes-vs-classic-guide/
  - Talents Forever data export: https://talentsforever.com/data.json
  - https://foreversixty.gg/sim/weights, https://foreversixty.gg/sim/drops and https://foreversixty.gg/planner?class=mage
