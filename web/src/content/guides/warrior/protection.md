---
title: Protection Warrior in Forever
classSlug: warrior
spec: protection
role: tank
description: 'Talents, tanking priority, stat priority, and race picks for Protection Warrior in Forever, with beta-versus-projection called out.'
updated: 2026-09-24
confidence: inferred
sources:
  - label: 'Blizzard, Deep Dive panel recap'
    url: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
    kind: blizzard
  - label: 'Talents Forever (demo transcription)'
    url: https://talentsforever.com/data.json
    kind: community
  - label: 'Icy Veins, the racial rework'
    url: https://www.icy-veins.com/wow-forever/news/the-racial-rework-gives-every-race-a-reason-to-matter-in-warcraft-forever/
    kind: community
  - label: 'Forever Sixty stat weights, data/curated/specs.json'
    url: https://foreversixty.gg/sim/weights
    kind: site
  - label: 'Forever Sixty rotation data, data/curated/apl/warrior-protection.json'
    url: https://foreversixty.gg/sim/specs
    kind: site
  - label: 'Forever Sixty raid loot data, data/curated/loot/forever-raid-phases.json'
    url: https://foreversixty.gg/sources
    kind: site
  - label: 'This site, character planner'
    url: https://foreversixty.gg/planner
    kind: site
---

## Overview

Protection is Forever's tanking spec, built around a one-handed weapon and shield rather than the two-handed or dual-wield choices Arms and Fury make. In a raid or group it holds enemy attention through Sunder Armor stacks and its own capstone finisher, Shield Slam, while a shield and a tree full of defensive talents keep its mitigation and Rage generation high enough to survive hits the other two specs wouldn't need to plan around. The beta only reaches level 30, so nothing below about level 60 play, raid tanking, or best-in-slot gear has actually been tested — it's a projection from the demo trees and 1.12 knowledge, not confirmed play.

## Talents and builds

Verified against this build's talent data, in roughly the order you'd take them:

- **Shield Specialization** — up to 5% more Block chance and a guaranteed 5 Rage whenever you Block at max rank, a Rage-generation talent as much as a mitigation one.
- **Anticipation** — up to 20 Defense Skill at max rank, which raises avoidance across Dodge, Parry, and Block together.
- **Toughness** — up to 10% more Armor from items at max rank, stacking with everything else you wear.
- **Defiance** — up to 15% more threat generated in Defensive Stance while a shield is equipped, on top of the stance's own threat bonus.
- **Concussion Blow** (one point, and the prerequisite for Shield Slam below) — a 5-second stun, useful crowd control on top of being a stepping stone.
- **Shield Slam** (requires 1 point in Concussion Blow) — the capstone: a shield bash dealing damage that scales with Block Value and generating a very high amount of threat.

A build reaching Shield Slam spends roughly 31 points in Protection, with the remainder commonly split into a few Fury points for Cruelty and Iron Will (crit for Rage generation, and shorter Stun and Fear durations) and a couple of Arms points for Deep Wounds. Exact splits vary by preference and aren't fixed by anything confirmed for Forever. Open the planner at [/planner?class=warrior](/planner?class=warrior) to build this out.

## Rotation and priority

This site's rotation data for Protection is unwritten — unlike Arms and Fury, there is no simulator priority list built for this spec yet, so nothing in this section comes from this site's own engine. Carried over from 1.12 tanking convention rather than confirmed for Forever, the general priority opens by stacking Sunder Armor toward its cap for threat and armor reduction, keeps Shield Slam on cooldown as the single highest-threat button, uses Revenge whenever it's available (it only becomes castable right after you Dodge, Parry, or Block an attack), and spends leftover Rage on Heroic Strike or Cleave depending on whether there's one target or several. Thunder Clap and Demoralizing Shout add area threat and a damage-reduction debuff against multiple attackers. Treat this whole section as inferred 1.12 knowledge until this site builds and simulates its own Protection rotation.

## Stat priority

This site's simulator weighting for Protection currently lists the same stats, in the same order, as the two DPS specs:

1. **Attack power**
2. **Strength**
3. **Agility**
4. **Critical strike**
5. **Hit**
6. **Melee haste**

This is an offense-oriented list, not a tanking-specific one built around Stamina, Defense Skill, and Block Value — the kind of priority a level 60 tank actually wants. Treat it as what this site's simulator currently outputs for Protection rather than tested tank stat advice; a Stamina- and avoidance-led list is the likely outcome once tank-specific simulation exists for this spec.

## Gear

Because the stat priority above doesn't yet reflect tanking needs, the itemization principle for Protection right now is qualitative rather than list-driven: prioritize Stamina and Defense Skill to raise health and avoidance, then Block Value and Strength for threat and mitigation, ahead of the raw Attack Power lean the current weighting shows. Specific pre-raid or raid-tier item picks can't be named with confidence yet: the beta caps at level 30, and this build's raid loot tables are themselves incomplete — nothing raids in-game until the first tier opens on 9 December 2026, and even then several of this client's raid item tables are missing most of what earlier sourcing expects. This section fills in once raid loot is itemized and tank-specific weighting exists.

## Enchants and consumables

Target Stamina and Defense Skill on enchants where they're available, rather than the Strength and Attack Power enchants the two DPS specs want. **Enchant Cloak - Superior Defense** and **Enchant Bracer - Superior Stamina** are both verified in this build's enchant data, and **Enchant Shield - Greater Stamina** is a shield-specific option also verified there. Specific consumable stacking for a level-60 tank isn't something this site can confirm yet at level 30.

## Races

Protection can be played by every race the Warrior class allows: Human, Dwarf, Night Elf, and Gnome on the Alliance; Orc, Undead, Tauren, and Troll on the Horde; and, new in Forever, the Skyborne on either faction, since Skyborne is a new race rather than one of Blizzard's six new class pairings. For Alliance, **Dwarf** is the strongest pick for a tank: the reworked Stoneform now reduces Physical damage taken instead of only raising Armor, a direct, on-demand mitigation cooldown that fits a spec absorbing sustained physical hits. For Horde, **Tauren** pairs well: War Stomp stuns up to 5 nearby enemies within 8 yards for 2 seconds on a 2-minute cooldown, useful for peeling adds off healers, and Endurance grants 5% more Health and 1% Hit chance, directly growing the health pool and threat reliability a tank leans on. Skyborne is open to Warrior on both factions with no restriction; its racials (a downward glide, a short full health-and-mana restore, 1% haste) read as utility rather than survivability, so it's a valid but not obviously stronger pick than Dwarf or Tauren for Protection specifically.

## Professions

Blacksmithing and Mining is also the standard convention for a tanking Warrior, with an extra reason beyond the DPS specs: Blacksmithing's socket-adding recipes let a tank add Stamina or Defense gems to gear that doesn't already have a socket. This is general Classic-community practice rather than anything confirmed for Forever specifically.

## Leveling

Protection is not generally the leveling spec of choice for Warrior — Fury, or Arms to a lesser extent, clears trash faster while leveling solo, and Protection's strengths show up in a group or raid tanking role instead. That comparison is carried over from 1.12 knowledge rather than tested — the beta only reaches level 30, so nothing about leveling or tanking above that has actually been played in Forever's client.
