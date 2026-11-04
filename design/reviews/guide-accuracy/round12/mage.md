# Guide accuracy round 12: Mage

Audited on the round 12 audit tree, main a8b147a6f: eleven guide rounds, sim round 15, leveling-order and guide-header-build. Client 1.60.1.70755, engine pin 9f7f1f6e4 (sim/enginever/version.go).

Files read line by line:
- web/src/content/guides/mage/index.md, arcane.md, fire.md and frost.md.
- Every note and action in data/curated/apl/mage-{arcane,fire,frost}.json, priorityList and prepullActions.

Read first:
- Mage reports for rounds 1 to 11, with their fix logs.
- round4/mage-levels.md, round2 to round5 race-sims.md and ledger.md.
- Commit 3bd24d3a2 (guide-header-build) and its "guide prose to fix" note.

## What changed since round 11

Since d54de48aa, main has these commits:
- The round 11 fixes (91da94f1e): arcane.md:42, :64, :74, :84 and :98, and fire.md:66.
- guide-header-build (3bd24d3a2):
  - The header rail card now shows and loads the raid build where one exists. On Arcane and Fire that is the raid string.
  - Key-talent chips show both builds' ranks where they differ.
  - "For the raid build." now sits under Rotation and priority.
- A planner budget fix (3dff67836) and merges.

No commit touches mage APL files, addon-builds.json, talents or bis/mage-*.json.

The commit note asks the guide lane to fix two lines: Arcane's "2.5 second" Frostbolt in the Presence of Mind card note, and the Improved Frostbolt key-talent line. Neither is wrong on the page now:
- No Arcane card note says 2.5 seconds. Note [3] and note [5] name Frostbolt only "before you have the Fire points".
- The key-talent line arcane.md:59 is marked "(10 points, leveling build)", so its 2.5 s Frostbolt is right. The new chips show Level 60 5/5 and Raid 0/5 beside it.

## Builds decoded (talents/mage.json, tier/column order; all five legal)

The decoded builds are unchanged from round 11, so see round11/mage.md for the full lists:
- Arcane leveling 31/4/16.
- Arcane raid 32/16/3.
- Fire leveling 17/31/3.
- Fire raid 17/31/3.
- Frost 15/5/31, the same string for both.

Shipped orders (addon-builds.json, index 0 = level 10), re-decoded this round:
- **Arcane leveling**:
  - Wand Specialization 10, Improved Channeling 11-14, Arcane Concentration 15-19.
  - Arcane Blast 20, Arcane Resilience 21, Missile Barrage 22-23, Arcane Geometry 24, Arcane Impact 25-27.
  - Arcane Meditation 28-30, Presence of Mind 31, Arcane Mind 32-36, Arcane Instability 37-39, Arcane Power 40.
  - Incineration 41-44, Improved Frostbolt 45-49, Precision 50-52, Ice Shards 53-57, **Frost Channeling 58-60**.
- **Fire leveling**:
  - Incineration 10-14, Accelerant 15, Ignite 16-19, Pyroblast 20, Ignite 21.
  - Flame Throwing 22-23, Burning Soul 24, Improved Fire Blast 25-26, Improved Scorch 27-29, Heating Up 30.
  - Master of Elements 31-33, Critical Mass 34-36, Fire Power 37-39, Combustion 40.
  - Then Arcane 41-57 and Precision 58-60.
- **Frost** (both builds):
  - Improved Frostbolt 10-14, Frost Warding 15-16, Precision 17-19, Ice Shards 20-24.
  - Frostbite 25, Permafrost 26, Frost Channeling 27-29, Ice Lance 30, Ice Block 31, Shatter 32-34.
  - Cold Snap 35, Fingers of Frost 36-37, Winter's Chill 38-39, Ice Barrier 40.
  - Then Arcane 41-55 and Incineration 56-60.

The Levels 1 to 19 and Levels 20 to 60 paragraphs match these orders level for level.

## Findings

| id | status | kind | file:line | wrong claim | evidence | fix |
|---|---|---|---|---|---|---|
| MAGE-R12-1 | CONFIRMED (low) | rotation, round 11 edit | arcane.md:74 | "On the way to 60 the loop is the same with Frostbolt as the filler and Fire Blast only to finish a mob or on the move." | The sentence before it, the raid loop, includes "Pyroblast on 3 Heating Up stacks" and "Arcane Wrath every time it comes off its 6 second cooldown while you hold 3 stacks". The leveling build (arcane.md:6, 31/4/16) has neither: no Arcane Wrath, no Pyroblast and no Heating Up. Its only Fire points are 4 in Incineration, and its shipped order never takes them. "The same" carries both lines into leveling. Round 11's edit replaced "no Fire spells" with "Fire Blast only...". That removed the one phrase that excluded Pyroblast, so the leveling loop now reads as if it casts Heating Up Pyroblasts. The page's own key-talent lines (arcane.md:60, :61) mark both talents "raid build". | "On the way to 60 there is no Arcane Wrath and no Pyroblast: Blast to 3 stacks, Missiles the moment Missile Barrage lights up, Frostbolt as the filler, and Fire Blast only to finish a mob or on the move." |
| MAGE-R12-2 | CONFIRMED (low) | leveling, talent | arcane.md:37, :82 | :37 "On the way to 60, the Frostbolt filler and Frost Channeling keep you slowing adds and light on threat". :82 "Frost Channeling cuts 30% of the threat from your Frostbolts on the way to 60". | The shipped Arcane leveling order (addon-builds.json mage-arcane leveling, indices 48-50) takes Frost Channeling last, at 58, 59 and 60. The page says so at arcane.md:116: "From 41 the points go into Incineration, then Improved Frostbolt, Precision, Ice Shards and Frost Channeling". From 20 to 57 an Arcane mage has no Frost Channeling, so it cannot keep them light on threat "on the way to 60". It is a talent of the finished leveling build at 60. Round 10 wrote both lines (MAGE-R10-2) and round 11 did not catch them. | :37 "On the way to 60, the Frostbolt slow keeps adds off you while you blast". :82 "The leveling build's Frost Channeling cuts 30% of the threat from your Frostbolts, and the raid build's Arcane Subtlety does the same for your Arcane spells." (Keep the rest of :82.) |

Counts: CONFIRMED 2 (both low), LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0.

The three worst:
1. MAGE-R12-1. The Arcane leveling loop is called "the same" as the raid loop, so it reads as if it casts Arcane Wrath and Heating Up Pyroblasts. The leveling build has neither. Round 11's edit dropped the "no Fire spells" that had excluded Pyroblast.
2. MAGE-R12-2. Arcane tells levelers that Frost Channeling keeps them light on threat on the way to 60. The shipped order takes it at 58 to 60.
3. No other finding is confirmed.

## Round 11's own fixes, re-checked

- **arcane.md:42.** "16 Frost points on Precision and the Frostbolt filler" holds: 5 + 3 + 5 + 3 = 16.
- **arcane.md:64 and fire.md:66.** "The leveling build" holds. The 31 Arcane points and the two Fire toll points check out.
- **arcane.md:74.** The edit is wrong; see MAGE-R12-1.
- **arcane.md:84.** Pyroblast r8 is a 6.0 s cast (18809), and 1.5 s at 3 Heating Up stacks. Adding it to the stand-still list is right.
- **arcane.md:98.**
  - Arcane Meditation rank 2 is 33% (18462) and Mage Armor is 50% (22783), so 83% of Spirit regeneration keeps running while you cast.
  - "Most of that Spirit's regeneration" is right.

## Header card (guide-header-build) against the prose

- **Arcane.** The rail now shows the gnome raid build, 32/16/3.
  - The prose calls it "the raid build" wherever it means it.
  - The opener's 3 s Frostbolt is right for the raid build: no Improved Frostbolt, and Accelerant does not touch Frostbolt (talents/mage.json Accelerant).
  - The key-talent counts agree with the two-rank chips:
    - "Arcane Meditation (3 points, 2 in the raid build)".
    - "Improved Frostbolt and Ice Shards (10 points, leveling build)".
    - "Arcane Wrath (1 point, raid build)".
    - "Ignite ... (9 points, raid build)".
    - Every other key talent has the same rank in both builds.
- **Fire.** The rail now shows the human raid build.
  - Every key talent has the same rank in both builds: Improved Scorch 3, Heating Up 1, Pyroblast 1, Ignite 5, Fire Power 3, Incineration plus Critical Mass 8, Master of Elements 3, Improved Fire Blast 2 and Combustion 1.
  - The builds differ only in Arcane Resilience, Magic Absorption, Accelerant and Burning Soul. fire.md:66 describes both builds correctly.
- **Frost.** One build, nothing to contradict.

## No findings in these areas

- **Spell values** (spelltips/mage.json):
  - Arcane Blast: 1.5 s cast and 15% base mana, so a 4-stack Blast costs five times a fresh one.
  - Arcane Wrath r4: 3.5 s cast, 6 s cooldown, 395 mana.
  - Pyroblast: 6 s cast, 440 mana. Fire Blast: 340 mana, 8 s cooldown, 20 yd.
  - Scorch r7: 150 mana, 1.5 s. Fireball r12: 3.5 s.
  - Frostbolt: 3.0 s base. Frostfire Bolt r3: 3.0 s, 370 mana, 292 damage.
  - Ice Lance r6: instant, 8 s cooldown, +300% against Frozen targets.
  - Ice Barrier: 30 s cooldown, 15% for 1 min plus 5%, no pushback. Ice Block: 10 s, 5 min. Cold Snap: 10 min, resets Frost spells including Ice Barrier.
  - Evocation: 8 min, +1500% for 8 s. Arcane Power, Presence of Mind and Combustion: 3 min.
  - Arcane Explosion: 390 mana, 10 yd. Flamestrike: 990 mana, 3 s cast, 8 s burn. Blizzard: 1,400 mana.
  - Mana Ruby restores 1,100 and Mana Citrine 850. Mage Armor: 50%.
- **Trainer levels** (trainables/mage.json):
  - Fireball 1/6/12/18. Frostbolt 4/8/14/20. Conjure Water 4. Fire Blast 6. Arcane Missiles 8. Frost Nova 10. Arcane Explosion 14. Flamestrike 16.
  - Remove Lesser Curse 18. Blink, Blizzard and Evocation 20. Scorch 22. Counterspell 24. Cone of Cold 26. Ice Armor 30. Mage Armor 34.
  - Frostfire Bolt 40/50/60. Teleport: Dalaran 50. Arcane Brilliance 56. Mana Ruby 58.
- **Talent text**, all rank texts as quoted on the pages:
  - Arcane: Improved Channeling, Arcane Subtlety (15 resistance and 30% Arcane threat, own spells only), Arcane Concentration, Missile Barrage, Arcane Geometry, Arcane Impact, Arcane Meditation, Presence of Mind, Arcane Mind, Arcane Instability.
  - Fire: Incineration (Fire Blast, Ice Lance, Arcane Blast and Scorch), Accelerant, Ignite, Flame Throwing, Burning Soul, Improved Fire Blast (2 s and 15 yd), Improved Scorch, Heating Up, Master of Elements, Critical Mass, Fire Power, Combustion.
  - Frost: Frost Warding, Improved Frostbolt, Precision, Ice Shards, Frost Channeling, Improved Blizzard, Ice Lance, Shatter, Fingers of Frost, Winter's Chill and Ice Barrier.
- **index.md**: Ice Block at 15 points (row 4) with Cold Snap below. 1.12's Elemental Precision covered only Fire and Frost. Orc is the third Horde mage race. Skyborne mages are Alliance only. The new-talent list.
- **Racials**:
  - races.json and spelltips/shared.json: Ruthless 1%, The Human Spirit 10%, Expansive Mind 5% mana, Blood Fury 10% spell power for 15 s every 2 min, Berserking 10% for 15 s every 3 min, Read Leyline 50% for 15 s every 2 min, Wind Blessed 2%.
  - Eureka! and Touch of the Grave use the mage rows, as settled in round 8.
  - Race order against round3/race-sims.md. Nothing measured since contradicts it.
- **Card notes**: Arcane [1]-[13], Fire [1]-[8] and Frost [1]-[7]. Fire [5] "most expensive spell for its damage" holds at 60 with spell power: Fire Blast and Scorch carry the same coefficient, and Fire Blast costs 340 against Scorch's 150.
- **Consumables**: Elixir of Frost Power (17708) is in simconsumes.json. Elixir of Greater Firepower is not in Forever (round 1 O1).
- **Voice**: no em or en dashes, no site, sim or pipeline words in body prose, and no quoted nightly numbers.
- **Source links**: all seven fetched live this round with HTTP 200 and match their labels:
  - Blizzard Deep Dive recap
  - Talents Forever data.json
  - /sim/weights
  - /sim/drops
  - Wowhead Skyborne First Look
  - Power Up Gaming talent changes
  - /planner?class=mage

## SHOULDs (wording, not findings)

- arcane.md:60 calls Arcane Wrath "a 3.5 second Arcane bolt" on a line marked "raid build". The raid build always has Accelerant 2, so it is a 3.3 s cast there. Round 11 accepted the base value because :66 adds Accelerant. Suggested: "a 3.3 second Arcane bolt with Accelerant".
- arcane.md:59 "a 2.5 second Frostbolt with full crit damage is the filler". The leveling build reaches 2.5 s at 49 and full Ice Shards at 57. Suggested: "...is the filler once both are in, by 57".
- Carried from round 10: Accelerant is missing from the index.md:32 new-talent list.

## Decisions

- **D1.** "On the way to 60" means the levels a player is leveling through, not the finished level 60 leveling build. So a talent the shipped order takes at 58 to 60 cannot be credited "on the way to 60" (MAGE-R12-2). Talents labelled "leveling build" describe the finished build and are not filed.
- **D2.** The guide-header-build note's two Arcane items, the card note and the Improved Frostbolt line, are not findings. Neither states a 2.5 s Frostbolt for the raid build.
- **D3.** A pre-pull Pyroblast against a pre-pull Frostbolt on the Arcane raid build is a measurable opener question. The prose follows the APL until it is measured, so it goes under Sims requested and is not a finding.

## Sims requested

- **PROVISIONAL (BiS re-rank).** bis/mage-*.json are still the d1d9f2775 runs (2026-11-04T03:13Z to 03:17Z), on the pre-round-15 strings.
  - Re-rank Arcane on 32/16/3. Fire power now covers half of Arcane's casts, but the Arcane statPriority frontmatter has no Fire power.
  - Re-rank Frost on 15/5/31.
  - Re-rank Fire on the current strings: the BiS file still carries Impact 1.
- **Race re-runs at 60 on the new builds.**
  - Arcane raid 32/16/3: Gnome, Human, Undead, Orc and Troll.
  - Frost 15/5/31: Gnome, Human, Undead, Troll and Orc.
- **Arcane raid 60, the pre-pull cast.** Frostbolt r10, the current pre-pull (3.0 s, can proc Missile Barrage), against Pyroblast r8 (6.0 s, the opener Fire uses).
  - Seeds 1 and 50001.
  - Adopt Pyroblast only on a gain of 1% or more beyond error. If adopted, the opener prose at arcane.md:72 changes with it.
- **Carried from round 11:**
  - The Presence of Mind target: Pyroblast against Arcane Wrath and against Arcane Blast.
  - Arcane Wrath at 3 stacks against 4.
  - S1, S2, SIM-R2-1, SIM-R2-3, SIM-5, SIM-6 and SIM-2.

## Sim findings

- None new. SIM-R7-1 is carried: the engine does not model Arcane Subtlety's 15 spell penetration.

## Sources

- data/builds/1.60.1.70755:
  - talents/mage.json: every rank text above.
  - trainables/mage.json.
  - spelltips/mage.json: 400574 to 1239700, 1326082, 18809, 10199, 10207, 10151, 25306, 10181, 25304, 1240047, 13033, 11958, 12472, 12051, 12042, 12043, 11129, 10202, 10216, 10187, 10054, 10053, 22783, 1237313 and 1953.
  - spelltips/shared.json: the racials.
  - races.json.
  - addon-builds.json: mage-arcane, mage-fire and mage-frost, leveling and raid orders.
  - simconsumes.json.
  - bis/mage-{arcane,fire,frost}.json: generated_at, engine_version and band 60 talents.
- data/curated/apl/mage-arcane.json, mage-fire.json and mage-frost.json.
- design/reviews/guide-accuracy/round3/race-sims.md: mage rows.
- Commits:
  - 91da94f1e: round 11 fixes.
  - 3bd24d3a2: guide-header-build.
  - eabba8f84 and e01dc2dff: the round 10 text and the Arcane raid build.
- Blizzard, Deep Dive panel recap: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
- Wowhead, Skyborne First Look: https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829
- Power Up Gaming, Forever talent changes vs Classic: https://powerupgaming.co.uk/2026/09/14/new-wow-forever-talents-changes-vs-classic-guide/
- Talents Forever data export: https://talentsforever.com/data.json

## Fix log (round 12)

- MAGE-R12-1: fixed, arcane.md:74: "On the way to 60 there is no Arcane Wrath and no Pyroblast: Blast to 3 stacks, Frostbolt as the filler, and Fire Blast only to finish a mob or on the move." (Missile Barrage is covered by the next sentences.)
- MAGE-R12-2: fixed, arcane.md:37 ("the Frostbolt slow keeps adds off you while you blast") and :82 ("The leveling build's Frost Channeling cuts 30%..."). Verified addon-builds.json mage-arcane leveling: Frost Channeling at indices 48 to 50 (levels 58 to 60).
- SHOULD arcane.md:60: applied, "a 3.3 second Arcane bolt with Accelerant". Verified the raid order takes Accelerant 2 (0.2 s, talents/mage.json) and Arcane Wrath 3.5 s base.
- SHOULD arcane.md:59: applied, "Frostbolt is half your casts while you level, and once both are in, by 57, a 2.5 second Frostbolt with full crit damage is the filler". Verified Improved Frostbolt 45 to 49 and Ice Shards 53 to 57 in the leveling order.
- Carried index.md:32 Accelerant SHOULD: left, not in this lane's scope.
