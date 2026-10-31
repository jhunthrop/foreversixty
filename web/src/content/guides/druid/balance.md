---
title: Balance Druid in Forever
classSlug: druid
spec: balance
role: dps
build: 'FS1:1.60.1.70658:druid:night-elf:5232221115400051/05/503301:'
raidBuild: 'FS1:1.60.1.70658:druid:night-elf:4132220115501051/05/5053:'
recommendedRaces: [night-elf, tauren]
statPriority:
  [Intellect, Spell power, Nature power, Hit, Arcane power, Critical strike, Spell haste, Spell penetration]
description: 'Talents, rotation, stats, and gear for Balance Druid in Forever, and what is confirmed versus projected from the beta.'
updated: 2026-09-24
confidence: inferred
sources:
  - label: 'Blizzard, Deep Dive panel recap'
    url: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
    kind: blizzard
  - label: 'Talents Forever (demo transcription)'
    url: https://talentsforever.com/data.json
    kind: community
  - label: 'Forever Sixty stat weights, data/curated/stat-weights.json'
    url: https://foreversixty.gg/sources
    kind: site
  - label: 'Forever Sixty rotation data, data/curated/apl/druid-balance.json'
    url: https://foreversixty.gg/sources
    kind: site
  - label: 'Forever Sixty raid loot data, data/curated/loot/forever-raid-phases.json'
    url: https://foreversixty.gg/sources
    kind: site
---

## Overview

Balance is the Druid ranged caster tree: Starfire, Moonfire and Insect Swarm damage from Moonkin Form, which gives the party a critical strike aura.

### Strengths

- Moonkin Form raises party critical strike chance by 3% within 45 yards
- Vengeance doubles critical damage on Arcane and Nature spells
- Starfire and Wrath cover both long and short windows
- Instant Moonfire finishes fights when no other cast will land

### Weaknesses

- Starfire's 3.5 second cast leaves little room to move
- Two DoTs to keep up alongside a long Starfire cast
- Damage is confined to the Nature and Arcane schools

### In a raid

- **Role:** Ranged caster damage
- **Brings:** Moonkin Form, a 3% critical strike chance aura for party members within 45 yards

## Talents and builds

- **Improved Wrath** — up to 0.5 seconds off Wrath's cast time and 50% off its mana cost at rank 5, a heavy early investment in the tree's other nuke.
- **Vengeance** — up to a 100% critical-strike-damage bonus at rank 5 on Arcane and Nature spells, the tree's biggest single multiplier.
- **Improved Starfire** — up to 0.5 seconds off Starfire's cast time and a 15% chance to stun the target at rank 5; this build takes rank 4 (0.4 seconds, 12% stun chance) to free a point for Insect Swarm below, since a stun rarely lands on a raid boss anyway.
- **Insect Swarm** (1 point) — a second cheap DoT the rotation keeps running alongside Moonfire; skipping it was an earlier error in this build, fixed by moving a point off Improved Starfire's last rank.
- **Nature's Splendor** (1 point) — the one-point talent that the simulator credits with the largest single gain in this tree, taken with the point Nature's Grace used to hold.
- **Moonfury** — up to 10% more Arcane and Nature damage at rank 5.
- **Moonkin Form** — the capstone: 360% more armor from items while shapeshifted, a doubled Omen of Clarity proc chance, and 3% more critical strike chance for party members within 45 yards.

## Rotation and priority

Be in Moonkin Form before the pull and stay in it: the form gives you 3% more critical strike chance, and every spell in this loop can be cast from it. Keep Moonfire and then Insect Swarm ticking on the target: with Improved Moonfire and Nature's Splendor, a Moonfire does more damage per second of casting than Starfire, and a raid-buffed Balance druid has the mana to keep both up. Starfire is everything else: cast it on repeat for the rest of the fight. An earlier search dropped Moonfire upkeep to save mana, but that run's character was running dry; with raid buffs and consumables, keeping Moonfire up gained about 5% in our level-60 raid run. Near the end of a fight, adjust for what will actually land: with less than 3.5 seconds remaining, switch to Wrath, since its 2-second cast still finishes where Starfire's 3.5-second cast would not; with less than 1.5 seconds left, only the instant Moonfire will land in time. Eclipse, if talented, rewards weaving in a Wrath cast periodically, since it shortens your next two Starfire casts, so it is worth folding into the priority once you have taken it.

## Stat priority

## Best in slot

## Enchants and consumables

## Races

## Professions

Balance benefits from the same caster-support professions as any spellcaster: Tailoring for crafted spellcaster gear and Enchanting to self-apply spell power enchants are the general Classic-era convention.

## Leveling

Balance is a reasonable leveling spec on its own, with Moonfire and Starfire providing ranged damage, but Feral is generally considered the stronger overall leveling choice for Druid thanks to Cat Form's mobility and Bear Form's durability.
