---
title: Feral Bear Druid in Forever
classSlug: druid
spec: feral-bear
role: tank
build: 'FS1:1.60.1.70707:druid:night-elf:0/5523033202310212551/051:'
recommendedRaces: [night-elf, tauren]
statPriority:
  [Stamina, Armor, Defense, Dodge, Strength, Agility, Attack power, Hit, Critical strike, Expertise]
description: 'Talents, rotation, stats, and gear for the Feral Druid tank in Forever, Bear Form from the same tree as Cat Form.'
updated: 2026-10-07
confidence: inferred
sources:
  - label: 'Blizzard, Deep Dive panel recap'
    url: https://news.blizzard.com/en-us/article/24303313/world-of-warcraft-forever-deep-dive-panel-recap
    kind: blizzard
  - label: 'Talents Forever (demo transcription)'
    url: https://talentsforever.com/data.json
    kind: community
  - label: 'Forever Sixty rotation data, data/curated/apl/druid-feral-bear.json'
    url: https://foreversixty.gg/sources
    kind: site
  - label: 'Forever Sixty tank fight, data/curated/tank-encounter.json'
    url: https://foreversixty.gg/sources
    kind: site
---

## Overview

Feral Bear is the Bear Form tank: a large health pool, form-multiplied Armor and Dodge, and Maul, Swipe, Lacerate and Primal Bite to hold a boss.

### Strengths

- Bear Form multiplies Stamina and Armor, and Thick Hide adds more Armor
- Primal Bite is a high-threat attack on its own cooldown
- Lacerate stacks a bleed that Rend and Tear reads, and its hit grows with every stack

### Weaknesses

- No shield, Parry or Block: Dodge is its only avoidance
- Armor lowers physical hits by a share that shrinks as you stack it
- Maul and Swipe spend Rage, so threat follows Rage income

### In a raid

- **Role:** Main tank in Bear Form; for cat DPS see [Feral Druid](/guides/druid/feral)
- **Brings:** Leader of the Pack, a critical strike aura for the group

## Talents and builds

This build is the Feral tree almost whole (45 points) plus 6 in Restoration, and it is the build this site's tank simulator runs, taken from the client's own talent text.

- **Ferocity** — Maul, Swipe, Primal Bite, Claw and Rake cost less Rage or Energy.
- **Furor** — Rage on shifting into Bear Form (it moved from Restoration to Feral's top row in client 1.60.1.70707).
- **Feral Swiftness** and **Natural Reaction** — Dodge, and Natural Reaction also pays Rage whenever you dodge.
- **Thick Hide** — extra base Armor while shapeshifted, growing with Defense Skill beyond five times your level.
- **Feral Instinct**, **Savage Fury** and **Rend and Tear** — more damage from Swipe, from your melee abilities, and from every ability against a bleeding target, which Lacerate keeps on the boss.
- **Lacerate** (1 point) — a talent since client 1.60.1.70707, and the bear's bleed: it stacks five times, and each cast hits for 10% of weapon damage for every stack already up.
- **Shredding Attacks** — cheaper Lacerate.
- **Sharpened Claws**, **Predatory Strikes**, **Natural Instinct** and **Blood Frenzy** — crit, Attack Power, crit damage, and Rage on a critical strike.
- **Leader of the Pack** — 3% more critical strike chance for the party and raid, on attacks and spells alike.
- **Primal Bite** — the capstone since 1.60.1.70707: a high-threat bite on its own cooldown, the best single button the tree gives a bear.

The Restoration points are Heart of the Wild, more Stamina in Bear Form (it moved from Feral to Restoration's top row), and one point of Naturalist. Berserk, which took the cooldown off Primal Bite, left the tree in 1.60.1.70707; its point went to Lacerate. Left out: Brutal Impact, Feral Charge, Predatory Discipline and Shifting Power, which are for the cat or for a fight with movement.

Open the planner at [/planner?class=druid](/planner?class=druid) to build this out.

## Rotation and priority

This is the rotation the simulator runs, measured in the simulator's tank fight (a level 63 boss that swings at you every two seconds, healers assumed). Everything is in Bear Form, which Dire Bear Form replaces at level 40.

1. **Primal Bite** on its cooldown.
2. **Lacerate** built to five stacks and kept ticking: Rend and Tear reads it, and a cast on a full stack hits hardest.
3. **Maul**, queued on your next swing with spare Rage.
4. **Swipe** as the Rage dump.
5. **Barkskin** on its cooldown, a small cut in damage taken for a smaller cut in threat; **Frenzied Regeneration** as the heal when you are under half health.

Left out on purpose: Demoralizing Roar, because the simulated boss has no attack power for it to lower (a boss that has some changes this), and Enrage, which buys a little threat for a little more damage taken because it costs armor while it lasts. Growl, Challenging Roar, Bash and Feral Charge do nothing in a fight with one stationary enemy.

## Stat priority

A tank is scored on one number that rewards, in this order, effective health (how much of the boss's damage your health pool can take once armor and avoidance have done their work), then a risk index that punishes spiky incoming damage, then threat. Each stat in the table is how much one point of it moves that score, shown in the same convention as every other spec: against the stat the table is anchored to.

Why the stats land where they do is mechanical, and stays true when the numbers move. Stamina is health, and a bear's Stamina is multiplied in Bear Form. Armor is multiplied by the form and lowers every physical hit by a fraction that shrinks as you stack it. Dodge is the only avoidance a bear has, because it cannot Parry or Block, and Defense Skill both lowers the chance to be critically hit and feeds Thick Hide. Agility is Dodge, Armor and crit at once. Strength and Attack Power are threat. Expertise lowers the boss's chance to dodge you, which is threat, not survival, and no item carries it yet.

## Best in slot

## Enchants and consumables

## Races

## Professions

Skinning and Leatherworking, as for the cat: leather is the armor a druid wears, and Leatherworking turns Skinning's leather into it.

## Leveling

Bear Form is available from level 10, and a bear is a safe way to pull and survive as a leveling druid at the cost of killing speed. This site's tank simulation covers bands from level 20, and its low-level bands are the least reliable part of it; treat them as a direction, not a verdict.
