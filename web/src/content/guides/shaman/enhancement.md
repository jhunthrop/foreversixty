---
title: Enhancement Shaman in Forever
classSlug: shaman
spec: enhancement
role: dps
build: 'FS1:1.60.1.70658:shaman:dwarf:320033/255130030005102051/052:'
raidBuild: 'FS1:1.60.1.70658:shaman:dwarf:0533001/253030031005102051/053:'
recommendedRaces: [dwarf, orc]
statPriority: [Strength, Attack power, Hit, Critical strike, Agility, Melee haste]
description: 'Enhancement Shaman overview, talent priority, rotation, stat weights, and race picks for Forever, with beta-versus-projection called out.'
updated: 2026-10-31
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

In rough priority order: **Thundering Strikes** raises critical strike chance with all attacks and spells, a broad early multiplier. **Elemental Weapons** increases the attack power bonus from Rockbiter Weapon and the proc value of Windfury Weapon and Flametongue Weapon, directly scaling the spec's weapon-imbue kit. **Flurry** grants an attack speed bonus for several swings after a melee crit, compounding with the crit chance Thundering Strikes already added. **Stormstrike**, used heavily in the rotation, is the highest damage-per-global-cooldown ability Enhancement has, on an 8-second cooldown. **Maelstrom Weapon**, new to the tree, can proc off a melee swing and stacks up a discount that shortens Lightning Bolt's cast and trims its mana cost, letting Enhancement weave in a ranged nuke without losing much melee uptime. Each stack takes 4% a rank off the bolt, so the raid build takes all five ranks: at five stacks Lightning Bolt is instant and free. It pays for those points with two of Ancestral Knowledge's five and with Convection and Call of Flame from Elemental, and takes Concussion (more Lightning Bolt and Earth Shock damage) and Reverberation (a shorter shock cooldown) there instead; with the Lightning Bolt line below, that gained about 3% in our level-60 raid run.

## Rotation and priority

Weapon imbues are class kit, not a rotation line — applied before the pull and reapplied whenever it falls off, the same way a Rogue keeps poisons up. Shamans don't dual-wield in Forever, so there's only ever one weapon to imbue: Rockbiter Weapon goes on it from level 1 (it's the only imbue learned that early), and Windfury Weapon replaces Rockbiter there once learned at level 30, staying on through 60. The off hand carries a shield or another held item rather than a second weapon.

The loop this site's simulator plays: keep Strength of Earth Totem down throughout the fight; the engine tracks totem uptime directly rather than needing an aura check, so it is treated as a simple refresh condition. Keep Windfury Totem down as well, since it is the melee group's largest damage totem and occupies the same Air-totem slot Grace of Air would otherwise use. Stormstrike is used on cooldown as the highest damage-per-global ability in the kit. At five Maelstrom Weapon stacks, cast Lightning Bolt right after Stormstrike, so Stormstrike's 20% on your next Lightning Bolt, Chain Lightning or Earth Shock goes to the bolt. Searing Totem is kept down for extra fire damage, but not refreshed with less than 20 seconds left on the fight, since there is no point paying totem mana for a totem that will barely tick before the encounter ends. Mana Spring Totem is dropped before the pull and kept up: the Water-totem slot is otherwise empty, and its mana regeneration was worth about 1.5% in our level-60 run. Fire Nova, the instant Fire-totem spell learned from level 12, is not in the loop either: every global spent on it cost more mana than its damage returned, at every mana threshold we tried. Flame Shock is not part of the loop: it cost more mana than it returned. Earth Shock is, last in the list and only while mana is above half: on the current engine, with the shock tables held to the client's numbers, the rotation search measured it as a large gain in the raid-ready run (more than +10%), and the mana gate keeps the bare character, which has no raid mana support, level.

## Stat priority

## Best in slot

Because hit is now a single stat covering both Enhancement's melee attacks and its shock spells, gear that used to be a tradeoff between the two kinds of hit rating is simply better across the board for this spec than it was in 1.12.

## Enchants and consumables

## Races

## Professions

Community convention for melee Shaman favors **Leatherworking** alongside **Skinning** for self-sourced agility and strength gear, matching general Classic-era melee-class practice.

## Leveling

Enhancement is generally the strongest leveling spec for Shaman: its melee durability, weapon imbues, and Stormstrike burst clear trash faster than Elemental's cast-time-dependent damage or Restoration's support-focused kit.
