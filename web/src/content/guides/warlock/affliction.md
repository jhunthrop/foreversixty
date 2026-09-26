---
title: Affliction Warlock in Forever
classSlug: warlock
spec: affliction
role: dps
build: 'FS1:1.60.1.69893:warlock:gnome:25552000030201051/235523/0:'
recommendedRaces: [gnome, troll]
statPriority: [Spell power, Intellect, Critical strike, Hit, Spell haste, Spell penetration, 'Shadow damage']
description: 'Talents, rotation, stats, and gear for Affliction Warlock in Forever, and what is confirmed versus projected from the beta.'
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
  - label: 'Forever Sixty rotation data, data/curated/apl/warlock-affliction.json'
    url: https://foreversixty.gg/sources
    kind: site
  - label: 'Forever Sixty raid loot data, data/curated/loot/forever-raid-phases.json'
    url: https://foreversixty.gg/sources
    kind: site
---

## Overview

Affliction is Warlock's damage-over-time spec: it layers multiple DoTs onto a target and lets them run, spending Life Tap to keep casting once mana runs low rather than worrying about health, since the drain spells and DoTs keep the Warlock topped up. In a raid it's a low-maintenance, high-uptime caster that keeps producing damage even during movement, once its DoTs are applied. The beta only reaches level 30, so nothing here about level 60 talents, raiding, or best-in-slot gear has actually been played — it is a projection from the demo trees and 1.12 knowledge, not tested content.

## Talents and builds

Blizzard confirmed the tree keeps its seven rows and 51 points, with a fourth one-point talent added at the 16-point mark alongside the existing ones at 11, 21, and 31; separately, Curses and Banes now run as independent categories so a Bane and a Curse can be active on the same target together. Verified against this build's talent data, Affliction's key picks in roughly the order you'd take them are:

- **Suppression** — up to 5% more hit chance and 20% less threat at rank 5, both scarce and valuable for a spec that casts constantly.
- **Improved Corruption** — cuts Corruption's cast time by up to 2 seconds and adds up to 10% more damage at rank 5.
- **Pandemic** — up to a 100% critical-strike-damage bonus at rank 3 on Corruption, both Banes, and the drain spells, the tree's biggest damage multiplier.
- **Nightfall** — up to a 4% chance per DoT tick at rank 2 to make your next Shadow Bolt instant, a real mana and time saver.
- **Shadow Mastery** — up to 5% more Shadow damage and life drained at rank 5.
- **Wrack** — the capstone: a DoT that also boosts your other Shadow DoTs on the same target by 10% while it's active.

A typical Affliction build spends roughly 31 points in this tree to reach Wrack at the bottom, with the remaining 20 points usually going into Demonology for pet survivability and mana talents — a common 1.12 hybrid pattern that likely still applies, though it isn't confirmed for Forever specifically. Open the planner at [/planner?class=warlock](/planner?class=warlock) to build this out.

## Rotation and priority

Pop a mana potion or gem at the start of the fight if your mana is already below its usual threshold, or save it for a closing burn if the fight is nearly over. From there, apply Bane of Agony first, since its damage ramps up over its duration and it benefits most from ticking the longest; refresh it the moment it falls off rather than letting the ramp reset early. Corruption goes up alongside it and stays up throughout the fight — it's the cheapest DoT per point of damage, so upkeep on it should never lapse. Once both DoTs are running, spend every remaining global on Shadow Bolt. Wrack, Affliction's own capstone DoT, belongs in that same DoT-upkeep rotation once it's actually castable; this site's simulator doesn't have it wired up yet, so today Shadow Bolt fills the globals it would otherwise use. Life Tap whenever mana drops below about 10%, since three running DoTs plus Drain Life give enough of a health buffer to spend some of it on mana.

## Stat priority

In this site's own simulator weighting, in priority order:

1. **Spell power** — the reference stat; every DoT tick and Shadow Bolt scales off it directly.
2. **Intellect** — a larger mana pool for a rotation that casts through the whole fight.
3. **Critical strike** — Pandemic turns every DoT and drain crit into far more damage.
4. **Hit** — needed to stop missing casts and letting DoTs fail to apply against raid-level bosses; partly covered by Suppression.
5. **Spell haste** — shortens Shadow Bolt's cast time, the main filler between DoT refreshes.
6. **Spell penetration** — only matters against targets with meaningful shadow resistance.
7. **Shadow damage (school power)** — the narrowest stat, appearing on very few items, but it stacks directly with Shadow Mastery's percentage bonus.

## Gear

Look for pieces that lead with spell power, then hit until capped, then crit, following the order above. Specific slot-by-slot picks can't be named honestly yet: the beta caps at level 30, so no one has gear-checked Affliction past the early game, and raid loot doesn't exist yet either — nothing raids until 9 December, and even then this client's item table is missing most of the classic raid loot tables it's meant to carry (Onyxia's Lair's own list is currently empty, and Barrow Deeps and Hyjal Summit aren't mapped to items at all). This section will fill in with real picks once raid loot is itemized.

## Enchants and consumables

Target spell power and hit on enchants and consumables, matching the stat priority above. **Enchant Weapon - Spell Power** (verified in this build's enchant data, +30 spell damage) is the standard weapon enchant for a caster. **Elixir of Shadow Power** is a verified school-specific damage elixir in this build's data and is the natural pick for Affliction over a generic spell-power elixir when the two conflict; **Flask of Supreme Power** covers the general spell-damage slot on a raid night. Beyond naming those, keep it principle-level until raid-tier consumables are confirmed for Forever specifically.

## Races

Warlock's playable races are Human, Orc, Undead, Gnome, and Troll on their respective factions; the class is not open to Skyborne at all. Troll is one of Blizzard's six confirmed new race and class pairs, giving the Horde a Warlock option it didn't have in 1.12.

For Alliance, **Gnome** is the strongest pick: Eureka! cuts the mana cost of your next three abilities by 50% and adds 10% more damage on a 2-minute cooldown, which pairs well with re-applying DoTs after a burst window, and Expansive Mind's mana bonus supports a rotation that keeps casting all fight. For Horde, **Troll** is the pick as the new pairing Blizzard specifically added for this class: Berserking grants 10% casting and attack speed on a 3-minute cooldown, which shortens Shadow Bolt casts between DoT refreshes, and Regeneration keeps some health regeneration running in combat, a fit for a spec that spends health through Life Tap.

## Professions

Tailoring and Enchanting is the standard caster pairing in Classic-era play: Tailoring's crafted spellcaster gear fills early slots, and Enchanting lets you apply your own weapon and gear enchants instead of paying for them. This is general Classic-community convention rather than anything Forever-specific, so treat it as inferred until this site has profession data from the beta.

## Leveling

Affliction is generally considered the strong leveling choice for Warlock, since its DoTs keep working while you move or fight multiple targets, and Drain Life covers the healing a solo caster otherwise lacks. That assessment is inferred from 1.12 knowledge, not tested — the beta only runs to level 30, so nothing about the leveling experience above that has been played in Forever's actual client.
