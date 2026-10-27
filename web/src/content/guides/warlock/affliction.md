---
title: Affliction Warlock in Forever
classSlug: warlock
spec: affliction
role: dps
build: 'FS1:1.60.1.70630:warlock:gnome:25522300130201351/200522/005:'
raidBuild: 'FS1:1.60.1.70630:warlock:gnome:25520300130201351/00052/0550001:'
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
- **Brings:** A curse on the target (Bane of Doom in its rotation), plus Healthstones and a Soulstone

## Talents and builds

- **Suppression** — up to 5% more hit chance and 20% less threat at rank 5, both scarce and valuable for a spec that casts constantly.
- **Improved Corruption** — cuts Corruption's cast time by up to 2 seconds and adds up to 10% more damage at rank 5.
- **Amplify Curse** (1 point) — a required pick, not an optional one, and missing from an earlier draft of this guide's talent section: this build's rotation prepull-casts it before the pull, so it has to be taken for that line to be legal.
- **Pandemic** — up to a 100% critical-strike-damage bonus at rank 3 on Corruption, both Banes, and the drain spells. This build takes two of its three points; the search run measured no damage in the last one, so it went elsewhere.
- **Nightfall** — up to a 4% chance per DoT tick at rank 2 to make your next Shadow Bolt instant, a real mana and time saver.
- **Shadow Mastery** — up to 5% more Shadow damage and life drained at rank 5.
- **Wrack** — the capstone: a DoT that also boosts your other Shadow DoTs on the same target by 10% while it's active.

## Rotation and priority

Pop a mana potion or gem at the start of the fight if your mana is already below its usual threshold, or save it for a closing burn if the fight is nearly over. From there, put Corruption up first and keep it up throughout the fight — it's the cheapest DoT per point of damage, so upkeep on it should never lapse. Bane of Doom goes up right behind it as your one curse: it does all of its damage in a single hit when its minute is up, so recast it the moment it fires. In our simulator it beat Bane of Agony by a small margin and Curse of the Elements by a wide one, since a lone caster's own damage gains less from Elements than Doom or Agony deal themselves; on a fight shorter than a minute Doom never lands, so use Agony there. Wrack is live in this site's simulator: an instant, no-cooldown DoT that also stacks a 10% Shadow-damage-taken vulnerability on the target, so recast it in its last second to keep the vulnerability from ever dropping. Immolate is the fourth DoT: the search found that a running Immolate out-damages the Shadow Bolt global it costs, so it joins the upkeep order ahead of the filler, and that recasting it when under 3 seconds remain, rather than after it lapses, measured a little better. Once all four DoTs are running, spend every remaining global on Shadow Bolt. Once mana falls under 50%, wand with Shoot instead: it costs nothing, so mana regenerates while the DoTs keep ticking, and Life Tap is only the backstop below about 10%. Together these changes were worth about +6% in our level-60 search run.

## Stat priority

## Best in slot

## Enchants and consumables

## Races

## Professions

Tailoring and Enchanting is the standard caster pairing in Classic-era play: Tailoring's crafted spellcaster gear fills early slots, and Enchanting lets you apply your own weapon and gear enchants instead of paying for them.

## Leveling

Affliction is generally considered the strong leveling choice for Warlock, since its DoTs keep working while you move or fight multiple targets, and Drain Life covers the healing a solo caster otherwise lacks.
