---
title: Restoration Shaman in Forever
classSlug: shaman
spec: restoration
role: healer
build: 'FS1:1.60.1.70009:shaman:dwarf:53/0/5032503315513151:'
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

Restoration is Shaman's dedicated healing tree, centered on Healing Wave and totem-based group sustain like Mana Spring Totem and Mana Tide Totem, with Water Shield and Riptide adding two new tools to the kit. Its group-wide mana support through totems is historically its distinguishing feature next to Priest and Druid healing. The beta caps at level 30, so how Restoration performs healing a level-60 raid, or how it stacks up against other healing classes at that scale, is a projection from the demo talent trees and 1.12 healing knowledge, not confirmed play.

## Talents and builds

The build is nearly all Restoration, with the rest in Elemental where the points cost a healer nothing. Chosen by simming candidate builds on the same gear against the same fight. Every figure behind this ranking comes from one stated incoming-damage profile (the scaled Phase 1 tank hits and raid pulses described on the [BiS page](/bis)): it ranks gear and builds for this spec, and says nothing about which healer class is stronger.

- **Improved Healing Wave** shortens Healing Wave, and **Healing Way** strengthens it; **Purification** is a flat multiplier on every heal.
- **Tidal Focus** makes the healing spells cheaper (and adds hit, which a healer does not use), **Mindfulness** keeps part of your regeneration running while you cast, and **Tidal Mastery** adds critical chance.
- **Water Shield** returns mana on a critical heal, **Mana Tide Totem** refills the group, **Restorative Totems** strengthens Healing Stream and Mana Spring, and **Healing Focus** protects your casts from pushback.
- **Nature's Swiftness** is the instant emergency heal, and **Riptide** is the capstone: a direct heal plus a heal over time, and a stronger Chain Heal on its target.

The sim found any build that keeps that core scores the same, because the points that differ change nothing a healer casts. Open the planner at [/planner?class=shaman](/planner?class=shaman) to adjust it.

## Rotation and priority

Water Shield and a Healing Stream Totem go up before the pull. In the fight the tank comes first: **Riptide** whenever it is ready and the tank is hurt, **Healing Wave** for a deeper hole, **Lesser Healing Wave** when it is serious, and **Nature's Swiftness** into an instant Healing Wave when the tank is nearly dead. **Chain Heal** is cast on the tank, so Riptide's bonus applies, and it jumps to the most injured party members. Members are healed directly only while your mana is healthy, so the tank keeps the mana. **Mana Tide Totem** goes down when you are low.

Under the profile the shaman's mana lasts the whole fight at raid-ready gear, so the rotation spends the spare on members. It does not downrank.

## Stat priority

In simulator-derived order, per point of stat: **MP5** first, then **Intellect**, **Spirit** after it, **healing power** and **critical strike** behind. The shaman's mana lasts the whole fight at raid-ready gear, so extra throughput is what turns into healing; before that gear the same order holds because regeneration is what keeps the casts coming. Items trade several points of one for few of another, so read the list as "which stat is cheap to give up".

Spell power adds nothing to a heal. Every figure behind this ranking comes from one stated incoming-damage profile (the scaled Phase 1 tank hits and raid pulses described on the [BiS page](/bis)): it ranks gear and builds for this spec, and says nothing about which healer class is stronger.

## Gear

Follow the stat priority above; spell power on a healing item adds nothing to a heal. Water Shield's proc-based mana return makes a slightly lower spirit floor more survivable than in 1.12, but spirit remains the tree's primary sustain stat. Specific pre-raid or raid-tier item picks can't be named with confidence yet: the beta caps at level 30 and this build's raid loot tables are still missing most items the community's sourcing expects, with nothing raiding in-game until the first tier opens on 9 December 2026.

## Enchants and consumables

Weapon enchants should target healing power; **Enchant Weapon - Healing Power** is a verified entry in this build's enchant data. For consumables, **Flask of Distilled Wisdom** is a verified option historically tied to mana regeneration for healers. A fuller consumable stack isn't something this site is naming until the class is played past level 30.

## Races

Restoration can be played by Orc and Troll on the Horde, Tauren on the Horde, Dwarf — new in Forever — on the Alliance, and the Windshaper Skyborne on the Horde only. Dwarf is, for now, the only Alliance option for Shaman at all. For Alliance, **Dwarf** is the only pick; its racials lean offensive rather than healer-focused, but Stoneform's physical damage reduction still helps a Restoration Shaman survive melee pressure while casting. For Horde, **Tauren** is the strongest pick for Restoration specifically: Endurance grants 5% health and 1% hit chance, both passive survivability and consistency gains that matter more for a support caster standing in raid damage than an offensive cooldown like Orc's Blood Fury does.

## Professions

Community convention for healing Shaman favors **Herbalism** alongside **Alchemy** for self-sourced healing and mana potions, or **Tailoring** for caster cloth-adjacent leather stats, matching general Classic-era healer practice. This is community convention, not confirmed for Forever.

## Leveling

Enhancement is the generally recommended leveling spec for Shaman, not Restoration — Restoration's group-support kit doesn't translate to fast solo leveling the way Enhancement's melee durability and burst does. The beta caps at level 30, so this is carried over from 1.12 leveling knowledge rather than tested in Forever.
