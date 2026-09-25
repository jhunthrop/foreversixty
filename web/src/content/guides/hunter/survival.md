---
title: Survival Hunter in Forever
classSlug: hunter
spec: survival
role: dps
description: 'Survival Hunter overview, talent priority, rotation, stat weights, and race picks for Forever, with beta-versus-projection called out.'
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
  - label: 'Warcraft Tavern, Survival Hunter rotation guide'
    url: https://www.warcrafttavern.com/wow-classic/guides/pve-survival-hunter-rotations-cooldowns/
    kind: community
---

## Overview

Survival blends the Hunter's usual ranged shots with melee weaving, using this tree's own talents to make standing in melee range worthwhile instead of a compromise. In a raid or group it plays a hybrid role: ranged damage backed by melee attacks and the class's trap kit, rather than committing fully to one range band the way Beast Mastery or Marksmanship do. The beta caps at level 30, so anything here about level 60 play, raid tuning, or how Survival compares to the other two specs at endgame is a projection from the demo trees and 1.12 knowledge, not confirmed play.

## Talents and builds

Reading the Survival tree in priority order: **Improved Wing Clip** and **Entrapment** raise the reliability of the Hunter's core control tools — Wing Clip's chance to immobilize and a trap's guaranteed short root — which matter more for a spec that plans to stay close to its target. **Predator's Edge** is a direct melee damage multiplier (bonus melee crit damage and off-hand damage), and **Counterattack** grants a free, unavoidable strike after a parry, rewarding the spec's melee stance. **Resourcefulness** trims what traps and melee abilities cost in mana and gives some critical hits a partial refund of casting regeneration — useful here since Survival draws mana from both its ranged shots and its melee attacks, unlike the other two specs. The tree's capstone, **Lacerating Strikes**, adds a bleed to Mongoose Bite, but Mongoose Bite doesn't appear in this site's current Survival rotation (see below), so that capstone's value can't be verified against a working priority list right now — flagged here as unconfirmed rather than asserted as a settled pick. Point allocation splits more evenly than the other two Hunter trees: a heavy Survival investment for the melee and control talents above, with supporting points spent in Marksmanship for Aimed Shot's own multipliers. Open the planner at [/planner?class=hunter](/planner?class=hunter) to build this out.

## Rotation and priority

This site's own simulator plays Survival on the same Aimed Shot and Multi-Shot base as the other two specs, but weaves in melee Raptor Strike, which this tree's talents make worth using. Aimed Shot is prioritized first whenever its cast won't clip the next auto shot; Raptor Strike then fills as the melee attack, gated so it never drops the Hunter below 20% mana, since Survival is spending mana on both its ranged and melee kit. Multi-Shot fills the ranged gaps between Aimed Shot casts, and Serpent Sting is refreshed last, only once it's close to falling off, kept deliberately behind both shots rather than ahead of them. One limitation worth flagging honestly: this site's simulated Hunters currently hold a fixed ranged distance from the target for the whole fight, which keeps the ranged shots working but leaves no room in the simulation to step into melee range on a stationary single-target fight — so the Raptor Strike line, while the correct priority order for how Survival is meant to play, isn't actually exercised in this site's own current sim numbers.

## Stat priority

In simulator-derived priority order: **attack power** (the core scalar behind auto shots, ranged abilities, and Raptor Strike alike), **ranged attack power** specifically (the ranged-only component that stacks on top of general attack power), **agility** (adds both attack power and crit indirectly, on both the ranged and melee sides), **crit** (extra damage on shots and, through Predator's Edge, extra melee crit damage too), **hit** (misses cost shot uptime and, for this spec specifically, melee uptime as well), and **melee haste** last, though it matters somewhat more here than for the other two specs given how much of this rotation is melee-based.

## Gear

Look for attack power and ranged attack power first, then agility, then crit and hit, in the order above. Specific pre-raid or raid-tier item picks can't be named with confidence yet: the beta caps at level 30, and this build's raid loot tables are themselves incomplete — most items the community's own sourcing expects are missing from the current client, and nothing raids in-game until the first tier opens on 9 December 2026. This section fills in once raid loot is itemized and the spec is played past 30.

## Enchants and consumables

Target agility and attack power on weapon and glove enchants — the weapon enchant matters for both the ranged weapon and any melee weapon this spec carries. **Enchant Gloves - Agility** and **Enchant Weapon - Agility** both exist in this build's enchant data, and a ranged weapon scope such as **Deadly Scope** or **Sniper Scope** is a verified item in this build. For consumables, **Elixir of Greater Agility** and **Elixir of the Mongoose** are both verified options in this build's consumable list. Beyond naming those, specific best-in-slot consumable stacking isn't something this site can confirm yet at level 30.

## Races

Survival can be played by every race Hunter allows: Human (new in Forever), Orc, Dwarf, Night Elf, Tauren, and Troll, plus the Skyborne on either faction. For Alliance, **Dwarf** is the strongest pairing: the reworked Dwarf racial, Big Game Hunter, adds bonus damage against Beasts specifically, which applies to Raptor Strike's melee hits against beast targets exactly as it applies to a ranged shot. For Horde, **Troll** pairs well, and pays off on both halves of this spec's kit: Berserking's 10% attack speed shortens the melee swing timer that drives Raptor Strike uptime as well as the ranged auto-shot loop, and Beast Slaying's 5% damage against Beasts stacks on top of it. Hunter is open to the Skyborne on either faction with no class restriction; their racials read as utility (a downward glide, a short full-resource restore, haste, bonus Elemental damage) rather than damage-focused, so Skyborne is a valid pick but not a clear upgrade over Dwarf or Troll for this spec.

## Professions

Community convention for Hunter favors **Skinning** for the cheap, plentiful leather this class already collects while killing beasts to level, often paired with **Leatherworking** for self-crafted agility gear, or **Engineering** for ranged-weapon scopes and trap-adjacent gadgets. This is general Classic-era community practice, not something confirmed for Forever specifically.

## Leveling

Beast Mastery is generally regarded as the strongest leveling spec for Hunter; Survival's melee-weaving payoff depends on a deeper Survival investment than a leveling character can usually afford early on, so it comes online later than Beast Mastery does. The beta caps at level 30, so this is carried over from 1.12 leveling knowledge rather than tested in Forever.
