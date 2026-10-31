---
title: Shadow Priest in Forever
classSlug: priest
spec: shadow
role: dps
build: 'FS1:1.60.1.70658:priest:gnome:52401003113/0/54342001201300051:'
raidBuild: 'FS1:1.60.1.70658:priest:gnome:32500003113/0/52352101201300251:'
recommendedRaces: [gnome, undead]
statPriority: [Intellect, Hit, Spell power, Shadow power, Critical strike, Spell haste, Spell penetration]
description: 'Shadow Priest overview, talent priority, rotation, stat weights, and race picks for Forever, with beta-versus-projection called out.'
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

Shadow is the Priest damage-over-time caster: keep Shadow Word: Pain up, then Mind Blast and Mind Flay, with Vampiric Embrace healing as sustain.

### Strengths

- Strongest Priest spec for solo play and leveling
- Shadow Weaving stacks a growing Shadow damage multiplier
- Shadow Word: Death adds a burst of damage below 20% health
- Vampiric Embrace gives self-sufficiency through damage dealt

### Weaknesses

- Mind Flay costs a talent point, and without it there is no filler spell
- Damage depends on keeping Shadow Word: Pain up
- A close fight wastes a late Shadow Word: Pain refresh, so it is skipped

### In a raid

- **Role:** Caster damage dealer
- **Brings:** Vampiric Embrace healing for the group, Power Word: Fortitude, and Shadow Weaving raising Shadow damage on the target

## Talents and builds

**Improved Mind Blast** still cuts Mind Blast's cooldown and holds 4 of its 5 points: the current client folded Shadow Affinity out of the tree, and the threat point it used to take goes here. The engine measures each of those points at no damage in this build, so the point moves to **Mental Agility**, taken at 3/3 (a mana-cost cut on Smite, Holy Fire and instant casts, measured by the engine as a real damage gain). **Silent Resolve**'s point (a threat talent, zero damage) and both points in **Early Demise** go to the same place, and the freed Shadow points go into a second rank of **Improved Shadow Word: Pain** (6 extra seconds of Shadow Word: Pain duration), which the probe credits as a clear damage gain. Early Demise, which raised Shadow Word: Death's critical strike chance below 20% health, is no longer taken: the engine measures it at zero damage, and Shadow Word: Death is still cast in the execute window. **Power in Light**, **Holy Precision**, **Improved Power Word: Shield** and **Blackout** hold their ranks: the engine has no code path for them, so this site cannot measure what they are worth and the build leaves them alone. This build measures about +6.9% over the previous spend in our level-60 search run.

Point allocation runs deep into Shadow to reach Shadowform at the bottom of the tree, 33 points, with the remaining 18 in Discipline: **Inner Focus** (1 point) is a required pick alongside the Meditation talent's mana regeneration while casting — it's the free, empowered Mind Blast cast this build's Rotation section calls out below, so it has to be taken for that line to be legal, not just Meditation. Open the planner at [/planner?class=priest](/planner?class=priest) to build this out.

## Rotation and priority

The core loop this site's simulator plays: keep Shadow Word: Pain up on the target unless the fight is close to ending, since refreshing it late would waste most of the remaining damage-over-time value. Keep Devouring Plague up too; in Forever every race learns it, not only Undead. Its 24-second cooldown lets you reapply it as it expires, and in Forever its ticks crit like Shadow Word: Pain's. Use Inner Focus on Mind Blast whenever it is ready, then cast Mind Blast on cooldown, ahead of Mind Flay: it is the bigger hit for its cast time, and Shadowform halves its cost and doubles its crit bonus like every other Shadow spell. Below 20% target health, Shadow Word: Death comes next. Mind Flay fills every other global. Mind Flay needs its talent taken, since Forever gates it behind that pick rather than granting it as a baseline trainer spell (see Talents and builds below for when to take it while leveling). Wand only when nothing else can be cast: a wand shot cannot start during the global cooldown, so it never fits between spells. Reordering this way, Mind Blast over Mind Flay and Devouring Plague kept up, was worth about +12% at the time of the audit. The raid build still carries two points in Wand Specialization: wand shots no longer fit between spells, so those points do almost nothing, but moving them measured under 1% and the build keeps them until a larger change comes along.

## Stat priority

## Best in slot

Because bonus healing gear in Forever now also carries a fraction of bonus damage, some cloth pieces itemized for healers may carry incidental value for Shadow that they wouldn't have in 1.12 — worth checking case by case rather than assuming.

## Enchants and consumables

## Races

## Professions

Community convention for Shadow Priests favors **Tailoring** for caster cloth stats alongside **Enchanting** for weapon and gear enchants, the same combination general caster practice recommends across classes.

## Leveling

Shadow is the strongest leveling spec for Priest by a wide margin: its damage-over-time kit provides both offense and self-healing that Discipline and Holy, built around supporting other players, don't have solo.

Take **Mind Flay**'s single point the moment it's available at level 20, ahead of the level-60 build's own point order above. That order maxes Shadow Focus, Blackout and Spirit Tap first for a level-60 raider's hit and threat needs, but a leveling character walking those same points in sequence doesn't reach Mind Flay's row until several levels later — and until then there is no ranged filler at all beyond the wand, since Forever moved Mind Flay out of the baseline kit. Grab the 1-point prerequisite first, then fill Shadow Focus/Blackout/Spirit Tap and the rest of the tree behind it.
