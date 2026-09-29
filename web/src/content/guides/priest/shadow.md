---
title: Shadow Priest in Forever
classSlug: priest
spec: shadow
role: dps
build: 'FS1:1.60.1.69893:priest:gnome:5241110013/0/443110501201300251:'
recommendedRaces: [gnome, undead]
statPriority: [Spell power, Intellect, Critical strike, Hit, Spell haste, Spell penetration, 'Shadow power']
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

Shadow is Priest's damage-over-time and mind-magic DPS spec, built around maintaining Shadow Word: Pain and layering Mind Blast and Mind Flay on top of it, with a small amount of self-sufficiency from Vampiric Embrace-style healing. It's also the class's strongest solo and leveling spec, since its damage kit doubles as sustain. The beta caps at level 30, so how Shadow's damage output compares to other casters at level 60, or performs in a raid, is a projection from the demo talent trees and 1.12 knowledge, not confirmed play.

## Talents and builds

In rough priority order: **Mind Flay** (1 point) is a required pick, not an optional one — Forever moved it out of the baseline trainer-taught kit and into this talent, so without it there is no filler spell to fill the GCDs between cooldowns at all. **Improved Mind Flay**, ranked further down the tree, is its scaling once the 1-point prerequisite is in. **Shadow Weaving** stacks a Shadow damage buff from spell crits, a compounding multiplier for a spec casting Shadow spells constantly. **Improved Mind Blast** cuts Mind Blast's cooldown, letting the hardest-hitting single spell in the rotation come up more often. **Darkness** adds a flat percentage to all Shadow damage, a broad multiplier late in the tree. **Early Demise**, new to the tree, raises Shadow Word: Death's critical strike chance against targets below 20% health, turning it into a dedicated execute. **Shadowform**, the tree's capstone, increases Shadow damage by 10% and reduces the mana cost of Shadow spells, at the cost of being unable to cast non-Shadow spells while it's active.

Point allocation runs deep into Shadow to reach Shadowform at the bottom of the tree, 33 points, with the remaining 18 in Discipline: **Inner Focus** (1 point) is a required pick alongside the Meditation talent's mana regeneration while casting — it's the free, empowered Devouring Plague cast this build's Rotation section calls out below, so it has to be taken for that line to be legal, not just Meditation. Open the planner at [/planner?class=priest](/planner?class=priest) to build this out.

## Rotation and priority

The core loop this site's simulator plays: keep Shadow Word: Pain up on the target unless the fight is close to ending, since refreshing it late would waste most of the remaining damage-over-time value. Below 20% target health, Shadow Word: Death is the priority — Forever's Early Demise talent turns it into a dedicated execute, at its full top rank — falling back to Mind Blast on cooldown and Mind Flay otherwise on the rare GCD where Shadow Word: Death is still on its own 15-second cooldown inside that window. Outside the execute window, Inner Focus into Devouring Plague is used as a free, empowered cast of the spec's biggest nuke, at its full top rank, followed by Mind Blast on cooldown as the hardest-hitting repeatable single-target spell. Mind Flay fills the remaining time between cooldowns — this needs the Mind Flay talent taken, since Forever gates it behind that pick rather than granting it as a baseline trainer spell (see Talents and builds below for when to take it while leveling).

## Stat priority

In simulator-derived priority order: **spell power** first, the direct multiplier on every Shadow spell. **Intellect** next for mana pool and a small crit contribution. **Crit** feeds both raw damage and Shadow Weaving's stacking buff. **Hit** keeps every cast in the rotation landing, which matters more for Shadow than for a healer since a missed Mind Blast is a missed cooldown. **Spell haste** speeds up the whole rotation. **Spell penetration** helps against magic-resistant targets. **Shadow power**, gear that specifically boosts Shadow damage rather than all spell schools, sits last as the most specialized and rarest stat to find.

## Gear

Prioritize spell power and intellect first, then crit and hit, then spell haste, spell penetration, and Shadow-specific power. Because bonus healing gear in Forever now also carries a fraction of bonus damage, some cloth pieces itemized for healers may carry incidental value for Shadow that they wouldn't have in 1.12 — worth checking case by case rather than assuming. Specific pre-raid or raid-tier item picks can't be named with confidence yet: the beta caps at level 30 and this build's raid loot tables are still missing most items the community's sourcing expects, with nothing raiding in-game until the first tier opens on 9 December 2026.

## Enchants and consumables

Weapon enchants should target spell power; **Enchant Weapon - Spell Power** is a verified entry in this build's enchant data. For consumables, **Elixir of Shadow Power** and **Greater Arcane Elixir** are both verified options tied historically to Shadow and spell damage respectively. A fuller consumable stack isn't something this site is naming until the class is played past level 30.

## Races

Shadow can be played by Human, Dwarf, Night Elf, and Gnome on the Alliance, and Undead and Troll on the Horde. Priest is not available to the Skyborne on either faction. For Alliance, **Gnome** is a strong pick for a caster DPS spec: Eureka! makes the next three abilities cost 50% less mana and deal 10% more damage on a 2-minute cooldown, functioning as a burst-and-sustain cooldown that lines up naturally with Mind Blast and Devouring Plague casts. For Horde, **Undead** is the strongest pick: Dark Sacrifice's health-to-mana conversion gives Shadow a way to keep casting through extended fights or pulls without downtime, and Touch of the Grave adds passive chance-on-hit damage that stacks with the spec's DoT-and-nuke kit.

## Professions

Community convention for Shadow Priests favors **Tailoring** for caster cloth stats alongside **Enchanting** for weapon and gear enchants, the same combination general caster practice recommends across classes. This is Classic-era community convention, not confirmed for Forever.

## Leveling

Shadow is the strongest leveling spec for Priest by a wide margin: its damage-over-time kit provides both offense and self-healing that Discipline and Holy, built around supporting other players, don't have solo. The beta caps at level 30, so this is carried over from 1.12 leveling knowledge rather than tested past that point in Forever.

Take **Mind Flay**'s single point the moment it's available at level 20, ahead of the level-60 build's own point order above. That order maxes Shadow Focus, Blackout and Spirit Tap first for a level-60 raider's hit and threat needs, but a leveling character walking those same points in sequence doesn't reach Mind Flay's row until several levels later — and until then there is no ranged filler at all beyond the wand, since Forever moved Mind Flay out of the baseline kit. Grab the 1-point prerequisite first, then fill Shadow Focus/Blackout/Spirit Tap and the rest of the tree behind it.
