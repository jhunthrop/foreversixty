---
title: Discipline Priest in Forever
classSlug: priest
spec: discipline
role: healer
description: 'Discipline Priest overview, talent priority, healing priority, stat weights, and race picks for Forever, with beta-versus-projection called out.'
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

Discipline is the shield-and-mitigation healing tree, leaning on Power Word: Shield uptime and cost-reduction talents rather than Holy's raw throughput, with Penance adding a new channeled heal-or-damage option to the kit. It tends to shine on tank healing and damage prevention rather than raw group-wide output. The beta caps at level 30, so anything about how Discipline performs healing a level-60 raid boss, or how it stacks up against Holy in that context, is a projection from the demo talent trees and 1.12 healing knowledge, not confirmed play.

## Talents and builds

In rough priority order: **Improved Power Word: Shield** raises the amount absorbed, a direct multiplier on the tree's signature spell. **Meditation** now allows 17% of mana regeneration to continue while casting, up sharply from Classic's 5%, which is the single biggest quality-of-life change in the tree and makes sustained casting far less punishing on mana. **Inner Focus** gives one free, empowered cast on demand, useful for an emergency Power Word: Shield or Penance. **Soul Warding** reduces Power Word: Shield's cooldown and mana cost, compounding with Improved Power Word: Shield. **Penance**, new to the tree, channels either a burst of healing into a friendly target or a burst of holy damage into an enemy, giving Discipline both a strong single-target heal and a usable damage option to fill downtime. **Divine Aegis**, also new, turns critical heals into a protective absorb shield, adding incidental mitigation on top of every crit.

Point allocation runs deep into Discipline to reach Penance and Divine Aegis, with the remainder split as a handful of points in Holy for survivability and utility talents. Open the planner at [/planner?class=priest](/planner?class=priest) to build this out.

## Rotation and priority

There's no curated rotation data for Discipline in this build — it's a healing spec, not a DPS one this site's simulator models. Based on general Classic-era Discipline practice plus what's changed for Forever: keep Power Word: Shield on cooldown against a tank or focus target ahead of expected damage, since Soul Warding shortens its cooldown enough to lean on it more than in vanilla 1.12. Use Penance as the primary direct heal once talented, since it's both mana-efficient and can double as filler damage during downtime. Fall back to Flash Heal or Greater Heal for burst healing outside Power Word: Shield's cooldown, and use Inner Focus to cover an emergency cast when mana is tight. This is this site's own inference from 1.12 healing conventions and the demo talent data, not a tested rotation.

## Stat priority

In simulator-derived priority order: **healing power** first, the direct multiplier on every heal cast. **Spell power** next, since Forever's itemization change lets bonus healing gear also carry a fraction of spell damage. **Spirit** feeds mana regeneration, which matters more for a spec casting Power Word: Shield and Penance repeatedly through a fight. **Mp5** provides flat regeneration independent of Spirit-based formulas. **Intellect** adds mana pool and a small crit chance. **Crit** last, both for direct heal size and for triggering Divine Aegis's absorb shield if talented.

## Gear

Prioritize healing power and spell power first, then spirit and mp5 for sustain, then intellect and crit. Because Meditation now lets a much larger share of mana regeneration continue while casting, spirit is somewhat less punishing to undervalue than in 1.12, but it's still the tree's main sustain stat. Specific pre-raid or raid-tier item picks can't be named with confidence yet: the beta caps at level 30 and this build's raid loot tables are still missing most items the community's sourcing expects, with nothing raiding in-game until the first tier opens on 9 December 2026.

## Enchants and consumables

Weapon and glove enchants should target healing power; **Enchant Weapon - Healing Power** and **Enchant Bracer - Healing Power** are both verified entries in this build's enchant data. For consumables, **Flask of Distilled Wisdom** is a verified option historically tied to mana regeneration for healers. A fuller consumable stack isn't something this site is naming until the class is played past level 30.

## Races

Discipline can be played by Human, Dwarf, Night Elf, and Gnome on the Alliance, and Undead and Troll on the Horde. Priest is not available to the Skyborne on either faction. For Alliance, **Human** is the strongest pick: the reworked Divine Grace snaps out an instant heal on any party member whose health has fallen under half, self-targeting excluded, giving Discipline a free emergency cooldown that stacks with Power Word: Shield rather than competing with it. Gnome, the new Alliance pairing, is also a reasonable pick through Contingency Plan, a defensive ward that triggers its own absorb and heal once a party member's health falls under roughly a third — a similar proactive-mitigation idea to Divine Grace, just automatic rather than cast. For Horde, **Undead** is the strongest pick: Dark Sacrifice converts health to mana, giving a Discipline healer a way to keep casting through a fight that's run the mana pool dry, which matters more for Discipline than for a burst-heavy Holy build.

## Professions

Community convention for healing Priests favors **Tailoring** for caster cloth stats alongside **Enchanting** for weapon and gear enchants, since neither profession competes with anything Discipline-specific. This is general Classic-era community practice, not confirmed for Forever.

## Leveling

Shadow is the generally recommended leveling spec for Priest, not Discipline — Discipline's shield-and-mitigation kit is built around keeping other people alive, which doesn't translate to fast solo leveling the way Shadow's self-sufficient damage and healing does. The beta caps at level 30, so this is carried over from 1.12 leveling knowledge rather than tested in Forever.
