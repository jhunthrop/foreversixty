---
title: Discipline Priest in Forever
classSlug: priest
spec: discipline
role: healer
build: 'FS1:1.60.1.70009:priest:human:02500303130510152/035050030301/0:'
recommendedRaces: [human, undead]
statPriority: [Healing power, Intellect, Spirit, MP5, Critical strike]
description: 'Discipline Priest overview, talent priority, healing priority, stat weights, and race picks for Forever, with beta-versus-projection called out.'
updated: 2026-10-07
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

Discipline is the shield-and-mitigation healing tree, leaning on Power Word: Shield uptime and cost-reduction talents rather than Holy's raw throughput, with Penance adding a new channeled heal-or-damage option to the kit. It tends to shine on tank healing and damage prevention rather than raw group-wide output. The beta caps at level 30, so anything about how Discipline performs healing a level-60 raid boss, or how it stacks up against Holy in that context, is a projection from the demo talent trees and 1.12 healing knowledge, not confirmed play.

## Talents and builds

The build spends its deep points in Discipline and fills Holy with the points that make the heals cast faster and cheaper. Chosen by simming candidate builds on the same gear against the same fight. Every figure behind this ranking comes from one stated incoming-damage profile (the scaled Phase 1 tank hits and raid pulses described on the [BiS page](/bis)): it ranks gear and builds for this spec, and says nothing about which healer class is stronger.

- **Twin Disciplines, Improved Power Word: Shield and Mental Agility** are the early rows: instant spells heal more, the shield absorbs more, and instants cost less.
- **Meditation** keeps part of your regeneration running while you cast and **Mental Strength** lifts Intellect; **Inner Focus** is a free, extra-critical heal on demand.
- **Penance**, **Renewed Hope** and **Divine Aegis** are the deep payoff: a channelled heal, a bigger critical chance on a target that carries Weakened Soul, and a shield from your critical heals.
- In Holy, **Holy Specialization, Divine Fury and Improved Healing** give critical chance, faster Heal and Greater Heal, and cheaper direct heals; **Inspiration**, **Improved Renew** and **Binding Heal** fill the remaining points.

Open the planner at [/planner?class=priest](/planner?class=priest) to adjust it.

## Rotation and priority

The curated rotation leads with **Penance** whenever it is ready, and keeps **Power Word: Shield** on a tank who is taking damage and is not under Weakened Soul. **Renew** and **Prayer of Mending** go on the tank; **Flash Heal** is the emergency button for a nearly dead tank or member; **Prayer of Healing** answers several party members being hurt together; **Heal** is the efficient filler.

The shield, the filler and Prayer of Healing are paced against your mana the way Holy's are: they fire at their normal thresholds only while your mana is keeping up with the share of the fight that is left. Discipline lasts longer on its mana than Holy does under the same fight, which is why the shield is worth pacing rather than dropping.

## Stat priority

In simulator-derived order, per point of stat: **MP5** first, then **Spirit**, **healing power** and **Intellect** close together, with **critical strike** last and measured with less certainty. The fight is mana-bound, so regeneration and the pool matter more than a point of throughput. Items trade several points of one for few of another, so read the list as "which stat is cheap to give up".

Spell power adds nothing to a heal; it is the damage side of a healing item. Every figure behind this ranking comes from one stated incoming-damage profile (the scaled Phase 1 tank hits and raid pulses described on the [BiS page](/bis)): it ranks gear and builds for this spec, and says nothing about which healer class is stronger.

## Gear

Follow the stat priority above: regeneration first, then healing power, Spirit and Intellect close together; spell power on a healing item adds nothing to a heal. Because Meditation now lets a much larger share of mana regeneration continue while casting, spirit is somewhat less punishing to undervalue than in 1.12, but it's still the tree's main sustain stat. Specific pre-raid or raid-tier item picks can't be named with confidence yet: the beta caps at level 30 and this build's raid loot tables are still missing most items the community's sourcing expects, with nothing raiding in-game until the first tier opens on 9 December 2026.

## Enchants and consumables

Weapon and glove enchants should target healing power; **Enchant Weapon - Healing Power** and **Enchant Bracer - Healing Power** are both verified entries in this build's enchant data. For consumables, **Flask of Distilled Wisdom** is a verified option historically tied to mana regeneration for healers. A fuller consumable stack isn't something this site is naming until the class is played past level 30.

## Races

Discipline can be played by Human, Dwarf, Night Elf, and Gnome on the Alliance, and Undead and Troll on the Horde. Priest is not available to the Skyborne on either faction. For Alliance, **Human** is the strongest pick: the reworked Divine Grace snaps out an instant heal on any party member whose health has fallen under half, self-targeting excluded, giving Discipline a free emergency cooldown that stacks with Power Word: Shield rather than competing with it. Gnome, the new Alliance pairing, is also a reasonable pick through Contingency Plan, a defensive ward that triggers its own absorb and heal once a party member's health falls under roughly a third — a similar proactive-mitigation idea to Divine Grace, just automatic rather than cast. For Horde, **Undead** is the strongest pick: Dark Sacrifice converts health to mana, giving a Discipline healer a way to keep casting through a fight that's run the mana pool dry, which matters more for Discipline than for a burst-heavy Holy build.

## Professions

Community convention for healing Priests favors **Tailoring** for caster cloth stats alongside **Enchanting** for weapon and gear enchants, since neither profession competes with anything Discipline-specific. This is general Classic-era community practice, not confirmed for Forever.

## Leveling

Shadow is the generally recommended leveling spec for Priest, not Discipline — Discipline's shield-and-mitigation kit is built around keeping other people alive, which doesn't translate to fast solo leveling the way Shadow's self-sufficient damage and healing does. The beta caps at level 30, so this is carried over from 1.12 leveling knowledge rather than tested in Forever.
