---
title: Feral Druid in Forever
classSlug: druid
spec: feral
role: dps
build: 'FS1:1.60.1.70532:druid:night-elf:01/55232032121032012001/50532:'
raidBuild: 'FS1:1.60.1.70532:druid:night-elf:05002/45211031021032212001/5053:'
recommendedRaces: [night-elf, tauren]
statPriority: [Strength, Agility, Critical strike, Attack power, Feral attack power, Hit, Melee haste]
description: 'Talents, rotation, stats, and gear for Feral Druid in Forever, covering both Cat Form DPS and Bear Form tanking from the one tree.'
updated: 2026-10-07
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
  - label: 'Forever Sixty rotation data, data/curated/apl/druid-feral.json'
    url: https://foreversixty.gg/sources
    kind: site
  - label: 'Forever Sixty raid loot data, data/curated/loot/forever-raid-phases.json'
    url: https://foreversixty.gg/sources
    kind: site
---

## Overview

Feral Combat is one tree for two jobs: Cat Form melee DPS built on Rip and Ferocious Bite, and Bear Form tanking. This guide covers both.

### Strengths

- Cat DPS spends combo points on Rip and Ferocious Bite
- Shifting Power refills 40 energy for mana every 16 seconds
- Tiger's Fury refunds energy on a 30 second cooldown
- The same tree tanks in Bear Form, so one Druid can fill either role

### Weaknesses

- The cat and bear builds split the 51 points differently, so switching roles means a respec
- A cat gains nothing from Thick Hide's Armor
- A bear has no shield, Parry or Block

### In a raid

- **Role:** Melee DPS in Cat Form, or main tank in Bear Form
- **Brings:** Leader of the Pack, a critical strike aura for the group

## Talents and builds

Two structural changes affect this tree specifically: Feral Charge splits into a Cat version, a gap-closing leap that dazes on landing, sitting next to the older Bear version built for interrupting; and Furor's cat-form energy refund now scales with how long you'd been out of form beforehand, rather than handing back a flat amount every time.

- **Ferocity** — up to 5 less Rage or Energy on Maul, Mangle, Swipe, Claw, and Rake at rank 5.
- **Thick Hide** — extra base armor while shapeshifted, scaling with defense skill, essential for a Bear tank and useful padding for a Cat.
- **Savage Fury** — up to 10% more damage at rank 2 on Claw, Rake, Shred, Maul, and Swipe.
- **Predatory Strikes** — up to 150% of your level added to melee attack power in Cat or Bear Form at rank 3.
- **Blood Frenzy** — up to a 100% chance at rank 2 for bonus Rage on a Bear crit, and a 100% chance for a bonus combo point on a Cat crit, supporting both forms from one talent. (Named "Primal Fury" in an earlier draft of this guide; this build's own data calls it Blood Frenzy, so the talent section and `build:` string below use that name.)
- **Berserk** — the capstone: Mangle hits up to 3 targets with no cooldown, and critical strike chance on combo-point generators rises 100%, for 15 seconds.

## Rotation and priority

**Cat DPS**: open with Tiger's Fury, since it costs nothing and refunds energy on its own 30-second cooldown, and keep it on cooldown throughout. Leave Rake out: it costs 40 energy for a small bleed and one combo point, and in this site's simulator that energy is worth more spent on Shred and Rip, which measured about 5% better on the bare character and about 9% better in the raid-ready run. At 5 combo points, cast Rip if it isn't already running; once Rip is up, spend 5 combo points on Ferocious Bite instead of letting them cap out uselessly. Shred fills every other global, building combo points the rest of the time.

**Shifting Power**: this is the new energy tool in the tree, replacing King of the Jungle. It is instant, turns a slice of your base mana into 40 energy, and has a 16 second cooldown (Improved Shifting Power takes 4 or 8 seconds off that). Press it every time it comes off cooldown, but only when your energy is 60 or lower, so the whole 40 fits under the 100 cap; if you are above that, spend a few energy on Shred first and press it on the next global. It sits ahead of Tiger's Fury and every builder in the order. The mana is otherwise unused in cat form, and Natural Shapeshifter makes it cheaper. Blizzard has not published whether it shares the global cooldown; our simulator treats it as an ordinary global, so if the real spell turns out to be off the global, you will do slightly better than this advice. The build above spends its one point in it, behind three points of Shredding Attacks.

**Bear tank**: Bear Form tanking has its own guide, [Feral Bear Druid](/guides/druid/feral-bear), with its own simulated rotation, build and stat table.

## Stat priority

## Best in slot

For Bear tanking, look for armor, stamina, and defense skill first.

## Enchants and consumables

## Races

## Professions

Skinning and Leatherworking is the standard Feral pairing in Classic-era play: Skinning is free crafting material from every kill, and Leatherworking turns it into armor a melee Druid can actually wear.

## Leveling

Feral is generally considered the strongest leveling spec for Druid: Cat Form gives questing speed and burst damage, and Bear Form gives the durability to solo tougher pulls or tank early dungeons.
