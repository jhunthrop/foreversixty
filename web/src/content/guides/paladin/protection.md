---
title: Protection Paladin in Forever
classSlug: paladin
spec: protection
role: tank
description: 'Talents, rotation, stats, gear, races, and professions for Protection Paladin tanking in Forever.'
updated: 2026-09-24
confidence: inferred
sources:
  - label: 'Blizzard, Deep Dive panel recap'
    url: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
    kind: blizzard
  - label: 'Forever Sixty build data, data/builds/1.60.1.69893/talents/paladin.json'
    url: https://foreversixty.gg/data/1.60.1.69893/talents/paladin.json
    kind: datamined
  - label: 'Forever Sixty stat weights, data/curated/specs.json'
    url: https://foreversixty.gg/data/curated/specs.json
    kind: site
  - label: 'Forever Sixty raid phase data, data/curated/loot/forever-raid-phases.json'
    url: https://foreversixty.gg/data/curated/loot/forever-raid-phases.json
    kind: site
  - label: 'Forever Sixty build data, data/builds/1.60.1.69893/races.json'
    url: https://foreversixty.gg/data/1.60.1.69893/races.json
    kind: datamined
---

## Overview

Protection is Forever's Paladin tanking tree, built around Seal of Fury for threat and mana return, Judgement as a ranged taunt, and Holy Shield for block uptime once a target is engaged. Blizzard confirmed Seal of Fury as a new addition specifically to give Protection a proper tanking Seal, with its Judgement taunting from farther away than other tanks get. In a group it holds the same job Protection has always held — pull threat, absorb damage, and keep a boss facing away from the raid — but the tools for doing it have changed since 1.12. The beta caps at level 30, so nothing here about level 60 raid tanking, threat ceilings against real boss damage, or best-in-slot gear has actually been played; it is a projection from the demo talent trees and 1.12 Protection Paladin knowledge, confirmed only where Blizzard's own recap says so directly.

## Talents and builds

Verified against this build's own Paladin talent data:

1. **Improved Seal of Fury** — when Seal of Fury's absorb shield is fully spent, it restores mana scaled by how far above your level the attacker is, which is Protection's core sustain tool given how often a tank eats big hits.
2. **Shield Specialization** — increases the shield's absorb amount and gives blocks a chance to restore mana, stacking with Improved Seal of Fury for mana return on both sides of a hit.
3. **Swift Judgement** — resets Judgement's cooldown and makes the next cast free, keeping the taunt and mana-return loop running without waiting out the full cooldown.
4. **Templar's Bulwark** — an activated cooldown granting a large absorb shield at the cost of Forbearance for a minute, a panic button for a dangerous hit rather than something to lean on constantly.
5. **Iron Creed** — Holy Strike generates more threat, and while Righteous Fury is active it also reduces damage taken for a few seconds after each cast, tying Protection's baseline attack directly into its survivability.
6. **Holy Shield** — the tree's capstone talent; it raises block chance for a window and deals Holy damage on every block during it, which is both a mitigation cooldown and one of Protection's better threat tools.

A rough point split at level 60 would put close to 31 points in Protection to reach Holy Shield, with the remainder split between Holy for early mana-sustain talents like Reverence and Retribution for Vindication's attack power debuff, depending on whether mana or threat is the bigger problem in a given fight. That split is a projection — the beta cap of 30 has not let anyone test it. Open the planner at [/planner?class=paladin](/planner?class=paladin) to build this out.

## Rotation and priority

This site's own rotation data for Protection Paladin is an unwritten stub as of this writing, so the following is written from general 1.12 Protection Paladin practice plus the baseline changes Forever confirmed, not from a simulated priority list. Open with Seal of Fury up before pulling, since its Judgement is both the taunt and the mana-return trigger. Judgement is free of its old cost of consuming the Seal, so fire it on cooldown for the taunt range, the mana return, and a bit of extra threat, then keep swinging while weaving in Holy Strike, which is baseline from level 6 and hits on its own twelve-second cooldown independent of the global cooldown. Consecration is baseline from level 20 and worth dropping under a pack for cleave threat, since it hits everything inside it lightly but the first few targets much harder, which keeps a Protection Paladin from accidentally pulling an entire room. Holy Shield is worth keeping active whenever the charges are available, both for the block chance and for the threat its damage generates, and Templar's Bulwark is best held for a hit that would otherwise be dangerous rather than used on cooldown, since it locks out Forbearance for a full minute afterward.

## Stat priority

Ordered by this site's own simulator-derived weights:

1. **Attack power** — the primary driver of Protection's melee and Seal of Fury damage, which in turn drives threat.
2. **Strength** — converts directly into attack power and adds a small amount of block value.
3. **Agility** — adds armor, crit, and dodge, all of which reduce incoming damage or add threat.
4. **Crit** — feeds Reckoning's extra-attack chance and Holy Shield's damage-on-block, both threat tools.
5. **Hit** — keeps Judgement and Holy Strike landing reliably, since a missed Judgement is a missed taunt.
6. **Melee haste** — more swings per minute, which is a smaller factor than the stats above it once basic threat and survivability are covered.

## Gear

Look for pieces that lead with attack power and strength, backed by stamina and the defensive stats a shield tank needs, matching the stat priority above. Beyond that principle, this section cannot get specific yet: the beta caps at level 30, nothing raids until the first tier opens on 9 December, and this build's own item data has most raid loot tables only partially re-itemized or not itemized at all for Forever, so a real Protection pre-raid or first-raid gear list would be guessing. This section fills in once raid loot data lands.

## Enchants and consumables

Target strength and stamina on weapon, bracer, and glove enchants, matching the stat priority above. This build's enchant data confirms Strength options for both bracers and gloves and a dedicated Enchant Shield line for stamina, so a Protection Paladin has verified enchant slots to chase its top stats and survivability. For consumables, look for elixirs or flasks that boost defense, strength, or stamina over general-purpose ones; this build's consumable data lists defense- and strength-specific elixirs, though which one is strongest for Forever's tank itemization isn't something this site can state with confidence yet.

## Races

Protection Paladin can only be Human or Dwarf on Alliance, or Undead on Horde — this build's race and class combination table confirms no other race can train Paladin at all. Between the two Alliance options, Dwarf is the better tanking pick: its reworked Stoneform now reduces physical damage taken for its duration instead of only raising armor, which is a direct mitigation cooldown a tank can time to a dangerous swing, and Mace Specialization adds crit chance that feeds Reckoning's extra-attack proc if paired with a one-handed mace and shield. Human's Sword Specialization only helps if wielding a sword, and its Spirit bonus does little for a tank build. On Horde, Undead is currently the only playable Paladin race; Will of the Forsaken's fear, charm, and sleep cleanse is useful raid utility for a tank that needs to stay locked onto a boss, even though it no longer grants brief immunity the way it did in 1.12.

## Professions

Blacksmithing suits Protection well since it can add sockets to armor pieces, letting a tank fill in stamina or defense the itemization doesn't provide on its own. Engineering offers utility trinkets and gadgets a tank can use situationally. This is general Classic-era community convention rather than anything Forever-specific, since profession bonuses have not been shown to change for Forever.

## Leveling

Retribution is the stronger leveling spec for Paladin — a shield and a threat rotation built for holding aggro on one enemy is a poor fit for the pace of solo questing. If leveling Protection anyway, expect it to feel closer to a Warrior's tank stance than a fast solo spec. All of this is restated 1.12 knowledge rather than tested Forever guidance: the beta caps at level 30, well short of where a full leveling route would matter.
