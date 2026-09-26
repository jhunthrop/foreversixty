---
title: Feral Druid in Forever
classSlug: druid
spec: feral
role: dps
build: 'FS1:1.60.1.69893:druid:night-elf:0/5523232020032010001/055:'
recommendedRaces: [night-elf, tauren]
statPriority:
  [Attack power, 'Feral-specific attack power', Strength, Agility, Critical strike, Hit, Melee haste]
description: 'Talents, rotation, stats, and gear for Feral Druid in Forever, covering both Cat Form DPS and Bear Form tanking from the one tree.'
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
  - label: 'Forever Sixty rotation data, data/curated/apl/druid-feral.json'
    url: https://foreversixty.gg/sources
    kind: site
  - label: 'Forever Sixty raid loot data, data/curated/loot/forever-raid-phases.json'
    url: https://foreversixty.gg/sources
    kind: site
---

## Overview

Feral Combat is one tree covering two different jobs: Cat Form melee DPS, built around combo points spent on Rip and Ferocious Bite, and Bear Form tanking, built around threat generation and the armor and dodge talents in the same tree. This guide covers both roles, since Classic-era Druid design puts them in a single talent tree rather than splitting them the way other classes' roles are split. The beta only reaches level 30, so nothing here about level 60 talents, raiding, or best-in-slot gear — for either role — has actually been played; it is a projection from the demo trees and 1.12 knowledge, not tested content.

## Talents and builds

Blizzard confirmed the tree keeps its seven rows and 51 points, with a fourth one-point talent added at the 16-point mark alongside the existing ones at 11, 21, and 31. Two structural changes affect this tree specifically: Feral Charge splits into a Cat version, a gap-closing leap that dazes on landing, sitting next to the older Bear version built for interrupting; and Furor's cat-form energy refund now scales with how long you'd been out of form beforehand, rather than handing back a flat amount every time. Verified against this build's talent data, Feral's key picks in roughly the order you'd take them, useful to both roles, are:

- **Ferocity** — up to 5 less Rage or Energy on Maul, Mangle, Swipe, Claw, and Rake at rank 5.
- **Thick Hide** — extra base armor while shapeshifted, scaling with defense skill, essential for a Bear tank and useful padding for a Cat.
- **Savage Fury** — up to 10% more damage at rank 2 on Claw, Rake, Shred, Maul, and Swipe.
- **Predatory Strikes** — up to 150% of your level added to melee attack power in Cat or Bear Form at rank 3.
- **Primal Fury** — up to a 100% chance at rank 2 for bonus Rage on a Bear crit, and a 100% chance for a bonus combo point on a Cat crit, supporting both forms from one talent.
- **Berserk** — the capstone: Mangle hits up to 3 targets with no cooldown, and critical strike chance on combo-point generators rises 100%, for 15 seconds.

The point split differs by role. A cat-DPS build typically spends roughly 31 points in Feral reaching Berserk through the offensive column — Savage Fury, Predatory Strikes, Primal Fury — with the remaining 10 usually going into Restoration for Furor's rage-and-energy-on-shift and Naturalist's flat damage bonus, a common 1.12 pattern. A bear-tank build stays in Feral for more of its points, weighting the defensive column instead — Feral Instinct, Thick Hide, Natural Reaction — before reaching Berserk, and typically only dips into Restoration for Furor rather than Naturalist. Open the planner at [/planner?class=druid](/planner?class=druid) to build this out.

## Rotation and priority

**Cat DPS**: open with Tiger's Fury, since it costs nothing and refunds energy on its own 30-second cooldown, and keep it on cooldown throughout. Keep Rake ticking on the target — it's a strong bleed that costs no combo points, so upkeep on it shouldn't lapse. At 5 combo points, cast Rip if it isn't already running; once Rip is up, spend 5 combo points on Ferocious Bite instead of letting them cap out uselessly. Shred fills every other global, building combo points the rest of the time.

**Bear tank**: this site doesn't yet have a curated rotation file for tanking, so this is general Classic-era tanking knowledge rather than a sourced priority list. Open threat with Growl, then keep the new baseline Lacerate bleed refreshed on the target alongside Mangle or Maul as rage allows; Swipe covers multiple targets, and Demoralizing Roar reduces incoming melee damage for the group. Shredding Attacks' rage discount on Lacerate makes upkeep cheaper once talented.

## Stat priority

This site's own simulator weighting only covers the Cat DPS role; in priority order:

1. **Attack power** — the reference stat; every melee ability in Cat Form scales off it.
2. **Feral-specific attack power** — a stat that applies only in shapeshifted forms, stacking on top of normal attack power.
3. **Strength** — converts into attack power and scales with Predatory Strikes and Heart of the Wild.
4. **Agility** — adds crit chance and dodge, both useful in Cat Form.
5. **Critical strike** — feeds combo-point generation through Primal Fury and raw damage through Savage Fury's affected abilities.
6. **Hit** — needed to stop missing melee swings against raid-level bosses.
7. **Melee haste** — more swings and faster combo-point generation.

Bear-tank stat priority isn't in this site's data at all and has to be inferred separately from general Classic-era tanking knowledge: armor and stamina for survivability, defense skill to push toward the uncrittable threshold, and dodge as the main avoidance stat, roughly in that order.

## Gear

For Cat DPS, look for pieces that lead with attack power and agility, then crit and hit, following the priority above. For Bear tanking, look for armor, stamina, and defense skill first. Specific slot-by-slot picks can't be named honestly yet for either role: the beta caps at level 30, so no one has gear-checked Feral past the early game, and raid loot doesn't exist yet either — nothing raids until 9 December, and even then this client's item table is missing most of the classic raid loot tables it's meant to carry (Onyxia's Lair's own list is currently empty, and Barrow Deeps and Hyjal Summit aren't mapped to items at all). This section will fill in with real picks once raid loot is itemized.

## Enchants and consumables

For Cat DPS, target attack power and agility on enchants and consumables. **Enchant Weapon - Agility** (verified in this build's enchant data, +15 agility) is the standard weapon enchant, and **Elixir of Greater Agility**, also verified, is a solid consumable pick before a raid-tier agility elixir is confirmed. For Bear tanking, target stamina and armor; this site doesn't have a verified tank-oriented enchant list yet, so keep it principle-level until that data is confirmed for Forever.

## Races

Druid's playable races are Night Elf and Tauren on their respective factions, plus Skyborne on both sides — Druid is one of the few classes open to both the Alliance's High Order Skyborne and the Horde's Windshaper Skyborne. None of Blizzard's six confirmed new race and class pairs touches Druid, so this list is unchanged from 1.12.

For Alliance, **Night Elf** is the strongest pick for either role: Shadowmeld lets you stealth even in combat on a 2-minute cooldown, which meshes with Feral's own Prowl mechanic for repositioning or a surprise opener, and Quickness's passive dodge helps a Bear tank's avoidance and a Cat's survivability alike. For Horde, **Tauren** is the pick: War Stomp stuns up to 5 enemies within 8 yards on a 2-minute cooldown, a strong tool for a Bear tank picking up multiple adds, and Endurance's 5% health and 1% hit help both the tanking and DPS role. Skyborne is available to Druid on both factions, but its racials — Read Ley Line's health and mana regen, Elemental Insight's damage bonus against Elementals — lean toward caster support rather than melee combat, so it's a less targeted pick for Feral than Night Elf or Tauren.

## Professions

Skinning and Leatherworking is the standard Feral pairing in Classic-era play: Skinning is free crafting material from every kill, and Leatherworking turns it into armor a melee Druid can actually wear. This is general Classic-community convention rather than anything Forever-specific, so treat it as inferred until this site has profession data from the beta.

## Leveling

Feral is generally considered the strongest leveling spec for Druid: Cat Form gives questing speed and burst damage, and Bear Form gives the durability to solo tougher pulls or tank early dungeons. That assessment is inferred from 1.12 knowledge, not tested — the beta only runs to level 30, so nothing about the leveling experience above that has been played in Forever's actual client.
