---
title: Restoration Shaman in Forever
classSlug: shaman
spec: restoration
role: healer
build: 'FS1:1.60.1.70532:shaman:dwarf:53/0/5532003315513151:'
recommendedRaces: [dwarf, tauren]
statPriority: [Healing power, Intellect, Spirit, MP5, Critical strike]
description: 'Restoration Shaman overview, talent priority, healing priority, stat weights, and race picks for Forever, with beta-versus-projection called out.'
updated: 2026-10-07
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

Restoration is Shaman's healing spec, built on Healing Wave, Chain Heal and Riptide, with totems that sustain the group's mana and health.

### Strengths

- Chain Heal jumps to the most injured party members
- Riptide and Nature's Swiftness give fast healing on the tank
- Mana Spring Totem and Mana Tide Totem support the group's mana
- Water Shield returns mana as it procs

### Weaknesses

- Needs raid-ready gear before its mana lasts a whole fight
- Slow solo leveling, since its group kit does not translate
- Totems sit in one place and must be dropped again as the group moves

### In a raid

- **Role:** Healer, centered on the tank with Chain Heal for the group
- **Brings:** Healing Stream Totem, Mana Spring Totem and Mana Tide Totem for the group

## Talents and builds

The build is nearly all Restoration, with the rest in Elemental where the points cost a healer nothing. Chosen by simming candidate builds on the same gear against the same fight. Every figure behind this ranking comes from one stated incoming-damage profile (the scaled Phase 1 tank hits and raid pulses described on the [BiS page](/bis)): it ranks gear and builds for this spec, and says nothing about which healer class is stronger.

- **Improved Healing Wave** shortens Healing Wave, and **Healing Way** strengthens it; **Purification** is a flat multiplier on every heal.
- **Tidal Focus** makes the healing spells cheaper (and adds hit, which a healer does not use), **Mindfulness** keeps part of your regeneration running while you cast, and **Tidal Mastery** adds critical chance.
- **Water Shield** returns mana on a critical heal, **Mana Tide Totem** refills the group, **Restorative Totems** strengthens Healing Stream and Mana Spring, and **Healing Focus** protects your casts from pushback.
- **Nature's Swiftness** is the instant emergency heal, and **Riptide** is the capstone: a direct heal plus a heal over time, and a stronger Chain Heal on its target.

The sim found any build that keeps that core scores the same, because the points that differ change nothing a healer casts. Open the planner at [/planner?class=shaman](/planner?class=shaman) to adjust it.

## Rotation and priority

Water Shield and a Healing Stream Totem go up before the pull. In the fight the tank comes first: **Riptide** whenever it is ready and the tank is hurt, **Healing Wave** for a deeper hole, **Lesser Healing Wave** when it is serious, and **Nature's Swiftness** into an instant Healing Wave when the tank is nearly dead. **Chain Heal** is cast on the tank, so Riptide's bonus applies, and it jumps to the most injured party members. Members are healed directly only while your mana is healthy, so the tank keeps the mana.

Under the profile the shaman's mana lasts the whole fight at raid-ready gear, so the rotation spends the spare on members. It does not downrank.

## Stat priority

In simulator-derived order, per point of stat: **MP5** first, then **Intellect**, **Spirit** after it, **healing power** and **critical strike** behind. The shaman's mana lasts the whole fight at raid-ready gear, so extra throughput is what turns into healing; before that gear the same order holds because regeneration is what keeps the casts coming. Items trade several points of one for few of another, so read the list as "which stat is cheap to give up".

Spell power adds nothing to a heal. Every figure behind this ranking comes from one stated incoming-damage profile (the scaled Phase 1 tank hits and raid pulses described on the [BiS page](/bis)): it ranks gear and builds for this spec, and says nothing about which healer class is stronger.

## Best in slot

Follow the stat priority above; spell power on a healing item adds nothing to a heal. Water Shield's proc-based mana return makes a slightly lower spirit floor more survivable than in 1.12, but spirit remains the tree's primary sustain stat.

## Enchants and consumables

## Races

Restoration can be played by Orc and Troll on the Horde, Tauren on the Horde, Dwarf — new in Forever — on the Alliance, and the Windshaper Skyborne on the Horde only. Dwarf is, for now, the only Alliance option for Shaman at all. For Alliance, **Dwarf** is the only pick; its racials lean offensive rather than healer-focused, but Stoneform's physical damage reduction still helps a Restoration Shaman survive melee pressure while casting. For Horde, **Tauren** is the strongest pick for Restoration specifically: Endurance grants 5% health and 1% hit chance, both passive survivability and consistency gains that matter more for a support caster standing in raid damage than an offensive cooldown like Orc's Blood Fury does.

## Professions

Community convention for healing Shaman favors **Herbalism** alongside **Alchemy** for self-sourced healing and mana potions, or **Tailoring** for caster cloth-adjacent leather stats, matching general Classic-era healer practice.

## Leveling

Enhancement is the generally recommended leveling spec for Shaman, not Restoration — Restoration's group-support kit doesn't translate to fast solo leveling the way Enhancement's melee durability and burst does.
