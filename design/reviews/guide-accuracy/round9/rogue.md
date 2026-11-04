# Guide accuracy round 9: Rogue

- Class: rogue. I audited every line of `web/src/content/guides/rogue/index.md`, `assassination.md`, `combat.md` and `subtlety.md` in the round 9 audit tree (land/guides-r8 645caa983, all eight rounds applied). I also read every card note and the top-level notes in `data/curated/apl/rogue-{assassination,combat,subtlety}.json`.
- Engine pin: d1d9f2775 (`sim/enginever/version.go`). I read `sim/rogue/shank.go` (Boot Knife) at that sha.
- No rogue data, APL or guide-rendering file changed after round 8, except round 8's two leveling edits (combat.md:132, subtlety.md:129). I re-checked both.
- Read first: the rogue reports for rounds 1 to 8 with their fix logs, `round4/rogue-levels.md`, the race-sims files for rounds 2 to 5 (only round 3 has rogue rows) and `ledger.md`.
- Not re-filed:
  - every carried sim item: SIM-1 to SIM-10, SIM-R2-1 to SIM-R2-3, SIM-R3-1, SIM-R3-2, SIM-R5-1, SIM-R6-1 and the Mutilate off-hand engine change
  - card actions: Subtlety's Shank and Ghostly Strike cast with no combo-point condition, and Slice and Dice shadowed by Eviscerate (SIM-3)
  - the Combat leveling order (ladder-first)
  - component-lane and data-lane items

## Decoded builds (re-verified, unchanged)

Builds were decoded positionally against `talents/rogue.json` tab order (`fs1.ts`).

| Spec | build | raidBuild |
|---|---|---|
| Assassination | 38/11/2 (night-elf) | 36/10/5 |
| Combat | 14/34/3 (gnome) | 12/34/5 |
| Subtlety | 12/8/31 (night-elf), single build | none |

**Checks that hold:**
- Every swap the prose describes recounts exactly:
  - Assassination is 7 for 7.
  - Combat is 11 for 11, tree by tree: Assassination +6/-4, Combat +2/-2 and Subtlety 3 for 5.
- Tier gates hold on all five strings:
  - Combat raid build: tiers 0 and 1 add up to exactly 15.
  - Subtlety: tiers 0 to 4 hold 25 only with the Improved Ambush point.
- Prerequisites hold: Mutilate needs Venom, Lethality needs Malice, Weapon Expertise needs Precision, Adrenaline Rush needs Blade Flurry, Hemorrhage needs Serrated Blades and Death Mark needs Preparation.

**Leveling orders** (`addon-builds.json`, index 0 = level 10) match every level named in the "Levels 1 to 19" and "Levels 20 to 60" prose:

| Spec | Order |
|---|---|
| Assassination | Improved Gouge 10-12, Remorseless Attacks 13-14, Malice 15-19, Venom 30, Relentless Strikes 39 and 41-44, Mutilate 40 |
| Combat | Coup de Grace 10-12, Efficient Strikes 13, Lightning Reflexes 14-18, Puncturing Wounds 19-21, Aggression 28-30, Flash of Steel 39 and 41-43, Adrenaline Rush 40, Improved Gouge 44-46, Remorseless Attacks 47-48 |
| Subtlety | Master of Deception 10-12, Opportunity 13-17, Dirty Deeds 18, Camouflage 19, Premeditation 20, Ghostly Strike 26, Hemorrhage 30, Boot Knife 35-39, Death Mark 40 |

## Round 9 focus checks

**Critical strike damage.** I applied the rounds 7 and 8 lesson: a "critical strike damage bonus" talent multiplies the +100% bonus.
- Lethality's rank 2 text: "Increases the critical strike damage bonus of your Sinister Strike, Gouge, Backstab, Mutilate, Ghostly Strike, and Hemorrhage abilities by 20%". So a crit is 100% + 100% × 1.20 = 220% of a normal hit.
- Every page line on this holds:
  - assassination.md:59: "Mutilate crits deal 220% instead of 200%"
  - combat.md:63: "takes Backstab crits from 200% to 220%"
  - combat.md:100: "Lethality makes your Backstab crits bigger still"
  - assassination.md:99: "Lethality makes it bigger"
  - assassination.md:72 and the Combat APL notes say Ambush gets none of Lethality, which is right.
- No other rogue talent or racial on 70755 changes crit damage. Cold Blood (assassination.md:36, :58) raises crit chance by 100%, so "a guaranteed crit" holds.
- combat.md:63's "more than half your Backstabs crit" is plausible. Puncturing Wounds gives +30%, Malice 5%, Hack and Slash 4/5 gives 4%, and agility and gear add more, well above boss crit suppression.

**Eureka!** The rogue row, Wowhead Forever 1259812, reads: "Your non-periodic damaging abilities have their Energy cost reduced by 10% and deal 10% more damage for 8 sec", on a 2 minute cooldown. `spelltips/shared.json` has the same text. Both Eureka! lines use exactly that limit:
- index.md:37
- combat.md:112

**Premeditation without Stealth** is re-proved on the tooltip header, not only on the description:
- Classic 14183 carries a "Requires Stealth" line.
- Wowhead Forever 14183 has none.
- Cheap Shot 1833 keeps "Requires Stealth" on both.

So subtlety.md:59 and :70 ("needs no Stealth in Forever") hold.

## What checks out (fresh pass)

**Round 8's edits**
- combat.md:132 is right: "Open with Garrote from stealth, or with Cheap Shot once you train it at 26". `trainables/rogue.json` lists Cheap Shot at level 26.
- subtlety.md:129 is right: "Open with Garrote, or Cheap Shot from 26".
- assassination.md:131 names Cheap Shot only after "Once Mutilate comes at 40", so it needs no level.

**Talent values against `talents/rogue.json` rank texts**

All re-read:
- **Assassination tree:** Improved Gouge 1.5 s; Remorseless Attacks 40%; Malice 5%; Ruthlessness 60%; Improved Slice and Dice; Vigor 5 a point; Blood Splatter 6 s and 9 s; Vile Poisons 20%; Improved Poisons 10% and 50%; Relentless Strikes 20% a point for 25 (four in five at four points); Venom 30%, 20% and 6 s; Seal Fate 100%; Mutilate 90% plus 79.2 at rank 3, 20% against poisoned targets.
- **Combat tree:** Coup de Grace 20% on Eviscerate and Death Mark (rate and damage); Efficient Strikes 3 and 5; Lightning Reflexes 5%; Puncturing Wounds 30%, 15% and 45%; Precision 3%; Flawless Execution 15; Weapon Expertise 2%; Aggression 15% on Sinister Strike and Eviscerate; Dual Wield Specialization 50%; Blade Flurry 20% and one extra target; Hack and Slash 5% crit for daggers and fists, 5% extra attack for swords and axes, 15% armor for maces; Flash of Steel 15 and 2.0 s; Adrenaline Rush 100%.
- **Subtlety tree:** Master of Deception 3 levels; Opportunity 20% on Backstab, Garrote, Ambush and Mutilate; Dirty Deeds rank 1 at 10 energy plus Garrote from the front; Camouflage 5% at rank 1; Initiative 100%; Premeditation 2 points and 20 s; Serrated Blades 9% and 30%; Elusiveness 90 s; Ghostly Strike 125%/180%, 15% dodge, 7 s, Ambush or Garrote without Stealth; Hemorrhage 100%/145% and 15%; Boot Knife 100% for 2 points; Death Mark 12% per point, twice for bleeds, 6 s, the cap scaled by `$mcrit` and Coup de Grace.

**Spells against `trainables/rogue.json` and `spelltips/rogue.json`**
- Training levels:
  - Sinister Strike, Eviscerate, Stealth and Redirect at 1.
  - Backstab 4, Gouge 6, Evasion 8, Slice and Dice and Sprint 10, Kick 12, Garrote 14, Ambush and Poised 18.
  - Instant Poison 20, Premeditation 20, Shank 24 (rank 2 at 44), Cheap Shot 26, Deadly Poison 30, Hemorrhage and Venom 30.
  - Mutilate and Death Mark 40. Kidney Shot rank 2 at 50 (6 s at five points).
- Costs and cooldowns:
  - Shank: 5 energy, 10 s, 15 or 25 back, cannot miss.
  - Backstab and Ambush 60 energy, Garrote 50, Cheap Shot 60, Ghostly Strike 40 on a 20 s cooldown, Hemorrhage 35, Death Mark 25 on a 10 s cooldown.
  - Kidney Shot and Kick 25 energy.
  - Blade Flurry 25 energy on a 2 min cooldown. Adrenaline Rush 5 min. Cold Blood 3 min. Premeditation 2 min. Preparation 10 min.
  - Blind 5 min (3.5 min with Elusiveness). Vanish 5 min. Feint 10 s. Evasion 50% for 15 s.
  - Poised: 10 at full energy out of combat. Redirect: 15 s, the next non-player enemy hit.
  - Thistle Tea: the Wowhead Forever item 7676 tooltip says "Instantly restores 100 energy. (5 Min Cooldown)".
- Pull arithmetic holds:
  - Combat: Poised Garrote 40 plus Backstab 60 is 100.
  - Subtlety opener: Premeditation 2 + Garrote 1 + Initiative 1 = 4. Preparation reset: 2 + 2 + 1 = 5.
  - Two Blade Flurries and one Adrenaline Rush on a 3 minute kill.
  - Ambush from stealth costs 50 with Poised.

**Stat priority**
- `statPriority` matches the band 60 raid scale factors in `bis/rogue-*.json`, both factions. Haste is outside the normalization.
  - Assassination: Hit, Agility, Crit, Strength, Attack power.
  - Combat: Hit, Agility, Strength, Attack power, Crit.
  - Subtlety: Hit, Agility, Crit, Strength, Attack power.
- Every stat paragraph gives reasons only and points to the table. `haste_on_items` is false, so "Haste is rare on gear" holds.

**Races, gear, professions, sources, voice**
- Racials match the 70755 data and the Wowhead Forever rows:
  - Touch of the Grave (1260189): 20%, 15 s, 5% of maximum health
  - Blood Fury: 10% attack power, 15 s, 2 min
  - Berserking: 10%, 3 min
  - Elune's Light: 10% crit, 15 s, 3 min
  - Ruthless: 1% crit
  - Artisan's Precision: 1% hit
  - Will of the Forsaken: charm, fear and sleep, 2 min
  - Wind Blessed: 2% haste
  - Expansive Mind, rogue row 1259803: "Maximum Energy increased by 5%"
- Every race order and "within a hair" claim matches `round3/race-sims.md:18-23`. Skyborne is named only as the paid option. `combos.json` (races 95 and 96) allows a rogue, and none of the six new pairs is a rogue.
- Gear and profession claims are unchanged since round 8 and still hold: Worn Dagger, throwing knives, Heartseeker (Blacksmithing, binds on equip), Frozen and Molten Heart of the Mountain (bind on pickup, Enchanting), and the Devilsaur and Stormshroud leatherworking crafts.
- Source links: all six returned 200 on 2026-11-04 (the two Blizzard articles, the Wowhead Skyborne article, /planner, /planner?class=rogue, /sim/weights). I re-read the deep dive. It still backs:
  - Mutilate needs two daggers and triggers Seal Fate once a cast.
  - Expose Armor and Sunder do not stack.
  - One-handed axes.
  - Poisons sit beside stones and have more charges.
  - Hack and Slash folds the old weapon specializations into one talent.
- Card notes state their own actions truthfully. Ghostly Strike's "10 more energy" holds: 40 against 30 with Efficient Strikes 2/2.
- No em or en dashes, no site, sim or APL talk, and no quoted nightly numbers.

## Findings

No findings this round: no CONFIRMED, no LIKELY, no NEEDS-SIM and no NEEDS-OWNER.

I re-derived every number, level, talent, item, race and source claim on the four pages from the 70755 data, the Wowhead Forever tooltips and the deep dive. None is false, and none contradicts the page around it.

### Shoulds (wording, not findings)

| id | file:line | note | suggested wording |
|---|---|---|---|
| ROGUE-R9-S1 | index.md:37, combat.md:112 | "Expansive Mind gives you 5% more energy." The rogue row (1259803) is "Maximum Energy increased by 5%": a bigger bar, not faster regeneration. The current words are not false, but "more energy" can be read as regeneration. | "**Expansive Mind** makes your energy bar 5% bigger." |
| ROGUE-R9-S2 | assassination.md:80 | "you go from nothing to 2 or 3, then to 4 or 5". With Ruthlessness 3/3 (60%), most finishers leave a point behind, so the usual run is 1 to 3 or 4, then 5 or more. The conclusion (spend at four) is unchanged and right. | "Mutilate gives two points and a Seal Fate crit gives a third, and Ruthlessness often leaves you one after a finisher, so two Mutilates usually put you on four or five: waiting for exactly five wastes a Mutilate." |
| ROGUE-R9-S3 (carried R8 should) | combat.md:104 | "The off hand only needs to be fast and hit hard." Shank's 50% weapon damage and Dual Wield Specialization favour a hard-hitting off hand. | "The off hand carries Instant Poison and every Shank, so it wants to hit hard; Dual Wield Specialization makes it worth a real upgrade." |

## Decisions

- No build, race, `statPriority`, card-action or prose changes.
- **Boot Knife stays at two combo points.** Wowhead Forever's Shank 1329549 tooltip shows a bracketed "[Boot Knife: 100% Chance to award 1 Combo Point]". The sources that set the rule all say 2:
  - the Boot Knife talent text (`talents/rogue.json` and Wowhead 1329286: "add 2 Combo Points")
  - the client effect (`spellconst/rogue.json` 1329286 effect 1, amount 2)
  - the engine (`shank.go:55` `bootKnifeComboPoints = 2`, pinned by `TestBootKnifeAwardsTwoComboPoints`)

  The bracketed "1" looks like a variable-display quirk on Shank's tooltip, so subtlety.md:55, :70, :76, :82 and :86 stand.
- **Premeditation needs no Stealth.** This is re-confirmed on the tooltip header (Forever has no "Requires Stealth" line; Classic does), not inferred from the description.
- **Lethality crit lines stand.** A 20% increase to the 100% bonus is 220%, which is exactly what the pages say.
- **Not filed: the Mutilate off-hand split** ("the off-hand strike takes the usual half weapon damage, the flat 79 does not"). It is the carried engine and owner item, and the page does not contradict itself.

## Sims requested

None new. Carried for the sim round:
- SIM-R6-1 (Combat crit scale factor)
- the Mutilate off-hand penalty
- SIM-1 to SIM-5, SIM-7, SIM-8 and SIM-10
- SIM-R2-1 to SIM-R2-3
- SIM-R3-1 and SIM-R3-2
- SIM-R5-1

Race re-runs still wait on the final best-in-slot re-rank.

## Sim findings

None new on d1d9f2775.
- `shank.go` gives Boot Knife 2 combo points at 20% a rank, which agrees with the talent text.
- SIM-3 and the Subtlety Shank and Ghostly Strike gating are still card actions against researched play. These are accepted until the sim round.

## Sources

- "World of Warcraft: Forever Class Deep Dives - Rogue and Warlock", Blizzard (read 2026-11-04, 200): https://news.blizzard.com/en-us/article/24310968/world-of-warcraft-forever-class-deep-dives-rogue-and-warlock
- "World of Warcraft: Forever Deep Dive Panel Recap", Blizzard (200): https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
- "Skyborne First Look: New Neutral Race in World of Warcraft: Forever", Wowhead (200): https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829
- foreversixty.gg/planner, /planner?class=rogue and /sim/weights (all 200)
- Wowhead tooltips, read 2026-11-04:
  - Forever, https://nether.wowhead.com/forever/tooltip/spell/<id>: Premeditation 14183, Cheap Shot 1833, Ghostly Strike 14278, Mutilate 1241584, Shank 1329549, Death Mark 1316709 and 1325758, Poised 1328825, Redirect 438040, Hemorrhage 16511, Cold Blood 14177, Blade Flurry 13877, Preparation 14185, Venom 1329328, Boot Knife 1329286, Eureka! 1259812, Expansive Mind 1259803, Touch of the Grave 1260189
  - Forever item, Thistle Tea: https://nether.wowhead.com/forever/tooltip/item/7676
  - Classic, for comparison: https://nether.wowhead.com/classic/tooltip/spell/14183, /1833, /14278
- Client data (this tree), in `data/builds/1.60.1.70755/`: `talents/rogue.json`, `trainables/rogue.json`, `spelltips/rogue.json`, `spelltips/shared.json`, `spellconst/rogue.json`, `races.json`, `combos.json`, `addon-builds.json`, `bis/rogue-*.json`
- Engine d1d9f2775: `sim/rogue/shank.go:54-88`, `sim/rogue/dps_rogue/venom_test.go:77-79`
- Round 3 race runs: `design/reviews/guide-accuracy/round3/race-sims.md:18-23`

## Counts

CONFIRMED 0, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0. SHOULDs 3 (one carried from round 8). Verdict: zero confirmed errors on the rogue pages.
