---
title: Protection Warrior in Forever
classSlug: warrior
spec: protection
role: tank
build: 'FS1:1.60.1.70532:warrior:dwarf:35305013/0/050533120330001311:'
recommendedRaces: [dwarf, tauren]
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
description: 'Talents, tanking priority, stat priority, and race picks for Protection Warrior in Forever, with beta-versus-projection called out.'
updated: 2026-10-07
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

Protection is Forever's tanking spec, using a one-handed weapon and shield to hold enemy attention and survive heavy hits.

### Strengths

- Shield Slam is its highest-threat button and grows with Block Value
- Sunder Armor stacks add threat on every spare global
- Defensive talents keep mitigation and Rage generation high
- Revenge comes up more often the better its avoidance

### Weaknesses

- Gives up the damage of the two-handed and dual-wield specs
- Needs healers to outlast sustained damage
- Threat leans on Sunder Armor filling every other global

### In a raid

- **Role:** Main tank
- **Brings:** Sunder Armor on the target, Thunder Clap's attack-speed slow and Demoralizing Shout

## Talents and builds

This build spends 20 points in Arms and 31 in Protection, the shortest road to Shield Slam. It is the build this site's tank simulator is run on, so everything below is what the simulator models, taken from the client's own talent text.

- **Shield Specialization** — more Block chance, and every Block pays Rage. A block is no longer only mitigation: it feeds the Rage that Shield Slam, Revenge and Heroic Strike spend.
- **Anticipation** — Defense Skill, which raises avoidance and lowers the chance to be critically hit.
- **Improved Revenge**, **Improved Thunder Clap** and **Improved Sunder Armor** — cheaper or harder-hitting versions of the three abilities that make most of a tank's threat.
- **Last Stand** — a short burst of health, then it is lost again. An emergency button, not a damage-taken talent.
- **Master of Defense** — Rage whenever you Dodge or Parry with a shield equipped, the avoidance counterpart of Shield Specialization.
- **Defiance** — more threat in Defensive Stance while a shield is equipped, added on top of the stance's own threat bonus.
- **Focused Rage** — cheaper offensive abilities, so the Rage the talents above produce goes further.
- **Bastion** — more damage from everything you do while a shield is equipped; one point is enough to keep the budget for the rest.
- **Concussion Blow** — one point, because Shield Slam requires it.
- **Shield Slam** — the capstone: a shield hit that adds your Block Value, with very high threat.

**Arms**, the cheapest useful twenty: Deflection for Parry, Improved Heroic Strike for cheaper Heroic Strikes, Improved Rend, Improved Tactical Mastery and Anger Management as the tier gates they are, and Deep Wounds with the points that are left.

What is left out, and why. Improved Bloodrage, Iron Will, Improved Disarm, Vanguard and Improved Shield Bash change nothing against a boss that stays in front of you and neither stuns, fears, silences nor disarms, which is the fight this site simulates. Improved Shield Wall shortens a cooldown the simulated fight never needs to use. Devastate exists in the client's spell list but is not a learnable ability in this build, so no rotation and no talent here depends on it.

Open the planner at [/planner?class=warrior](/planner?class=warrior) to build this out.

## Rotation and priority

This is the rotation the simulator runs, and each line in it was kept because taking it out made the tank worse in the simulator's tank fight (a level 63 boss that swings at you every two seconds, with healers assumed). You stand in Defensive Stance with a shield for the whole fight.

1. **Shield Slam** on cooldown. It is your highest-threat button and the one whose damage grows with Block Value.
2. **Revenge** whenever it is available. It becomes castable after you dodge, parry or block, so the better your avoidance, the more often it is up.
3. **Shield Block** whenever its buff is down. Of everything you press, it is the one whose removal costs the most damage taken.
4. **Thunder Clap** held on the boss. Its attack-speed slow is worth the Rage for the damage it removes, at a small cost in threat.
5. **Sunder Armor** as the filler for every other global cooldown. Stacking it is the biggest single source of threat outside Shield Slam and Revenge.
6. **Heroic Strike** with spare Rage. At this gear a Forever tank is limited by cooldowns rather than Rage, so the exact threshold you queue it at barely matters.

Last Stand and Shield Wall are for a fight the healers are losing; the simulator fires them on low health and they cut the chance of death on a harder boss, but they never fire on the standard fight. Demoralizing Shout is not in the list because it lowers the boss's attack power and the simulated boss has none to lower; against a real boss with attack power it earns its place. Taunt, Disarm and Concussion Blow do nothing in a fight with one stationary enemy and are left out.

Rage comes from your swings (a fixed amount per landed swing from the weapon's speed, more on a critical strike), from the damage you take, and from Shield Specialization and Master of Defense, so a tank that is being hit and is avoiding well is never short of it.

## Stat priority

The table above is this level range's own tank simulation at level 60, re-run by the nightly pipeline whenever the build or its gear data changes; these numbers are never hand-entered. A tank is scored on one number that rewards, in this order, effective health (how much of the boss's damage your health pool can take once armor and avoidance have done their work), then a risk index that punishes spiky incoming damage, then threat. Each stat in the table is how much one point of it moves that score, shown in the same convention as every other spec: against the stat the table is anchored to.

Why the stats land where they do is mechanical, and stays true when the numbers move. Stamina is health, and health is a straight multiplier on effective health. Armor lowers every physical hit by a fraction that shrinks as you stack it. Defense Skill, Dodge and Parry take hits out entirely, and Defense also lowers the chance to be critically hit and crushed. Block takes a flat amount off the hits it catches, so it depends on your Block Value. Expertise lowers the boss's chance to dodge or parry you, which is threat and Rage, not survival; no item carries it yet, so it is a stat for the day one does. Strength and Attack Power add threat and, for Strength, Block Value.

Two things the table cannot tell you. The item tables used here fold a shield's Block Value into its Block line, so a shield's Block Value is not weighed apart from its Block chance. And a stat shown as "not separable from zero" is not a verdict that it is worthless; the simulation's own error on it is too wide, at this level range's sample size, to tell its value apart from zero.

## Best in slot

Follow the table above: health, then the avoidance stats and armor, with the threat stats behind them.

## Enchants and consumables

Target Stamina and Defense Skill on enchants where they're available, rather than the Strength and Attack Power enchants the two DPS specs want.

## Races

Protection can be played by every race the Warrior class allows: Human, Dwarf, Night Elf, and Gnome on the Alliance; Orc, Undead, Tauren, and Troll on the Horde; and, new in Forever, the Skyborne on either faction, since Skyborne is a new race rather than one of Blizzard's six new class pairings. For Alliance, **Dwarf** is the strongest pick for a tank: the reworked Stoneform now reduces Physical damage taken instead of only raising Armor, a direct, on-demand mitigation cooldown that fits a spec absorbing sustained physical hits. For Horde, **Tauren** pairs well: War Stomp stuns up to 5 nearby enemies within 8 yards for 2 seconds on a 2-minute cooldown, useful for peeling adds off healers, and Endurance grants 5% more Health and 1% Hit chance, directly growing the health pool and threat reliability a tank leans on. Skyborne is open to Warrior on both factions with no restriction; its racials (a downward glide, a short full health-and-mana restore, 1% haste) read as utility rather than survivability, so it's a valid but not obviously stronger pick than Dwarf or Tauren for Protection specifically.

## Professions

Blacksmithing and Mining is also the standard convention for a tanking Warrior, with an extra reason beyond the DPS specs: Blacksmithing's socket-adding recipes let a tank add Stamina or Defense gems to gear that doesn't already have a socket.

## Leveling

Protection is not generally the leveling spec of choice for Warrior — Fury, or Arms to a lesser extent, clears trash faster while leveling solo, and Protection's strengths show up in a group or raid tanking role instead.
