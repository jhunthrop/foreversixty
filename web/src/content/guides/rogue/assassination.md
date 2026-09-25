---
title: Assassination Rogue in Forever
classSlug: rogue
spec: assassination
role: dps
description: 'Assassination Rogue overview, talent priority, rotation, stat weights, and race picks for Forever, with beta-versus-projection called out.'
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

Assassination is a poison-focused melee DPS spec that trades Combat's flat weapon damage for consistent poison uptime and a guaranteed-critical finisher window through Cold Blood. In a group it plays a single-target damage role rather than the cleave utility Combat or Subtlety can offer. The beta caps out at level 30, so nothing below about how this spec performs at level 60, in a full poison-uptime rotation, or against raid-tuned targets is confirmed play — it's read off the demo trees and 1.12 knowledge, not tested.

## Talents and builds

Reading the tree in priority order for a single-target build: **Malice** (flat crit chance with attacks and poisons) and **Lethality** (bonus critical strike damage on Sinister Strike, Gouge, Backstab, Mutilate, Ghostly Strike, and Hemorrhage) come first as damage multipliers that scale everything after them. **Seal Fate** follows, since it converts combo-point-generating crits into bonus combo points and rewards the crit investment already made. **Cold Blood** is the spec's signature burst cooldown, guaranteeing a critical strike on the next attack. **Vile Poisons** and **Improved Poisons** raise both poison damage and application chance, which matters more here than in the other two trees. The two talent-granted abilities new to this tree in Forever, **Mutilate** (a dual-wield combo builder that hits harder against a target already carrying your poisons) and **Venom** (a finisher that raises how hard and how often your poisons land), round out the kit but aren't verified against a working rotation yet — see Rotation below.

Point allocation is heavily weighted into Assassination: roughly 31 points are needed to reach Venom at the bottom of the tree, leaving the rest split as a handful of points in Combat for weapon-skill and survivability talents and a few in Subtlety for utility. Open the planner at [/planner?class=rogue](/planner?class=rogue) to build this out.

## Rotation and priority

This site's own simulator currently plays Assassination close to the 1.12 Sinister Strike into Eviscerate loop, because Mutilate and Venom — the two abilities the Forever Assassination tree grants — don't have a working implementation in the simulator yet and are skipped when the rotation reaches them. Until that lands, the practical priority is: keep Slice and Dice up so combo points aren't wasted on idle swing time, use Cold Blood immediately before a finisher for the guaranteed crit, and dump five combo points into Eviscerate. Sinister Strike fills the remaining energy as the combo-point builder. Once Mutilate and Venom are simulated, expect Mutilate to replace Sinister Strike as the builder and Venom to compete with Eviscerate as the finisher, per the tree's intent — but that's this site's own projection, not a confirmed rotation.

## Stat priority

In simulator-derived priority order: **attack power** (the primary damage driver for every physical attack in the rotation), **agility** (adds both attack power and crit indirectly), **crit** (feeds Seal Fate's bonus combo points and Cold Blood's guaranteed hit), **hit** (avoiding misses keeps combo-point generation consistent), and **melee haste** last (more swings and faster energy-limited combo generation).

## Gear

Look for agility and attack power first, then the merged hit and crit ratings the itemization change consolidated from their old melee/spell splits, then melee haste. A new stat that reduces the target's dodge and parry chance may show up on gear as itemization fills in, which would sit alongside hit rating in value for a dagger-and-poison spec that wants attacks to land. Specific pre-raid or raid-tier item picks can't be named with confidence yet: the beta caps at level 30, and this build's raid loot tables are themselves incomplete — most raid items the community's own sourcing expects are absent from the current client, and nothing raids in-game until the first tier opens on 9 December 2026. This section fills in once raid loot is itemized and the class is played past 30.

## Enchants and consumables

Target agility and attack power on weapon and glove enchants; **Enchant Gloves - Superior Agility** and **Enchant Weapon - Agility** both exist in this build's enchant data. For consumables, **Elixir of the Mongoose** (agility) is a verified option in this build's consumable list. Beyond naming those two, specific best-in-slot consumable stacking is not something this site can confirm yet at level 30.

## Races

Assassination can be played by every race the Rogue class allows: Human, Dwarf, Night Elf, and Gnome on the Alliance; Orc, Undead, and Troll on the Horde; and, new in Forever, the Skyborne on either faction, since Skyborne is a new race rather than one of Blizzard's six new class pairings. For Alliance, **Night Elf** is the strongest pairing for a stealth-based melee class: Shadowmeld stealths on demand, works in combat, and is usable on a 2-minute cooldown, giving Assassination a reliable way to vanish out of a bad pull or reset an opener that Human or Dwarf can't match. For Horde, **Troll** pairs well through Berserking, a 10% attack and casting speed cooldown on a 3-minute timer that stacks on top of Cold Blood for a bigger burst window. Rogue is open to the Skyborne on both factions with no class restriction; their racials read as more utility- and survival-focused (a downward glide, a short full-resource restore) than damage-oriented, so they're a valid but not obviously stronger pick than Night Elf or Troll for this spec.

## Professions

Community convention for Rogue favors **Engineering** for its combat utility items (bombs, gadgets) alongside a gathering profession like Herbalism or Mining for self-funded leveling, since Rogue has no crafting profession that directly boosts poison or weapon damage the way casters benefit from tailored gear. This is general Classic-era community practice, not something confirmed for Forever specifically.

## Leveling

Combat is generally regarded as the stronger leveling spec for Rogue, since Adrenaline Rush and heavier weapon damage clear trash faster than Assassination's poison-uptime playstyle, which pays off more in sustained single-target fights. The beta caps at level 30, so this is carried over from 1.12 leveling knowledge rather than tested in Forever.
