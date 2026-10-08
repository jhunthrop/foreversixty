---
title: Fire Mage in Forever
classSlug: mage
spec: fire
role: dps
build: 'FS1:1.60.1.70009:mage:gnome:2050151/23552100130103051/005:'
raidBuild: 'FS1:1.60.1.70009:mage:gnome:2050151/23252100130133051/005:'
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

Fire is Mage's sustained single-target burner: it builds a stacking vulnerability debuff on its target through Scorch, then cashes it in for guaranteed critical strikes through Combustion, filling the rest of the fight with Scorch and Fire Blast. In a raid it's a steady, high-uptime caster rather than a spike-and-wait spec. The beta only reaches level 30, so nothing here about level 60 talents, raiding, or best-in-slot gear has actually been played — it is a projection from the demo trees and 1.12 knowledge, not tested content.

## Talents and builds

Blizzard confirmed the tree keeps its seven rows and 51 points, with a fourth one-point talent added at the 16-point mark alongside the existing ones at 11, 21, and 31. Verified against this build's talent data, the Fire tree's key picks in roughly the order you'd take them are:

- **Improved Fireball** — up to 0.5 seconds off Fireball's and Frostfire Bolt's cast time at rank 5, a direct throughput gain taken early.
- **Ignite** — Fire critical strikes burn the target for a further 40% of the spell's damage over 4 seconds at rank 5, turning every crit into extra sustained damage.
- **Improved Scorch** — up to a guaranteed chance for Scorch to stack a 3%-per-stack Fire vulnerability debuff, up to 5 stacks, on the target.
- **Critical Mass** — up to 6% more Fire critical strike chance at rank 3, feeding both Ignite and Combustion.
- **Fire Power** — a flat 10% more Fire damage at rank 5, the tree's biggest raw damage talent.
- **Pyroblast** and **Heating Up** — a pair of points in the middle of the tree. Heating Up is the renamed Hot Streak: every non-periodic critical strike with Fireball, Frostfire Bolt, Fire Blast, or Scorch cuts the cast time of your next Pyroblast by 25% for 20 seconds, stacking up to three times, so a Pyroblast cast on three stacks takes a quarter of its usual 6 seconds. Casting Pyroblast spends every stack.
- **Combustion** — the capstone: each Fire spell hit adds 10% Fire crit chance, lasting until you land three non-periodic Fire crits.

This build spends 32 points in Fire to reach Combustion at the bottom, and two of them are Pyroblast and Heating Up. They come out of **Master of Elements** (2/3 → 0/3), which refunded 20% of the base mana cost on your Fire and Frost critical strikes: the sim measures the swap, with the Pyroblast line added to the rotation, as about 3.5% more damage than the previous spend and rotation. Taking those two points from Wake of Fire or Incineration instead measures the same within the sim's error, so that is a matter of taste if you would rather keep the mana refund.

Outside Fire, the spend drops from 20 points in Arcane to 14 — losing **Arcane Focus** entirely (5/5 → 0/5) and **Arcane Subtlety** entirely (1/2 → 0/2), neither of which registered any damage in the sim's model for this rotation — and picks up 5 points in Frost instead, all into **Elemental Precision** (5/5, 5% more hit chance with Frost and Fire spells), which this build's Fire-heavy rotation measures as a real damage gain since Scorch and Fire Blast both benefit directly. Open the planner at [/planner?class=mage](/planner?class=mage) to build this out.

**Raid build.** For a raid the search found a small gain, about +1% over the leveling build above in the raid-ready run, from moving the three Improved Fireball points into Master of Elements. Improved Scorch, Combustion, Flame Throwing and Impact all stay.

## Rotation and priority

Open by stacking Improved Scorch's vulnerability debuff with a few Scorch casts before your main damage window, then hold Combustion until the debuff is fully stacked so every Fire spell hit during Combustion lands as a guaranteed crit. Heating Up decides when Pyroblast is worth a global: cast it only once Heating Up has reached three stacks, which cuts it to a 1.5 second cast, and never at the full 6 seconds. In the sim a three-stack Pyroblast is worth roughly 3.5% more damage over the whole fight than leaving it out, while casting it on two stacks measured no better than not casting it and casting it on one stack lost damage. Fire Blast slots in on cooldown as a free extra hit that doesn't compete with your Scorch casts for a global cooldown, and its crits feed Heating Up like the rest. Scorch is the filler for every global that isn't spent on a three-stack Pyroblast, Fire Blast, or Combustion: it costs 150 mana for a 1.5 second cast against Fireball's 395 for 3 seconds, a Fire mage runs short of mana long before a fight ends, and every Scorch also renews the debuff, so it never falls off. In the sim, Scorch as the filler beat Fireball by about 10 percent on a level 60 character over a three to five minute fight, and by 5 to 8 percent over one minute, at levels 40, 50 and 60 alike; a mix that casts Fireball while mana is high and Scorch when it is low measured below Scorch alone. Fireball is only the filler before you train Scorch at level 22. Frostfire Bolt (trained at levels 40, 50 and 60) feeds Heating Up and Ignite like Fireball and measured within error of it, so it is no upgrade on Scorch either. Pyroblast's own damage-over-time is minor next to the hit; the point of the buff is the fast cast. This site's simulator only models the Combustion-stacking loop when Improved Scorch is actually talented, and without Heating Up it skips the Pyroblast line, so take Improved Scorch if you want the Combustion-stacking loop above to matter.

## Stat priority

The table above is this band's own simulation at level 60, re-run by the nightly pipeline every time the build or its gear data changes — these numbers are never hand-entered. A stat shown as "not separable from zero" is not a verdict that it is worthless; the sim's measured error on it is too wide, at this band's sample size, to tell its true value apart from zero. These numbers come from this site's own level-60 simulator, not from beta play, which only reaches level 30.

## Gear

Look for pieces that lead with spell power, then hit until capped, then crit, following the order above. Specific slot-by-slot picks can't be named honestly yet: the beta caps at level 30, so no one has gear-checked Fire past the early game, and raid loot doesn't exist yet either — nothing raids until 9 December, and even then this client's item table is missing most of the classic raid loot tables it's meant to carry (Onyxia's Lair's own list is currently empty, and Barrow Deeps and Hyjal Summit aren't mapped to items at all). This section will fill in with real picks once raid loot is itemized.

## Enchants and consumables

Target spell power and hit on enchants and consumables, matching the stat priority above. **Enchant Weapon - Spell Power** (verified in this build's enchant data, +30 spell damage) is the standard weapon enchant for a caster. **Elixir of Fire Power** is a verified school-specific damage elixir in this build's data and is the natural pick for Fire over a generic spell-power elixir when the two conflict; **Flask of Supreme Power** covers the general spell-damage slot on a raid night. Beyond naming those, keep it principle-level until raid-tier consumables are confirmed for Forever specifically.

## Races

Mage's playable races are Human, Orc, Undead, Gnome, and Troll on their respective factions, plus Skyborne — but only the Alliance-side Skyborne, since Mage is Alliance-only for that race. Orc is one of Blizzard's six confirmed new race and class pairs, giving Horde a Mage option it didn't have in 1.12.

For Alliance, **Gnome** is the strongest pick: Eureka! cuts the mana cost of your next three abilities by 50% and adds 10% more damage on a 2-minute cooldown, which lines up well with a burst of Scorch casts around a Combustion window, and Expansive Mind's mana bonus helps sustain a spec that rarely stops casting. For Horde, **Orc** is the pick: Blood Fury grants 10% spell power for 15 seconds on a 2-minute cooldown, a straightforward damage cooldown to pair with Combustion, and it's the new pairing Blizzard specifically added for this class. Skyborne is available to Alliance Mages as a third option; its Read Ley Line racial restores all health and mana over 15 seconds, useful out of combat, but it doesn't add spell damage the way Gnome's Eureka! does, so it's a pick for the race and story rather than raw throughput.

## Professions

Tailoring and Enchanting is the standard caster pairing in Classic-era play: Tailoring's crafted spellcaster gear fills early slots, and Enchanting lets you apply your own weapon and gear enchants instead of paying for them. This is general Classic-community convention rather than anything Forever-specific, so treat it as inferred until this site has profession data from the beta.

## Leveling

Frost is the spec most players lean on for leveling, since Frostbolt's slow and Frost Nova's root make it safer to solo than Fire's more contact-heavy kit. Fire can still level well once Improved Scorch and Combustion are both online, but that's inferred from 1.12 knowledge, not tested — the beta only runs to level 30, so nothing about the leveling experience above that has been played in Forever's actual client.
