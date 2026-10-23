---
title: Enhancement Shaman in Forever
classSlug: shaman
spec: enhancement
role: dps
build: 'FS1:1.60.1.70532:shaman:dwarf:320033/255130030005102051/052:'
raidBuild: 'FS1:1.60.1.70532:shaman:dwarf:3230031/255030031005102031/053:'
recommendedRaces: [dwarf, orc]
statPriority: [Strength, Attack power, Hit, Critical strike, Agility, Melee haste]
description: 'Enhancement Shaman overview, talent priority, rotation, stat weights, and race picks for Forever, with beta-versus-projection called out.'
updated: 2026-09-24
confidence: inferred
sources:
  - label: 'Blizzard, Deep Dive panel recap'
    url: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
    kind: blizzard
  - label: 'Talents Forever (demo transcription)'
    url: https://talentsforever.com/data.json
    kind: community
  - label: 'This site, stat weights'
    url: https://foreversixty.gg/sim/weights
    kind: site
  - label: 'This site, character planner'
    url: https://foreversixty.gg/planner
    kind: site
---

## Overview

Enhancement is Shaman's melee spec, fighting with a weapon imbue, Stormstrike and group totems such as Strength of Earth and Windfury.

### Strengths

- Stormstrike is its highest damage-per-global ability
- Windfury Weapon adds melee damage once learned at level 30
- Melee and spell hit are one stat, so hit gear serves both attacks and shocks
- Melee durability clears trash faster than the caster specs

### Weaknesses

- Must stay in melee range, unlike the other Shaman specs
- Imbue and totems need constant upkeep
- Only one weapon carries an imbue, since Shamans do not dual-wield in Forever

### In a raid

- **Role:** Melee damage
- **Brings:** Windfury Totem and Strength of Earth Totem for the melee group

## Talents and builds

In rough priority order: **Thundering Strikes** raises critical strike chance with all attacks and spells, a broad early multiplier. **Elemental Weapons** increases the attack power bonus from Rockbiter Weapon and the proc value of Windfury Weapon and Flametongue Weapon, directly scaling the spec's weapon-imbue kit. **Flurry** grants an attack speed bonus for several swings after a melee crit, compounding with the crit chance Thundering Strikes already added. **Stormstrike**, used heavily in the rotation, is the highest damage-per-global-cooldown ability Enhancement has, on an 8-second cooldown. **Maelstrom Weapon**, new to the tree, can proc off a melee swing and stacks up a discount that shortens Lightning Bolt's cast and trims its mana cost, letting Enhancement weave in a ranged nuke without losing much melee uptime. **Rage of the Farseer**, the tree's new capstone, increases melee attack speed by 30% for 25 seconds as a burst cooldown; it no longer touches casting speed, a change the beta notes made in late September and this site's engine now follows.

Point allocation runs deep into Enhancement to reach Maelstrom Weapon and Rage of the Farseer near the bottom of the tree (33 points, with Ancestral Knowledge raised to five), with 11 in Elemental (Convection 3, Concussion 2, Elemental Alacrity 3, Call of Flame 3) and 7 in Restoration for Tidal Focus (maxed, for its spell hit) and 2 in Mindfulness. The current client moved Elemental Warding and Totemic Focus a tier deeper, past what this budget reaches, so Elemental Alacrity and Tidal Focus hold those points. This replaces the earlier 20/31/0 build, and the talent search on the current engine put it about 2% ahead, beyond the run's error. The points came from Convection, Concussion, Reverberation and Elemental Devastation, which the search measured at little or no damage for this spec, and went to Ancestral Knowledge, Tidal Focus, Call of Flame and Mindfulness; Stormstrike, Maelstrom Weapon, Flurry and the rest of the melee core are untouched. Open the planner at [/planner?class=shaman](/planner?class=shaman) to build this out.

**Raid build.** The raid-ready search found a clearly better spend, about +9% over the leveling build above in the raid-ready run once Earth Shock is back in the loop (see Rotation). It puts three points into Elemental Devastation, which gives a critical strike chance after a spell crit and so only pays when the loop casts a shock, plus single points into Elemental Focus, Shamanistic Focus and Mindfulness, and pays with the three Elemental Alacrity points, two Maelstrom Weapon points and the Guardian Totems point. Stormstrike, Call of Flame, Mindfulness and Spirit Weapons are kept, as are the group buffs the raid already counts on and the threat and survival talents.

## Rotation and priority

Weapon imbues are class kit, not a rotation line — applied before the pull and reapplied whenever it falls off, the same way a Rogue keeps poisons up. Shamans don't dual-wield in Forever, so there's only ever one weapon to imbue: Rockbiter Weapon goes on it from level 1 (it's the only imbue learned that early), and Windfury Weapon replaces Rockbiter there once learned at level 30, staying on through 60. The off hand carries a shield or another held item rather than a second weapon.

The loop this site's simulator plays: keep Strength of Earth Totem down throughout the fight; the engine tracks totem uptime directly rather than needing an aura check, so it is treated as a simple refresh condition. Keep Windfury Totem down as well, since it is the melee group's largest damage totem and occupies the same Air-totem slot Grace of Air would otherwise use. Stormstrike is used on cooldown as the highest damage-per-global ability in the kit. Flame Shock is not part of the loop: it cost more mana than it returned.

## Stat priority

The table above is this level range's own simulation at level 60, re-run by the nightly pipeline every time the build or its gear data changes — these numbers are never hand-entered. A stat shown as "not separable from zero" is not a verdict that it is worthless; the sim's measured error on it is too wide, at this level range's sample size, to tell its true value apart from zero.

## Best in slot

Prioritize attack power, strength, and agility first, then crit and hit, then melee haste. Because hit is now a single stat covering both Enhancement's melee attacks and its shock spells, gear that used to be a tradeoff between the two kinds of hit rating is simply better across the board for this spec than it was in 1.12.

## Enchants and consumables

## Races

Enhancement can be played by Orc and Troll on the Horde, Tauren on the Horde, Dwarf — new in Forever — on the Alliance, and the Windshaper Skyborne on the Horde only. Dwarf is, for now, the only Alliance option for Shaman at all. For Alliance, **Dwarf** is the only pick, and it fits a melee spec reasonably well: Stoneform now reduces physical damage taken rather than only raising armor, a direct survivability gain for a spec standing in melee range. For Horde, **Orc** is the strongest pick: Blood Fury's 10% attack power for 15 seconds on a 2-minute cooldown lines up directly with Enhancement's melee damage kit, and Axe Specialization's crit bonus rewards an Orc Enhancement Shaman for choosing an axe as their main-hand weapon.

## Professions

Community convention for melee Shaman favors **Leatherworking** alongside **Skinning** for self-sourced agility and strength gear, matching general Classic-era melee-class practice.

## Leveling

Enhancement is generally the strongest leveling spec for Shaman: its melee durability, weapon imbues, and Stormstrike burst clear trash faster than Elemental's cast-time-dependent damage or Restoration's support-focused kit.
