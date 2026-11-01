---
title: Assassination Rogue in Forever
classSlug: rogue
spec: assassination
role: dps
build: 'FS1:1.60.1.70707:rogue:night-elf:32502110551501001/302303/512:'
raidBuild: 'FS1:1.60.1.70707:rogue:night-elf:0053231032150105/324303000012/002:'
recommendedRaces: [night-elf, troll]
statPriority: [Agility, Attack power, Strength, Critical strike, Hit, Melee haste]
description: 'Assassination Rogue overview, talent priority, rotation, stat weights, and race picks for Forever, with beta-versus-projection called out.'
updated: 2026-10-31
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

Assassination is the Rogue poison-based melee damage tree: constant poison uptime, Mutilate as the builder, and Cold Blood for a guaranteed-critical finisher.

### Strengths

- Mutilate strikes with a dagger in each hand for 2 combo points a cast
- Cold Blood guarantees a critical strike on the next attack
- Vile Poisons and Improved Poisons raise poison damage and application chance
- Relentless Strikes is the largest single damage gain in the build

### Weaknesses

- Single-target damage role, with less cleave than Combat or Subtlety
- Gives up Combat's flat weapon damage
- Mutilate's bonus needs a lingering poison, so poison upkeep is required

### In a raid

- **Role:** Single-target melee damage dealer
- **Brings:** Its own damage; no raid-wide buff or debuff

## Talents and builds

Reading the tree in priority order for a single-target build: **Malice** (flat crit chance with attacks and poisons) and **Lethality** (bonus critical strike damage on Sinister Strike, Gouge, Backstab, Mutilate, Ghostly Strike, and Hemorrhage) come first as damage multipliers that scale everything after them. **Seal Fate** turns a combo-point builder's crit into one more combo point. Mutilate procs it once a cast when either hand crits, so the raid build takes all five ranks: about +3% in our level-60 search run. For a boss kill it pays with Remorseless Attacks (a bonus after each kill, which helps on trash, not on a single boss), a point each from Lethality and Lightning Reflexes and three from Dual Wield Specialization, and it adds a second point of Improved Sinister Strike. The leveling build below still skips Seal Fate. **Cold Blood** is the spec's signature burst cooldown, guaranteeing a critical strike on the next attack. **Vile Poisons** and **Improved Poisons** raise both poison damage and application chance, which matters more here than in the other two trees. The two talent-granted abilities new to this tree in Forever, **Mutilate** (a combo builder that strikes with a dagger in each hand and hits harder against a target already carrying your poisons) and **Venom** (a finisher that raises how hard and how often your poisons land), round out the kit; Mutilate is live in this site's simulated rotation once talented, while Venom is not cast (see Rotation below).

The leveling build's points are still heavily weighted into Assassination: 32 points reach Venom at the bottom of the tree, with 11 in Combat (Improved Eviscerate, Lightning Reflexes, Puncturing Wounds, Precision) and 8 in Subtlety (Camouflage, Master of Deception, Opportunity). Open the planner at [/planner?class=rogue](/planner?class=rogue) to build this out.

## Rotation and priority

Poisons come first: Deadly Poison on the main hand and Instant Poison on the off hand once Deadly is learned at 30. Mutilate's bonus against poisoned targets counts only a poison that lingers, so two Instants leave it on the table; this site's simulator carries that kit for Assassination, and it measured well ahead of two Instants. Mutilate is your builder, striking with both daggers for 2 combo points a cast (it needs a dagger in each hand), and you spend combo points from four rather than waiting for five: Mutilate gives points in twos, so a fifth point wastes one of them and a whole Mutilate on top. Slice and Dice goes first at four or more combo points whenever it has lapsed (opened when it is down, refreshed under 3 seconds remaining), for the attack-speed bonus; otherwise Eviscerate is the finisher whenever its 35 energy is there, and Cold Blood goes off right ahead of it for a guaranteed crit. Cold Blood costs no combo points and no global cooldown, so it simply waits for the Eviscerate. Venom, this tree's poison finisher, is not cast. The simulator models it to the client's numbers (30% more poison damage, 10 points more poison proc chance, 6 seconds plus 3 per combo point), but it multiplies Venom's bonus with Vile Poisons where the client adds the two, which flatters Venom, and even so the rotation measured about 1 to 2% better on the bare character without it and no different in the raid-ready run, where the Eviscerate and Slice and Dice lines take every set of combo points first. Thistle Tea goes in whenever energy is near empty: 100 Energy on a 5 minute cooldown. Without the Mutilate talent, Sinister Strike builds instead. Poisons do much of this spec's work, so keep them applied; Mutilate hits harder against a poisoned target.

## Stat priority

## Best in slot

A new stat that reduces the target's dodge and parry chance may show up on gear as itemization fills in, which would sit alongside hit rating in value for a dagger-and-poison spec that wants attacks to land.

## Enchants and consumables

A raid shaman's **Windfury Totem** is an aura on you, not a weapon enchant, so it stacks with your poisons; this site's raid-preset numbers include it.

## Races

## Professions

Community convention for Rogue favors **Engineering** for its combat utility items (bombs, gadgets) alongside a gathering profession like Herbalism or Mining for self-funded leveling, since Rogue has no crafting profession that directly boosts poison or weapon damage the way casters benefit from tailored gear.

## Leveling

Combat is generally regarded as the stronger leveling spec for Rogue, since Adrenaline Rush and heavier weapon damage clear trash faster than Assassination's poison-uptime playstyle, which pays off more in sustained single-target fights.
