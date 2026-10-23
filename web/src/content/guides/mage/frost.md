---
title: Frost Mage in Forever
classSlug: mage
spec: frost
role: dps
build: 'FS1:1.60.1.70532:mage:gnome:2030005/113023/25351013010030025:'
raidBuild: 'FS1:1.60.1.70532:mage:gnome:2030005/13102/253510130101301251:'
recommendedRaces: [gnome, troll]
statPriority: [Hit, Spell power, Frost power, Intellect, Critical strike, Spell haste, Spell penetration]
description: 'Talents, rotation, stats, and gear for Frost Mage in Forever, and what is confirmed versus projected from the beta.'
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
  - label: 'Forever Sixty rotation data, data/curated/apl/mage-frost.json'
    url: https://foreversixty.gg/sources
    kind: site
  - label: 'Forever Sixty raid loot data, data/curated/loot/forever-raid-phases.json'
    url: https://foreversixty.gg/sources
    kind: site
---

## Overview

Frost is the Mage control and survivability spec: Frostbolt keeps enemies slowed, with Frost Nova and Ice Block for emergencies and Ice Lance for burst.

### Strengths

- Frostbolt slows enemies and keeps them at range
- Frost Nova and Ice Block cover emergencies
- Ice Shards adds up to 100% Frost critical strike damage
- Frost Channeling cuts Frost mana cost by 15% and threat by 30%

### Weaknesses

- Less sustained output than Fire
- No burst cooldown
- Shatter works on a boss only through Fingers of Frost

### In a raid

- **Role:** Caster DPS with personal safety and utility
- **Brings:** Arcane Intellect for the group; Frostbolt slows the target

## Talents and builds

Blizzard confirmed the tree keeps its seven rows and 51 points, with a fourth one-point talent added at the 16-point mark alongside the existing ones at 11, 21, and 31.

- **Improved Frostbolt** — up to 0.5 seconds off Frostbolt's cast time at rank 5, the tree's most direct throughput gain.
- **Ice Shards** — up to 100% more Frost critical strike damage at rank 5, the biggest single damage multiplier in the tree.
- **Frost Channeling** — up to 15% off the mana cost of your Frost spells at rank 3, and 30% less threat from them. With Piercing Ice gone from the tree, it is where those three points went.
- **Shatter** — up to 50% more critical strike chance against a frozen target at rank 3. A raid boss cannot be frozen by Frost Nova, so against a boss this talent only works through Fingers of Frost, below.
- **Fingers of Frost** — each Chill effect (Frostbolt, Frostfire Bolt, Cone of Cold, Improved Blizzard) has a 15% chance to grant a 15 second effect that treats your next 1 spell (2 at rank 2) as if the target were Frozen. That switches on Shatter's crit bonus and Ice Lance's frozen damage bonus against a target that Frost Nova cannot freeze. It needs Ice Lance, so that is a one-point cost.
- **Winter's Chill** — a chance for Frost hits to stack a debuff that raises Ice Lance's and Frostbolt's crit chance against that target, up to 5 stacks at rank 5.

The build above comes from this site's talent search, held to the rule that a replacement must keep every talent the simulator cannot yet credit, stay in the Frost tree, and gain at least 1% beyond the run's measured error. It takes **Frost Channeling** at 3/3 (taking the place of Piercing Ice, which the client removed), **Wand Specialization** at 2/2 (25% more wand damage), **Improved Channeling** at 3/5 (a 42–60% chance to keep channeling Arcane Missiles through incoming damage — the engine has no code path yet to credit that resistance, so it's an honest unknown rather than a confirmed zero) and **Arcane Concentration** at 5/5 (a 10% chance per damage spell cast to enter Clearcasting, making your very next cast free — by far the largest single swing measured in this build). Fire keeps only the points that unlock Impact and Flame Throwing, which the simulator cannot credit but a Frost mage may still want for their control effects.

The latest change is Fingers of Frost, now modeled: the search's previous build had Shatter at 3/3 but nothing to make a boss count as frozen, so Shatter did nothing. Spending the Arcane Blast, Cold Snap, and Improved Frost Nova points on Ice Lance and Fingers of Frost 2/2 lifts the build by about 10% over the previous spend on the Frostbolt-only rotation, and the Ice Lance line below adds roughly 3% on top. Those three dropped talents are utility the simulator measures as no damage, not damage you lose; Cold Snap and Improved Frost Nova are real control tools, and a player who values them more than that gain can take them back. Ice Barrier stays out for the same reason as before: a pure-damage search does not weigh survivability. Open the planner at [/planner?class=mage](/planner?class=mage) to build this out.

**Raid build.** For a raid the search found about +3% over the leveling build above in the raid-ready run: it moves two ranks of Improved Fireball and the three Impact points (a stun) into Incineration and a third rank of Frost Channeling, which also takes 30% of the threat off Frost spells. The client removed Piercing Ice and trimmed Elemental Precision to 3 ranks, and the five points they held went to Frostbite (1/3), Ice Block, Cold Snap and Ice Barrier plus that third Frost Channeling rank: every Frost talent the simulator credits as damage was already maxed, so the spare points buy survivability and the cooldown reset. Ice Barrier needs 30 points below it, which the 34 Frost points meet. Winter's Chill, Ice Lance, Fingers of Frost, Frost Warding, Flame Throwing and Improved Channeling stay.

## Rotation and priority

Frostbolt is the filler, with the wand coming out once mana runs low. When Fingers of Frost procs, cast **Ice Lance** while the effect is up: the held charge makes the target count as frozen, so Ice Lance deals triple damage and Shatter's crit bonus applies, and each cast spends one charge. In the simulator the Ice Lance line gains roughly 3% over Frostbolt alone, beyond the run's error. Frostbolt's own Chill can roll the next charge, so keep casting while charges are held rather than waiting.

## Stat priority

The table above is this level range's own simulation at level 60, re-run by the nightly pipeline every time the build or its gear data changes — these numbers are never hand-entered. A stat shown as "not separable from zero" is not a verdict that it is worthless; the sim's measured error on it is too wide, at this level range's sample size, to tell its true value apart from zero.

## Best in slot

Look for pieces that lead with spell power, then hit until capped, then crit, following the order above.

## Enchants and consumables

Target spell power and hit on enchants and consumables, matching the stat priority above.

## Races

Mage's playable races are Human, Orc, Undead, Gnome, and Troll on their respective factions, plus Skyborne — but only the Alliance-side Skyborne, since Mage is Alliance-only for that race. Orc is one of Blizzard's six confirmed new race and class pairs, giving Horde a Mage option it didn't have in 1.12.

For Alliance, **Gnome** is the strongest pick: Escape Artist breaks and grants brief immunity to movement impairment on a 2-minute cooldown, which is a direct complement to a spec that already relies on kiting and control, and Expansive Mind's mana bonus supports a rotation that runs on continuous Frostbolt casts. For Horde, **Troll** is the pick: Berserking grants 10% casting and attack speed on a 3-minute cooldown, a straightforward way to squeeze more Frostbolts and Ice Lances into a burst window; it's also the new pairing Blizzard specifically added for this class, though for Warlock rather than Mage — Orc remains Mage's new Horde option, and its Blood Fury spell power cooldown is a reasonable alternative pick if Troll's timing doesn't line up with your raid's needs. Skyborne is available to Alliance Mages as a third option; its Read Ley Line racial restores all health and mana over 15 seconds, useful out of combat, but it doesn't add spell damage or control the way Gnome's kit does, so it's a pick for the race and story rather than raw throughput.

## Professions

Tailoring and Enchanting is the standard caster pairing in Classic-era play: Tailoring's crafted spellcaster gear fills early slots, and Enchanting lets you apply your own weapon and gear enchants instead of paying for them.

## Leveling

Frost is the strong leveling choice for Mage generally, and that holds for itself: Frostbolt's slow and Frost Nova's root make solo play safer than either Arcane's burst-dependent kit or Fire's closer-range playstyle.
