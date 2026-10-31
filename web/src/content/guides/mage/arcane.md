---
title: Arcane Mage in Forever
classSlug: mage
spec: arcane
role: dps
build: 'FS1:1.60.1.70707:mage:gnome:141005213100301531/32023/0532:'
raidBuild: 'FS1:1.60.1.70707:mage:gnome:131205213100301531/3/05350003:'
recommendedRaces: [gnome, orc]
statPriority: [Spell power, Arcane power, Hit, Critical strike, Intellect, Spell haste, Spell penetration]
description: 'Talents, rotation, stats, and gear for Arcane Mage in Forever, and what is confirmed versus projected from the beta.'
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
  - label: 'Forever Sixty rotation data, data/curated/apl/mage-arcane.json'
    url: https://foreversixty.gg/sources
    kind: site
  - label: 'Forever Sixty raid loot data, data/curated/loot/forever-raid-phases.json'
    url: https://foreversixty.gg/sources
    kind: site
---

## Overview

Arcane is the Mage burst spec: Arcane Power and Presence of Mind concentrate damage into short windows on the Arcane Blast loop.

### Strengths

- Arcane Power adds 30% spell damage for 15 seconds
- Presence of Mind makes the next cast instant
- Arcane Concentration gives free Clearcasting casts
- Missile Barrage halves Arcane Missiles' channel and makes it free

### Weaknesses

- Less sustained damage than Fire once the target's debuffs are stacked
- Arcane Power costs 30% more mana per cast
- Each Arcane Blast stack raises the next Blast's mana cost by 100%

### In a raid

- **Role:** Single-target caster DPS with burst on cooldown
- **Brings:** Arcane Intellect for the group; no raid-wide debuff

## Talents and builds

- **Precision** (Frost tree) — up to 6% more hit chance with all your spells at rank 3. The client removed Arcane Focus and widened Elemental Precision, now named Precision, from Frost and Fire to every school, so the three Frost points this build already spent carry the Arcane hit Arcane Focus used to.
- **Arcane Concentration** — up to a 10% chance for a damage spell to grant Clearcasting, making your next cast free; a large mana saver over a long fight.
- **Arcane Blast** (1 point) — a required pick: every cast gives you 12% of your Intellect as spell damage, stacking four times for 18 seconds from the first Blast, and it opens every cycle of the Blast loop described below in Rotation and priority.
- **Missile Barrage** — a proc off Arcane Blast, Arcane Wrath, Fireball, Frostbolt, or Frostfire Bolt that halves Arcane Missiles' channel time and drops its mana cost to zero. It now sits one row lower in the tree, and both builds take both ranks.
- **Presence of Mind** — turns your next sub-10-second cast instant, with no cast time of its own.
- **Arcane Mind** — up to 10% Intellect and a 100% bonus to Arcane critical strike damage at rank 5.
- **Arcane Power** — the capstone: 30% more spell damage for 15 seconds at the cost of 30% more mana per cast.

This build spends 31 points in Arcane, exactly the path to Arcane Power, and keeps every Arcane pick the sim measures as damage: **Arcane Impact** (3/3, 6% more Arcane critical strike chance) and **Arcane Instability** (3/3, 3% more spell damage and 3% more critical strike chance) stay maxed. The points that earlier carried the Arcane tree past 31 (both ranks of **Magic Absorption** and the second rank of **Wand Specialization**) came out because the sim measures none of them as damage for this rotation; **Missile Barrage** takes both its ranks, since a 40% proc on Arcane Blast is the engine of the spec (about +10% over one rank in our audit run).

Arcane Focus's five points had to stay in the top four rows of the tree, or Presence of Mind, Arcane Mind and Arcane Power fall out of reach. The leveling build puts them into the second rank of **Missile Barrage**, all three ranks of **Arcane Meditation** (3/3, in the fourth row; half your mana regeneration keeps running while you cast, which the Blast loop's mana cost leans on) and a fourth rank of **Improved Channeling**, so damage taken while leveling interrupts Arcane Blast less often. The raid build makes the same Missile Barrage and Arcane Meditation picks and spends the rest on both ranks of **Arcane Subtlety** (15 less resistance to all your spells on the target) instead, since a boss pushes back fewer of your casts than a pack of mobs does.

The other 20 points split 10 and 10. Fire keeps **Incineration** (3/5), **Accelerant** (the renamed Improved Fireball) at 2/5, **Flame Throwing** (2/2) and **Impact** (3/3); **Improved Fire Blast** (the talent in Wake of Fire's old slot) and **Ignite** stay at 0, as before. Those picks are held on purpose: the sim does not model Flame Throwing or Impact, so it cannot show what they grant, and the guide keeps them rather than trade away a talent nobody has measured. Frost now takes 10 points instead of 5: **Improved Frostbolt** (5/5) and **Ice Shards** (2/5) are new, and **Precision** stays in at its 3 ranks (the client trimmed it from 5). Improved Frostbolt lifts Frostbolt's damage in the two-Blast loop and is the single largest gain in this spend; Ice Shards adds to Frostbolt's critical strike damage, and the two points Precision no longer holds went there. This build measures about +1.9% over the previous spend in our level-60 search run, and every talent the engine cannot model (Improved Channeling, Arcane Resilience, Arcane Geometry, Flame Throwing, Impact) is still in it. Open the planner at [/planner?class=mage](/planner?class=mage) to build this out.

## Rotation and priority

Use Arcane Power the instant it's off cooldown — it has no cast time, so there's never a reason to hold it. Presence of Mind is the same: free to activate, so pop it whenever it's up and spend the instant cast on whichever spell benefits most in the moment. This site's simulator has Arcane Blast, its stacking buff and Missile Barrage wired up, including the part of Arcane Blast's tooltip that matters most: it costs 15% of your base mana, and every stack of its buff raises the next Blast's cost by 100%. Whenever Missile Barrage's proc is active, fire an empowered Arcane Missiles at top rank (half the channel time, no mana cost) to cash in the proc. Otherwise cast Arcane Blast until its buff carries three stacks. Past three, keep blasting while your mana is ahead of the fight (a larger share of your mana bar left than of the fight) and above 20%; once you fall behind, cast Frostbolt instead. The buff lasts 18 seconds from the first Blast and does not refresh, so Frostbolt fills until it runs out and the next Blast starts a fresh, cheap set. The mana-paced fourth stack gained about 10% at the raid band in our run and was even at the bare band, where mana runs out first. Frostfire Bolt can fill the Frostbolt slot (it also rolls Missile Barrage's chance) but measured lower, so Frostbolt stays. If the build hasn't taken Arcane Blast yet, plain Arcane Missiles fills instead, the same way pre-Arcane-Blast Era Arcane played.

## Stat priority

## Best in slot

## Enchants and consumables

## Races

## Professions

Tailoring and Enchanting is the standard caster pairing in Classic-era play: Tailoring's crafted spellcaster gear fills early slots, and Enchanting lets you apply your own weapon and gear enchants instead of paying for them.

## Leveling

Frost is the spec most players lean on for leveling, since Frostbolt's slow and Frost Nova's root make it safer to solo than Arcane's burst-focused kit. Arcane can still level fine, especially once Arcane Blast and Missile Barrage are actually castable.
