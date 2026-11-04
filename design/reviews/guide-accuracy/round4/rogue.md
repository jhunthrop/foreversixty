# Guide accuracy round 4: Rogue

- Class: rogue. I audited every line of `web/src/content/guides/rogue/index.md`, `assassination.md`, `combat.md` and `subtlety.md` at land/guides-r4 752c02e43. I also read the card notes in `data/curated/apl/rogue-*.json`.
- Client 1.60.1.70755: `talents/rogue.json`, `trainables/rogue.json`, `addon-builds.json`, `races.json`, `items.json`, `items/rogue.json`, `loot.json`. Engine pin d1d9f2775 (`sim/enginever/version.go`).
- Read first: the rogue reports for rounds 1, 2 and 3, `round2/race-sims.md`, `round3/race-sims.md`, `ledger.md` and `round4/rogue-levels.md`. Not re-filed: SIM-1 to SIM-10, SIM-R2-1 to SIM-R2-3, SIM-R3-1 and SIM-R3-2, the Mutilate off-hand engine change, the spelltips regen (R2-6, data lane), card actions, and the Combat leveling order (SIM-R2-3, which needs a ladder-first row).

## Decoded builds (unchanged, re-verified)

All five frontmatter strings decode as in rounds 1 to 3: Assassination 38/11/2 and 36/10/5, Combat 14/34/3 and 12/34/5, Subtlety 12/8/31. Each is legal on 70755 (tier gates met; prerequisites Venom for Mutilate, Malice for Lethality, Precision for Weapon Expertise, Blade Flurry for Adrenaline Rush, Serrated Blades for Hemorrhage, Preparation for Death Mark). Each equals the final state of its `addon-builds.json` order. Combat's race segment is `gnome`, matching `recommendedRaces[0]`.

Shipped leveling orders (index 0 = level 10), re-derived:
- **Assassination:** Improved Gouge 10-12, Remorseless Attacks 13-14, Malice 15-19, Ruthlessness 20-22, Improved Slice and Dice 23-24, Lethality 25-26, Vile Poisons 27-29 and 31-32, Venom 30, Cold Blood 33, Improved Poisons 34-38, Relentless Strikes 39 and 41-44, Mutilate 40, Seal Fate 45-47, Combat 48-58, Master of Deception 59, Opportunity 60.
- **Combat:** Coup de Grace 10-12, Efficient Strikes 13, Lightning Reflexes 14-18, Puncturing Wounds 19-21, Precision 22-24, Flawless Execution 25, Weapon Expertise 26-27, Aggression 28-30, Dual Wield Specialization 31-32, Blade Flurry 33, Hack and Slash 34-38, Flash of Steel 39 and 41-43, Adrenaline Rush 40, Improved Gouge 44-46, Remorseless Attacks 47-48, Malice 49-53, Ruthlessness 54-56, Improved Slice and Dice 57, Master of Deception 58-60.
- **Subtlety:** Master of Deception 10-12, Opportunity 13-17, Dirty Deeds 18, Camouflage 19, 24 and 27, Premeditation 20, Serrated Blades 21-23, Elusiveness 25 and 33, Ghostly Strike 26, Improved Ambush 28, Initiative 29 and 31-32, Hemorrhage 30, Preparation 34, Boot Knife 35-39, Death Mark 40, Assassination 41-52, Combat 53-60.

## The new "Levels 1 to 19" blocks: what checks out

- **Training levels** (`trainables/rogue.json`): Sinister Strike, Eviscerate, Stealth and Redirect at 1; Backstab 4; Gouge 6; Evasion 8; Sap, Slice and Dice and Sprint 10; Kick 12; Garrote and Expose Armor 14; Feint 16; Ambush and Poised 18. Every level the three blocks name is right, and none of them names Cheap Shot (26) or a poison (20) under 20.
- **Costs:** Ambush rank 1 costs 60. Poised 1328825 reads "While at maximum Energy, and out of combat, reduce the Energy cost of your next ability by 10", so Subtlety's "Ambush from stealth costs 50" is right.
- **Talent text against the shipped order:**
  - Assassination: Improved Gouge 3/3 is 1.5 s. Remorseless Attacks 2/2 is 40% on Sinister Strike, Backstab, Ambush, Mutilate or Ghostly Strike after a kill. Malice 5/5 is 5%.
  - Combat: Coup de Grace 3/3 is 20% Eviscerate. Efficient Strikes 1/2 is 3 energy off Sinister Strike. Lightning Reflexes 5/5 is 5% dodge.
  - Subtlety: Master of Deception 3/3 is "as if you were 3 levels higher". Opportunity 5/5 is 20% to Backstab, Garrote, Ambush and Mutilate. Dirty Deeds 1/2 is 10 energy off Garrote and lifts the behind requirement. Camouflage 1/3 cuts the Stealth speed penalty.
  - All of these match the levels the prose gives.
- **Items:** Worn Dagger (2092, dagger) and Small Throwing Knife (2947, thrown) exist in `items.json`.
- **Consistency with "Levels 20 to 60" and the strip:**
  - Assassination's 20-60 text ("Remorseless Attacks and Improved Gouge make chain pulls quicker") agrees with the 10-14 points.
  - Combat's 20-60 text no longer repeats Efficient Strikes. It keeps "Aggression by 30" (28-30) and "Improved Gouge and Remorseless Attacks come in the mid 40s" (44-48).
  - Subtlety's 20-60 text starts at Premeditation 20 and matches 26/30/35/40.
  - The strip shows each band's first-level build from the same `addon-builds.json` orders, so no block contradicts it.

## Fresh pass on the rest of the pages: what checks out

- Every talent value in the prose matches its 70755 rank text: Mutilate 79 a hand at the top rank, Venom 30%/20%/6 s, Relentless Strikes 20% per point for 25, Lethality 200% to 220%, Coup de Grace 20%, Puncturing Wounds 30%/45%, Flash of Steel +15/-2 s, Aggression 15%, Flawless Execution 15, Hack and Slash, Serrated Blades 9%/30%, Boot Knife, Death Mark, Ghostly Strike, Hemorrhage, Initiative, Blood Splatter 6/9 s, Dirty Deeds, Efficient Strikes.
- I recounted the build summaries: Assassination trades 7 points for 7; Combat trades 11 points for 11. Combat's raid row-3 gate is 15 (Coup de Grace 3, Efficient Strikes 2, Lightning Reflexes 4, Puncturing Wounds 3, Precision 3).
- Racial text matches `races.json`: Touch of the Grave, Blood Fury, Berserking (10% for 15 s, 3 min), Elune's Light, Eureka!, Expansive Mind 5%, Will of the Forsaken, Wind Blessed 2%. Skyborne is named only as the paid option.
- The Blizzard Rogue and Warlock deep dive (read 2026-11-04) backs these prose lines:
  - Expose Armor: "no longer prevents Sunder Armor from being used ... Sunder Armor still won't overwrite it unless its effect is weaker".
  - Poisons stack with stones.
  - Mutilate needs two daggers and costs 60.
  - Assassination's finishers: "damaging finishers like Rupture and Eviscerate are a lower priority".
- Round 3 fixes re-checked:
  - R3-2 to R3-7 are right.
  - R3-1 is right on every page except one clause in combat.md:112 (ROGUE-R4-2).
- Card notes: every note states its own action truthfully. The accepted card-action items stand.
- Voice: no em or en dashes; no site, sim, APL or band talk; no quoted nightly numbers.

## Source links

| Link | Pages | Status | Says what its label says? |
|---|---|---|---|
| https://www.warcrafttavern.com/forever/news/poison-pets-rogue-warlock-deep-dive-for-wow-forever/ | all four | 200 | **No**. The label is "Blizzard, Rogue and Warlock class deep dive" with kind `blizzard`, but the page is a Warcraft Tavern summary by Luxrah (7 Oct 2026). It links the Blizzard original itself. See ROGUE-R4-1. |
| https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap | all four | 200 | Yes: "World of Warcraft: Forever Deep Dive Panel Recap". It only touches rogues in passing (poisons in itemization, level-up weapons). |
| https://foreversixty.gg/sim/weights | three specs | 200 | Yes |
| https://foreversixty.gg/planner | all four | 200 | Yes |
| https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829 | index | 200 | Yes. The title matches; the body renders client-side, so only the headline could be read. |
| /planner?class=rogue (body) | three specs | site route | Yes |

## Findings

| id | status | kind | file:line | wrong claim | evidence | exact fix |
|---|---|---|---|---|---|---|
| ROGUE-R4-1 | CONFIRMED | source | index.md:8-10, assassination.md:14-16, combat.md:14-16, subtlety.md:13-15 | `label: 'Blizzard, Rogue and Warlock class deep dive'`, `url: https://www.warcrafttavern.com/forever/news/poison-pets-rogue-warlock-deep-dive-for-wow-forever/`, `kind: blizzard` | The fetched page is a Warcraft Tavern article by Luxrah that summarizes the Blizzard post and links it as a reference. Blizzard's own post, "World of Warcraft: Forever Class Deep Dives — Rogue and Warlock" (7 Oct), resolves at https://news.blizzard.com/en-us/article/24310968/world-of-warcraft-forever-class-deep-dives-rogue-and-warlock (200; https://worldofwarcraft.blizzard.com/en-us/news/24310968 lands on the same post). The druid pages already cite their Blizzard deep dive on news.blizzard.com in this form. | In all four files: `url: https://news.blizzard.com/en-us/article/24310968/world-of-warcraft-forever-class-deep-dives-rogue-and-warlock`, keeping label 'Blizzard, Rogue and Warlock class deep dive' and kind `blizzard`. Every claim the pages lean on it for appears in the Blizzard post: Mutilate on two daggers for 60 energy, Expose Armor and Sunder, poisons with stones, Assassination's finisher priority. |
| ROGUE-R4-2 | CONFIRMED (minor; a round 3 edit) | race (measured) | combat.md:112 | "Orc and Night Elf are next; Troll trails." | `round3/race-sims.md:20-21`: Combat Alliance Gnome 759.3, Night Elf 755.9, Human 755.1, Dwarf 753.4; Horde Undead 759.3, Orc 756.6, Troll 756.0 (±0.32). Per seed: Troll 756.39/755.68 against Night Elf 755.98/755.85 (lines 62, 65). Troll is level with Night Elf, ahead of Human and Dwarf, and within error of Orc. "Trails" is false for Combat. Troll does trail for Subtlety (636.0 vs Orc 638.1), so the index and the other spec pages are fine. | "Orc, Troll and Night Elf all land a step behind and within a hair of each other." |
| ROGUE-R4-3 | LIKELY (minor) | leveling (new block) | combat.md:124 | "and Puncturing Wounds starts at 19 for Backstab crits and the odd extra combo point." | Puncturing Wounds 1224716 names only Backstab and Mutilate. The same block (line 125) says to replace the Worn Dagger with a sword, axe or mace, and line 129 levels on Sinister Strike. So this point does nothing for the rogue the page describes until a dagger goes back in the main hand. The order itself is the accepted ladder item (SIM-R2-3); only the wording is filed. | "...and Puncturing Wounds from 19 only pays off when you have a dagger in the main hand and get behind the mob for Backstab." |
| ROGUE-R4-4 | LIKELY (minor) | leveling (new block) | assassination.md:124 | "Drop the Worn Dagger the moment a quest offers a better weapon" | Line 121 of the same block builds the whole pull on "With a dagger in hand, Gouge, step behind and Backstab", and Backstab and Ambush need a main-hand dagger. A quest sword at 12 would take both away. The Subtlety block already says "the first better quest dagger". | "Drop the Worn Dagger the moment a quest offers a better dagger, so Backstab and Ambush stay open to you." |

## Decisions

- Source: cite Blizzard's own post for "Blizzard, ..." labels, in the news.blizzard.com form the druid pages use. Do not relabel the Warcraft Tavern summary.
- Combat race wording: the ranking follows `round3/race-sims.md` as measured. Gnome and Undead are unchanged; only the order of the runners-up is corrected.
- No build or race changes.

## Sims requested

None new. Carried for the sim round: the Mutilate off-hand penalty, SIM-7 (Preparation resets), SIM-1 to SIM-5, SIM-8, SIM-10, SIM-R2-1 to SIM-R2-3, SIM-R3-1 (Gnome Assassination Eureka! masks) and SIM-R3-2 (Combat Thistle Tea inside Adrenaline Rush).

## Sim findings

None new on d1d9f2775.

## Sources

- "World of Warcraft: Forever Class Deep Dives — Rogue and Warlock", Blizzard: https://news.blizzard.com/en-us/article/24310968/world-of-warcraft-forever-class-deep-dives-rogue-and-warlock
- "Poison & Pets: Rogue & Warlock Deep Dive for WoW Forever", Warcraft Tavern (Luxrah): https://www.warcrafttavern.com/forever/news/poison-pets-rogue-warlock-deep-dive-for-wow-forever/
- "World of Warcraft: Forever Deep Dive Panel Recap", Blizzard: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
- "Skyborne First Look: New Neutral Race in World of Warcraft: Forever", Wowhead: https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829
- Round 3 race runs: `design/reviews/guide-accuracy/round3/race-sims.md` (this tree)

## Fix log (round 4)

- ROGUE-R4-1 fixed: deep-dive `url` now https://news.blizzard.com/en-us/article/24310968/world-of-warcraft-forever-class-deep-dives-rogue-and-warlock, label and kind `blizzard` kept. index.md:9, assassination.md:15, combat.md:15, subtlety.md:14.
- ROGUE-R4-2 fixed: combat.md:112 now "Orc, Troll and Night Elf all land a step behind and within a hair of each other."
- ROGUE-R4-3 fixed (verified: Puncturing Wounds 1224716 names only Backstab and Mutilate): combat.md:124 now says the point pays off only with a main-hand dagger and Backstab from behind.
- ROGUE-R4-4 fixed: assassination.md:124 now "a better dagger, so Backstab and Ambush stay open to you".
- Round 4 source rules: site labels changed to "Forever Sixty, stat weights" (assassination.md:20, combat.md:20, subtlety.md:19; all damage specs, link stays /sim/weights) and "Forever Sixty, character planner" (index.md:17, assassination.md:23, combat.md:23, subtlety.md:22). No /sources, wago, "Talents Forever", rotation or tank-fight entries on rogue pages. Planner body links already in the one form.
- Decisions applied: Blizzard original cited; Combat runners-up reordered as measured; no build, race or card-action changes. Card notes unchanged (all truthful per the audit).
- New SIM findings: none.
