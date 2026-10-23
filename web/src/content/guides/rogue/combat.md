---
title: Combat Rogue in Forever
classSlug: rogue
spec: combat
role: dps
build: 'FS1:1.60.1.70532:rogue:night-elf:32531/32530300001515201/51:'
raidBuild: 'FS1:1.60.1.70532:rogue:night-elf:005320105/31530300001515231/002:'
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

Combat is the Rogue weapon-damage melee tree: Sinister Strike into Eviscerate, with Adrenaline Rush doubling energy regeneration during burst windows.

### Strengths

- Most reliable sustained single-target damage of the three Rogue trees
- Does not depend on poison uptime
- Adrenaline Rush doubles energy regeneration in burst windows
- Blade Flurry adds attack speed and a second melee target

### Weaknesses

- Burst is tied to Adrenaline Rush and Blade Flurry cooldowns
- Best damage wants a dagger in the main hand for Backstab
- Little stealth or stun utility in the raid build

### In a raid

- **Role:** Melee damage dealer, sustained single target
- **Brings:** Its own damage; no raid-wide buff or debuff

## Talents and builds

In rough priority order: **Improved Sinister Strike** lowers the energy cost of the spec's main combo builder early, which compounds for the rest of a leveling or raiding career. **Dual Wield Specialization** adds flat off-hand damage, a straightforward multiplier once dual-wielding. **Hack and Slash** now treats axes the same way it always treated swords, folded into one shared weapon-type group — this is the talent behind Rogues gaining access to one-handed axes at all in Forever. **Weapon Expertise**, reworked to reduce the chance enemies dodge or parry attacks rather than raising weapon skill, is a flat damage-uptime gain. **Blade Flurry** adds attack speed and a second melee target, useful whenever there's more than one target to hit. **Adrenaline Rush**, the tree's capstone, doubles energy regeneration for 15 seconds and is the spec's signature burst cooldown.

One point moved in the latest talent search: **Deflection** became **Flawless Execution**, about +2.5% in the search run. Deflection is a defensive talent the simulation credits at zero damage, and Flawless Execution adds damage on top of the same fight with nothing the rotation casts given up. A larger rebuild toward Assassination talents scored much higher in the same search but would drop talents the engine cannot yet measure (Improved Gouge, Remorseless Attacks, Camouflage, Master of Deception), so this guide keeps the conservative build until the engine can price them.

Point allocation runs deep into Combat to reach Adrenaline Rush at the bottom of the tree, roughly 31 points, with the remainder split as a few points in Assassination for crit and poison utility and a few in Subtlety for utility talents. Open the planner at [/planner?class=rogue](/planner?class=rogue) to build this out.

**Raid build.** For a raid the search found about +16% over the leveling build above in the raid-ready run, the largest of any spec. It moves the points the simulator measures at no damage or cannot see and a raid rogue does not use (Improved Gouge, a stun; Camouflage and Master of Deception, stealth; Remorseless Attacks, a kill-chain bonus for solo play) and one Improved Sinister Strike into five Lethality, three Aggression, two Opportunity, Murder and Relentless Strikes. Adrenaline Rush and Blade Flurry stay.

## Rotation and priority

Open with Slice and Dice at three combo points, then spend combo points on Eviscerate at five: in this site's simulator, keeping Slice and Dice up the whole fight measured worse than putting those combo points into Eviscerate, so after the opener the simulated priority does not spend more points on it. In this site's simulator the Backstab version beat the Sinister Strike one by about 10%. With a sword or mace in the main hand Backstab cannot be cast, and Sinister Strike builds instead.

## Stat priority

The table above is this level range's own simulation at level 60, re-run by the nightly pipeline every time the build or its gear data changes — these numbers are never hand-entered. A stat shown as "not separable from zero" is not a verdict that it is worthless; the sim's measured error on it is too wide, at this level range's sample size, to tell its true value apart from zero.

## Best in slot

Prioritize agility and attack power first, then the merged hit and crit ratings the itemization change consolidated, then melee haste. Dual Wield Specialization makes off-hand weapon damage worth checking specifically when comparing two weapon options. A new stat reducing the target's dodge and parry chance may appear on gear as itemization fills in, which is worth tracking alongside hit rating for a spec that depends on landing every Backstab or Sinister Strike.

## Enchants and consumables

A raid shaman's **Windfury Totem** is an aura on you, not a weapon enchant, so it stacks with your poisons; this site's raid-preset numbers include it.

## Races

Combat can be played by every race the Rogue class allows on both factions, plus the Skyborne on either faction since Rogue's Skyborne availability isn't restricted by faction. For Alliance, **Night Elf** is a strong pick even for a non-stealth-focused spec like Combat: Shadowmeld works in combat on a 2-minute cooldown, giving a melee-range escape or a way to drop threat and reposition that Human and Dwarf don't have. For Horde, **Troll** is the strongest fit: Berserking's 10% attack and casting speed on a 3-minute cooldown stacks directly with Adrenaline Rush for a larger combined burst window than any other Horde racial offers Combat. Rogue is open to the Skyborne on both factions with no class restriction, but their read racials (a downward glide, a short full-resource restore) lean utility rather than the sustained melee output Combat is built around.

## Professions

Community convention favors **Engineering** for combat utility items alongside a gathering profession such as Mining or Herbalism for self-funded leveling and consumables, since Combat's damage doesn't scale from a crafting profession the way a caster's does.

## Leveling

Combat is generally the strongest of the three Rogue trees for leveling: Adrenaline Rush and consistent weapon damage clear trash faster than Assassination's poison-uptime approach or Subtlety's stealth-oriented kit.
