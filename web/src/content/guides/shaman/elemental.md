---
title: Elemental Shaman in Forever
classSlug: shaman
spec: elemental
role: dps
build: 'FS1:1.60.1.69893:shaman:dwarf:5532300500103031/0/55334:'
recommendedRaces: [dwarf, orc]
statPriority: [Spell power, Intellect, Critical strike, Hit, Spell haste, Spell penetration, 'Nature power']
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

Elemental is Shaman's ranged spellcasting DPS spec, built on Lightning Bolt and Chain Lightning backed by Flame Shock and totem support, with Lava Burst adding a new high-damage nuke that hits harder when Flame Shock is already on the target. Forever's baseline cast-time cut to Lightning Bolt and Chain Lightning, before any talent points are spent, is a meaningful floor-raise for the spec's damage even before talents are considered. The beta caps at level 30, so how Elemental's damage compares to other casters at level 60, or performs in a raid, is a projection from the demo talent trees and 1.12 knowledge, not confirmed play.

## Talents and builds

In rough priority order: **Concussion** raises the damage of Lightning Bolt, Chain Lightning, and Earth Shock, a broad multiplier across the entire single-target and AoE kit. **Call of Flame**, reworked in Forever, now boosts Lava Burst and Fire Nova alongside Flame Shock within a single talent instead of touching only the Fire Totems it used to buff, effectively unifying the tree's fire-damage kit under one pick. **Elemental Fury** raises critical strike damage on the totems and elements the spec relies on, moved earlier in the tree with a lower first-rank value than 1.12's single large-rank version. **Lightning Overload** gives Lightning Bolt and Chain Lightning a chance to cast a second, weaker bolt for free, a straight damage-per-cast increase. **Elemental Alacrity** cuts the cast time of Lightning Bolt, Chain Lightning, and Lava Burst further on top of the baseline speed-up Forever already gave those spells — maxed, it brings deep Elemental's effective cast speed roughly back to where 1.12 Elemental topped out, since the starting point is already faster. **Lava Burst**, the tree's new capstone nuke, hits harder against a target already carrying Flame Shock.

Point allocation runs deep into Elemental to reach Lava Burst at the bottom of the tree, with the remainder split as a handful of points in Restoration for mana-sustain talents. Open the planner at [/planner?class=shaman](/planner?class=shaman) to build this out.

## Rotation and priority

The loop this site's simulator plays: against a single target, drop Searing Totem for the higher fire damage; against two or more targets, drop Magma Totem instead so everyone standing in it takes damage. Flame Shock is kept up on the target throughout, since it's cheap relative to its damage and doesn't compete heavily for the global cooldown budget. Against two or more targets, Chain Lightning is used since it hits every target for close to single-target Lightning Bolt damage. Earth Shock is slotted in as a mana dump once mana is comfortably above half, since it shares a cooldown only with Flame Shock rather than with Chain Lightning, making it safe to weave in. Lightning Bolt fills the rest of the single-target rotation.

## Stat priority

In simulator-derived priority order: **spell power** first, the direct multiplier on every offensive spell. **Intellect** next for mana pool and a small crit contribution. **Crit** raises both direct damage and Elemental Fury's bonus crit damage on totems and spells. **Hit** keeps casts landing consistently. **Spell haste** speeds up the whole rotation, compounding with Elemental Alacrity. **Spell penetration** helps against magic-resistant targets. **Nature power**, gear boosting Nature damage specifically, sits last as the most specialized stat.

## Gear

Prioritize spell power and intellect first, then crit and hit, then spell haste, spell penetration, and Nature-specific power. Specific pre-raid or raid-tier item picks can't be named with confidence yet: the beta caps at level 30 and this build's raid loot tables are still missing most items the community's sourcing expects, with nothing raiding in-game until the first tier opens on 9 December 2026. This section fills in once raid loot is itemized and the class is played past 30.

## Enchants and consumables

Weapon enchants should target spell power; **Enchant Weapon - Spell Power** is a verified entry in this build's enchant data. For consumables, **Greater Arcane Elixir** (spell power and crit) is a verified option. A fuller consumable stack isn't something this site is naming until the class is played past level 30.

## Races

Elemental can be played by Orc and Troll on the Horde, Tauren on the Horde, Dwarf — new in Forever — on the Alliance, and the Windshaper Skyborne on the Horde only. Dwarf is, for now, the only Alliance option for Shaman at all. For Alliance, **Dwarf** is the only pick, but it's a reasonable one for Elemental: Big Game Hunter adds damage against Beasts, useful for leveling and questing damage even if it doesn't apply in most raid encounters. For Horde, **Orc** is the strongest pick: Blood Fury grants 10% attack power and spell power for 15 seconds on a 2-minute cooldown, and the spell power half applies directly to Elemental's entire damage kit, making it a straightforward burst cooldown with no downside for a caster build.

## Professions

Community convention for caster Shaman favors **Enchanting** alongside **Tailoring** or a gathering profession for spell-power gear and enchant materials, though Shaman's totem and weapon-imbue kit doesn't lean as hard on any one profession as a pure caster class does. This is general Classic-era community practice, not confirmed for Forever.

## Leveling

Enhancement is generally the stronger leveling spec for Shaman, since its melee durability and weapon-imbue burst clear trash faster than Elemental's cast-time-dependent damage. The beta caps at level 30, so this is carried over from 1.12 leveling knowledge rather than tested in Forever.
