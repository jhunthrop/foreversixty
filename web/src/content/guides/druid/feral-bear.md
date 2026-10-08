---
title: Feral Bear Druid in Forever
classSlug: druid
spec: feral-bear
role: tank
build: 'FS1:1.60.1.70009:druid:night-elf:0/55230332020132012551/051:'
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

The Feral Combat tree serves two jobs, and this is the second: Bear Form tanking. The cat's DPS guide is [Feral Druid](/guides/druid/feral); the two share a talent tree and a trainer and little else. A bear has no shield, no Parry and no Block. It stays alive on a large health pool, Armor multiplied by its form and by Thick Hide, and Dodge, and it holds a boss with Maul, Swipe, Lacerate and Primal Bite. The beta only reaches level 30, so everything here about level 60 is this site's own simulation of the client's tables, not play.

## Talents and builds

This build is the Feral tree almost whole (45 points) plus 6 in Restoration, and it is the build this site's tank simulator runs, taken from the client's own talent text.

- **Ferocity** — Maul, Swipe, Primal Bite, Claw and Rake cost less Rage or Energy.
- **Heart of the Wild** — in Bear Form, more Stamina, which is health.
- **Feral Swiftness** and **Natural Reaction** — Dodge, and Natural Reaction also pays Rage whenever you dodge.
- **Thick Hide** — extra base Armor while shapeshifted, growing with Defense Skill beyond five times your level.
- **Feral Instinct**, **Savage Fury** and **Rend and Tear** — more damage from Swipe, from your melee abilities, and from every ability against a bleeding target, which Lacerate keeps on the boss.
- **Shredding Attacks** — cheaper Lacerate.
- **Sharpened Claws**, **Predatory Strikes**, **Predatory Instincts** and **Blood Frenzy** — crit, Attack Power, crit damage, and Rage on a critical strike.
- **Primal Bite** — a high-threat bite on its own cooldown, the best single button the tree gives a bear.
- **Leader of the Pack** and **Berserk** — the raid's crit, and Berserk takes the cooldown off Primal Bite and lets it strike three targets.

The Restoration points are Furor, for Rage on shifting into Bear Form, and the rest of the budget. Left out: Brutal Impact, Feral Charge, Shifting Power and Improved Shifting Power, which are for the cat or for a fight with movement.

Open the planner at [/planner?class=druid](/planner?class=druid) to build this out.

## Rotation and priority

This is the rotation the simulator runs, measured in the simulator's tank fight (a level 63 boss that swings at you every two seconds, healers assumed). Everything is in Bear Form, which Dire Bear Form replaces at level 40.

1. **Primal Bite** on its cooldown.
2. **Lacerate** kept ticking: it stacks, and Rend and Tear reads it.
3. **Maul**, queued on your next swing with spare Rage.
4. **Swipe** as the Rage dump.
5. **Barkskin** on its cooldown, a small cut in damage taken for a smaller cut in threat; **Frenzied Regeneration** as the heal when you are under half health.

Berserk on its cooldown adds threat. Left out on purpose: Demoralizing Roar, because the simulated boss has no attack power for it to lower (a boss that has some changes this), and Enrage, which buys a little threat for a little more damage taken because it costs armor while it lasts. Growl, Challenging Roar, Bash and Feral Charge do nothing in a fight with one stationary enemy.

## Stat priority

The table above is this band's own tank simulation at level 60, re-run by the nightly pipeline whenever the build or its gear data changes; these numbers are never hand-entered. A tank is scored on one number that rewards, in this order, effective health (how much of the boss's damage your health pool can take once armor and avoidance have done their work), then a risk index that punishes spiky incoming damage, then threat. Each stat in the table is how much one point of it moves that score, shown in the same convention as every other spec: against the stat the table is anchored to.

Why the stats land where they do is mechanical, and stays true when the numbers move. Stamina is health, and a bear's Stamina is multiplied in Bear Form. Armor is multiplied by the form and lowers every physical hit by a fraction that shrinks as you stack it. Dodge is the only avoidance a bear has, because it cannot Parry or Block, and Defense Skill both lowers the chance to be critically hit and feeds Thick Hide. Agility is Dodge, Armor and crit at once. Strength and Attack Power are threat. Expertise lowers the boss's chance to dodge you, which is threat, not survival, and no item carries it yet.

A stat shown as "not separable from zero" is not a verdict that it is worthless; the simulation's own error on it is too wide, at this band's sample size, to tell its value apart from zero.

## Gear

Follow the table above: health, then Armor and Dodge, with the threat stats behind them. This site's ranker builds the level 60 set for you in the BiS tab, scored on the same tank fight. Specific raid picks cannot be named honestly yet: nothing raids until 9 December, and this build's raid loot tables are only partly itemized for Forever.

## Enchants and consumables

Target stamina and the defensive stats on enchants where they are available. For consumables the BiS tab's Raid-ready set shows the bear's: a health flask, stamina food and drink, strength and attack power for threat, and a Stoneshield potion.

## Races

Night Elf and Tauren, as for the cat. **Night Elf** has Quickness, a passive Dodge bonus that a bear, with Dodge as its only avoidance, values more than the cat does. **Tauren** has Endurance, more health, which is a straight gain in effective health for a tank, and War Stomp for a pack of adds.

## Professions

Skinning and Leatherworking, as for the cat: leather is the armor a druid wears, and Leatherworking turns Skinning's leather into it. This is general Classic convention rather than something confirmed for Forever.

## Leveling

Bear Form is available from level 10, and a bear is a safe way to pull and survive as a leveling druid at the cost of killing speed. This site's tank simulation covers bands from level 20, and its low-level bands are the least reliable part of it; treat them as a direction, not a verdict.
