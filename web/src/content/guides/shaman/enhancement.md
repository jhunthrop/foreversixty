---
title: Enhancement Shaman in Forever
classSlug: shaman
spec: enhancement
role: dps
build: 'FS1:1.60.1.70009:shaman:dwarf:550333/254130031005002051/0:'
recommendedRaces: [dwarf, orc]
statPriority: [Strength, Attack power, Hit, Critical strike, Agility, Melee haste]
description: 'Enhancement Shaman overview, talent priority, rotation, stat weights, and race picks for Forever, with beta-versus-projection called out.'
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

Enhancement is Shaman's melee DPS spec, combining weapon imbues like Windfury Weapon with Stormstrike and totem support to fight in melee range rather than at range. Forever's unification of melee and spell hit into a single stat is a particularly large indirect buff to this spec, since 1.12 Enhancement had to gear separately for melee hit against its physical attacks and spell hit against its shock spells. The beta caps at level 30, so how Enhancement's damage compares to other melee specs at level 60, or performs in a raid, is a projection from the demo talent trees and 1.12 knowledge, not confirmed play.

## Talents and builds

In rough priority order: **Thundering Strikes** raises critical strike chance with all attacks and spells, a broad early multiplier. **Elemental Weapons** increases the attack power bonus from Rockbiter Weapon and the proc value of Windfury Weapon and Flametongue Weapon, directly scaling the spec's weapon-imbue kit. **Flurry** grants an attack speed bonus for several swings after a melee crit, compounding with the crit chance Thundering Strikes already added. **Maelstrom Weapon**, new to the tree, can proc off a melee swing and stacks up a discount that shortens Lightning Bolt's cast and trims its mana cost, letting Enhancement weave in a ranged nuke without losing much melee uptime. **Rage of the Farseer**, the tree's new capstone, increases both melee attack speed and spell casting speed for 25 seconds as a burst cooldown.

This re-spend adds **Call of Flame** at 3/3, pushing damage from Fire Totems and from Flame Shock and Fire Nova up 15%. **Elemental Devastation** goes to 3/3; its tooltip turns offensive spell crits into a 9% melee crit chance for 10 seconds, but the engine has no code path yet for that interaction, so it's an honest unknown rather than a confirmed zero — it costs nothing extra beyond the other adds, so it rides along. **Ancestral Knowledge** goes to 4/5 for an 8% Intellect boost at that rank, feeding more spell power into the shock-and-totem side of the kit. **Shamanistic Focus** — a single point — is the single largest contributor in this re-spend, cutting Shock and Lightning Shield mana cost by 45%.

To pay for them, this build drops **Elemental Warding** entirely (down from 3/3), a real 10%-at-max reduction to Fire/Frost/Nature damage taken that the raid-DPS sim simply can't credit as damage. It also drops **Stormstrike** to 0/1. This is the harder call: Stormstrike is Enhancement's signature attack, and this guide's own Rotation section still calls it "the highest damage-per-global ability in the kit," used on cooldown — Stormstrike is a genuine damage ability, and the engine's own probe confirms removing it in isolation would cost DPS, but in this build's overall re-spend, moving that single point into Call of Flame, Elemental Devastation, and Shamanistic Focus together nets more total damage under the current rotation model. That is a real tension between this talent search's result and the spec's usual signature-ability framing, worth flagging honestly rather than papering over: this page's Rotation section has not yet been revisited for a build that drops Stormstrike, and treating this spend as final would mean losing Enhancement's core melee hit entirely. This build measures about +15.4% over the previous spend in our level-60 search run.

Point allocation still runs deep into Enhancement to reach Maelstrom Weapon and Rage of the Farseer near the bottom of the tree, picks up slightly more in Elemental for Concussion, Call of Flame, and Elemental Devastation, and spends nothing in Restoration. The beta caps at level 30, so none of this spend has actually been played — it's a projection from the demo talent trees, this site's own simulator, and 1.12 knowledge, not confirmed play. Open the planner at [/planner?class=shaman](/planner?class=shaman) to build this out.

## Rotation and priority

Weapon imbues are class kit, not a rotation line — applied before the pull and reapplied whenever it falls off, the same way a Rogue keeps poisons up. Shamans don't dual-wield in Forever, so there's only ever one weapon to imbue: Rockbiter Weapon goes on it from level 1 (it's the only imbue learned that early), and Windfury Weapon replaces Rockbiter there once learned at level 30, staying on through 60. The off hand carries a shield or another held item rather than a second weapon.

The loop this site's simulator plays: keep Strength of Earth Totem down throughout the fight — the engine tracks totem uptime directly rather than needing an aura check, so it's treated as a simple refresh condition. Keep Windfury Totem down as well, since it's the melee group's largest damage totem and occupies the same Air-totem slot Grace of Air would otherwise use. Stormstrike is used on cooldown as the highest damage-per-global ability in the kit. Searing Totem is kept down for extra fire damage, but not refreshed with less than 20 seconds left on the fight, since there's no point paying totem mana for a totem that will barely tick before the encounter ends. Earth Shock is used as a mana dump once mana is comfortably above half.

## Stat priority

The table above is this band's own simulation at level 60, re-run by the nightly pipeline every time the build or its gear data changes — these numbers are never hand-entered. A stat shown as "not separable from zero" is not a verdict that it is worthless; the sim's measured error on it is too wide, at this band's sample size, to tell its true value apart from zero. These numbers come from this site's own level-60 simulator, not from beta play, which only reaches level 30.

## Gear

Prioritize attack power, strength, and agility first, then crit and hit, then melee haste. Because hit is now a single stat covering both Enhancement's melee attacks and its shock spells, gear that used to be a tradeoff between the two kinds of hit rating is simply better across the board for this spec than it was in 1.12. Specific pre-raid or raid-tier item picks can't be named with confidence yet: the beta caps at level 30 and this build's raid loot tables are still missing most items the community's sourcing expects, with nothing raiding in-game until the first tier opens on 9 December 2026.

## Enchants and consumables

Weapon enchants should target strength or agility depending on the piece; **Enchant Weapon - Strength** and **Enchant Weapon - Agility** are both verified entries in this build's enchant data, alongside **Enchant Weapon - Crusader**, a proc-based option also present in the data. For consumables, **R.O.I.D.S.** (strength) is a verified option. A fuller consumable stack isn't something this site is naming until the class is played past level 30.

## Races

Enhancement can be played by Orc and Troll on the Horde, Tauren on the Horde, Dwarf — new in Forever — on the Alliance, and the Windshaper Skyborne on the Horde only. Dwarf is, for now, the only Alliance option for Shaman at all. For Alliance, **Dwarf** is the only pick, and it fits a melee spec reasonably well: Stoneform now reduces physical damage taken rather than only raising armor, a direct survivability gain for a spec standing in melee range. For Horde, **Orc** is the strongest pick: Blood Fury's 10% attack power for 15 seconds on a 2-minute cooldown lines up directly with Enhancement's melee damage kit, and Axe Specialization's crit bonus rewards an Orc Enhancement Shaman for choosing an axe as their main-hand weapon.

## Professions

Community convention for melee Shaman favors **Leatherworking** alongside **Skinning** for self-sourced agility and strength gear, matching general Classic-era melee-class practice. This is community convention, not confirmed for Forever.

## Leveling

Enhancement is generally the strongest leveling spec for Shaman: its melee durability, weapon imbues, and Stormstrike burst clear trash faster than Elemental's cast-time-dependent damage or Restoration's support-focused kit. The beta caps at level 30, so this is carried over from 1.12 leveling knowledge rather than tested in Forever.
