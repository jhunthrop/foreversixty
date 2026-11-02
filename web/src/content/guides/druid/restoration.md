---
title: Restoration Druid in Forever
classSlug: druid
spec: restoration
role: healer
build: 'FS1:1.60.1.70707:druid:night-elf:45322001/0/50530051531132:'
recommendedRaces: [night-elf, tauren]
statPriority: [Healing power, Intellect, Spirit, MP5, Critical strike]
description: 'Talents, rotation, stats, and gear for Restoration Druid in Forever, and what is confirmed versus projected from the beta.'
updated: 2026-10-07
confidence: inferred
sources:
  - label: 'Blizzard, Deep Dive panel recap'
    url: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
    kind: blizzard
  - label: 'Talents Forever (demo transcription)'
    url: https://talentsforever.com/data.json
    kind: community
  - label: 'Forever Sixty stat weights, data/curated/stat-weights.json'
    url: https://foreversixty.gg/sources
    kind: site
  - label: 'Forever Sixty raid loot data, data/curated/loot/forever-raid-phases.json'
    url: https://foreversixty.gg/sources
    kind: site
---

## Overview

Restoration is the Druid healer tree: Rejuvenation and Regrowth on many targets, Healing Touch for real damage, and mana efficiency to last a long fight.

### Strengths

- Heals over time on several targets keep a group healthy at low cost
- Mana efficiency talents let it out-sustain other healers over a long fight
- Swiftmend turns a running heal over time into an instant heal
- Nature's Swiftness into Healing Touch answers a nearly dead tank

### Weaknesses

- Heals over time are slow against sudden burst damage
- Healing Touch is held back so mana lasts the fight
- Spell power adds nothing to a heal; gear needs healing power

### In a raid

- **Role:** Healer, keeping the tank up with heals over time and group healing on pulses
- **Brings:** Mark of the Wild for the group and Innervate for casters

## Talents and builds

The build spends most of its points in Restoration and fills Balance with the periodic-healing and critical-chance talents. Chosen by simming candidate builds on the same gear against the same fight. Every figure behind this ranking comes from one stated incoming-damage profile (the scaled Phase 1 tank hits and raid pulses described on the [BiS page](/bis)): it ranks gear and builds for this spec, and says nothing about which healer class is stronger.

- **Naturalist** shortens Healing Touch, **Gift of Nature** strengthens every heal, and **Improved Rejuvenation** strengthens Rejuvenation.
- **Reflection** keeps part of your regeneration running while you cast, **Tranquil Spirit** makes Healing Touch and Tranquility cheaper, and **Living Spirit** lifts Spirit.
- **Bloom** (1 point, new in client 1.60.1.70707) heals the target for your Spirit whenever a Rejuvenation or Regrowth runs its full duration; with heals over time doing most of the work it was the largest single gain the sim found for this build, about +10% healing over the alternatives for that point. It took the point of Gift of the Earthmother, which left the tree in the same patch.
- **Swiftmend** (it eats the heal over time it reads) and **Nature's Swiftness** are the burst tools.
- From Balance, **Genesis** adds 10% to every heal over time, **Nature's Majesty** adds critical chance, and **Nature's Splendor** lengthens your heal-over-time spells.

The build leaves out the capstone Wild Growth and Improved Regrowth: the sim measured both against spending those points elsewhere, and under this profile the points did better elsewhere. That is a statement about this rotation's mana pacing, not about the spells; since 1.60.1.70707 Wild Growth also has a 3 second cast. Measured again in round 6, with the raid pulses now reaching you too, the best Wild Growth line came out level with this rotation, within error, not ahead of it: Rejuvenation with Bloom heals more for its mana, and this fight runs your mana down before it runs out your time. In a fight with heavy group damage in bursts, take it. Nature's Focus stays over Heart of the Wild (10% more Intellect, now on Restoration's top row): the raid's damage reaches you while you cast, and Nature's Focus keeps Healing Touch, Regrowth and Tranquility from losing casting time to it, which measured level with the Intellect. Open the planner at [/planner?class=druid](/planner?class=druid) to adjust it.

## Rotation and priority

The tank keeps **Regrowth** and **Rejuvenation** up for the whole fight, and the heals over time carry most of the work. **Healing Touch** fires only on a real hole in the tank's health and only while your mana is keeping up with the share of the fight that is left, so the mana lasts instead of running dry at the first long fight. A hurt party member gets **Rejuvenation**, you included, and **Tranquility** goes out when two of the five are well down. **Nature's Swiftness** into an instant Healing Touch is the emergency answer for a tank who is nearly dead, and **Innervate** goes on you when you are low.

The rotation does not downrank.

## Stat priority

In simulator-derived order, per point of stat: **MP5** first, **Spirit** behind it, then **healing power**, **Intellect** and **critical strike**. Spirit also reaches the Living Spirit bonus. Items trade several points of one for few of another, so read the list as "which stat is cheap to give up".

Spell power adds nothing to a heal. Every figure behind this ranking comes from one stated incoming-damage profile (the scaled Phase 1 tank hits and raid pulses described on the [BiS page](/bis)): it ranks gear and builds for this spec, and says nothing about which healer class is stronger.

## Best in slot

## Enchants and consumables

## Races

## Professions

Herbalism and Alchemy is the general Classic-era pairing for a Restoration healer, giving access to self-crafted mana and healing consumables.

## Leveling

Feral is the stronger leveling choice for Druid overall, and Restoration is generally considered the slowest of the three specs to solo with, since its kit is built around sustaining others rather than dealing damage.
