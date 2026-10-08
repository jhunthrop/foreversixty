---
title: Beast Mastery Hunter in Forever
classSlug: hunter
spec: beast-mastery
role: dps
build: 'FS1:1.60.1.70291:hunter:dwarf:5420001505001251/0053502001/4:'
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

Beast Mastery is a pet-focused ranged DPS spec: a large share of total damage in Forever comes from a strengthened pet rather than the Hunter's own shots, with Bestial Wrath as the spec's signature burst cooldown. In a raid or group it plays consistent single-target damage with a pet that can absorb hits and threat the Hunter would otherwise take directly. The beta caps at level 30, so anything here about level 60 play, raid tuning, or how Beast Mastery compares to the other two specs at endgame is a projection from the demo trees and 1.12 knowledge, not confirmed play.

## Talents and builds

Reading the Beast Mastery tree in priority order for a pet-damage build: **Unleashed Fury** (a flat damage increase for pets and hawks) and **Ferocity** (pet critical strike chance) come first, since both scale every hit the pet lands afterward. **Focused Fire** follows — its damage bonus applies to the Hunter as well as the pet, unlike the two talents above it. **Frenzy** gives the pet a chance at bonus attack speed off its own critical strikes, compounding with Ferocity. **Bestial Discipline** helps mana sustain by letting some regeneration continue through casting. The tree's capstone, **Bestial Wrath**, is the spec's defining cooldown: an 18-second window of 50% additional pet damage.

The remaining points move out of a thin Marksmanship/Survival utility spread and into talents the search confirmed actually move the parse. **Efficiency** (5/5) cuts 15% off the mana cost of shots, stings, and melee abilities, paying for a shot-heavy rotation. **Trueshot Aura** (1 point) adds 30 Ranged Attack Power to the whole party for 30 minutes — a raid-wide buff on top of its personal value. **Improved Stings** (3/3) adds 20% more Serpent Sting damage, a shorter Viper Sting cooldown, and a longer Scorpid Sting duration, and **Rapid Killing** (2/2) shortens Rapid Fire's cooldown and grants a temporary damage buff to the next shot on a kill — both of these measure close to zero in the sim today because the engine doesn't yet have code paths for Serpent Sting's damage scaling or Rapid Killing's proc, so treat them as an honest unknown rather than a confirmed gain, taken because nothing else was better with those points.

To pay for that, the build drops **Hawk Eye** (ranged weapon range) and **Improved Concussive Shot** (a stun chance on Concussive Shot) entirely, and **Deflection** (parry chance) entirely — all three measured no damage benefit at all in the sim, so the points are better spent elsewhere for a pure DPS build even though they're real defensive or utility effects in a fight where parrying or kiting matters. It also gives back the fifth point in **Improved Tracking** (the per-point damage bonus against your tracked creature type), keeping it at 4/5 instead of 5/5 — that fifth point is a real, measured damage pick, just a smaller one than what replaced it, so it's a deliberate trade rather than a wasted point. This build measures about +8.5% over the previous spend in our level-60 search run. Open the planner at [/planner?class=hunter](/planner?class=hunter) to build this out.

**Raid build.** The raid-ready search found no build that beats this one beyond error while keeping what the rotation casts, the group buffs the raid already counts on and the threat and survival talents, so the raid build is the leveling build and one tree serves both.

## Rotation and priority

This site's own simulator plays Beast Mastery on the same Aimed Shot and Multi-Shot shell the other two specs share, with Bestial Wrath layered in as this spec's own cooldown. Before the pull, Aspect of the Hawk goes up and the first Aimed Shot begins casting so it lands right as the fight opens. Once engaged, Bestial Wrath is used on cooldown first and timed to land together with Rapid Fire, since a fight of Classic length only gives room for one real burst window — stacking both cooldowns into it beats spreading them apart. Aimed Shot is then recast every time it comes off cooldown, as long as the cast won't clip the next auto shot. Multi-Shot comes next on every free global between Aimed Shot casts, then Serpent Sting is refreshed whenever it's close to falling off and neither Rapid Fire nor Aimed Shot is close to ready (putting Multi-Shot ahead of the sting was a small gain, about +0.4% in the raid-ready run, and level for the bare character), and Arcane Shot is the last filler whenever Aimed Shot itself is more than 1.5 seconds from ready and mana is above a 25% reserve — Multi-Shot's own cooldown comes up often enough that gating Arcane Shot on it too would zero out Arcane Shot's casts entirely, so the priority no longer waits on it. Arcane Shot is instant, so it never clips an auto shot, but it's the least mana-efficient shot the spec has, hence the reserve floor. The priority list doesn't call for active pet repositioning beyond keeping the pet alive and attacking through the Bestial Wrath window.

## Stat priority

The table above is this band's own simulation at level 60, re-run by the nightly pipeline every time the build or its gear data changes — these numbers are never hand-entered. A stat shown as "not separable from zero" is not a verdict that it is worthless; the sim's measured error on it is too wide, at this band's sample size, to tell its true value apart from zero. These numbers come from this site's own level-60 simulator, not from beta play, which only reaches level 30.

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
