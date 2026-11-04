# Guide accuracy round 7: Rogue

- Class: rogue. I audited every line of `web/src/content/guides/rogue/index.md`, `assassination.md`, `combat.md` and `subtlety.md` in the round 7 audit tree (land/guides-r6 394562f20, all six rounds applied). I also read every rendered card note (`rotation.priorityList[*].notes`) and the prepull actions in `data/curated/apl/rogue-*.json`.
- Client 1.60.1.70755 data read:
  - `talents/rogue.json`: every rank text and every prerequisite.
  - `trainables/rogue.json`: every rank level, cost and cooldown.
  - `spelltips/rogue.json`: descriptions for Backstab, Ambush, Garrote, Gouge, Mutilate, Shank, Redirect, Poised, Kidney Shot, Death Mark, Hemorrhage, Ghostly Strike, the poisons and others.
  - `races.json`, `items.json`, `items/rogue.json` and `simconsumes.json` (Thistle Tea).
  - `addon-builds.json`.
  - `bis/rogue-{assassination,combat,subtlety}.json`: band 60 bare and raid, both factions, with weights, `hit_to_cap` and weapon slots.
- Engine pin: d1d9f2775 (`sim/enginever/version.go`).
- Render order checked in `web/src/pages/guides/[class]/[spec].astro`:
  - Stat pills and the stat table come before the Stat priority prose.
  - The best in slot table comes before the Best in slot prose.
  - The race pills and the race table come before the Races prose, and Professions follows straight after the race table.
- Read first: the rogue reports for rounds 1 to 6 (with every fix log and fix pass), `round4/rogue-levels.md`, the race-sims files for rounds 2 to 5 (only round 3 has rogue rows) and `ledger.md`.
- Not re-filed:
  - sim items: SIM-1 to SIM-10, SIM-R2-1 to SIM-R2-3, SIM-R3-1, SIM-R3-2, SIM-R5-1 and SIM-R6-1
  - the Mutilate off-hand engine change
  - card actions and the Combat leveling order (ladder-first)
  - component-lane items and data-lane items, including the BiS race captions, which still read night-elf and troll
  - the index's generic "Troll trails" (round 4 accepted it)

## Decoded builds (re-verified, unchanged)

| Spec | build | raidBuild |
|---|---|---|
| Assassination | 38/11/2 | 36/10/5 |
| Combat | 14/34/3 (race `gnome`) | 12/34/5 (race `gnome`) |
| Subtlety | 12/8/31 (single build) | none |

- Every talent matches rounds 1 to 6.
- Every build is legal on 70755. The prerequisites check out against `talents/rogue.json`: Lethality needs Malice 5, Mutilate needs Venom, Weapon Expertise needs Precision 3, Adrenaline Rush needs Blade Flurry, Hemorrhage needs Serrated Blades 3, and Death Mark needs Preparation.
- Each build equals the final state of its `addon-builds.json` order.

I re-decoded the leveling orders level by level (index 0 = level 10). They are identical to round 6:
- **Assassination:** Improved Gouge 10-12, Remorseless Attacks 13-14, Malice 15-19, Ruthlessness 20-22, Improved Slice and Dice 23-24, Lethality 25-26, Vile Poisons 27-29 and 31-32, Venom 30, Cold Blood 33, Improved Poisons 34-38, Relentless Strikes 39 and 41-44, Mutilate 40, Seal Fate 45-47, Combat 48-58, Master of Deception 59, Opportunity 60.
- **Combat:** Coup de Grace 10-12, Efficient Strikes 13, Lightning Reflexes 14-18, Puncturing Wounds 19-21, Precision 22-24, Flawless Execution 25, Weapon Expertise 26-27, Aggression 28-30, Dual Wield Specialization 31-32, Blade Flurry 33, Hack and Slash 34-38, Flash of Steel 39 and 41-43, Adrenaline Rush 40, Improved Gouge 44-46, Remorseless Attacks 47-48, Malice 49-53, Ruthlessness 54-56, Improved Slice and Dice 57, Master of Deception 58-60.
- **Subtlety:** Master of Deception 10-12, Opportunity 13-17, Dirty Deeds 18, Camouflage 19, 24 and 27, Premeditation 20, Serrated Blades 21-23, Elusiveness 25 and 33, Ghostly Strike 26, Improved Ambush 28, Initiative 29 and 31-32, Hemorrhage 30, Preparation 34, Boot Knife 35-39, Death Mark 40, Assassination 41-52, Combat 53-60.

## The band 60 tables the stat paragraphs sit under (current files, generated 2026-11-04 04:27 to 04:36)

| Spec | Raid order (both factions unless noted) | `statPriority` |
|---|---|---|
| Assassination | Hit 1.99, Agility 1.00, Crit 0.82, Strength 0.59, Attack power 0.53, haste | same order |
| Combat | Hit 2.39, Agility 1.00, Strength 0.74, Attack power 0.67, Crit 0.52, haste | same order (round 6 fix) |
| Subtlety | Alliance: Hit 1.05, Agility 1.00, Crit 0.83, Strength 0.57, Attack power 0.52. Horde: Hit 1.36, Agility 1.00, Crit 0.92, Strength 0.52, Attack power 0.48 | same order |

- No rogue item in `items/rogue.json` carries a haste stat. The only attack-speed effects are a few procs and use effects (Empyrean Demolisher, Eskhandar's Right Claw, Kiss of the Spider), so "haste is rare on gear" is true.
- Weapons on the band 60 raid tables:
  - Assassination: Electrified Dagger or Glacial Blade, with Heartseeker.
  - Combat: Electrified Dagger or Glacial Blade, with Sentinel's Blade or Scout's Blade. All of these are daggers.
  - Subtlety: Teebu's Blazing Longsword (a main-hand sword, `inventory_type` 21) with Persuader (a mace), on both factions.

## What checks out (fresh pass)

**Stat priority prose (round 6 rewrite)**
- No paragraph ranks stats. Every stated reason holds on 70755:
  - White swings with two weapons miss far more than specials: `hit_to_cap` white 18.8 to 25 against specials 0 to 6.
  - Agility gives attack power and crit.
  - Seal Fate (3/3 at 100%) and Lethality (2/2 at 20%) make builder crits worth more for Assassination.
  - Death Mark's cap carries `$mcrit/100+1`, and bleed ticks crit (forever-dots-crit).
  - Strength and attack power add flat damage. Haste is rare.
- The "table above" in each Stat priority paragraph is right: the table renders above the prose.

**Ranking by a nightly value elsewhere:** none found. The weapon lines are kept on mechanics:
- Assassination: the main hand carries Deadly Poison and the white hits.
- Combat: Backstab, Puncturing Wounds, Lethality and Hack and Slash crit on daggers.
- Subtlety: "the weapons in the table above".
- The poison pairing and race lines are measured picks, not nightly values.
- Professions name the Hearts only as bind-on-pickup enchanting crafts.

**Levels 1 to 19, then 20 to 60**
- Every training level matches `trainables/rogue.json`:
  - Sinister Strike, Eviscerate, Stealth and Redirect at 1.
  - Backstab 4, Gouge 6, Evasion 8, Slice and Dice and Sprint 10, Kick 12, Garrote 14, Ambush and Poised 18.
  - Instant Poison 20, Shank 24 (rank 2 at 44), Cheap Shot 26, Deadly Poison 30, Mutilate 40 (ranks 50 and 60), Death Mark 40.
- Ambush rank 1 costs 60, so "costs 50" with Poised is right.
- Every talent named at a level is in the order at that level, with the rank value quoted:
  - Improved Gouge 1.5 s, Remorseless Attacks 40%, Malice 5%.
  - Coup de Grace 20%, Efficient Strikes 3, Lightning Reflexes 5%.
  - Master of Deception 3 levels, Opportunity 20%, Dirty Deeds rank 1 (10 energy and Garrote from the front), Camouflage rank 1.
- In each spec the two blocks read correctly in order. Subtlety's 1 to 19 block now opens with "level Combat and respec", and round 6's R6-5 and R6-6 edits read right.
- Worn Dagger (2092) and Small Throwing Knife (2947) exist.

**Talents, spells and arithmetic**
- Every value the prose quotes matches its rank or spell text:
  - Mutilate: "79.2 with each weapon" in the spelltip.
  - Shank: 15 or 25 energy, cannot miss, off-hand 50%.
  - Venom: 30%/20%/6 s. Relentless Strikes: 20% a point for 25.
  - Lethality: 200% to 220%. Vile Poisons: 20%. Improved Poisons: 10% and 50%.
  - Vigor: 5 a point. Blood Splatter: 6 s and 9 s.
  - Coup de Grace: 20% (Eviscerate and Death Mark). Flash of Steel: 15 and 2.0 s. Hack and Slash: 5% at five points.
  - Dual Wield Specialization: 50%. Puncturing Wounds: 30%/15%/45%. Aggression: 15%. Flawless Execution: 15.
  - Serrated Blades: 9%/30%. Hemorrhage: 145% with a dagger, Rupture and Garrote. Ghostly Strike: 180% with a main-hand dagger, 15% dodge, 7 s.
  - Boot Knife: 2 points. Death Mark: 25 energy, 10 s, 6 s, 12% a point, bleeds at double rate. Initiative: 100%.
  - Premeditation: no Stealth gate, 20 s, 2 min. Preparation: 10 min. Elusiveness: 90 s.
  - Kidney Shot rank 2: 25 energy, 20 s, 6 s at five points. Kick: 25. Blade Flurry: 25 energy, 2 min, 20%.
  - Adrenaline Rush: +100% regeneration, 5 min. Cold Blood: 3 min. Blind: 5 min. Feint: 10 s.
  - Redirect: 15 s, next enemy hit. Poised: 10 at full energy out of combat. Thistle Tea: item 7676, spell 9512.
- Build swaps recount 7 for 7 (Assassination) and 11 for 11 (Combat). The tolls hold: Combat raid tier 3 needs exactly 15, and Assassination's tier 1 needs Coup de Grace 3 and Lightning Reflexes 2. Subtlety's Improved Ambush point is what makes the 25 for Boot Knife.
- The pull and reset arithmetic holds:
  - Combat's Garrote opener is 40 with Poised, plus a 60 Backstab.
  - Subtlety's four-point opener and its five-point Preparation reset.
  - Two Blade Flurries and one Adrenaline Rush on a 3 minute kill.

**Races, sources, card notes and voice**
- Racials match `races.json` 70755: Touch of the Grave, Blood Fury, Berserking, Elune's Light, Ruthless, Artisan's Precision, Eureka!, Expansive Mind, Will of the Forsaken and Wind Blessed. The race orders follow `round3/race-sims.md:18-23`. Skyborne is named only as the paid option.
- Source links (all 200 on 2026-11-04):
  - the Blizzard Rogue and Warlock deep dive (24310968)
  - the Deep Dive panel recap (24303313)
  - Wowhead Skyborne First Look
  - foreversixty.gg/planner, /planner?class=rogue and /sim/weights
- Each source follows the round 4 rule. The deep dive still backs:
  - poisons beside stones and more charges
  - Expose Armor against Sunder
  - one-handed axes
  - Mutilate on two daggers for 60 energy, with Seal Fate once a cast
  - Assassination's finisher priority
- Card notes all state their own action truthfully. One example: Combat's Sinister Strike note, "held to 60 energy", matches its `currentEnergy >= 60` condition.
- No em or en dashes, no non-ASCII characters, no site, sim, APL, preset or band talk, and no quoted nightly numbers.

## Findings

| id | status | kind | file:line | wrong claim | evidence | exact fix |
|---|---|---|---|---|---|---|
| ROGUE-R7-1 | CONFIRMED (minor) | spell vs the page's own gear | subtlety.md:70 (and :62) | :70 "Ambush is the opener for anything that will die fast, not for a boss." :62 "Opportunity (5 points) adds 20% to Garrote and Ambush." | **Ambush needs a main-hand dagger.** Ambush 8676 and 11269 read: "Must be stealthed and behind the target. Requires a dagger in the main hand." (`spelltips/rogue.json`; Backstab 53 has the same clause).<br>**The page's own table gives a sword.** On both factions, bare and raid, the Subtlety main hand is Teebu's Blazing Longsword 1728, a main-hand sword. Line 101 tells the reader to "take the weapons in the table above, dagger or not".<br>**The result.** A Subtlety rogue geared as this page says cannot cast Ambush at all. The trash-opener advice and Opportunity's Ambush half do nothing for them. With a sword, Opportunity at 60 adds only to Garrote, since Backstab and Mutilate need daggers too. | :70: "Ambush needs a dagger in the main hand, so with the sword your table gives you it is off your bar; on trash that dies fast, open with Cheap Shot, which Initiative makes three combo points, or Garrote." :62: "Opportunity (5 points) adds 20% to Garrote, and to Ambush and Backstab when you hold a dagger." |
| ROGUE-R7-2 | LIKELY (minor) | stats (reason reads wrong) | combat.md:100 | "...but Combat has no Seal Fate and keeps no bleed running, so a crit never adds a combo point." | **The "so" clause does not follow.** A bleed has nothing to do with combo points. Seal Fate (14186: "critical strikes from abilities that add Combo Points ... add an additional Combo Point") is the only crit-to-combo-point link, and the Combat raid build (12/34/5) has none.<br>**The bleed point is a separate reason.** It is about what a crit is worth: no Rupture or Garrote of yours ticks all fight, so no bleed tick can crit (forever-dots-crit), which Subtlety's paragraph uses. Both facts are true; joined with "so", the sentence states a cause that is not one. | "A crit makes one hit bigger, and Lethality makes your Backstab crits bigger still, but Combat has no Seal Fate, so a crit never adds a combo point, and no bleed of yours runs all fight for a crit to tick on." |
| ROGUE-R7-3 | LIKELY (minor) | wording (page reference) | assassination.md:115 | "...bind-on-pickup trinkets only an enchanter can make; the table above shows which of them your set wants." | **The nearest table is the race table.** `[spec].astro` renders Professions straight after `RaceTable` and the Races prose, so the table directly above this line is the race table. The trinkets are in the best in slot table, two sections up. Round 6's ledger already counted wrong "above" and "below" references as errors on other classes. | "...bind-on-pickup trinkets only an enchanter can make; your best in slot table shows which of them your set wants." |
| ROGUE-R7-4 | LIKELY (minor) | consumable (overstated) | assassination.md:88 | "Refresh poisons before every pull; a Mutilate into an unpoisoned boss is a weak one." | **A coat outlasts many pulls.** Deadly Poison V 25351: "lasts for 30 minutes ... 240 charges". Instant Poison VI 11340: "30 minutes ... 250 charges". The raid build's Improved Poisons 5/5 gives "a 50% chance to not consume a charge".<br>**Blizzard says the same.** The deep dive says "Charges increased on all poisons", and that Improved Poisons "helps their Poisons last much longer without needing to be reapplied".<br>**What the advice should be.** A rogue does not re-coat before every pull. They check both coats before a boss, which is what combat.md:71 already says ("Check your poisons and stones before the pull"). | "Check both coats before every boss pull and re-apply any that is close to running out; a Mutilate into an unpoisoned boss is a weak one." |

## Decisions

- No build, race, `statPriority` or card-action changes. The three `statPriority` lists match the current band 60 tables.
- **R7-1 is filed as CONFIRMED.** Two hard data lines (the Ambush spell text and the table's main-hand sword) make the sentence false for the reader the page itself gears. The Initiative and Ghostly Strike lines stay: they describe the mechanic correctly, and Garrote still benefits.
- **Not filed: weapon skills.** The Combat 1 to 19 advice to take "the first quest sword, axe or mace" does not mention training the weapon skill. I could not confirm which weapon skills a Forever rogue starts with. Warcraft Wiki says weapon skills "remain in Classic and Forever" but lists no rogue starting set. So this is not a finding. A fix lane may add "train the skill at a weapon master" if the client data shows it.
- **Not re-filed:** the index's Horde "Troll trails" (accepted in round 4) and Dwarf's position against Human in the Assassination tie (the Assassination page says all three land in the same place).

## Sims requested

None new. Carried for the sim round:
- SIM-R6-1 (Combat crit: is the 0.52 scale factor the white crit cap and the lack of Seal Fate, or a gap in how Hack and Slash, Puncturing Wounds or Lethality are modelled?)
- the Mutilate off-hand penalty
- SIM-1 to SIM-5, SIM-7, SIM-8 and SIM-10
- SIM-R2-1 to SIM-R2-3
- SIM-R3-1 and SIM-R3-2
- SIM-R5-1

## Sim findings

None new on d1d9f2775. SIM-3 is still open: the Subtlety Slice and Dice line (priority 6) is still shadowed by the Eviscerate line (priority 5).

## Sources

- "World of Warcraft: Forever Class Deep Dives — Rogue and Warlock", Blizzard (read 2026-11-04, 200): https://news.blizzard.com/en-us/article/24310968/world-of-warcraft-forever-class-deep-dives-rogue-and-warlock
- "World of Warcraft: Forever Deep Dive Panel Recap", Blizzard (200): https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
- "Skyborne First Look: New Neutral Race in World of Warcraft: Forever", Wowhead (200): https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829
- "Weapon skill", Warcraft Wiki (used for the weapon-skill decision only): https://warcraft.wiki.gg/wiki/Weapon_skill
- foreversixty.gg/planner, /planner?class=rogue and /sim/weights (all 200)
- Client data (this tree), in `data/builds/1.60.1.70755/`:
  - `spelltips/rogue.json`: 8676, 11269, 53, 25351, 11340, 14186 and the rest listed above
  - `bis/rogue-subtlety.json`: band 60 main hand 1728
  - `items.json`: 1728, inventory_type 21
  - `talents/rogue.json`, `trainables/rogue.json`, `addon-builds.json`, `races.json`, `items/rogue.json`, `simconsumes.json`
- Render order: `web/src/pages/guides/[class]/[spec].astro:337-366`
- Round 3 race runs: `design/reviews/guide-accuracy/round3/race-sims.md:18-23`

## Fix log (round 7)

Client text checked first (`spelltips/rogue.json`, `talents/rogue.json`): Ambush 8676 to 11269 and Backstab 53 say "Requires a dagger in the main hand"; Cheap Shot 1833 "Must be stealthed. Awards 2 combo points"; Garrote has no weapon clause; Initiative 3/3 adds a point to Ambush, Garrote and Cheap Shot; Dirty Deeds removes Garrote's behind-the-target rule; Opportunity covers Backstab, Garrote, Ambush and Mutilate.

- ROGUE-R7-1 (CONFIRMED): fixed.
  - subtlety.md:62: Opportunity now "adds 20% to Garrote, and to Ambush and Backstab when you hold a dagger"; the Improved Ambush sentence now says Ambush needs a main-hand dagger and is not a boss opener even with one.
  - subtlety.md:70: "Ambush needs a dagger in the main hand, so with a sword it is off your bar. On trash that will die fast, open with Cheap Shot, which Initiative makes three combo points and which holds the mob still for 4 seconds, or with Garrote." Boss opener (Premeditation, Garrote) unchanged. Ghostly Strike (:58) and Initiative (:60) lines left: they state the mechanic correctly.
  - The top-level `notes` in rogue-subtlety.json mention Ambush only as history of an earlier action change; no card note claims Ambush is cast. Left.
- ROGUE-R7-2 (LIKELY, verified by Seal Fate 14186 and the 12/34/5 build): fixed, combat.md:100, auditor's wording.
- ROGUE-R7-3 (LIKELY, verified by the render order): fixed, assassination.md:115, "your best in slot table shows which of them your set wants".
- ROGUE-R7-4 (LIKELY, verified by Deadly Poison V 25351 and Instant Poison VI 11340 charges and Improved Poisons 5/5): fixed, assassination.md:88, auditor's wording.
- New stat rule: every rogue Stat priority section (assassination.md:99, combat.md:100, subtlety.md:97) already gives reasons only and points to the table; no other line ranks stats or gear by nightly value. Prose test `-t rogue` passes before and after the edits.
- `updated:` is already 2026-11-04 on all three spec files; index.md untouched.
- Decisions applied: no build, race, `statPriority` or card-action changes. Weapon-skill note not added (no client data on rogue starting skills).
- SIM findings: none new.
