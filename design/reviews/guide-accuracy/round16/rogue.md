# Guide accuracy round 16: Rogue

Audited on the round 16 audit tree (land/guides-r15 2c3135a2e: fifteen guide rounds on launch-state main), client 1.60.1.70755, engine pin b2c4a2d05 (`sim/enginever/version.go`).

**Scope.**
- Pages: `web/src/content/guides/rogue/{index,assassination,combat,subtlety}.md`, every line.
- Card notes: every rogue line in `data/builds/1.60.1.70755/addon-data.json` (`rotations.rogue-*`, all 22 or 23 bands per spec), and the notes in `data/curated/apl/rogue-{assassination,combat,subtlety}.json`.
- Shipped orders: `addon-builds.json` (index 0 = level 10).
- BiS: `bis/rogue-*.json`, every band on both factions (weapons) and band 60 bare and raid (all slots, for gear and professions prose).

**Read first:**
- the rogue reports for rounds 1 to 15 with their fix logs
- `round3/race-sims.md`, which is still the only race file with rogue rows
- `ledger.md` through round 15

## What changed since round 15's tree

`git diff 4600b7d82 HEAD` touches rogue in four places, and all of them are round 15's own fixes:
- `rogue-subtlety.json` Backstab note (R15-1): now "every point of Opportunity, from 13, adds 4% to it".
- `rogue-combat.json` Backstab note (R15-2): now "once you have Puncturing Wounds it crits more often ...".
- `combat.md:132` (R15-3): adds "In a dungeon, where the tank holds the mob and you stand at its back, from 20 on, a dagger and Backstab beat the sword, which is what your best in slot table lists."
- `addon-data.json` has been regenerated, and both new notes are on every band.

Re-check of those fixes:
- **R15-1.** Opportunity is 4% a rank, from "by 4%" to "by 20%" (`talents/rogue.json`, Opportunity). The shipped Subtlety order takes it at 13 to 17. Holds on every band.
- **R15-2.** Puncturing Wounds comes at 19 to 21 in the leveling order and 17 to 19 in the raid order, so "once you have" is right. Lethality and Opportunity both list Backstab in their 70755 text, and both are in the Combat `raidBuild`. Holds.
- **R15-3.** Every Combat band 20 to 60 holds a dagger main hand on both factions:
  - 20: Blackfang
  - 30: Toxic Revenger
  - 40 and 50: Gut Ripper
  - 60: Electrified Dagger or Glacial Blade

  Read with "Level as Combat with a sword" before it, the sentence explains the dagger tables without contradicting the solo advice. Holds.

## Builds decoded

Decoded positionally against `talents/rogue.json`, ordered by tier then column.

**Combat `build` 14/34/3:**
- Assassination: Improved Gouge 3, Remorseless Attacks 2, Malice 5, Ruthlessness 3, Improved Slice and Dice 1.
- Combat: Coup de Grace 3, Efficient Strikes 1, Lightning Reflexes 5, Puncturing Wounds 3, Precision 3, Flawless Execution 1, Weapon Expertise 2, Aggression 3, Dual Wield Specialization 2, Blade Flurry 1, Hack and Slash 5, Flash of Steel 4, Adrenaline Rush 1.
- Subtlety: Master of Deception 3.

**Combat `raidBuild` 12/34/5:**
- Assassination: Malice 5, Ruthlessness 2, Improved Slice and Dice 3, Lethality 2.
- Combat: Coup de Grace 3, Lightning Reflexes 4, Puncturing Wounds 3, Precision 3, Composure 2, Flawless Execution 1, Weapon Expertise 2, Aggression 3, Dual Wield Specialization 2, Blade Flurry 1, Hack and Slash 4, Flash of Steel 5, Adrenaline Rush 1.
- Subtlety: Opportunity 5.

**Assassination `build` 38/11/2:**
- Assassination: Improved Gouge 3, Remorseless Attacks 2, Malice 5, Ruthlessness 3, Improved Slice and Dice 2, Lethality 2, Vile Poisons 5, Cold Blood 1, Improved Poisons 5, Relentless Strikes 5, Venom 1, Seal Fate 3, Mutilate 1.
- Combat: Coup de Grace 3, Lightning Reflexes 2, Puncturing Wounds 3, Precision 3.
- Subtlety: Master of Deception 1, Opportunity 1.

**Assassination `raidBuild` 36/10/5:**
- Assassination: Malice 5, Ruthlessness 3, Improved Slice and Dice 3, Vigor 1, Blood Splatter 1, Lethality 2, Vile Poisons 5, Cold Blood 1, Improved Poisons 5, Relentless Strikes 5, Venom 1, Seal Fate 3, Mutilate 1.
- Combat: Coup de Grace 3, Lightning Reflexes 2, Puncturing Wounds 3, Precision 2.
- Subtlety: Opportunity 5.

**Subtlety `build` 12/8/31:** the same decode as round 15.

**Legality.** All five builds are legal:
- Mutilate needs Venom (105712).
- Adrenaline Rush needs Blade Flurry, and Weapon Expertise needs Precision.
- Death Mark needs Preparation, and Hemorrhage needs Serrated Blades.
- Every tier gate holds.

**Build paragraphs.** All three swap paragraphs recount:
- Combat :52: Assassination 5 for 5. In Combat, the Efficient Strikes point and the fifth points of Hack and Slash and Lightning Reflexes replace the fifth point of Flash of Steel and both points of Composure (3 for 3). In Subtlety, Master of Deception 3 replaces Opportunity 5.
- Assassination :52 trades 7 points for 7.
- Combat :63: Composure's toll holds. Tier 0 and tier 1 hold 13 points, and Composure 2 makes the 15 that Flawless Execution, Weapon Expertise and Aggression need.

**Shipped orders** match every level the prose names:
- Assassination: Improved Gouge 10-12, Remorseless Attacks 13-14, Malice 15-19, Venom 30, Mutilate 40, Relentless Strikes 39 and 41-44.
- Combat: Coup de Grace 10-12, Efficient Strikes 13, Lightning Reflexes 14-18, Puncturing Wounds 19-21, Aggression 28-30, Improved Gouge 44-46, Remorseless Attacks 47-48.
- Subtlety: Master of Deception 10-12, Opportunity 13-17, Dirty Deeds 18-19, Premeditation 20, Camouflage 24, Ghostly Strike 26, Hemorrhage 30, Boot Knife 35-39, Death Mark 40.
- **Efficient Strikes comes at 56-57 in both Subtlety orders.** This matters for R16-1.

## Findings

| id | status | kind | file:line | quote | evidence | fix |
|---|---|---|---|---|---|---|
| ROGUE-R16-1 | CONFIRMED | card note | `rogue-subtlety.json` Ghostly Strike notes, rendered on the Subtlety Ghostly Strike line of every band (`addon-data.json` rogue-subtlety, bands 10 to 60) | "Ghostly Strike on cooldown at four combo points or fewer: it hits harder than Hemorrhage for 10 more energy (Hemorrhage costs 30 with Efficient Strikes) and gives a combo point." | Ghostly Strike 14278 costs 40 and Hemorrhage 16511 costs 35 (`spelltips/rogue.json`). Efficient Strikes is 3 at rank 1 and 5 at rank 2, and both Subtlety orders take it at 56 and 57 (`addon-builds.json` rogue-subtlety indices 46 and 47). On every band where both spells are on the bar but Efficient Strikes is not (30 to 54), Ghostly Strike costs 5 more, not 10. At band 56 it costs 8 more. "10 more" is true only at 58 and up. Rounds 1, 9 and 15 checked the note at 60 only. This is the R15 rule: a note that names a talent must hold on every band it renders on. | Note: "Ghostly Strike on cooldown at four combo points or fewer: it hits harder than Hemorrhage for a little more energy (40 against Hemorrhage's 35, or 30 with both points of Efficient Strikes) and gives a combo point." Regenerate `addon-data.json` and Data.lua. |
| ROGUE-R16-2 | CONFIRMED | card note | `rogue-subtlety.json` Garrote notes, rendered on the Subtlety Garrote line of bands 14 to 60 | "... and a fresh Garrote is a bleed that Hemorrhage, Opportunity and Death Mark all feed." | Death Mark is the level 40 capstone (index 30 of both Subtlety orders), and Hemorrhage comes at 30 (index 20). On the bands 14 to 38 cards the note credits the Garrote with a mark the player does not have. On bands 14 to 28 it also credits a Hemorrhage debuff the player cannot apply. Opportunity, from 13, is the only one that holds on every band the line shows. This is the same pattern as R15-1 and R15-2.<br><br>The line also shows from 14, before Ghostly Strike at 26 ("Garrote after every Ghostly Strike"). That is the whole-row issue R15 left with the component lane, and it is not re-filed here. | Note: "Garrote after every Ghostly Strike, at four combo points or fewer, even when the old one is still ticking: Ghostly Strike lets it in without Stealth, and a fresh Garrote is a bleed that Opportunity feeds, and Hemorrhage and Death Mark too once you have them." Regenerate `addon-data.json` and Data.lua. |

No page-prose finding. Every line of the four pages was re-derived against 70755 (list below).

## Shoulds (wording, not findings)

| id | file:line | note | suggested wording |
|---|---|---|---|
| ROGUE-R16-S1 | subtlety.md:80 | Carried from R15-S1. | As R15-S1. |
| ROGUE-R16-S2 | subtlety.md:90 | Carried from R15-S2 ("only damage cooldown"). | As R15-S2. |
| ROGUE-R16-S3 | subtlety.md:117, assassination.md:115, combat.md:116 | Carried from R15-S3, now checked on all three pages against the band 60 raid rows. Each page names "Devilsaur and Stormshroud", but each spec's rows use only one family:<br>- Assassination wears Stormshroud Shoulders, Armor and Pants, and no Devilsaur.<br>- Combat wears Devilsaur Gauntlets and Leggings, plus Outlaw's Collar and Dawn Armor, and no Stormshroud.<br>- Subtlety wears neither: Outlaw's Collar and Dawn Armor, with Devilsaur only on the Horde bare row.<br><br>Not false, since all of them are leatherworking crafts. | Assassination: "the Stormshroud set is a leatherworking craft". Combat: "Outlaw's Collar, Dawn Armor and the Devilsaur pieces are leatherworking crafts". Subtlety: as R15-S3. |
| ROGUE-R16-S4 | `rogue-combat.json` Sinister Strike note | Carried from R15-S4. The note gives "no dagger in the main hand" as the only reason. combat.md:132 adds the solo reason ("mobs face you"), and Assassination's and Subtlety's notes include "or the target is facing you". | As R15-S4. |

## Checked and right (re-derived this round)

**Spell texts** (`spelltips/rogue.json`):
- Cooldowns: Adrenaline Rush 15 s on 5 min, Blade Flurry 25 energy on 2 min with 20% speed for 15 s, Cold Blood 3 min.
- Combat :87 "two Blade Flurries and one Adrenaline Rush" on a 3 minute kill holds.
- Costs and cooldowns: Kidney Shot 25 energy on 20 s, up to 6 s at five (rank 2). Kick 25, Feint 20 on 10 s. Cheap Shot 60, 4 s, 2 points.
- Defensives: Gouge 4 s, target facing you. Evasion 15 s. Blind and Vanish 5 min, so 3.5 min with Elusiveness 2/2.
- Preparation 10 min. Premeditation 2 points for 20 s on 2 min, and its text has no Stealth clause.
- Builders and openers: Garrote 50 (30 with Dirty Deeds 2/2), Ambush 60 (dagger, behind, stealthed), Backstab 60.
- Utility: Poised 10, Redirect 15 s, Shank 5 energy for 15 then 25 on 10 s, cannot miss.
- Mutilate rank 3 is 90% plus 79.2 a hand. Hemorrhage is 145% with a dagger, +15% to Rupture and Garrote. Ghostly Strike is 180% with a dagger, 15% dodge for 7 s, on 20 s.
- Death Mark: 25 energy, 10 s, a 6 s mark, Garrote and Rupture at twice the rate.

**Talent texts:**
- Lethality is +20% on the crit bonus, for Sinister Strike, Gouge, Backstab, Mutilate, Ghostly Strike and Hemorrhage. It does not cover Ambush, so assassination.md:72 holds.
- Opportunity covers Backstab, Garrote, Ambush and Mutilate.
- Puncturing Wounds gives Backstab 30% crit and Mutilate 15%, plus a 45% second point.
- Remorseless Attacks 40%, Improved Gouge 1.5 s.
- Vigor 5 a point, Composure 5 a point.
- Seal Fate 100% at 3/3.
- Relentless Strikes 20% a point for 25.
- Venom 30% and 20%.
- Flash of Steel +15 and -2 s.
- Hack and Slash dagger crit 5% at 5/5.
- Coup de Grace 20%, Serrated Blades 9% and 30%, Blood Splatter 6 s and 9 s.
- Dirty Deeds 20 and frontal Garrote.
- Master of Deception "3 levels higher".

**Races.** Every race line on the four pages matches `round3/race-sims.md:18-23`. No newer measured rogue race result is in the tree. Racial texts match `races.json`:
- Horde: Touch of the Grave, Blood Fury, Berserking, Will of the Forsaken.
- Alliance: Elune's Light, Eureka! on the rogue row (non-periodic damaging), Expansive Mind, Ruthless, Artisan's Precision.
- Skyborne: Wind Blessed 2% haste. Skyborne is named only as the paid Heroic Pack option.
- BiS race captions (night-elf and troll) are known, and not filed.

**Gear.**
- Assassination and Combat hold two daggers on every band and faction. Subtlety's band 60 rows hold Teebu's Blazing Longsword and Persuader, which agrees with subtlety.md:33, :70 and :105.
- Heartseeker is on the Assassination band 60 rows and the Combat band 60 bare rows, so combat.md:116 "when it is on your list" holds.
- The Molten and Frozen Hearts are on every spec's band 60 rows.

**Ranking claims.**
- Band 60 raid set DPS ranks Combat over Assassination over Subtlety on both factions, so subtlety.md:40 "lowest raid damage" and combat.md:30 "steadiest" hold.

**Sources.** All six URLs returned 200 on 2026-11-04 with a browser agent:
- the two Blizzard articles
- the Wowhead Skyborne article
- foreversixty.gg /planner, /sim/weights and /planner?class=rogue

**Voice.**
- No em or en dashes in the pages or the notes.
- No site, sim, APL or band talk, and no nightly numbers.

## Decisions

- **The R15 band rule applies to costs as well as talents.** A card note's number must hold on every band where the line renders and its talents are present. A row that renders before its own talent stays with the component lane. R16-1 is filed only for bands 30 to 56, where Ghostly Strike and Hemorrhage both exist. R16-2 is filed for the Hemorrhage and Death Mark clause only.
- **No prose, build, order, `statPriority` or race change.**

## Sims requested

None new. Carried, routed to sim round 18 or the data lane, and not re-filed:
- Subtlety band 20 BiS: Diamond Hammer should become a dagger.
- Subtlety BiS and weights on the Dirty Deeds build. The band 60 rows still carry `035131031321001151`.
- Race re-runs for all three specs on the current builds.
- Combat band 60 raid re-rank on the shipped `raidBuild`. The BiS rows carry `3243030000012321451`, the old Efficient Strikes string.

## Sim findings

None new. The open items from sim round 17 are untouched:
- K1: Death Mark storage, and the Mutilate and Shank off-hand penalty.
- K2/C1: Combat Eviscerate at four.
- X1: Occult Poison learnability.
- Round 16 S4: Hemorrhage's dagger hand.

The pages make no claim on any of these that the client contradicts.

## Sources

- "World of Warcraft: Forever Class Deep Dives - Rogue and Warlock", Blizzard: https://news.blizzard.com/en-us/article/24310968/world-of-warcraft-forever-class-deep-dives-rogue-and-warlock
- "World of Warcraft: Forever Deep Dive Panel Recap", Blizzard: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
- "Skyborne First Look: New Neutral Race in World of Warcraft: Forever", Wowhead: https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829
- foreversixty.gg: /planner, /sim/weights, /planner?class=rogue
- Client data, `data/builds/1.60.1.70755/`:
  - `talents/rogue.json`
  - `spelltips/rogue.json`: 14278, 16511, 11290, 1833, 13877, 13750, 14177, 8643, 1329283, 1329549, 1241584, 1325758, 14183, 14185, 1328825, 438040
  - `races.json`
  - `addon-builds.json`, `addon-data.json`
  - `bis/rogue-*.json`, `items.json`, `dps-per-point.json`
- APLs: `data/curated/apl/rogue-{assassination,combat,subtlety}.json`
- Previous rounds: `round15/rogue.md`, `round14/rogue.md`, `round3/race-sims.md`, `ledger.md`

## Counts

CONFIRMED 2, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0. SHOULDs: 4 (all carried). New sim findings: 0.

Verdict: NOT zero for rogue. Both confirmed findings are in the Subtlety card notes. Each states something that holds only at 60 on cards that render from much earlier. The page prose of all four rogue pages, and round 15's three fixes, are clean.

## Fix log (round 16)

- ROGUE-R16-1: fixed, `data/curated/apl/rogue-subtlety.json` Ghostly Strike notes: "for a little more energy (40 against Hemorrhage's 35, or 30 with both points of Efficient Strikes)". Costs checked in 70755 spelltips (14278 40, 16511 35) and talents (Efficient Strikes 3/5). Notes only; `addon-data.json` and Data.lua regenerate at the integrator.
- ROGUE-R16-2: fixed, same file, Garrote notes: "a bleed that Opportunity feeds, and Hemorrhage and Death Mark too once you have them". Notes only.
- S1 to S4: left (not in this round's scope).
