---
title: Holy Priest in Forever
classSlug: priest
spec: holy
role: healer
description: 'Holy Priest overview, talent priority, healing priority, stat weights, and race picks for Forever, with beta-versus-projection called out.'
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

Holy is the throughput healing tree for Priest, leaning on efficient direct and periodic heals rather than Discipline's shield-and-mitigation approach, with Prayer of Mending adding a new jumping heal to the kit that rewards keeping it moving through a group. It's the tree most Classic-era raids expect to carry the bulk of group-wide healing. The beta caps at level 30, so anything about how Holy performs healing a level-60 raid, or how it compares to Discipline at that scale, is a projection from the demo talent trees and 1.12 healing knowledge, not confirmed play.

## Talents and builds

In rough priority order: **Improved Healing** reduces the mana cost of Lesser Heal, Heal, Greater Heal, Penance, and Prayer of Mending, a broad efficiency gain across most of the direct-heal kit. **Holy Specialization** adds flat crit chance to Holy spells, which both increases throughput and, if Divine Aegis is reached through Discipline points, adds incidental shielding. **Spiritual Healing** increases the amount healed by all spells, a flat multiplier that scales with everything else in the build. **Spirit of Redemption**, now lasting 15 seconds rather than 10, gives a dying Holy Priest a longer window to land a last heal or cast before falling. **Litany of Light**, new to the tree, refunds mana whenever you switch which heal you're casting instead of repeating the same one, rewarding a varied toolkit over spamming a single spell. **Prayer of Mending**, the tree's new capstone, places a heal on a target that jumps to a new target each time it triggers, extending single-cast value across a whole group.

Point allocation runs deep into Holy to reach Prayer of Mending at the bottom of the tree, with the remainder split as a handful of points in Discipline for Power Word: Shield support. Open the planner at [/planner?class=priest](/planner?class=priest) to build this out.

## Rotation and priority

There's no curated rotation data for Holy in this build — it's a healing spec, not a DPS one this site's simulator models. Based on general Classic-era Holy practice plus what's changed for Forever: use Prayer of Mending on cooldown against a group taking spread damage, since it's a single cast that heals multiple targets over time as it jumps. Lean on Greater Heal for tank healing and Flash Heal for burst-raid-damage windows, alternating heal types where possible to trigger Litany of Light's mana return. Renew is worth maintaining on targets expected to take steady chip damage. This is this site's own inference from 1.12 healing conventions and the demo talent data, not a tested rotation.

## Stat priority

In simulator-derived priority order: **healing power** first, the direct multiplier on every heal cast. **Spell power** next, since Forever's itemization change lets bonus healing gear also carry a fraction of spell damage. **Spirit** feeds mana regeneration for a spec that casts frequently through a fight. **Mp5** provides flat regeneration independent of Spirit-based formulas. **Intellect** adds mana pool and a small crit chance. **Crit** last, both for direct heal size and, at high crit, for extra value from Holy Specialization.

## Gear

Prioritize healing power and spell power first, then spirit and mp5 for sustained casting through a full fight, then intellect and crit. Holy's efficiency talents make mana management more forgiving than in 1.12, but spirit remains the tree's core sustain stat. Specific pre-raid or raid-tier item picks can't be named with confidence yet: the beta caps at level 30 and this build's raid loot tables are still missing most items the community's sourcing expects, with nothing raiding in-game until the first tier opens on 9 December 2026.

## Enchants and consumables

Weapon and glove enchants should target healing power; **Enchant Weapon - Healing Power** and **Enchant Bracer - Healing Power** are both verified entries in this build's enchant data. For consumables, **Flask of Distilled Wisdom** is a verified option historically tied to mana regeneration for healers. A fuller consumable stack isn't something this site is naming until the class is played past level 30.

## Races

Holy can be played by Human, Dwarf, Night Elf, and Gnome on the Alliance, and Undead and Troll on the Horde. Priest is not available to the Skyborne on either faction. For Alliance, **Gnome** is the strongest pick for a throughput-focused Holy build: Contingency Plan is a proactive Holy ward that absorbs and heals an ally automatically when they drop below 35% health, functioning as a passive extra heal that costs no cast time — a direct fit for a tree built around efficient, repeated healing. Human's Divine Grace is a close second through its instant heal on demand. For Horde, **Troll** is the stronger pick over Undead for Holy specifically: Regeneration keeps 10% of health regeneration running in combat, a passive survivability boost that complements a spec focused on keeping others alive rather than sustaining its own mana pool the way Undead's Dark Sacrifice does.

## Professions

Community convention for healing Priests favors **Tailoring** for caster cloth stats alongside **Enchanting** for weapon and gear enchants, since neither profession competes with anything Holy-specific. This is general Classic-era community practice, not confirmed for Forever.

## Leveling

Shadow is the generally recommended leveling spec for Priest, not Holy — Holy's group-throughput kit doesn't translate to fast solo leveling the way Shadow's self-sufficient damage and healing does. The beta caps at level 30, so this is carried over from 1.12 leveling knowledge rather than tested in Forever.
