---
title: Retribution Paladin in Forever
classSlug: paladin
spec: retribution
role: dps
build: 'FS1:1.60.1.70658:paladin:human:52003003/052/05025331001330311:'
raidBuild: 'FS1:1.60.1.70658:paladin:human:52003/503/0522533100133032:'
recommendedRaces: [human, undead]
statPriority: [Strength, Attack power, Critical strike, Hit, Agility, Melee haste]
description: 'Talents, rotation, stats, gear, races, and professions for Retribution Paladin melee damage in Forever.'
updated: 2026-10-31
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
- **Brings:** Its own damage, plus Blessing of Might and Blessing of Wisdom

## Talents and builds

**Improved Seals** (3/3), **Divine Strength** (5/5), **Champion of the Light** (3/3), **Vengeance**, **Two-Handed Weapon Specialization**, **Conviction**, **Benediction**, **Seal of Command**, **Vindication**, and **Sacred Arbiter** hold unchanged. **Sanctified Judgement**, held at 3/3 (a 33% chance for Judgement to refund 20% of the judged Seal's mana cost), is still the single largest damage contributor in this build's own probe, since sustaining more Judgements over a fight is worth real DPS even though the tooltip reads like a mana talent. **Holy Conduit**, held at 2/2 (20% less mana cost on Consecration, Holy Wrath, Exorcism, and Hammer of Wrath), is a real if smaller damage gain rather than pure sustain. The rotation casts exactly what it did before: no ability in the priority list goes idle at any level of the ladder. This build measures about +6% over the previous spend in our level-60 search run. The raid build also takes **Precision** (3/3), 3% hit behind five points of Toughness, with the points moved out of Reverence, Purifying Power, Divine Favor and three of Divine Intellect: none of those adds damage to a melee paladin, and the build measured about 2% ahead in our level-60 raid search run. Open the planner at [/planner?class=paladin](/planner?class=paladin) to build this out.

## Rotation and priority

Apply Seal of Righteousness about a second and a half before the pull so it is active the moment combat starts. Once in combat, Judgement comes first whenever it is off cooldown: it is off the global cooldown, costs 6% of your base mana, unloads the active Seal's Judgement effect and immediately starts its own cooldown again. Exorcism follows on its own 15-second cooldown against Undead and Demons only; it is free Holy damage on those encounters, simply unavailable against anything else. Below roughly 20% target health, Hammer of Wrath takes priority as the execute option, cast at its highest rank available at level 60. Next, reapply Seal of Righteousness just before it expires so uptime never lapses; that line also catches the seal being consumed by Judgement's own cast, which fires this same reapply on the very next decision. Putting the seal refresh ahead of Holy Strike was the one change the search found, worth about +3.5% in our level-60 search run. Holy Strike, Forever's new baseline Paladin attack added to the client this week, comes last, on its own 10-second cooldown, cast whenever it is up: it is what spends the globals that Judgement and seal upkeep alone leave idle.

This rotation does not seal-twist. Older community guides for 1.12 Retribution (and an earlier version of this page) describe alternating Seal of Righteousness and Seal of Command swing-by-swing for a DPS gain. Measured directly against Forever's own engine, that twist is a DPS **loss** here: casting a new Seal always fully replaces whatever Seal was active a moment before (there is no overlap window for a swing to benefit from both, unlike the timing trick 1.12's version of the technique exploited), and Judgement always ends by reapplying Seal of Righteousness regardless of which Seal was up when it fired, so twisting for Judgement's benefit does not work either. Each Seal recast still costs a full global and 200+ mana at level 60, so twisting on every swing mostly just burns mana for nothing. A forever-sim comparison at level 60 in band-60 gear found the old twist-every-swing rotation at 153.8 DPS, the same rotation with the twist removed at 165.6 DPS, and this page's current rotation (twist removed, Holy Strike added) at 190.0 DPS — a 23.5% gain over the twisting version. Just keep Seal of Righteousness up and spend every other global on Judgement, Exorcism, Hammer of Wrath and Holy Strike in that priority order.

Consecration, also new in Forever's baseline kit this week, is absent from the leveling and bare rotation above and joins it last in the raid-ready one: against a single target it is a loss on a bare character, where its mana cost (scaling steeply with rank, 135 at rank 1 up to 565 at rank 5) crowds out Judgement, Exorcism and Holy Strike, but with the raid's mana support the rotation search measured it as a clear gain, about +9% in the raid-ready run, and it holds the bare character level when it is only cast while mana is above 65%. The client states no damage on the cast spell itself; the damage is on a companion row per rank, which the engine's numbers are held to by a test. While leveling, or fighting several enemies at once, it is still worth reaching for (see Leveling, below).

## Stat priority

## Best in slot

## Enchants and consumables

## Races

## Professions

Blacksmithing suits Retribution well since it can add sockets to weapons and armor, letting a melee Paladin fill in strength or crit the itemization doesn't provide on its own. Enchanting offers weapon and ring enchants a Retribution Paladin can apply without paying another crafter.

## Leveling

Retribution is the strong leveling choice for Paladin — it kills things fastest solo of the three specs, which is why it is worth considering even for a character planning to heal or tank later.

Consecration (baseline from level 20, no talent needed) is worth reaching for while leveling even though this page's single-target rotation leaves it out: questing and farming routinely put two or three mobs on you at once, and Consecration's ground effect hits all of them for the one cast and mana cost, which is a better trade than it is against a single raid target.
