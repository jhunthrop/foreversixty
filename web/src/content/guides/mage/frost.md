---
title: Frost Mage in Forever
classSlug: mage
spec: frost
role: dps
build: 'FS1:1.60.1.70707:mage:gnome:230005/23023/253510130100301231:'
raidBuild: 'FS1:1.60.1.70707:mage:gnome:2300052/5102/253510130101301221:'
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

Frost is the Mage control and survivability spec: Frostbolt keeps enemies slowed, with Frost Nova and Ice Block for emergencies and Ice Lance on a Fingers of Frost charge.

### Strengths

- Frostbolt slows enemies and keeps them at range
- Frost Nova and Ice Block cover emergencies
- Ice Shards adds up to 100% Frost critical strike damage
- Frost Channeling cuts Frost mana cost by 15% and threat by 30%

### Weaknesses

- Less sustained output than Fire
- No burst cooldown: Icy Veins is not in this client's trees
- Shatter works on a boss only through Fingers of Frost

### In a raid

- **Role:** Caster DPS with personal safety and utility
- **Brings:** Arcane Intellect for the group

## Talents and builds

- **Improved Frostbolt** — up to 0.5 seconds off Frostbolt's cast time at rank 5, the tree's most direct throughput gain.
- **Ice Shards** — up to 100% more Frost critical strike damage at rank 5, the biggest single damage multiplier in the tree.
- **Frost Channeling** — up to 15% off the mana cost of your Frost spells at rank 3, and 30% less threat from them.
- **Shatter** — up to 50% more critical strike chance against a frozen target at rank 3. A raid boss cannot be frozen by Frost Nova, so against a boss this talent only works through Fingers of Frost, below.
- **Fingers of Frost** — each Chill effect (Frostbolt, Frostfire Bolt, Cone of Cold, Improved Blizzard) has a 15% chance to grant a 15 second effect that treats your next 1 spell (2 at rank 2) as if the target were Frozen. That switches on Shatter's crit bonus and Ice Lance's frozen damage bonus against a target that Frost Nova cannot freeze. It needs Ice Lance, so that is a one-point cost.
- **Winter's Chill** — reworked in Forever: every Frostbolt and every Fingers of Frost Ice Lance stacks Winter's Chill on the target, up to 5, and only your own next Cone of Cold spends them, taking 20% more damage and 7% more of your spell damage per stack. A raid frost mage stands at range and does not cast Cone of Cold, so the raid build keeps just the 2 points Ice Barrier needs below it; the leveling build keeps 3, since you are often in Cone of Cold range while leveling.
- **Missile Barrage** — Missile Barrage 2/2 gives Frostbolt a 20% chance to make your next Arcane Missiles free, halve its channel and fire the missiles every half second, and Precision (below) now covers those Arcane missiles' hit. Together with an Arcane Missiles cast on every proc it was worth about +7% in our audit run.
- **Ice Barrier** — the shield also raises your Frost damage by 15% for a minute, and by another 5% while the shield holds, so it is a damage talent first. Both builds take it: the leveling build moves two Winter's Chill points into **Cold Snap**, its prerequisite, and Ice Barrier itself, about +12% in our level-60 search run.

Outside Frost, **Precision** (3/3) is Elemental Precision renamed, now adding 6% hit chance to all your spells. The client moved **Improved Fire Blast** two rows down the Fire tree, out of reach of a ten-point dip, so its point went to **Incineration**, which also raises Ice Lance's critical strike chance; the raid build's Incineration goes to 5/5, taking that point and one of the three Arcane Focus held before the client removed it. Precision now does Arcane Focus's job for the Missile Barrage missiles, and its other two points go to **Wand Specialization**, which the Missile Barrage row needs above it.

## Rotation and priority

Put up **Ice Barrier** before the pull and recast it whenever its minute runs out, for its Frost damage bonus. Frostbolt is the filler. Cast Arcane Missiles whenever Missile Barrage procs. Wand only when you are out of mana: a wand shot cannot start during the global cooldown. When Fingers of Frost procs, cast **Ice Lance** while the effect is up: the held charge makes the target count as frozen, so Ice Lance deals four times its damage and Shatter's crit bonus applies, and each cast spends one charge. In the simulator the Ice Lance line gains roughly 3% over Frostbolt alone, beyond the run's error. Frostbolt's own Chill can roll the next charge, so keep casting while charges are held rather than waiting. Frostfire Bolt, which you can train from level 40, is not a Frost filler: swapping it in for Frostbolt measured roughly a third lower damage, because its direct hit is smaller and the Frost talents above scale Frostbolt, not it. Frost Nova and the other freezes are not part of the boss rotation, since a boss cannot be frozen. Keep Ice Barrier up at all times: it is the largest single Frost damage bonus you have. **Cone of Cold** only reaches 10 yards (12 with Arctic Reach), and raid frost mages stand at range, so skip it on most bosses. If a fight lets you stand that close (no cleave, tail or melee-range mechanic), Cone of Cold is worth a global on five Winter's Chill stacks: standing in range with more Winter's Chill points measured about 6% ahead of the ranged build in our raid run, but it puts you in melee range, so the guide keeps the ranged build.

## Stat priority

## Best in slot

## Enchants and consumables

## Races

## Professions

Tailoring and Enchanting is the standard caster pairing in Classic-era play: Tailoring's crafted spellcaster gear fills early slots, and Enchanting lets you apply your own weapon and gear enchants instead of paying for them.

## Leveling

Frost is the strong leveling choice for Mage generally, and that holds for itself: Frostbolt's slow and Frost Nova's root make solo play safer than either Arcane's burst-dependent kit or Fire's closer-range playstyle.
