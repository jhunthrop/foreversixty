---
title: Marksmanship Hunter in Forever
classSlug: hunter
spec: marksmanship
role: dps
build: 'FS1:1.60.1.70532:hunter:dwarf:5320000501/005155000150305/5:'
recommendedRaces: [dwarf, troll]
statPriority: [Agility, Critical strike, Ranged attack power, Hit, Melee haste]
description: 'Marksmanship Hunter overview, talent priority, rotation, stat weights, and race picks for Forever, with beta-versus-projection called out.'
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
  - label: 'Warcraft Tavern, Marksmanship Hunter rotation guide'
    url: https://www.warcrafttavern.com/wow-classic/guides/pve-marksmanship-hunter-rotations-cooldowns/
    kind: community
---

## Overview

Marksmanship is the Hunter ranged spec built on Aimed Shot, the hardest single shot of the three trees, boosted by Careful Aim and Mortal Shots.

### Strengths

- Aimed Shot is the highest single-shot damage of the three Hunter trees
- Careful Aim adds Attack Power from Intellect
- Mortal Shots adds critical damage to ranged abilities
- Lethal Attacks and Ranged Weapon Specialization add flat crit chance and damage

### Weaknesses

- Stationary: the whole rotation is cast at range from one spot
- Aimed Shot's slow cast must not clip the next auto shot
- Arcane Shot is its least mana-efficient filler

### In a raid

- **Role:** Stationary ranged DPS with high uptime
- **Brings:** Trueshot Aura, ranged attack power for the party

## Talents and builds

Reading the Marksmanship tree in priority order for a single-target Aimed Shot build: **Careful Aim** (adds Attack Power equal to a percentage of Intellect) and **Mortal Shots** (bonus critical strike damage on ranged abilities) come first, since both are flat multipliers on every shot that follows. **Trueshot Aura** is next: in Forever it grants ranged attack power only to the party, a narrower version than 1.12's melee-and-ranged buff per the demo notes. **Barrage** adds a damage bonus specifically to Multi-Shot, Aimed Shot, and Volley — all abilities this spec's rotation already leans on. **Efficiency** helps sustain the mana cost of a shot-heavy rotation. **Improved Stings** (1/3) adds 6% to Serpent Sting, which this rotation keeps up throughout.

On top of that spread, this build reaches further into Beast Mastery and down to the bottom of its own tree for three more measured damage gains. **Unleashed Fury** (5/5) adds 15% more pet and hawk damage and **Ferocity** (1/5) adds 2% pet and hawk critical strike chance. **Lethal Attacks** (5/5) adds a flat 5% critical strike chance with all attacks, and **Ranged Weapon Specialization** (5/5) adds a flat 5% more damage with ranged weapons — both are the kind of unconditional multiplier that's hard to pass up once the points are available.

The last round of the talent search moved three points: two from **Endurance Training** (pet health and armor, now 3/5) and the one in **Lone Wolf** (20% more damage with no active pet) went to the fifth point of Unleashed Fury, the first of Ferocity and the first of Improved Stings. Endurance Training and Lone Wolf both measured no damage in the sim, because a stationary single-target parse shows neither pet durability nor a pet-less bonus this build uses; they are real effects, just not ones the parse can show value from. The build keeps Improved Aspect of the Monkey, Hawk Eye, Improved Concussive Shot and Deflection at zero, as before. This build measures about +1.5% over the previous spend in our level-60 search run. Open the planner at [/planner?class=hunter](/planner?class=hunter) to build this out.

**Raid build.** The raid-ready search found no build that beats this one beyond error while keeping what the rotation casts, the group buffs the raid already counts on and the threat and survival talents, so the raid build is the leveling build and one tree serves both.

## Rotation and priority

This site's own simulator plays Marksmanship as the most direct read of the shared Hunter shot rotation: heavy Aimed Shot weaving, with no Serpent Sting: the rotation search measured the sting as a loss in the raid-ready run, about +1.2% for dropping it, because the global and the mana are worth more to Aimed Shot and Multi-Shot (the engine's Serpent Sting damage is held to the client's numbers, so this is the sim's honest verdict, and the bare character comes out about level).

## Stat priority

The table above is this level range's own simulation at level 60, re-run by the nightly pipeline every time the build or its gear data changes — these numbers are never hand-entered. A stat shown as "not separable from zero" is not a verdict that it is worthless; the sim's measured error on it is too wide, at this level range's sample size, to tell its true value apart from zero.

## Best in slot

Look for attack power and ranged attack power first, then agility, then crit and hit, in the order above.

## Enchants and consumables

Target agility and attack power on weapon and glove enchants.

## Races

Marksmanship can be played by every race Hunter allows: Human (new in Forever), Orc, Dwarf, Night Elf, Tauren, and Troll, plus the Skyborne on either faction. For Alliance, **Dwarf** is the strongest pairing: the reworked Dwarf racial, Big Game Hunter, adds bonus damage against Beasts specifically, which applies to Aimed Shot the same as any other physical attack. For Horde, **Troll** pairs well and pays off a little differently for this spec than for the other two: Berserking's 10% casting speed applies directly to Aimed Shot's own cast time, shortening the one ability Marksmanship revolves around, on top of the 10% attack speed and Beast Slaying's 5% damage against Beasts. Hunter is open to the Skyborne on either faction with no class restriction; their racials read as utility (a downward glide, a short full-resource restore, haste, bonus Elemental damage) rather than damage-focused, so Skyborne is a valid pick but not a clear upgrade over Dwarf or Troll for this spec.

## Professions

Community convention for Hunter favors **Skinning** for the cheap, plentiful leather this class already collects while killing beasts to level, often paired with **Leatherworking** for self-crafted agility gear, or **Engineering** for ranged-weapon scopes and trap-adjacent gadgets.

## Leveling

Beast Mastery is generally regarded as the stronger leveling spec for Hunter overall; Marksmanship's single-target focus doesn't clear trash as quickly without a pet built up the same way.
