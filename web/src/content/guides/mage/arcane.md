---
title: Arcane Mage in Forever
classSlug: mage
spec: arcane
role: dps
build: 'FS1:1.60.1.70532:mage:gnome:1531005113100001531/032023/0532:'
raidBuild: 'FS1:1.60.1.70532:mage:gnome:1531005113100101531/03/05350003:'
recommendedRaces: [gnome, orc]
statPriority: [Spell power, Arcane power, Hit, Critical strike, Intellect, Spell haste, Spell penetration]
description: 'Talents, rotation, stats, and gear for Arcane Mage in Forever, and what is confirmed versus projected from the beta.'
updated: 2026-09-24
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
  - label: 'Forever Sixty rotation data, data/curated/apl/mage-arcane.json'
    url: https://foreversixty.gg/sources
    kind: site
  - label: 'Forever Sixty raid loot data, data/curated/loot/forever-raid-phases.json'
    url: https://foreversixty.gg/sources
    kind: site
---

## Overview

Arcane is the Mage burst spec: Arcane Power and Presence of Mind concentrate damage into short windows on the Arcane Blast loop.

### Strengths

- Arcane Power adds 30% spell damage for 15 seconds
- Presence of Mind makes the next cast instant
- Arcane Concentration gives free Clearcasting casts
- Missile Barrage halves Arcane Missiles' channel and makes it free

### Weaknesses

- Less sustained damage than Fire once the target's debuffs are stacked
- Arcane Power costs 30% more mana per cast
- Each Arcane Blast stack raises the next Blast's mana cost by 175%

### In a raid

- **Role:** Single-target caster DPS with burst on cooldown
- **Brings:** Arcane Intellect for the group; no raid-wide debuff

## Talents and builds

Blizzard confirmed the tree keeps its seven rows and 51 points, with a fourth one-point talent added at the 16-point mark alongside the existing ones at 11, 21, and 31.

- **Arcane Focus** — up to 5% additional hit chance with Arcane spells, worth taking early since hit is scarce while leveling.
- **Arcane Concentration** — up to a 10% chance for a damage spell to grant Clearcasting, making your next cast free; a large mana saver over a long fight.
- **Arcane Blast** (1 point) — a required pick, not an optional one, and missing from an earlier draft of this guide's talent section: it opens every cycle of the two-Blast loop described below in Rotation and priority, so without it there's no spell to stack the buff with in the first place.
- **Missile Barrage** — a proc off Arcane Blast, Fireball, Frostbolt, or Frostfire Bolt that halves Arcane Missiles' channel time and drops its mana cost to zero.
- **Presence of Mind** — turns your next sub-10-second cast instant, with no cast time of its own.
- **Arcane Mind** — up to 10% Intellect and a 100% bonus to Arcane critical strike damage at rank 5.
- **Arcane Power** — the capstone: 30% more spell damage for 15 seconds at the cost of 30% more mana per cast.

This build spends 31 points in Arcane, exactly the path to Arcane Power, and keeps every Arcane pick the sim measures as damage: **Arcane Impact** (3/3, 6% more Arcane critical strike chance) and **Arcane Instability** (3/3, 3% more spell damage and 3% more critical strike chance) stay maxed. The five points that earlier carried the Arcane tree past 31 (both ranks of **Arcane Subtlety**, both of **Magic Absorption**, one rank of **Wand Specialization**) came out because the sim measures none of them as damage for this rotation.

The other 20 points split 10 and 10. Fire keeps **Incineration** (3/3), **Improved Fireball** at 2/5, **Flame Throwing** (2/2) and **Impact** (3/3); **Improved Fire Blast** (the talent in Wake of Fire's old slot) and **Ignite** stay at 0, as before. Those picks are held on purpose: the sim does not model Flame Throwing or Impact, so it cannot show what they grant, and the guide keeps them rather than trade away a talent nobody has measured. Frost now takes 10 points instead of 5: **Improved Frostbolt** (5/5) and **Ice Shards** (2/5) are new, and **Elemental Precision** stays in at its 3 ranks (the client trimmed it from 5). Improved Frostbolt lifts Frostbolt's damage in the two-Blast loop and is the single largest gain in this spend; Ice Shards adds to Frostbolt's critical strike damage, and the two points Elemental Precision no longer holds went there. This build measures about +1.9% over the previous spend in our level-60 search run, and every talent the engine cannot model (Improved Channeling, Arcane Resilience, Arcane Geometry, Flame Throwing, Impact) is still in it. Open the planner at [/planner?class=mage](/planner?class=mage) to build this out.

**Raid build.** For a raid the search found about +1.5% over the leveling build above in the raid-ready run: it moves two points of Improved Fireball, the two Flame Throwing points (range on a school this rotation does not cast) and the three Impact points (a stun) into five of Ice Shards and three of **Frost Channeling**, which lift the Frostbolt in the two-Blast loop (Frost Channeling takes 15% off Frostbolt's mana cost). The client dropped Piercing Ice and cut Elemental Precision to 3 ranks, so the points they held went to Frost Channeling and to one rank of **Arcane Meditation**, which keeps part of your mana regeneration running while you cast. The rotation's talents, Improved Channeling, Arcane Resilience and Arcane Geometry all stay.

## Rotation and priority

Use Arcane Power the instant it's off cooldown — it has no cast time, so there's never a reason to hold it. Presence of Mind is the same: free to activate, so pop it whenever it's up and spend the instant cast on whichever spell benefits most in the moment. This site's simulator has Arcane Blast, its stacking buff and Missile Barrage wired up, including the part of Arcane Blast's tooltip that matters most: it costs 15% of your base mana, and every stack of its buff raises the next Blast's cost by 175%.

## Stat priority

The table above is this level range's own simulation at level 60, re-run by the nightly pipeline every time the build or its gear data changes — these numbers are never hand-entered. A stat shown as "not separable from zero" is not a verdict that it is worthless; the sim's measured error on it is too wide, at this level range's sample size, to tell its true value apart from zero.

## Best in slot

Look for pieces that lead with spell power, then hit until capped, then crit, following the order above.

## Enchants and consumables

Target spell power and hit on enchants and consumables, matching the stat priority above.

## Races

Mage's playable races are Human, Orc, Undead, Gnome, and Troll on their respective factions, plus Skyborne — but only the Alliance-side Skyborne, since Mage is Alliance-only for that race. Orc is one of Blizzard's six confirmed new race and class pairs, giving Horde a Mage option it didn't have in 1.12.

For Alliance, **Gnome** is the strongest pick: Eureka! cuts the mana cost of your next three abilities by 50% and adds 10% more damage on a 2-minute cooldown, a direct damage-and-mana cooldown that lines up with Arcane's burst-window playstyle, and Expansive Mind adds to the mana pool a spec built around Arcane Power's extra mana cost benefits from. For Horde, **Orc** is the pick: Blood Fury grants 10% spell power for 15 seconds on a 2-minute cooldown, a real damage cooldown that stacks with Arcane Power's own burst window, and it's the new pairing Blizzard specifically added for this class. Skyborne is available to Alliance Mages as a third option; its Read Ley Line racial restores all health and mana over 15 seconds, useful out of combat, but it doesn't add spell damage the way Gnome's Eureka! does, so it's a pick for the race and story rather than raw throughput.

## Professions

Tailoring and Enchanting is the standard caster pairing in Classic-era play: Tailoring's crafted spellcaster gear fills early slots, and Enchanting lets you apply your own weapon and gear enchants instead of paying for them.

## Leveling

Frost is the spec most players lean on for leveling, since Frostbolt's slow and Frost Nova's root make it safer to solo than Arcane's burst-focused kit. Arcane can still level fine, especially once Arcane Blast and Missile Barrage are actually castable.
