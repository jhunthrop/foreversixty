---
title: Frost Mage in Forever
classSlug: mage
spec: frost
role: dps
build: 'FS1:1.60.1.69893:mage:gnome:0/2305/2555100300000301051:'
recommendedRaces: [gnome, troll]
statPriority: [Spell power, Intellect, Critical strike, Hit, Spell haste, Spell penetration, 'Frost damage']
description: 'Talents, rotation, stats, and gear for Frost Mage in Forever, and what is confirmed versus projected from the beta.'
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
  - label: 'Forever Sixty rotation data, data/curated/apl/mage-frost.json'
    url: https://foreversixty.gg/sources
    kind: site
  - label: 'Forever Sixty raid loot data, data/curated/loot/forever-raid-phases.json'
    url: https://foreversixty.gg/sources
    kind: site
---

## Overview

Frost is Mage's control and survivability spec: Frostbolt's slow keeps enemies at range, Frost Nova and Ice Block cover emergencies, and the tree's damage talents lean on the Shatter interaction — bonus crit chance against a frozen target — rather than a burst cooldown. In a raid it trades some of Fire's sustained output for more personal safety and utility; solo and leveling, that safety is the whole appeal. The beta only reaches level 30, so nothing here about level 60 talents, raiding, or best-in-slot gear has actually been played — it is a projection from the demo trees and 1.12 knowledge, not tested content.

## Talents and builds

Blizzard confirmed the tree keeps its seven rows and 51 points, with a fourth one-point talent added at the 16-point mark alongside the existing ones at 11, 21, and 31. Verified against this build's talent data, the Frost tree's key picks in roughly the order you'd take them are:

- **Improved Frostbolt** — up to 0.5 seconds off Frostbolt's cast time at rank 5, the tree's most direct throughput gain.
- **Ice Shards** — up to 100% more Frost critical strike damage at rank 5, the biggest single damage multiplier in the tree.
- **Piercing Ice** — a flat 6% more Frost damage at rank 3.
- **Shatter** — up to 50% more critical strike chance against a frozen target at rank 3, the talent the whole "freeze then burst" playstyle is built around.
- **Winter's Chill** — a chance for Frost hits to stack a debuff that raises Ice Lance's and Frostbolt's crit chance against that target, up to 5 stacks at rank 5.
- **Ice Barrier** — the capstone: an instant shield that also stops your casts from being interrupted or delayed while it holds.

A typical Frost build spends roughly 31 points in this tree to reach Ice Barrier at the bottom, with the remaining 10 points usually going into Fire for Ignite — a common 1.12 hybrid pattern that likely still applies, though it isn't confirmed for Forever specifically. Open the planner at [/planner?class=mage](/planner?class=mage) to build this out.

## Rotation and priority

Right now this site's simulator models Frost as Frostbolt on repeat, with nothing else in the loop — that's an accurate reflection of a no-talent Frost rotation, but not the full Shatter-combo playstyle Frost is built around once its talents are online. The intended loop layers Frost Nova or another freeze effect onto a target, then fires Ice Lance while it's frozen for a guaranteed, heavily boosted critical strike through Shatter, before returning to Frostbolt as the default filler between freezes. Winter's Chill stacks passively off your Frost hits and should be allowed to build before you commit to the Ice Lance burst. Treat the freeze-into-Ice-Lance combo as the spec's real rotation once talents and cooldowns are available to test, and plain Frostbolt spam as the floor.

## Stat priority

In this site's own simulator weighting, in priority order:

1. **Spell power** — the reference stat; every Frost cast scales off it directly.
2. **Intellect** — a larger mana pool for a spec that casts Frostbolt continuously.
3. **Critical strike** — Ice Shards turns every crit into far more damage, and Shatter multiplies your effective crit chance against frozen targets.
4. **Hit** — needed to stop missing casts against raid-level bosses; only partly covered by Elemental Precision.
5. **Spell haste** — shortens Frostbolt's cast time, more casts per minute.
6. **Spell penetration** — only matters against targets with meaningful frost resistance.
7. **Frost damage (school power)** — the narrowest stat, appearing on very few items, but it stacks directly with Piercing Ice's percentage bonus.

## Gear

Look for pieces that lead with spell power, then hit until capped, then crit, following the order above. Specific slot-by-slot picks can't be named honestly yet: the beta caps at level 30, so no one has gear-checked Frost past the early game, and raid loot doesn't exist yet either — nothing raids until 9 December, and even then this client's item table is missing most of the classic raid loot tables it's meant to carry (Onyxia's Lair's own list is currently empty, and Barrow Deeps and Hyjal Summit aren't mapped to items at all). This section will fill in with real picks once raid loot is itemized.

## Enchants and consumables

Target spell power and hit on enchants and consumables, matching the stat priority above. **Enchant Weapon - Spell Power** (verified in this build's enchant data, +30 spell damage) is the standard weapon enchant for a caster. **Elixir of Frost Power** is a verified school-specific damage elixir in this build's data and is the natural pick for Frost over a generic spell-power elixir when the two conflict; **Flask of Supreme Power** covers the general spell-damage slot on a raid night. Beyond naming those, keep it principle-level until raid-tier consumables are confirmed for Forever specifically.

## Races

Mage's playable races are Human, Orc, Undead, Gnome, and Troll on their respective factions, plus Skyborne — but only the Alliance-side Skyborne, since Mage is Alliance-only for that race. Orc is one of Blizzard's six confirmed new race and class pairs, giving Horde a Mage option it didn't have in 1.12.

For Alliance, **Gnome** is the strongest pick: Escape Artist breaks and grants brief immunity to movement impairment on a 2-minute cooldown, which is a direct complement to a spec that already relies on kiting and control, and Expansive Mind's mana bonus supports a rotation that runs on continuous Frostbolt casts. For Horde, **Troll** is the pick: Berserking grants 10% casting and attack speed on a 3-minute cooldown, a straightforward way to squeeze more Frostbolts and Ice Lances into a burst window; it's also the new pairing Blizzard specifically added for this class, though for Warlock rather than Mage — Orc remains Mage's new Horde option, and its Blood Fury spell power cooldown is a reasonable alternative pick if Troll's timing doesn't line up with your raid's needs. Skyborne is available to Alliance Mages as a third option; its Read Ley Line racial restores all health and mana over 15 seconds, useful out of combat, but it doesn't add spell damage or control the way Gnome's kit does, so it's a pick for the race and story rather than raw throughput.

## Professions

Tailoring and Enchanting is the standard caster pairing in Classic-era play: Tailoring's crafted spellcaster gear fills early slots, and Enchanting lets you apply your own weapon and gear enchants instead of paying for them. This is general Classic-community convention rather than anything Forever-specific, so treat it as inferred until this site has profession data from the beta.

## Leveling

Frost is the strong leveling choice for Mage generally, and that holds for itself: Frostbolt's slow and Frost Nova's root make solo play safer than either Arcane's burst-dependent kit or Fire's closer-range playstyle. That assessment is inferred from 1.12 knowledge, not tested — the beta only runs to level 30, so nothing about the leveling experience above that has been played in Forever's actual client.
