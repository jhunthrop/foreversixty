---
title: Destruction Warlock in Forever
classSlug: warlock
spec: destruction
role: dps
description: 'Talents, rotation, stats, and gear for Destruction Warlock in Forever, and what is confirmed versus projected from the beta.'
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
  - label: 'Forever Sixty rotation data, data/curated/apl/warlock-destruction.json'
    url: https://foreversixty.gg/sources
    kind: site
  - label: 'Forever Sixty raid loot data, data/curated/loot/forever-raid-phases.json'
    url: https://foreversixty.gg/sources
    kind: site
---

## Overview

Destruction is Warlock's direct-damage spec: it keeps Immolate ticking on the target so Conflagrate has something to consume and Incinerate has something to boost, and it closes out low-health targets with Shadowburn. Compared to Affliction's set-and-forget DoTs, Destruction spends more globals on hard-hitting single casts and instants, trading some sustain for burst. The beta only reaches level 30, so nothing here about level 60 talents, raiding, or best-in-slot gear has actually been played — it is a projection from the demo trees and 1.12 knowledge, not tested content.

## Talents and builds

Blizzard confirmed the tree keeps its seven rows and 51 points, with a fourth one-point talent added at the 16-point mark alongside the existing ones at 11, 21, and 31. Verified against this build's talent data, Destruction's key picks in roughly the order you'd take them are:

- **Bane** — up to 0.5 seconds off Shadow Bolt, Immolate, and Incinerate cast time, and up to 2 seconds off Soul Fire, at rank 5.
- **Ruin** — up to a 100% critical-strike-damage bonus at rank 5 on all Destruction spells, the tree's biggest single multiplier.
- **Conflagrate** — an instant that ignites a target already afflicted by Immolate, dealing Fire damage and consuming the Immolate effect.
- **Shadowburn** — an instant Shadow nuke that refunds a Soul Shard if the target dies within 8 seconds of being hit.
- **Shadow and Flame** — up to a 10% damage buff at rank 5 from landing Conflagrate or Shadowburn, plus a chance for Conflagrate not to consume Immolate.
- **Incinerate** — the capstone: extra Fire damage that gains a further 25% if the target is afflicted by Immolate, Destruction's signature filler.

A typical Destruction build spends roughly 31 points in this tree to reach Incinerate at the bottom, with the remaining 20 points usually going into Affliction for Corruption and Suppression — a common 1.12 hybrid pattern that likely still applies, though it isn't confirmed for Forever specifically. Open the planner at [/planner?class=warlock](/planner?class=warlock) to build this out.

## Rotation and priority

Pop a mana potion at the start of the fight if your mana is already below its usual threshold. Immolate is the DoT this whole rotation is built around: refresh it the instant it falls off, since both Conflagrate and Incinerate depend on it being active. Conflagrate is only worth casting while Immolate is up, since it consumes the DoT for an instant burst of damage — cast it as soon as it's off cooldown with Immolate running. Once a target drops below 20% health, Shadowburn becomes the priority finisher. Incinerate is meant to be the default filler for every other global, dealing bonus damage while Immolate is active; this site's simulator doesn't have it wired up yet, so today Shadow Bolt fills that role instead — without it, once Immolate is up and Conflagrate is on cooldown, there would be nothing else to cast.

## Stat priority

In this site's own simulator weighting, in priority order:

1. **Spell power** — the reference stat; Immolate, Shadow Bolt, and every Destruction nuke scale off it directly.
2. **Intellect** — a larger mana pool for a rotation with several expensive direct-damage casts.
3. **Critical strike** — Ruin turns every crit into far more damage, and Shadowburn's execute value climbs with it.
4. **Hit** — needed to stop missing casts against raid-level bosses.
5. **Spell haste** — shortens Shadow Bolt and Immolate cast time, more casts per minute.
6. **Spell penetration** — only matters against targets with meaningful shadow or fire resistance.
7. **Shadow damage and Fire damage (school power)** — the narrowest stats, appearing on very few items, both relevant to Destruction's mixed-school kit.

## Gear

Look for pieces that lead with spell power, then hit until capped, then crit, following the order above. Specific slot-by-slot picks can't be named honestly yet: the beta caps at level 30, so no one has gear-checked Destruction past the early game, and raid loot doesn't exist yet either — nothing raids until 9 December, and even then this client's item table is missing most of the classic raid loot tables it's meant to carry (Onyxia's Lair's own list is currently empty, and Barrow Deeps and Hyjal Summit aren't mapped to items at all). This section will fill in with real picks once raid loot is itemized.

## Enchants and consumables

Target spell power and hit on enchants and consumables, matching the stat priority above. **Enchant Weapon - Spell Power** (verified in this build's enchant data, +30 spell damage) is the standard weapon enchant for a caster. Destruction splits its damage between Shadow (Shadow Bolt, Shadowburn) and Fire (Immolate, Conflagrate, Incinerate), so **Elixir of Fire Power**, verified in this build's data, is a reasonable school-specific pick given Immolate's constant uptime, while **Flask of Supreme Power** covers the general spell-damage slot on a raid night. Beyond naming those, keep it principle-level until raid-tier consumables are confirmed for Forever specifically.

## Races

Warlock's playable races are Human, Orc, Undead, Gnome, and Troll on their respective factions; the class is not open to Skyborne at all. Troll is one of Blizzard's six confirmed new race and class pairs, giving the Horde a Warlock option it didn't have in 1.12.

For Alliance, **Gnome** is the strongest pick: Eureka!'s 50% mana-cost reduction and 10% damage bonus on the next three casts is a direct fit for a spec that leans on expensive direct-damage spells like Incinerate and Soul Fire, and Expansive Mind's mana bonus keeps that rotation running longer. For Horde, **Troll** is the pick as the new pairing Blizzard specifically added for this class: Berserking's 10% casting and attack speed on a 3-minute cooldown shortens the cast time on every direct-damage spell in the rotation, a straightforward fit for Destruction's burst-oriented kit.

## Professions

Tailoring and Enchanting is the standard caster pairing in Classic-era play: Tailoring's crafted spellcaster gear fills early slots, and Enchanting lets you apply your own weapon and gear enchants instead of paying for them. This is general Classic-community convention rather than anything Forever-specific, so treat it as inferred until this site has profession data from the beta.

## Leveling

Affliction is generally the stronger leveling choice for Warlock; Destruction can level well too, since Shadowburn and Conflagrate give it solid burst against single targets, but it lacks Affliction's DoT-and-move safety against multiple enemies. That assessment is inferred from 1.12 knowledge, not tested — the beta only runs to level 30, so nothing about the leveling experience above that has been played in Forever's actual client.
