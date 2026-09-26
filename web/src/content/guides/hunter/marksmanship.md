---
title: Marksmanship Hunter in Forever
classSlug: hunter
spec: marksmanship
role: dps
build: 'FS1:1.60.1.69893:hunter:dwarf:5522/35305500115003/51:'
recommendedRaces: [dwarf, troll]
statPriority: [Attack power, 'Ranged attack power', Agility, Critical strike, Hit, Melee haste]
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

Marksmanship is a ranged DPS spec built around Aimed Shot: a slow, hard-hitting cast reinforced by Intellect-scaled Attack Power and bonus ranged crit damage, aiming for the highest single-shot damage of the three Hunter trees rather than pet or melee contribution. In a raid or group it plays a stationary, high-uptime ranged role. The beta caps at level 30, so anything here about level 60 play, raid tuning, or how Marksmanship compares to Beast Mastery or Survival at endgame is a projection from the demo trees and 1.12 knowledge, not confirmed play.

## Talents and builds

Reading the Marksmanship tree in priority order for a single-target Aimed Shot build: **Careful Aim** (adds Attack Power equal to a percentage of Intellect) and **Mortal Shots** (bonus critical strike damage on ranged abilities) come first, since both are flat multipliers on every shot that follows. **Trueshot Aura** is next: in Forever it grants ranged attack power only to the party, a narrower version than 1.12's melee-and-ranged buff per the demo notes. **Barrage** adds a damage bonus specifically to Multi-Shot, Aimed Shot, and Volley — all abilities this spec's rotation already leans on. **Efficiency** helps sustain the mana cost of a shot-heavy rotation. **Lone Wolf**, new for Forever, is notable but situational: it grants 20% more damage with no active pet, making a pet-less Marksmanship build viable, at the cost of any pet utility or damage. Point allocation is heavily weighted into Marksmanship to reach Trueshot Aura and Barrage — roughly 30 or more points — with the remainder split between a supporting pet talent or two in Beast Mastery and utility points in Survival. Open the planner at [/planner?class=hunter](/planner?class=hunter) to build this out.

## Rotation and priority

This site's own simulator plays Marksmanship as the most direct read of the shared Hunter shot rotation: heavy Aimed Shot weaving, with Serpent Sting kept up throughout for the extra sustained damage. Before the pull, Aspect of the Hawk goes up and the opening Aimed Shot begins casting so it lands close to the start of the fight. In combat, Rapid Fire is timed to land together with an incoming auto shot right as Aimed Shot is about to come off cooldown, so both benefit from the same window. Aimed Shot is then recast on cooldown whenever the cast won't clip the next auto shot, Serpent Sting is refreshed whenever it's close to falling off and neither Rapid Fire nor Aimed Shot is close to ready, and Multi-Shot fills the remaining gaps between Aimed Shot casts. There's no melee or pet-management component to this priority list — the spec stays at range for the whole fight.

## Stat priority

In simulator-derived priority order: **attack power** (the core scalar behind auto shots and every shot in this rotation), **ranged attack power** specifically (the ranged-only component that stacks on top of general attack power), **agility** (adds both attack power and ranged crit indirectly), **crit** (extra shot damage, and, through Mortal Shots, extra crit damage on top of that), **hit** (misses waste an Aimed Shot cast, which is costly in a spec built around that one slow ability), and **melee haste** last, since this spec's damage comes almost entirely from ranged shots rather than melee swings.

## Gear

Look for attack power and ranged attack power first, then agility, then crit and hit, in the order above. Specific pre-raid or raid-tier item picks can't be named with confidence yet: the beta caps at level 30, and this build's raid loot tables are themselves incomplete — most items the community's own sourcing expects are missing from the current client, and nothing raids in-game until the first tier opens on 9 December 2026. This section fills in once raid loot is itemized and the spec is played past 30.

## Enchants and consumables

Target agility and attack power on weapon and glove enchants. **Enchant Gloves - Agility** and **Enchant Weapon - Agility** both exist in this build's enchant data, and a high-tier ranged weapon scope such as **Sniper Scope** is a verified item in this build — a natural fit given how much of this spec's output runs through one ranged weapon. For consumables, **Elixir of Greater Agility** and **Elixir of the Mongoose** are both verified options in this build's consumable list. Beyond naming those, specific best-in-slot consumable stacking isn't something this site can confirm yet at level 30.

## Races

Marksmanship can be played by every race Hunter allows: Human (new in Forever), Orc, Dwarf, Night Elf, Tauren, and Troll, plus the Skyborne on either faction. For Alliance, **Dwarf** is the strongest pairing: the reworked Dwarf racial, Big Game Hunter, adds bonus damage against Beasts specifically, which applies to Aimed Shot the same as any other physical attack. For Horde, **Troll** pairs well and pays off a little differently for this spec than for the other two: Berserking's 10% casting speed applies directly to Aimed Shot's own cast time, shortening the one ability Marksmanship revolves around, on top of the 10% attack speed and Beast Slaying's 5% damage against Beasts. Hunter is open to the Skyborne on either faction with no class restriction; their racials read as utility (a downward glide, a short full-resource restore, haste, bonus Elemental damage) rather than damage-focused, so Skyborne is a valid pick but not a clear upgrade over Dwarf or Troll for this spec.

## Professions

Community convention for Hunter favors **Skinning** for the cheap, plentiful leather this class already collects while killing beasts to level, often paired with **Leatherworking** for self-crafted agility gear, or **Engineering** for ranged-weapon scopes and trap-adjacent gadgets. This is general Classic-era community practice, not something confirmed for Forever specifically.

## Leveling

Beast Mastery is generally regarded as the stronger leveling spec for Hunter overall; Marksmanship's single-target focus doesn't clear trash as quickly without a pet built up the same way. The beta caps at level 30, so this is carried over from 1.12 leveling knowledge rather than tested in Forever.
