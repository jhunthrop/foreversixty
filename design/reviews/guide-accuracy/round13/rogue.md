# Guide accuracy round 13: Rogue

Audited on the round 13 audit tree (main 604b78a80: guide rounds 1 to 12, sim round 15, guide-header-build), client 1.60.1.70755, engine pin 9f7f1f6e4 (engine source read at `scratchpad/gr10/engine`).

**Scope.**
- Pages: `web/src/content/guides/rogue/{index,assassination,combat,subtlety}.md`, every line.
- Card notes: every rogue note in `data/builds/1.60.1.70755/addon-data.json` (`rotations.rogue-*`, all seven bands).
- APLs: `data/curated/apl/rogue-{assassination,combat,subtlety}.json`.

**Read first:**
- the rogue reports for rounds 1 to 12, with their fix logs
- `round4/rogue-levels.md`
- `round3/race-sims.md`, the only race file with rogue rows
- `ledger.md`, through round 12

## What changed since round 12's tree

`git diff a8b147a6f HEAD` touches rogue only through the round 12 fix commit:
- `assassination.md:78`: Venom's 30% now reads "on top of Vile Poisons" (ROGUE-R12-1).
- The Shank note in `rogue-assassination.json` priorityList[1] and its generated copy in `addon-data.json`, with the same wording.
- `subtlety.md:70`: the trash Cheap Shot sentence moved to the end of the paragraph and now ends "then Eviscerate" (R12-S2).
- `subtlety.md:129`: "Eviscerate at five on mobs. On anything bigger, Rupture at four and Eviscerate at three while it ticks, until Death Mark comes at 40." (R12-S1).

Nothing else moved: no rogue build string, `addon-builds.json` row, `bis/rogue-*.json` or other APL line.

All three round 12 edits are re-checked below and hold.

## Builds decoded (re-verified, all legal on 70755)

Decoded positionally against `talents/rogue.json` tab order (`fs1.ts`).

- **Assassination `build` 38/11/2** (night-elf):
  - Assassination: Improved Gouge 3, Remorseless Attacks 2, Malice 5, Ruthlessness 3, Improved Slice and Dice 2, Lethality 2, Vile Poisons 5, Cold Blood 1, Improved Poisons 5, Relentless Strikes 5, Venom 1, Seal Fate 3, Mutilate 1
  - Combat: Coup de Grace 3, Lightning Reflexes 2, Puncturing Wounds 3, Precision 3
  - Subtlety: Master of Deception 1, Opportunity 1
- **Assassination `raidBuild` 36/10/5:**
  - Assassination: Malice 5, Ruthlessness 3, Improved Slice and Dice 3, Vigor 1, Blood Splatter 1, Lethality 2, Vile Poisons 5, Cold Blood 1, Improved Poisons 5, Relentless Strikes 5, Venom 1, Seal Fate 3, Mutilate 1
  - Combat: Coup de Grace 3, Lightning Reflexes 2, Puncturing Wounds 3, Precision 2
  - Subtlety: Opportunity 5
- **Combat `build` 14/34/3** (gnome):
  - Assassination: Improved Gouge 3, Remorseless Attacks 2, Malice 5, Ruthlessness 3, Improved Slice and Dice 1
  - Combat: Coup de Grace 3, Efficient Strikes 1, Lightning Reflexes 5, Puncturing Wounds 3, Precision 3, Flawless Execution 1, Weapon Expertise 2, Aggression 3, Dual Wield Specialization 2, Blade Flurry 1, Hack and Slash 5, Flash of Steel 4, Adrenaline Rush 1
  - Subtlety: Master of Deception 3
- **Combat `raidBuild` 12/34/5:**
  - Assassination: Malice 5, Ruthlessness 2, Improved Slice and Dice 3, Lethality 2
  - Combat: Coup de Grace 3, Lightning Reflexes 4, Puncturing Wounds 3, Precision 3, Composure 2, Flawless Execution 1, Weapon Expertise 2, Aggression 3, Dual Wield Specialization 2, Blade Flurry 1, Hack and Slash 4, Flash of Steel 5, Adrenaline Rush 1
  - Subtlety: Opportunity 5
- **Subtlety `build` 12/8/31** (night-elf, single build):
  - Assassination: Malice 5, Ruthlessness 3, Improved Slice and Dice 2, Blood Splatter 1, Lethality 1
  - Combat: Coup de Grace 3, Efficient Strikes 2, Precision 3
  - Subtlety: Master of Deception 3, Opportunity 5, Dirty Deeds 1, Camouflage 3, Improved Ambush 1, Initiative 3, Premeditation 1, Serrated Blades 3, Elusiveness 2, Ghostly Strike 1, Preparation 1, Hemorrhage 1, Boot Knife 5, Death Mark 1

**Tier gates hold:**
- Subtlety's Assassination tier 2: Malice 5, Ruthlessness 3 and Improved Slice and Dice 2 make 10.
- Subtlety's Combat tier 1: Coup de Grace 3 and Efficient Strikes 2 make 5.
- Assassination's Combat tier 1: Coup de Grace 3 and Lightning Reflexes 2 make 5.
- Combat raid tier 3: Coup de Grace 3, Lightning Reflexes 4, Puncturing Wounds 3, Precision 3 and Composure 2 make 15.

**The swap paragraphs recount exactly:**
- Assassination :52, 7 for 7: Vigor 1, Blood Splatter 1, Improved Slice and Dice 1 and Opportunity 4, against Improved Gouge 3, Remorseless Attacks 2, Precision 1 and Master of Deception 1.
- Combat :52, 12 for 12.

**Shipped orders** (`addon-builds.json`, index 0 = level 10), decoded again this round:

- Assassination: Improved Gouge 10-12, Remorseless Attacks 13-14, Malice 15-19, Ruthlessness 20-22, Improved Slice and Dice 23-24, Lethality 25-26, Vile Poisons 27-29 and 31-32, Venom 30, Cold Blood 33, Improved Poisons 34-38, Relentless Strikes 39 and 41-44, Mutilate 40, Seal Fate 45-47, Coup de Grace 48-50, Lightning Reflexes 51-52, Puncturing Wounds 53-55, Precision 56-58, Master of Deception 59, Opportunity 60.
- Combat: Coup de Grace 10-12, Efficient Strikes 13, Lightning Reflexes 14-18, Puncturing Wounds 19-21, Precision 22-24, Flawless Execution 25, Weapon Expertise 26-27, Aggression 28-30, Dual Wield Specialization 31-32, Blade Flurry 33, Hack and Slash 34-38, Flash of Steel 39 and 41-43, Adrenaline Rush 40, Improved Gouge 44-46, Remorseless Attacks 47-48, Malice 49-53, Ruthlessness 54-56, Improved Slice and Dice 57, Master of Deception 58-60.
- Subtlety: Master of Deception 10-12, Opportunity 13-17, Dirty Deeds 18, Camouflage 19, 24 and 27, Premeditation 20, Serrated Blades 21-23, Elusiveness 25 and 33, Ghostly Strike 26, Improved Ambush 28, Initiative 29 and 31-32, Hemorrhage 30, Preparation 34, Boot Knife 35-39, Death Mark 40, then Assassination and Combat 41-60.

Every level the prose names matches:
- Assassination: 10-12, 13 and 14, 15 to 19, Venom at 30, Mutilate at 40, Relentless Strikes "39 to 44".
- Combat: 10-12, 13, 14 to 18, 19, "Aggression by 30", "mid 40s" for 44-48.
- Subtlety: 10-12, 13 to 17, 18, 19, 20, 26, 30, "from 35", 40.

## Findings

None.

## Re-check of round 12's own edits

- **assassination.md:78** ("for the next 6 seconds your poisons get another 30% on top of Vile Poisons and Instant Poison lands on half your off-hand hits"):
  - Venom 1329328 reads "increase your Poison damage by 30% and your chance to apply Poisons by 20% for 6 sec".
  - Instant Poison VI 11340 has a 20% base chance, Improved Poisons 5/5 adds 10 and Venom adds 20, so it lands on 50% of hits.
  - The sentence agrees with :56 and is true whether the two damage modifiers add or multiply.
- **Shank card note** (all bands from 30): it matches :56 and :78. At bands 30 and 38 the card casts rank 1 (1329283: 15 energy), and "pays back more energy than it costs" still holds.
- **subtlety.md:70.**
  - The boss line, Garrote then Rupture at four then Death Mark at five, now reads as one sequence.
  - The trash line follows it. Cheap Shot 1833 "Awards 2 combo points" and Initiative 3/3 adds one, so the "three combo points" holds, as does "4 sec".
- **subtlety.md:129.**
  - Rupture trains at 20 (1943), which fits "Rupture at four".
  - "Eviscerate at three while it ticks" matches the band 20 to 38 Eviscerate card (three or more with Rupture ticking, before Death Mark).
  - "Death Mark comes at 40" matches the order (index 30) and trainables (1316709 at 40).
  - The paragraph no longer contradicts :70 or :76 to :78.

## Shoulds (wording, not findings)

| id | file:line | note | suggested wording |
|---|---|---|---|
| ROGUE-R13-S1 | subtlety.md:86 | "Death Mark is your only damage cooldown" reads loosely: Ghostly Strike (20 s) and Shank (10 s) are damage abilities on cooldowns too. The point the sentence makes, that nothing needs holding, is right. | "Death Mark is your burst, and it comes back every 10 seconds, so there is nothing to hold: ..." |

## Checked and right (re-derived this round)

**Spells** (`trainables/rogue.json`, `spelltips/rogue.json`, `spellconst/rogue.json`):
- **Training levels.**
  - From level 1: Sinister Strike 1752, Eviscerate 2098, Stealth 1784 and Redirect 438040.
  - Backstab 53 at 4, Gouge 1776 at 6, Evasion 5277 at 8, Slice and Dice 5171 and Sprint 2983 at 10, Kick 1766 at 12, Garrote 703 at 14, Ambush 8676 and Poised 1328825 at 18.
  - Instant Poison 8679 and Rupture 1943 at 20, Shank 1329283 at 24 and 1329549 at 44, Cheap Shot 1833 at 26, Deadly Poison 2823 and Kidney Shot 408 at 30, Kidney Shot 8643 at 50.
  - The talent spells: Premeditation 20, Hemorrhage 30, Venom 30, Mutilate 40, Death Mark 40.
- **Shank.** Rank 1 gives 15 energy, rank 2 gives 25, on a 10 s cooldown for 5 energy. Both deal 50% off-hand weapon damage and "cannot Miss", so index.md:30 "25 energy back at rank 2" holds.
- **Mutilate** 1241584: 90% weapon damage plus 79.2 with each weapon, 20% more against poisoned targets, 2 combo points. There is no facing line, so "no facing requirement" (assassination.md:34, :131) holds. The engine keeps the dagger-in-each-hand check (`mutilate.go`, ExtraCastCondition).
- **Garrote** 11290 costs 50 and reads "Must be stealthed and behind the target". That makes combat.md:63's Garrote at 30 with Composure 2/2 right. Subtlety's Garrote from the front comes from Dirty Deeds, whose rank 1 text carries the facing clause.
- **Poised, Ambush and Cheap Shot.** Poised takes 10 off the next ability. Ambush costs 60 (50 with Poised). Cheap Shot costs 60, is "Must be stealthed", stuns for 4 s and gives 2 combo points.
- **Builders.** Backstab and Ambush require a main-hand dagger. Ghostly Strike: 40 energy, 20 s, 180% with a dagger, 15% dodge for 7 s, and the next Ambush or Garrote needs no Stealth. Hemorrhage: 35 energy, 145% with a dagger, 15% more Rupture and Garrote damage.
- **Death Mark** (1316709 to 1325758): 25 energy, 10 s cooldown, 6 s duration. The engine stores 0.12 per combo point (`death_mark.go:38`), and Garrote and Rupture are stored at twice the rate.
- **Cooldowns and costs.**
  - Blade Flurry: 25 energy, 2 min, 20% speed, one extra target.
  - Adrenaline Rush: +100% for 15 s on a 5 min cooldown.
  - Cold Blood: 3 min, no cost. Its list includes Mutilate and Eviscerate.
  - Premeditation: 2 min, 20 s window, no Stealth line.
  - Preparation: 10 min, "your other Rogue abilities".
  - Blind: 5 min, so 3.5 min with Elusiveness 2/2 (90 s).
  - Kidney Shot: 25 energy, 20 s, 6 s at five points (rank 2).
  - Kick: 25. Feint: 10 s. Vanish and Evasion: 5 min.
  - Redirect: 15 s, next non-player enemy hit.
- **Poisons.** Deadly Poison V: 30% base. Instant Poison VI: 20% base.

**Talents.** Every rank text the prose quotes matches `talents/rogue.json`:
- Improved Gouge 3/3 is 1.5 s. Remorseless Attacks 2/2 is 40% after a non-trivial kill, on Sinister Strike, Backstab, Ambush, Mutilate and Ghostly Strike. Malice 5/5 is 5%.
- Ruthlessness 3/3 is 60%. Improved Slice and Dice is 15% a point. Vigor 1/2 is +5 energy. Blood Splatter is 6 s on Rupture and 9 s on Garrote.
- Lethality 2/2 is +20% on the crit bonus, which makes 220% on Sinister Strike, Gouge, Backstab, Mutilate, Ghostly Strike and Hemorrhage, not Ambush.
- Vile Poisons 5/5 is 20%. Improved Poisons 5/5 is +10% and a 50% chance to keep the charge. Relentless Strikes is 20% a point for 25 energy. Venom is 30% damage and 20 points of chance for 6 s. Seal Fate is 100%.
- Coup de Grace 3/3 is 20% on Eviscerate and on Death Mark's damage and storage. Efficient Strikes is 3, then 5, on Sinister Strike and Hemorrhage. Lightning Reflexes 5/5 is 5%.
- Puncturing Wounds 3/3 is 30% Backstab crit, 15% Mutilate crit and a 45% chance at a second combo point. Precision is 1% a point. Composure 2/2 is +10.
- Flawless Execution is 15. Weapon Expertise 2/2 is 2% dodge and parry. Aggression is 15% on Sinister Strike and Eviscerate only. Dual Wield Specialization 2/2 is 50%.
- Hack and Slash 5/5 is 5% crit with a dagger or fist, a 5% extra attack with a sword or axe, and 15% armor ignored with a mace. Flash of Steel 5/5 is +15 energy and 2.0 s off the cooldown.
- Master of Deception is three levels. Opportunity 5/5 is 20% on Backstab, Garrote, Ambush and Mutilate. Dirty Deeds 1/2 is 10 energy, with Garrote from the front. Camouflage 1/3 is a 5% smaller Stealth slow.
- Initiative 3/3 is 100%. Serrated Blades 3/3 is 9% and 30%. Elusiveness 2/2 is 90 s. Boot Knife 5/5 is 100% for 2 combo points.
- **Seal Fate once per Mutilate.** Blizzard's deep dive says it directly: "Mutilate will trigger this exactly once if either or both hands generates a critical strike". The engine matches (`mutilate.go`, a 500 ms cooldown on Seal Fate).

**Aggression and the deep dive.** The deep dive says Aggression "Now also increases Backstab damage". The 70755 rank text (18427) names only Sinister Strike and Eviscerate. Under the round 1 lesson, the client description wins, so combat.md:63 ("not to Backstab") stands. The client rows carry no family mask that would say otherwise.

**Rotation.**
- **Prepull.** All three APLs cast Stealth and then Garrote 11290; Subtlety adds Premeditation before the Garrote. That matches the three "The pull" paragraphs.
- **Card notes** (all seven bands) agree with the prose:
  - Assassination: Slice and Dice at four, Rupture only with Blood Splatter, Cold Blood, then Eviscerate at four.
  - Combat: Shank above Eviscerate at five, Slice and Dice at one or two, Backstab, Adrenaline Rush below Backstab. This fits "low on energy, never at the cap".
  - Subtlety: Shank at three or fewer, Ghostly Strike at four or fewer, Garrote back after Ghostly Strike, Rupture at four, Death Mark at five while Rupture ticks.
- **Combat's 3 minute kill.** It holds two Blade Flurries (2 min cooldown) and one Adrenaline Rush (5 min), as :87 says.
- **Subtlety's Preparation line.** Premeditation 2, Shank with Boot Knife 2 and Ghostly Strike 1 make five points, as :86 says.

**Stats.**
- `statPriority` matches the band 60 raid scale factors in the three `bis/rogue-*.json` files, haste set aside:
  - Assassination: hit 1.99, agility 1, crit 0.82, strength 0.59, attack power 0.53.
  - Combat: hit 2.39, agility 1, strength 0.74, attack power 0.67, crit 0.52.
  - Subtlety: hit, agility, crit, strength, attack power on both factions.
- Every stat paragraph explains and never ranks. Every reason is true on 70755:
  - Combat has no Seal Fate in either build.
  - Death Mark's cap carries `$mcrit`.
  - Bleeds crit (`forever-dots-crit`).

**Gear and professions.**
- Daggers in both hands for Combat and Assassination at band 60 raid:
  - Electrified Dagger or Glacial Blade in the main hand.
  - Off hand: Heartseeker for Assassination; Sentinel's Blade or Scout's Blade for Combat.
- Subtlety raids on Teebu's Blazing Longsword and Persuader. That fits "dagger or not", and the main-hand sword fits subtlety.md:70 ("with a sword it is off your bar").
- Molten, Frozen and Dormant Heart of the Mountain are bind on pickup.
- Heartseeker 12783 binds when equipped.
- Hearts appear in every band 60 trinket list, so assassination.md:115 "your best in slot table shows which of them your set wants" holds.

**Races.**
- Every line matches `round3/race-sims.md:18-23`:
  - Assassination: Night Elf, Dwarf and Human within 1.2. Undead, then Orc, then Troll.
  - Combat: Gnome and Undead at 759.3. Orc, Troll and Night Elf at 755.9 to 756.6.
  - Subtlety: Night Elf, then Human; Undead, then Orc.
- Racial text matches `races.json`:
  - Eureka! is the rogue row (non-periodic damaging abilities). Expansive Mind adds 5% energy.
  - Touch of the Grave: up to 5% of maximum health, once every 15 s.
  - Blood Fury: 10% attack power, 15 s, 2 min. Berserking: 10%, 15 s, 3 min.
  - Elune's Light: 10% crit, 15 s, 3 min. Ruthless: 1% crit. Artisan's Precision: 1% hit.
  - Will of the Forsaken: charm, fear and sleep, 2 min. Wind Blessed: 2% haste.
- Skyborne appears only as the paid Heroic Pack option and is never the pick. Every Skyborne row sits below the free leader on both factions.

**Sources** (fetched 2026-11-04). All six returned 200:
- the two Blizzard articles
- the Wowhead Skyborne article
- foreversixty.gg /sim/weights, /planner and /planner?class=rogue

The rogue deep dive still says what the pages cite:
- poisons beside stones, with more charges
- one-handed axes
- Expose Armor not stacking with Sunder ("Sunder Armor will still not overwrite Expose Armor")
- Seal Fate's single Mutilate proc
- Mutilate on two daggers for 60 energy

**Voice.** No em or en dashes. No site, sim or APL talk, and no quoted nightly numbers.

## Decisions

- **Aggression stays "not to Backstab".** The 70755 tooltip beats the pre-70755 deep dive line, under the round 1 lesson.
- **Race lines stand.** None contradicts a measured result. Re-runs on the round 15 builds are still with the sim lane.
- **No build, order, `statPriority` or BiS change.**

## Sims requested

Carried and still open (none new):
1. Subtlety raid 60: the shipped APL against the APL with the old Eviscerate line.
2. Subtlety race re-run, both factions, on the Death Mark rotation.
3. Assassination race re-run with the Rupture line.
4. PROVISIONAL: re-rank `bis/rogue-combat.json` band 60 raid on the shipped `raidBuild`.
   - The current rows were ranked on `005203002-3243030000012321451-005` (Efficient Strikes 2, no Composure).
   - The shipped string is `005203002/3043030002012321451/005` (Composure 2, no Efficient Strikes).
   - The stat order still matches `statPriority`.

## Sim findings

Carried, not re-filed:
- **SIM-R11-1.** Subtlety's Eviscerate and Slice and Dice lines still require `dotIsActive(Rupture)`, so band 10 never spends combo points. subtlety.md:129's "Eviscerate at five on mobs" is the right play; the card lags it.
- **Component lane.**
  - Talent spells with trainable level 0 (Ghostly Strike, Cold Blood) still show on the band 10 and 20 cards.
  - The Hemorrhage card note is cut off before its "Hold it for a moment" clause.

No new engine finding.

## Sources

- "World of Warcraft: Forever Class Deep Dives - Rogue and Warlock", Blizzard (200, 2026-11-04): https://news.blizzard.com/en-us/article/24310968/world-of-warcraft-forever-class-deep-dives-rogue-and-warlock
- "World of Warcraft: Forever Deep Dive Panel Recap", Blizzard (200): https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
- "Skyborne First Look: New Neutral Race in World of Warcraft: Forever", Wowhead (200): https://www.wowhead.com/forever/news/skyborne-first-look-new-neutral-race-in-world-of-warcraft-forever-382829
- foreversixty.gg: /planner, /planner?class=rogue and /sim/weights (all 200)
- Client data in `data/builds/1.60.1.70755/`:
  - `talents/rogue.json`, `trainables/rogue.json`, `addon-builds.json` and `addon-data.json`
  - `races.json`, `items/rogue.json` and `bis/rogue-*.json`
  - `spelltips/rogue.json`: 1329283, 1329549, 1328825, 1833, 408, 8643, 14183, 1856, 5277, 2094, 14185, 2983, 1966, 11290, 703, 6774, 5171, 11275, 438040, 1766, 8676, 25300, 53, 1752, 11294, 31016, 399956, 1241582, 1241584, 1316709, 1325758, 16511, 14278, 13877, 13750, 14177, 11300, 11281, 2823, 25351, 11340, 8679
  - `spellconst/rogue.json`: 18427, 14057, 13732, 1329265, 1316709 to 1325758
- APLs: `data/curated/apl/rogue-{assassination,combat,subtlety}.json`
- Engine 9f7f1f6e4: `sim/rogue/mutilate.go` (hands, dagger check, Seal Fate cooldown) and `death_mark.go:13-38`
- Race runs: `design/reviews/guide-accuracy/round3/race-sims.md:18-23`
- Previous rounds: `round12/rogue.md` (fix log), `round11/rogue.md` and `ledger.md`

## Counts

CONFIRMED 0, LIKELY 0, NEEDS-SIM 0, NEEDS-OWNER 0. SHOULDs: 1. New sim findings: 0. One PROVISIONAL BiS re-rank request is carried.

Verdict: zero for rogue. All three round 12 edits hold, and no line on the four pages is false on 70755 or contradicts the page.
