---
title: Beast Mastery Hunter in Forever
classSlug: hunter
spec: beast-mastery
role: dps
build: 'FS1:1.60.1.70658:hunter:dwarf:5420001505001251/0053542001/0:'
recommendedRaces: [dwarf, troll]
statPriority: [Agility, Critical strike, Ranged attack power, Hit, Attack speed]
description: 'Beast Mastery Hunter overview, talent priority, rotation, stat weights, and race picks for Forever, with beta-versus-projection called out.'
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
  - label: 'Warcraft Tavern, Beast Mastery Hunter rotation guide'
    url: https://www.warcrafttavern.com/wow-classic/guides/pve-beast-mastery-hunter-rotations-cooldowns/
    kind: community
---

## Overview

Beast Mastery is the pet-focused ranged Hunter spec: the pet, drawing its stats from the Hunter's gear, deals much of the damage, and Bestial Wrath is its burst.

### Strengths

- A large share of damage comes from the pet
- Bestial Wrath gives the pet 50% more damage for 18 seconds
- The pet absorbs hits and threat the Hunter would otherwise take
- Efficiency cuts shot and sting mana cost by 15%

### Weaknesses

- Damage depends on the pet staying alive and attacking
- A fight of Classic length allows only one real burst window
- Arcane Shot is its least mana-efficient filler

### In a raid

- **Role:** Ranged single-target DPS with a pet
- **Brings:** Trueshot Aura, ranged attack power for the party

## Talents and builds

Reading the Beast Mastery tree in priority order for a pet-damage build: **Unleashed Fury** (a flat damage increase for pets and hawks) and **Ferocity** (pet critical strike chance) come first, since both scale every hit the pet lands afterward. **Focused Fire** follows — its damage bonus applies to the Hunter as well as the pet, unlike the two talents above it. **Frenzy** gives the pet a chance at bonus attack speed off its own critical strikes, compounding with Ferocity. **Bestial Discipline** helps mana sustain by letting some regeneration continue through casting. The tree's capstone, **Bestial Wrath**, is the spec's defining cooldown: an 18-second window of 50% additional pet damage.

To pay for that, the build drops **Hawk Eye** (ranged weapon range) and **Improved Concussive Shot** (a stun chance on Concussive Shot) entirely, and **Deflection** (parry chance) entirely — all three measured no damage benefit at all in the sim, so the points are better spent elsewhere for a pure DPS build even though they're real defensive or utility effects in a fight where parrying or kiting matters. The last four points go to **Careful Aim** (4/5) in Marksmanship, which turns Intellect into melee and ranged attack power, instead of **Improved Tracking**: about +2% in our level-60 search run. In Forever the pet takes a tenth of your attack power and all of your critical strike chance, and it gets no raid buffs of its own, so the Hunter's own gear and talents now raise the pet's damage too. Open the planner at [/planner?class=hunter](/planner?class=hunter) to build this out.

## Rotation and priority

This site's own simulator plays Beast Mastery on the same shot shell as Marksmanship, with Bestial Wrath layered in as this spec's own cooldown. In Forever, Multi-Shot and Aimed Shot share one 6-second cooldown. Instant Multi-Shot does more over a fight than Aimed Shot, whose cast pauses the auto shot, so Aimed Shot is only the opener: before the pull, Aspect of the Hawk goes up and the first Aimed Shot begins casting so it lands right as the fight opens. Once engaged, Bestial Wrath is used on cooldown first and timed to land together with Rapid Fire, since a fight of Classic length only gives room for one real burst window — stacking both cooldowns into it beats spreading them apart. Multi-Shot is then cast every time the shared cooldown comes up. Serpent Sting is refreshed whenever it's close to falling off and Multi-Shot is not about to come back, and Arcane Shot, which has its own cooldown, is the last filler whenever Multi-Shot is more than 1.5 seconds from ready and mana is above a 25% reserve. Arcane Shot is instant, so it never clips an auto shot, but it's the least mana-efficient shot the spec has, hence the reserve floor. The priority list doesn't call for active pet repositioning beyond keeping the pet alive and attacking through the Bestial Wrath window.

## Stat priority

## Best in slot

## Enchants and consumables

## Races

## Professions

Community convention for Hunter favors **Skinning** for the cheap, plentiful leather this class already collects while killing beasts to level, often paired with **Leatherworking** for self-crafted agility gear, or **Engineering** for ranged-weapon scopes and trap-adjacent gadgets.

## Leveling

Beast Mastery is generally regarded as the strongest leveling spec for Hunter: a durable, hard-hitting pet absorbs damage that would otherwise land on the Hunter and clears trash faster than Marksmanship's single-target focus or Survival's melee weaving.
