# Guide accuracy round 14: Mage

Audited on the round 14 audit tree, main d7fc51fe7 (thirteen guide rounds, launch state). Client 1.60.1.70755, engine pin f637ec32c (sim/enginever/version.go).

Files read line by line:
- web/src/content/guides/mage/index.md, arcane.md, fire.md and frost.md.
- Every note and action in data/curated/apl/mage-{arcane,fire,frost}.json, priorityList and prepullActions.
- The live BiS files bis/mage-{arcane,fire,frost}.json (de044125c picks, engine d1d9f2775): band 60 raid rows for both factions, scale factors and preset consumables.

Read first: the mage reports for rounds 1 to 13 with their fix logs, round4/mage-levels.md, the race-sims files for rounds 2 to 5, and ledger.md.

## What changed since round 13

Main moved from 604b78a80 to d7fc51fe7. The mage changes are:
- **arcane.md:48.** The Brings line now reads "Arcane Subtlety's resistance cut helps only your own spells, not the raid's."
- **arcane.md:72.** The opener has a new raid pre-pull: "In a raid, start Pyroblast 6 seconds before the pull ... and its crit is your first Ignite. With the leveling build, start a Frostbolt 2.5 seconds out instead".
- **mage-arcane.json.** The prepull now casts Pyroblast r8 at -6 s when Heating Up (400625) is known, and Frostbolt r10 at -2.5 s otherwise. The Presence of Mind Pyroblast line now also checks that Presence of Mind (12043) is known.

addon-builds.json, talents, trainables and spelltips have no mage changes. The mage leveling orders are the ones round 13 decoded, and the Levels 1 to 19 and Levels 20 to 60 prose still agree with them level for level.

## Builds decoded

These are the same as round 13 (talents/mage.json, tier/column order). All five builds are legal at 51 points.
- **Arcane leveling, 31/4/16.**
- **Arcane raid, 32/16/3.**
- **Fire leveling, 17/31/3.**
- **Fire raid, 17/31/3.**
- **Frost, 15/5/31.**

Every point count the prose quotes was re-added this round:
- Arcane raid: 3 points into Arcane Subtlety and Arcane Wrath, paid for by a rank each of Arcane Meditation and Improved Channeling plus a 32nd point. The 16 Fire points are 5/2/5/1/2/1.
- Fire: Combustion's 28 + 2.
- Frost: 31 / 15 / 5.

## Findings

| id | status | kind | file:line | claim | evidence | fix |
|---|---|---|---|---|---|---|
| MAGE-R14-1 | CONFIRMED | gear/professions | fire.md:102, frost.md:103 | "Tailoring and Enchanting are the standard pair ... Gnomes who want the explosives take Engineering." | The live BiS head for Fire and Frost, on both factions (bis/mage-fire.json and mage-frost.json, band 60 raid, Alliance and Horde), is Spellpower Goggles Xtreme Plus (15999). It needs Engineering 270 to equip: raw/wowhead/items.json at 70755 gives reqskill 202 and reqskillrank 270, and the Wowhead Forever tooltip says "Requires Engineering (270)". The prose presents Engineering as a Gnome-only extra for explosives. On a Human Fire mage, the guide's own pick, the page's best-in-slot head cannot be worn without it. The same rows also hold two bind on pickup crafts: Robe of the Archmage (14152, Tailoring, bonding pickup) and Frozen Heart of the Mountain (249469, Enchanting, bonding pickup). | fire.md:102: "Tailoring and Enchanting are the standard pair. Robe of the Archmage binds when it is made, so only a Tailor wears it, and Enchanting puts the spell power enchants on your own gear. The goggles in the head slot need 270 Engineering to wear: without it, take the best head you can wear and keep your two professions. Alchemy pays for itself once you raid every week, because Fire drinks a Major Mana Potion on every boss." frost.md:103: the same first three sentences, then "Frost drinks fewer potions than the other specs, so it can skip Alchemy." |
| MAGE-R14-2 | CONFIRMED (low) | rotation | arcane.md:72 | "In a raid, start Pyroblast 6 seconds before the pull ... the whole cast happens while nothing else is going on, and its crit is your first Ignite." | Pyroblast has no guaranteed crit. The client gives it only the normal spell crit roll: the Pyroblast tooltip has no crit clause, and no Arcane raid talent (Incineration covers Fire Blast, Ice Lance, Arcane Blast and Scorch, not Pyroblast) or APL condition forces one. Ignite (talents/mage.json) triggers only on "critical strikes from Fire damage spells". Most pre-pull Pyroblasts start no Ignite, and the sentence says every one does. | "In a raid, start Pyroblast 6 seconds before the pull so it lands just after the tank's first hit: the whole cast happens while nothing else is going on, and if it crits, that is your first Ignite." |

Counts: CONFIRMED 2, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0.

## Round 13's state and this round's new text, re-checked

- **arcane.md:48 Brings.** Arcane Subtlety r2 reads "Reduces your target's resistance to all your spells by 15 ...". "Helps only your own spells, not the raid's" is right.
- **arcane.md:72 opener.** It matches the new prepull:
  - Pyroblast r8 at -6 s is gated on Heating Up being known, which only the raid build has. Frostbolt r10 at -2.5 s goes out otherwise, and the leveling build's Improved Frostbolt is 5/5 from 49.
  - Pyroblast r8 is a 6.0 s cast. Accelerant does not touch it.
  - Arcane Power and Presence of Mind are gated at 3 s or later in the fight, which fits "Blast twice while the tank puts a couple of Sunders up".
  - The only error in the paragraph is "its crit" (R14-2).
- **Presence of Mind lines** (arcane.md:56, :72, :84). APL lines 4 and 5 now check spellIsKnown 12043 and the aura. "An instant Frostbolt before you have the Fire points" stays a SHOULD (round 13 D3).

## Race lines against the current race data

- **Racial values** (races.json, spelltips/shared.json, and the mage rows settled in rounds 8 and 9):
  - Eureka!: 10% more damage and 10% lower cost on damage spells, 8 s, 2 min. The shared row is the rogue's Energy row, and the mage row 1259817 says Mana.
  - Touch of the Grave: 30% on the mage row. The shared tip shows 20%, but that is the rogue, warrior, hunter, paladin row. Engine sim/core/racial_touch_of_the_grave.go:42 and racials_test.go:670 give 1260201, 0.30 for mages. 5% of maximum health, 15 s.
  - Will of the Forsaken: removes Charm, Fear and Sleep.
  - Blood Fury: 10% spell power, 15 s, 2 min.
  - Berserking: 10% cast speed, 15 s, 3 min.
  - Ruthless: 1% crit.
  - The Human Spirit: 10% Spirit.
  - Expansive Mind: 5% mana.
  - Wind Blessed: 2% haste.
  - Read Leyline: 50% health and mana regeneration, 15 s, 2 min.
  - Every value on all four pages holds.
- **recommendedRaces** (Arcane [gnome, undead], Fire [human, undead], Frost [gnome, undead]) matches each Races paragraph and index.md:38.
- **Orders against round3/race-sims.md**, the only measured mage race run; the new builds are still unmeasured:
  - Arcane: Gnome 614.0, Human 611.8, Skyborne 607.5; Undead 623.3, Orc 619.3, Troll 616.0.
  - Fire: Human 667.2 against Gnome 666.8 ("a hair ahead"). Skyborne is lowest at 655.9. Troll 667.3 is behind both Undead and Orc.
  - Frost: Skyborne 630.9 against Gnome 627.3 ("edges a Gnome by a hair"); Undead 633.4, Orc 632.4, Troll 632.2.
  - No line contradicts a measured result or the client.
- **Skyborne** is named only as the paid option, never as the pick. Mage Skyborne are High Order only (combos.json, index.md:37).

## Gear, consumables and stat prose against the live tables

- **Robe of the Archmage** (14152), named in every Mana paragraph, is the chest in all six band 60 raid rows.
- **Arcane professions** (arcane.md:102): the Arcane rows need a Tailor for the bind on pickup Robe and an Enchanter for the bind on pickup Frozen Heart of the Mountain and Brilliant Wand (249385). Neither Arcane head needs a profession (Dreadmist Mask for Alliance, Red Mageweave Headband for Horde). So "Tailoring and Enchanting" is right. The stated reasons are thin; see the SHOULDs.
- **Stat prose** explains and does not rank. Each stated reason holds against the band 60 raid scale factors:
  - Fire: spell power and Fire power are equal at 2.689, so "the same stat to you" is right.
  - Frost: Frost power (4.37) is below spell power (5.12), which fits "only Missiles and the odd Fire Blast miss out".
  - Arcane: Arcane power (1.04) is below spell power (1.52), which fits "not your Fire spells in a raid".
- **Consumables:** Fire's "Greater Arcane Elixir is your elixir" sits above a table that also lists Elixir of Fire Power (SHOULD below). Frost's Elixir of Frost Power is in its table. Major Mana Potion and Demonic Rune are in all three.

## No findings in these areas (re-derived this round)

- **Talent text** against talents/mage.json (every value the pages quote):
  - Arcane Meditation 17/33/50.
  - Improved Channeling 4/5 gives Missiles 80%.
  - Burning Soul is 23%/10% at one rank and 47%/20% at two.
  - Winter's Chill 2/5 is 40%, and 5 stacks double Cone of Cold.
  - Missile Barrage is 40% from Blast and 20% from the others.
  - Heating Up procs from crits on Fireball, Frostfire Bolt, Fire Blast and Scorch, plus every Fire Blast hit.
  - Shatter 50%, Ice Shards +100% crit bonus, Arcane Mind +100% Arcane crit bonus.
  - Fire Power 3% per effect to 12%, with mixed schools excluded.
  - Master of Elements refunds 30% of base cost.
  - Ice Barrier 15% plus 5%, with no pushback.
  - Fingers of Frost: 15% for 2 spells over 15 s.
- **Spell values and trainer levels:** unchanged since round 13 and still matching (spelltips and trainables have no diff).
- **Leveling 1 to 19 and 20 to 60:** match the unchanged addon-builds.json orders, as decoded in round 13.
- **Voice:**
  - No em or en dashes.
  - No site, sim, APL or band words in body prose.
  - No quoted nightly numbers.
- **Source links:** all seven fetched live this round. All returned HTTP 200, and each matches its label:
  - the Deep Dive recap
  - talentsforever.com/data.json
  - /sim/weights and /sim/drops
  - Wowhead Skyborne First Look
  - Power Up Gaming
  - /planner?class=mage

## SHOULDs (wording, not findings)

- **arcane.md:102.** "Tailoring makes the early cloth sets and Enchanting puts the spell power enchants on your own gear." The real reason for both professions is the bind on pickup gear: the Robe for Tailoring, and the Frozen Heart of the Mountain and Brilliant Wand for Enchanting. Suggested: "Tailoring and Enchanting are the pair: Robe of the Archmage binds when a Tailor makes it, and an Enchanter makes their own trinket and wand."
- **fire.md:94** "Greater Arcane Elixir is your elixir." The consumables table under it also lists Elixir of Fire Power, and the two stack. Suggested: "Greater Arcane Elixir is your main elixir, and the small Elixir of Fire Power stacks with it."
- Carried from round 13:
  - arcane.md:72 Missile Barrage "before the fight starts" (now on the leveling-build Frostbolt only).
  - frost.md:71 "Recast Ice Barrier the moment it breaks", which has a 30 s cooldown.
  - fire.md:116 "keeps a mob's hits from pushing back every cast", which is only 23% at one rank.
  - arcane.md:56/:72/:84 "before you have the Fire points".
  - index.md:32 is missing Accelerant from the new-talent list.

## Decisions

- **D1.** The professions line is a finding on Fire and Frost only. Their rows put an Engineering-locked head in the BiS slot on both factions. Arcane's rows do not, and Arcane's Tailoring and Enchanting pair is what its bind on pickup gear needs.
- **D2.** Touch of the Grave stays at 30% on mage pages. The shared tooltip's 20% is a different class row, and the engine pins the mage row at 0.30 (round 3, re-checked at f637ec32c).
- **D3.** Race lines are judged against round 3, the only measured mage race run, and against the client. Race re-runs on the shipped builds stay with the sim lane.
- **D4.** R14-2 is a finding, not a SHOULD: the sentence says every pre-pull Pyroblast crits.

## Sims requested

- **BiS (post-launch re-rank; data lane):**
  - The Fire and Frost band 60 raid sets need three professions at once: Engineering 270 for Spellpower Goggles Xtreme Plus, Tailoring for the bind on pickup Robe of the Archmage, and Enchanting for the bind on pickup Frozen Heart of the Mountain. No character can wear the full set. Rank with a two-profession limit, or show the best non-profession alternative beside each locked piece.
  - Frozen Heart of the Mountain (249469) is a BiS trinket for Arcane and Fire. Its on-use helps only Frost and Shadow damage, and its passive is hit and resistances.
  - Thunderbrew's Boot Flask (744; 8 Spirit, an on-use Fire cone) is the Arcane Alliance trinket2.
  - Check both of these after the re-rank.
  - The Arcane raid band still runs 31/3/16, not the shipped 32/16/3. This is known; re-rank on the shipped strings. Then decide whether Arcane's statPriority should select Fire power.
- **Race re-runs at 60 on the shipped builds:**
  - Arcane 32/16/3: Gnome, Human, Undead, Orc, Troll.
  - Frost 15/5/31: Gnome, Human, Skyborne, Undead, Troll, Orc.
- **Carried from earlier rounds:**
  - The Arcane raid pre-pull: Pyroblast r8 at -6 s against Frostbolt r10 at -2.5 s, now adopted in the APL. Confirm the gain on seeds 1 and 50001.
  - The Presence of Mind target: Pyroblast against Arcane Wrath and Arcane Blast.
  - Arcane Wrath at 3 stacks against 4.
  - S1, S2, SIM-R2-1, SIM-R2-3, SIM-5, SIM-6 and SIM-2.

## Sim findings

- None new. SIM-R7-1 is carried: the engine does not model Arcane Subtlety's 15 spell penetration.

## Sources

- data/builds/1.60.1.70755:
  - talents/mage.json: every rank text.
  - spelltips/shared.json: the racials.
  - races.json.
  - combos.json.
  - items/mage.json: bonding for 15999, 14152, 249469, 249385, 254013, 744 and 19120.
  - raw/wowhead/items.json: 15999, reqskill 202 at rank 270, in the live set.
  - raw/items/item-sources.json: 249469, 249385 and 254013 crafted by Enchanting, Enchanting and Tailoring.
  - addon-builds.json, unchanged for mage.
  - bis/mage-{arcane,fire,frost}.json: band 60 raid rows for both factions, weights and preset consumables.
- data/curated/apl/mage-arcane.json (prepullActions; priorityList 2-12), mage-fire.json and mage-frost.json.
- Engine f637ec32c: sim/core/racial_touch_of_the_grave.go:9-42 and sim/core/racials_test.go:652-671.
- design/reviews/guide-accuracy/round3/race-sims.md, and round1 to round13/mage.md.
- Wowhead Forever, Spellpower Goggles Xtreme Plus tooltip: https://nether.wowhead.com/forever/tooltip/item/15999
- Blizzard, Deep Dive panel recap: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
- Wowhead, Skyborne First Look: https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829
- Power Up Gaming, Forever talent changes vs Classic: https://powerupgaming.co.uk/2026/09/14/new-wow-forever-talents-changes-vs-classic-guide/
- Talents Forever data export: https://talentsforever.com/data.json

## Fix log (round 14)

- MAGE-R14-1: fixed, fire.md:102 and frost.md:103. Both name Robe of the Archmage (Tailoring) and Frozen Heart of the Mountain (Enchanting) as BoP raid-set crafts, and the goggles' Engineering 270 (raw/wowhead reqskill 202/270). The Gnome-only Engineering sentence is dropped. Verified on band 60 raid, both factions; Fire bare rows carry no Heart, so the text says "your raid set".
- MAGE-R14-2: fixed, arcane.md:72: "and if it crits, that is your first Ignite." Coordinator constraint: the "before you have the Fire points" lines were not touched. Note that the same paragraph line holds one of them.
- SHOULDs: not applied.

All files were already at updated: 2026-11-04. The prose and sections tests pass (145/145).
