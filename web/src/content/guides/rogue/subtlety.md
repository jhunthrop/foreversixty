---
title: Subtlety Rogue in Forever
classSlug: rogue
spec: subtlety
role: dps
build: 'FS1:1.60.1.69893:rogue:night-elf:321/32003/532322131000300105:'
recommendedRaces: [night-elf, troll]
statPriority: [Attack power, Agility, Critical strike, Hit, Melee haste]
description: 'Subtlety Rogue overview, talent priority, rotation, stat weights, and race picks for Forever, with beta-versus-projection called out.'
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

Subtlety trades Combat's raw weapon damage and Assassination's poison uptime for cheap, mobile combo generation and stealth-oriented utility, using Hemorrhage as a bleed-based builder instead of Sinister Strike. It's historically the least raid-damage-competitive of the three Rogue trees in the 1.12 era, and nothing in the Forever demo data suggests that's changed at the tree-design level. The beta caps at level 30, so how Subtlety's damage actually compares to Assassination or Combat at level 60 or in a raid is a projection from demo-tree reading and 1.12 knowledge, not confirmed play.

## Talents and builds

In rough priority order: **Opportunity** raises the damage of Backstab, Ambush, and Mutilate, which matters even for a Hemorrhage-based build since Ambush remains the spec's stealth opener. **Initiative** adds a chance at an extra combo point when opening with Cheap Shot or a similar ability, accelerating the path to a finisher. **Ghostly Strike** is a cheap, high-damage-relative-to-cost filler that also applies its own attack power debuff. **Serrated Blades** adds armor penetration and boosts finisher damage, a flat multiplier once you're spending combo points regularly. **Hemorrhage**, the tree's signature ability, replaces Sinister Strike as the combo builder and needs its bleed reapplied once its charges run out. **Cutthroat**, new to the tree in Forever, occasionally lets a Backstab set up your following Ambush so it lands without needing Stealth first — a build-around for repeated openers mid-fight, though it isn't verified in this site's simulated rotation (see Rotation below).

Point allocation runs deep into Subtlety to reach the bottom-row talents, with the remainder split as a handful of points in Combat for utility (Deflection) and a few in Assassination for crit. Open the planner at [/planner?class=rogue](/planner?class=rogue) to build this out.

## Rotation and priority

This site's simulator plays the fully supported version of the loop: keep Slice and Dice active so combo points aren't wasted on idle swing time, use Ghostly Strike on cooldown since it's cheaper energy than Backstab and carries its own debuff, spend five combo points on Eviscerate, and use Hemorrhage as the combo-point builder, reapplying its bleed once the charges expire. Cutthroat's proc-based free Ambush isn't reflected in that priority list yet, so treat any Ambush-heavy opener sequencing as this site's own projection of the tree's intent rather than a confirmed rotation.

## Stat priority

In simulator-derived priority order: **attack power** first, driving every physical attack in the loop. **Agility** next for its indirect attack power and crit. **Crit** feeds Seal Fate-style bonus combo points if talented and raises Hemorrhage and Eviscerate damage directly. **Hit** keeps the combo-point loop consistent by avoiding misses. **Melee haste** last, for more swings and faster combo generation.

## Gear

Look for agility and attack power first, then the merged hit and crit ratings, then melee haste. Since Hemorrhage's bleed needs periodic reapplication rather than constant uptime like a DoT, raw weapon damage and attack power outweigh haste more than in a pure swing-speed spec. A new stat reducing target dodge and parry chance may show up on gear as itemization fills in. Specific pre-raid or raid-tier item picks aren't something this site can name with confidence yet: the beta caps at level 30 and this build's raid loot tables are still missing most items the community's sourcing expects, with nothing raiding in-game until the first tier opens on 9 December 2026.

## Enchants and consumables

Weapon and glove enchants should target agility and attack power; **Enchant Gloves - Superior Agility** and **Enchant Weapon - Agility** are both verified in this build's enchant data. **Elixir of the Mongoose** (agility) is a verified consumable option. A fuller consumable stack isn't something this site is naming until the class is played past level 30.

## Races

Subtlety can be played by every race the Rogue class allows on both factions, plus the Skyborne on either faction since Rogue's Skyborne availability has no faction restriction. For Alliance, **Night Elf** is the clearest fit for a stealth-oriented spec: Shadowmeld works in combat on a 2-minute cooldown, giving Subtlety a second stealth tool beyond its own Vanish and Preparation for repositioning into a fresh opener. For Horde, **Troll** pairs well through Berserking's 10% attack and casting speed on a 3-minute cooldown, which lines up with a burst window after an opener lands. Rogue is open to the Skyborne on both factions with no class restriction; their read racials skew toward mobility and resource restoration rather than the stealth-and-burst identity Subtlety leans on, making them a workable but not clearly stronger pick than Night Elf or Troll here.

## Professions

Community convention favors **Engineering** for combat utility alongside a gathering profession such as Herbalism or Mining, matching general Rogue practice across all three trees rather than anything Subtlety-specific. This is Classic-era community convention, not confirmed for Forever.

## Leveling

Combat is the generally recommended leveling spec for Rogue over Subtlety, since Adrenaline Rush and consistent weapon damage clear trash faster than Subtlety's stealth-and-reposition playstyle, which pays off more in single-target or PvP scenarios. The beta caps at level 30, so this is carried over from 1.12 leveling knowledge rather than tested in Forever.
