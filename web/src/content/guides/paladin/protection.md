---
title: Protection Paladin in Forever
classSlug: paladin
spec: protection
role: tank
build: 'FS1:1.60.1.70009:paladin:dwarf:55313003/5530513301301001/0:'
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
updated: 2026-10-07
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

This build spends 31 points in Protection to reach Holy Shield and 20 in Holy, and it is the build this site's tank simulator runs, taken from the client's own talent text.

**Protection**, in roughly the order you would take them:

- **Toughness** — more armor from your items.
- **Redoubt** — a chance, whenever a melee attack damages you, to gain extra Block chance for a short window or until it has absorbed a few blocks.
- **Precision** — hit, so the Judgement and Holy Strike that carry your threat land.
- **Anticipation** — Defense Skill, which raises avoidance and lowers the chance to be critically hit.
- **Improved Seal of Fury** — when Seal of Fury's shield is spent, you get mana back; it is the talent that makes the shield a mana source.
- **Improved Righteous Fury** — while Righteous Fury is on, all damage you take is reduced. In Forever it is a mitigation talent, not a threat one.
- **Shield Specialization** — a stronger shield absorb and mana back on a Block, no more than once every few seconds.
- **Swift Judgement** — finishes Judgement's cooldown and makes the next one free.
- **One-Handed Weapon Specialization** — more damage with the one-hander you tank with.
- **Templar's Bulwark** — an activated absorb shield worth a large share of your health, followed by Forbearance. It is one point because Holy Shield requires it.
- **Holy Shield** — the 31-point talent: more Block for a short window, Holy damage on every Block, and a little extra threat on that damage.

**Holy**, the twenty that remain: Divine Strength and Divine Intellect for the raw stats, Improved Seals for more Seal of Fury and Judgement damage, Reverence for mana while you fight, and the tier gates they stand on.

Left out to fit the budget: Reckoning (an extra attack after you block or are critically hit), Iron Creed (more threat on Holy Strike, and less damage taken after it while Righteous Fury is on) and Sacred Duty (more stamina and shorter defensive cooldowns). Each is worth a look as you gain points; the simulator can price them in the planner. Blessing of Sanctuary is not in this build of the game, so the Protection tree gives a paladin no blessing of its own.

Open the planner at [/planner?class=paladin](/planner?class=paladin) to build this out.

## Rotation and priority

This is the rotation the simulator runs. It was measured in the simulator's tank fight (a level 63 boss that swings at you every two seconds, healers assumed), and mana, not cooldowns, is what limits it, so the order below is also the order in which mana is spent.

1. **Righteous Fury** is a setting, not a button: keep it on. It multiplies the threat of your Holy damage, and with Improved Righteous Fury it also lowers the damage you take. A Righteous-Fury-less rotation of the same gear makes far less threat.
2. **Seal of Fury** up. Judgement no longer consumes your seal in Forever, so you recast the seal only when it runs out, not after every Judgement.
3. **Holy Shield** on cooldown. It is both the best mitigation you have and the best use of your mana, so every other spell leaves its cost in the pool.
4. **Hammer of the Righteous** and **Holy Strike** on their cooldowns, the first of which is the largest single threat gain in this list.
5. **Judgement** on cooldown; with Swift Judgement the next one is free.
6. **Consecration** only with spare mana, where it pays for itself in threat.

Templar's Bulwark, Divine Protection and Divine Shield are for a fight you are losing: the simulator uses them on low health, where they cut the chance of death on a harder boss, and they never fire on the standard fight. Divine Shield and Divine Protection put Forbearance on you, which shares a timer with Templar's Bulwark.

## Stat priority

The table above is this band's own tank simulation at level 60, re-run by the nightly pipeline whenever the build or its gear data changes; these numbers are never hand-entered. A tank is scored on one number that rewards, in this order, effective health (how much of the boss's damage your health pool can take once armor and avoidance have done their work), then a risk index that punishes spiky incoming damage, then threat. Each stat in the table is how much one point of it moves that score, shown in the same convention as every other spec: against the stat the table is anchored to.

Why the stats land where they do is mechanical, and stays true when the numbers move. Stamina is health, and health is a straight multiplier on effective health. Armor lowers every physical hit by a fraction that shrinks as you stack it. Defense Skill, Dodge and Parry take hits out entirely, and Defense also lowers the chance to be critically hit and crushed. Block takes a flat amount off the hits it catches and powers Holy Shield's damage and Redoubt. Strength adds threat and Block Value; Hit keeps Judgement and Holy Strike landing; Expertise lowers the boss's chance to dodge or parry you, which is threat, not survival, and no item carries it yet.

Two things the table cannot tell you. The item tables used here fold a shield's Block Value into its Block line, so Block Value is not weighed apart from Block chance. And a stat shown as "not separable from zero" is not a verdict that it is worthless; the simulation's own error on it is too wide, at this band's sample size, to tell its value apart from zero.

## Gear

Follow the table above: health, then the avoidance stats and armor, with the threat stats behind them. This site's ranker builds the level 60 set for you in the BiS tab, scored on the same tank fight, in a shield and a one-handed weapon. Specific raid picks cannot be named honestly yet: nothing raids until 9 December, and this build's raid loot tables are only partly itemized for Forever, so the set is built from what the item tables do carry.

## Enchants and consumables

Target stamina and the defensive stats on enchants where they are available. This build's enchant data confirms a dedicated Enchant Shield line for stamina. For consumables the BiS tab's Raid-ready set shows the tank's: a health flask, stamina food and drink, and strength and attack power for threat.

## Races

Protection Paladin can only be Human or Dwarf on Alliance, or Undead on Horde — this build's race and class combination table confirms no other race can train Paladin at all. Between the two Alliance options, Dwarf is the better tanking pick: its reworked Stoneform now reduces physical damage taken for its duration instead of only raising armor, which is a direct mitigation cooldown a tank can time to a dangerous swing, and Mace Specialization adds crit chance that feeds Reckoning's extra-attack proc if paired with a one-handed mace and shield. Human's Sword Specialization only helps if wielding a sword, and its Spirit bonus does little for a tank build. On Horde, Undead is currently the only playable Paladin race; Will of the Forsaken's fear, charm, and sleep cleanse is useful raid utility for a tank that needs to stay locked onto a boss, even though it no longer grants brief immunity the way it did in 1.12.

## Professions

Blacksmithing suits Protection well since it can add sockets to armor pieces, letting a tank fill in stamina or defense the itemization doesn't provide on its own. Engineering offers utility trinkets and gadgets a tank can use situationally. This is general Classic-era community convention rather than anything Forever-specific, since profession bonuses have not been shown to change for Forever.

## Leveling

Retribution is the stronger leveling spec for Paladin — a shield and a threat rotation built for holding aggro on one enemy is a poor fit for the pace of solo questing. If leveling Protection anyway, expect it to feel closer to a Warrior's tank stance than a fast solo spec. All of this is restated 1.12 knowledge rather than tested Forever guidance: the beta caps at level 30, well short of where a full leveling route would matter.
