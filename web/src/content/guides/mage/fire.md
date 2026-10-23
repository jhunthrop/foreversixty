---
title: Fire Mage in Forever
classSlug: mage
spec: fire
role: dps
build: 'FS1:1.60.1.70532:mage:gnome:20500151/23552100130125031/003:'
raidBuild: 'FS1:1.60.1.70532:mage:gnome:20500151/23452100130135031/003:'
recommendedRaces: [gnome, orc]
statPriority: [Hit, Intellect, Spell power, Fire power, Critical strike, Spell haste, Spell penetration]
description: 'Talents, rotation, stats, and gear for Fire Mage in Forever, and what is confirmed versus projected from the beta.'
updated: 2026-10-07
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

Fire is the Mage sustained single-target spec: Scorch stacks a Fire vulnerability debuff, and Combustion turns it into guaranteed critical strikes.

### Strengths

- Combustion makes Fire spell hits critical until three non-periodic crits land
- Ignite turns every Fire crit into extra damage over time
- Scorch is cheap, at 150 mana for a 1.5 second cast
- Fire Blast adds a free hit without competing for a global

### Weaknesses

- Mana runs short before a long fight ends
- Damage is confined to the Fire school
- Pyroblast pays off only on three stacks of Heating Up

### In a raid

- **Role:** Steady single-target caster DPS
- **Brings:** Improved Scorch, a Fire vulnerability debuff of 3% per stack, up to 5 stacks, on the target

## Talents and builds

Blizzard confirmed the tree keeps its seven rows and 51 points, with a fourth one-point talent added at the 16-point mark alongside the existing ones at 11, 21, and 31.

- **Improved Fireball** — up to 0.5 seconds off Fireball's and Frostfire Bolt's cast time at rank 5, a direct throughput gain taken early.
- **Ignite** — Fire critical strikes burn the target for a further 40% of the spell's damage over 4 seconds at rank 5, turning every crit into extra sustained damage.
- **Improved Scorch** — up to a guaranteed chance for Scorch to stack a 3%-per-stack Fire vulnerability debuff, up to 5 stacks, on the target.
- **Critical Mass** — more Fire critical strike chance, feeding both Ignite and Combustion. The current client gives it five ranks and makes all five the prerequisite for Combustion, so the build takes it at 5/5.
- **Fire Power** — a flat Fire damage bonus, the tree's biggest raw damage talent, now capped at three ranks; the two points it no longer holds are the two that went into Critical Mass.
- **Pyroblast** and **Heating Up** — a pair of points in the middle of the tree. Heating Up is the renamed Hot Streak: every non-periodic critical strike with Fireball, Frostfire Bolt, Fire Blast, or Scorch cuts the cast time of your next Pyroblast by 25% for 20 seconds, stacking up to three times, so a Pyroblast cast on three stacks takes a quarter of its usual 6 seconds. Casting Pyroblast spends every stack.
- **Combustion** — the capstone: each Fire spell hit adds 10% Fire crit chance, lasting until you land three non-periodic Fire crits.

This build spends 34 points in Fire to reach Combustion at the bottom, and two of them are Pyroblast and Heating Up. They were bought with the points of **Master of Elements** (2/3 → 0/3), which refunded 20% of the base mana cost on your Fire and Frost critical strikes: the sim measured the swap, with the Pyroblast line added to the rotation, as about 3.5% more damage than the previous spend and rotation. The client has since trimmed Elemental Precision to 3 ranks, and the two points that freed went back into Master of Elements (2/3), so the mana refund is back and Pyroblast and Heating Up stay. Taking Pyroblast and Heating Up's points from Wake of Fire or Incineration instead measured the same within the sim's error, so that is a matter of taste.

Outside Fire, the spend drops from 20 points in Arcane to 14 — losing **Arcane Focus** entirely (5/5 → 0/5) and **Arcane Subtlety** entirely (1/2 → 0/2), neither of which registered any damage in the sim's model for this rotation — and picks up 3 points in Frost instead, all into **Elemental Precision** (3/3, 6% more hit chance with Frost and Fire spells), which this build's Fire-heavy rotation measures as a real damage gain since Scorch and Fire Blast both benefit directly. Open the planner at [/planner?class=mage](/planner?class=mage) to build this out.

**Raid build.** For a raid the search found a small gain, about +1% over the leveling build above in the raid-ready run, from moving one Improved Fireball point (the build keeps 4 of 5) into a full three ranks of Master of Elements. Improved Scorch, Combustion, Flame Throwing and Impact all stay.

## Rotation and priority

Open by stacking Improved Scorch's vulnerability debuff with a few Scorch casts before your main damage window, then hold Combustion until the debuff is fully stacked so every Fire spell hit during Combustion lands as a guaranteed crit. Heating Up decides when Pyroblast is worth a global: cast it only once Heating Up has reached three stacks, which cuts it to a 1.5 second cast, and never at the full 6 seconds. Fireball is only the filler before you train Scorch at level 22.

## Stat priority

The table above is this level range's own simulation at level 60, re-run by the nightly pipeline every time the build or its gear data changes — these numbers are never hand-entered. A stat shown as "not separable from zero" is not a verdict that it is worthless; the sim's measured error on it is too wide, at this level range's sample size, to tell its true value apart from zero.

## Best in slot

Look for pieces that lead with spell power, then hit until capped, then crit, following the order above.

## Enchants and consumables

Target spell power and hit on enchants and consumables, matching the stat priority above.

## Races

Mage's playable races are Human, Orc, Undead, Gnome, and Troll on their respective factions, plus Skyborne — but only the Alliance-side Skyborne, since Mage is Alliance-only for that race. Orc is one of Blizzard's six confirmed new race and class pairs, giving Horde a Mage option it didn't have in 1.12.

For Alliance, **Gnome** is the strongest pick: Eureka! cuts the mana cost of your next three abilities by 50% and adds 10% more damage on a 2-minute cooldown, which lines up well with a burst of Scorch casts around a Combustion window, and Expansive Mind's mana bonus helps sustain a spec that rarely stops casting. For Horde, **Orc** is the pick: Blood Fury grants 10% spell power for 15 seconds on a 2-minute cooldown, a straightforward damage cooldown to pair with Combustion, and it's the new pairing Blizzard specifically added for this class. Skyborne is available to Alliance Mages as a third option; its Read Ley Line racial restores all health and mana over 15 seconds, useful out of combat, but it doesn't add spell damage the way Gnome's Eureka! does, so it's a pick for the race and story rather than raw throughput.

## Professions

Tailoring and Enchanting is the standard caster pairing in Classic-era play: Tailoring's crafted spellcaster gear fills early slots, and Enchanting lets you apply your own weapon and gear enchants instead of paying for them.

## Leveling

Frost is the spec most players lean on for leveling, since Frostbolt's slow and Frost Nova's root make it safer to solo than Fire's more contact-heavy kit. Fire can still level well once Improved Scorch and Combustion are both online.
