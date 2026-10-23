---
title: Retribution Paladin in Forever
classSlug: paladin
spec: retribution
role: dps
build: 'FS1:1.60.1.70532:paladin:human:52003003/052/05025331001330311:'
raidBuild: 'FS1:1.60.1.70532:paladin:human:550030031001/0/0522533100133032:'
recommendedRaces: [human, undead]
statPriority: [Strength, Attack power, Critical strike, Hit, Agility, Melee haste]
description: 'Talents, rotation, stats, gear, races, and professions for Retribution Paladin melee damage in Forever.'
updated: 2026-09-24
confidence: inferred
sources:
  - label: 'Blizzard, Deep Dive panel recap'
    url: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
    kind: blizzard
  - label: 'Forever Sixty build data, data/builds/1.60.1.69893/talents/paladin.json'
    url: https://foreversixty.gg/data/1.60.1.69893/talents/paladin.json
    kind: datamined
  - label: 'Forever Sixty stat weights, data/curated/specs.json'
    url: https://foreversixty.gg/data/curated/specs.json
    kind: site
  - label: 'wowsims-forever default Retribution priority list'
    url: https://github.com/wowsims/classic/blob/master/ui/retribution_paladin/apls/basic_ret.apl.json
    kind: community
  - label: 'Warcraft Tavern, Classic Retribution Paladin rotation'
    url: https://www.warcrafttavern.com/wow-classic/guides/pve-retribution-paladin-rotations-cooldowns/
    kind: community
  - label: 'Forever Sixty raid phase data, data/curated/loot/forever-raid-phases.json'
    url: https://foreversixty.gg/data/curated/loot/forever-raid-phases.json
    kind: site
  - label: 'Forever Sixty build data, data/builds/1.60.1.69893/races.json'
    url: https://foreversixty.gg/data/1.60.1.69893/races.json
    kind: datamined
---

## Overview

Retribution is the Paladin melee damage tree: keep a Seal up, Judge on cooldown, and add Hammer of Wrath on targets low enough to execute.

### Strengths

- Strongest of the three Paladin specs for solo play and leveling
- Judgement is off the global cooldown and drives the damage loop
- Hammer of Wrath adds an execute below roughly 20% health
- Sanctified Judgement sustains more Judgements over a fight

### Weaknesses

- Exorcism works only against Undead and Demons
- Consecration costs too much mana on a single target without raid mana support
- Every Seal recast costs a global and over 200 mana, and seal twisting loses damage

### In a raid

- **Role:** Melee damage dealer
- **Brings:** Its own damage, plus blessings and an aura like any Paladin

## Talents and builds

A newer talent search against this build's own level-60 gear and engine moved this guide's spend deeper into Retribution, 31 points in the tree with 13 in Holy and 7 in Protection. It adds **Reverence** at 3/3 (an engine-measured damage talent in Holy) and **Twist of Light** at 1/1, the capstone of the Retribution tree, which the engine's per-talent probe credits as one of the largest single-point gains available. To pay for them, it drops **Divine Intellect** from 4 points to 2, **Precision** from 3 points to 2, and **Instrument of Law** from 2/2 to 1/2 (the probe measures that talent at zero damage, so the second point is the cheapest one to give up), and it moves the Protection points out of **Toughness** (3/5 → 0/5) and into **Redoubt** (2/5 → 5/5). Both are defensive talents the engine models and measures at zero damage, which is why the swap is free.

**Improved Seals** (3/3), **Divine Strength** (5/5), **Champion of the Light** (3/3), **Vengeance**, **Two-Handed Weapon Specialization**, **Conviction**, **Benediction**, **Seal of Command**, **Vindication**, and **Sacred Arbiter** hold unchanged. **Sanctified Judgement**, held at 3/3 (a 33% chance for Judgement to refund 20% of the judged Seal's mana cost), is still the single largest damage contributor in this build's own probe, since sustaining more Judgements over a fight is worth real DPS even though the tooltip reads like a mana talent. **Holy Conduit**, held at 2/2 (20% less mana cost on Consecration, Holy Wrath, Exorcism, and Hammer of Wrath), is a real if smaller damage gain rather than pure sustain. The rotation casts exactly what it did before: no ability in the priority list goes idle at any level of the ladder. This build measures about +6% over the previous spend in our level-60 search run. Open the planner at [/planner?class=paladin](/planner?class=paladin) to build this out.

**Raid build.** For a raid the search found about +3% over the previous raid build in the raid-ready run, with Consecration in the rotation (below). It moves the five Redoubt points and three Precision, which the simulator measures at no damage or less than their replacements, into Reverence at 3/3, a second point each of Divine Intellect and Purifying Power, Divine Favor and two Holy Conduit points. Seal of Command, Vindication and Instrument of Law stay.

## Rotation and priority

Apply Seal of Righteousness about a second and a half before the pull so it is active the moment combat starts. Once in combat, Judgement comes first whenever it is off cooldown: it is off the global cooldown, costs 6% of your base mana, unloads the active Seal's Judgement effect and immediately starts its own cooldown again. Exorcism follows on its own 15-second cooldown against Undead and Demons only; it is free Holy damage on those encounters, simply unavailable against anything else.

This rotation does not seal-twist. Older community guides for 1.12 Retribution (and an earlier version of this page) describe alternating Seal of Righteousness and Seal of Command swing-by-swing for a DPS gain. Each Seal recast still costs a full global and 200+ mana at level 60, so twisting on every swing mostly just burns mana for nothing. Just keep Seal of Righteousness up and spend every other global on Judgement, Exorcism, Hammer of Wrath and Holy Strike in that priority order.

Consecration, also new in Forever's baseline kit this week, is absent from the leveling and bare rotation above and joins it last in the raid-ready one: against a single target it is a loss on a bare character, where its mana cost (scaling steeply with rank, 135 at rank 1 up to 565 at rank 5) crowds out Judgement, Exorcism and Holy Strike, but with the raid's mana support the rotation search measured it as a clear gain, about +9% in the raid-ready run, and it holds the bare character level when it is only cast while mana is above 65%.

## Stat priority

The table above is this level range's own simulation at level 60, re-run by the nightly pipeline every time the build or its gear data changes — these numbers are never hand-entered. A stat shown as "not separable from zero" is not a verdict that it is worthless; the sim's measured error on it is too wide, at this level range's sample size, to tell its true value apart from zero.

## Best in slot

Look for pieces that lead with attack power and strength, matching the stat priority above.

## Enchants and consumables

## Races

Retribution Paladin can only be Human or Dwarf on Alliance, or Undead on Horde — this build's race and class combination table confirms no other race can train Paladin at all. Between the two Alliance options, Human is the better damage pick if wielding a two-handed sword: its reworked Sword Specialization grants a flat crit chance bonus with swords, which applies to every Judgement and melee swing. Dwarf's Mace Specialization offers the same kind of bonus but only for maces, so the choice mostly follows weapon type rather than one race being flatly stronger. On Horde, Undead is currently the only playable Paladin race; Will of the Forsaken's fear, charm, and sleep cleanse is a useful tool for a melee spec that otherwise has few ways to shrug off crowd control mid-fight, even though it no longer grants brief immunity the way it did in 1.12.

## Professions

Blacksmithing suits Retribution well since it can add sockets to weapons and armor, letting a melee Paladin fill in strength or crit the itemization doesn't provide on its own. Enchanting offers weapon and ring enchants a Retribution Paladin can apply without paying another crafter.

## Leveling

Retribution is the strong leveling choice for Paladin — it kills things fastest solo of the three specs, which is why it is worth considering even for a character planning to heal or tank later.

Consecration (baseline from level 20, no talent needed) is worth reaching for while leveling even though this page's single-target rotation leaves it out: questing and farming routinely put two or three mobs on you at once, and Consecration's ground effect hits all of them for the one cast and mana cost, which is a better trade than it is against a single raid target.
