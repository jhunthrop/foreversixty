# Guide accuracy round 12: Rogue

Audited on the round 12 audit tree (main a8b147a6f: guide rounds 1 to 11, sim round 15, leveling-order, guide-header-build), client 1.60.1.70755, engine pin 9f7f1f6e4 (engine source read at `scratchpad/gr10/engine`). Files: `web/src/content/guides/rogue/{index,assassination,combat,subtlety}.md`, every rogue card note in `data/builds/1.60.1.70755/addon-data.json` (`rotations.rogue-*`, all seven bands), and the curated APLs `data/curated/apl/rogue-{assassination,combat,subtlety}.json`. I read the rogue reports for rounds 1 to 11 with their fix logs, `round4/rogue-levels.md`, `round3/race-sims.md` (the only race file with rogue rows) and `ledger.md`.

What changed since round 11's tree:
- `git diff 91da94f1e HEAD -- web/src/content/guides/rogue data/curated/apl data/builds/1.60.1.70755/addon-builds.json data/builds/1.60.1.70755/bis/` is empty. No rogue prose, card note, point order or BiS file moved after the round 11 fix commit.
- The only change that reaches these pages is guide-header-build (3bd24d3a2 and its merges):
  - The header build card shows `raidBuild` where a guide has one, titled "Raid build" (Assassination and Combat).
  - The rotation cards, cooldown cuts and pre-pull timings read that build, under a "For the raid build." caption.
  - Key-talent chips show both builds' ranks where they differ.
- Subtlety has no `raidBuild`, so its header still shows `build`.

## Builds decoded (re-verified, all legal on 70755)

Decoded positionally against `talents/rogue.json` tab order (`fs1.ts`).

- Assassination `build` 38/11/2 (night-elf): Improved Gouge 3, Remorseless Attacks 2, Malice 5, Ruthlessness 3, Improved Slice and Dice 2, Lethality 2, Vile Poisons 5, Cold Blood 1, Improved Poisons 5, Relentless Strikes 5, Venom 1, Seal Fate 3, Mutilate 1 / Coup de Grace 3, Lightning Reflexes 2, Puncturing Wounds 3, Precision 3 / Master of Deception 1, Opportunity 1.
- Assassination `raidBuild` 36/10/5 (header card): Malice 5, Ruthlessness 3, Improved Slice and Dice 3, Vigor 1, Blood Splatter 1, Lethality 2, Vile Poisons 5, Cold Blood 1, Improved Poisons 5, Relentless Strikes 5, Venom 1, Seal Fate 3, Mutilate 1 / Coup de Grace 3, Lightning Reflexes 2, Puncturing Wounds 3, Precision 2 / Opportunity 5.
- Combat `build` 14/34/3 (gnome): Improved Gouge 3, Remorseless Attacks 2, Malice 5, Ruthlessness 3, Improved Slice and Dice 1 / Coup de Grace 3, Efficient Strikes 1, Lightning Reflexes 5, Puncturing Wounds 3, Precision 3, Flawless Execution 1, Weapon Expertise 2, Aggression 3, Dual Wield Specialization 2, Blade Flurry 1, Hack and Slash 5, Flash of Steel 4, Adrenaline Rush 1 / Master of Deception 3.
- Combat `raidBuild` 12/34/5 (header card): Malice 5, Ruthlessness 2, Improved Slice and Dice 3, Lethality 2 / Coup de Grace 3, Lightning Reflexes 4, Puncturing Wounds 3, Precision 3, Composure 2, Flawless Execution 1, Weapon Expertise 2, Aggression 3, Dual Wield Specialization 2, Blade Flurry 1, Hack and Slash 4, Flash of Steel 5, Adrenaline Rush 1 / Opportunity 5.
- Subtlety `build` 12/8/31 (night-elf, single build): Malice 5, Ruthlessness 3, Improved Slice and Dice 2, Blood Splatter 1, Lethality 1 / Coup de Grace 3, Efficient Strikes 2, Precision 3 / Master of Deception 3, Opportunity 5, Dirty Deeds 1, Camouflage 3, Improved Ambush 1, Initiative 3, Premeditation 1, Serrated Blades 3, Elusiveness 2, Ghostly Strike 1, Preparation 1, Hemorrhage 1, Boot Knife 5, Death Mark 1.

Prerequisites hold: Lethality behind Malice 5, Mutilate behind Venom, Weapon Expertise behind Precision 3, Adrenaline Rush behind Blade Flurry, Hemorrhage behind Serrated Blades 3, Death Mark behind Preparation. Every swap paragraph recounts exactly (Assassination 7 for 7, Combat 12 for 12).

Shipped orders (`addon-builds.json`, index 0 = level 10), re-derived:
- Assassination: Improved Gouge 10-12, Remorseless Attacks 13-14, Malice 15-19, Ruthlessness 20-22, Improved Slice and Dice 23-24, Lethality 25-26, Vile Poisons 27-29 and 31-32, Venom 30, Cold Blood 33, Improved Poisons 34-38, Relentless Strikes 39 and 41-44, Mutilate 40, Seal Fate 45-47, then Combat 48-58, Master of Deception 59, Opportunity 60.
- Combat: Coup de Grace 10-12, Efficient Strikes 13, Lightning Reflexes 14-18, Puncturing Wounds 19-21, Precision 22-24, Flawless Execution 25, Weapon Expertise 26-27, Aggression 28-30, Dual Wield Specialization 31-32, Blade Flurry 33, Hack and Slash 34-38, Flash of Steel 39 and 41-43, Adrenaline Rush 40, Improved Gouge 44-46, Remorseless Attacks 47-48, then Assassination and Master of Deception 49-60.
- Subtlety: Master of Deception 10-12, Opportunity 13-17, Dirty Deeds 18, Camouflage 19, Premeditation 20, Serrated Blades 21-23, Ghostly Strike 26, Hemorrhage 30, Preparation 34, Boot Knife 35-39, Death Mark 40, then Assassination and Combat 41-60.

Every level the prose names matches.

## Findings

| id | status | kind | file:line | quote | evidence | fix |
|---|---|---|---|---|---|---|
| ROGUE-R12-1 | LIKELY | talent (round 11 decision left half applied) | assassination.md:78; Shank card note, `rogue-assassination.json` / `addon-data.json` `rotations.rogue-assassination` bands 30 to 60 | ":78 with Venom, for the next 6 seconds your poisons do 30% more damage" / card: "with Venom, makes your poisons hit 30% harder and land more often for 6 seconds." | See the evidence below the table. | :78: "...and with Venom, for the next 6 seconds your poisons get another 30% on top of Vile Poisons and Instant Poison lands on half your off-hand hits." Card: "Shank on cooldown: it cannot miss, pays back more energy than it costs and, with Venom, adds 30% to your poison damage on top of Vile Poisons and makes them land more often for 6 seconds." |

### Evidence for ROGUE-R12-1

**Round 11's own finding and decision.**
- ROGUE-R11-1 quoted ":56 every Shank makes your poisons hit 30% harder" as the error.
- Its Decisions section says: "Damage is a percent add that sums with Vile Poisons (aura 108, same misc and masks)", so poisons go from 120% to 150% of base, "25% harder than without Venom, not 30%".
- The fix rewrote :56 to "adds 30% to your poison damage, on top of Vile Poisons", which is right under that model.
- The same fix wrote ":78 your poisons do 30% more damage", which is the wording R11-1 called wrong.
- Round 11 also kept the Shank card note. It said the note "gives no number and stays true", but the note does say "hit 30% harder", the same phrase R11-1 quoted from :56.

**Client and engine.**
- `spellconst/rogue.json` 1310703 (Venom's buff): e0 aura 108 misc 0, amount 30; e1 aura 108 misc 22, amount 30.
- `spellconst/rogue.json` 16513 (Vile Poisons 5/5): e0 aura 108 misc 0, amount 20; e1 aura 108 misc 22, amount 20. These are the same modifier kind and op.
- Engine `sim/rogue/venom.go:10-22`: "The aura 108 modifiers share Vile Poisons' (16513) kind, misc value and masks, so the two add, never multiply". `venomDamageScale()` returns (vile + 0.30) / vile, which is 1.25 at Vile Poisons 5/5.

So under the model the site measures, :78 and the card overstate Venom's gain on a page that already has Vile Poisons.

**Why LIKELY, not CONFIRMED.** Whether two percent spell modifiers of the same op add or multiply on Forever's server is not settled by a client row. Additive is the vanilla rule; multiplicative is the modern engine's rule. If they multiply, "30% more" is exact. The fix above is true under both rules, costs nothing, and removes the contradiction with :56 and with round 11's decision.

## Header-build check (round 12 focus)

**Assassination** (header card and rotation cards are now 36/10/5):
- :52 opens on "The raid build goes 36/10/5" and names the level 60 build's trades from it, so the prose agrees with the card.
- Every key-talent bullet (:54 to :62) has the same rank in both builds, so no chip shows two ranks the prose leaves out.
- Raid-only lines are labelled: :78 "In the raid build, Rupture next" and :80 "Without Blood Splatter, in the level 60 build, skip Rupture too". They match the card's Rupture gate, `auraIsKnown(1329256)`.

**Combat** (header card and rotation cards are now 12/34/5):
- :52 opens on "The raid build goes 12/34/5".
- :56 "Hack and Slash (5 points, 4 in the raid build)" and :57 "Flash of Steel (5 points, 4 in the level 60 build)" carry both ranks, as the chips now do.
- The Shank card ("every 8 with five points of Flash of Steel") and :77 ("gives 40 back with Flash of Steel") read the raid build's five points: 25 + 15.
- :63 and :73 tie Composure and the 20-energy Poised opener to the raid build. Garrote 11290 costs 50, so it costs 30 in the raid build.

**Subtlety:** there is no `raidBuild`, so nothing moved.

No prose on any of the four pages calls the header card the level 60 build or describes it as the build you level with.

## Shoulds (wording, not findings)

| id | file:line | note | suggested wording |
|---|---|---|---|
| ROGUE-R12-S1 | subtlety.md:129 | "Eviscerate at three while Rupture ticks on anything bigger" assumes a Rupture that the 20 to 40 text never tells you to put up. Rupture trains at 20, and the band 20 to 38 cards open Rupture at four or more. | "...keep Slice and Dice up on anything that lives, and Eviscerate at five on mobs. On anything bigger, Rupture at four and Eviscerate at three while it ticks, until Death Mark comes at 40." |
| ROGUE-R12-S2 | subtlety.md:70 | The sentence "Then **Rupture** right away at four points, and **Death Mark**..." follows the trash Cheap Shot sentence, so it reads as trash advice. :129 says to Eviscerate a mob that dies before the mark goes off. | Move the trash sentence to the end of the paragraph: "...Then **Rupture** right away at four points, and **Death Mark** as soon as you have five with both bleeds ticking. On trash that will die fast, open with **Cheap Shot**, which Initiative makes three combo points and which holds the mob still for 4 seconds, and Eviscerate." |

## Checked and right (re-derived this round)

**Spells** (`trainables/rogue.json`, `spelltips/rogue.json`, `spellconst/rogue.json`):
- Every training level named holds: Sinister Strike, Eviscerate, Stealth and Redirect at 1, Backstab 4, Gouge 6, Evasion 8, Slice and Dice and Sprint 10, Kick 12, Garrote 14, Ambush and Poised 18, Instant Poison and Rupture 20, Shank 24 and 44, Cheap Shot 26, Deadly Poison and Kidney Shot 30, Kidney Shot rank 2 at 50.
- Mutilate 1241584: "90% weapon damage plus an additional 79.2 with each weapon", 20% against a poisoned target, 60 energy, 2 combo points.
- Shank 1329549: 5 energy, 25 back, 10 s, "cannot Miss", off-hand 50% weapon damage, no combo point of its own.
- Poised 1328825: 10 off the next ability at maximum energy out of combat. Composure 2/2 adds 10.
- Ambush costs 60, so 50 with Poised. Garrote 50. Cheap Shot 60, "Awards 2 combo points", 4 s stun. Ghostly Strike 40 energy on a 20 s cooldown. Hemorrhage 35, so 30 with Efficient Strikes 2/2.
- Death Mark: 25 energy, 10 s cooldown, 6 s. Engine `death_mark.go:20-27` applies Coup de Grace to both the storage rate and the cap, as subtlety.md:57 says.
- Premeditation 14183: 2 points, 20 s, 2 min cooldown, with no stealth line (Cheap Shot's tooltip has "Must be stealthed"; Premeditation's does not).
- Preparation: 10 min. Blind: 5 min, so 3.5 min with Elusiveness 2/2. Blade Flurry: 25 energy, 2 min. Adrenaline Rush: 5 min. Cold Blood: 3 min, off the global, no cost.
- Kidney Shot: 25 energy, 20 s, 6 s at five points. Kick: 25. Redirect: 15 s, next enemy hit.

**Talents.** Every rank text the prose quotes matches `talents/rogue.json`:
- Lethality on the crit bonus: 220%. Ambush is not on Lethality's list, so Ambush "gets none of Lethality's bonus" holds.
- Puncturing Wounds has no crit condition. Seal Fate is 100%. Relentless Strikes is 20% a point for 25 energy.
- Venom's chance is a flat 20 points (aura 107 misc 18, like Improved Poisons). 20 + 10 + 20 = 50%, so "half your off-hand hits" holds.
- Blood Splatter is 6 s on Rupture and 9 s on Garrote. Master of Deception is three levels. Camouflage rank 1 is 5%. Dirty Deeds rank 1 is 10 energy and Garrote from the front.
- Boot Knife is 100% at 5/5. Initiative is 100% at 3/3. Ghostly Strike is 180% with a dagger, 15% dodge, 7 s. Hemorrhage is 145% and 15%.

**Rotation.**
- Opener: all three APLs cast Stealth at -3 s and Garrote 11290 at 0 s. Subtlety adds Premeditation at -2 s, which matches "a couple of seconds before the pull".
- The card lines and gates match the prose on all three pages, including Subtlety's Eviscerate note as fixed in round 11.

**Stats.**
- `statPriority` matches the band 60 raid scale factors in the three `bis/rogue-*.json` files, haste set aside:
  - Assassination: hit, agility, crit, strength, attack power.
  - Combat: hit, agility, strength, attack power, crit.
  - Subtlety: hit, agility, crit, strength, attack power (both factions).
- Each stat paragraph explains and never ranks.

**Races.**
- Every line matches `round3/race-sims.md:18-23`.
- No Skyborne pick is recommended; Skyborne appears only as the paid Heroic Pack option.
- `combos.json` has both Skyborne rogues and no rogue among the six new pairs.

**Gear and professions.** Unchanged since round 9, and they hold.

**Sources.** All six links returned 200 on 2026-11-04: the two Blizzard articles, the Wowhead Skyborne article, /planner, /planner?class=rogue and /sim/weights. Each follows the source-link rule.

**Voice.** No em or en dashes, no site, sim or APL talk, and no quoted nightly numbers.

## Decisions

- **Venom's damage wording** is filed LIKELY, not CONFIRMED. The client rows do not settle whether same-op percent modifiers add or multiply on Forever's server. The suggested text is true either way.
- **Race lines stand.** No race line contradicts a measured result. Race picks on the round 15 builds are still unmeasured, as the round 11 notes say.
- **No build, order, `statPriority` or BiS change.**

## Sims requested

Carried from rounds 10 and 11 and still open in this tree:
1. Subtlety raid 60: the shipped APL against the APL with the old Eviscerate line.
2. Subtlety race re-run, both factions, on the Death Mark rotation.
3. Assassination race re-run with the Rupture line.

PROVISIONAL (BiS rows wait for the final re-rank):
4. The `bis/rogue-combat.json` band 60 raid rows were ranked on talents `005203002-3243030000012321451-005`. That build has Efficient Strikes 2 and no Composure. The shipped `raidBuild` is `005203002/3043030002012321451/005`: Composure 2, no Efficient Strikes. The stat table order still matches `statPriority`. Re-rank Combat raid on the shipped `raidBuild`. Assassination's and Subtlety's BiS strings match their builds.

No new sim request.

## Sim findings

Carried, not re-filed:
- **SIM-R11-1.** Subtlety's Eviscerate (priorityList[6]) and Slice and Dice (priorityList[7]) still require `dotIsActive(Rupture)`, so band 10 never spends combo points.
- **Component lane.** Talent spells with trainable level 0 (Ghostly Strike, Cold Blood) still show on the band 10 and 20 cards.

No new engine finding.

## Sources

- "World of Warcraft: Forever Class Deep Dives - Rogue and Warlock", Blizzard (200, 2026-11-04): https://news.blizzard.com/en-us/article/24310968/world-of-warcraft-forever-class-deep-dives-rogue-and-warlock
- "World of Warcraft: Forever Deep Dive Panel Recap", Blizzard (200): https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
- "Skyborne First Look: New Neutral Race in World of Warcraft: Forever", Wowhead (200): https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829
- foreversixty.gg/planner, /planner?class=rogue and /sim/weights (all 200)
- Client data, in `data/builds/1.60.1.70755/`:
  - `talents/rogue.json`, `trainables/rogue.json`, `addon-builds.json`, `combos.json` and `races.json`
  - `spellconst/rogue.json`: 1310703, 16513, 14113, 1329328, 1241584, 1241586, 1329549, 1328825, 1328826, 14183, plus the cost and cooldown rows listed above
  - `spelltips/rogue.json`: 1241584, 1329283, 1329549, 1328825, 1316709, 1325758, 438040, 14183, 14185, 14278, 16511, 13877, 13750, 8643, 1833, 14177
  - `addon-data.json` (`rotations.rogue-*`) and `bis/rogue-*.json`
- APLs: `data/curated/apl/rogue-{assassination,combat,subtlety}.json`
- Engine 9f7f1f6e4: `sim/rogue/venom.go:10-40`, `poisons.go:24-42`, `death_mark.go:10-62`
- Race runs: `design/reviews/guide-accuracy/round3/race-sims.md:18-23`
- Previous rounds: `round11/rogue.md` (ROGUE-R11-1 and Decisions), `ledger.md`

## Counts

CONFIRMED 0, LIKELY 1, NEEDS-SIM 0, NEEDS-OWNER 0. SHOULDs 2. New sim findings: 0. One PROVISIONAL BiS re-rank request.

Verdict: zero confirmed for rogue. The one LIKELY item is round 11's Venom fix left half applied: :78 and the Shank card still say the "30% harder" that R11-1 called wrong.

## Fix log (round 12)

- ROGUE-R12-1: fixed, assassination.md:78 ("your poisons get another 30% on top of Vile Poisons") and the Shank card note in data/curated/apl/rogue-assassination.json ("adds 30% to your poison damage on top of Vile Poisons and makes them land more often for 6 seconds"). Verified spellconst/rogue.json 1310703 (aura 108 misc 0 and 22, 30) and 16513 (aura 108 misc 0 and 22, 20). Wording holds whether the modifiers add or multiply. addon-data.json is generated and not touched.
- ROGUE-R12-S1: applied, subtlety.md:129 ("Eviscerate at five on mobs. On anything bigger, Rupture at four and Eviscerate at three while it ticks, until Death Mark comes at 40."). Verified Rupture trains at 20 (trainables) and the subtlety Rupture card fires at four or more.
- ROGUE-R12-S2: applied, subtlety.md:70: the trash Cheap Shot sentence now closes the paragraph, ending "and Eviscerate", after the Rupture and Death Mark sentence.
