---
title: Affliction Warlock in Forever
classSlug: warlock
spec: affliction
role: dps
build: 'FS1:1.60.1.70707:warlock:gnome:25522300130201351/200522/005:'
raidBuild: 'FS1:1.60.1.70707:warlock:gnome:25520300130201351/00052/0550001:'
recommendedRaces: [gnome, troll]
statPriority: [Spell power, Shadow power, Intellect, Hit, Critical strike, Spell haste, Spell penetration]
description: 'Talents, rotation, stats, and gear for Affliction Warlock in Forever, and what is confirmed versus projected from the beta.'
updated: 2026-10-08
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
  - label: 'Forever Sixty rotation data, data/curated/apl/warlock-affliction.json'
    url: https://foreversixty.gg/sources
    kind: site
  - label: 'Forever Sixty raid loot data, data/curated/loot/forever-raid-phases.json'
    url: https://foreversixty.gg/sources
    kind: site
---

## Overview

Affliction is Warlock's damage-over-time spec: it stacks several DoTs on a target and keeps casting with Life Tap and drain spells.

### Strengths

- High DoT uptime with little upkeep once the DoTs are applied
- Keeps dealing damage while moving
- Life Tap and drain spells keep mana and health topped up
- Corruption is the cheapest DoT per point of damage

### Weaknesses

- Life Tap spends health to fund mana
- Damage ramps up slowly, since DoTs must be applied first
- Less burst than Destruction

### In a raid

- **Role:** Ranged shadow damage over time
- **Brings:** A curse on the target (Bane of Agony in its rotation), plus Healthstones and a Soulstone

## Talents and builds

- **Suppression** — up to 5% more hit chance and 20% less threat at rank 5, both scarce and valuable for a spec that casts constantly.
- **Improved Corruption** — cuts Corruption's cast time by up to 2 seconds and adds up to 10% more damage at rank 5.
- **Amplify Curse** (1 point) — a required pick, not an optional one, and missing from an earlier draft of this guide's talent section: this build's rotation prepull-casts it before the pull, so it has to be taken for that line to be legal.
- **Pandemic** — up to a 100% critical-strike-damage bonus at rank 3 on Corruption, both Banes, and the drain spells. This build takes two of its three points; the search run measured no damage in the last one, so it went elsewhere.
- **Nightfall** — up to a 4% chance per DoT tick at rank 2 to make your next Shadow Bolt instant, a real mana and time saver.
- **Shadow Mastery** — up to 5% more Shadow damage and life drained at rank 5.
- **Wrack** — the capstone: a DoT that also boosts your other Shadow DoTs on the same target by 10% while it's active.

## Rotation and priority

Pop a mana potion or gem at the start of the fight if your mana is already below its usual threshold, or save it for a closing burn if the fight is nearly over. From there, put Corruption up first and keep it up throughout the fight — it's the cheapest DoT per point of damage, so upkeep on it should never lapse. Bane of Agony goes up right behind it as your one curse. In Forever your damage over time can crit, Wrack adds 10% to your other Shadow DoTs, and Agony is an Affliction effect for Soul Siphon, so the ticking curse beats Bane of Doom's single hit. Wrack is an instant, no-cooldown DoT that also stacks a 10% Shadow-damage-taken vulnerability on the target, so recast it in its last second to keep the vulnerability from ever dropping. Siphon Life comes next: the build takes it, its ticks crit and heal you, and every Affliction effect on the target raises Soul Siphon's bonus. Immolate is the fifth DoT: a running Immolate out-damages the Shadow Bolt global it costs, and recasting it when under 3 seconds remain measured a little better than waiting for it to lapse. Once every DoT is running, spend each remaining global on Shadow Bolt. Life Tap keeps you casting once mana falls below about 10%. Wand only when you are out of mana: a wand shot cannot start during the global cooldown, so it never fits between spells. Keeping Siphon Life up and switching to Agony were worth about +4% in our audit run.

## Stat priority

## Best in slot

## Enchants and consumables

## Races

## Professions

Tailoring and Enchanting is the standard caster pairing in Classic-era play: Tailoring's crafted spellcaster gear fills early slots, and Enchanting lets you apply your own weapon and gear enchants instead of paying for them.

## Leveling

Affliction is generally considered the strong leveling choice for Warlock, since its DoTs keep working while you move or fight multiple targets, and Drain Life covers the healing a solo caster otherwise lacks.
