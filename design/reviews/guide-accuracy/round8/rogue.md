# Guide accuracy round 8: Rogue

- Class: rogue. I audited every line of `web/src/content/guides/rogue/index.md`, `assassination.md`, `combat.md` and `subtlety.md` in the round 8 audit tree (land/guides-r7 09b98d9e4, all seven rounds applied). I also read every card note and prepull action in `data/curated/apl/rogue-{assassination,combat,subtlety}.json`.
- Engine pin: d1d9f2775 (`sim/enginever/version.go`). I read `sim/rogue/death_mark.go`, `sim/rogue/mutilate.go` and `sim/core/consumes.go` (Thistle Tea) at that sha.
- No rogue data, APL or guide-rendering file changed between round 7's tree (394562f20) and this one. The only rogue diffs are round 7's four guide edits. I re-checked each of those edits.
- Read first: the rogue reports for rounds 1 to 7 with their fix logs, `round4/rogue-levels.md`, the race-sims files for rounds 2 to 5 (only round 3 has rogue rows) and `ledger.md`.
- Not re-filed:
  - every carried sim item: SIM-1 to SIM-10, SIM-R2-1 to SIM-R2-3, SIM-R3-1, SIM-R3-2, SIM-R5-1, SIM-R6-1 and the Mutilate off-hand engine change
  - card actions: Subtlety's Shank and Ghostly Strike cast with no combo-point condition, and Slice and Dice shadowed by Eviscerate (SIM-3)
  - the Combat leveling order (ladder-first)
  - component-lane and data-lane items

## Decoded builds (re-verified, unchanged)

| Spec | build | raidBuild |
|---|---|---|
| Assassination | 38/11/2 (night-elf) | 36/10/5 |
| Combat | 14/34/3 (gnome) | 12/34/5 |
| Subtlety | 12/8/31 (night-elf), single build | none |

- Every swap the prose describes recounts exactly:
  - **Assassination:** 7 for 7. Vigor 1, Blood Splatter, the third point of Improved Slice and Dice and Opportunity 2 to 5 are traded for Improved Gouge 3, Remorseless Attacks 2, Precision 3 and Master of Deception 1.
  - **Combat:** 11 for 11.
- The tier tolls hold:
  - Assassination: Coup de Grace 3 plus Lightning Reflexes 2 is 5.
  - Combat raid build: tiers 0 and 1 add up to exactly 15.
  - Subtlety: Improved Ambush is the 25th point before Boot Knife.
- Prerequisites are all met. Mutilate needs Venom, Lethality needs Malice 5, Weapon Expertise needs Precision 3, Adrenaline Rush needs Blade Flurry, Hemorrhage needs Serrated Blades 3, and Death Mark needs Preparation.
- The leveling orders are identical to round 7's level-by-level decode. Every talent and level in the "Levels 1 to 19" and "Levels 20 to 60" prose matches them:
  - **Assassination:** Improved Gouge 10-12, Remorseless Attacks 13-14, Malice 15-19, Venom 30, Relentless Strikes 39 and 41-44, Mutilate 40.
  - **Combat:** Coup de Grace 10-12, Efficient Strikes 13, Lightning Reflexes 14-18, Puncturing Wounds 19-21, Aggression 28-30, Improved Gouge 44-46, Remorseless Attacks 47-48.
  - **Subtlety:** Master of Deception 10-12, Opportunity 13-17, Dirty Deeds 18, Camouflage 19, Premeditation 20, Ghostly Strike 26, Hemorrhage 30, Boot Knife 35-39, Death Mark 40.

## What checks out (fresh pass)

**Round 7's edits**
- subtlety.md:62 is right. Opportunity 14057 covers Backstab, Garrote, Ambush and Mutilate, and Ambush and Backstab need a main-hand dagger.
- subtlety.md:70 is right:
  - Ambush 8676/11269 says "Requires a dagger in the main hand".
  - Cheap Shot 1833 says "Stuns the target for 4 sec ... Awards 2 combo points", and Initiative 3/3 adds one more at 100%, so three.
  - Garrote has no weapon clause.
- combat.md:100 is right. Seal Fate is not in 12/34/5, and Garrote's 18 seconds is not "all fight".
- assassination.md:88 is right. Deadly Poison V has 240 charges and Instant Poison VI has 250, both for 30 minutes, and Improved Poisons gives a 50% chance to keep the charge.
- assassination.md:115 is right. The best in slot table holds the Hearts.
- The race lines' "non-periodic damaging abilities" for Eureka! is right.

**Talent values against the rank texts in `talents/rogue.json`**
- Every quoted value matches its rank text:
  - **Assassination tree:** Improved Gouge 1.5 s; Remorseless Attacks 40%; Malice 5%; Vigor 5 a point; Blood Splatter 6 s and 9 s; Lethality +20% on Mutilate, Backstab and Hemorrhage, not Ambush; Vile Poisons 20%; Improved Poisons 10% and 50%; Relentless Strikes 20% a point for 25; Venom 30%, 20% and 6 s; Seal Fate 100%; Cold Blood on Mutilate and Eviscerate.
  - **Combat tree:** Coup de Grace 20% on Eviscerate and Death Mark, rate and damage. Efficient Strikes 3 and 5. Lightning Reflexes 5%. Puncturing Wounds 30%, 15% and 45%. Precision 3%. Weapon Expertise 2%. Aggression 15% on Sinister Strike and Eviscerate only; the client beats the deep dive's "also Backstab". Flawless Execution 15; the client beats the deep dive's 10. Dual Wield Specialization 50%. Blade Flurry 20% and one extra target. Hack and Slash 5% at five points. Flash of Steel 15 and 2.0 s, so Shank gives 40 every 8 s. Adrenaline Rush 100%.
  - **Subtlety tree:** Master of Deception 3 levels; Opportunity 20%; Dirty Deeds rank 1 at 10 energy plus Garrote from the front; Camouflage rank 1; Initiative 100%; Premeditation 20 s with no Stealth clause; Serrated Blades 9% and 30%; Elusiveness 90 s; Ghostly Strike 180% with a main-hand dagger, 15% dodge, 7 s; Hemorrhage 145% with a dagger, 15% on Rupture and Garrote; Boot Knife 2 points; Death Mark 12% a point, bleeds at twice the rate, 6 s, cap carrying `$mcrit`.
- The engine's `death_mark.go` applies Coup de Grace to both the storage rate and the cap. This matches "stores 20% more and its cap rises 20%".

**Spells against `trainables/rogue.json` and `spelltips/rogue.json`**
- Training levels:
  - Sinister Strike, Eviscerate, Stealth and Redirect at 1.
  - Backstab 4, Gouge 6, Evasion 8, Slice and Dice and Sprint 10, Kick 12, Garrote 14, Ambush and Poised 18.
  - Instant Poison 20, Premeditation 20, Shank 24 (rank 2 at 44), Cheap Shot 26, Deadly Poison 30, Hemorrhage and Venom 30.
  - Mutilate and Death Mark 40. Kidney Shot rank 2 at 50.
- Costs and cooldowns:
  - Shank: 5 energy, 10 s, 15 or 25 back, cannot miss, off-hand 50%, no weapon-type clause.
  - Mutilate: 60 energy, 79.2 each hand at rank 3, 20% against poisoned targets.
  - Garrote 50 energy (40 with Poised). Backstab and Ambush 60. Ghostly Strike 40 energy on a 20 s cooldown. Hemorrhage 35 (30 with Efficient Strikes). Death Mark 25 energy on a 10 s cooldown.
  - Kidney Shot 25 energy, 20 s, 6 s at five points. Kick 25. Blade Flurry 25 energy, 2 min.
  - Adrenaline Rush 5 min. Cold Blood 3 min. Premeditation 2 min. Preparation 10 min.
  - Blind 5 min (3.5 min with Elusiveness). Vanish 5 min. Feint 10 s. Evasion 50% for 15 s.
  - Poised: 10 at full energy out of combat. Redirect: 15 s, the next non-player enemy hit.
  - Thistle Tea: 100 energy on a 5 minute cooldown (engine `consumes.go:1397`).
- Pull arithmetic holds:
  - Combat: Garrote at 40 plus Backstab at 60 is 100.
  - Subtlety: Premeditation 2 + Garrote 1 + Initiative 1 = 4, and the Preparation reset is Premeditation 2 + Shank 2 + Ghostly Strike 1 = 5.
  - Combat on a 3 minute kill: two Blade Flurries and one Adrenaline Rush.
- Windfury Totem 10613/10614 grants an extra attack to party members. It is not a weapon imbue, so "an aura on you" holds.

**Races, gear, professions, sources, voice**
- Racials match `races.json`. Skyborne pairs are races 95 and 96, so "none of the six new pairs is a rogue" holds.
- Every race order and "within a hair" claim matches `round3/race-sims.md:18-23`. Skyborne is never first among free and paid in any rogue row except as the paid option.
- Items checked:
  - Worn Dagger 2092 and Small Throwing Knife 2947 exist.
  - Heartseeker 12783 is a dagger, binds on equip and is crafted by Blacksmithing.
  - Frozen and Molten Heart of the Mountain (249469/249470) bind on pickup.
  - Teebu's Blazing Longsword 1728 is a main-hand sword. Devilsaur and Stormshroud pieces bind on equip.
- Source links: all six returned 200 on 2026-11-04, and each follows the round 4 source rule.
  - The deep dive still backs: Expose Armor not stacking with Sunder (Sunder will not overwrite it), one-handed axes, Mutilate on two daggers for 60 energy, poisons beside stones and more charges.
- Card notes state their own actions truthfully. Ghostly Strike "hits harder than Hemorrhage" holds on a sword too: 125% against 100%.
- No em or en dashes, no site, sim or APL talk, and no quoted nightly numbers. Stat priority sections give reasons only and point to the table above.

## Findings

| id | status | kind | file:line | wrong claim | evidence | exact fix |
|---|---|---|---|---|---|---|
| ROGUE-R8-1 | LIKELY (minor) | leveling (spell level) | combat.md:132 | "Open with Cheap Shot or Garrote from stealth, ..." (first sentences of "Levels 20 to 60") | **Cheap Shot trains at 26** (`trainables/rogue.json`: Cheap Shot rank 0 level 26). The block gives a level for everything else it introduces: Shank at 24, Instant Poison at 20, Deadly Poison at 30. So a reader at 20 to 25 is told to open with a spell they cannot have. The 1 to 19 block correctly never names it. | "Open with Garrote from stealth, or with Cheap Shot once you train it at 26, Slice and Dice at one or two points on anything that will live long enough, ..." |
| ROGUE-R8-2 | LIKELY (minor) | leveling (spell level) | subtlety.md:129 | "Shank comes at 24 for every rogue: press it on cooldown. Open with Cheap Shot or Garrote, ..." | Same data line: Cheap Shot trains at 26. The sentence comes straight after "Shank comes at 24", so it reads as available from 24. | "Shank comes at 24 for every rogue: press it on cooldown. Open with Garrote, or Cheap Shot from 26, keep Slice and Dice up on anything that lives, and Eviscerate at five." |

No CONFIRMED finding. I re-derived every number, level, talent, item and race claim above from the 70755 data and found none false. Both findings are imprecise level wording in a block that spans 20 to 60, so they are LIKELY rather than hard errors.

### Shoulds (wording, not findings)
- combat.md:104: "The off hand only needs to be fast and hit hard." These two pull against each other: Shank's 50% weapon damage favours a slow, hard-hitting off hand, and Instant Poison procs favour a fast one. "The off hand carries Instant Poison and every Shank, so it wants to hit hard" (as subtlety.md:101 says) would be cleaner.

## Decisions

- No build, race, `statPriority` or card-action changes. The `statPriority` lists still match the band 60 tables, which are unchanged since round 7.
- Round 7's four edits are correct and stay.
- **Not filed: Garrote over Ambush as the Assassination boss opener.** The page gives true mechanical reasons: armor and Lethality's ability list. The relative value was decided in earlier rounds.
- **Not filed: Aggression and Flawless Execution against the deep dive.** The deep dive's "Aggression now also increases Backstab damage" and "Flawless Execution ... by 10" are older than the 70755 rank texts, which say Sinister Strike and Eviscerate, and 15. The guide follows the client, which is the rule.

## Sims requested

None new. Carried for the sim round:
- SIM-R6-1 (Combat crit scale factor)
- the Mutilate off-hand penalty
- SIM-1 to SIM-5, SIM-7, SIM-8 and SIM-10
- SIM-R2-1 to SIM-R2-3
- SIM-R3-1 and SIM-R3-2
- SIM-R5-1

## Sim findings

None new on d1d9f2775.
- `death_mark.go` models Coup de Grace on both the storage rate and the cap, the same as the prose.
- SIM-3 is still open: the Subtlety Slice and Dice line (priority 6) sits behind Eviscerate (priority 5) with the same combo-point gate.
- The Subtlety APL casts Shank and Ghostly Strike with no combo-point condition, while the prose holds Shank at four or more points and Ghostly Strike at five. This is a card action against researched play, which is accepted until the sim round.

## Sources

- "World of Warcraft: Forever Class Deep Dives - Rogue and Warlock", Blizzard (read 2026-11-04, 200): https://news.blizzard.com/en-us/article/24310968/world-of-warcraft-forever-class-deep-dives-rogue-and-warlock
- "World of Warcraft: Forever Deep Dive Panel Recap", Blizzard (200): https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
- "Skyborne First Look: New Neutral Race in World of Warcraft: Forever", Wowhead (200): https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829
- foreversixty.gg/planner, /planner?class=rogue and /sim/weights (all 200)
- Client data (this tree), in `data/builds/1.60.1.70755/`:
  - `trainables/rogue.json`: Cheap Shot level 26
  - `talents/rogue.json`, `spelltips/rogue.json`, `spelltips/shaman.json` (10613, 10614)
  - `items/rogue.json`, `items.json`, `races.json`, `combos.json`
  - `bis/rogue-*.md`, `addon-builds.json`
- Engine d1d9f2775: `sim/rogue/death_mark.go:20-62`, `sim/core/consumes.go:1397-1425`
- Round 3 race runs: `design/reviews/guide-accuracy/round3/race-sims.md:18-23`

## Fix log (round 8)

- ROGUE-R8-1 fixed, combat.md:132: "Open with Garrote from stealth, or with Cheap Shot once you train it at 26" (trainables/rogue.json 1833 level 26).
- ROGUE-R8-2 fixed, subtlety.md:129: "Open with Garrote, or Cheap Shot from 26".
- Decisions unchanged. Prose check: rogue clean. No new SIM findings.
