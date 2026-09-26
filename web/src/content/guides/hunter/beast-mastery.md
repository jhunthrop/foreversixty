---
title: Beast Mastery Hunter in Forever
classSlug: hunter
spec: beast-mastery
role: dps
build: 'FS1:1.60.1.69893:hunter:dwarf:5420001505001251/3551/51:'
recommendedRaces: [dwarf, troll]
statPriority: [Attack power, 'Ranged attack power', Agility, Critical strike, Hit, Melee haste]
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

Beast Mastery is a pet-focused ranged DPS spec: a large share of total damage in Forever comes from a strengthened pet rather than the Hunter's own shots, with Bestial Wrath as the spec's signature burst cooldown. In a raid or group it plays consistent single-target damage with a pet that can absorb hits and threat the Hunter would otherwise take directly. The beta caps at level 30, so anything here about level 60 play, raid tuning, or how Beast Mastery compares to the other two specs at endgame is a projection from the demo trees and 1.12 knowledge, not confirmed play.

## Talents and builds

Reading the Beast Mastery tree in priority order for a pet-damage build: **Unleashed Fury** (a flat damage increase for pets and hawks) and **Ferocity** (pet critical strike chance) come first, since both scale every hit the pet lands afterward. **Focused Fire** follows — its damage bonus applies to the Hunter as well as the pet, unlike the two talents above it. **Frenzy** gives the pet a chance at bonus attack speed off its own critical strikes, compounding with Ferocity. **Bestial Discipline** helps mana sustain by letting some regeneration continue through casting. The tree's capstone, **Bestial Wrath**, is the spec's defining cooldown: an 18-second window of 50% additional pet damage. Getting to that capstone takes a heavy investment — roughly 30 or more points are needed to unlock the sixth row of the tree — leaving the remainder split as a handful of points into Marksmanship for Aimed Shot support and a few utility points into Survival. Open the planner at [/planner?class=hunter](/planner?class=hunter) to build this out.

## Rotation and priority

This site's own simulator plays Beast Mastery on the same Aimed Shot and Multi-Shot shell the other two specs share, with Bestial Wrath layered in as this spec's own cooldown. Before the pull, Aspect of the Hawk goes up and the first Aimed Shot begins casting so it lands right as the fight opens. Once engaged, Bestial Wrath is used on cooldown first and timed to land together with Rapid Fire, since a fight of Classic length only gives room for one real burst window — stacking both cooldowns into it beats spreading them apart. Aimed Shot is then recast every time it comes off cooldown, as long as the cast won't clip the next auto shot. Serpent Sting is refreshed whenever it's close to falling off and neither Rapid Fire nor Aimed Shot is close to ready, and Multi-Shot fills the remaining gaps between Aimed Shot casts. The priority list doesn't call for active pet repositioning beyond keeping the pet alive and attacking through the Bestial Wrath window.

## Stat priority

In simulator-derived priority order: **attack power** (the core scalar behind auto shots and every ability in this rotation), **ranged attack power** specifically (the ranged-only component that stacks on top of general attack power), **agility** (adds both attack power and ranged crit indirectly), **crit** (extra shot damage, and, through Ferocity and Frenzy, more value out of the pet), **hit** (misses cost both shot uptime and pet-buffing talent value), and **melee haste** last, since this spec's damage is mostly ranged and pet-driven rather than built around the Hunter's own melee swing.

## Gear

Look for attack power and ranged attack power first, then agility, then crit and hit, in the order above. Specific pre-raid or raid-tier item picks can't be named with confidence yet: the beta caps at level 30, and this build's raid loot tables are themselves incomplete — most items the community's own sourcing expects are missing from the current client, and nothing raids in-game until the first tier opens on 9 December 2026. This section fills in once raid loot is itemized and the spec is played past 30.

## Enchants and consumables

Target agility and attack power on weapon and glove enchants. **Enchant Gloves - Agility** and **Enchant Weapon - Agility** both exist in this build's enchant data, and a ranged weapon scope such as **Deadly Scope** or **Sniper Scope** is a verified item in this build. For consumables, **Elixir of Greater Agility** and **Elixir of the Mongoose** are both verified options in this build's consumable list. Beyond naming those, specific best-in-slot consumable stacking isn't something this site can confirm yet at level 30.

## Races

Beast Mastery can be played by every race Hunter allows: Human (new in Forever), Orc, Dwarf, Night Elf, Tauren, and Troll, plus the Skyborne on either faction. For Alliance, **Dwarf** is the strongest pairing: the reworked Dwarf racial, Big Game Hunter, adds bonus damage against Beasts specifically, which lines up with how much of this spec's own damage comes from a beast-type pet and from beast-heavy leveling content. For Horde, **Troll** pairs well: Berserking grants 10% attack and casting speed on a 3-minute cooldown that stacks directly with Bestial Wrath for a bigger burst window, and Troll's own Beast Slaying racial adds another 5% damage against Beasts on top of it. Hunter is open to the Skyborne on either faction with no class restriction; their racials read as utility (a downward glide, a short full-resource restore, haste, bonus Elemental damage) rather than beast-focused, so Skyborne is a valid pick but not a clear damage upgrade over Dwarf or Troll for this spec.

## Professions

Community convention for Hunter favors **Skinning** for the cheap, plentiful leather this class already collects while killing beasts to level, often paired with **Leatherworking** for self-crafted agility gear, or **Engineering** for ranged-weapon scopes and trap-adjacent gadgets. This is general Classic-era community practice, not something confirmed for Forever specifically.

## Leveling

Beast Mastery is generally regarded as the strongest leveling spec for Hunter: a durable, hard-hitting pet absorbs damage that would otherwise land on the Hunter and clears trash faster than Marksmanship's single-target focus or Survival's melee weaving. The beta caps at level 30, so this is carried over from 1.12 leveling knowledge rather than tested in Forever.
