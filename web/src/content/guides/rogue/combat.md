---
title: Combat Rogue in Forever
classSlug: rogue
spec: combat
role: dps
build: 'FS1:1.60.1.70009:rogue:night-elf:32531/32531300000515201/51:'
recommendedRaces: [night-elf, troll]
statPriority: [Agility, Attack power, Strength, Critical strike, Hit, Melee haste]
description: 'Combat Rogue overview, talent priority, rotation, stat weights, and race picks for Forever, with beta-versus-projection called out.'
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

Combat is the straightforward, weapon-damage melee DPS spec for Rogue: a Sinister Strike into Eviscerate loop with a burst window from Adrenaline Rush doubling energy regeneration. It brings the most reliable sustained single-target damage of the three Rogue trees, and unlike Assassination it doesn't depend on poison uptime to hit its numbers. The beta caps at level 30, so anything about how Combat scales at level 60 or performs against raid-tuned encounters is a projection from the demo trees and 1.12 knowledge, not confirmed play.

## Talents and builds

In rough priority order: **Improved Sinister Strike** lowers the energy cost of the spec's main combo builder early, which compounds for the rest of a leveling or raiding career. **Dual Wield Specialization** adds flat off-hand damage, a straightforward multiplier once dual-wielding. **Hack and Slash** now treats axes the same way it always treated swords, folded into one shared weapon-type group — this is the talent behind Rogues gaining access to one-handed axes at all in Forever. **Weapon Expertise**, reworked to reduce the chance enemies dodge or parry attacks rather than raising weapon skill, is a flat damage-uptime gain. **Blade Flurry** adds attack speed and a second melee target, useful whenever there's more than one target to hit. **Adrenaline Rush**, the tree's capstone, doubles energy regeneration for 15 seconds and is the spec's signature burst cooldown.

Point allocation runs deep into Combat to reach Adrenaline Rush at the bottom of the tree, roughly 31 points, with the remainder split as a few points in Assassination for crit and poison utility and a few in Subtlety for survivability talents. Open the planner at [/planner?class=rogue](/planner?class=rogue) to build this out.

## Rotation and priority

Keep Slice and Dice up at all times: it is the first thing you cast, at any number of combo points, and from then on you refresh it at five combo points once under 3 seconds remain, so its attack-speed bonus never lapses. Spend combo points on Eviscerate at four or more; in this site's simulator that beat waiting for a fifth. Use Adrenaline Rush on cooldown as the burst window (Blade Flurry goes up alongside it whenever it is ready), and build with Backstab when a dagger is in your main hand: it hits for 150% weapon damage against Sinister Strike's 100%, and Puncturing Wounds adds crit and a chance at a second combo point to it. In this site's simulator the Backstab version beat the Sinister Strike one by about 10%. With a sword or mace in the main hand Backstab cannot be cast, and Sinister Strike builds instead. There is no opener-specific sequencing beyond that in the current simulated priority, and the tree's talent changes (cheaper Sinister Strike, stronger off-hand, Blade Flurry cleave) affect the numbers behind each swing rather than the order abilities are used in.

## Stat priority

The table above is this band's own simulation at level 60, re-run by the nightly pipeline every time the build or its gear data changes — these numbers are never hand-entered. A stat shown as "not separable from zero" is not a verdict that it is worthless; the sim's measured error on it is too wide, at this band's sample size, to tell its true value apart from zero. These numbers come from this site's own level-60 simulator, not from beta play, which only reaches level 30.

## Gear

Prioritize agility and attack power first, then the merged hit and crit ratings the itemization change consolidated, then melee haste. Dual Wield Specialization makes off-hand weapon damage worth checking specifically when comparing two weapon options. A new stat reducing the target's dodge and parry chance may appear on gear as itemization fills in, which is worth tracking alongside hit rating for a spec that depends on landing every Backstab or Sinister Strike. Specific pre-raid or raid-tier item recommendations aren't something this site can name with confidence yet: the beta caps at level 30 and this build's raid loot tables are still missing most items the community's sourcing expects, with nothing raiding in-game until the first tier opens on 9 December 2026.

## Enchants and consumables

Weapon and glove enchants should target agility and attack power; **Enchant Gloves - Superior Agility** and **Enchant Weapon - Agility** are both verified entries in this build's enchant data. **Elixir of the Mongoose** (agility) is a verified consumable option. Beyond those, this site isn't naming a full consumable stack until the class is played past level 30.

## Races

Combat can be played by every race the Rogue class allows on both factions, plus the Skyborne on either faction since Rogue's Skyborne availability isn't restricted by faction. For Alliance, **Night Elf** is a strong pick even for a non-stealth-focused spec like Combat: Shadowmeld works in combat on a 2-minute cooldown, giving a melee-range escape or a way to drop threat and reposition that Human and Dwarf don't have. For Horde, **Troll** is the strongest fit: Berserking's 10% attack and casting speed on a 3-minute cooldown stacks directly with Adrenaline Rush for a larger combined burst window than any other Horde racial offers Combat. Rogue is open to the Skyborne on both factions with no class restriction, but their read racials (a downward glide, a short full-resource restore) lean utility rather than the sustained melee output Combat is built around.

## Professions

Community convention favors **Engineering** for combat utility items alongside a gathering profession such as Mining or Herbalism for self-funded leveling and consumables, since Combat's damage doesn't scale from a crafting profession the way a caster's does. This is general Classic-era practice, not confirmed for Forever specifically.

## Leveling

Combat is generally the strongest of the three Rogue trees for leveling: Adrenaline Rush and consistent weapon damage clear trash faster than Assassination's poison-uptime approach or Subtlety's stealth-oriented kit. The beta caps at level 30, so this is carried over from 1.12 leveling knowledge, not tested in Forever.
