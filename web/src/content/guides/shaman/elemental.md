---
title: Elemental Shaman in Forever
classSlug: shaman
spec: elemental
role: dps
build: 'FS1:1.60.1.70532:shaman:dwarf:551033130000135/03/503352:'
raidBuild: 'FS1:1.60.1.70532:shaman:dwarf:551033130000135/03/053352:'
recommendedRaces: [dwarf, orc]
statPriority: [Intellect, Spell power, Hit, Nature power, Critical strike, Spell haste, Spell penetration]
description: 'Elemental Shaman overview, talent priority, rotation, stat weights, and race picks for Forever, with beta-versus-projection called out.'
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

Elemental is Shaman's ranged spellcasting spec, dealing damage with Lightning Bolt, Chain Lightning, Flame Shock and Lava Burst, backed by fire totems.

### Strengths

- Forever's baseline cast-time cut to Lightning Bolt and Chain Lightning
- Lava Burst hits harder when Flame Shock is already on the target
- Chain Lightning hits every target for close to single-target Lightning Bolt damage
- Searing Totem for one target and Magma Totem for several add steady damage

### Weaknesses

- Slower leveling than Enhancement, because its damage depends on cast times
- Hard casts leave it idle while moving
- Raid build spends points in Restoration to reach spell hit

### In a raid

- **Role:** Ranged spell damage
- **Brings:** Party totems alongside its own damage; no raid-wide buff or debuff

## Talents and builds

In rough priority order: **Concussion** raises the damage of Lightning Bolt, Chain Lightning, and Earth Shock, a broad multiplier across the entire single-target and AoE kit. **Call of Flame**, reworked in Forever, now boosts Lava Burst and Fire Nova alongside Flame Shock within a single talent instead of touching only the Fire Totems it used to buff, effectively unifying the tree's fire-damage kit under one pick. **Elemental Fury** raises critical strike damage on the totems and elements the spec relies on, moved earlier in the tree with a lower first-rank value than 1.12's single large-rank version. **Lightning Overload** gives Lightning Bolt and Chain Lightning a chance to cast a second, weaker bolt for free, a straight damage-per-cast increase. **Elemental Alacrity** cuts the cast time of Lightning Bolt, Chain Lightning, and Lava Burst further on top of the baseline speed-up Forever already gave those spells — maxed, it brings deep Elemental's effective cast speed roughly back to where 1.12 Elemental topped out, since the starting point is already faster. **Lava Burst**, the tree's new capstone nuke, hits harder against a target already carrying Flame Shock.

Point allocation runs deep into Elemental (30 points: Convection, Concussion, Elemental Warding, Call of Flame, Elemental Devastation, Elemental Focus, Elemental Alacrity, Lightning Overload, Elemental Fury and Lava Burst, with Convection, Concussion, Lightning Overload and Elemental Fury maxed), with 3 points in Enhancement for Thundering Strikes and the remaining 18 in Restoration for Improved Healing Wave and Totemic Focus (both maxed), Mindfulness, Natural Grace, and Improved Reincarnation. This replaces the 32/1/18 build. The talent search re-run on the current engine, once with the bare character and once under the Phase 1 raid context, found one change that keeps every talent the engine does not yet model, stays Elemental-majority and drops nothing the loop casts: Reverberation's two points go to Thundering Strikes. It came out about 1.5% ahead in both runs, beyond the run's error. Reverberation shortens the shock cooldown and the loop casts no shock, so the points were doing nothing. The Lava Burst point is in the build: the current client removed Call of Thunder from the tree, and the point it held goes to Lava Burst, which the loop below casts on every cooldown. Every talent the engine does not yet model (Elemental Warding, Mindfulness, Natural Grace, Improved Healing Wave, Improved Reincarnation) stays as before, since the search cannot judge it. Open the planner at [/planner?class=shaman](/planner?class=shaman) to build this out.

**Raid build.** For a raid the search found about +4% over the previous raid build in the raid-ready run. It takes **Tidal Focus** at 5/5, a Restoration talent whose healing-cost cut does nothing here but whose five percent spell hit the engine counts in full, and it pays for it with the Thundering Strikes and Ancestral Knowledge points the previous raid build spent. Against the leveling build above the raid build moves the five Improved Healing Wave points, a self-heal for solo play, into Tidal Focus. Elemental Warding, Natural Grace (less threat), Mindfulness and Improved Reincarnation stay.

## Rotation and priority

The loop this site's simulator plays: against a single target, drop Searing Totem for the higher fire damage; against two or more targets, drop Magma Totem instead so everyone standing in it takes damage. Against two or more targets, Chain Lightning is used since it hits every target for close to single-target Lightning Bolt damage. Lightning Bolt fills the rest of the single-target rotation, on every free global. Lava Burst is the first line when you have the talent: it is cast on every cooldown, ahead of the totems and the cooldowns, with no Flame Shock condition.

## Stat priority

The table above is this level range's own simulation at level 60, re-run by the nightly pipeline every time the build or its gear data changes — these numbers are never hand-entered. A stat shown as "not separable from zero" is not a verdict that it is worthless; the sim's measured error on it is too wide, at this level range's sample size, to tell its true value apart from zero.

## Best in slot

Prioritize spell power and intellect first, then crit and hit, then spell haste, spell penetration, and Nature-specific power.

## Enchants and consumables

## Races

Elemental can be played by Orc and Troll on the Horde, Tauren on the Horde, Dwarf — new in Forever — on the Alliance, and the Windshaper Skyborne on the Horde only. Dwarf is, for now, the only Alliance option for Shaman at all. For Alliance, **Dwarf** is the only pick, but it's a reasonable one for Elemental: Big Game Hunter adds damage against Beasts, useful for leveling and questing damage even if it doesn't apply in most raid encounters. For Horde, **Orc** is the strongest pick: Blood Fury grants 10% attack power and spell power for 15 seconds on a 2-minute cooldown, and the spell power half applies directly to Elemental's entire damage kit, making it a straightforward burst cooldown with no downside for a caster build.

## Professions

Community convention for caster Shaman favors **Enchanting** alongside **Tailoring** or a gathering profession for spell-power gear and enchant materials, though Shaman's totem and weapon-imbue kit doesn't lean as hard on any one profession as a pure caster class does.

## Leveling

Enhancement is generally the stronger leveling spec for Shaman, since its melee durability and weapon-imbue burst clear trash faster than Elemental's cast-time-dependent damage.
