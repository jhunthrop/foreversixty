---
title: Fire Mage in Forever
classSlug: mage
spec: fire
role: dps
build: 'FS1:1.60.1.70707:mage:gnome:251015/55521001230123031/003:'
raidBuild: 'FS1:1.60.1.70707:mage:gnome:251015/54521001230133031/003:'
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

- **Accelerant** — up to 0.5 seconds off Fireball's, Frostfire Bolt's and Arcane Wrath's cast time at rank 5, a direct throughput gain taken early. It is Improved Fireball renamed, with Arcane Wrath added, in the same slot.
- **Incineration** — up to 10% more critical strike chance on Fire Blast and Scorch at rank 5. The client gave it two more ranks when it took two from Critical Mass, and the build moves its two points with them: 5/5.
- **Ignite** — Fire critical strikes burn the target for a further 40% of the spell's damage over 4 seconds at rank 5, turning every crit into extra sustained damage.
- **Improved Scorch** — up to a guaranteed chance for Scorch to stack a 3%-per-stack Fire vulnerability debuff, up to 5 stacks, on the target.
- **Critical Mass** — more Fire critical strike chance, feeding both Ignite and Combustion. The current client gives it three ranks and makes all three the prerequisite for Combustion, so the build takes it at 3/3.
- **Fire Power** — a flat Fire damage bonus, the tree's biggest raw damage talent, now capped at three ranks.
- **Pyroblast** and **Heating Up** — a pair of points in the middle of the tree. Heating Up is the renamed Hot Streak: every non-periodic critical strike with Fireball, Frostfire Bolt, Fire Blast, or Scorch cuts the cast time of your next Pyroblast by 25% for 20 seconds, stacking up to three times, so a Pyroblast cast on three stacks takes a quarter of its usual 6 seconds. Casting Pyroblast spends every stack.
- **Combustion** — the capstone: each Fire spell hit adds 10% Fire crit chance, lasting until you land three non-periodic Fire crits.

This build spends 34 points in Fire to reach Combustion at the bottom, and two of them are Pyroblast and Heating Up. They were bought with the points of **Master of Elements** (2/3 → 0/3), which refunded 20% of the base mana cost on your Fire and Frost critical strikes: the sim measured the swap, with the Pyroblast line added to the rotation, as about 3.5% more damage than the previous spend and rotation. The client has since trimmed Precision (then Elemental Precision) to 3 ranks, and the two points that freed went back into Master of Elements (2/3), so the mana refund is back and Pyroblast and Heating Up stay. Taking Pyroblast and Heating Up's points from Wake of Fire or Incineration instead measured the same within the sim's error, so that is a matter of taste.

Outside Fire, the spend drops from 20 points in Arcane to 14 — losing **Arcane Focus** (since removed from the game) and **Arcane Subtlety** entirely, neither of which registered any damage in the sim's model for this rotation — and picks up 3 points in Frost instead, all into **Precision** (3/3, 6% more hit chance with all your spells; the client renamed Elemental Precision and widened it from Frost and Fire to every school), a real damage gain for this rotation since Fireball, Scorch and Fire Blast all benefit directly. Open the planner at [/planner?class=mage](/planner?class=mage) to build this out.

## Rotation and priority

Open by stacking Improved Scorch's vulnerability debuff with a few Scorch casts before your main damage window, then hold Combustion until the debuff is fully stacked so every Fire spell hit during Combustion lands as a guaranteed crit. Heating Up decides when Pyroblast is worth a global: cast it only once Heating Up has reached three stacks, which cuts it to a 1.5 second cast, and never at the full 6 seconds. In the sim a three-stack Pyroblast is worth roughly 3.5% more damage over the whole fight than leaving it out, while casting it on two stacks measured no better than not casting it and casting it on one stack lost damage. Fire Blast slots in on cooldown as a free extra hit that doesn't compete with your Scorch casts for a global cooldown, and its crits feed Heating Up like the rest. Scorch is the filler for every global that isn't spent on a three-stack Pyroblast, Fire Blast, or Combustion: it costs 150 mana for a 1.5 second cast against Fireball's 395 for 3 seconds, a Fire mage runs short of mana long before a fight ends, and every Scorch also renews the debuff, so it never falls off. In the sim, Scorch as the filler beat Fireball by about 10 percent on a level 60 character over a three to five minute fight, and by 5 to 8 percent over one minute, at levels 40, 50 and 60 alike; a mix that casts Fireball while mana is high and Scorch when it is low measured below Scorch alone. Fireball is only the filler before you train Scorch at level 22. Frostfire Bolt (trained at levels 40, 50 and 60) feeds Heating Up and Ignite like Fireball and measured within error of it, so it is no upgrade on Scorch either. Pyroblast's own damage-over-time is minor next to the hit; the point of the buff is the fast cast. This site's simulator only models the Combustion-stacking loop when Improved Scorch is actually talented, and without Heating Up it skips the Pyroblast line, so take Improved Scorch if you want the Combustion-stacking loop above to matter.

## Stat priority

## Best in slot

## Enchants and consumables

## Races

## Professions

Tailoring and Enchanting is the standard caster pairing in Classic-era play: Tailoring's crafted spellcaster gear fills early slots, and Enchanting lets you apply your own weapon and gear enchants instead of paying for them.

## Leveling

Frost is the spec most players lean on for leveling, since Frostbolt's slow and Frost Nova's root make it safer to solo than Fire's more contact-heavy kit. Fire can still level well once Improved Scorch and Combustion are both online.
