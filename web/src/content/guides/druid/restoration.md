---
title: Restoration Druid in Forever
classSlug: druid
spec: restoration
role: healer
build: 'FS1:1.60.1.69893:druid:night-elf:503/0/5552005003113001:'
recommendedRaces: [night-elf, tauren]
statPriority: [Healing power, Spell power, Spirit, MP5, Intellect, Critical strike]
description: 'Talents, rotation, stats, and gear for Restoration Druid in Forever, and what is confirmed versus projected from the beta.'
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
  - label: 'Forever Sixty raid loot data, data/curated/loot/forever-raid-phases.json'
    url: https://foreversixty.gg/sources
    kind: site
---

## Overview

Restoration is Druid's healer tree, keeping Rejuvenation rolling on multiple targets while spending bigger casts like Regrowth and Healing Touch on whoever's taking real damage, with Swiftmend as a way to turn an existing heal-over-time into an instant burst. Its mana efficiency talents let it out-sustain other healers over a long fight rather than out-burst them. The beta only reaches level 30, so nothing here about level 60 talents, raiding, or best-in-slot gear has actually been played — it is a projection from the demo trees and 1.12 knowledge, not tested content.

## Talents and builds

Blizzard confirmed the tree keeps its seven rows and 51 points, with a fourth one-point talent added at the 16-point mark alongside the existing ones at 11, 21, and 31. Verified against this build's talent data, Restoration's key picks in roughly the order you'd take them are:

- **Nature's Focus** — up to a 70% chance at rank 5 to avoid a cast being pushed back by damage, essential for healing through incoming raid damage.
- **Gift of Nature** — a flat 10% more effective healing at rank 5 across every healing spell, the tree's single biggest throughput talent.
- **Improved Rejuvenation** — up to 15% more effective Rejuvenation at rank 3.
- **Swiftmend** — an instant heal that consumes an active Rejuvenation or Regrowth on the target for a burst equal to the rest of that heal-over-time's value.
- **Nature's Swiftness** — makes your next Nature spell instant on activation, a saved emergency cooldown.
- **Wild Growth** — the capstone: an area heal on the target and nearby party members, weighted to land more up front and taper off over its duration.

A typical Restoration build spends roughly 31 points in this tree, with the remaining 8 usually going into Balance for Moonglow's mana discount — a common 1.12 hybrid pattern that likely still applies, though it isn't confirmed for Forever specifically. Open the planner at [/planner?class=druid](/planner?class=druid) to build this out.

## Rotation and priority

This site doesn't have a curated rotation file for Restoration, since it's a healer spec rather than a damage one, so this section is written from general 1.12/Classic-era healing priority plus what's confirmed to have changed for Forever, not from a sourced simulator build. Keep Rejuvenation rolling on the tank or whoever is taking steady damage, refreshing it before it falls off rather than letting it lapse. Reach for Regrowth when a target needs a faster, larger heal than Rejuvenation alone provides, and hold Swiftmend for moments when you need an instant heal right now and already have a Rejuvenation or Regrowth running to consume. Wild Growth and Tranquility cover spikes of raid-wide damage. Save Nature's Swiftness for genuine emergencies, since it only affects a single cast. One change worth noting for a leveling or soloing Restoration Druid: starting at level 10, gear's healing bonus also adds roughly a third of that number as bonus damage, so a Resto-specced Druid hits harder while questing than in 1.12 without needing to swap specs.

## Stat priority

In this site's own simulator weighting, in priority order:

1. **Healing power** — the reference stat; it scales every heal directly.
2. **Spell power** — contributes to healing output alongside healing power on modern itemization.
3. **Spirit** — mana regeneration for a spec built around sustaining a fight rather than burning cooldowns.
4. **MP5** — flat mana regeneration, a direct complement to Spirit.
5. **Intellect** — a larger mana pool.
6. **Critical strike** — a smaller factor than for a DPS spec, but Improved Regrowth and Nature's Grace both make it worth some weight.

## Gear

Look for pieces that lead with healing power, then spirit and mp5 for mana sustain, following the order above. Specific slot-by-slot picks can't be named honestly yet: the beta caps at level 30, so no one has gear-checked Restoration past the early game, and raid loot doesn't exist yet either — nothing raids until 9 December, and even then this client's item table is missing most of the classic raid loot tables it's meant to carry (Onyxia's Lair's own list is currently empty, and Barrow Deeps and Hyjal Summit aren't mapped to items at all). This section will fill in with real picks once raid loot is itemized.

## Enchants and consumables

Target healing power and spirit on enchants and consumables, matching the stat priority above. **Enchant Bracer - Healing Power** and **Enchant Gloves - Healing Power**, both verified in this build's enchant data, are standard picks for a healer. **Elixir of Greater Spirit**, verified in this build's data, supports mana sustain over a long fight; **Flask of Distilled Wisdom** is the raid-night mana-pool alternative. Beyond naming those, keep it principle-level until raid-tier consumables are confirmed for Forever specifically.

## Races

Druid's playable races are Night Elf and Tauren on their respective factions, plus Skyborne on both sides — Druid is one of the few classes open to both the Alliance's High Order Skyborne and the Horde's Windshaper Skyborne. None of Blizzard's six confirmed new race and class pairs touches Druid, so this list is unchanged from 1.12.

For Alliance, **Night Elf** is the strongest pick: Quickness's passive dodge and speed help a healer stay alive while kiting or repositioning, and Shadowmeld's in-combat stealth is a real emergency tool when a healer draws unwanted attention. For Horde, **Tauren** is the pick: Endurance's 5% health gives a healer more of a buffer, and War Stomp's group stun is a useful panic button if adds close in on a healer standing at range. Skyborne is available to Druid on both factions; Read Ley Line's 100% health and mana regeneration over 15 seconds is a genuinely strong out-of-combat sustain tool for a mana-focused healer, making it a reasonable third option rather than purely a flavor pick.

## Professions

Herbalism and Alchemy is the general Classic-era pairing for a Restoration healer, giving access to self-crafted mana and healing consumables. This is community convention rather than anything Forever-specific, so treat it as inferred until this site has profession data from the beta.

## Leveling

Feral is the stronger leveling choice for Druid overall, and Restoration is generally considered the slowest of the three specs to solo with, since its kit is built around sustaining others rather than dealing damage. That assessment is inferred from 1.12 knowledge, not tested — the beta only runs to level 30, so nothing about the leveling experience above that has been played in Forever's actual client.
