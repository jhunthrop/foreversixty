---
title: Restoration Shaman in Forever
classSlug: shaman
spec: restoration
role: healer
build: 'FS1:1.60.1.69893:shaman:dwarf:0/005/5532500010513001:'
recommendedRaces: [dwarf, tauren]
statPriority: [Healing power, Spell power, Spirit, MP5, Intellect, Critical strike]
description: 'Restoration Shaman overview, talent priority, healing priority, stat weights, and race picks for Forever, with beta-versus-projection called out.'
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

Restoration is Shaman's dedicated healing tree, centered on Healing Wave and totem-based group sustain like Mana Spring Totem and Mana Tide Totem, with Water Shield and Riptide adding two new tools to the kit. Its group-wide mana support through totems is historically its distinguishing feature next to Priest and Druid healing. The beta caps at level 30, so how Restoration performs healing a level-60 raid, or how it stacks up against other healing classes at that scale, is a projection from the demo talent trees and 1.12 healing knowledge, not confirmed play.

## Talents and builds

In rough priority order: **Improved Healing Wave** cuts Healing Wave's cast time, a direct throughput gain on the tree's main heal. **Tidal Focus** reduces the mana cost of healing spells and improves hit chance, an efficiency gain that compounds over a full fight. **Restorative Totems** increases the effect of Mana Spring Totem and Healing Stream Totem, raising the value of Restoration's group-support totems specifically. **Mana Tide Totem** is a dedicated raid mana-regeneration cooldown, valuable enough that most Restoration builds reach it. **Water Shield**, new to the tree, surrounds the caster with globes of water that restore mana when the Shaman is hit or lands a healing crit, adding a passive mana-sustain layer that didn't exist in 1.12. **Riptide**, the tree's new capstone, heals a target instantly and again over time while increasing the effectiveness of subsequent heals on that target, combining an instant heal with a HoT and a healing-taken buff in one cast.

Point allocation runs deep into Restoration to reach Riptide at the bottom of the tree, with the remainder split as a handful of points in Enhancement for Ancestral Knowledge or a similar utility talent. Open the planner at [/planner?class=shaman](/planner?class=shaman) to build this out.

## Rotation and priority

There's no curated rotation data for Restoration in this build — it's a healing spec, not a DPS one this site's simulator models. Based on general Classic-era Restoration practice plus what's changed for Forever: open with Riptide where talented, since it front-loads an instant heal before its HoT and healing-taken buff continue working. Use Healing Wave as the primary throughput heal, backed by Lesser Healing Wave for cheaper, faster top-offs. Keep totems appropriate to the group's needs down — Mana Spring Totem for sustained fights, Healing Stream Totem for spread raid damage — and use Mana Tide Totem as a group mana cooldown when the raid's casters are running low. This is this site's own inference from 1.12 healing conventions and the demo talent data, not a tested rotation.

## Stat priority

In simulator-derived priority order: **healing power** first, the direct multiplier on every heal cast. **Spell power** next, since Forever's itemization change lets bonus healing gear also carry a fraction of spell damage. **Spirit** feeds mana regeneration, compounding with Water Shield's proc-based mana return. **Mp5** provides flat regeneration independent of Spirit-based formulas. **Intellect** adds mana pool and a small crit chance. **Crit** last, both for direct heal size and Tidal Mastery's bonus crit if talented.

## Gear

Prioritize healing power and spell power first, then spirit and mp5 for sustain, then intellect and crit. Water Shield's proc-based mana return makes a slightly lower spirit floor more survivable than in 1.12, but spirit remains the tree's primary sustain stat. Specific pre-raid or raid-tier item picks can't be named with confidence yet: the beta caps at level 30 and this build's raid loot tables are still missing most items the community's sourcing expects, with nothing raiding in-game until the first tier opens on 9 December 2026.

## Enchants and consumables

Weapon enchants should target healing power; **Enchant Weapon - Healing Power** is a verified entry in this build's enchant data. For consumables, **Flask of Distilled Wisdom** is a verified option historically tied to mana regeneration for healers. A fuller consumable stack isn't something this site is naming until the class is played past level 30.

## Races

Restoration can be played by Orc and Troll on the Horde, Tauren on the Horde, Dwarf — new in Forever — on the Alliance, and the Windshaper Skyborne on the Horde only. Dwarf is, for now, the only Alliance option for Shaman at all. For Alliance, **Dwarf** is the only pick; its racials lean offensive rather than healer-focused, but Stoneform's physical damage reduction still helps a Restoration Shaman survive melee pressure while casting. For Horde, **Tauren** is the strongest pick for Restoration specifically: Endurance grants 5% health and 1% hit chance, both passive survivability and consistency gains that matter more for a support caster standing in raid damage than an offensive cooldown like Orc's Blood Fury does.

## Professions

Community convention for healing Shaman favors **Herbalism** alongside **Alchemy** for self-sourced healing and mana potions, or **Tailoring** for caster cloth-adjacent leather stats, matching general Classic-era healer practice. This is community convention, not confirmed for Forever.

## Leveling

Enhancement is the generally recommended leveling spec for Shaman, not Restoration — Restoration's group-support kit doesn't translate to fast solo leveling the way Enhancement's melee durability and burst does. The beta caps at level 30, so this is carried over from 1.12 leveling knowledge rather than tested in Forever.
