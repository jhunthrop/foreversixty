---
title: Beast Mastery Hunter in Forever
classSlug: hunter
spec: beast-mastery
role: dps
build: 'FS1:1.60.1.70658:hunter:dwarf:5420001505001251/0053502001/4:'
recommendedRaces: [dwarf, troll]
statPriority: [Agility, Critical strike, Ranged attack power, Hit, Melee haste]
description: 'Beast Mastery Hunter overview, talent priority, rotation, stat weights, and race picks for Forever, with beta-versus-projection called out.'
updated: 2026-09-24
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

Beast Mastery is the pet-focused ranged Hunter spec: a strengthened pet deals a large share of the damage, and Bestial Wrath is its burst cooldown.

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

To pay for that, the build drops **Hawk Eye** (ranged weapon range) and **Improved Concussive Shot** (a stun chance on Concussive Shot) entirely, and **Deflection** (parry chance) entirely — all three measured no damage benefit at all in the sim, so the points are better spent elsewhere for a pure DPS build even though they're real defensive or utility effects in a fight where parrying or kiting matters. It also gives back the fifth point in **Improved Tracking** (the per-point damage bonus against your tracked creature type), keeping it at 4/5 instead of 5/5 — that fifth point is a real, measured damage pick, just a smaller one than what replaced it, so it's a deliberate trade rather than a wasted point. This build measures about +8.5% over the previous spend in our level-60 search run. Open the planner at [/planner?class=hunter](/planner?class=hunter) to build this out.

## Rotation and priority

This site's own simulator plays Beast Mastery on the same Aimed Shot and Multi-Shot shell the other two specs share, with Bestial Wrath layered in as this spec's own cooldown. Before the pull, Aspect of the Hawk goes up and the first Aimed Shot begins casting so it lands right as the fight opens. Once engaged, Bestial Wrath is used on cooldown first and timed to land together with Rapid Fire, since a fight of Classic length only gives room for one real burst window — stacking both cooldowns into it beats spreading them apart. Aimed Shot is then recast every time it comes off cooldown, as long as the cast won't clip the next auto shot. Multi-Shot comes next on every free global between Aimed Shot casts, then Serpent Sting is refreshed whenever it's close to falling off and neither Rapid Fire nor Aimed Shot is close to ready (putting Multi-Shot ahead of the sting was a small gain, about +0.4% in the raid-ready run, and level for the bare character), and Arcane Shot is the last filler whenever Aimed Shot itself is more than 1.5 seconds from ready and mana is above a 25% reserve — Multi-Shot's own cooldown comes up often enough that gating Arcane Shot on it too would zero out Arcane Shot's casts entirely, so the priority no longer waits on it. Arcane Shot is instant, so it never clips an auto shot, but it's the least mana-efficient shot the spec has, hence the reserve floor. The priority list doesn't call for active pet repositioning beyond keeping the pet alive and attacking through the Bestial Wrath window.

## Stat priority

## Best in slot

## Enchants and consumables

## Races

## Professions

Community convention for Hunter favors **Skinning** for the cheap, plentiful leather this class already collects while killing beasts to level, often paired with **Leatherworking** for self-crafted agility gear, or **Engineering** for ranged-weapon scopes and trap-adjacent gadgets.

## Leveling

Beast Mastery is generally regarded as the strongest leveling spec for Hunter: a durable, hard-hitting pet absorbs damage that would otherwise land on the Hunter and clears trash faster than Marksmanship's single-target focus or Survival's melee weaving.
