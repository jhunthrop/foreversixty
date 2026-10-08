---
title: Marksmanship Hunter in Forever
classSlug: hunter
spec: marksmanship
role: dps
build: 'FS1:1.60.1.70009:hunter:dwarf:5320000501/005155000150305/5:'
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

Marksmanship is a ranged DPS spec built around Aimed Shot: a slow, hard-hitting cast reinforced by Intellect-scaled Attack Power and bonus ranged crit damage, aiming for the highest single-shot damage of the three Hunter trees rather than pet or melee contribution. In a raid or group it plays a stationary, high-uptime ranged role. The beta caps at level 30, so anything here about level 60 play, raid tuning, or how Marksmanship compares to Beast Mastery or Survival at endgame is a projection from the demo trees and 1.12 knowledge, not confirmed play.

## Talents and builds

Reading the Marksmanship tree in priority order for a single-target Aimed Shot build: **Careful Aim** (adds Attack Power equal to a percentage of Intellect) and **Mortal Shots** (bonus critical strike damage on ranged abilities) come first, since both are flat multipliers on every shot that follows. **Trueshot Aura** is next: in Forever it grants ranged attack power only to the party, a narrower version than 1.12's melee-and-ranged buff per the demo notes. **Barrage** adds a damage bonus specifically to Multi-Shot, Aimed Shot, and Volley — all abilities this spec's rotation already leans on. **Efficiency** helps sustain the mana cost of a shot-heavy rotation. **Improved Stings** (1/3) adds 6% to Serpent Sting, which this rotation keeps up throughout.

On top of that spread, this build reaches further into Beast Mastery and down to the bottom of its own tree for three more measured damage gains. **Unleashed Fury** (5/5) adds 15% more pet and hawk damage and **Ferocity** (1/5) adds 2% pet and hawk critical strike chance. **Lethal Attacks** (5/5) adds a flat 5% critical strike chance with all attacks, and **Ranged Weapon Specialization** (5/5) adds a flat 5% more damage with ranged weapons — both are the kind of unconditional multiplier that's hard to pass up once the points are available.

The last round of the talent search moved three points: two from **Endurance Training** (pet health and armor, now 3/5) and the one in **Lone Wolf** (20% more damage with no active pet) went to the fifth point of Unleashed Fury, the first of Ferocity and the first of Improved Stings. Endurance Training and Lone Wolf both measured no damage in the sim, because a stationary single-target parse shows neither pet durability nor a pet-less bonus this build uses; they are real effects, just not ones the parse can show value from. The build keeps Improved Aspect of the Monkey, Hawk Eye, Improved Concussive Shot and Deflection at zero, as before. This build measures about +1.5% over the previous spend in our level-60 search run. Open the planner at [/planner?class=hunter](/planner?class=hunter) to build this out.

**Raid build.** The raid-ready search found no build that beats this one beyond error while keeping what the rotation casts, the group buffs the raid already counts on and the threat and survival talents, so the raid build is the leveling build and one tree serves both.

## Rotation and priority

This site's own simulator plays Marksmanship as the most direct read of the shared Hunter shot rotation: heavy Aimed Shot weaving, with Serpent Sting kept up throughout for the extra sustained damage. Before the pull, Aspect of the Hawk goes up and the opening Aimed Shot begins casting so it lands close to the start of the fight. In combat, Rapid Fire is timed to land together with an incoming auto shot right as Aimed Shot is about to come off cooldown, so both benefit from the same window. Aimed Shot is then recast on cooldown whenever the cast won't clip the next auto shot, Serpent Sting is refreshed whenever it's close to falling off and neither Rapid Fire nor Aimed Shot is close to ready, Multi-Shot fills the remaining gaps between Aimed Shot casts, and Arcane Shot is the last filler whenever Aimed Shot itself is more than 1.5 seconds from ready and mana is above a 25% reserve — Multi-Shot's own cooldown comes up often enough that gating Arcane Shot on it too would zero out Arcane Shot's casts entirely, so the priority no longer waits on it. Arcane Shot is instant, so it never clips an auto shot, but it's the least mana-efficient shot the spec has, hence the reserve floor. There's no melee or pet-management component to this priority list — the spec stays at range for the whole fight.

## Stat priority

The table above is this band's own simulation at level 60, re-run by the nightly pipeline every time the build or its gear data changes — these numbers are never hand-entered. A stat shown as "not separable from zero" is not a verdict that it is worthless; the sim's measured error on it is too wide, at this band's sample size, to tell its true value apart from zero. These numbers come from this site's own level-60 simulator, not from beta play, which only reaches level 30.

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
