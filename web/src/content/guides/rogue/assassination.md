---
title: Assassination Rogue in Forever
classSlug: rogue
spec: assassination
role: dps
build: 'FS1:1.60.1.70291:rogue:night-elf:32502110551501001/302303/512:'
raidBuild: 'FS1:1.60.1.70291:rogue:night-elf:01532310421501/315303000015/002:'
recommendedRaces: [night-elf, troll]
statPriority: [Agility, Attack power, Strength, Critical strike, Hit, Melee haste]
description: 'Assassination Rogue overview, talent priority, rotation, stat weights, and race picks for Forever, with beta-versus-projection called out.'
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

Assassination is a poison-focused melee DPS spec that trades Combat's flat weapon damage for consistent poison uptime and a guaranteed-critical finisher window through Cold Blood. In a group it plays a single-target damage role rather than the cleave utility Combat or Subtlety can offer. The beta caps out at level 30, so nothing below about how this spec performs at level 60, in a full poison-uptime rotation, or against raid-tuned targets is confirmed play — it's read off the demo trees and 1.12 knowledge, not tested.

## Talents and builds

Reading the tree in priority order for a single-target build: **Malice** (flat crit chance with attacks and poisons) and **Lethality** (bonus critical strike damage on Sinister Strike, Gouge, Backstab, Mutilate, Ghostly Strike, and Hemorrhage) come first as damage multipliers that scale everything after them. **Seal Fate** converts combo-point-generating crits into bonus combo points, but this build no longer takes it (see below). **Cold Blood** is the spec's signature burst cooldown, guaranteeing a critical strike on the next attack. **Vile Poisons** and **Improved Poisons** raise both poison damage and application chance, which matters more here than in the other two trees. The two talent-granted abilities new to this tree in Forever, **Mutilate** (a dual-wield combo builder that hits harder against a target already carrying your poisons) and **Venom** (a finisher that raises how hard and how often your poisons land), round out the kit; Mutilate is live in this site's simulated rotation once talented, while Venom is not cast (see Rotation below).

The point allocation was re-searched against the simulated rotation on the current engine, and the build moved. It gained about +19% in the search run by re-spending points the engine measures at zero damage: **Seal Fate** (all five points), **Improved Sinister Strike** and three of the five **Lightning Reflexes** went to **Murder**, **Improved Slice and Dice**, **Relentless Strikes** (the largest single gain on the list), **Puncturing Wounds**, **Precision** and **Opportunity**. The trade is honest: Seal Fate's bonus combo points are not credited by the simulation, so its value here is unmeasured rather than proven nil, and a player who values it can take it back at a real cost in the measured damage. The search also wanted to drop **Venom**, because the engine scores it slightly negative, and the rotation has since stopped casting it, so the leveling build keeps the point only as a spare. Every other point of the old build stays, including **Cold Blood**, **Mutilate**, the poison talents, and the Improved Gouge, Remorseless Attacks, Camouflage and Master of Deception points the engine cannot yet measure.

Point allocation is still heavily weighted into Assassination: 32 points reach Venom at the bottom of the tree, with 11 in Combat (Improved Eviscerate, Lightning Reflexes, Puncturing Wounds, Precision) and 8 in Subtlety (Camouflage, Master of Deception, Opportunity). Open the planner at [/planner?class=rogue](/planner?class=rogue) to build this out.

**Raid build.** For a raid the search found about +9% over the leveling build above in the raid-ready run, and it drops Venom. The rotation no longer casts Venom (see Rotation below), so the point goes to Ruthlessness, Improved Slice and Dice, Improved Sinister Strike, Lightning Reflexes, Flawless Execution and five points of Dual Wield Specialization, paid for by three points of Vile Poisons and one of Lethality, which the simulator measures at little damage without Venom, and by the utility the engine cannot see and a raid rogue does not use: Improved Gouge (a stun), Camouflage and Master of Deception (stealth) and Remorseless Attacks (a kill-chain bonus for solo play). Cold Blood and Mutilate stay.

## Rotation and priority

Poisons come first: Deadly Poison on the main hand and Instant Poison on the off hand once Deadly is learned at 30. Mutilate's bonus against poisoned targets counts only a poison that lingers, so two Instants leave it on the table; this site's simulator carries that kit for Assassination, and it measured well ahead of two Instants. Mutilate is your builder, striking with both weapons for 2 combo points a cast, and you spend combo points from four rather than waiting for five: Mutilate gives points in twos, so a fifth point wastes one of them and a whole Mutilate on top. Slice and Dice goes first at four or more combo points whenever it has lapsed (opened when it is down, refreshed under 3 seconds remaining), for the attack-speed bonus; otherwise Eviscerate is the finisher whenever its 35 energy is there, and Cold Blood goes off right ahead of it for a guaranteed crit. Cold Blood costs no combo points and no global cooldown, so it simply waits for the Eviscerate. Venom, this tree's poison finisher, is not cast. The simulator models it to the client's numbers (30% more poison damage, 10 points more poison proc chance, 6 seconds plus 3 per combo point), but it multiplies Venom's bonus with Vile Poisons where the client adds the two, which flatters Venom, and even so the rotation measured about 1 to 2% better on the bare character without it and no different in the raid-ready run, where the Eviscerate and Slice and Dice lines take every set of combo points first. Without the Mutilate talent, Sinister Strike builds instead. Poisons do much of this spec's work, so keep them applied; Mutilate hits harder against a poisoned target.

## Stat priority

The table above is this band's own simulation at level 60, re-run by the nightly pipeline every time the build or its gear data changes — these numbers are never hand-entered. A stat shown as "not separable from zero" is not a verdict that it is worthless; the sim's measured error on it is too wide, at this band's sample size, to tell its true value apart from zero. These numbers come from this site's own level-60 simulator, not from beta play, which only reaches level 30.

## Gear

Look for agility and attack power first, then the merged hit and crit ratings the itemization change consolidated from their old melee/spell splits, then melee haste. A new stat that reduces the target's dodge and parry chance may show up on gear as itemization fills in, which would sit alongside hit rating in value for a dagger-and-poison spec that wants attacks to land. Specific pre-raid or raid-tier item picks can't be named with confidence yet: the beta caps at level 30, and this build's raid loot tables are themselves incomplete — most raid items the community's own sourcing expects are absent from the current client, and nothing raids in-game until the first tier opens on 9 December 2026. This section fills in once raid loot is itemized and the class is played past 30.

## Enchants and consumables

Target agility and attack power on weapon and glove enchants; **Enchant Gloves - Superior Agility** and **Enchant Weapon - Agility** both exist in this build's enchant data. For consumables, **Elixir of the Mongoose** (agility) is a verified option in this build's consumable list. Beyond naming those two, specific best-in-slot consumable stacking is not something this site can confirm yet at level 30. A raid shaman's **Windfury Totem** is an aura on you, not a weapon enchant, so it stacks with your poisons; this site's raid-preset numbers include it.

## Races

Assassination can be played by every race the Rogue class allows: Human, Dwarf, Night Elf, and Gnome on the Alliance; Orc, Undead, and Troll on the Horde; and, new in Forever, the Skyborne on either faction, since Skyborne is a new race rather than one of Blizzard's six new class pairings. For Alliance, **Night Elf** is the strongest pairing for a stealth-based melee class: Shadowmeld stealths on demand, works in combat, and is usable on a 2-minute cooldown, giving Assassination a reliable way to vanish out of a bad pull or reset an opener that Human or Dwarf can't match. For Horde, **Troll** pairs well through Berserking, a 10% attack and casting speed cooldown on a 3-minute timer that stacks on top of Cold Blood for a bigger burst window. Rogue is open to the Skyborne on both factions with no class restriction; their racials read as more utility- and survival-focused (a downward glide, a short full-resource restore) than damage-oriented, so they're a valid but not obviously stronger pick than Night Elf or Troll for this spec.

## Professions

Community convention for Rogue favors **Engineering** for its combat utility items (bombs, gadgets) alongside a gathering profession like Herbalism or Mining for self-funded leveling, since Rogue has no crafting profession that directly boosts poison or weapon damage the way casters benefit from tailored gear. This is general Classic-era community practice, not something confirmed for Forever specifically.

## Leveling

Combat is generally regarded as the stronger leveling spec for Rogue, since Adrenaline Rush and heavier weapon damage clear trash faster than Assassination's poison-uptime playstyle, which pays off more in sustained single-target fights. The beta caps at level 30, so this is carried over from 1.12 leveling knowledge rather than tested in Forever.
