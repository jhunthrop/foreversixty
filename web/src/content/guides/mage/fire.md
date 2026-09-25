---
title: Fire Mage in Forever
classSlug: mage
spec: fire
role: dps
description: 'Talents, rotation, stats, and gear for Fire Mage in Forever, and what is confirmed versus projected from the beta.'
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
  - label: 'Forever Sixty rotation data, data/curated/apl/mage-fire.json'
    url: https://foreversixty.gg/sources
    kind: site
  - label: 'Forever Sixty raid loot data, data/curated/loot/forever-raid-phases.json'
    url: https://foreversixty.gg/sources
    kind: site
---

## Overview

Fire is Mage's sustained single-target burner: it builds a stacking vulnerability debuff on its target through Scorch, then cashes it in for guaranteed critical strikes through Combustion, chaining Fireball and Fire Blast around that loop for the rest of the fight. In a raid it's a steady, high-uptime caster rather than a spike-and-wait spec. The beta only reaches level 30, so nothing here about level 60 talents, raiding, or best-in-slot gear has actually been played — it is a projection from the demo trees and 1.12 knowledge, not tested content.

## Talents and builds

Blizzard confirmed the tree keeps its seven rows and 51 points, with a fourth one-point talent added at the 16-point mark alongside the existing ones at 11, 21, and 31. Verified against this build's talent data, the Fire tree's key picks in roughly the order you'd take them are:

- **Improved Fireball** — up to 0.5 seconds off Fireball's and Frostfire Bolt's cast time at rank 5, a direct throughput gain taken early.
- **Ignite** — Fire critical strikes burn the target for a further 40% of the spell's damage over 4 seconds at rank 5, turning every crit into extra sustained damage.
- **Improved Scorch** — up to a guaranteed chance for Scorch to stack a 3%-per-stack Fire vulnerability debuff, up to 5 stacks, on the target.
- **Critical Mass** — up to 6% more Fire critical strike chance at rank 3, feeding both Ignite and Combustion.
- **Fire Power** — a flat 10% more Fire damage at rank 5, the tree's biggest raw damage talent.
- **Combustion** — the capstone: each Fire spell hit adds 10% Fire crit chance, lasting until you land four non-periodic Fire crits.

A typical Fire build spends roughly 31 points in this tree to reach Combustion at the bottom, with the remaining 20 points usually going into Arcane for Arcane Concentration's Clearcasting and Arcane Mind's crit damage bonus — a common 1.12 hybrid pattern that likely still applies, though it isn't confirmed for Forever specifically. Open the planner at [/planner?class=mage](/planner?class=mage) to build this out.

## Rotation and priority

Open by stacking Improved Scorch's vulnerability debuff with a few Scorch casts before your main damage window, then hold Combustion until the debuff is fully stacked so every Fire spell hit during Combustion lands as a guaranteed crit. Keep Scorch refreshed through the fight so the debuff never falls off, weaving it in whenever it's about to expire. Fire Blast slots in on cooldown as a free extra hit that doesn't compete with your Scorch or Fireball casts for a global cooldown. Fireball is the primary filler for every global that isn't spent on Scorch upkeep, Fire Blast, or Combustion. This site's simulator currently only models this loop when Improved Scorch is actually talented — without it, the debuff-stacking logic has nothing to track and the rotation collapses to plain Fireball filler, so take Improved Scorch if you want the loop above to matter.

## Stat priority

In this site's own simulator weighting, in priority order:

1. **Spell power** — the reference stat; every Fire cast scales off it directly.
2. **Intellect** — a larger mana pool for a rotation that never really stops casting.
3. **Critical strike** — Ignite and Combustion both turn directly into more damage the higher your crit chance runs.
4. **Hit** — needed to stop missing casts against raid-level bosses; not covered by any Fire talent.
5. **Spell haste** — shortens Fireball and Scorch cast times, more casts per minute.
6. **Spell penetration** — only matters against targets with meaningful fire resistance.
7. **Fire damage (school power)** — the narrowest stat, appearing on very few items, but it stacks directly with Fire Power's percentage bonus.

## Gear

Look for pieces that lead with spell power, then hit until capped, then crit, following the order above. Specific slot-by-slot picks can't be named honestly yet: the beta caps at level 30, so no one has gear-checked Fire past the early game, and raid loot doesn't exist yet either — nothing raids until 9 December, and even then this client's item table is missing most of the classic raid loot tables it's meant to carry (Onyxia's Lair's own list is currently empty, and Barrow Deeps and Hyjal Summit aren't mapped to items at all). This section will fill in with real picks once raid loot is itemized.

## Enchants and consumables

Target spell power and hit on enchants and consumables, matching the stat priority above. **Enchant Weapon - Spell Power** (verified in this build's enchant data, +30 spell damage) is the standard weapon enchant for a caster. **Elixir of Fire Power** is a verified school-specific damage elixir in this build's data and is the natural pick for Fire over a generic spell-power elixir when the two conflict; **Flask of Supreme Power** covers the general spell-damage slot on a raid night. Beyond naming those, keep it principle-level until raid-tier consumables are confirmed for Forever specifically.

## Races

Mage's playable races are Human, Orc, Undead, Gnome, and Troll on their respective factions, plus Skyborne — but only the Alliance-side Skyborne, since Mage is Alliance-only for that race. Orc is one of Blizzard's six confirmed new race and class pairs, giving Horde a Mage option it didn't have in 1.12.

For Alliance, **Gnome** is the strongest pick: Eureka! cuts the mana cost of your next three abilities by 50% and adds 10% more damage on a 2-minute cooldown, which lines up well with weaving a burst of Fireballs around a Combustion window, and Expansive Mind's mana bonus helps sustain a spec that rarely stops casting. For Horde, **Orc** is the pick: Blood Fury grants 10% spell power for 15 seconds on a 2-minute cooldown, a straightforward damage cooldown to pair with Combustion, and it's the new pairing Blizzard specifically added for this class. Skyborne is available to Alliance Mages as a third option; its Read Ley Line racial restores all health and mana over 15 seconds, useful out of combat, but it doesn't add spell damage the way Gnome's Eureka! does, so it's a pick for the race and story rather than raw throughput.

## Professions

Tailoring and Enchanting is the standard caster pairing in Classic-era play: Tailoring's crafted spellcaster gear fills early slots, and Enchanting lets you apply your own weapon and gear enchants instead of paying for them. This is general Classic-community convention rather than anything Forever-specific, so treat it as inferred until this site has profession data from the beta.

## Leveling

Frost is the spec most players lean on for leveling, since Frostbolt's slow and Frost Nova's root make it safer to solo than Fire's more contact-heavy kit. Fire can still level well once Improved Scorch and Combustion are both online, but that's inferred from 1.12 knowledge, not tested — the beta only runs to level 30, so nothing about the leveling experience above that has been played in Forever's actual client.
