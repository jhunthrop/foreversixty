---
title: Elemental Shaman in Forever
classSlug: shaman
spec: elemental
role: dps
build: 'FS1:1.60.1.70658:shaman:dwarf:551033130000135/03/503352:'
raidBuild: 'FS1:1.60.1.70658:shaman:dwarf:551033130000135/03/053352:'
recommendedRaces: [dwarf, orc]
statPriority: [Intellect, Spell power, Hit, Nature power, Critical strike, Spell haste, Spell penetration]
description: 'Elemental Shaman overview, talent priority, rotation, stat weights, and race picks for Forever, with beta-versus-projection called out.'
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

Elemental is Shaman's ranged spellcasting spec, dealing damage with Lightning Bolt, Chain Lightning, Flame Shock and Lava Burst, backed by fire totems.

### Strengths

- Forever's baseline cast-time cut to Lightning Bolt and Chain Lightning
- Lava Burst hits harder when Flame Shock is already on the target
- Chain Lightning hits every target for close to single-target Lightning Bolt damage
- Searing Totem for one target and Magma Totem for several add steady damage

### Weaknesses

- Slower leveling than Enhancement, because its damage depends on cast times
- Hard casts leave it idle while moving
- Raid build spends points in Restoration to reach spell hit

### In a raid

- **Role:** Ranged spell damage
- **Brings:** Party totems alongside its own damage; no raid-wide buff or debuff

## Talents and builds

In rough priority order: **Concussion** raises the damage of Lightning Bolt, Chain Lightning, and Earth Shock, a broad multiplier across the entire single-target and AoE kit. **Call of Flame**, reworked in Forever, now boosts Lava Burst and Fire Nova alongside Flame Shock within a single talent instead of touching only the Fire Totems it used to buff, effectively unifying the tree's fire-damage kit under one pick. **Elemental Fury** raises critical strike damage on the totems and elements the spec relies on, moved earlier in the tree with a lower first-rank value than 1.12's single large-rank version. **Lightning Overload** gives Lightning Bolt and Chain Lightning a chance to cast a second, weaker bolt for free, a straight damage-per-cast increase. **Elemental Alacrity** cuts the cast time of Lightning Bolt, Chain Lightning, and Lava Burst further on top of the baseline speed-up Forever already gave those spells — maxed, it brings deep Elemental's effective cast speed roughly back to where 1.12 Elemental topped out, since the starting point is already faster. **Lava Burst**, the tree's new capstone nuke, hits harder against a target already carrying Flame Shock.

## Rotation and priority

The loop this site's simulator plays: against a single target, drop Searing Totem for the higher fire damage; against two or more targets, drop Magma Totem instead so everyone standing in it takes damage. Against two or more targets, Chain Lightning is used since it hits every target for close to single-target Lightning Bolt damage. Lightning Bolt fills the rest of the single-target rotation, on every free global. Keep **Flame Shock** on the target: Call of Flame stretches it to 21 seconds, so it costs about one global in fourteen, and its damage plus Lava Burst's 20% bonus beats the Lightning Bolt it replaces. **Lava Burst** is the first line when you have the talent, cast on every cooldown while your Flame Shock is on the target, so every cast carries the bonus. An earlier search left Flame Shock out to save mana, but that run's character was running dry; with raid buffs and consumables, Flame Shock upkeep with Lava Burst waiting for it gained about 9% in our level-60 raid run. Chain Lightning on cooldown against a single target is also worse than plain Lightning Bolt casts, because it costs more than twice the mana for less damage per cast. Earth Shock is not in the loop: it shares the shock cooldown with Flame Shock, which earns that cooldown. Mana Spring Totem goes down first: it lasts five minutes, returns 10 mana every 2 seconds at its top rank, so the 100 mana it costs comes back many times over; adding it was worth roughly 4% in our level-60 run, and starting it one global before the cooldowns was worth about 1.5% more. Fire Nova is a real spell in Forever (learned from level 12, instant, on a 10-second cooldown, and it needs a Fire totem standing), but it is not in the loop: it deals about a third of Lightning Bolt's damage per point of mana, and at no mana threshold we tried did adding it beat plain Lightning Bolt casts. We re-checked these numbers against the client's own spell data, and every one of them got smaller; the order of the loop did not change.

## Stat priority

## Best in slot

## Enchants and consumables

## Races

## Professions

Community convention for caster Shaman favors **Enchanting** alongside **Tailoring** or a gathering profession for spell-power gear and enchant materials, though Shaman's totem and weapon-imbue kit doesn't lean as hard on any one profession as a pure caster class does.

## Leveling

Enhancement is generally the stronger leveling spec for Shaman, since its melee durability and weapon-imbue burst clear trash faster than Elemental's cast-time-dependent damage.
