---
title: Protection Paladin in Forever
classSlug: paladin
spec: protection
role: tank
build: 'FS1:1.60.1.70707:paladin:dwarf:05005/15302102513501031/5003:'
recommendedRaces: [dwarf, undead]
statPriority:
  [
    Stamina,
    Armor,
    Defense,
    Dodge,
    Parry,
    Block,
    Strength,
    Agility,
    Attack power,
    Hit,
    Critical strike,
    Expertise,
  ]
description: 'Talents, rotation, stats, gear, races, and professions for Protection Paladin tanking in Forever.'
updated: 2026-10-31
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

Protection is the Paladin tanking tree: Seal of Fury and Judgement hold threat from range, while Holy Shield, block and avoidance absorb damage.

### Strengths

- Judgement taunts from farther away than other tanks can
- Seal of Fury gives a tanking Seal with a small absorb on every proc
- Holy Shield raises block and adds threat on every block
- Righteous Fury lowers damage taken, and Templar's Bulwark gives an emergency absorb

### Weaknesses

- Mana, not cooldowns, limits the rotation
- The emergency defensives share one Forbearance timer
- Hit is needed for Judgement and Holy Strike to land and hold threat

### In a raid

- **Role:** Main tank; pulls threat and holds a boss facing away from the raid
- **Brings:** Devotion Aura for the group's armor; no blessing of its own from the Protection tree

## Talents and builds

This build spends 33 points in Protection, 8 in Retribution and 10 in Holy, and it is the build this site's tank simulator runs, taken from the client's own talent text. Pricing each choice in the simulator's tank fight is what put the points past the Protection core into Deflection, Sacred Duty, Iron Creed and Shieldsworn.

- **Divine Intellect** — one point on the first row, for a little more mana.
- **Redoubt** — a chance, whenever a melee attack damages you, to gain extra Block chance for a short window or until it has absorbed a few blocks.
- **Precision** — hit, so the Judgement and Holy Strike that carry your threat land.
- **Sacred Duty** — shorter cooldowns on your three emergency defensives, and a shorter Forbearance.
- **Templar's Bulwark** — an activated absorb shield worth your whole health bar, followed by Forbearance. Swift Judgement requires it.
- **Shield Specialization** — a stronger shield absorb and mana back on a Block, no more than once every few seconds.
- **Toughness** — more armor from your items.
- **Swift Judgement** — a second Judgement of your Seal on its own one minute cooldown, free and off the global cooldown.
- **Improved Righteous Fury** — while Righteous Fury is on, all damage you take is reduced. In Forever it is a mitigation talent, not a threat one.
- **Anticipation** — Defense Skill, which raises avoidance and lowers the chance to be critically hit.
- **Holy Shield** — now on the 21-point row: more Block for a short window, Holy damage on every Block, and a little extra threat on that damage.
- **Iron Creed** — more threat on Holy Strike, and while Righteous Fury is on, Holy Strike also cuts the damage you take for a few seconds.
- **Shieldsworn** — the capstone: Holy Strike's cooldown drops to six seconds, which keeps Iron Creed up all the time, and each Holy Strike also fires Holy Shield's damage and Shield Specialization's mana while Holy Shield is up. It is the largest single gain in the build, mostly in threat.

**Retribution**: five points in **Deflection**, five percent more Parry and the only parry a paladin can get from a talent, then **Conviction** for critical strike on every attack and spell.

**Holy**: **Divine Strength** for the raw Strength, and **Improved Seals** for more Seal of Fury and Judgement damage.

Left out to fit the budget: Reckoning (an extra attack after you block or are critically hit; priced in the simulator it adds less than Conviction), Stoicism, Guardian's Favor and Improved Hammer of Justice (utility). Blessing of Sanctuary is not in this build of the game, so the Protection tree gives a paladin no blessing of its own; the paladin's own contribution to its armor is Devotion Aura, which the simulator keeps on.

Open the planner at [/planner?class=paladin](/planner?class=paladin) to build this out.

## Rotation and priority

This is the rotation the simulator runs. It was measured in the simulator's tank fight (a level 63 boss that swings at you every two seconds, healers assumed), and mana, not cooldowns, is what limits it, so the order below is also the order in which mana is spent.

1. **Righteous Fury** is a setting, not a button: keep it on. It multiplies the threat of your Holy damage, and with Improved Righteous Fury it also lowers the damage you take. A Righteous-Fury-less rotation of the same gear makes far less threat.
2. **Seal of Fury** up. Judgement no longer consumes your seal in Forever, so you recast the seal only when it runs out, not after every Judgement.
3. **Holy Shield** on cooldown. It is both the best mitigation you have and the best use of your mana, so every other spell leaves its cost in the pool.
4. **Hammer of the Righteous** and **Holy Strike** on their cooldowns, the first of which is the largest single threat gain in this list; with Shieldsworn, Holy Strike comes back every six seconds.
5. **Judgement** on cooldown, and **Swift Judgement** whenever it is ready: it is a second Judgement on its own cooldown.
6. **Consecration** only with spare mana, where it pays for itself in threat.

Templar's Bulwark, Divine Protection and Divine Shield are for a fight you are losing: the simulator uses them on low health, where they cut the chance of death on a harder boss, and they never fire on the standard fight. Divine Shield and Divine Protection put Forbearance on you, which shares a timer with Templar's Bulwark.

## Stat priority

A tank is scored on one number that rewards, in this order, effective health (how much of the boss's damage your health pool can take once armor and avoidance have done their work), then a risk index that punishes spiky incoming damage, then threat. Each stat in the table is how much one point of it moves that score, shown in the same convention as every other spec: against the stat the table is anchored to.

Why the stats land where they do is mechanical, and stays true when the numbers move. Stamina is health, and health is a straight multiplier on effective health. Armor lowers every physical hit by a fraction that shrinks as you stack it. Defense Skill, Dodge and Parry take hits out entirely, and Defense also lowers the chance to be critically hit and crushed. Block takes a flat amount off the hits it catches and powers Holy Shield's damage and Redoubt. Strength adds threat and Block Value; Hit keeps Judgement and Holy Strike landing; Expertise lowers the boss's chance to dodge or parry you, which is threat, not survival, and no item carries it yet.

Two things the table cannot tell you. The item tables used here fold a shield's Block Value into its Block line, so Block Value is not weighed apart from Block chance.

## Best in slot

## Enchants and consumables

## Races

## Professions

Blacksmithing suits Protection well since it can add sockets to armor pieces, letting a tank fill in stamina or defense the itemization doesn't provide on its own. Engineering offers utility trinkets and gadgets a tank can use situationally.

## Leveling

Retribution is the stronger leveling spec for Paladin — a shield and a threat rotation built for holding aggro on one enemy is a poor fit for the pace of solo questing. If leveling Protection anyway, expect it to feel closer to a Warrior's tank stance than a fast solo spec.
