---
title: Enhancement Shaman in Forever
classSlug: shaman
spec: enhancement
role: dps
build: 'FS1:1.60.1.70009:shaman:dwarf:553322/253130030005102051/0:'
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

Enhancement is Shaman's melee DPS spec, combining weapon imbues like Windfury Weapon with Stormstrike and totem support to fight in melee range rather than at range. Forever's unification of melee and spell hit into a single stat is a particularly large indirect buff to this spec, since 1.12 Enhancement had to gear separately for melee hit against its physical attacks and spell hit against its shock spells. The beta caps at level 30, so how Enhancement's damage compares to other melee specs at level 60, or performs in a raid, is a projection from the demo talent trees and 1.12 knowledge, not confirmed play.

## Talents and builds

In rough priority order: **Thundering Strikes** raises critical strike chance with all attacks and spells, a broad early multiplier. **Elemental Weapons** increases the attack power bonus from Rockbiter Weapon and the proc value of Windfury Weapon and Flametongue Weapon, directly scaling the spec's weapon-imbue kit. **Flurry** grants an attack speed bonus for several swings after a melee crit, compounding with the crit chance Thundering Strikes already added. **Stormstrike**, used heavily in the rotation, is the highest damage-per-global-cooldown ability Enhancement has, on an 8-second cooldown. **Maelstrom Weapon**, new to the tree, can proc off a melee swing and stacks up a discount that shortens Lightning Bolt's cast and trims its mana cost, letting Enhancement weave in a ranged nuke without losing much melee uptime. **Rage of the Farseer**, the tree's new capstone, increases melee attack speed by 30% for 25 seconds as a burst cooldown; it no longer touches casting speed, a change the beta notes made in late September and this site's engine now follows.

Point allocation runs deep into Enhancement to reach Maelstrom Weapon and Rage of the Farseer near the bottom of the tree (31 points), with the remaining 20 in Elemental for Concussion's shock damage — Convection and Concussion both maxed, then Elemental Warding, Reverberation, Call of Flame, and Elemental Devastation for the rest of the budget. Open the planner at [/planner?class=shaman](/planner?class=shaman) to build this out.

## Rotation and priority

Weapon imbues are class kit, not a rotation line — applied before the pull and reapplied whenever it falls off, the same way a Rogue keeps poisons up. Shamans don't dual-wield in Forever, so there's only ever one weapon to imbue: Rockbiter Weapon goes on it from level 1 (it's the only imbue learned that early), and Windfury Weapon replaces Rockbiter there once learned at level 30, staying on through 60. The off hand carries a shield or another held item rather than a second weapon.

The loop this site's simulator plays: keep Strength of Earth Totem down throughout the fight; the engine tracks totem uptime directly rather than needing an aura check, so it is treated as a simple refresh condition. Keep Windfury Totem down as well, since it is the melee group's largest damage totem and occupies the same Air-totem slot Grace of Air would otherwise use. Stormstrike is used on cooldown as the highest damage-per-global ability in the kit. Searing Totem is kept down for extra fire damage, but not refreshed with less than 20 seconds left on the fight, since there is no point paying totem mana for a totem that will barely tick before the encounter ends. Mana Spring Totem is dropped before the pull and kept up: the Water-totem slot is otherwise empty, and its mana regeneration was worth about 1.5% in our level-60 run. Fire Nova, the instant Fire-totem spell learned from level 12, is not in the loop either: every global spent on it cost more mana than its damage returned, at every mana threshold we tried. Flame Shock and Earth Shock are no longer part of the loop: in our level-60 search run each cost more mana than it returned, and dropping both was worth about +4%.

## Stat priority

The table above is this band's own simulation at level 60, re-run by the nightly pipeline every time the build or its gear data changes — these numbers are never hand-entered. A stat shown as "not separable from zero" is not a verdict that it is worthless; the sim's measured error on it is too wide, at this band's sample size, to tell its true value apart from zero. These numbers come from this site's own level-60 simulator, not from beta play, which only reaches level 30.

## Gear

Prioritize attack power, strength, and agility first, then crit and hit, then melee haste. Because hit is now a single stat covering both Enhancement's melee attacks and its shock spells, gear that used to be a tradeoff between the two kinds of hit rating is simply better across the board for this spec than it was in 1.12. Specific pre-raid or raid-tier item picks can't be named with confidence yet: the beta caps at level 30 and this build's raid loot tables are still missing most items the community's sourcing expects, with nothing raiding in-game until the first tier opens on 9 December 2026.

## Enchants and consumables

Weapon enchants should target strength or agility depending on the piece; **Enchant Weapon - Strength** and **Enchant Weapon - Agility** are both verified entries in this build's enchant data, alongside **Enchant Weapon - Crusader**, a proc-based option also present in the data. For consumables, **R.O.I.D.S.** (strength) is a verified option. A fuller consumable stack isn't something this site is naming until the class is played past level 30.

## Races

Enhancement can be played by Orc and Troll on the Horde, Tauren on the Horde, Dwarf — new in Forever — on the Alliance, and the Windshaper Skyborne on the Horde only. Dwarf is, for now, the only Alliance option for Shaman at all. For Alliance, **Dwarf** is the only pick, and it fits a melee spec reasonably well: Stoneform now reduces physical damage taken rather than only raising armor, a direct survivability gain for a spec standing in melee range. For Horde, **Orc** is the strongest pick: Blood Fury's 10% attack power for 15 seconds on a 2-minute cooldown lines up directly with Enhancement's melee damage kit, and Axe Specialization's crit bonus rewards an Orc Enhancement Shaman for choosing an axe as their main-hand weapon.

## Professions

Community convention for melee Shaman favors **Leatherworking** alongside **Skinning** for self-sourced agility and strength gear, matching general Classic-era melee-class practice. This is community convention, not confirmed for Forever.

## Leveling

Enhancement is generally the strongest leveling spec for Shaman: its melee durability, weapon imbues, and Stormstrike burst clear trash faster than Elemental's cast-time-dependent damage or Restoration's support-focused kit. The beta caps at level 30, so this is carried over from 1.12 leveling knowledge rather than tested in Forever.
