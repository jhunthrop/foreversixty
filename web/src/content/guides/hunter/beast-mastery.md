---
title: Beast Mastery Hunter in Forever
classSlug: hunter
spec: beast-mastery
role: dps
build: 'FS1:1.60.1.70532:hunter:dwarf:5420001505001251/0053502001/4:'
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

The remaining points move out of a thin Marksmanship/Survival utility spread and into talents the search confirmed actually move the parse. **Efficiency** (5/5) cuts 15% off the mana cost of shots, stings, and melee abilities, paying for a shot-heavy rotation. **Trueshot Aura** (1 point) adds 30 Ranged Attack Power to the whole party for 30 minutes — a raid-wide buff on top of its personal value. **Improved Stings** (3/3) adds 20% more Serpent Sting damage, a shorter Viper Sting cooldown, and a longer Scorpid Sting duration, and **Rapid Killing** (2/2) shortens Rapid Fire's cooldown and grants a temporary damage buff to the next shot on a kill — both of these measure close to zero in the sim today because the engine doesn't yet have code paths for Serpent Sting's damage scaling or Rapid Killing's proc, so treat them as an honest unknown rather than a confirmed gain, taken because nothing else was better with those points.

To pay for that, the build drops **Hawk Eye** (ranged weapon range) and **Improved Concussive Shot** (a stun chance on Concussive Shot) entirely, and **Deflection** (parry chance) entirely — all three measured no damage benefit at all in the sim, so the points are better spent elsewhere for a pure DPS build even though they're real defensive or utility effects in a fight where parrying or kiting matters. It also gives back the fifth point in **Improved Tracking** (the per-point damage bonus against your tracked creature type), keeping it at 4/5 instead of 5/5 — that fifth point is a real, measured damage pick, just a smaller one than what replaced it, so it's a deliberate trade rather than a wasted point. This build measures about +8.5% over the previous spend in our level-60 search run. Open the planner at [/planner?class=hunter](/planner?class=hunter) to build this out.

**Raid build.** The raid-ready search found no build that beats this one beyond error while keeping what the rotation casts, the group buffs the raid already counts on and the threat and survival talents, so the raid build is the leveling build and one tree serves both.

## Rotation and priority

This site's own simulator plays Beast Mastery on the same Aimed Shot and Multi-Shot shell the other two specs share, with Bestial Wrath layered in as this spec's own cooldown. Before the pull, Aspect of the Hawk goes up and the first Aimed Shot begins casting so it lands right as the fight opens. Once engaged, Bestial Wrath is used on cooldown first and timed to land together with Rapid Fire, since a fight of Classic length only gives room for one real burst window — stacking both cooldowns into it beats spreading them apart.

## Stat priority

The table above is this level range's own simulation at level 60, re-run by the nightly pipeline every time the build or its gear data changes — these numbers are never hand-entered. A stat shown as "not separable from zero" is not a verdict that it is worthless; the sim's measured error on it is too wide, at this level range's sample size, to tell its true value apart from zero.

## Best in slot

Look for attack power and ranged attack power first, then agility, then crit and hit, in the order above.

## Enchants and consumables

Target agility and attack power on weapon and glove enchants.

## Races

Beast Mastery can be played by every race Hunter allows: Human (new in Forever), Orc, Dwarf, Night Elf, Tauren, and Troll, plus the Skyborne on either faction. For Alliance, **Dwarf** is the strongest pairing: the reworked Dwarf racial, Big Game Hunter, adds bonus damage against Beasts specifically, which lines up with how much of this spec's own damage comes from a beast-type pet and from beast-heavy leveling content. For Horde, **Troll** pairs well: Berserking grants 10% attack and casting speed on a 3-minute cooldown that stacks directly with Bestial Wrath for a bigger burst window, and Troll's own Beast Slaying racial adds another 5% damage against Beasts on top of it. Hunter is open to the Skyborne on either faction with no class restriction; their racials read as utility (a downward glide, a short full-resource restore, haste, bonus Elemental damage) rather than beast-focused, so Skyborne is a valid pick but not a clear damage upgrade over Dwarf or Troll for this spec.

## Professions

Community convention for Hunter favors **Skinning** for the cheap, plentiful leather this class already collects while killing beasts to level, often paired with **Leatherworking** for self-crafted agility gear, or **Engineering** for ranged-weapon scopes and trap-adjacent gadgets.

## Leveling

Beast Mastery is generally regarded as the strongest leveling spec for Hunter: a durable, hard-hitting pet absorbs damage that would otherwise land on the Hunter and clears trash faster than Marksmanship's single-target focus or Survival's melee weaving.
