---
title: Subtlety Rogue in Forever
classSlug: rogue
spec: subtlety
role: dps
build: 'FS1:1.60.1.70532:rogue:night-elf:005323101014/0/5323220310013011031:'
recommendedRaces: [night-elf, troll]
statPriority: [Agility, Attack power, Strength, Critical strike, Hit, Melee haste]
description: 'Subtlety Rogue overview, talent priority, rotation, stat weights, and race picks for Forever, with beta-versus-projection called out.'
updated: 2026-09-24
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

Subtlety is the Rogue stealth-utility melee tree: Hemorrhage builds combo points cheaply and on the move, with Rupture and Ghostly Strike behind it.

### Strengths

- Cheap, mobile combo point generation
- Premeditation and Ambush give a strong stealth opener
- Hemorrhage debuff makes Rupture hit harder
- Vanish and Preparation allow a fresh opener mid-fight

### Weaknesses

- Least competitive of the three Rogue trees in raid damage
- Gives up Combat's weapon damage and Assassination's poison uptime
- Opener depends on stealth

### In a raid

- **Role:** Melee damage dealer with stealth utility
- **Brings:** Its own damage; no raid-wide buff or debuff

## Talents and builds

In rough priority order: **Opportunity** raises the damage of Backstab, Ambush, and Mutilate, which matters even for a Hemorrhage-based build since Ambush remains the spec's stealth opener. **Initiative** adds a chance at an extra combo point when opening with Cheap Shot or a similar ability, accelerating the path to a finisher. **Ghostly Strike** is a cheap, high-damage-relative-to-cost filler that also applies its own attack power debuff. **Premeditation** (1 point) is a required pick, not an optional one — missing from an earlier draft of this guide's talent section — since this build's rotation opens with it for two free combo points before the first Ambush. **Serrated Blades** adds armor penetration and boosts finisher damage, a flat multiplier once you're spending combo points regularly. **Hemorrhage**, the tree's signature ability, replaces Sinister Strike as the combo builder and needs its bleed reapplied once its charges run out. **Cutthroat**, new to the tree in Forever, occasionally lets a Backstab set up your following Ambush so it lands without needing Stealth first — a build-around for repeated openers mid-fight, though it isn't verified in this site's simulated rotation (see Rotation below). **Thousand Cuts**, the tree's capstone, is the bottom-row talent this build reaches.

The build changed in the latest talent search, by about +18% in the search run. The 15 points that sat in Combat (Improved Eviscerate, Improved Sinister Strike, Lightning Reflexes, Puncturing Wounds, Deflection, Precision) were re-spent into Assassination: **Malice** stays at five, joined by **Ruthlessness**, **Murder**, **Improved Slice and Dice**, **Relentless Strikes**, **Lethality**, **Cold Blood** and four points of **Improved Poisons**, with the remaining points going to **Setup** and **Dirty Tricks**. Two points of **Cutthroat** were given up to pay for it. Ghostly Strike, Premeditation, Hemorrhage and Thousand Cuts are all kept, because the rotation casts them. The dropped Combat points were credited at zero damage by the simulation, so the cost of the move is whatever those utility and defence talents did for you in play, not damage.

Point allocation now runs 31 points deep into Subtlety to reach Thousand Cuts at the bottom row, with the remaining 20 in Assassination. Open the planner at [/planner?class=rogue](/planner?class=rogue) to build this out.

**Raid build.** The raid-ready search found no build that beats this one beyond error while keeping what the rotation casts, the group buffs the raid already counts on and the threat and survival talents, so the raid build is the leveling build and one tree serves both.

## Rotation and priority

Open from stealth with Stealth, Premeditation and Ambush, then play from Hemorrhage. Hemorrhage is the builder, hitting for 145% weapon damage with a dagger and awarding a combo point, and its debuff makes the target take 15% more damage from your Rupture for 15 seconds, so keep it going. Rupture is the main finisher: cast it at four or more combo points whenever it is not already running. Ghostly Strike goes on cooldown, hitting for 180% weapon damage with a dagger for a combo point.

## Stat priority

The table above is this level range's own simulation at level 60, re-run by the nightly pipeline every time the build or its gear data changes — these numbers are never hand-entered. A stat shown as "not separable from zero" is not a verdict that it is worthless; the sim's measured error on it is too wide, at this level range's sample size, to tell its true value apart from zero.

## Best in slot

Look for agility and attack power first, then the merged hit and crit ratings, then melee haste. Since Hemorrhage and Rupture are both driven by weapon damage and attack power, those outweigh haste more than in a pure swing-speed spec. A new stat reducing target dodge and parry chance may show up on gear as itemization fills in.

## Enchants and consumables

A raid shaman's **Windfury Totem** is an aura on you, not a weapon enchant, so it stacks with your poisons; this site's raid-preset numbers include it.

## Races

Subtlety can be played by every race the Rogue class allows on both factions, plus the Skyborne on either faction since Rogue's Skyborne availability has no faction restriction. For Alliance, **Night Elf** is the clearest fit for a stealth-oriented spec: Shadowmeld works in combat on a 2-minute cooldown, giving Subtlety a second stealth tool beyond its own Vanish and Preparation for repositioning into a fresh opener. For Horde, **Troll** pairs well through Berserking's 10% attack and casting speed on a 3-minute cooldown, which lines up with a burst window after an opener lands. Rogue is open to the Skyborne on both factions with no class restriction; their read racials skew toward mobility and resource restoration rather than the stealth-and-burst identity Subtlety leans on, making them a workable but not clearly stronger pick than Night Elf or Troll here.

## Professions

Community convention favors **Engineering** for combat utility alongside a gathering profession such as Herbalism or Mining, matching general Rogue practice across all three trees rather than anything Subtlety-specific.

## Leveling

Combat is the generally recommended leveling spec for Rogue over Subtlety, since Adrenaline Rush and consistent weapon damage clear trash faster than Subtlety's stealth-and-reposition playstyle, which pays off more in single-target or PvP scenarios.
