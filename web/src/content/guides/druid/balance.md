---
title: Balance Druid in Forever
classSlug: druid
spec: balance
role: dps
build: 'FS1:1.60.1.69893:druid:night-elf:5532220005501001/0/5:'
recommendedRaces: [night-elf, tauren]
statPriority:
  [
    Spell power,
    Intellect,
    Critical strike,
    Hit,
    Spell haste,
    Spell penetration,
    'Nature damage and Arcane damage',
  ]
description: 'Talents, rotation, stats, and gear for Balance Druid in Forever, and what is confirmed versus projected from the beta.'
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
  - label: 'Forever Sixty rotation data, data/curated/apl/druid-balance.json'
    url: https://foreversixty.gg/sources
    kind: site
  - label: 'Forever Sixty raid loot data, data/curated/loot/forever-raid-phases.json'
    url: https://foreversixty.gg/sources
    kind: site
---

## Overview

Balance is Druid's ranged caster tree, keeping a Moonfire DoT on the target while filling every other global with Starfire, and offering the option to shapeshift into Moonkin Form for a party-wide critical strike buff. In a raid it's the class's only pure-damage caster role, and Moonkin Form gives it a support angle Balance-only builds elsewhere don't get. The beta only reaches level 30, so nothing here about level 60 talents, raiding, or best-in-slot gear has actually been played — it is a projection from the demo trees and 1.12 knowledge, not tested content.

## Talents and builds

Blizzard confirmed the tree keeps its seven rows and 51 points, with a fourth one-point talent added at the 16-point mark alongside the existing ones at 11, 21, and 31; separately, both Omen of Clarity and Nature's Grasp moved to the trainer instead of the tree, and Improved Mark of the Wild became a baseline passive, together freeing points every Druid used to spend elsewhere. Verified against this build's talent data, Balance's key picks in roughly the order you'd take them are:

- **Improved Wrath** — up to 0.5 seconds off Wrath's cast time and 50% off its mana cost at rank 5, a heavy early investment in the tree's other nuke.
- **Vengeance** — up to a 100% critical-strike-damage bonus at rank 5 on Arcane and Nature spells, the tree's biggest single multiplier.
- **Improved Starfire** — up to 0.5 seconds off Starfire's cast time and a 15% chance to stun the target at rank 5.
- **Nature's Grace** — a non-periodic spell crit speeds up your casting and cuts your next global cooldown by 10% for 3 seconds.
- **Moonfury** — up to 10% more Arcane and Nature damage at rank 5.
- **Moonkin Form** — the capstone: 360% more armor from items while shapeshifted, a doubled Omen of Clarity proc chance, and 3% more critical strike chance for party members within 45 yards.

A typical Balance build spends roughly 31 points in this tree to reach Moonkin Form at the bottom, with the remaining 5 points usually going into Restoration for Nature's Focus and mana talents — a common 1.12 hybrid pattern that likely still applies, though it isn't confirmed for Forever specifically. Open the planner at [/planner?class=druid](/planner?class=druid) to build this out.

## Rotation and priority

Keep Moonfire ticking on the target throughout the fight — it's the cheapest DoT available and, once it's up, everything else is filler. Starfire is that filler: cast it on repeat for the rest of the fight once Moonfire is running. Near the end of a fight, adjust for what will actually land: with less than 3.5 seconds remaining, switch to Wrath, since its 2-second cast still finishes where Starfire's 3.5-second cast would not; with less than 1.5 seconds left, only the instant Moonfire itself will land in time. Eclipse, if talented, rewards weaving in a Wrath cast periodically, since it shortens your next two Starfire casts — worth folding into the priority once you've taken it.

## Stat priority

In this site's own simulator weighting, in priority order:

1. **Spell power** — the reference stat; Moonfire and Starfire both scale off it directly.
2. **Intellect** — a larger mana pool for a rotation that casts continuously.
3. **Critical strike** — Vengeance and Moonfury both turn crits into more damage, and Nature's Grace rewards a crit with a faster next cast.
4. **Hit** — needed to stop missing casts against raid-level bosses.
5. **Spell haste** — shortens Starfire's long cast time, the rotation's main filler.
6. **Spell penetration** — only matters against targets with meaningful nature or arcane resistance.
7. **Nature damage and Arcane damage (school power)** — the narrowest stats, appearing on very few items, both relevant since Balance's damage spans both schools.

## Gear

Look for pieces that lead with spell power, then hit until capped, then crit, following the order above. Specific slot-by-slot picks can't be named honestly yet: the beta caps at level 30, so no one has gear-checked Balance past the early game, and raid loot doesn't exist yet either — nothing raids until 9 December, and even then this client's item table is missing most of the classic raid loot tables it's meant to carry (Onyxia's Lair's own list is currently empty, and Barrow Deeps and Hyjal Summit aren't mapped to items at all). This section will fill in with real picks once raid loot is itemized.

## Enchants and consumables

Target spell power and hit on enchants and consumables, matching the stat priority above. **Enchant Weapon - Spell Power** (verified in this build's enchant data, +30 spell damage) is the standard weapon enchant for a caster. **Greater Arcane Elixir** is a verified spell-damage consumable in this build's data that covers both of Balance's damage schools better than a single-school elixir would; **Flask of Supreme Power** is the raid-night alternative. Beyond naming those, keep it principle-level until raid-tier consumables are confirmed for Forever specifically.

## Races

Druid's playable races are Night Elf and Tauren on their respective factions, plus Skyborne on both sides — Druid is one of the few classes open to both the Alliance's High Order Skyborne and the Horde's Windshaper Skyborne. None of Blizzard's six confirmed new race and class pairs touches Druid, so this list is unchanged from 1.12.

For Alliance, **Night Elf** is the strongest pick: Elune's Light grants 10% critical strike chance for 15 seconds on a 3-minute cooldown, a direct damage cooldown for a caster spec, and Quickness's passive dodge and run speed add a little survivability while casting from range. For Horde, **Tauren** is the pick: Endurance grants 5% health and 1% hit chance, and that hit is a straightforward complement to a spec that needs to stay under the spell hit cap. Skyborne is available to both factions for Druid; its Elemental Insight racial adds 5% damage against Elementals, a narrow niche bonus, and Wind Blessed's 1% haste is a small, safe fit for a caster rotation, but neither is as strong a Balance pick as Night Elf's or Tauren's racials above.

## Professions

Balance benefits from the same caster-support professions as any spellcaster: Tailoring for crafted spellcaster gear and Enchanting to self-apply spell power enchants are the general Classic-era convention. This is community convention rather than anything Forever-specific, so treat it as inferred until this site has profession data from the beta.

## Leveling

Balance is a reasonable leveling spec on its own, with Moonfire and Starfire providing ranged damage, but Feral is generally considered the stronger overall leveling choice for Druid thanks to Cat Form's mobility and Bear Form's durability. That assessment is inferred from 1.12 knowledge, not tested — the beta only runs to level 30, so nothing about the leveling experience above that has been played in Forever's actual client.
