---
title: Holy Priest in Forever
classSlug: priest
spec: holy
role: healer
build: 'FS1:1.60.1.70707:priest:gnome:02500303133/03505003030121431/0:'
recommendedRaces: [gnome, troll]
statPriority: [Healing power, Intellect, Spirit, MP5, Critical strike]
description: 'Holy Priest overview, talent priority, healing priority, stat weights, and race picks for Forever, with beta-versus-projection called out.'
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

Holy is the Priest throughput healing tree: efficient direct and periodic heals, with Prayer of Mending and Prayer of Healing carrying the group-wide healing.

### Strengths

- Prayer of Mending jumps through the group once placed on the tank
- Prayer of Healing answers several members being hurt together
- Spiritual Healing multiplies every heal
- Spiritual Guidance turns Spirit into bonus healing

### Weaknesses

- Mana is the limit, and it runs dry sooner than Discipline
- No shield or other damage prevention
- Spell damage on healing items is wasted, since it adds nothing to a heal

### In a raid

- **Role:** Raid and group healer; carries the bulk of group-wide healing
- **Brings:** Power Word: Fortitude for the group, with Prayer of Mending and Renew on the tank

## Talents and builds

The build spends its deep points in Holy and fills Discipline with the mana and spell-cost talents a healer wants. Chosen by simming candidate builds on the same gear against the same fight, not by recollection. Every figure behind this ranking comes from one stated incoming-damage profile (the scaled Phase 1 tank hits and raid pulses described on the [BiS page](/bis)): it ranks gear and builds for this spec, and says nothing about which healer class is stronger.

- **Holy Specialization, Divine Fury and Improved Renew** open the tree: more critical heals, Heal and Greater Heal cast faster, and a stronger Renew.
- **Improved Healing** makes the direct heals cheaper, and **Litany of Light** pays mana back when you switch heals instead of repeating one.
- **Holy Binding** takes the slot of Binding Heal, which the client removed: your Flash Heals on others heal you for half as much. The simulator does not model it (the fight never damages the healer), and nothing else in the rows above Litany of Light measured more healing, so it is the point that keeps you alive.
- **Spiritual Guidance** turns Spirit into bonus healing, which is why Spirit outranks its raw look on the stat list below.
- **Spiritual Healing** is a flat multiplier on every heal, and **Prayer of Mending** (with Spirit of Redemption as its prerequisite) is the capstone.
- In Discipline, **Meditation** keeps part of your regeneration running while you cast and **Mental Strength** lifts Intellect; **Twin Disciplines** and **Improved Power Word: Shield** are the points that open the rows beneath them.

On the same gear the build with Holy deepest scored ahead of the one with Discipline deepest. Open the planner at [/planner?class=priest](/planner?class=priest) to adjust it.

## Rotation and priority

The curated rotation tops up from health, tank first. **Prayer of Mending** and **Renew** go on the tank and stay there. **Flash Heal** is for emergencies only: a tank or party member who is nearly dead. **Prayer of Healing** answers several party members being hurt at once, and **Heal** is the efficient filler for a tank or member with a real hole in their health bar.

The filler and Prayer of Healing are paced against your mana: they only fire at their normal thresholds while your mana is keeping up with the share of the fight that is left, and wait for a more serious hole when it is not. Mana is the limit for this spec under the profile, so the pacing matters. The filler is Heal at its top rank.

## Stat priority

In simulator-derived order, per point of stat, the mana stats outrank raw healing: **MP5** and **Spirit** buy more healing over the fight than a point of anything else, then **Intellect**, then **healing power**, with **critical strike** last. That order follows from the fight being mana-bound: a healer who runs dry stops healing, so a point of regeneration is worth more than a point of throughput. Items still trade many points of one for few of another, so read the list as "which stat is cheap to give up", not "which stat to hunt".

Spell power, the damage side of a healing item, adds nothing to a heal. Forever puts a third of every healing bonus on an item as spell damage; it is the healing number that heals. Every figure behind this ranking comes from one stated incoming-damage profile (the scaled Phase 1 tank hits and raid pulses described on the [BiS page](/bis)): it ranks gear and builds for this spec, and says nothing about which healer class is stronger.

## Best in slot

Spirit stays the tree's core sustain stat.

## Enchants and consumables

## Races

## Professions

Community convention for healing Priests favors **Tailoring** for caster cloth stats alongside **Enchanting** for weapon and gear enchants, since neither profession competes with anything Holy-specific.

## Leveling

Shadow is the generally recommended leveling spec for Priest, not Holy — Holy's group-throughput kit doesn't translate to fast solo leveling the way Shadow's self-sufficient damage and healing does.
