---
title: Assassination Rogue in Forever
classSlug: rogue
spec: assassination
role: dps
build: 'FS1:1.60.1.70009:rogue:night-elf:3250001055050105/325201/51:'
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

Reading the tree in priority order for a single-target build: **Malice** (flat crit chance with attacks and poisons) and **Lethality** (bonus critical strike damage on Sinister Strike, Gouge, Backstab, Mutilate, Ghostly Strike, and Hemorrhage) come first as damage multipliers that scale everything after them. **Seal Fate** follows, since it converts combo-point-generating crits into bonus combo points and rewards the crit investment already made. **Vile Poisons** and **Improved Poisons** raise both poison damage and application chance, which matters more here than in the other two trees. **Mutilate**, the talent-granted dual-wield combo builder that hits harder against a target already carrying your poisons, rounds out the Assassination kit and is live in this site's simulated rotation once talented — see Rotation below.

This spend adds a single point each in **Relentless Strikes** and **Precision**. Relentless Strikes gives every finishing move a 20% chance per combo point spent to refund 25 Energy, which the sim measures as a real +8.9 DPS from that one point — more Energy back means more global cooldowns spent on Mutilate and finishers instead of waiting to recover. Precision, taken at 1/3, adds a flat 1% chance to hit, worth +10.3 DPS in the sim's model by cutting down on wasted swings.

To pay for them, this build drops **Cold Blood** and **Venom** entirely. Cold Blood is still a real ability — a guaranteed critical strike on the next attack — but it's an on-demand burst cooldown, and the current sustained single-target sim doesn't model cooldown-timing gains like that, so it measures as no DPS gained for the point spent, an honest gap in what the sim can credit rather than a verdict that the ability itself is bad. Venom is a genuine damage talent — the engine's own probe confirms it's doing real work on its own — but in this build's overall re-spend, directing that point into Relentless Strikes and Precision instead nets more total damage, so it's a calculated trade within the full search result, not a talent the sim dismisses as worthless. This build measures about +9.3% over the previous spend in our level-60 search run.

Point allocation is heavily weighted into Assassination, with the rest split as a handful of points in Combat for weapon-skill and survivability talents and a few in Subtlety for utility. The beta caps out at level 30, so none of this has actually been played past the early game — it's read off the demo trees, this site's own simulator, and 1.12 knowledge, not tested content. Open the planner at [/planner?class=rogue](/planner?class=rogue) to build this out.

## Rotation and priority

This site's own simulator now has both Mutilate and Venom implemented, and gates each on its talent: Cold Blood goes off first, right before a finisher, for the guaranteed crit — it costs no combo points, so it's checked ahead of the finishers regardless of which one ends up spending them. At five combo points, a talented rogue spends them on Venom for its poison damage and proc-chance buff; without the talent, Eviscerate is the finisher instead. Slice and Dice is only refreshed once five combo points are banked, and only when it actually needs it — fully down, or under 3 seconds remaining — so a full five-point finisher is never given up for the refresh. Mutilate, once talented, replaces Sinister Strike as the combo-point builder, striking with both weapons for 2 combo points a hit; Sinister Strike remains the fallback builder for a rogue without the Mutilate talent.

## Stat priority

The table above is this band's own simulation at level 60, re-run by the nightly pipeline every time the build or its gear data changes — these numbers are never hand-entered. A stat shown as "not separable from zero" is not a verdict that it is worthless; the sim's measured error on it is too wide, at this band's sample size, to tell its true value apart from zero. These numbers come from this site's own level-60 simulator, not from beta play, which only reaches level 30.

## Gear

Look for agility and attack power first, then the merged hit and crit ratings the itemization change consolidated from their old melee/spell splits, then melee haste. A new stat that reduces the target's dodge and parry chance may show up on gear as itemization fills in, which would sit alongside hit rating in value for a dagger-and-poison spec that wants attacks to land. Specific pre-raid or raid-tier item picks can't be named with confidence yet: the beta caps at level 30, and this build's raid loot tables are themselves incomplete — most raid items the community's own sourcing expects are absent from the current client, and nothing raids in-game until the first tier opens on 9 December 2026. This section fills in once raid loot is itemized and the class is played past 30.

## Enchants and consumables

Target agility and attack power on weapon and glove enchants; **Enchant Gloves - Superior Agility** and **Enchant Weapon - Agility** both exist in this build's enchant data. For consumables, **Elixir of the Mongoose** (agility) is a verified option in this build's consumable list. Beyond naming those two, specific best-in-slot consumable stacking is not something this site can confirm yet at level 30.

## Races

Assassination can be played by every race the Rogue class allows: Human, Dwarf, Night Elf, and Gnome on the Alliance; Orc, Undead, and Troll on the Horde; and, new in Forever, the Skyborne on either faction, since Skyborne is a new race rather than one of Blizzard's six new class pairings. For Alliance, **Night Elf** is the strongest pairing for a stealth-based melee class: Shadowmeld stealths on demand, works in combat, and is usable on a 2-minute cooldown, giving Assassination a reliable way to vanish out of a bad pull or reset an opener that Human or Dwarf can't match. For Horde, **Troll** pairs well through Berserking, a 10% attack and casting speed cooldown on a 3-minute timer that stacks on top of Cold Blood for a bigger burst window. Rogue is open to the Skyborne on both factions with no class restriction; their racials read as more utility- and survival-focused (a downward glide, a short full-resource restore) than damage-oriented, so they're a valid but not obviously stronger pick than Night Elf or Troll for this spec.

## Professions

Community convention for Rogue favors **Engineering** for its combat utility items (bombs, gadgets) alongside a gathering profession like Herbalism or Mining for self-funded leveling, since Rogue has no crafting profession that directly boosts poison or weapon damage the way casters benefit from tailored gear. This is general Classic-era community practice, not something confirmed for Forever specifically.

## Leveling

Combat is generally regarded as the stronger leveling spec for Rogue, since Adrenaline Rush and heavier weapon damage clear trash faster than Assassination's poison-uptime playstyle, which pays off more in sustained single-target fights. The beta caps at level 30, so this is carried over from 1.12 leveling knowledge rather than tested in Forever.
