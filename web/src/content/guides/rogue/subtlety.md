---
title: Subtlety Rogue in Forever
classSlug: rogue
spec: subtlety
role: dps
build: 'FS1:1.60.1.70009:rogue:night-elf:005323101014/0/5323220310013011031:'
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

Subtlety trades Combat's raw weapon damage and Assassination's poison uptime for cheap, mobile combo generation and stealth-oriented utility, using Hemorrhage as a bleed-based builder instead of Sinister Strike. It's historically the least raid-damage-competitive of the three Rogue trees in the 1.12 era, and nothing in the Forever demo data suggests that's changed at the tree-design level. The beta caps at level 30, so how Subtlety's damage actually compares to Assassination or Combat at level 60 or in a raid is a projection from demo-tree reading and 1.12 knowledge, not confirmed play.

## Talents and builds

In rough priority order: **Opportunity** raises the damage of Backstab, Ambush, and Mutilate, which matters even for a Hemorrhage-based build since Ambush remains the spec's stealth opener. **Initiative** adds a chance at an extra combo point when opening with Cheap Shot or a similar ability, accelerating the path to a finisher. **Ghostly Strike** is a cheap, high-damage-relative-to-cost filler that also applies its own attack power debuff. **Premeditation** (1 point) is a required pick, not an optional one — missing from an earlier draft of this guide's talent section — since this build's rotation opens with it for two free combo points before the first Ambush. **Serrated Blades** adds armor penetration and boosts finisher damage, a flat multiplier once you're spending combo points regularly. **Hemorrhage**, the tree's signature ability, replaces Sinister Strike as the combo builder and needs its bleed reapplied once its charges run out. **Cutthroat**, new to the tree in Forever, occasionally lets a Backstab set up your following Ambush so it lands without needing Stealth first — a build-around for repeated openers mid-fight, though it isn't verified in this site's simulated rotation (see Rotation below). **Thousand Cuts**, the tree's capstone, is the bottom-row talent this build reaches.

The build changed in the latest talent search, by about +18% in the search run. The 15 points that sat in Combat (Improved Eviscerate, Improved Sinister Strike, Lightning Reflexes, Puncturing Wounds, Deflection, Precision) were re-spent into Assassination: **Malice** stays at five, joined by **Ruthlessness**, **Murder**, **Improved Slice and Dice**, **Relentless Strikes**, **Lethality**, **Cold Blood** and four points of **Improved Poisons**, with the remaining points going to **Setup** and **Dirty Tricks**. Two points of **Cutthroat** were given up to pay for it. Ghostly Strike, Premeditation, Hemorrhage and Thousand Cuts are all kept, because the rotation casts them. The dropped Combat points were credited at zero damage by the simulation, so the cost of the move is whatever those utility and defence talents did for you in play, not damage.

Point allocation now runs 31 points deep into Subtlety to reach Thousand Cuts at the bottom row, with the remaining 20 in Assassination. Open the planner at [/planner?class=rogue](/planner?class=rogue) to build this out.

## Rotation and priority

Open from stealth with Stealth, Premeditation and Ambush, then play from Hemorrhage. Hemorrhage is the builder, hitting for 145% weapon damage with a dagger and awarding a combo point, and its debuff makes the target take 15% more damage from your Rupture for 15 seconds, so keep it going. Rupture is the main finisher: cast it at four or more combo points whenever it is not already running. Ghostly Strike goes on cooldown, hitting for 180% weapon damage with a dagger for a combo point. While Rupture is running, spend the combo points that build behind it: Eviscerate at three or more, and Slice and Dice when it is down or under 3 seconds from ending, so nothing idles at the cap. In this site's simulator, dropping Eviscerate and Slice and Dice entirely and putting every point into Rupture measured a few percent higher still; the priority above keeps all three finishers in play rather than adopting that. Cutthroat's proc-based free Ambush isn't reflected in that priority list yet, so treat any Ambush-heavy opener sequencing as this site's own projection of the tree's intent rather than a confirmed rotation. Poisons are your weapon imbues: this site's simulator carries Deadly Poison on the main hand and Instant Poison on the off hand once Deadly is learned at 30, which measured well ahead of two Instants; at level 60 the mirrored pair measured about 1% higher still.

## Stat priority

The table above is this band's own simulation at level 60, re-run by the nightly pipeline every time the build or its gear data changes — these numbers are never hand-entered. A stat shown as "not separable from zero" is not a verdict that it is worthless; the sim's measured error on it is too wide, at this band's sample size, to tell its true value apart from zero. These numbers come from this site's own level-60 simulator, not from beta play, which only reaches level 30.

## Gear

Look for agility and attack power first, then the merged hit and crit ratings, then melee haste. Since Hemorrhage and Rupture are both driven by weapon damage and attack power, those outweigh haste more than in a pure swing-speed spec. A new stat reducing target dodge and parry chance may show up on gear as itemization fills in. Specific pre-raid or raid-tier item picks aren't something this site can name with confidence yet: the beta caps at level 30 and this build's raid loot tables are still missing most items the community's sourcing expects, with nothing raiding in-game until the first tier opens on 9 December 2026.

## Enchants and consumables

Weapon and glove enchants should target agility and attack power; **Enchant Gloves - Superior Agility** and **Enchant Weapon - Agility** are both verified in this build's enchant data. **Elixir of the Mongoose** (agility) is a verified consumable option. A fuller consumable stack isn't something this site is naming until the class is played past level 30. A raid shaman's **Windfury Totem** is an aura on you, not a weapon enchant, so it stacks with your poisons; this site's raid-preset numbers include it.

## Races

Subtlety can be played by every race the Rogue class allows on both factions, plus the Skyborne on either faction since Rogue's Skyborne availability has no faction restriction. For Alliance, **Night Elf** is the clearest fit for a stealth-oriented spec: Shadowmeld works in combat on a 2-minute cooldown, giving Subtlety a second stealth tool beyond its own Vanish and Preparation for repositioning into a fresh opener. For Horde, **Troll** pairs well through Berserking's 10% attack and casting speed on a 3-minute cooldown, which lines up with a burst window after an opener lands. Rogue is open to the Skyborne on both factions with no class restriction; their read racials skew toward mobility and resource restoration rather than the stealth-and-burst identity Subtlety leans on, making them a workable but not clearly stronger pick than Night Elf or Troll here.

## Professions

Community convention favors **Engineering** for combat utility alongside a gathering profession such as Herbalism or Mining, matching general Rogue practice across all three trees rather than anything Subtlety-specific. This is Classic-era community convention, not confirmed for Forever.

## Leveling

Combat is the generally recommended leveling spec for Rogue over Subtlety, since Adrenaline Rush and consistent weapon damage clear trash faster than Subtlety's stealth-and-reposition playstyle, which pays off more in single-target or PvP scenarios. The beta caps at level 30, so this is carried over from 1.12 leveling knowledge rather than tested in Forever.
