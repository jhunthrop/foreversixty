---
title: Survival Hunter in Forever
classSlug: hunter
spec: survival
role: dps
build: 'FS1:1.60.1.70532:hunter:dwarf:0/32500500005/500230131051120151:'
recommendedRaces: [dwarf, troll]
statPriority: [Agility, Attack power, Strength, Critical strike, Hit, Melee haste]
description: 'Survival Hunter overview, talent priority, rotation, stat weights, and race picks for Forever, with beta-versus-projection called out.'
updated: 2026-10-07
confidence: inferred
sources:
  - label: 'Blizzard, Deep Dive panel recap'
    url: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
    kind: blizzard
  - label: 'Talents Forever (demo transcription)'
    url: https://talentsforever.com/data.json
    kind: community
  - label: 'Forever Sixty build data, data/builds/1.60.1.70374/talents/hunter.json'
    url: https://foreversixty.gg/data/1.60.1.70532/talents/hunter.json
    kind: datamined
  - label: 'Forever Sixty build data, data/builds/1.60.1.70374/spellconst/hunter.json'
    url: https://foreversixty.gg/data/1.60.1.70532/spellconst/hunter.json
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
  - label: 'wowhead Forever, Mongoose Bite and Counterattack tooltips'
    url: https://www.wowhead.com/forever/spell=1495
    kind: community
  - label: 'wowsims-forever default hunter priority list'
    url: https://github.com/wowsims/classic/blob/master/ui/hunter/apls/p1.apl.json
    kind: community
  - label: 'Warcraft Tavern, Survival Hunter rotation guide'
    url: https://www.warcrafttavern.com/wow-classic/guides/pve-survival-hunter-rotations-cooldowns/
    kind: community
---

## Overview

Survival is a melee spec in Forever: the Hunter stands in weapon range for the whole fight and uses Raptor Strike, Mongoose Bite and bleeds instead of shots.

### Strengths

- Counterattack is a free strike that cannot be blocked, dodged or parried
- Lacerating Strikes adds a bleed to Mongoose Bite
- Raptor Strike replaces the next auto-attack without costing a swing
- Strider Kick adds a third melee special on a short cooldown

### Weaknesses

- Stays in melee range and fires no ranged shots
- Counterattack and Mongoose Bite fire only when their windows open
- Wing Clip filler pays off only with a Windfury Totem in the group

### In a raid

- **Role:** Melee DPS
- **Brings:** Hunter's Mark on the target; no raid-wide buff

## Talents and builds

Reading the Survival tree in priority order for a melee build: **Savage Strikes** and **Predator's Edge** are direct melee damage — crit chance and crit damage on every melee special, plus a bonus to off-hand damage if this spec carries a second weapon. **Deterrence** unlocks **Counterattack**, a free, unavoidable strike after a parry (cannot be blocked, dodged, or parried) — a genuine "free damage" proc this build takes specifically to use. **Expose Prey** is a damage talent in its own right: each melee or ranged attack against a target carrying your Hunter's Mark has a 5% (10% at 2/2) chance to activate Mongoose Bite for 5 seconds, and that window is the only way a boss that never attacks you lets you cast it at all — the two points measure about +7% in our level-60 search run, so keep Hunter's Mark up. It also unlocks the capstone, **Lacerating Strikes**, which adds a bleed to Mongoose Bite worth 40% of that hit's damage over 21 seconds. **Strider Kick** is a third melee special on its own short cooldown. **Improved Wing Clip** and **Entrapment** are left at 0 here: they raise the reliability of control tools this rotation doesn't use against a stationary target, and the points are worth more spent reaching Counterattack and Lacerating Strikes sooner. The remaining points go into Marksmanship, spent reaching **Mortal Shots** (its own prerequisite, **Careful Aim**, needs all 5 points) — in this engine, Mortal Shots adds its crit-damage bonus to every Hunter special, melee included, rather than into Beast Mastery or further Survival utility. The five points once held in **Efficiency** (shot mana cost, which a melee rotation that never fires a shot does not use) now sit in **Lethal Attacks** (5/5), a flat 5% critical strike chance on all attacks, and that swap measures about +5% in our level-60 search run. Open the planner at [/planner?class=hunter](/planner?class=hunter) to build this out.

**Raid build.** The raid-ready search found no build that beats this one beyond error while keeping what the rotation casts, the group buffs the raid already counts on and the threat and survival talents, so the raid build is the leveling build and one tree serves both.

## Rotation and priority

This site's simulator starts Survival in melee range and keeps it there for the whole fight: Raptor Strike is queued on the swing timer (it replaces the next melee auto-attack rather than costing its own swing) with no mana gate at all; an earlier version held it back under 20% mana to protect Mongoose Bite, and the search measured that gate as a small loss. Ranged shots (Aimed Shot, Multi-Shot, Arcane Shot) and Serpent Sting are not part of this rotation at all: they require standing outside melee range, which this build never does.

## Stat priority

The table above is this level range's own simulation at level 60, re-run by the nightly pipeline every time the build or its gear data changes — these numbers are never hand-entered. A stat shown as "not separable from zero" is not a verdict that it is worthless; the sim's measured error on it is too wide, at this level range's sample size, to tell its true value apart from zero.

## Best in slot

Look for attack power first, then agility and strength, then crit and hit, in the order above. Prioritize weapon slots: a main-hand (and off-hand, if this build carries one — Predator's Edge's off-hand damage bonus only matters with one equipped) over a ranged weapon, which this rotation no longer uses for anything but a stat stick.

## Enchants and consumables

Target agility and attack power on weapon and glove enchants — the weapon enchant matters most here, since this spec's melee weapon (or weapons, if dual-wielding) carries the whole rotation now, not a ranged one.

## Races

Survival can be played by every race Hunter allows: Human (new in Forever), Orc, Dwarf, Night Elf, Tauren, and Troll, plus the Skyborne on either faction. For Alliance, **Dwarf** is the strongest pairing: the reworked Dwarf racial, Big Game Hunter, adds bonus damage against Beasts specifically, which applies to every melee special here against a beast target exactly as it would to a shot. For Horde, **Troll** pairs well: Berserking's 10% attack speed shortens the melee swing timer that now drives this entire rotation, not just the Raptor Strike weave, and Beast Slaying's 5% damage against Beasts stacks on top of it. Hunter is open to the Skyborne on either faction with no class restriction; their racials read as utility (a downward glide, a short full-resource restore, haste, bonus Elemental damage) rather than damage-focused, so Skyborne is a valid pick but not a clear upgrade over Dwarf or Troll for this spec.

## Professions

Community convention for Hunter favors **Skinning** for the cheap, plentiful leather this class already collects while killing beasts to level, often paired with **Leatherworking** for self-crafted agility gear, or **Blacksmithing** for a melee-focused Hunter's own weapon and armor work.

## Leveling

Beast Mastery is generally regarded as the strongest leveling spec for Hunter; a melee-committed Survival needs deeper Survival investment (Deterrence for Counterattack, Expose Prey for Lacerating Strikes) than a leveling character can usually afford early, so it comes online later than Beast Mastery does, and later than it did as a part-ranged spec.
