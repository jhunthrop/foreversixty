---
title: Arcane Mage in Forever
classSlug: mage
spec: arcane
role: dps
build: 'FS1:1.60.1.70009:mage:gnome:253225113100011531/032023/005:'
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

Arcane is Mage's burst-window spec: it trades a flatter damage curve for cooldowns that spike hard, built around Arcane Power and instant-cast tools that let you dump damage in a short window rather than sustain it evenly. In a raid it's a single-target caster with strong burst on cooldown and less baseline sustained output than Fire once a target's debuffs are stacked. The beta only reaches level 30, so nothing here about level 60 talents, raiding, or best-in-slot gear has actually been played — it is a projection from the demo trees and 1.12 knowledge, not tested content.

## Talents and builds

Blizzard confirmed the tree keeps its seven rows and 51 points, with a fourth one-point talent added at the 16-point mark alongside the existing ones at 11, 21, and 31. Verified against this build's talent data, the Arcane tree's key picks in roughly the order you'd take them are:

- **Arcane Focus** — up to 5% additional hit chance with Arcane spells, worth taking early since hit is scarce while leveling.
- **Arcane Concentration** — up to a 10% chance for a damage spell to grant Clearcasting, making your next cast free; a large mana saver over a long fight.
- **Arcane Blast** (1 point) — a required pick, not an optional one, and missing from an earlier draft of this guide's talent section: it opens every cycle of the two-Blast loop described below in Rotation and priority, so without it there's no spell to stack the buff with in the first place.
- **Missile Barrage** — a proc off Arcane Blast, Fireball, Frostbolt, or Frostfire Bolt that halves Arcane Missiles' channel time and drops its mana cost to zero.
- **Presence of Mind** — turns your next sub-10-second cast instant, with no cast time of its own.
- **Arcane Mind** — up to 10% Intellect and a 100% bonus to Arcane critical strike damage at rank 5.
- **Arcane Power** — the capstone: 30% more spell damage for 15 seconds at the cost of 30% more mana per cast.

This build spends 36 points in Arcane to reach Arcane Power at the bottom, two more than a bare 31-point path needs, because the sim's own level-60 search found two extra Arcane picks worth maxing on the way down: **Arcane Impact** (3/3, 6% more Arcane critical strike chance) and **Arcane Instability** (3/3, 3% more spell damage and 3% more critical strike chance), both measured as real damage gains rather than assumed.

That leaves 15 points to split between Fire and Frost, down from 20 in Fire alone. Fire drops to 10 points: **Wake of Fire** goes to 0/2, **Improved Fireball** falls from 5/5 to 2/5, and **Ignite** goes to 0/5 — none of those three measured any damage in the sim's model for this rotation, so giving them up cost nothing the sim could see. The other 5 points go into Frost for **Elemental Precision** (5/5, 5% more hit chance with Frost and Fire spells); this rotation never actually casts a Frost or Fire spell, so the sim can't register a damage gain from it either, but it's a real stat for a build that does mix in an off-school cast, and it was still the strongest use of those 5 points among the options the sim could measure. This build measures about +9.2% over the previous spend in our level-60 search run. Open the planner at [/planner?class=mage](/planner?class=mage) to build this out.

## Rotation and priority

Use Arcane Power the instant it's off cooldown — it has no cast time, so there's never a reason to hold it. Presence of Mind is the same: free to activate, so pop it whenever it's up and spend the instant cast on whichever spell benefits most in the moment. This site's simulator has Arcane Blast, its stacking buff and Missile Barrage wired up, including the part of Arcane Blast's tooltip that matters most: it costs 15% of your base mana, and every stack of its buff raises the next Blast's cost by 175%. Whenever Missile Barrage's proc is active, fire an empowered Arcane Missiles at top rank (half the channel time, no mana cost) to cash in the proc. Otherwise the loop is two Arcane Blasts, then Frostbolt. The second Blast is still affordable and brings the buff to two stacks; a third would cost several times a Frostbolt for the same damage, so Frostbolt takes the next global instead. It hits harder under the two-stack buff, it can proc Missile Barrage itself, and casting it ends the buff, which is exactly what you want: the next Blast starts a fresh, cheap pair. The search that settled this measured the two-Blast loop about 10% ahead of blasting without a stack limit, and well ahead of Arcane Missiles as the filler, which costs more mana per second of damage. If the build hasn't taken Arcane Blast yet, plain Arcane Missiles fills instead, the same way pre-Arcane-Blast Era Arcane played.

## Stat priority

The table above is this band's own simulation at level 60, re-run by the nightly pipeline every time the build or its gear data changes — these numbers are never hand-entered. A stat shown as "not separable from zero" is not a verdict that it is worthless; the sim's measured error on it is too wide, at this band's sample size, to tell its true value apart from zero. These numbers come from this site's own level-60 simulator, not from beta play, which only reaches level 30.

## Gear

Look for pieces that lead with spell power, then hit until capped, then crit, following the order above. Specific slot-by-slot picks can't be named honestly yet: the beta caps at level 30, so no one has gear-checked Arcane past the early game, and raid loot doesn't exist yet either — nothing raids until 9 December, and even then this client's item table is missing most of the classic raid loot tables it's meant to carry (Onyxia's Lair's own list is currently empty, and Barrow Deeps and Hyjal Summit aren't mapped to items at all). This section will fill in with real picks once raid loot is itemized.

## Enchants and consumables

Target spell power and hit on enchants and consumables, matching the stat priority above. **Enchant Weapon - Spell Power** (verified in this build's enchant data, +30 spell damage) is the standard weapon enchant for a caster. For flasks and elixirs, **Greater Arcane Elixir** and **Flask of Supreme Power** are both verified spell-damage consumables in this build's data; a weapon oil such as **Brilliant Wizard Oil** adds more of the same. Beyond naming those, keep it principle-level until raid-tier consumables are confirmed for Forever specifically.

## Races

Mage's playable races are Human, Orc, Undead, Gnome, and Troll on their respective factions, plus Skyborne — but only the Alliance-side Skyborne, since Mage is Alliance-only for that race. Orc is one of Blizzard's six confirmed new race and class pairs, giving Horde a Mage option it didn't have in 1.12.

For Alliance, **Gnome** is the strongest pick: Eureka! cuts the mana cost of your next three abilities by 50% and adds 10% more damage on a 2-minute cooldown, a direct damage-and-mana cooldown that lines up with Arcane's burst-window playstyle, and Expansive Mind adds to the mana pool a spec built around Arcane Power's extra mana cost benefits from. For Horde, **Orc** is the pick: Blood Fury grants 10% spell power for 15 seconds on a 2-minute cooldown, a real damage cooldown that stacks with Arcane Power's own burst window, and it's the new pairing Blizzard specifically added for this class. Skyborne is available to Alliance Mages as a third option; its Read Ley Line racial restores all health and mana over 15 seconds, useful out of combat, but it doesn't add spell damage the way Gnome's Eureka! does, so it's a pick for the race and story rather than raw throughput.

## Professions

Tailoring and Enchanting is the standard caster pairing in Classic-era play: Tailoring's crafted spellcaster gear fills early slots, and Enchanting lets you apply your own weapon and gear enchants instead of paying for them. This is general Classic-community convention rather than anything Forever-specific, so treat it as inferred until this site has profession data from the beta.

## Leveling

Frost is the spec most players lean on for leveling, since Frostbolt's slow and Frost Nova's root make it safer to solo than Arcane's burst-focused kit. Arcane can still level fine, especially once Arcane Blast and Missile Barrage are actually castable, but that's inferred from 1.12 knowledge, not tested — the beta only runs to level 30, so nothing about the leveling experience above that has been played in Forever's actual client.
