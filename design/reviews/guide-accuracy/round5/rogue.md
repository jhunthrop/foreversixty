# Guide accuracy round 5: Rogue

- Class: rogue. I audited every line of `web/src/content/guides/rogue/index.md`, `assassination.md`, `combat.md` and `subtlety.md` in the round 5 audit tree (land/guides-r4 with the round 4 fixes: a901bae80, 77bd7c12b, 752c02e43). I also read the rendered card notes (`rotation.priorityList[*].notes`) in `data/curated/apl/rogue-*.json`.
- Client 1.60.1.70755 data: `talents/rogue.json` (every rank text), `trainables/rogue.json`, `spelltips/rogue.json`, `races.json`, `items.json`, `addon-builds.json` and `stat-weights.json`. Engine pin d1d9f2775 (`sim/enginever/version.go`).
- Read first: the rogue reports for rounds 1 to 4 (including the round 4 fix log), `round4/rogue-levels.md`, `round2/race-sims.md`, `round3/race-sims.md`, `round4/race-sims.md` (no rogue rows) and `ledger.md`, including the round 4 source-link rule.
- Not re-filed:
  - SIM-1 to SIM-10, SIM-R2-1 to SIM-R2-3, SIM-R3-1 and SIM-R3-2
  - the Mutilate off-hand engine change
  - the spelltips regen (R2-6, data lane)
  - card actions
  - the Combat leveling order (ladder-first)

## Decoded builds (re-verified, unchanged)

| Spec | build | raidBuild |
|---|---|---|
| Assassination | 38/11/2 | 36/10/5 |
| Combat | 14/34/3 (race field `gnome`) | 12/34/5 (race field `gnome`) |
| Subtlety | 12/8/31 (single build) | none |

Every build is legal on 70755 and equals the final state of its `addon-builds.json` order. Talent by talent:

- **Assassination `build`:**
  - Assassination: Improved Gouge 3, Remorseless Attacks 2, Malice 5, Ruthlessness 3, Improved Slice and Dice 2, Lethality 2, Vile Poisons 5, Cold Blood 1, Improved Poisons 5, Relentless Strikes 5, Venom 1, Seal Fate 3, Mutilate 1
  - Combat: Coup de Grace 3, Lightning Reflexes 2, Puncturing Wounds 3, Precision 3
  - Subtlety: Master of Deception 1, Opportunity 1
- **Assassination `raidBuild`:**
  - Assassination: Malice 5, Ruthlessness 3, Improved Slice and Dice 3, Vigor 1, Blood Splatter 1, Lethality 2, Vile Poisons 5, Cold Blood 1, Improved Poisons 5, Relentless Strikes 5, Venom 1, Seal Fate 3, Mutilate 1
  - Combat: Coup de Grace 3, Lightning Reflexes 2, Puncturing Wounds 3, Precision 2
  - Subtlety: Opportunity 5
- **Combat `build`:**
  - Assassination: Improved Gouge 3, Remorseless Attacks 2, Malice 5, Ruthlessness 3, Improved Slice and Dice 1
  - Combat: Coup de Grace 3, Efficient Strikes 1, Lightning Reflexes 5, Puncturing Wounds 3, Precision 3, Flawless Execution 1, Weapon Expertise 2, Aggression 3, Dual Wield Specialization 2, Blade Flurry 1, Hack and Slash 5, Flash of Steel 4, Adrenaline Rush 1
  - Subtlety: Master of Deception 3
- **Combat `raidBuild`:**
  - Assassination: Malice 5, Ruthlessness 2, Improved Slice and Dice 3, Lethality 2
  - Combat: Coup de Grace 3, Efficient Strikes 2, Lightning Reflexes 4, Puncturing Wounds 3, Precision 3, Flawless Execution 1, Weapon Expertise 2, Aggression 3, Dual Wield Specialization 2, Blade Flurry 1, Hack and Slash 4, Flash of Steel 5, Adrenaline Rush 1
  - Subtlety: Opportunity 5
- **Subtlety `build`:**
  - Assassination: Malice 5, Ruthlessness 3, Improved Slice and Dice 2, Blood Splatter 1, Lethality 1
  - Combat: Coup de Grace 3, Efficient Strikes 2, Precision 3
  - Subtlety: Master of Deception 3, Opportunity 5, Dirty Deeds 1, Camouflage 3, Improved Ambush 1, Initiative 3, Premeditation 1, Serrated Blades 3, Elusiveness 2, Ghostly Strike 1, Preparation 1, Hemorrhage 1, Boot Knife 5, Death Mark 1

Shipped leveling orders, re-derived from `addon-builds.json` (index 0 = level 10). They are identical to round 4's:

- **Assassination:** Improved Gouge 10-12, Remorseless Attacks 13-14, Malice 15-19, Ruthlessness 20-22, Improved Slice and Dice 23-24, Lethality 25-26, Vile Poisons 27-29 and 31-32, Venom 30, Cold Blood 33, Improved Poisons 34-38, Relentless Strikes 39 and 41-44, Mutilate 40, Seal Fate 45-47, Combat 48-58, Master of Deception 59, Opportunity 60.
- **Combat:** Coup de Grace 10-12, Efficient Strikes 13, Lightning Reflexes 14-18, Puncturing Wounds 19-21, Precision 22-24, Flawless Execution 25, Weapon Expertise 26-27, Aggression 28-30, Dual Wield Specialization 31-32, Blade Flurry 33, Hack and Slash 34-38, Flash of Steel 39 and 41-43, Adrenaline Rush 40, Improved Gouge 44-46, Remorseless Attacks 47-48, Malice 49-53, Ruthlessness 54-56, Improved Slice and Dice 57, Master of Deception 58-60.
- **Subtlety:** Master of Deception 10-12, Opportunity 13-17, Dirty Deeds 18, Camouflage 19, 24 and 27, Premeditation 20, Serrated Blades 21-23, Elusiveness 25 and 33, Ghostly Strike 26, Improved Ambush 28, Initiative 29 and 31-32, Hemorrhage 30, Preparation 34, Boot Knife 35-39, Death Mark 40, Assassination 41-52, Combat 53-60.

## What checks out (fresh pass)

**"Levels 1 to 19" paragraphs**

- Training levels match `trainables/rogue.json`: Sinister Strike, Eviscerate and Stealth at 1; Backstab 4; Gouge 6; Evasion 8; Sap, Slice and Dice and Sprint 10; Kick 12; Garrote 14; Ambush and Poised 18.
- Ambush rank 1 costs 60, and Poised 1328825 reads "reduce the Energy cost of your next ability by 10". Subtlety's "Ambush from stealth costs 50" is right.
- Talent ranks at the levels named:
  - Improved Gouge 3/3 is 1.5 s at 12.
  - Remorseless Attacks is 20% at 13 and 40% at 14 ("next Sinister Strike, Backstab, Ambush, Mutilate, or Ghostly Strike").
  - Malice 5/5 is 5%.
  - Coup de Grace 3/3 is 20%.
  - Efficient Strikes 1/2 is 3 energy.
  - Lightning Reflexes 5/5 is 5% dodge.
  - Puncturing Wounds names only Backstab and Mutilate, which matches the round 4 wording.
  - Master of Deception 3/3 is "3 levels higher".
  - Opportunity 5/5 is 20% on Backstab, Garrote, Ambush and Mutilate.
  - Dirty Deeds rank 1 is 10 energy, and it already drops the behind requirement at one point.
  - Camouflage rank 1 is 5% speed and 2 s off the cooldown.
- Every row gate is met at the levels the prose names:
  - Subtlety reaches 10 points at 19 (Premeditation at 20), 16 at 25 (Ghostly Strike at 26) and 20 at 29 (Hemorrhage at 30).
  - Assassination reaches 20 at 29 (Venom at 30).
  - Combat reaches 9 at 18 (Puncturing Wounds at 19).
- Worn Dagger, Small Throwing Knife and Thistle Tea exist in `items.json`.
- No block names Cheap Shot (26) or a poison (20) before 20. No block contradicts its "Levels 20 to 60" text or the strip.

**Talent and spell values elsewhere on the pages**

- Every value matches its 70755 rank or spell text:
  - Mutilate: 79 a hand at the top rank, 20% against poisoned targets.
  - Venom: 30%/20%/6 s.
  - Relentless Strikes: 20% a point for 25 energy (four in five at four points).
  - Lethality: 200% to 220%.
  - Vile Poisons: 20%. Improved Poisons: 10% and a 50% chance to keep the charge.
  - Vigor: 5 at one point.
  - Blood Splatter: 6 s and 9 s.
  - Coup de Grace: 20%.
  - Flash of Steel: 15 and 2.0 s at 5/5.
  - Hack and Slash: 5% at five points.
  - Dual Wield Specialization: 50%.
  - Puncturing Wounds: 30%/45%.
  - Aggression: 15% to Sinister Strike and Eviscerate.
  - Flawless Execution: 15.
  - Serrated Blades: 9%/30%.
  - Hemorrhage: 145% with a dagger, 15% to Rupture and Garrote.
  - Ghostly Strike: 180%, 15% dodge, 7 s.
  - Boot Knife: 2 points at 100%.
  - Death Mark: 12% a point, bleeds at double rate, 6 s.
  - Initiative: 100%.
  - Premeditation: 2 points, 20 s, no Stealth gate.
  - Preparation: "finishes the cooldown on your other Rogue abilities".
  - Elusiveness 2/2: 90 s off Vanish and Blind.
- Spell values match `spelltips/rogue.json`, and these do not move with level:
  - Kidney Shot rank 2: 25 energy, 20 s, 6 s at five points.
  - Kick: 25 energy.
  - Garrote: 50 energy, 1 point, stealthed and behind.
  - Cheap Shot: 2 points.
  - Shank rank 2: 5 energy for 25, 10 s, cannot miss.
  - Redirect: 10 s cooldown, stored 15 s.
  - Sprint, Vanish and Blind cooldowns.
- Build summaries recounted: Assassination trades 7 points for 7 and Combat trades 11 for 11. Combat's raid row-3 gate is exactly 15.
- Pull arithmetic still holds:
  - Combat: Garrote with Poised is 40, plus a 60 Backstab, which fits a 100 energy bar.
  - Subtlety: the opener is four points (Premeditation 2, Garrote 1, Initiative 1). The Preparation reset is five points (Premeditation 2, Boot Knife Shank 2, Ghostly Strike 1).

**Races, stats, sources, card notes and voice**

- Racial text matches `races.json` 70755: Touch of the Grave (5% of maximum health, once per 15 s), Blood Fury, Berserking (10% attack speed, 15 s, 3 min), Elune's Light, Ruthless, Artisan's Precision, Eureka!, Expansive Mind (5% energy), Will of the Forsaken and Wind Blessed. Skyborne is named only as the paid option. The race orderings in the prose match `round3/race-sims.md:18-23`, and the round 4 Combat runners-up line is right.
- `statPriority` is unchanged and still matches the band-60 scale factors.
- Source links: every frontmatter and body link resolves (200), and each one follows the round 4 rule (no /sources, /sim/weights on damage specs only, the planner for talents). The Blizzard deep dive (24310968, "World of Warcraft: Forever Class Deep Dives — Rogue and Warlock") backs everything the pages lean on it for:
  - Mutilate on two daggers for 60 energy.
  - Expose Armor equals 5 Sunders, and Sunder will not overwrite it unless weaker.
  - Poisons go on beside stones.
  - One-handed axes.
  - Premeditation lasts 20 s.
  - Hack and Slash combines the weapon specializations.
  - Assassination finishers: Rupture and Eviscerate are a lower priority.
- The panel recap (24303313) is correctly labelled. The Wowhead Skyborne page's title matches; its body renders client-side.
- Card notes: every rendered note states its own action truthfully.
- Voice: no em or en dashes; no site, sim, APL or band talk; no quoted nightly numbers.

## Findings

| id | status | kind | file:line | wrong claim | evidence | exact fix |
|---|---|---|---|---|---|---|
| ROGUE-R5-1 | CONFIRMED (minor; a round 1 edit) | talent | subtlety.md:62 | "The point in Improved Ambush only opens the next row; Ambush is not a boss opener for you." | **Row 3 does not need it.** Row 3 (tier 2: Initiative, Premeditation, Serrated Blades) needs 10 points. Rows 1 and 2 hold 12 without Improved Ambush (Master of Deception 3, Opportunity 5, Dirty Deeds 1, Camouflage 3).<br>**Boot Knife's row does.** In the final build, tiers 0-4 hold 8 + 5 + 7 + 3 + 2 = 25, exactly the 25 that Boot Knife's tier 5 needs. Without the Improved Ambush point they hold 24, and Boot Knife and Death Mark are illegal.<br>**The leveling order agrees.** That point at 28 is what brings the tree to 20 by 29, so Hemorrhage can come at 30.<br>**Where the error came from.** Round 1's note ("needed to open row 3") was wrong, and the round 1 fix carried it into the prose. | "The point in Improved Ambush is there for the points total that opens Boot Knife and Death Mark; Ambush is not a boss opener for you." |
| ROGUE-R5-2 | LIKELY (minor) | leveling (consistency) | assassination.md:131 vs :121, :127 | "Level Combat or Assassination with Sinister Strike until 40." | **The 1 to 19 block contradicts it.** The block just above builds the leveling pull on "Gouge, step behind and Backstab" and says to keep a dagger "so Backstab and Ambush stay open to you". The shipped Assassination leveling build backs that block: Improved Gouge 10-12, then Remorseless Attacks, which also covers Backstab and Ambush.<br>**Wrong guide.** "Combat or Assassination" reads as if the page might send you down another tree. This is the Assassination guide, and its build is the Assassination order. | "From 20 to 40, Sinister Strike is your everyday builder on mobs that face you, with Gouge into Backstab whenever you can get behind." |
| ROGUE-R5-3 | LIKELY (minor) | wording | combat.md:35 | "Adrenaline Rush doubles your energy for 15 seconds" | Adrenaline Rush 13750: "Increases your Energy regeneration rate by 100% for 15 sec." It doubles regeneration, not the bar. The same page gets it right at :54 ("double energy regeneration"). | "Adrenaline Rush doubles your energy regeneration for 15 seconds, and Blade Flurry cleaves every swing onto a second target" |
| ROGUE-R5-4 | LIKELY (minor) | voice | subtlety.md:90 | "Elusiveness puts Blind on a short leash." | "On a short leash" means held back, the opposite of what is meant. Elusiveness 2/2 takes 90 s off Blind's 5 minute cooldown (`spelltips` 2094: 300000 ms), so it comes back every 3.5 minutes. | "Kick, Gouge and Blind handle the rest; with Elusiveness, Blind comes back every three and a half minutes." |

## Decisions

- No build, race or card-action changes. The race picks follow `round3/race-sims.md` and were not re-measured; no rogue spec changed builds since that run.
- Round 4's fixes were re-checked and all are right. That covers R4-1 (Blizzard original cited), R4-2 (the Combat runners-up as measured), R4-3 (Puncturing Wounds wording) and R4-4 (a better dagger), plus the source relabels.
- The Venom description in the Blizzard deep dive ("new 25 Energy Finisher") predates the 70755 hotfix that made Venom a passive on Shank. The guides follow the 70755 rank text, which is correct. The source is still the right citation for the claims the pages draw from it.

## Sims requested

None new. Carried for the sim round:
- the Mutilate off-hand penalty
- SIM-1 to SIM-5, SIM-7, SIM-8 and SIM-10
- SIM-R2-1 to SIM-R2-3
- SIM-R3-1 (Gnome Assassination Eureka! masks)
- SIM-R3-2 (Combat Thistle Tea inside Adrenaline Rush)

## Sim findings

None new on d1d9f2775.

## Sources

- "World of Warcraft: Forever Class Deep Dives — Rogue and Warlock", Blizzard (read 2026-11-03): https://news.blizzard.com/en-us/article/24310968/world-of-warcraft-forever-class-deep-dives-rogue-and-warlock
- "World of Warcraft: Forever Deep Dive Panel Recap", Blizzard: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
- "Skyborne First Look: New Neutral Race in World of Warcraft: Forever", Wowhead (title only; the body is client-rendered): https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829
- Round 3 race runs: `design/reviews/guide-accuracy/round3/race-sims.md` (this tree)
- Client data: `data/builds/1.60.1.70755/talents/rogue.json` (Improved Ambush, Dirty Deeds, Camouflage, Remorseless Attacks, Elusiveness, Adrenaline Rush rank texts and tiers), `trainables/rogue.json`, `spelltips/rogue.json` (2094 Blind, 703 Garrote, 1833 Cheap Shot, 8643 Kidney Shot, 1766 Kick, 1329549 Shank, 1328825 Poised), `addon-builds.json`

## Fix log (round 5)

- ROGUE-R5-1 (CONFIRMED): fixed, `subtlety.md:62`. Verified in `talents/rogue.json`: Improved Ambush is tier 1, Boot Knife tier 5; the build's tiers 0-4 hold exactly 25 only with that point. The line now says the point is there for the total that opens Boot Knife and Death Mark.
- ROGUE-R5-2 (LIKELY, verified against the shipped Assassination order and the 1 to 19 block): fixed, `assassination.md:131`. It now reads "From 20 to 40, Sinister Strike is your everyday builder on mobs that face you, with Gouge into Backstab whenever you can get behind."
- ROGUE-R5-3 (LIKELY, verified: Adrenaline Rush rank text "Increases your Energy regeneration rate by 100% for 15 sec."): fixed, `combat.md:35` ("doubles your energy regeneration").
- ROGUE-R5-4 (LIKELY, verified: Blind `cooldown_ms` 300000 in `spelltips` 2094, Elusiveness 2/2 "by 90 sec"): fixed, `subtlety.md:90` ("with Elusiveness, Blind comes back every three and a half minutes").
- Decisions: applied as written. No build, race, card-action or card-note changes, because no note repeated any of the four claims. No em or en dashes were added.
- New SIM findings: none.

## Fix log (round 5, review pass)

Evidence check before the weapon rewrites:
- Client 70755: Hemorrhage 16511 "100% weapon damage (145% if a Dagger is equipped)"; Ghostly Strike 14278 "125% (180% if a Dagger is equipped in your Main Hand)" (`spelltips/rogue.json`, `talents/rogue.json`).
- Engine pin `sim/rogue/client_values.go:97` `mainHandStrikePct` applies the dagger figure when `HasDagger(core.MainHand)`, used by `hemorrhage.go:64` and `ghostly_strike.go:74`. The engine models the dagger bonus, so the Subtlety table's sword main hand and mace off hand stand, and the prose follows the table. The Subtlety build has no Hack and Slash, so the sword's win comes from white swings and the weapon itself; Rupture, Garrote, Death Mark and Shank do not check weapon type.

- MUST 1 (index all three trees on daggers): fixed. `index.md:4` description, `index.md:22` intro and `index.md:26` 1.12 bullet now put Combat and Assassination on daggers (both band 60 tables carry a dagger main hand on both factions) and say Subtlety takes the best weapons it finds, dagger or not. The "swords, axes, maces and fists are leveling weapons" clause is gone, since Subtlety raids with a sword and a mace.
- MUST 2 (subtlety.md:101 "Daggers in both hands"): fixed, `subtlety.md:101`. No pieces named; it keeps the stable why (the dagger bonus is on two builders, the bulk of the damage ignores weapon type), sends the player to the table, and the off hand line now says it carries Instant Poison and every Shank (Shank 1329549 strikes with the off-hand weapon).
- MUST 3 (subtlety.md:97 hit prose vs "Hit cap: reached (special attacks), 28% (white swings)"): fixed, `subtlety.md:97`. Specials are at the cap with level 60 gear and gain nothing more; hit stays first for white swings. Order follows the table: Hit, Agility, Critical strike, Strength, Attack power, Melee haste. No numbers.
- SHOULD (combat.md:116 Heartseeker): fixed, `combat.md:116`. Heartseeker 12783 is a dagger with `bonding: equip` (`items/rogue.json`), so the line now says it binds when equipped and to buy one when it is on your list. It no longer implies it is in the table at every preset (raid off hand is Sentinel's Blade).

New SIM finding:
- SIM-R5-1 (LIKELY, minor): client Hemorrhage text says "if a Dagger is equipped" with no hand named, while the engine gives the 145% only to a main-hand dagger (`client_values.go:97`). For the Subtlety raid set (no dagger in either hand) this changes nothing; it matters only for a sword main hand with a dagger off hand. Needs an in-game check of the effect's item requirement before the engine changes.
