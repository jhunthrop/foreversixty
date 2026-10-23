---
title: Destruction Warlock in Forever
classSlug: warlock
spec: destruction
role: dps
build: 'FS1:1.60.1.70532:warlock:gnome:25522/0/2053235103101351:'
raidBuild: 'FS1:1.60.1.70532:warlock:gnome:255/0005/2053045103101351:'
recommendedRaces: [gnome, troll]
statPriority:
  [Spell power, Intellect, Critical strike, Hit, Spell haste, Spell penetration, Shadow power, Fire power]
description: 'Talents, rotation, stats, and gear for Destruction Warlock in Forever, and what is confirmed versus projected from the beta.'
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
  - label: 'Forever Sixty rotation data, data/curated/apl/warlock-destruction.json'
    url: https://foreversixty.gg/sources
    kind: site
  - label: 'Forever Sixty raid loot data, data/curated/loot/forever-raid-phases.json'
    url: https://foreversixty.gg/sources
    kind: site
---

## Overview

Destruction is Warlock's direct-damage spec: it keeps Immolate on the target to fuel Conflagrate and Incinerate, and finishes with Shadowburn.

### Strengths

- Conflagrate and Incinerate both build on Immolate
- Shadowburn closes out low-health targets
- Hard-hitting casts and instants give high burst
- Corruption adds cheap sustained damage under the direct casts

### Weaknesses

- Gives up sustain compared with Affliction
- Expensive casts such as Incinerate and Soul Fire press its mana
- Loses value on every cast if Immolate lapses

### In a raid

- **Role:** Ranged fire and shadow burst damage
- **Brings:** A curse on the target (Bane of Doom in its rotation), plus Healthstones and a Soulstone

## Talents and builds

Blizzard confirmed the tree keeps its seven rows and 51 points, with a fourth one-point talent added at the 16-point mark alongside the existing ones at 11, 21, and 31.

- **Bane** — up to 0.5 seconds off Shadow Bolt, Immolate, and Incinerate cast time, and up to 2 seconds off Soul Fire, at rank 5.
- **Ruin** — up to a 100% critical-strike-damage bonus at rank 5 on all Destruction spells, the tree's biggest single multiplier.
- **Conflagrate** — an instant that ignites a target already afflicted by Immolate, dealing Fire damage and consuming the Immolate effect.
- **Shadowburn** — an instant Shadow nuke that refunds a Soul Shard if the target dies within 8 seconds of being hit.
- **Shadow and Flame** — up to a 10% damage buff at rank 5 from landing Conflagrate or Shadowburn, plus a chance for Conflagrate not to consume Immolate.
- **Incinerate** — the capstone: extra Fire damage that gains a further 25% if the target is afflicted by Immolate, Destruction's signature filler.

This build spends 34 points in Destruction to reach Incinerate at the bottom, with the remaining 16 in Affliction — Improved Corruption and Suppression both maxed, then Improved Life Tap, Soul Harvest and Grim Reach (the two-rank range talent the current client put in Malediction's slot) for the rest of the budget; Malediction's third point went to a third rank of Cataclysm. Against the first draft it drops three Improved Drains points (no damage in this rotation) and three Improved Shadow Bolt points, and spends them on Fire and Brimstone and Agonizing Flames. Improved Shadow Bolt is worth nothing to a solo caster here because the rotation fills with Incinerate rather than Shadow Bolt, so it is the cheapest place to find points; Shadow Bolt itself stays castable. The change was worth about +8% in the search run, and about +7% on the level-60 row of the rotation ladder. Molten Skin and Destructive Reach are kept even though the simulator cannot see them, and Conflagrate, Shadowburn and Incinerate are untouched. Open the planner at [/planner?class=warlock](/planner?class=warlock) to build this out.

**Raid build.** For a raid the search found about +3% over the leveling build above in the raid-ready run: it moves Soul Harvest (a solo-play mana return), the two Grim Reach points and three Cataclysm, which the simulator measures at little damage, into five Unholy Power and two Aftermath. Suppression, Molten Skin, Destructive Reach and every spell the rotation casts stay.

## Rotation and priority

Pop a mana potion at the start of the fight if your mana is already below its usual threshold. Open with Corruption and keep it up throughout: it's cheap sustained damage, and adding it was the single biggest gain the search found. Keep Bane of Doom running behind it, recast the moment it fires; it edged out Bane of Agony in our simulator, and on a fight shorter than a minute Agony is the curse to use. Together these changes were worth about +21% in our level-60 search run.

## Stat priority

The table above is this level range's own simulation at level 60, re-run by the nightly pipeline every time the build or its gear data changes — these numbers are never hand-entered. A stat shown as "not separable from zero" is not a verdict that it is worthless; the sim's measured error on it is too wide, at this level range's sample size, to tell its true value apart from zero.

## Best in slot

Look for pieces that lead with spell power, then hit until capped, then crit, following the order above.

## Enchants and consumables

Target spell power and hit on enchants and consumables, matching the stat priority above.

## Races

Warlock's playable races are Human, Orc, Undead, Gnome, and Troll on their respective factions; the class is not open to Skyborne at all. Troll is one of Blizzard's six confirmed new race and class pairs, giving the Horde a Warlock option it didn't have in 1.12.

For Alliance, **Gnome** is the strongest pick: Eureka!'s 50% mana-cost reduction and 10% damage bonus on the next three casts is a direct fit for a spec that leans on expensive direct-damage spells like Incinerate and Soul Fire, and Expansive Mind's mana bonus keeps that rotation running longer. For Horde, **Troll** is the pick as the new pairing Blizzard specifically added for this class: Berserking's 10% casting and attack speed on a 3-minute cooldown shortens the cast time on every direct-damage spell in the rotation, a straightforward fit for Destruction's burst-oriented kit.

## Professions

Tailoring and Enchanting is the standard caster pairing in Classic-era play: Tailoring's crafted spellcaster gear fills early slots, and Enchanting lets you apply your own weapon and gear enchants instead of paying for them.

## Leveling

Affliction is generally the stronger leveling choice for Warlock; Destruction can level well too, since Shadowburn and Conflagrate give it solid burst against single targets, but it lacks Affliction's DoT-and-move safety against multiple enemies.
