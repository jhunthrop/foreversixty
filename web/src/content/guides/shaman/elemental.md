---
title: Elemental Shaman in Forever
classSlug: shaman
spec: elemental
role: dps
build: 'FS1:1.60.1.70009:shaman:dwarf:553231130010305/01/553302:'
recommendedRaces: [dwarf, orc]
statPriority: [Intellect, Spell power, Hit, Nature power, Critical strike, Spell haste, Spell penetration]
description: 'Elemental Shaman overview, talent priority, rotation, stat weights, and race picks for Forever, with beta-versus-projection called out.'
updated: 2026-09-24
confidence: inferred
sources:
  - label: 'Blizzard, Deep Dive panel recap'
    url: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
    kind: blizzard
  - label: 'Talents Forever (demo transcription)'
    url: https://talentsforever.com/data.json
    kind: community
  - label: 'This site, stat weights'
    url: https://foreversixty.gg/sim/weights
    kind: site
  - label: 'This site, character planner'
    url: https://foreversixty.gg/planner
    kind: site
---

## Overview

Elemental is Shaman's ranged spellcasting DPS spec, built on Lightning Bolt and Chain Lightning backed by Flame Shock and totem support, with Lava Burst adding a new high-damage nuke that hits harder when Flame Shock is already on the target. Forever's baseline cast-time cut to Lightning Bolt and Chain Lightning, before any talent points are spent, is a meaningful floor-raise for the spec's damage even before talents are considered. The beta caps at level 30, so how Elemental's damage compares to other casters at level 60, or performs in a raid, is a projection from the demo talent trees and 1.12 knowledge, not confirmed play.

## Talents and builds

In rough priority order: **Concussion** raises the damage of Lightning Bolt, Chain Lightning, and Earth Shock, a broad multiplier across the entire single-target and AoE kit. **Call of Flame**, reworked in Forever, now boosts Lava Burst and Fire Nova alongside Flame Shock within a single talent instead of touching only the Fire Totems it used to buff, effectively unifying the tree's fire-damage kit under one pick. **Elemental Fury** raises critical strike damage on the totems and elements the spec relies on, moved earlier in the tree with a lower first-rank value than 1.12's single large-rank version. **Lightning Overload** gives Lightning Bolt and Chain Lightning a chance to cast a second, weaker bolt for free, a straight damage-per-cast increase. **Elemental Alacrity** cuts the cast time of Lightning Bolt, Chain Lightning, and Lava Burst further on top of the baseline speed-up Forever already gave those spells — maxed, it brings deep Elemental's effective cast speed roughly back to where 1.12 Elemental topped out, since the starting point is already faster. **Lava Burst**, the tree's new capstone nuke, hits harder against a target already carrying Flame Shock.

Point allocation runs deep into Elemental to reach the bottom of the tree (32 points: Convection, Concussion, Call of Flame, Elemental Fury and Lightning Overload as above, plus Elemental Focus and Call of Thunder), with 1 point in Enhancement for Thundering Strikes and the remaining 18 in Restoration for Improved Healing Wave and Totemic Focus (both maxed), Mindfulness, Natural Grace, and Improved Reincarnation. This replaces the earlier 31/0/20 build, and the talent search on the current engine put it about 10% ahead, beyond the run's error. Two things moved. Elemental Focus is now taken: it was the single largest per-point gain the search found. Convection takes its fifth rank and Thundering Strikes its first. The points came from Tidal Focus, which adds no damage in the engine, and from the Lava Burst talent itself, a fire-and-forget capstone this loop never casts; the loop never casts Lava Burst and the search measured the point as a small net loss, so it went to Elemental Focus. Every talent the engine does not yet model (Elemental Warding, Mindfulness, Natural Grace, Improved Healing Wave, Improved Reincarnation) stays as before, since the search cannot judge it. Open the planner at [/planner?class=shaman](/planner?class=shaman) to build this out.

## Rotation and priority

The loop this site's simulator plays: against a single target, drop Searing Totem for the higher fire damage; against two or more targets, drop Magma Totem instead so everyone standing in it takes damage. Against two or more targets, Chain Lightning is used since it hits every target for close to single-target Lightning Bolt damage. Lightning Bolt fills the rest of the single-target rotation, on every free global. Flame Shock, Lava Burst and Earth Shock are not part of the loop. Flame Shock costs more mana than the damage it adds, and Lava Burst is only worth its extra damage on a Flame-Shocked target, so the pair together do not beat plain Lightning Bolt casts. Lava Burst on its own comes out level with Lightning Bolt within our measurement error, so it is not a reason to rearrange the loop. Earth Shock as a mana dump sits behind Lightning Bolt and rarely gets a global. Mana Spring Totem goes down first: it lasts five minutes, returns 10 mana every 2 seconds at its top rank, and the rest of the loop is limited by mana long before a three-minute fight ends, so the 100 mana it costs comes back many times over; adding it was worth roughly 4% in our level-60 run, and starting it one global before the cooldowns was worth about 1.5% more. Fire Nova is a real spell in Forever (learned from level 12, instant, on a 10-second cooldown, and it needs a Fire totem standing), but it is not in the loop: it deals about a third of Lightning Bolt's damage per point of mana, and at no mana threshold we tried did adding it beat plain Lightning Bolt casts. We re-checked these numbers against the client's own spell data, and every one of them got smaller; the order of the loop did not change.

## Stat priority

The table above is this band's own simulation at level 60, re-run by the nightly pipeline every time the build or its gear data changes — these numbers are never hand-entered. A stat shown as "not separable from zero" is not a verdict that it is worthless; the sim's measured error on it is too wide, at this band's sample size, to tell its true value apart from zero. These numbers come from this site's own level-60 simulator, not from beta play, which only reaches level 30.

## Gear

Prioritize spell power and intellect first, then crit and hit, then spell haste, spell penetration, and Nature-specific power. Specific pre-raid or raid-tier item picks can't be named with confidence yet: the beta caps at level 30 and this build's raid loot tables are still missing most items the community's sourcing expects, with nothing raiding in-game until the first tier opens on 9 December 2026. This section fills in once raid loot is itemized and the class is played past 30.

## Enchants and consumables

Weapon enchants should target spell power; **Enchant Weapon - Spell Power** is a verified entry in this build's enchant data. For consumables, **Greater Arcane Elixir** (spell power and crit) is a verified option. A fuller consumable stack isn't something this site is naming until the class is played past level 30.

## Races

Elemental can be played by Orc and Troll on the Horde, Tauren on the Horde, Dwarf — new in Forever — on the Alliance, and the Windshaper Skyborne on the Horde only. Dwarf is, for now, the only Alliance option for Shaman at all. For Alliance, **Dwarf** is the only pick, but it's a reasonable one for Elemental: Big Game Hunter adds damage against Beasts, useful for leveling and questing damage even if it doesn't apply in most raid encounters. For Horde, **Orc** is the strongest pick: Blood Fury grants 10% attack power and spell power for 15 seconds on a 2-minute cooldown, and the spell power half applies directly to Elemental's entire damage kit, making it a straightforward burst cooldown with no downside for a caster build.

## Professions

Community convention for caster Shaman favors **Enchanting** alongside **Tailoring** or a gathering profession for spell-power gear and enchant materials, though Shaman's totem and weapon-imbue kit doesn't lean as hard on any one profession as a pure caster class does. This is general Classic-era community practice, not confirmed for Forever.

## Leveling

Enhancement is generally the stronger leveling spec for Shaman, since its melee durability and weapon-imbue burst clear trash faster than Elemental's cast-time-dependent damage. The beta caps at level 30, so this is carried over from 1.12 leveling knowledge rather than tested in Forever.
