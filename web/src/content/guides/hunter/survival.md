---
title: Survival Hunter in Forever
classSlug: hunter
spec: survival
role: dps
build: 'FS1:1.60.1.69893:hunter:dwarf:0/32005500005/500230131051120151:'
recommendedRaces: [dwarf, troll]
statPriority: [Attack power, Agility, Strength, Critical strike, Hit, Melee haste]
description: 'Survival Hunter overview, talent priority, rotation, stat weights, and race picks for Forever, with beta-versus-projection called out.'
updated: 2026-09-28
confidence: inferred
sources:
  - label: 'Blizzard, Deep Dive panel recap'
    url: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
    kind: blizzard
  - label: 'Talents Forever (demo transcription)'
    url: https://talentsforever.com/data.json
    kind: community
  - label: 'Forever Sixty build data, data/builds/1.60.1.70009/talents/hunter.json'
    url: https://foreversixty.gg/data/1.60.1.70009/talents/hunter.json
    kind: datamined
  - label: 'Forever Sixty build data, data/builds/1.60.1.70009/spellconst/hunter.json'
    url: https://foreversixty.gg/data/1.60.1.70009/spellconst/hunter.json
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

Survival is a melee spec in Forever (owner direction, 2026-09-28): it commits fully to standing in weapon range instead of splitting time between shots and Raptor Strike the way earlier builds of this guide described. The old "ranged base with a melee weave" reading came from the shared Hunter spell package every spec starts from; Survival's own tree — Counterattack, Strider Kick, Predator's Edge, Lacerating Strikes — only pays off in melee, and this site's simulator now plays it that way: in melee range for the whole fight, no shots. The beta caps at level 30, so anything here about level 60 play, raid tuning, or how Survival compares to the other two specs at endgame is a projection from the demo trees and 1.12 knowledge, not confirmed play.

## Talents and builds

Reading the Survival tree in priority order for a melee build: **Savage Strikes** and **Predator's Edge** are direct melee damage — crit chance and crit damage on every melee special, plus a bonus to off-hand damage if this spec carries a second weapon. **Deterrence** unlocks **Counterattack**, a free, unavoidable strike after a parry (cannot be blocked, dodged, or parried) — a genuine "free damage" proc this build takes specifically to use. **Expose Prey** unlocks the capstone, **Lacerating Strikes**, which adds a bleed to Mongoose Bite worth 40% of that hit's damage over 21 seconds — Mongoose Bite itself is not talent-gated (it's a class ability, like Raptor Strike), so this capstone is a straightforward damage-per-point pick once the tree is deep enough to reach it. **Strider Kick** is a third melee special on its own short cooldown. **Improved Wing Clip** and **Entrapment** are left at 0 here: they raise the reliability of control tools this rotation doesn't use against a stationary target, and the points are worth more spent reaching Counterattack and Lacerating Strikes sooner. The remaining points go into Marksmanship, spent reaching **Mortal Shots** (its own prerequisite, **Careful Aim**, needs all 5 points) — in this engine, Mortal Shots adds its crit-damage bonus to every Hunter special, melee included, rather than into Beast Mastery or further Survival utility. Open the planner at [/planner?class=hunter](/planner?class=hunter) to build this out.

## Rotation and priority

This site's simulator starts Survival in melee range and keeps it there for the whole fight: Raptor Strike is queued on the swing timer (it replaces the next melee auto-attack rather than costing its own swing), gated at 20% mana so it never starves Mongoose Bite. Counterattack is a bare cast that only lands while its parry-triggered window is open — the same "free line, falls through otherwise" shape this site's own Warrior Overpower line uses — so it is checked early, right after Raptor Strike's queue. Mongoose Bite is cast on cooldown: Forever's own tooltip for it lists no dodge requirement (vanilla Classic's version needed one), so its 5-second cooldown and mana cost are the only real gates. Strider Kick fills as the lowest-priority special on its own 8-second cooldown. There's no instant, spammable melee filler beyond that in Classic-style Hunter kits, so a global with nothing else ready is spent on nothing beyond the swing timer itself — the same shape live Classic Survival play has always had. Ranged shots (Aimed Shot, Multi-Shot, Arcane Shot) and Serpent Sting are not part of this rotation at all: they require standing outside melee range, which this build never does.

## Stat priority

In simulator-derived priority order: **attack power** (the core scalar behind every melee special and the auto-attack), **agility** (adds attack power and crit both), **strength** (a smaller but real attack-power contributor a melee Hunter didn't need to weigh before), **crit** (more crits, and bigger ones through Predator's Edge and Mortal Shots), **hit** (a miss costs melee uptime, and this spec has no ranged fallback to cover for it), and **melee haste** last — it speeds up the swing timer Raptor Strike rides, but everything ahead of it matters more per point.

## Gear

Look for attack power first, then agility and strength, then crit and hit, in the order above. Prioritize weapon slots: a main-hand (and off-hand, if this build carries one — Predator's Edge's off-hand damage bonus only matters with one equipped) over a ranged weapon, which this rotation no longer uses for anything but a stat stick. Specific pre-raid or raid-tier item picks can't be named with confidence yet: the beta caps at level 30, and this build's raid loot tables are themselves incomplete — most items the community's own sourcing expects are missing from the current client, and nothing raids in-game until the first tier opens on 9 December 2026. This section fills in once raid loot is itemized and the spec is played past 30.

## Enchants and consumables

Target agility and attack power on weapon and glove enchants — the weapon enchant matters most here, since this spec's melee weapon (or weapons, if dual-wielding) carries the whole rotation now, not a ranged one. **Enchant Gloves - Agility** and **Enchant Weapon - Agility** both exist in this build's enchant data. For consumables, **Elixir of Greater Agility** and **Elixir of the Mongoose** are both verified options in this build's consumable list. Beyond naming those, specific best-in-slot consumable stacking isn't something this site can confirm yet at level 30.

## Races

Survival can be played by every race Hunter allows: Human (new in Forever), Orc, Dwarf, Night Elf, Tauren, and Troll, plus the Skyborne on either faction. For Alliance, **Dwarf** is the strongest pairing: the reworked Dwarf racial, Big Game Hunter, adds bonus damage against Beasts specifically, which applies to every melee special here against a beast target exactly as it would to a shot. For Horde, **Troll** pairs well: Berserking's 10% attack speed shortens the melee swing timer that now drives this entire rotation, not just the Raptor Strike weave, and Beast Slaying's 5% damage against Beasts stacks on top of it. Hunter is open to the Skyborne on either faction with no class restriction; their racials read as utility (a downward glide, a short full-resource restore, haste, bonus Elemental damage) rather than damage-focused, so Skyborne is a valid pick but not a clear upgrade over Dwarf or Troll for this spec.

## Professions

Community convention for Hunter favors **Skinning** for the cheap, plentiful leather this class already collects while killing beasts to level, often paired with **Leatherworking** for self-crafted agility gear, or **Blacksmithing** for a melee-focused Hunter's own weapon and armor work. This is general Classic-era community practice, not something confirmed for Forever specifically.

## Leveling

Beast Mastery is generally regarded as the strongest leveling spec for Hunter; a melee-committed Survival needs deeper Survival investment (Deterrence for Counterattack, Expose Prey for Lacerating Strikes) than a leveling character can usually afford early, so it comes online later than Beast Mastery does, and later than it did as a part-ranged spec. The beta caps at level 30, so this is carried over from 1.12 leveling knowledge rather than tested in Forever.
