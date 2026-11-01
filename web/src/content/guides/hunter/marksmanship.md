---
title: Marksmanship Hunter in Forever
classSlug: hunter
spec: marksmanship
role: dps
build: 'FS1:1.60.1.70707:hunter:dwarf:5320000501/005155000150305/5:'
recommendedRaces: [dwarf, troll]
statPriority: [Agility, Critical strike, Ranged attack power, Hit, Attack speed]
description: 'Marksmanship Hunter overview, talent priority, rotation, stat weights, and race picks for Forever, with beta-versus-projection called out.'
updated: 2026-10-31
confidence: inferred
sources:
  - label: 'Blizzard, Deep Dive panel recap'
    url: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
    kind: blizzard
  - label: 'Talents Forever (demo transcription)'
    url: https://talentsforever.com/data.json
    kind: community
  - label: 'Forever Sixty build data, data/builds/1.60.1.69893/talents/hunter.json'
    url: https://foreversixty.gg/data/1.60.1.69893/talents/hunter.json
    kind: datamined
  - label: 'Forever Sixty build data, data/builds/1.60.1.69893/combos.json'
    url: https://foreversixty.gg/data/1.60.1.69893/combos.json
    kind: datamined
  - label: 'Forever Sixty build data, data/builds/1.60.1.69893/races.json'
    url: https://foreversixty.gg/data/1.60.1.69893/races.json
    kind: datamined
  - label: 'Forever Sixty build data, data/builds/1.60.1.69893/enchants.json'
    url: https://foreversixty.gg/data/1.60.1.69893/enchants.json
    kind: datamined
  - label: 'Forever Sixty build data, data/builds/1.60.1.69893/simconsumes.json'
    url: https://foreversixty.gg/data/1.60.1.69893/simconsumes.json
    kind: datamined
  - label: 'This site, stat weights'
    url: https://foreversixty.gg/sim/weights
    kind: site
  - label: 'This site, default rotation'
    url: https://foreversixty.gg/sim/specs
    kind: site
  - label: 'This site, character planner'
    url: https://foreversixty.gg/planner
    kind: site
  - label: 'wowsims-forever default hunter priority list'
    url: https://github.com/wowsims/classic/blob/master/ui/hunter/apls/p1.apl.json
    kind: community
  - label: 'Warcraft Tavern, Marksmanship Hunter rotation guide'
    url: https://www.warcrafttavern.com/wow-classic/guides/pve-marksmanship-hunter-rotations-cooldowns/
    kind: community
---

## Overview

Marksmanship is the Hunter ranged spec built on the shots themselves: Multi-Shot on its cooldown, an Aimed Shot opener, boosted by Careful Aim and Mortal Shots.

### Strengths

- Barrage and Mortal Shots both raise Multi-Shot, the shot this rotation casts most
- Careful Aim adds melee and ranged attack power from Intellect
- Mortal Shots adds critical damage to ranged abilities
- Lethal Attacks and Ranged Weapon Specialization add flat crit chance and damage

### Weaknesses

- Stationary: the whole rotation is cast at range from one spot
- Multi-Shot and Aimed Shot share one cooldown, so Aimed Shot only opens the fight
- Arcane Shot is its least mana-efficient filler

### In a raid

- **Role:** Stationary ranged DPS with high uptime
- **Brings:** Trueshot Aura, ranged attack power for the party

## Talents and builds

Reading the Marksmanship tree in priority order for a single-target shot build: **Careful Aim** (adds melee and ranged attack power equal to a percentage of Intellect) and **Mortal Shots** (bonus critical strike damage on ranged abilities) come first, since both are flat multipliers on every shot that follows. **Trueshot Aura** is next: in Forever it grants ranged attack power only to the party, a narrower version than 1.12's melee-and-ranged buff per the demo notes. **Barrage** adds a damage bonus specifically to Multi-Shot, Aimed Shot, and Volley — all abilities this spec's rotation already leans on. **Efficiency** helps sustain the mana cost of a shot-heavy rotation. **Improved Stings** (1/3) is there only to open the tier that holds **Ranged Weapon Specialization**: this rotation casts no Serpent Sting, so the point buys nothing else.

On top of that spread, this build reaches further into Beast Mastery and down to the bottom of its own tree for three more measured damage gains. **Unleashed Fury** (5/5) adds 15% more pet and hawk damage and **Ferocity** (1/5) adds 2% pet and hawk critical strike chance. **Lethal Attacks** (5/5) adds a flat 5% critical strike chance with all attacks, and **Ranged Weapon Specialization** (5/5) adds a flat 5% more damage with ranged weapons — both are the kind of unconditional multiplier that's hard to pass up once the points are available.

## Rotation and priority

This site's own simulator plays Marksmanship as the most direct read of the shared Hunter shot rotation, with no Serpent Sting: the rotation search measured the sting as a loss in the raid-ready run, about +1.2% for dropping it, because the global and the mana are worth more to the shots (the engine's Serpent Sting damage is held to the client's numbers, so this is the sim's honest verdict, and the bare character comes out about level). In Forever, Multi-Shot and Aimed Shot share one 6-second cooldown. Aimed Shot's 2-second cast pauses the auto shot while instant Multi-Shot does not, so on the shared cooldown Multi-Shot does more over a fight, and Aimed Shot is the opener only: before the pull, Aspect of the Hawk goes up and the opening Aimed Shot begins casting so it lands close to the start of the fight. In combat, Rapid Fire is timed to land together with an incoming auto shot right as Multi-Shot is about to come off cooldown, so both benefit from the same window. Multi-Shot is then cast every time the shared cooldown comes up, and Arcane Shot, which has its own cooldown, is the last filler whenever Multi-Shot is more than 1.5 seconds from ready and mana is above a 25% reserve. Arcane Shot is instant, so it never clips an auto shot, but it's the least mana-efficient shot the spec has, hence the reserve floor. There's no melee or pet-management component to this priority list — the spec stays at range for the whole fight.

## Stat priority

## Best in slot

## Enchants and consumables

## Races

## Professions

Community convention for Hunter favors **Skinning** for the cheap, plentiful leather this class already collects while killing beasts to level, often paired with **Leatherworking** for self-crafted agility gear, or **Engineering** for ranged-weapon scopes and trap-adjacent gadgets.

## Leveling

Beast Mastery is generally regarded as the stronger leveling spec for Hunter overall; Marksmanship's single-target focus doesn't clear trash as quickly without a pet built up the same way.
